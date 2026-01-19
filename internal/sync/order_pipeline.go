package sync

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"
	lxorder "lingxingipass/internal/lingxing/order"
	"lingxingipass/internal/platform/retry"
	"lingxingipass/internal/store"
)

const jobNamePullDSCOOrders = "pull_dsco_orders"
const jobNamePushOrdersToLingXing = "push_orders_to_lingxing"
const jobNameAckToDSCO = "ack_to_dsco"
const jobNameShipToDSCO = "ship_to_dsco"
const jobNameInvoiceToDSCO = "invoice_to_dsco"

// OrderPipeline 串联一期订单闭环的编排逻辑（先落地 Step3/Step4）。
type OrderPipeline struct {
	dscoCli     *dsco.Client
	lingxingCli *lingxing.Client

	orderState *store.OrderStateStore
	watermark  *store.WatermarkStore
	manualTask *store.ManualTaskStore
	orderRaw   *store.DscoOrderRawStore

	now func() time.Time

	// maxRetryPerOrder 表示同一 dscoOrderId 在单个环节的最大重试次数；达到上限后转人工，避免无限重试。
	maxRetryPerOrder int
	// shipDateSource 表示发货回传的 shipDate 取值来源（delivered_at/stock_delivered_at/none）。
	shipDateSource string
}

type PullDSCOOrdersWatermark struct {
	Mode  string `json:"mode"`  // 一期固定 updatedSince
	Since string `json:"since"` // RFC3339
}

func NewOrderPipeline(dscoCli *dsco.Client, lingxingCli *lingxing.Client, orderState *store.OrderStateStore, watermark *store.WatermarkStore, manualTask *store.ManualTaskStore, orderRaw *store.DscoOrderRawStore, now func() time.Time, maxRetryPerOrder int, shipDateSource string) (*OrderPipeline, error) {
	if dscoCli == nil {
		return nil, errors.New("dscoCli 不能为空")
	}
	if orderState == nil {
		return nil, errors.New("orderState 不能为空")
	}
	if watermark == nil {
		return nil, errors.New("watermark 不能为空")
	}
	if manualTask == nil {
		return nil, errors.New("manualTask 不能为空")
	}
	if orderRaw == nil {
		return nil, errors.New("orderRaw 不能为空")
	}
	if now == nil {
		now = time.Now
	}
	if maxRetryPerOrder <= 0 {
		maxRetryPerOrder = 5
	}
	shipDateSource = strings.TrimSpace(shipDateSource)
	if shipDateSource == "" {
		shipDateSource = "delivered_at"
	}
	switch shipDateSource {
	case "delivered_at", "stock_delivered_at", "none":
	default:
		return nil, errors.New("shipDateSource 仅支持 delivered_at/stock_delivered_at/none")
	}
	return &OrderPipeline{
		dscoCli:          dscoCli,
		lingxingCli:      lingxingCli,
		orderState:       orderState,
		watermark:        watermark,
		manualTask:       manualTask,
		orderRaw:         orderRaw,
		now:              now,
		maxRetryPerOrder: maxRetryPerOrder,
		shipDateSource:   shipDateSource,
	}, nil
}

func (p *OrderPipeline) tryTurnManualOnRetryExceeded(ctx context.Context, dscoOrderID string, stage string, reason string, markManual func(ctx context.Context, dscoOrderID string, reason string) error) {
	if p.maxRetryPerOrder <= 0 || markManual == nil {
		return
	}
	n, ok, err := p.orderState.GetRetryCount(ctx, dscoOrderID)
	if err != nil || !ok {
		return
	}
	if n < p.maxRetryPerOrder {
		return
	}
	_ = p.manualTask.Create(ctx, store.ManualTask{
		TaskType:    "max_retry_exceeded",
		DscoOrderID: dscoOrderID,
		Payload:     []byte(fmt.Sprintf(`{"stage":%q,"reason":%q,"retry_count":%d,"max_retry":%d}`, stage, reason, n, p.maxRetryPerOrder)),
	})
	_ = markManual(ctx, dscoOrderID, "max_retry_exceeded")
}

