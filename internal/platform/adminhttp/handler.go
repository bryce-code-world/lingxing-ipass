package adminhttp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/internal/store"
)

// JobRunner 用于手动触发任务（同步执行一次）。
type JobRunner func(ctx context.Context) error

type Handler struct {
	watermark *store.WatermarkStore
	manual    *store.ManualTaskStore
	order     *store.OrderStateStore
	runners   map[string]JobRunner
}

func NewHandler(watermark *store.WatermarkStore, manual *store.ManualTaskStore, order *store.OrderStateStore, runners map[string]JobRunner) (*Handler, error) {
	if watermark == nil {
		return nil, errors.New("watermark 不能为空")
	}
	if manual == nil {
		return nil, errors.New("manual 不能为空")
	}
	if order == nil {
		return nil, errors.New("order 不能为空")
	}
	if runners == nil {
		runners = map[string]JobRunner{}
	}
	return &Handler{watermark: watermark, manual: manual, order: order, runners: runners}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := h.withTrace(r)
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/healthz":
		h.writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	case r.URL.Path == "/admin/run" && r.Method == http.MethodPost:
		h.handleRun(ctx, w, r)
		return
	case r.URL.Path == "/admin/order_state/get" && r.Method == http.MethodGet:
		h.handleOrderStateGet(ctx, w, r)
		return
	case r.URL.Path == "/admin/order_states" && r.Method == http.MethodGet:
		h.handleOrderStates(ctx, w, r)
		return
	case r.URL.Path == "/admin/watermark/get" && r.Method == http.MethodGet:
		h.handleWatermarkGet(ctx, w, r)
		return
	case r.URL.Path == "/admin/watermark/set" && r.Method == http.MethodPost:
		h.handleWatermarkSet(ctx, w, r)
		return
	case r.URL.Path == "/admin/manual_tasks" && r.Method == http.MethodGet:
		h.handleManualTasks(ctx, w, r)
		return
	default:
		h.writeJSON(w, http.StatusNotFound, map[string]any{"error": "not found"})
		return
	}
}

func (h *Handler) withTrace(r *http.Request) context.Context {
	ctx := r.Context()
	traceID := strings.TrimSpace(r.Header.Get("X-Trace-Id"))
	if traceID == "" {
		traceID = logger.NewTraceID()
	}
	return logger.WithTraceID(ctx, traceID)
}

func (h *Handler) handleRun(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	job := strings.TrimSpace(r.URL.Query().Get("job"))
	if job == "" {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "missing job"})
		return
	}
	fn, ok := h.runners[job]
	if !ok {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "unknown job"})
		return
	}

	start := time.Now()
	logger.Info(ctx, "admin_run_start", "job", job)
	err := fn(ctx)
	cost := time.Since(start).Milliseconds()
	if err != nil {
		logger.Error(ctx, "admin_run_end", "job", job, "cost_ms", cost, "err", err.Error())
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	logger.Info(ctx, "admin_run_end", "job", job, "cost_ms", cost)
	h.writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (h *Handler) handleWatermarkGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	job := strings.TrimSpace(r.URL.Query().Get("job"))
	if job == "" {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "missing job"})
		return
	}
	raw, ok, err := h.watermark.Get(ctx, job)
	if err != nil {
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	if !ok {
		h.writeJSON(w, http.StatusNotFound, map[string]any{"error": "not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(raw)
}

func (h *Handler) handleWatermarkSet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	job := strings.TrimSpace(r.URL.Query().Get("job"))
	if job == "" {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "missing job"})
		return
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "read body failed"})
		return
	}
	body = []byte(strings.TrimSpace(string(body)))
	if len(body) == 0 {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "empty body"})
		return
	}
	var tmp any
	if err := json.Unmarshal(body, &tmp); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid json"})
		return
	}

	if err := h.watermark.Set(ctx, job, body); err != nil {
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	logger.Info(ctx, "watermark_set", "job", job, "watermark", logger.MarshalAny(tmp))
	h.writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (h *Handler) handleOrderStateGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	dscoOrderID := strings.TrimSpace(r.URL.Query().Get("dsco_order_id"))
	if dscoOrderID == "" {
		dscoOrderID = strings.TrimSpace(r.URL.Query().Get("dscoOrderId"))
	}
	if dscoOrderID == "" {
		h.writeJSON(w, http.StatusBadRequest, map[string]any{"error": "missing dsco_order_id"})
		return
	}
	row, ok, err := h.order.GetByDscoOrderID(ctx, dscoOrderID)
	if err != nil {
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	if !ok {
		h.writeJSON(w, http.StatusNotFound, map[string]any{"error": "not found"})
		return
	}
	h.writeJSON(w, http.StatusOK, row)
}

func (h *Handler) handleOrderStates(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	limit := mustInt(r.URL.Query().Get("limit"), 50)
	offset := mustInt(r.URL.Query().Get("offset"), 0)

	pushedStatus, ok := optionalInt(r.URL.Query().Get("push_status"))
	var pushedPtr *int
	if ok {
		pushedPtr = &pushedStatus
	}
	ackedStatus, ok := optionalInt(r.URL.Query().Get("ack_status"))
	var ackedPtr *int
	if ok {
		ackedPtr = &ackedStatus
	}
	shippedStatus, ok := optionalInt(r.URL.Query().Get("ship_status"))
	var shippedPtr *int
	if ok {
		shippedPtr = &shippedStatus
	}
	invoicedStatus, ok := optionalInt(r.URL.Query().Get("invoice_status"))
	var invoicedPtr *int
	if ok {
		invoicedPtr = &invoicedStatus
	}

	items, err := h.order.List(ctx, store.OrderStateListQuery{
		PushedToLXStatus:     pushedPtr,
		AckedToDSCOStatus:    ackedPtr,
		ShippedToDSCOStatus:  shippedPtr,
		InvoicedToDSCOStatus: invoicedPtr,
		Limit:                limit,
		Offset:               offset,
	})
	if err != nil {
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusOK, map[string]any{"items": items, "count": len(items)})
}

func (h *Handler) handleManualTasks(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	status := mustInt(r.URL.Query().Get("status"), 0)
	limit := mustInt(r.URL.Query().Get("limit"), 50)
	offset := mustInt(r.URL.Query().Get("offset"), 0)

	items, err := h.manual.ListByStatus(ctx, status, limit, offset)
	if err != nil {
		h.writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusOK, map[string]any{"items": items, "count": len(items)})
}

func mustInt(raw string, def int) int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return def
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return def
	}
	return n
}

func optionalInt(raw string) (int, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0, false
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, false
	}
	return n, true
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
