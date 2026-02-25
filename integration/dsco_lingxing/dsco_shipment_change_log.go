package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
)

type dscoShipmentChangeLogPayload struct {
	PoNumber string `json:"poNumber"`
}

func pollDSCOShipmentChangeLog(
	ctx context.Context,
	cli *dsco.Client,
	requestID string,
	want map[string]struct{},
	maxWait time.Duration,
	interval time.Duration,
) (map[string]bool, map[string]any, bool) {
	requestID = strings.TrimSpace(requestID)
	if cli == nil || requestID == "" || len(want) == 0 {
		return map[string]bool{}, map[string]any{}, true
	}
	if maxWait <= 0 {
		maxWait = 10 * time.Second
	}
	if interval <= 0 {
		interval = 1 * time.Second
	}

	deadline := time.Now().Add(maxWait)
	result := make(map[string]bool, len(want))
	detail := make(map[string]any, len(want))

	for {
		out, err := cli.Order.GetChangeLog(ctx, dsco.OrderChangeLogQuery{
			RequestID: requestID,
			Status:    "success_or_failure",
		})
		if err == nil && out != nil {
			for _, l := range out.Logs {
				po := extractPONumberFromDSCOChangeLogPayload(l.Payload)
				if po == "" {
					continue
				}
				if _, ok := want[po]; !ok {
					continue
				}
				st := strings.ToLower(strings.TrimSpace(l.Status))
				switch st {
				case "success":
					// 若同一 PO 出现多条日志（极少见），失败优先于成功，避免误推进状态。
					if prev, ok := result[po]; ok && !prev {
						continue
					}
					result[po] = true
				case "failure":
					result[po] = false
					detail[po] = l
				default:
					// pending/unknown：等待下一轮
				}
			}

			if strings.EqualFold(strings.TrimSpace(out.Status), "COMPLETED") {
				return result, detail, true
			}
			allDone := true
			for po := range want {
				if _, ok := result[po]; !ok {
					allDone = false
					break
				}
			}
			if allDone {
				return result, detail, true
			}
		}

		if time.Now().After(deadline) {
			return result, detail, false
		}
		time.Sleep(interval)
	}
}

func extractPONumberFromDSCOChangeLogPayload(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var p dscoShipmentChangeLogPayload
	if err := json.Unmarshal(raw, &p); err == nil {
		if s := strings.TrimSpace(p.PoNumber); s != "" {
			return s
		}
	}
	// 兜底：防止 DSCO 返回 payload 结构变化导致无法提取 poNumber。
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return ""
	}
	if v, ok := m["poNumber"]; ok {
		if s, ok := v.(string); ok {
			return strings.TrimSpace(s)
		}
	}
	return ""
}