// PullOrders 增量拉取 DSCO 订单，只落库 dscoOrderId，并推进水位。
func (p *OrderPipeline) PullOrders(ctx context.Context) error {
	logger.Info(ctx, "pull_orders_start", "job", jobNamePullDSCOOrders)
	raw, ok, err := p.watermark.Get(ctx, jobNamePullDSCOOrders)
	if err != nil {
		return err
	}
	if !ok {
		// 起始水位从 0 开始：updatedSince=1970-01-01T00:00:00Z
		raw = []byte(`{"mode":"updatedSince","since":"1970-01-01T00:00:00Z"}`)
		_ = p.watermark.Set(ctx, jobNamePullDSCOOrders, raw)
	}

	var wm PullDSCOOrdersWatermark
	if err := json.Unmarshal(raw, &wm); err != nil {
		return err
	}
	if wm.Mode == "" {
		wm.Mode = "updatedSince"
	}
	if wm.Mode != "updatedSince" {
		return errors.New("pull_dsco_orders.watermark.mode 仅支持 updatedSince")
	}
	if wm.Since == "" {
		return errors.New("pull_dsco_orders.watermark.since 不能为空")
	}
	logger.Info(ctx, "pull_orders_watermark", "job", jobNamePullDSCOOrders, "since", wm.Since)

	// DSCO 要求 until 必须在过去至少 5 秒；这里保守取 10 秒。
	until := p.now().UTC().Add(-10 * time.Second).Format(time.RFC3339)

	q := dsco.OrderPageQuery{
		OrdersUpdatedSince: wm.Since,
		Until:              until,
		OrdersPerPage:      1000,
		Status:             []string{"created", "shipment_pending"},
	}

	for {
		var page *dsco.PagedOrderResultRaw
		callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
			var err error
			page, err = p.dscoCli.Order.GetPageRaw(ctx, q)
			return err
		})
		if callErr != nil {
			return callErr
		}

		var ids []string
		for _, rawOrder := range page.Orders {
			var tmp struct {
				DscoOrderID string `json:"dscoOrderId"`
			}
			if err := json.Unmarshal(rawOrder, &tmp); err != nil || tmp.DscoOrderID == "" {
				_ = p.manualTask.Create(ctx, store.ManualTask{
					TaskType:    "bad_payload",
					DscoOrderID: "",
					Payload:     []byte(`{"reason":"dscoOrderId 为空或无法解析"}`),
				})
				continue
			}
			ids = append(ids, tmp.DscoOrderID)
			// 保存 DSCO 原始订单快照（只存最新一份）
			_ = p.orderRaw.UpsertLatest(ctx, tmp.DscoOrderID, rawOrder, p.now().UTC())
		}
		logger.Info(ctx, "pull_orders_upsert", "job", jobNamePullDSCOOrders, "count", len(ids), "scroll_id", page.ScrollID)
		if err := p.orderState.UpsertOrderIDs(ctx, ids); err != nil {
			return err
		}

		// 翻页结束条件：空页 + 没有 scrollId。
		if page.ScrollID == "" {
			break
		}
		if len(page.Orders) == 0 {
			break
		}
		q = dsco.OrderPageQuery{ScrollID: page.ScrollID}
	}

	next := PullDSCOOrdersWatermark{
		Mode:  "updatedSince",
		Since: until,
	}
	nextRaw, _ := json.Marshal(next)
	if err := p.watermark.Set(ctx, jobNamePullDSCOOrders, nextRaw); err != nil {
		return err
	}
	logger.Info(ctx, "pull_orders_end", "job", jobNamePullDSCOOrders, "next_since", until)
	return nil
}

