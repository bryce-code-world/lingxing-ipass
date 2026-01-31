package integration

import (
	"encoding/json"
	"fmt"
)

const logJSONMaxBytes = 32 * 1024

// JSONForLog 将对象转为 JSON 字符串（用于日志打印），并对超长内容做截断。
//
// 注意：
// - 对 json.RawMessage/[]byte：认为其本身就是 JSON 文本，直接转 string。
// - 对其它类型：使用 json.Marshal。
func JSONForLog(v any) string {
	switch t := v.(type) {
	case nil:
		return "null"
	case json.RawMessage:
		return truncateForLog(string(t), logJSONMaxBytes)
	case []byte:
		return truncateForLog(string(t), logJSONMaxBytes)
	case string:
		return truncateForLog(t, logJSONMaxBytes)
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return fmt.Sprintf("<json_marshal_failed:%v>", err)
		}
		return truncateForLog(string(b), logJSONMaxBytes)
	}
}

func truncateForLog(s string, maxBytes int) string {
	if maxBytes <= 0 || len(s) <= maxBytes {
		return s
	}
	return s[:maxBytes] + fmt.Sprintf("...(truncated,total_bytes=%d)", len(s))
}
