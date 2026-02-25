package lingxing

import "fmt"

// APIError 表示领星 API 请求失败（包括 HTTP 非2xx 或业务 code 非成功）。
type APIError struct {
	StatusCode int
	Method     string
	URL        string

	Code      Code
	Message   string
	RequestID string

	Body []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Code.Raw() != "" || e.Message != "" || e.RequestID != "" {
		return fmt.Sprintf("领星 API 请求失败: %s %s, http_status=%d, code=%s, msg=%s, request_id=%s",
			e.Method, e.URL, e.StatusCode, e.Code.Raw(), e.Message, e.RequestID)
	}
	if len(e.Body) == 0 {
		return fmt.Sprintf("领星 API 请求失败: %s %s, http_status=%d", e.Method, e.URL, e.StatusCode)
	}
	return fmt.Sprintf("领星 API 请求失败: %s %s, http_status=%d, body=%s", e.Method, e.URL, e.StatusCode, string(e.Body))
}