// PushOrdersToLingXing 将待推单的 DSCO 订单推送到领星（一期：每次仅创建单笔订单）。
//
// 规则（一期最小实现）：
// - 候选集：sync_order_state.pushed_to_lx_status in (0,2)（由 store 做抢占）
// - 回源：DSCO GET /order/?orderKey=dscoOrderId&value=<dscoOrderId>
// - 映射：缺必填字段 => manual_task(bad_payload) + pushed_to_lx_status=3
// - 推单：领星 /pb/mp/order/v2/create，成功写入 global_order_no
// - 若领星返回“已存在/重复”导致无法获取 global_order_no，则尝试用 /pb/mp/order/v2/list 回查并补齐本地状态
func (p *OrderPipeline) PushOrdersToLingXing(ctx context.Context, platformCode int, storeID string, batchSize int) error {
	if p.lingxingCli == nil {
		return errors.New("lingxingCli 不能为空")
	}
	if platformCode <= 0 {
		return errors.New("platformCode 必须为正整数")
	}
	storeID = strings.TrimSpace(storeID)
	if storeID == "" {
		return errors.New("storeID 不能为空")
	}
	if batchSize <= 0 {
		return errors.New("batchSize 必须大于 0")
	}

	logger.Info(ctx, "push_orders_start", "job", jobNamePushOrdersToLingXing, "batch_size", batchSize)
	ids, err := p.orderState.ClaimForPush(ctx, batchSize)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		logger.Info(ctx, "push_orders_noop", "job", jobNamePushOrdersToLingXing)
		return nil
	}
	logger.Info(ctx, "push_orders_claimed", "job", jobNamePushOrdersToLingXing, "count", len(ids))

	for _, dscoOrderID := range ids {
		dscoOrderID = strings.TrimSpace(dscoOrderID)
		if dscoOrderID == "" {
			continue
		}

		var dscoOrder *dsco.Order
		dscoErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
			var err error
			dscoOrder, err = p.dscoCli.Order.GetByKey(ctx, "dscoOrderId", dscoOrderID, nil)
			return err
		})
		if dscoErr != nil {
			_ = p.orderState.MarkPushFailure(ctx, dscoOrderID, dscoErr.Error())
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "push_to_lingxing", dscoErr.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkPushManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "push_orders_dsco_failed", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "err", dscoErr.Error())
			continue
		}

		lxOrder, mapErr := lxorder.MapCreateOrderV2FromDSCO(dscoOrder)
		if mapErr != nil {
			_ = p.manualTask.Create(ctx, store.ManualTask{
				TaskType:    "bad_payload",
				DscoOrderID: dscoOrderID,
				Payload:     []byte(fmt.Sprintf(`{"reason":%q}`, mapErr.Error())),
			})
			_ = p.orderState.MarkPushManual(ctx, dscoOrderID, "bad_payload")
			logger.Warn(ctx, "push_orders_bad_payload", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "err", mapErr.Error())
			continue
		}

		createReq := lingxing.CreateOrdersV2Request{
			PlatformCode: platformCode,
			StoreID:      storeID,
			Orders:       []lingxing.CreateOrderV2{lxOrder},
		}

		var createResp lingxing.CreateOrdersV2ResponseData
		lxErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
			var err error
			createResp, err = p.lingxingCli.Order.CreateOrdersV2(ctx, createReq)
			return err
		})
		if lxErr != nil {
			_ = p.orderState.MarkPushFailure(ctx, dscoOrderID, lxErr.Error())
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "push_to_lingxing", lxErr.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkPushManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "push_orders_lingxing_failed", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "err", lxErr.Error())
			continue
		}

		if globalOrderNo, ok := pickCreateOrdersV2GlobalOrderNo(createResp, dscoOrderID); ok {
			_ = p.orderState.MarkPushSuccess(ctx, dscoOrderID, globalOrderNo)
			logger.Info(ctx, "push_orders_success", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "lingxing_global_order_no", globalOrderNo)
			continue
		}

		// 如果创建接口没返回 success_detail，尝试回查领星订单列表补齐 global_order_no。
		globalOrderNo, ok, qErr := p.queryLingXingGlobalOrderNo(ctx, platformCode, storeID, dscoOrderID)
		if qErr == nil && ok {
			_ = p.orderState.MarkPushSuccess(ctx, dscoOrderID, globalOrderNo)
			logger.Info(ctx, "push_orders_success", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "lingxing_global_order_no", globalOrderNo)
			continue
		}

		errMsg := pickCreateOrdersV2ErrorMessage(createResp, dscoOrderID)
		if errMsg == "" && qErr != nil {
			errMsg = qErr.Error()
		}
		if errMsg == "" {
			errMsg = "领星创建订单未返回 success_detail"
		}
		_ = p.orderState.MarkPushFailure(ctx, dscoOrderID, errMsg)
		p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "push_to_lingxing", errMsg, func(ctx context.Context, dscoOrderID string, reason string) error {
			return p.orderState.MarkPushManual(ctx, dscoOrderID, reason)
		})
		logger.Error(ctx, "push_orders_failed", "job", jobNamePushOrdersToLingXing, "dsco_order_id", dscoOrderID, "err", errMsg)
	}

	logger.Info(ctx, "push_orders_end", "job", jobNamePushOrdersToLingXing)
	return nil
}

func shouldRetryDSCO(err error) bool {
	var apiErr *dsco.APIError
	if errors.As(err, &apiErr) {
		if apiErr.StatusCode == 429 || apiErr.StatusCode >= 500 {
			return true
		}
		return false
	}
	return false
}

func shouldRetryLingXing(err error) bool {
	var apiErr *lingxing.APIError
	if errors.As(err, &apiErr) {
		if apiErr.StatusCode == 429 || apiErr.StatusCode >= 500 {
			return true
		}
		return false
	}
	return false
}

