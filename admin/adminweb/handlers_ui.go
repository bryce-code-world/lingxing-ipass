package adminweb

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/internal/store"
)

func (s *Server) uiLoginGet(c *gin.Context) {
	next := strings.TrimSpace(c.Query("next"))
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
		"next":  next,
	})
}

func (s *Server) uiLoginPost(c *gin.Context) {
	pass := strings.TrimSpace(c.PostForm("password"))
	if pass == "" || pass != strings.TrimSpace(s.getAdminPassword()) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "登录",
			"error": "密码错误",
			"next":  strings.TrimSpace(c.Query("next")),
		})
		return
	}
	if err := s.setSessionCookie(c); err != nil {
		logger.Error(c.Request.Context(), "admin_login_failed", "err", err.Error())
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "登录",
			"error": "登录失败",
		})
		return
	}
	next := strings.TrimSpace(c.Query("next"))
	if next == "" {
		next = "/admin/ui/"
	}
	c.Redirect(http.StatusFound, next)
}

func (s *Server) uiLogoutPost(c *gin.Context) {
	s.clearSessionCookie(c)
	c.Redirect(http.StatusFound, "/admin/ui/login")
}

func (s *Server) uiDashboard(c *gin.Context) {
	rows, err := s.watermark.ListAll(c.Request.Context())
	if err != nil {
		logger.Error(c.Request.Context(), "admin_dashboard_watermark_failed", "err", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "错误", "error": err.Error()})
		return
	}
	jobs := make([]string, 0, len(s.runners))
	for k := range s.runners {
		jobs = append(jobs, k)
	}
	sort.Strings(jobs)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title":      "总览",
		"watermarks": rows,
		"jobs":       jobs,
	})
}

func (s *Server) uiOrders(c *gin.Context) {
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
		logger.Error(c.Request.Context(), "admin_orders_failed", "err", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "错误", "error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "orders.html", gin.H{
		"title": "订单列表",
		"items": items,
		"q":     q,
	})
}

func (s *Server) uiOrderDetail(c *gin.Context) {
	dscoOrderID := strings.TrimSpace(c.Query("dsco_order_id"))
	if dscoOrderID == "" {
		dscoOrderID = strings.TrimSpace(c.Query("dscoOrderId"))
	}
	if dscoOrderID == "" {
		c.HTML(http.StatusOK, "order_detail.html", gin.H{
			"title": "订单详情",
			"error": "缺少 dsco_order_id",
		})
		return
	}
	row, ok, err := s.order.GetByDscoOrderID(c.Request.Context(), dscoOrderID)
	if err != nil {
		logger.Error(c.Request.Context(), "admin_order_get_failed", "dsco_order_id", dscoOrderID, "err", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "错误", "error": err.Error()})
		return
	}
	if !ok {
		c.HTML(http.StatusOK, "order_detail.html", gin.H{
			"title":         "订单详情",
			"dsco_order_id": dscoOrderID,
			"error":         "未找到该订单",
		})
		return
	}
	c.HTML(http.StatusOK, "order_detail.html", gin.H{
		"title": "订单详情",
		"item":  row,
	})
}

func (s *Server) uiManualTasks(c *gin.Context) {
	status := mustInt(c.Query("status"), 0, 0, 3)
	limit := mustInt(c.Query("limit"), 50, 1, 200)
	offset := mustInt(c.Query("offset"), 0, 0, 1<<30)

	items, err := s.manual.ListByStatus(c.Request.Context(), status, limit, offset)
	if err != nil {
		logger.Error(c.Request.Context(), "admin_manual_tasks_failed", "err", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "错误", "error": err.Error()})
		return
	}

	type manualTaskUIItem struct {
		ID          int64
		TaskType    string
		DscoOrderID string
		Status      int
		UpdatedAt   time.Time
	}
	viewItems := make([]manualTaskUIItem, 0, len(items))
	for _, it := range items {
		var dscoOrderID string
		if it.DscoOrderID != nil {
			dscoOrderID = *it.DscoOrderID
		}
		viewItems = append(viewItems, manualTaskUIItem{
			ID:          it.ID,
			TaskType:    it.TaskType,
			DscoOrderID: dscoOrderID,
			Status:      it.Status,
			UpdatedAt:   it.UpdatedAt,
		})
	}

	c.HTML(http.StatusOK, "manual_tasks.html", gin.H{
		"title":  "人工任务",
		"items":  viewItems,
		"status": status,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) uiWatermarks(c *gin.Context) {
	rows, err := s.watermark.ListAll(c.Request.Context())
	if err != nil {
		logger.Error(c.Request.Context(), "admin_watermarks_failed", "err", err.Error())
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "错误", "error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "watermarks.html", gin.H{
		"title":     "水位管理",
		"items":     rows,
		"selected":  strings.TrimSpace(c.Query("job")),
		"watermark": strings.TrimSpace(c.Query("watermark")),
	})
}

func (s *Server) uiJobs(c *gin.Context) {
	jobs := make([]string, 0, len(s.runners))
	for k := range s.runners {
		jobs = append(jobs, k)
	}
	sort.Strings(jobs)
	c.HTML(http.StatusOK, "jobs.html", gin.H{
		"title": "任务执行",
		"jobs":  jobs,
	})
}

func mustInt(raw string, def, min, max int) int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return def
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return def
	}
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

func optionalIntPtr(raw string) *int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return nil
	}
	return &n
}

