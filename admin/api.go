package admin

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"lingxingipass/infra/runtimecfg"
	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

type apiResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, apiResp{Code: 0, Message: "ok", Data: data})
}

func fail(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, apiResp{Code: code, Message: msg})
}

func (s *Server) apiLogin(c *gin.Context) {
	pass := strings.TrimSpace(c.PostForm("password"))
	if pass == "" {
		var body struct {
			Password string `json:"password"`
		}
		_ = c.ShouldBindJSON(&body)
		pass = strings.TrimSpace(body.Password)
	}
	if pass == "" || pass != s.env.Admin.Password {
		fail(c, http.StatusUnauthorized, 401, "invalid password")
		return
	}
	s.setSession(c)
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiLogout(c *gin.Context) {
	s.clearSession(c)
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiMe(c *gin.Context) {
	cookie, _ := c.Cookie(sessionCookieName)
	ok(c, map[string]any{"authed": cookie == s.sessionToken()})
}

func (s *Server) apiGetRuntimeConfig(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		domain = runtimecfg.DomainDSCOLingXing
	}
	rc, okk := s.cfgMgr.Snapshot(domain)
	if !okk {
		fail(c, http.StatusServiceUnavailable, 503, "runtime config not loaded")
		return
	}
	ok(c, map[string]any{
		"domain":     domain,
		"updated_at": rc.UpdatedAt,
		"config":     rc.Config,
	})
}

func (s *Server) apiUpdateRuntimeConfig(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		domain = runtimecfg.DomainDSCOLingXing
	}

	var body runtimecfg.Config
	if err := c.ShouldBindJSON(&body); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	if body.Domain == "" {
		body.Domain = domain
	}
	if err := s.cfgMgr.Update(c.Request.Context(), domain, body, s.reg.SupportedJobsSet()); err != nil {
		fail(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiRunJobs(c *gin.Context) {
	var body struct {
		Jobs []runtimecfg.JobName `json:"jobs"`
	}
	_ = c.ShouldBindJSON(&body)
	jobs := body.Jobs
	if len(jobs) == 0 {
		jobs = s.reg.Jobs()
	}
	var results []map[string]any
	for _, j := range jobs {
		err := s.runner.Run(c.Request.Context(), integration.RunRequest{
			Domain:  runtimecfg.DomainDSCOLingXing,
			Job:     j,
			Trigger: integration.TriggerManual,
		})
		results = append(results, map[string]any{"job": j, "error": errString(err)})
	}
	ok(c, map[string]any{"results": results})
}

func (s *Server) apiRunOneJob(c *gin.Context) {
	var body struct {
		Job runtimecfg.JobName `json:"job"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Job == "" {
		fail(c, http.StatusBadRequest, 400, "missing job")
		return
	}
	err := s.runner.Run(c.Request.Context(), integration.RunRequest{
		Domain:  runtimecfg.DomainDSCOLingXing,
		Job:     body.Job,
		Trigger: integration.TriggerManual,
	})
	if err != nil {
		if errors.Is(err, integration.ErrJobRunning) {
			fail(c, http.StatusConflict, 409, "job running")
			return
		}
		fail(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiListOrders(c *gin.Context) {
	filter := store.DSCOOrderSyncListFilter{
		Offset:                  parseInt(c.Query("offset"), 0),
		Limit:                   parseInt(c.Query("limit"), 50),
		PONumberLike:            c.Query("po_number"),
		DSCOOrderID:             c.Query("dsco_order_id"),
		ConsumerOrderNumberLike: c.Query("consumer_order_number"),
		Channel:                 c.Query("channel"),
		MSKU:                    c.Query("msku"),
	}
	if v := c.Query("start"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			filter.StartTime = &n
		}
	}
	if v := c.Query("end"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			filter.EndTime = &n
		}
	}
	if v := c.Query("status"); v != "" {
		parts := strings.Split(v, ",")
		for _, p := range parts {
			if n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 16); err == nil {
				filter.StatusIn = append(filter.StatusIn, int16(n))
			}
		}
	}
	items, total, err := s.orderStore.List(c.Request.Context(), filter)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	ok(c, map[string]any{"items": items, "total": total})
}

func (s *Server) apiOrderDetail(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		fail(c, http.StatusBadRequest, 400, "missing id")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid id")
		return
	}
	// Simple: query by primary key.
	row, ok2, err := s.orderStore.GetByID(c.Request.Context(), id)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	if !ok2 {
		fail(c, http.StatusNotFound, 404, "not found")
		return
	}
	ok(c, row)
}

func (s *Server) apiManualPullOrders(c *gin.Context) {
	var body struct {
		Start  int64 `json:"start"`
		End    int64 `json:"end"`
		Status int16 `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	if body.Start <= 0 || body.End <= 0 || body.End <= body.Start {
		fail(c, http.StatusBadRequest, 400, "invalid time range")
		return
	}
	if body.Status < 1 || body.Status > 5 {
		fail(c, http.StatusBadRequest, 400, "invalid status")
		return
	}
	maxSec := int64(s.env.Admin.Export.MaxRangeDays) * 86400
	if maxSec > 0 && body.End-body.Start > maxSec {
		fail(c, http.StatusBadRequest, 400, "range too large")
		return
	}

	runCtx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Minute)
	defer cancel()

	err := s.runner.Run(runCtx, integration.RunRequest{
		Domain:   runtimecfg.DomainDSCOLingXing,
		Job:      runtimecfg.JobPullDSCOOrders,
		Trigger:  integration.TriggerManual,
		Size:     0,
		Override: integration.PullDSCOOrdersOverride{Start: body.Start, End: body.End, Status: body.Status},
	})
	if err != nil {
		switch {
		case errors.Is(err, integration.ErrJobRunning):
			fail(c, http.StatusConflict, 409, err.Error())
		case errors.Is(err, integration.ErrJobDisabled):
			fail(c, http.StatusBadRequest, 400, err.Error())
		default:
			fail(c, http.StatusInternalServerError, 500, err.Error())
		}
		return
	}
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiListWarehouseSync(c *gin.Context) {
	filter := store.DSCOWarehouseSyncListFilter{
		Offset:               parseInt(c.Query("offset"), 0),
		Limit:                parseInt(c.Query("limit"), 50),
		DSCOWarehouseID:      c.Query("dsco_warehouse_id"),
		DSCOWarehouseSKU:     c.Query("dsco_warehouse_sku"),
		LingXingWarehouseID:  c.Query("lingxing_warehouse_id"),
		LingXingWarehouseSKU: c.Query("lingxing_warehouse_sku"),
	}
	if v := c.Query("start"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			filter.StartTime = &n
		}
	}
	if v := c.Query("end"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			filter.EndTime = &n
		}
	}
	if v := c.Query("status"); v != "" {
		parts := strings.Split(v, ",")
		for _, p := range parts {
			if n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 16); err == nil {
				filter.StatusIn = append(filter.StatusIn, int16(n))
			}
		}
	}
	items, total, err := s.warehouseStore.List(c.Request.Context(), filter)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	ok(c, map[string]any{"items": items, "total": total})
}

func (s *Server) apiExportOrders(c *gin.Context) {
	// For MVP: export last query via request body filters (same as list subset).
	var filter store.DSCOOrderSyncListFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		// allow empty body -> export all
	}
	// enforce max range days
	if filter.StartTime != nil && filter.EndTime != nil {
		maxSec := int64(s.env.Admin.Export.MaxRangeDays) * 86400
		if *filter.EndTime-*filter.StartTime > maxSec {
			fail(c, http.StatusBadRequest, 400, "range too large")
			return
		}
	}
	filter.Offset = 0
	filter.Limit = 500

	dir := s.env.Admin.Export.Dir
	_ = os.MkdirAll(dir, 0755)
	tmp, err := os.CreateTemp(dir, "dsco_order_sync_*.csv.tmp")
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	defer tmp.Close()

	w := csv.NewWriter(tmp)
	_ = w.Write([]string{"id", "po_number", "dsco_create_time", "status", "warehouse_id", "shipment", "shipped_tracking_no", "dsco_invoice_id"})

	offset := 0
	for {
		filter.Offset = offset
		items, total, err := s.orderStore.List(c.Request.Context(), filter)
		if err != nil {
			fail(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		for _, it := range items {
			_ = w.Write([]string{
				strconv.FormatInt(it.ID, 10),
				it.PONumber,
				strconv.FormatInt(it.DSCOCreateTime, 10),
				strconv.FormatInt(int64(it.Status), 10),
				it.WarehouseID,
				it.Shipment,
				it.ShippedTrackingNo,
				it.DSCOInvoiceID,
			})
		}
		offset += len(items)
		if offset >= int(total) || len(items) == 0 {
			break
		}
	}
	w.Flush()
	_ = tmp.Sync()

	finalName := filepath.Join(dir, "dsco_order_sync_"+time.Now().UTC().Format("20060102_150405")+".csv")
	_ = os.Rename(tmp.Name(), finalName)
	c.FileAttachment(finalName, filepath.Base(finalName))
}

func (s *Server) apiExportWarehouseSync(c *gin.Context) {
	var filter store.DSCOWarehouseSyncListFilter
	_ = c.ShouldBindJSON(&filter)
	if filter.StartTime != nil && filter.EndTime != nil {
		maxSec := int64(s.env.Admin.Export.MaxRangeDays) * 86400
		if *filter.EndTime-*filter.StartTime > maxSec {
			fail(c, http.StatusBadRequest, 400, "range too large")
			return
		}
	}
	filter.Offset = 0
	filter.Limit = 500

	dir := s.env.Admin.Export.Dir
	_ = os.MkdirAll(dir, 0755)
	tmp, err := os.CreateTemp(dir, "dsco_warehouse_sync_*.csv.tmp")
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	defer tmp.Close()

	w := csv.NewWriter(tmp)
	_ = w.Write([]string{"id", "sync_time", "dsco_warehouse_id", "dsco_warehouse_sku", "dsco_warehouse_num", "lingxing_warehouse_id", "lingxing_warehouse_sku", "lingxing_warehouse_num", "status", "reason"})

	offset := 0
	for {
		filter.Offset = offset
		items, total, err := s.warehouseStore.List(c.Request.Context(), filter)
		if err != nil {
			fail(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		for _, it := range items {
			_ = w.Write([]string{
				strconv.FormatInt(it.ID, 10),
				strconv.FormatInt(it.SyncTime, 10),
				it.DSCOWarehouseID,
				it.DSCOWarehouseSKU,
				strconv.Itoa(it.DSCOWarehouseNum),
				it.LingXingWarehouseID,
				it.LingXingWarehouseSKU,
				strconv.Itoa(it.LingXingWarehouseNum),
				strconv.FormatInt(int64(it.Status), 10),
				it.Reason,
			})
		}
		offset += len(items)
		if offset >= int(total) || len(items) == 0 {
			break
		}
	}
	w.Flush()
	_ = tmp.Sync()

	finalName := filepath.Join(dir, "dsco_warehouse_sync_"+time.Now().UTC().Format("20060102_150405")+".csv")
	_ = os.Rename(tmp.Name(), finalName)
	c.FileAttachment(finalName, filepath.Base(finalName))
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func parseInt(s string, def int) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}

// Keep json import used (payload debug).
var _ = json.RawMessage{}