func pickCreateOrdersV2GlobalOrderNo(resp lingxing.CreateOrdersV2ResponseData, platformOrderNo string) (string, bool) {
	for _, it := range resp.SuccessDetails {
		if strings.TrimSpace(it.PlatformOrderNo) == platformOrderNo && strings.TrimSpace(it.GlobalOrderNo) != "" {
			return strings.TrimSpace(it.GlobalOrderNo), true
		}
	}
	return "", false
}

func pickCreateOrdersV2ErrorMessage(resp lingxing.CreateOrdersV2ResponseData, platformOrderNo string) string {
	for _, it := range resp.ErrorDetails {
		if strings.TrimSpace(it.PlatformOrderNo) == platformOrderNo {
			return strings.TrimSpace(it.ErrorMessage)
		}
	}
	if len(resp.ErrorDetails) > 0 {
		return strings.TrimSpace(resp.ErrorDetails[0].ErrorMessage)
	}
	return ""
}

func (p *OrderPipeline) queryLingXingGlobalOrderNo(ctx context.Context, platformCode int, storeID string, dscoOrderID string) (string, bool, error) {
	// 领星 list 接口必须填 offset/length；一期按单号精确查，长度给 20 足够。
	var resp lingxing.OrderListV2ResponseData
	callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
		var err error
		resp, err = p.lingxingCli.Order.ListOrdersV2(ctx, lingxing.OrderListV2Request{
			Offset:           0,
			Length:           20,
			StoreID:          []string{storeID},
			PlatformCode:     []int{platformCode},
			PlatformOrderNos: []string{dscoOrderID},
		})
		return err
	})
	if callErr != nil {
		return "", false, callErr
	}

	// list 元素字段很多，这里按一期口径提取 global_order_no + platform_info.platform_order_no。
	for _, raw := range resp.List {
		var item lingxing.OrderListV2Item
		if err := json.Unmarshal(raw, &item); err != nil {
			continue
		}
		globalOrderNo := strings.TrimSpace(item.GlobalOrderNo)
		if globalOrderNo == "" {
			continue
		}
		for _, pi := range item.PlatformInfo {
			if strings.TrimSpace(pi.PlatformOrderNo) == dscoOrderID {
				return globalOrderNo, true, nil
			}
		}
	}
	return "", false, nil
}

type AckToDSCOWatermark struct {
	Mode  string `json:"mode"`  // 一期固定 update_time
	Since int64  `json:"since"` // Unix 秒
}

// AckToDSCO 从领星“待发货(5)”订单中筛选平台单号，并回传 DSCO acknowledge。
//
// 注意：该任务使用 job_watermark.ack_to_dsco 控制查询时间窗口；一期不做自动初始化，避免误拉全量。
func (p *OrderPipeline) AckToDSCO(ctx context.Context, platformCode int, storeID string) error {
	if p.lingxingCli == nil {
		return errors.New("lingxingCli 不能为空")
	}
	if platformCode <= 0 {
		return errors.New("platformCode 必须为正整数")
	}
	storeID = strings.TrimSpace(storeID)
	if storeID == "" {
		return errors.New("storeID 不能为空")
	}

	logger.Info(ctx, "ack_start", "job", jobNameAckToDSCO)

	raw, ok, err := p.watermark.Get(ctx, jobNameAckToDSCO)
	if err != nil {
		return err
	}
	if !ok {
		// 起始水位从 0 开始：since=0 表示“未初始化/从头”，实际查询时会按领星限制裁剪为最近 30 天窗口。
		raw = []byte(`{"mode":"update_time","since":0}`)
		_ = p.watermark.Set(ctx, jobNameAckToDSCO, raw)
	}

	var wm AckToDSCOWatermark
	if err := json.Unmarshal(raw, &wm); err != nil {
		return err
	}
	if wm.Mode == "" {
		wm.Mode = "update_time"
	}
	if wm.Mode != "update_time" {
		return errors.New("ack_to_dsco.watermark.mode 仅支持 update_time")
	}
	if wm.Since < 0 {
		return errors.New("ack_to_dsco.watermark.since 不能为负数")
	}

	until := p.now().UTC().Add(-10 * time.Second).Unix()
	startTime := wm.Since
	if startTime == 0 {
		startTime = until - int64((30 * 24 * time.Hour).Seconds())
		if startTime < 0 {
			startTime = 0
		}
	}

	offset := 0
	length := 200
	for {
		var resp lingxing.OrderListV2ResponseData
		callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
			var err error
			resp, err = p.lingxingCli.Order.ListOrdersV2(ctx, lingxing.OrderListV2Request{
				Offset:       offset,
				Length:       length,
				DateType:     "update_time",
				StartTime:    startTime,
				EndTime:      until,
				StoreID:      []string{storeID},
				PlatformCode: []int{platformCode},
				OrderStatus:  5,
			})
			return err
		})
		if callErr != nil {
			return callErr
		}
		if len(resp.List) == 0 {
			break
		}

		for _, rawOrder := range resp.List {
			dscoOrderID := parseLingXingPlatformOrderNo(rawOrder)
			if dscoOrderID == "" {
				continue
			}

			_ = p.orderState.UpsertOrderIDs(ctx, []string{dscoOrderID})
			claimed, err := p.orderState.TryClaimAck(ctx, dscoOrderID)
			if err != nil {
				continue
			}
			if !claimed {
				continue
			}

			_, err = p.dscoCli.Order.Acknowledge(ctx, []dsco.OrderAcknowledgeRequest{
				{ID: dscoOrderID, Type: dsco.OrderAcknowledgeIDTypeDscoOrderID},
			})
			if err != nil {
				_ = p.orderState.MarkAckFailure(ctx, dscoOrderID, err.Error())
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ack_to_dsco", err.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkAckManual(ctx, dscoOrderID, reason)
				})
				logger.Error(ctx, "ack_failed", "job", jobNameAckToDSCO, "dsco_order_id", dscoOrderID, "err", err.Error())
				continue
			}
			_ = p.orderState.MarkAckSuccess(ctx, dscoOrderID)
			logger.Info(ctx, "ack_success", "job", jobNameAckToDSCO, "dsco_order_id", dscoOrderID)
		}

		offset += length
		if int64(offset) >= resp.Total {
			break
		}
	}

	next := AckToDSCOWatermark{Mode: "update_time", Since: until}
	nextRaw, _ := json.Marshal(next)
	if err := p.watermark.Set(ctx, jobNameAckToDSCO, nextRaw); err != nil {
		return err
	}
	logger.Info(ctx, "ack_end", "job", jobNameAckToDSCO, "next_since", until)
	return nil
}

