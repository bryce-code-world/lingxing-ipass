package dsco_lingxing

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"

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
	// mapping.sku is DSCO -> 领星；回传侧需要 领星 -> DSCO，所以这里做严格反向。
	lingSKUToDSCOPartner, err := reverseMapStrict(cfg.Mapping.SKU)
	if err != nil {
		return nil, err
	}
	return lingSKUToDSCOPartner, nil
}
