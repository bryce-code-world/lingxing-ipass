package opshttp

import (
	"net/http"
	"strings"
)

// Wrap 用于给业务侧 ops 接口加一层最小鉴权：
// - /healthz 放行（便于探活）
// - 其他路径必须携带请求头 X-Ops-Password 且等于配置密码
func Wrap(next http.Handler, password string) http.Handler {
	password = strings.TrimSpace(password)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}
		if password == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"error":"missing ops password"}`))
			return
		}
		if strings.TrimSpace(r.Header.Get("X-Ops-Password")) != password {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