func parseLingXingPlatformOrderNo(raw json.RawMessage) string {
	var item lingxing.OrderListV2Item
	if err := json.Unmarshal(raw, &item); err != nil {
		return ""
	}
	if len(item.PlatformInfo) == 0 {
		return ""
	}
	return strings.TrimSpace(item.PlatformInfo[0].PlatformOrderNo)
}

type ShipToDSCOWatermark struct {
	Mode  string `json:"mode"`  // 一期固定 update_time
	Since int64  `json:"since"` // Unix 秒
}

// ShipToDSCO 从领星“已发货(6)”订单中筛选平台单号，权威以 WMS 出库单 status=3 + tracking_no 非空为准，并回传 DSCO singleShipment。
//
// 一期限制：
// - 只回传一个 tracking_no；如果同一平台单号出现多个出库单/多个 tracking_no，写 manual_task(multi_shipment) 转人工。
func (p *OrderPipeline) ShipToDSCO(ctx context.Context, platformCode int, storeID string, sid int) error {
	if p.lingxingCli == nil {
		return errors.New("lingxingCli 不能为空")
	}
	if platformCode <= 0 {
		return errors.New("platformCode 必须为正整数")
	}
	storeID = strings.TrimSpace(storeID)
	if storeID == "" {
		return errors.New("storeID 不能为空")
	}
	if sid <= 0 {
		return errors.New("sid 必须为正整数")
	}

	logger.Info(ctx, "ship_start", "job", jobNameShipToDSCO, "sid", sid)

	raw, ok, err := p.watermark.Get(ctx, jobNameShipToDSCO)
	if err != nil {
		return err
	}
	if !ok {
		raw = []byte(`{"mode":"update_time","since":0}`)
		_ = p.watermark.Set(ctx, jobNameShipToDSCO, raw)
	}

	var wm ShipToDSCOWatermark
	if err := json.Unmarshal(raw, &wm); err != nil {
		return err
	}
	if wm.Mode == "" {
		wm.Mode = "update_time"
	}
	if wm.Mode != "update_time" {
		return errors.New("ship_to_dsco.watermark.mode 仅支持 update_time")
	}
	if wm.Since < 0 {
		return errors.New("ship_to_dsco.watermark.since 不能为负数")
	}

	until := p.now().UTC().Add(-10 * time.Second).Unix()
	startTime := wm.Since
	if startTime == 0 {
		startTime = until - int64((30 * 24 * time.Hour).Seconds())
		if startTime < 0 {
			startTime = 0
		}
	}

	offset := 0
	length := 200
	for {
		var resp lingxing.OrderListV2ResponseData
		callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
			var err error
			resp, err = p.lingxingCli.Order.ListOrdersV2(ctx, lingxing.OrderListV2Request{
				Offset:       offset,
				Length:       length,
				DateType:     "update_time",
				StartTime:    startTime,
				EndTime:      until,
				StoreID:      []string{storeID},
				PlatformCode: []int{platformCode},
				OrderStatus:  6,
			})
			return err
		})
		if callErr != nil {
			return callErr
		}
		if len(resp.List) == 0 {
			break
		}

		for _, rawOrder := range resp.List {
			dscoOrderID := parseLingXingPlatformOrderNo(rawOrder)
			if dscoOrderID == "" {
				continue
			}

			_ = p.orderState.UpsertOrderIDs(ctx, []string{dscoOrderID})
			claimed, err := p.orderState.TryClaimShipment(ctx, dscoOrderID)
			if err != nil || !claimed {
				continue
			}

			wms, err := p.queryWmsShipment(ctx, dscoOrderID, sid)
			if err != nil {
				_ = p.orderState.MarkShipmentFailure(ctx, dscoOrderID, err.Error())
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ship_to_dsco", err.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkShipmentManual(ctx, dscoOrderID, reason)
				})
				logger.Error(ctx, "ship_wms_failed", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID, "err", err.Error())
				continue
			}
			if len(wms) == 0 {
				_ = p.orderState.MarkShipmentFailure(ctx, dscoOrderID, "wmsOrderList 无记录")
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ship_to_dsco", "wmsOrderList 无记录", func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkShipmentManual(ctx, dscoOrderID, reason)
				})
				logger.Error(ctx, "ship_wms_empty", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID)
				continue
			}

			tracking, ok := pickSingleTrackingNo(wms)
			if !ok {
				_ = p.manualTask.Create(ctx, store.ManualTask{
					TaskType:    "multi_shipment",
					DscoOrderID: dscoOrderID,
					Payload:     []byte(`{"reason":"同一订单匹配到多个出库单/多个跟踪号"}`),
				})
				_ = p.orderState.MarkShipmentManual(ctx, dscoOrderID, "multi_shipment")
				logger.Warn(ctx, "ship_multi_shipment", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID)
				continue
			}
			items := buildShipmentLineItems(wms)
			if len(items) == 0 {
				_ = p.orderState.MarkShipmentFailure(ctx, dscoOrderID, "wms product_info 为空")
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ship_to_dsco", "wms product_info 为空", func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkShipmentManual(ctx, dscoOrderID, reason)
				})
				continue
			}

			shipDate := pickShipDateRFC3339(wms, p.shipDateSource)
			resp, err := p.dscoCli.Order.SingleShipment(ctx, dsco.ShipmentsForUpdate{
				DscoOrderID: dscoOrderID,
				Shipments: []dsco.ShipmentForUpdate{
					{
						TrackingNumber: tracking,
						ShipDate:       shipDate,
						LineItems:      items,
					},
				},
			})
			if err != nil {
				_ = p.orderState.MarkShipmentFailure(ctx, dscoOrderID, err.Error())
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ship_to_dsco", err.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkShipmentManual(ctx, dscoOrderID, reason)
				})
				logger.Error(ctx, "ship_dsco_failed", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID, "err", err.Error())
				continue
			}
			if resp == nil || !resp.Success {
				_ = p.orderState.MarkShipmentFailure(ctx, dscoOrderID, "DSCO singleShipment 返回 success=false")
				p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "ship_to_dsco", "DSCO singleShipment 返回 success=false", func(ctx context.Context, dscoOrderID string, reason string) error {
					return p.orderState.MarkShipmentManual(ctx, dscoOrderID, reason)
				})
				logger.Error(ctx, "ship_dsco_failed", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID, "err", "DSCO singleShipment 返回 success=false")
				continue
			}
			_ = p.orderState.MarkShipmentSuccess(ctx, dscoOrderID, tracking)
			logger.Info(ctx, "ship_success", "job", jobNameShipToDSCO, "dsco_order_id", dscoOrderID, "tracking_no", tracking)
		}

		offset += length
		if int64(offset) >= resp.Total {
			break
		}
	}

	next := ShipToDSCOWatermark{Mode: "update_time", Since: until}
	nextRaw, _ := json.Marshal(next)
	if err := p.watermark.Set(ctx, jobNameShipToDSCO, nextRaw); err != nil {
		return err
	}
	logger.Info(ctx, "ship_end", "job", jobNameShipToDSCO, "next_since", until)
	return nil
}

