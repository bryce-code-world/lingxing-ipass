package dsco_lingxing

import (
	"context"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

// PushToLingXing 将 DSCO 订单推送到领星创建订单（CreateOrdersV2）。
//
// 流程（一期口径）：
//  1. 从 dsco_order_sync 中筛选 status=1 的订单（待推单）。
//  2. 解析 DSCO 原始订单 payload，校验关键字段（poNumber/地址/行项目）。
//  3. 映射（匹配不到一律跳过）：
//     - mapping.shop：dscoRetailerId -> 领星 store_id
//     - mapping.warehouse：requestedWarehouseCode -> 领星 WID
//     - mapping.shipment：<shipWarehouseCode(空则用 requestedWarehouseCode)>-<shippingServiceLevelCode> -> 领星 logistics_type_id
//  4. 幂等：创建前先查领星是否已存在该 poNumber；存在则直接推进状态到 2，不重复创建。
//  5. 调用 CreateOrdersV2（最小必要字段），成功后推进状态到 2。
func (d *Domain) PushToLingXing(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "push_to_lingxing")...)

	var (
		total   int
		okCount int
		skip    int
		fail    int
	)
	defer func() {
		fields := append(base,
			"task", "push_to_lingxing",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okCount,
			"skip", skip,
			"fail", fail,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	// 1) 取待推单（status=1）
	items, err := d.orderStore.FindByStatus(taskCtx, 1, ctx.Size)
	if err != nil {
		retErr = err
		return retErr
	}
	if len(items) == 0 {
		return nil
	}
	total = len(items)

	// 2) 初始化领星客户端
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}

	// 4) 幂等检查（批量）：本批次所有 poNumber 一次性查询，避免逐单调用导致触发 API 限制。
	//
	// 说明：
	// - SDK 的 GetOrderDetailV2 底层也是调用 /pb/mp/order/v2/list，但一次只查 1 个。
	// - 这里直接按批次查 list，并按需降级到逐单查询，兼顾效率与稳定性。
	poNumbers := make([]string, 0, len(items))
	for _, it := range items {
		poNumbers = append(poNumbers, it.PONumber)
	}
	poNumbers = uniqueNonEmptyStrings(poNumbers)
	existing := make(map[string]lingxing.OrderDetailV2, len(poNumbers))
	includeDelete := true
	const maxBatch = 50
	for _, chunk := range chunkStrings(poNumbers, maxBatch) {
		out, err := lx.Order.ListOrdersV2(taskCtx, lingxing.OrderListV2Request{
			Offset:           0,
			Length:           len(chunk),
			PlatformOrderNos: chunk,
			IncludeDelete:    &includeDelete,
		})
		if err != nil {
			// 降级：逐单查（尽量不因为一次 list 失败影响整个批次）
			logger.Warn(taskCtx, "lingxing list orders failed, fallback to per-order check", "err", err)
			for _, po := range chunk {
				detail, derr := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
				if derr == nil {
					existing[po] = detail
				}
			}
			continue
		}
		for _, detail := range out.List {
			if po := poNumberFromLingXingOrderDetail(detail); po != "" {
				existing[po] = detail
			}
		}
	}
	logger.Info(taskCtx, "lingxing existing orders checked",
		append(base,
			"task", "push_to_lingxing",
			"po_count", len(poNumbers),
			"existing", len(existing),
		)...,
	)

	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		// 4.1) 若领星已存在该 poNumber，则不再创建，直接推进状态到 2。
		if detail, ok := existing[po]; ok {
			if err := d.orderStore.UpdateStatusAndFields(taskCtx, po, 2, "", ""); err != nil {
				fail++
				logger.Warn(taskCtx, "order done",
					append(base,
						"task", "push_to_lingxing",
						"po_number", po,
						"result", "fail",
						"reason", "update_local_status_failed_for_existing_order",
						"lingxing_status", int(detail.Status),
						"lingxing_raw", integration.JSONForLog(detail),
						"err", err,
					)...,
				)
				continue
			}
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_order_already_exists",
					"lingxing_status", int(detail.Status),
					"lingxing_raw", integration.JSONForLog(detail),
					"new_status", 2,
				)...,
			)
			continue
		}

		// 2.1) 解析 DSCO 原始订单 payload（payload 保留原始 JSON，方便审计/排查）
		order, err := decodeDSCOOrder(row.Payload)
		if err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", po,
					"result", "fail",
					"reason", "decode_dsco_payload_failed",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"err", err,
				)...,
			)
			continue
		}
		if strings.TrimSpace(order.PoNumber) == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", "",
					"result", "skip",
					"reason", "dsco_po_number_empty",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 3) 店铺映射：dscoRetailerId -> 领星 store_id（用于领星创建订单）
		shopKey := strings.TrimSpace(derefString(order.DscoRetailerID))
		storeID := ""
		if shopKey != "" {
			storeID = strings.TrimSpace(ctx.Config.Mapping.Shop[shopKey])
		}
		if storeID == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "skip",
					"reason", "missing_mapping_shop_store_id",
					"shop_key", shopKey,
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 3.1) 仓库映射：requestedWarehouseCode -> 领星 WID（空则允许不传）
		wid := ""
		dscoWarehouse := getDSCOWarehouseCode(order)
		if dscoWarehouse != "" {
			if v, ok := ctx.Config.Mapping.Warehouse[dscoWarehouse]; ok {
				wid = v
			}
		}

		// 3.2) 物流方式映射（用于领星 LogisticsTypeID）：
		// key = <shipWarehouseCode>-<shippingServiceLevelCode>
		// shipWarehouseCode 缺失时允许使用 requestedWarehouseCode 兜底（已确认）。
		logisticsTypeID := ""
		if wid != "" {
			shipWarehouseCode := strings.TrimSpace(derefString(order.ShipWarehouseCode))
			if shipWarehouseCode == "" {
				shipWarehouseCode = strings.TrimSpace(derefString(order.RequestedWarehouseCode))
			}
			serviceLevelCode := strings.TrimSpace(getDSCOShippingServiceLevelCode(order))
			key := shipWarehouseCode + "-" + serviceLevelCode
			logisticsTypeID = strings.TrimSpace(ctx.Config.Mapping.Shipment[key])
			if logisticsTypeID == "" {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "push_to_lingxing",
						"po_number", order.PoNumber,
						"result", "skip",
						"reason", "missing_mapping_shipment_logistics_type_id",
						"key", key,
						"dsco_raw", integration.JSONForLog(row.Payload),
					)...,
				)
				continue
			}
		}

		// 2.2) 地址校验（MVP：仅校验最小必填；更严格校验后续可根据日志再补）
		addr := order.Shipping
		if addr == nil {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "skip",
					"reason", "missing_shipping_address",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}
		country := strings.TrimSpace(derefString(addr.Country))
		if country == "" {
			country = "US"
		}
		name := strings.TrimSpace(derefString(addr.Name))
		if name == "" {
			name = strings.TrimSpace(strings.TrimSpace(derefString(addr.FirstName)) + " " + strings.TrimSpace(derefString(addr.LastName)))
		}
		line1 := ""
		if addr.Address1 != nil && strings.TrimSpace(*addr.Address1) != "" {
			line1 = strings.TrimSpace(*addr.Address1)
		} else if len(addr.Address) > 0 && strings.TrimSpace(addr.Address[0]) != "" {
			line1 = strings.TrimSpace(addr.Address[0])
		}
		if line1 == "" || name == "" || strings.TrimSpace(addr.City) == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "skip",
					"reason", "missing_required_address_fields",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 2.3) 行项目组装：
		// - SKU 不映射：直接使用 DSCO sku/partnerSku 赋值到领星 MSKU（领星侧自动匹配）。
		// - 单价：按字段优先级选择，尽量避免领星校验失败。
		createItems := make([]lingxing.CreateOrderItemV2, 0, len(order.LineItems))
		for _, li := range order.LineItems {
			msku := ""
			// 一期口径：优先使用 DSCO partnerSku，否则使用 sku。
			if li.PartnerSKU != nil && *li.PartnerSKU != "" {
				msku = *li.PartnerSKU
			} else if li.SKU != nil && *li.SKU != "" {
				msku = *li.SKU
			}
			if msku == "" {
				continue
			}
			unitPrice, ok := pickUnitPrice(li)
			if !ok {
				// 行项目级别的跳过，不终止整单；仍保留在最终 createItems 为空时做订单级别 skip。
				logger.Warn(taskCtx, "dsco line item unit price missing",
					append(base,
						"task", "push_to_lingxing",
						"po_number", order.PoNumber,
						"sku", msku,
					)...,
				)
				continue
			}
			createItems = append(createItems, lingxing.CreateOrderItemV2{
				MSKU:      msku,
				Quantity:  li.Quantity,
				UnitPrice: unitPrice,
			})
		}
		if len(createItems) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "skip",
					"reason", "no_valid_items",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 5) 调用领星 CreateOrdersV2（最小必要字段）
		req := lingxing.CreateOrdersV2Request{
			PlatformCode: lingxing.PlatformCode(d.env.Integration.LingXing.PlatformCode),
			StoreID:      storeID,
			Orders: []lingxing.CreateOrderV2{
				{
					PlatformOrderNo: order.PoNumber,
					// 收件信息
					ReceiverName:        name,                                               // 收件人
					BuyerName:           name,                                               // 买家姓名（暂与收件人相同）
					BuyerNote:           strings.TrimSpace(derefString(order.BuyerMessage)), // 买家留言
					BuyerEmail:          *addr.Email,                                        // 买家邮箱
					ReceiverMobile:      *addr.Phone,                                        // 手机号
					ReceiverCountryCode: country,                                            // 国家代码
					City:                strings.TrimSpace(addr.City),                       // 城市
					AddressLine1:        line1,                                              // 地址行1
					PostalCode:          addr.Postal,                                        // 邮编
					// 仓库和物流信息
					WID:             wid,             // 领星仓库ID
					LogisticsTypeID: logisticsTypeID, // 领星物流 type_id
					// 币种
					AmountCurrency: "USD", // 一期口径：DSCO 订单均为美国站点，币种固定 USD
					// 订单
					Items: createItems,
				},
			},
		}

		resp, err := lx.Order.CreateOrdersV2(taskCtx, req)
		if err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "fail",
					"reason", "lingxing_create_order_failed",
					"create_req", integration.JSONForLog(req),
					"dsco_raw", integration.JSONForLog(row.Payload),
					"err", err,
				)...,
			)
			continue
		}
		if len(resp.SuccessDetails) == 0 && len(resp.ErrorDetails) > 0 {
			// 保持简单：不推进状态，下次定时任务自动重试；失败原因仅日志。
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "fail",
					"reason", "lingxing_create_order_error_detail",
					"create_req", integration.JSONForLog(req),
					"resp", integration.JSONForLog(resp),
					"msg", resp.ErrorDetails[0].ErrorMessage,
				)...,
			)
			continue
		}

		// 6) 推单成功：推进状态到 2（待回传 ACK）
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, order.PoNumber, 2, "", ""); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "push_to_lingxing",
					"po_number", order.PoNumber,
					"result", "fail",
					"reason", "update_local_status_failed_after_create",
					"resp", integration.JSONForLog(resp),
					"err", err,
				)...,
			)
			continue
		}
		okCount++
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "push_to_lingxing",
				"po_number", order.PoNumber,
				"result", "ok",
				"reason", "lingxing_create_order_ok",
				"create_req", integration.JSONForLog(req),
				"resp", integration.JSONForLog(resp),
				"new_status", 2,
			)...,
		)
	}
	return nil
}

// pickUnitPrice 选择“最可能可用”的单价字段。
//
// 说明：
// - 一期目标是“先跑通闭环”，当遇到领星侧对价格字段的校验时，优先保证有值。
// - 若后续需要严格口径（例如税前/税后/平台价），再根据业务确认调整优先级。
func pickUnitPrice(li dsco.OrderLineItem) (float64, bool) {
	if li.ConsumerPriceWithTax != nil && *li.ConsumerPriceWithTax > 0 {
		return *li.ConsumerPriceWithTax, true
	}
	if li.ConsumerPrice != nil && *li.ConsumerPrice > 0 {
		return *li.ConsumerPrice, true
	}
	if li.RetailPrice != nil && *li.RetailPrice > 0 {
		return *li.RetailPrice, true
	}
	if li.ExpectedCost != nil && *li.ExpectedCost > 0 {
		return *li.ExpectedCost, true
	}
	return 0, false
}
