package dsco_lingxing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"

	"lingxingipass/infra/runtimecfg"
)

var (
	ErrMissingRequiredField = errors.New("missing required field")
)

func parseRFC3339ToUnixSec(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("time empty")
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return 0, err
	}
	return t.UTC().Unix(), nil
}

func parseLingXingDateTimeToRFC3339UTC(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("time empty")
	}
	// LingXing often returns "2006-01-02 15:04:05" without timezone.
	// Current project standardizes all internal time to UTC.
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
	if err != nil {
		return "", err
	}
	return t.UTC().Format(time.RFC3339), nil
}

func parseInt64(v any) (int64, error) {
	var s string
	switch t := v.(type) {
	case nil:
		return 0, errors.New("value is nil")
	case string:
		s = t
	case interface{ String() string }:
		s = t.String()
	default:
		s = fmt.Sprint(v)
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("value empty")
	}
	return strconv.ParseInt(s, 10, 64)
}

func reverseMapStrict(m map[string]string) (map[string]string, error) {
	out := make(map[string]string, len(m))
	for k, v := range m {
		if k == "" || v == "" {
			return nil, errors.New("mapping contains empty key/value")
		}
		if old, ok := out[v]; ok && old != k {
			return nil, fmt.Errorf("mapping reverse conflict: %q maps to both %q and %q", v, old, k)
		}
		out[v] = k
	}
	return out, nil
}

func decodeDSCOOrder(payload json.RawMessage) (dsco.Order, error) {
	var o dsco.Order
	if err := json.Unmarshal(payload, &o); err != nil {
		return dsco.Order{}, err
	}
	return o, nil
}

func uniqueNonEmptyStrings(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, it := range items {
		s := strings.TrimSpace(it)
		if s == "" {
			continue
		}
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func chunkStrings(items []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		chunkSize = 50
	}
	if len(items) == 0 {
		return nil
	}
	var out [][]string
	for i := 0; i < len(items); i += chunkSize {
		j := i + chunkSize
		if j > len(items) {
			j = len(items)
		}
		out = append(out, items[i:j])
	}
	return out
}

// poNumberFromLingXingOrderDetail 从领星订单详情中提取 platform_order_no（即 DSCO poNumber）。
//
// 说明：OrderDetailV2 的“平台单号”字段在 PlatformInfo 列表中。
func poNumberFromLingXingOrderDetail(d lingxing.OrderDetailV2) string {
	for _, p := range d.PlatformInfo {
		if s := strings.TrimSpace(p.PlatformOrderNo); s != "" {
			return s
		}
	}
	return ""
}

const dscoOrderKeyPoNumber = "poNumber"

// fetchDSCOOrdersByPONumbers 根据 poNumber 批量获取 DSCO 订单对象（用于检查 dscoStatus）。
//
// 注意：
//   - DSCO OpenAPI 并未提供“按 poNumber 列表批量查询订单对象”的接口；
//     /order/page 只能按时间窗或 consumerOrderNumber 查。
//   - 因此这里采用“并发受限的逐单 GET /order/（orderKey=poNumber）”，以减少总耗时，同时避免触发 DSCO 的限流。
func fetchDSCOOrdersByPONumbers(ctx context.Context, cli *dsco.Client, poNumbers []string, maxConcurrent int) map[string]dsco.Order {
	poNumbers = uniqueNonEmptyStrings(poNumbers)
	if len(poNumbers) == 0 || cli == nil {
		return map[string]dsco.Order{}
	}
	if maxConcurrent <= 0 {
		maxConcurrent = 5
	}

	sem := make(chan struct{}, maxConcurrent)
	out := make(map[string]dsco.Order, len(poNumbers))
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, po := range poNumbers {
		po := po
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			o, err := cli.Order.GetByKey(ctx, dscoOrderKeyPoNumber, po, nil)
			if err != nil || o == nil {
				return
			}
			mu.Lock()
			out[po] = *o
			mu.Unlock()
		}()
	}

	wg.Wait()
	return out
}

func getDSCOShipMethod(o dsco.Order) string {
	if o.RequestedShipMethod != nil && strings.TrimSpace(*o.RequestedShipMethod) != "" {
		return strings.TrimSpace(*o.RequestedShipMethod)
	}
	if o.ShipMethod != nil && strings.TrimSpace(*o.ShipMethod) != "" {
		return strings.TrimSpace(*o.ShipMethod)
	}
	return ""
}

func getDSCOWarehouseCode(o dsco.Order) string {
	if o.RequestedWarehouseCode != nil && strings.TrimSpace(*o.RequestedWarehouseCode) != "" {
		return strings.TrimSpace(*o.RequestedWarehouseCode)
	}
	// Fallbacks exist in DSCO order schema, but are not required by our MVP.
	return ""
}

func getDSCOShippingServiceLevelCode(o dsco.Order) string {
	if o.ShippingServiceLevelCode != nil && strings.TrimSpace(*o.ShippingServiceLevelCode) != "" {
		return strings.TrimSpace(*o.ShippingServiceLevelCode)
	}
	if o.RequestedShippingServiceLevelCode != nil && strings.TrimSpace(*o.RequestedShippingServiceLevelCode) != "" {
		return strings.TrimSpace(*o.RequestedShippingServiceLevelCode)
	}
	if o.RequestedShippingServiceLevelCodeUnmapped != nil && strings.TrimSpace(*o.RequestedShippingServiceLevelCodeUnmapped) != "" {
		return strings.TrimSpace(*o.RequestedShippingServiceLevelCodeUnmapped)
	}
	return ""
}

func buildReverseSKUMap(cfg runtimecfg.Config) (map[string]string, error) {
	// 一期（已确认）：推单到领星时 SKU 不做映射，直接把 DSCO SKU 赋值到领星 MSKU，由领星侧自动匹配。
	// 因此 runtime_config.mapping.sku 暂不使用，保留字段仅为未来扩展位。
	return map[string]string{}, nil
}

func shopKeyFromDSCOOrder(o dsco.Order) string {
	if o.DscoRetailerID != nil && strings.TrimSpace(*o.DscoRetailerID) != "" {
		return strings.TrimSpace(*o.DscoRetailerID)
	}
	return ""
}

func lingxingSIDFromMapping(cfg runtimecfg.Config, o dsco.Order) (int, bool) {
	key := shopKeyFromDSCOOrder(o)
	if key == "" {
		return 0, false
	}
	raw := strings.TrimSpace(cfg.Mapping.Shop[key])
	if raw == "" {
		return 0, false
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v <= 0 {
		return 0, false
	}
	return v, true
}