func (p *OrderPipeline) queryWmsShipment(ctx context.Context, platformOrderNo string, sid int) ([]lingxing.WmsOrder, error) {
	var items []lingxing.WmsOrder
	callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
		list, _, err := p.lingxingCli.Warehouse.WmsOrderList(ctx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           20,
			SIDArr:             []int{sid},
			StatusArr:          []int{3},
			PlatformOrderNoArr: []string{platformOrderNo},
		})
		if err != nil {
			return err
		}
		items = list
		return nil
	})
	if callErr != nil {
		return nil, callErr
	}
	return items, nil
}

func pickSingleTrackingNo(items []lingxing.WmsOrder) (string, bool) {
	set := make(map[string]struct{})
	for _, it := range items {
		tn := strings.TrimSpace(it.TrackingNo)
		if tn == "" {
			continue
		}
		set[tn] = struct{}{}
	}
	if len(set) != 1 {
		return "", false
	}
	for tn := range set {
		return tn, true
	}
	return "", false
}

func buildShipmentLineItems(items []lingxing.WmsOrder) []dsco.ShipmentLineItemForUpdate {
	m := make(map[string]int)
	for _, it := range items {
		for _, p := range it.ProductInfo {
			sku := strings.TrimSpace(p.SKU)
			if sku == "" || p.Count <= 0 {
				continue
			}
			m[sku] += p.Count
		}
	}
	out := make([]dsco.ShipmentLineItemForUpdate, 0, len(m))
	for sku, qty := range m {
		out = append(out, dsco.ShipmentLineItemForUpdate{SKU: sku, Quantity: qty})
	}
	return out
}

