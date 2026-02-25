package dsco_lingxing

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
	"lingxingipass/golib/v2/sdk/lingxing"
	"lingxingipass/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

type dscoOrderBrief struct {
	PONumber         string   `json:"po_number"`
	DscoOrderID      string   `json:"dsco_order_id,omitempty"`
	DscoStatus       string   `json:"dsco_status,omitempty"`
	CreatedDate      string   `json:"created_date,omitempty"`
	PackageTrackings []string `json:"package_trackings,omitempty"`
	LineItems        []struct {
		DscoItemID string `json:"dsco_item_id,omitempty"`
		PartnerSKU string `json:"partner_sku,omitempty"`
		LineNumber int    `json:"line_number,omitempty"`
		Quantity   int    `json:"quantity,omitempty"`
	} `json:"line_items,omitempty"`
}

type dscoInvoiceBrief struct {
	InvoiceID   string  `json:"invoice_id"`
	PONumber    string  `json:"po_number,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	LineItems   []struct {
		DscoItemID string `json:"dsco_item_id,omitempty"`
		PartnerSKU string `json:"partner_sku,omitempty"`
		Quantity   int    `json:"quantity,omitempty"`
	} `json:"line_items,omitempty"`
}

type checkDetail struct {
	Local store.DSCOOrderSyncRow `json:"local"`

	DSCOOrder    *dscoOrderBrief    `json:"dsco_order,omitempty"`
	DSCOInvoices []dscoInvoiceBrief `json:"dsco_invoices,omitempty"`

	LingXingOrderFound bool `json:"lingxing_order_found,omitempty"`
}

func (d *Domain) CheckOrders(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "check_orders")...)

	var (
		total int
		okN   int
		failN int
		skipN int
	)
	defer func() {
		fields := append(base,
			"task", "check_orders",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okN,
			"skip", skipN,
			"fail", failN,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	dscoCli, err := d.dscoClient()
	if err != nil {
		retErr = err
		return retErr
	}
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}

	var override integration.CheckOrdersOverride
	if ctx.Override != nil {
		if v, ok := ctx.Override.(integration.CheckOrdersOverride); ok {
			override = v
		} else {
			retErr = errors.New("invalid override type for check_orders")
			return retErr
		}
	}

	poOnly := strings.TrimSpace(ctx.OnlyPONumber)
	var localRows []store.DSCOOrderSyncRow
	var dscoOrders map[string]dsco.Order

	if poOnly != "" {
		row, ok, gerr := d.orderStore.GetByPONumber(taskCtx, poOnly)
		if gerr != nil {
			retErr = gerr
			return retErr
		}
		if ok {
			localRows = []store.DSCOOrderSyncRow{row}
		}
		dscoOrders = fetchDSCOOrdersByPONumbers(taskCtx, dscoCli, []string{poOnly}, 1)
	} else {
		if override.Start <= 0 || override.End <= 0 || override.End <= override.Start {
			retErr = errors.New("missing or invalid time range")
			return retErr
		}

		// 1) 读取本地范围内订单（按 dsco_create_time）。
		localRows, err = d.listLocalOrdersByCreateTimeRange(taskCtx, override.Start, override.End)
		if err != nil {
			retErr = err
			return retErr
		}

		// 2) 拉取 DSCO 订单（按创建时间范围）。
		dscoOrders, err = fetchDSCOOrdersByCreatedRange(taskCtx, dscoCli, override.Start, override.End, 500, maxInt(ctx.Size, 200))
		if err != nil {
			retErr = err
			return retErr
		}
	}

	localByPO := make(map[string]store.DSCOOrderSyncRow, len(localRows))
	for _, r := range localRows {
		if po := strings.TrimSpace(r.PONumber); po != "" {
			localByPO[po] = r
		}
	}

	// 构造待检查 PO 列表：优先 DSCO 侧订单（用于发现 missing_local_record），并补上本地有但 DSCO 拉不到的单据。
	poSet := make(map[string]struct{}, len(localByPO)+len(dscoOrders))
	for po := range dscoOrders {
		if strings.TrimSpace(po) != "" {
			poSet[po] = struct{}{}
		}
	}
	for po := range localByPO {
		if strings.TrimSpace(po) != "" {
			poSet[po] = struct{}{}
		}
	}
	if poOnly != "" {
		poSet[poOnly] = struct{}{}
	}
	poList := make([]string, 0, len(poSet))
	for po := range poSet {
		poList = append(poList, po)
	}
	poList = uniqueNonEmptyStrings(poList)
	total = len(poList)

	// status=2：批量检查领星订单是否存在。
	lxExistByPO := map[string]bool{}
	{
		var need []string
		for _, po := range poList {
			if r, ok := localByPO[po]; ok && r.Status == 2 {
				need = append(need, po)
			}
		}
		lxExistByPO, err = lingxingOrdersExistByPlatformOrderNos(taskCtx, lx, need, 50)
		if err != nil {
			// 领星查询失败不应导致整个检查失败：按单标记异常即可。
			logger.Warn(taskCtx, "check_orders: lingxing list orders failed", append(base, "task", "check_orders", "err", err)...)
		}
	}

	// status=5：批量检查 DSCO 发票是否存在（按 poNumber 查询）。
	dscoInvByPO := map[string][]dscoInvoiceBrief{}
	{
		var need []string
		for _, po := range poList {
			if r, ok := localByPO[po]; ok && r.Status == 5 {
				need = append(need, po)
			}
		}
		dscoInvByPO = fetchDSCOInvoicesByPONumbers(taskCtx, dscoCli, need, 5)
	}

	var results []integration.OrderCheckResult
	var missingLocal []integration.MissingLocalOrder

	for _, po := range poList {
		local, hasLocal := localByPO[po]
		dscoOrder, hasDSCO := dscoOrders[po]

		if !hasLocal {
			if !hasDSCO {
				failN++
				results = append(results, integration.OrderCheckResult{
					PONumber:    po,
					OK:          false,
					Code:        "order_not_found",
					Message:     "order not found in DSCO and local store",
					LocalStatus: 0,
					DSCOStatus:  "",
				})
				continue
			}

			// DSCO 侧存在但本地没入库：异常
			if hasDSCO {
				missingLocal = append(missingLocal, integration.MissingLocalOrder{
					PONumber:    strings.TrimSpace(dscoOrder.PoNumber),
					DSCOStatus:  strings.TrimSpace(dscoOrder.DscoStatus),
					DSCOOrderID: strings.TrimSpace(dscoOrder.DscoOrderID),
				})
			}
			failN++
			var br *dscoOrderBrief
			if hasDSCO {
				v := buildDSCOOrderBrief(dscoOrder)
				br = &v
			}
			results = append(results, integration.OrderCheckResult{
				PONumber:    po,
				OK:          false,
				Code:        "missing_local_record",
				Message:     "DSCO order exists but local record missing",
				LocalStatus: 0,
				DSCOStatus:  strings.TrimSpace(dscoOrder.DscoStatus),
				Detail: map[string]any{
					"dsco_order": br,
				},
			})
			continue
		}

		// status=1：不检查
		if local.Status == 1 {
			skipN++
			continue
		}

		// 本地有但 DSCO 拉不到：异常（也可能是 DSCO 范围查询遗漏/接口延迟，仍提示）。
		if !hasDSCO {
			failN++
			results = append(results, integration.OrderCheckResult{
				PONumber:    po,
				OK:          false,
				Code:        "dsco_order_not_found",
				Message:     "local record exists but DSCO order not found",
				LocalStatus: local.Status,
				Detail: map[string]any{
					"local": local,
				},
			})
			continue
		}

		dscoSt := strings.TrimSpace(dscoOrder.DscoStatus)
		br := buildDSCOOrderBrief(dscoOrder)
		detail := checkDetail{
			Local:        local,
			DSCOOrder:    &br,
			DSCOInvoices: dscoInvByPO[po],
		}

		var ok bool
		var code string
		var msg string

		switch local.Status {
		case 2:
			found := lxExistByPO[po]
			detail.LingXingOrderFound = found
			ok = found
			if ok {
				code = "ok"
				msg = "lingxing order found"
			} else {
				code = "lingxing_order_missing"
				msg = "lingxing order not found"
			}
		case 3:
			ok = isAnyOf(strings.ToLower(dscoSt), "shipment_pending", "shipped", "completed")
			if ok {
				code = "ok"
				msg = "dsco status allowed for local status=3"
			} else {
				code = "dsco_status_mismatch"
				msg = "dsco status not allowed for local status=3"
			}
		case 4:
			ok = isAnyOf(strings.ToLower(dscoSt), "shipped", "completed")
			if ok {
				code = "ok"
				msg = "dsco status allowed for local status=4"
			} else {
				code = "dsco_status_mismatch"
				msg = "dsco status not allowed for local status=4"
			}
		case 5:
			ok = len(dscoInvByPO[po]) > 0
			if ok {
				code = "ok"
				msg = "dsco invoice exists"
			} else {
				code = "dsco_invoice_missing"
				msg = "dsco invoice not found by poNumber"
			}
		default:
			ok = false
			code = "status_not_supported"
			msg = "local status not supported"
		}

		if ok {
			okN++
		} else {
			failN++
		}

		results = append(results, integration.OrderCheckResult{
			PONumber:    po,
			OK:          ok,
			Code:        code,
			Message:     msg,
			LocalStatus: local.Status,
			DSCOStatus:  dscoSt,
			Detail:      detail,
		})
	}

	resp := integration.CheckOrdersResponse{
		Results:      results,
		MissingLocal: missingLocal,
		Meta: map[string]any{
			"dsco_order_count":  len(dscoOrders),
			"local_order_count": len(localByPO),
		},
	}

	if ctx.Report != nil {
		ctx.Report(resp)
	}
	return nil
}

func (d *Domain) listLocalOrdersByCreateTimeRange(ctx context.Context, start, end int64) ([]store.DSCOOrderSyncRow, error) {
	var out []store.DSCOOrderSyncRow
	offset := 0
	for {
		f := store.DSCOOrderSyncListFilter{
			StartTime: &start,
			EndTime:   &end,
			Offset:    offset,
			Limit:     500,
		}
		items, total, err := d.orderStore.List(ctx, f)
		if err != nil {
			return nil, err
		}
		out = append(out, items...)
		offset += len(items)
		if offset >= int(total) || len(items) == 0 {
			break
		}
	}
	return out, nil
}

func fetchDSCOOrdersByCreatedRange(ctx context.Context, cli *dsco.Client, startSec, endSec int64, perPage int, maxOrders int) (map[string]dsco.Order, error) {
	if cli == nil {
		return map[string]dsco.Order{}, errors.New("dsco client is nil")
	}
	if perPage <= 0 {
		perPage = 200
	}
	if perPage > 1000 {
		perPage = 1000
	}
	if maxOrders <= 0 {
		maxOrders = 200
	}

	until := time.Unix(endSec, 0).UTC()
	now := time.Now().UTC()
	// DSCO 要求 until 必须至少早于当前 5 秒，这里保守一点。
	if until.After(now.Add(-10 * time.Second)) {
		until = now.Add(-10 * time.Second)
	}
	since := time.Unix(startSec, 0).UTC()
	if since.After(until) {
		return map[string]dsco.Order{}, fmt.Errorf("invalid range: since=%s, until=%s", since.Format(time.RFC3339), until.Format(time.RFC3339))
	}

	out := make(map[string]dsco.Order, 256)
	scrollID := ""
	for page := 0; page < 50; page++ {
		q := dsco.OrderPageQuery{
			OrdersPerPage:      perPage,
			OrdersCreatedSince: since.Format(time.RFC3339),
			Until:              until.Format(time.RFC3339),
		}
		if scrollID != "" {
			q.ScrollID = scrollID
			// scrollId 模式下 DSCO 会忽略其它字段，但保留这些字段对我们无害。
		}
		resp, err := cli.Order.GetPage(ctx, q)
		if err != nil {
			return nil, err
		}
		if resp == nil || len(resp.Orders) == 0 {
			break
		}
		for _, o := range resp.Orders {
			po := strings.TrimSpace(o.PoNumber)
			if po == "" {
				continue
			}
			out[po] = o
			if len(out) >= maxOrders {
				return out, nil
			}
		}
		scrollID = strings.TrimSpace(resp.ScrollID)
		if scrollID == "" {
			break
		}
	}
	return out, nil
}

func fetchDSCOInvoicesByPONumbers(ctx context.Context, cli *dsco.Client, poNumbers []string, maxConcurrent int) map[string][]dscoInvoiceBrief {
	poNumbers = uniqueNonEmptyStrings(poNumbers)
	if len(poNumbers) == 0 || cli == nil {
		return map[string][]dscoInvoiceBrief{}
	}
	if maxConcurrent <= 0 {
		maxConcurrent = 5
	}

	sem := make(chan struct{}, maxConcurrent)
	type invOut struct {
		po  string
		val []dscoInvoiceBrief
	}
	ch := make(chan invOut, len(poNumbers))

	for _, po := range poNumbers {
		po := po
		go func() {
			sem <- struct{}{}
			defer func() { <-sem }()
			resp, err := cli.Invoice.GetByID(ctx, dsco.InvoiceGetQuery{Key: "poNumber", Value: po})
			if err != nil || resp == nil || len(resp.Invoices) == 0 {
				ch <- invOut{po: po, val: nil}
				return
			}
			var arr []dscoInvoiceBrief
			for _, inv := range resp.Invoices {
				b := dscoInvoiceBrief{
					InvoiceID:   inv.InvoiceID,
					PONumber:    inv.PoNumber,
					TotalAmount: inv.TotalAmount,
				}
				for _, li := range inv.LineItems {
					item := struct {
						DscoItemID string `json:"dsco_item_id,omitempty"`
						PartnerSKU string `json:"partner_sku,omitempty"`
						Quantity   int    `json:"quantity,omitempty"`
					}{Quantity: li.Quantity, PartnerSKU: li.PartnerSKU}
					if strings.TrimSpace(li.DscoItemID) != "" {
						item.DscoItemID = strings.TrimSpace(li.DscoItemID)
					}
					b.LineItems = append(b.LineItems, item)
				}
				arr = append(arr, b)
			}
			ch <- invOut{po: po, val: arr}
		}()
	}

	out := make(map[string][]dscoInvoiceBrief, len(poNumbers))
	for i := 0; i < len(poNumbers); i++ {
		r := <-ch
		if len(r.val) > 0 {
			out[r.po] = r.val
		}
	}
	return out
}

func lingxingOrdersExistByPlatformOrderNos(ctx context.Context, lx *lingxing.Client, poNumbers []string, chunkSize int) (map[string]bool, error) {
	poNumbers = uniqueNonEmptyStrings(poNumbers)
	out := make(map[string]bool, len(poNumbers))
	if len(poNumbers) == 0 {
		return out, nil
	}
	if lx == nil {
		return out, errors.New("lingxing client is nil")
	}
	if chunkSize <= 0 {
		chunkSize = 50
	}

	for _, chunk := range chunkStrings(poNumbers, chunkSize) {
		// 领星要求 length 最小 20
		length := len(chunk)
		if length < 20 {
			length = 20
		}
		if length > 200 {
			length = 200
		}
		resp, err := lx.Order.ListOrdersV2(ctx, lingxing.OrderListV2Request{
			Offset:           0,
			Length:           length,
			PlatformOrderNos: chunk,
		})
		if err != nil {
			return out, err
		}
		for _, d := range resp.List {
			po := poNumberFromLingXingOrderDetail(d)
			if po != "" {
				out[po] = true
			}
		}
		// 未返回的视为不存在
		for _, po := range chunk {
			if _, ok := out[po]; !ok {
				out[po] = false
			}
		}
	}
	return out, nil
}

func buildDSCOOrderBrief(o dsco.Order) dscoOrderBrief {
	createdDate := ""
	if o.DscoCreateDate != nil && strings.TrimSpace(*o.DscoCreateDate) != "" {
		createdDate = strings.TrimSpace(*o.DscoCreateDate)
	} else if o.RetailerCreateDate != nil && strings.TrimSpace(*o.RetailerCreateDate) != "" {
		createdDate = strings.TrimSpace(*o.RetailerCreateDate)
	}

	b := dscoOrderBrief{
		PONumber:    strings.TrimSpace(o.PoNumber),
		DscoStatus:  strings.TrimSpace(o.DscoStatus),
		DscoOrderID: strings.TrimSpace(o.DscoOrderID),
		CreatedDate: createdDate,
	}
	for _, p := range o.Packages {
		if s := strings.TrimSpace(p.TrackingNumber); s != "" {
			b.PackageTrackings = append(b.PackageTrackings, s)
		}
	}
	for _, li := range o.LineItems {
		item := struct {
			DscoItemID string `json:"dsco_item_id,omitempty"`
			PartnerSKU string `json:"partner_sku,omitempty"`
			LineNumber int    `json:"line_number,omitempty"`
			Quantity   int    `json:"quantity,omitempty"`
		}{Quantity: li.Quantity}
		if li.DscoItemID != nil {
			item.DscoItemID = strings.TrimSpace(*li.DscoItemID)
		}
		if li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "" {
			item.PartnerSKU = strings.TrimSpace(*li.PartnerSKU)
		} else if li.SKU != nil && strings.TrimSpace(*li.SKU) != "" {
			item.PartnerSKU = strings.TrimSpace(*li.SKU)
		}
		if li.LineNumber != nil && *li.LineNumber > 0 {
			item.LineNumber = *li.LineNumber
		}
		b.LineItems = append(b.LineItems, item)
	}
	b.PackageTrackings = uniqueNonEmptyStrings(b.PackageTrackings)
	return b
}

func isAnyOf(v string, allowed ...string) bool {
	for _, a := range allowed {
		if v == a {
			return true
		}
	}
	return false
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
