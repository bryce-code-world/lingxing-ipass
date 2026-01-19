package adminweb

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/admin/store"
)

func (s *Server) apiRunJob(c *gin.Context) {
	job := strings.TrimSpace(c.Query("job"))
	if job == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing job"})
		return
	}

	start := time.Now()
	logger.Info(c.Request.Context(), "admin_run_start", "job", job)
	err := s.callOpsRun(c, job)
	cost := time.Since(start).Milliseconds()
	if err != nil {
		logger.Error(c.Request.Context(), "admin_run_end", "job", job, "cost_ms", cost, "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}
	logger.Info(c.Request.Context(), "admin_run_end", "job", job, "cost_ms", cost)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (s *Server) callOpsRun(c *gin.Context, job string) error {
	base := strings.TrimRight(strings.TrimSpace(s.opsBaseURL), "/")
	if base == "" {
		return errors.New("未配置 ADMIN_OPS_BASE_URL")
	}
	pass := strings.TrimSpace(s.opsPassword)
	if pass == "" {
		return errors.New("未配置 ADMIN_OPS_PASSWORD")
	}

	u := base + "/admin/run?job=" + url.QueryEscape(job)
	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodPost, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Ops-Password", pass)
	if traceID := strings.TrimSpace(c.Request.Header.Get("X-Trace-Id")); traceID != "" {
		req.Header.Set("X-Trace-Id", traceID)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
	body = bytes.TrimSpace(body)
	if len(body) == 0 {
		return errors.New("ops 调用失败: status=" + resp.Status)
	}
	return errors.New("ops 调用失败: status=" + resp.Status + ", body=" + string(body))
}

func (s *Server) apiWatermarkGet(c *gin.Context) {
	job := strings.TrimSpace(c.Query("job"))
	if job == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing job"})
		return
	}
	raw, ok, err := s.watermark.Get(c.Request.Context(), job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	_, _ = c.Writer.Write(raw)
}

func (s *Server) apiWatermarkSet(c *gin.Context) {
	job := strings.TrimSpace(c.Query("job"))
	if job == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing job"})
		return
	}
	body, err := mustJSONBody(c, 1<<20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.watermark.Set(c.Request.Context(), job, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info(c.Request.Context(), "watermark_set", "job", job, "watermark", string(body))
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (s *Server) apiManualTasks(c *gin.Context) {
	status := mustInt(c.Query("status"), 0, 0, 3)
	limit := mustInt(c.Query("limit"), 50, 1, 200)
	offset := mustInt(c.Query("offset"), 0, 0, 1<<30)

	items, err := s.manual.ListByStatus(c.Request.Context(), status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "count": len(items)})
}

func (s *Server) apiOrderStateGet(c *gin.Context) {
	dscoOrderID := strings.TrimSpace(c.Query("dsco_order_id"))
	if dscoOrderID == "" {
		dscoOrderID = strings.TrimSpace(c.Query("dscoOrderId"))
	}
	if dscoOrderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing dsco_order_id"})
		return
	}
	row, ok, err := s.order.GetByDscoOrderID(c.Request.Context(), dscoOrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, row)
}

func (s *Server) apiOrderStates(c *gin.Context) {
	limit := mustInt(c.Query("limit"), 50, 1, 200)
	offset := mustInt(c.Query("offset"), 0, 0, 1<<30)

	q := store.OrderStateListQuery{
		PushedToLXStatus:     optionalIntPtr(c.Query("push_status")),
		AckedToDSCOStatus:    optionalIntPtr(c.Query("ack_status")),
		ShippedToDSCOStatus:  optionalIntPtr(c.Query("ship_status")),
		InvoicedToDSCOStatus: optionalIntPtr(c.Query("invoice_status")),
		Limit:                limit,
		Offset:               offset,
	}
	items, err := s.order.List(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "count": len(items)})
}