func pickShipDateRFC3339(items []lingxing.WmsOrder, source string) string {
	source = strings.TrimSpace(source)
	if source == "" {
		source = "delivered_at"
	}
	if source == "none" {
		return ""
	}
	for _, it := range items {
		var raw string
		switch source {
		case "stock_delivered_at":
			raw = strings.TrimSpace(it.StockDeliveredAt)
		default:
			raw = strings.TrimSpace(it.DeliveredAt)
		}
		if raw == "" {
			continue
		}
		if v, ok := parseLingXingDatetimeToRFC3339(raw); ok {
			return v
		}
	}
	return ""
}

func parseLingXingDatetimeToRFC3339(raw string) (string, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false
	}
	// 先尝试 RFC3339（部分接口可能直接返回 ISO8601）。
	if ts, err := time.Parse(time.RFC3339, raw); err == nil {
		return ts.UTC().Format(time.RFC3339), true
	}
	// 领星 WMS 常见格式：2006-01-02 15:04:05（文档为字符串时间）。
	if ts, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.UTC); err == nil {
		return ts.UTC().Format(time.RFC3339), true
	}
	return "", false
}

// InvoiceToDSCO 生成并回传 DSCO 发票（一期最小字段集）。
//
// 幂等：
// - 先查 DSCO `GET /invoice?key=invoiceId&value=...`，若已存在直接视为成功。
// - 本地以 sync_order_state.invoiced_to_dsco_status 控制重复执行。
func (p *OrderPipeline) InvoiceToDSCO(ctx context.Context, batchSize int) error {
	if batchSize <= 0 {
		return errors.New("batchSize 必须大于 0")
	}

	logger.Info(ctx, "invoice_start", "job", jobNameInvoiceToDSCO, "batch_size", batchSize)

	ids, err := p.orderState.ClaimForInvoice(ctx, batchSize)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		logger.Info(ctx, "invoice_noop", "job", jobNameInvoiceToDSCO)
		return nil
	}

	for _, dscoOrderID := range ids {
		dscoOrderID = strings.TrimSpace(dscoOrderID)
		if dscoOrderID == "" {
			continue
		}

		invoiceID := "INV-" + dscoOrderID

		// 幂等检查：DSCO 是否已存在该 invoiceId
		var invResp *dsco.GetInvoicesByIDResponse
		getErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
			var err error
			invResp, err = p.dscoCli.Invoice.GetByID(ctx, dsco.InvoiceGetQuery{Key: "invoiceId", Value: invoiceID})
			return err
		})
		if getErr == nil && invResp != nil && len(invResp.Invoices) > 0 {
			_ = p.orderState.MarkInvoiceSuccess(ctx, dscoOrderID, invoiceID)
			logger.Info(ctx, "invoice_exists", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "invoice_id", invoiceID)
			continue
		}
		if getErr != nil {
			_ = p.orderState.MarkInvoiceFailure(ctx, dscoOrderID, getErr.Error())
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "invoice_to_dsco", getErr.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkInvoiceManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "invoice_get_failed", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "err", getErr.Error())
			continue
		}

		var o *dsco.Order
		orderErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
			var err error
			o, err = p.dscoCli.Order.GetByKey(ctx, "dscoOrderId", dscoOrderID, nil)
			return err
		})
		if orderErr != nil {
			_ = p.orderState.MarkInvoiceFailure(ctx, dscoOrderID, orderErr.Error())
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "invoice_to_dsco", orderErr.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkInvoiceManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "invoice_order_get_failed", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "err", orderErr.Error())
			continue
		}

		inv, err := buildInvoiceFromDSCOOrder(invoiceID, o, p.now().UTC())
		if err != nil {
			_ = p.manualTask.Create(ctx, store.ManualTask{
				TaskType:    "bad_payload",
				DscoOrderID: dscoOrderID,
				Payload:     []byte(fmt.Sprintf(`{"reason":%q}`, err.Error())),
			})
			_ = p.orderState.MarkInvoiceManual(ctx, dscoOrderID, "bad_payload")
			logger.Warn(ctx, "invoice_bad_payload", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "err", err.Error())
			continue
		}

		var createResp *dsco.SuccessFailResponse
		createErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
			var err error
			createResp, err = p.dscoCli.Invoice.CreateSingle(ctx, inv)
			return err
		})
		if createErr != nil {
			_ = p.orderState.MarkInvoiceFailure(ctx, dscoOrderID, createErr.Error())
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "invoice_to_dsco", createErr.Error(), func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkInvoiceManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "invoice_create_failed", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "err", createErr.Error())
			continue
		}
		if createResp == nil || !createResp.Success {
			_ = p.orderState.MarkInvoiceFailure(ctx, dscoOrderID, "DSCO invoice 返回 success=false")
			p.tryTurnManualOnRetryExceeded(ctx, dscoOrderID, "invoice_to_dsco", "DSCO invoice 返回 success=false", func(ctx context.Context, dscoOrderID string, reason string) error {
				return p.orderState.MarkInvoiceManual(ctx, dscoOrderID, reason)
			})
			logger.Error(ctx, "invoice_create_failed", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "err", "DSCO invoice 返回 success=false")
			continue
		}

		_ = p.orderState.MarkInvoiceSuccess(ctx, dscoOrderID, invoiceID)
		logger.Info(ctx, "invoice_success", "job", jobNameInvoiceToDSCO, "dsco_order_id", dscoOrderID, "invoice_id", invoiceID)
	}

	logger.Info(ctx, "invoice_end", "job", jobNameInvoiceToDSCO)
	return nil
}

