package dsco

import (
	"fmt"
)

// APIError 表示 HTTP 非 2xx 响应。
type APIError struct {
	StatusCode int
	Method     string
	URL        string
	Body       []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if len(e.Body) == 0 {
		return fmt.Sprintf("Dsco API 请求失败: %s %s, status=%d", e.Method, e.URL, e.StatusCode)
	}
	return fmt.Sprintf("Dsco API 请求失败: %s %s, status=%d, body=%s", e.Method, e.URL, e.StatusCode, string(e.Body))
}
