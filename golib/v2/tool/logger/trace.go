package logger

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

type ctxKeyTraceID struct{}

// NewTraceID 生成一个新的 trace_id（32 位 hex）。
func NewTraceID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

// WithTraceID 将 trace_id 写入 context，贯穿整个执行链路。
func WithTraceID(ctx context.Context, traceID string) context.Context {
	traceID = strings.TrimSpace(traceID)
	if traceID == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxKeyTraceID{}, traceID)
}

// TraceIDFromContext 从 context 读取 trace_id。
func TraceIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(ctxKeyTraceID{})
	s, ok := v.(string)
	s = strings.TrimSpace(s)
	if !ok || s == "" {
		return "", false
	}
	return s, true
}

// EnsureTraceID 确保 ctx 中存在 trace_id；如果没有则生成并写入。
func EnsureTraceID(ctx context.Context) (context.Context, string) {
	if tid, ok := TraceIDFromContext(ctx); ok {
		return ctx, tid
	}
	tid := NewTraceID()
	return WithTraceID(ctx, tid), tid
}