func buildInvoiceFromDSCOOrder(invoiceID string, o *dsco.Order, now time.Time) (*dsco.Invoice, error) {
	if o == nil {
		return nil, errors.New("dsco order 不能为空")
	}
	if strings.TrimSpace(o.DscoOrderID) == "" {
		return nil, errors.New("缺少 dscoOrderId")
	}
	currency := strings.TrimSpace(o.CurrencyCode)
	if currency == "" {
		return nil, errors.New("缺少 currencyCode")
	}
	if len(o.LineItems) == 0 {
		return nil, errors.New("缺少 lineItems")
	}

	var total float64
	items := make([]dsco.InvoiceLineItem, 0, len(o.LineItems))
	for i, li := range o.LineItems {
		if li.Quantity <= 0 {
			return nil, fmt.Errorf("lineItems[%d] quantity 非法", i)
		}
		sku := strings.TrimSpace(li.SKU)
		if sku == "" {
			sku = strings.TrimSpace(li.PartnerSKU)
		}
		if sku == "" {
			return nil, fmt.Errorf("lineItems[%d] 缺少 sku/partnerSku", i)
		}

		var unitPrice *float64
		if li.ConsumerPrice != nil {
			unitPrice = li.ConsumerPrice
		} else if li.RetailPrice != nil {
			unitPrice = li.RetailPrice
		}
		if unitPrice == nil {
			return nil, fmt.Errorf("lineItems[%d] 缺少 consumerPrice/retailPrice", i)
		}

		items = append(items, dsco.InvoiceLineItem{
			SKU:       sku,
			Quantity:  li.Quantity,
			UnitPrice: *unitPrice,
		})
		total += float64(li.Quantity) * (*unitPrice)
	}

	return &dsco.Invoice{
		InvoiceID:    invoiceID,
		DscoOrderID:  strings.TrimSpace(o.DscoOrderID),
		InvoiceDate:  now.UTC().Format(time.RFC3339),
		CurrencyCode: currency,
		TotalAmount:  total,
		LineItems:    items,
	}, nil
}
