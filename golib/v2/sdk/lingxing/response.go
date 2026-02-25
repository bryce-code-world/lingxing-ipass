package lingxing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Code 表示领星接口返回的 code 字段。
//
// 注意：文档与示例中 code 既可能是 number，也可能是 string（且可能带前导0），所以用 Raw 字符串保存。
type Code struct {
	raw string
}

func (c Code) Raw() string { return c.raw }

func (c Code) IsSuccess() bool {
	switch strings.TrimSpace(c.raw) {
	case "0", "200":
		return true
	default:
		return false
	}
}

func (c *Code) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		c.raw = ""
		return nil
	}
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		c.raw = s
		return nil
	}
	// 数字等情况，直接保留原始字符串；若是 0.0 之类，按整型收敛为 "0"。
	s := string(b)
	if strings.Contains(s, ".") {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			c.raw = strconv.FormatInt(int64(f), 10)
			return nil
		}
	}
	c.raw = s
	return nil
}

// ResponseEnvelope 是领星接口的通用响应壳，兼容不同接口的字段命名差异。
type ResponseEnvelope struct {
	Code Code

	// 消息字段：有的接口使用 msg，有的使用 message。
	Msg         string
	MessageText string

	// 链路字段：request_id 或 requestId。
	RequestIDSnake string
	RequestIDCamel string

	// 错误详情：error_details 或 errorDetails。
	ErrorDetailsSnake json.RawMessage
	ErrorDetailsCamel json.RawMessage

	Data  json.RawMessage
	Total json.RawMessage
}

func (e ResponseEnvelope) IsSuccess() bool { return e.Code.IsSuccess() }

func (e ResponseEnvelope) Message() string {
	if strings.TrimSpace(e.Msg) != "" {
		return e.Msg
	}
	return e.MessageText
}

func (e ResponseEnvelope) RequestID() string {
	if strings.TrimSpace(e.RequestIDSnake) != "" {
		return e.RequestIDSnake
	}
	return e.RequestIDCamel
}

func (e *ResponseEnvelope) UnmarshalJSON(b []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	decode := func(key string, out any) error {
		raw, ok := m[key]
		if !ok {
			return nil
		}
		return json.Unmarshal(raw, out)
	}

	if err := decode("code", &e.Code); err != nil {
		return fmt.Errorf("解析 code 失败: %w", err)
	}
	_ = decode("msg", &e.Msg)
	_ = decode("message", &e.MessageText)
	_ = decode("request_id", &e.RequestIDSnake)
	_ = decode("requestId", &e.RequestIDCamel)
	_ = decode("error_details", &e.ErrorDetailsSnake)
	_ = decode("errorDetails", &e.ErrorDetailsCamel)
	_ = decode("data", &e.Data)
	_ = decode("total", &e.Total)
	return nil
}

func parseIntFromRaw(raw json.RawMessage) (int, bool) {
	if len(raw) == 0 {
		return 0, false
	}
	b := bytes.TrimSpace(raw)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		return 0, false
	}
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return 0, false
		}
		s = strings.TrimSpace(s)
		if s == "" {
			return 0, false
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, false
		}
		return n, true
	}
	var n int
	if err := json.Unmarshal(b, &n); err == nil {
		return n, true
	}
	var f float64
	if err := json.Unmarshal(b, &f); err == nil {
		return int(f), true
	}
	return 0, false
}
