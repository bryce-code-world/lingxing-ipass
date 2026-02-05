package admin

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
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

func loadDisplayLocation(displayTimezone string) *time.Location {
	tz := strings.TrimSpace(displayTimezone)
	if tz == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.UTC
	}
	return loc
}

func formatUnixSecToDisplay(sec int64, loc *time.Location) string {
	if sec <= 0 {
		return ""
	}
	if loc == nil {
		loc = time.UTC
	}
	return time.Unix(sec, 0).In(loc).Format("2006-01-02 15:04:05.000")
}

func copyFile(srcPath string, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func() { _ = dst.Close() }()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return dst.Sync()
}

func writeCSVExportFile(dir string, baseName string, write func(w *csv.Writer) error) (string, error) {
	if strings.TrimSpace(dir) == "" {
		return "", errors.New("export dir is empty")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	tmp, err := os.CreateTemp(dir, baseName+"_*.csv.tmp")
	if err != nil {
		return "", err
	}
	tmpName := tmp.Name()

	w := csv.NewWriter(tmp)
	if err := write(w); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
		return "", err
	}
	w.Flush()
	if err := w.Error(); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
		return "", err
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
		return "", err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpName)
		return "", err
	}

	nano := time.Now().UTC().UnixNano()
	finalName := filepath.Join(dir, baseName+"_"+time.Now().UTC().Format("20060102_150405")+"_"+strconv.FormatInt(nano, 10)+".csv")

	if err := os.Rename(tmpName, finalName); err != nil {
		if err2 := copyFile(tmpName, finalName); err2 != nil {
			_ = os.Remove(tmpName)
			return "", err
		}
		_ = os.Remove(tmpName)
	}
	return finalName, nil
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
	type orderItem struct {
		ID                int64    `json:"id"`
		PONumber          string   `json:"po_number"`
		DSCOCreateTime    int64    `json:"dsco_create_time"`
		CreatedAt         int64    `json:"created_at"`
		UpdatedAt         int64    `json:"updated_at"`
		DSCOStatus        string   `json:"dsco_status"`
		Status            int16    `json:"status"`
		MSKUs             []string `json:"mskus"`
		WarehouseID       string   `json:"warehouse_id"`
		Shipment          string   `json:"shipment"`
		DSCOREtailerID    string   `json:"dsco_retailer_id"`
		ShippedTrackingNo string   `json:"shipped_tracking_no"`
		DSCOInvoiceID     string   `json:"dsco_invoice_id"`
	}

	filter := store.DSCOOrderSyncListFilter{
		Offset:         parseInt(c.Query("offset"), 0),
		Limit:          parseInt(c.Query("limit"), 50),
		DSCOStatus:     c.Query("dsco_status"),
		PONumberLike:   c.Query("po_number"),
		DSCOREtailerID: c.Query("dsco_retailer_id"),
		MSKU:           c.Query("msku"),
		WarehouseID:    c.Query("warehouse_id"),
		Shipment:       c.Query("shipment"),
		Tracking:       c.Query("tracking"),
		InvoiceID:      c.Query("invoice_id"),
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
	out := make([]orderItem, 0, len(items))
	for _, it := range items {
		out = append(out, orderItem{
			ID:                it.ID,
			PONumber:          it.PONumber,
			DSCOCreateTime:    it.DSCOCreateTime,
			CreatedAt:         it.CreatedAt,
			UpdatedAt:         it.UpdatedAt,
			DSCOStatus:        it.DSCOStatus,
			Status:            it.Status,
			MSKUs:             []string(it.MSKUs),
			WarehouseID:       it.WarehouseID,
			Shipment:          it.Shipment,
			DSCOREtailerID:    it.DSCOREtailerID,
			ShippedTrackingNo: it.ShippedTrackingNo,
			DSCOInvoiceID:     it.DSCOInvoiceID,
		})
	}
	ok(c, map[string]any{"items": out, "total": total})
}

func (s *Server) apiUpdateOrderStatus(c *gin.Context) {
	var body struct {
		PONumber string `json:"po_number"`
		Status   int16  `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	if strings.TrimSpace(body.PONumber) == "" {
		fail(c, http.StatusBadRequest, 400, "missing po_number")
		return
	}
	if body.Status < 1 || body.Status > 6 {
		fail(c, http.StatusBadRequest, 400, "invalid status")
		return
	}
	if err := s.orderStore.UpdateStatus(c.Request.Context(), body.PONumber, body.Status); err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	ok(c, map[string]any{"ok": true})
}

func (s *Server) apiRunOneOrderByStatus(c *gin.Context) {
	var body struct {
		PONumber string `json:"po_number"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || strings.TrimSpace(body.PONumber) == "" {
		fail(c, http.StatusBadRequest, 400, "missing po_number")
		return
	}
	po := strings.TrimSpace(body.PONumber)

	before, okk, err := s.orderStore.GetByPONumber(c.Request.Context(), po)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	if !okk {
		fail(c, http.StatusNotFound, 404, "order not found")
		return
	}

	// status=5/6（已完成/取消）禁用手动执行。
	if before.Status == 5 || before.Status == 6 {
		fail(c, http.StatusBadRequest, 400, "status not runnable")
		return
	}

	var job runtimecfg.JobName
	switch before.Status {
	case 1:
		job = runtimecfg.JobPushToLingXing
	case 2:
		job = runtimecfg.JobAckToDSCO
	case 3:
		job = runtimecfg.JobShipToDSCO
	case 4:
		job = runtimecfg.JobInvoiceToDSCO
	default:
		fail(c, http.StatusBadRequest, 400, "unsupported status")
		return
	}

	err = s.runner.Run(c.Request.Context(), integration.RunRequest{
		Domain:       runtimecfg.DomainDSCOLingXing,
		Job:          job,
		Trigger:      integration.TriggerManual,
		Size:         1,
		OnlyPONumber: po,
	})
	if err != nil {
		if errors.Is(err, integration.ErrJobRunning) {
			fail(c, http.StatusConflict, 409, "job running")
			return
		}
		fail(c, http.StatusBadRequest, 400, err.Error())
		return
	}

	afterStatus := before.Status
	if after, okk, err := s.orderStore.GetByPONumber(c.Request.Context(), po); err == nil && okk {
		afterStatus = after.Status
	}
	ok(c, map[string]any{
		"po_number":      po,
		"job":            job,
		"status_before":  before.Status,
		"status_after":   afterStatus,
		"status_changed": afterStatus != before.Status,
	})
}

func (s *Server) apiCheckOrders(c *gin.Context) {
	var body struct {
		Start    int64  `json:"start"`
		End      int64  `json:"end"`
		PONumber string `json:"po_number"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	po := strings.TrimSpace(body.PONumber)
	if po == "" {
		if body.Start <= 0 || body.End <= 0 || body.End <= body.Start {
			fail(c, http.StatusBadRequest, 400, "missing or invalid time range")
			return
		}
	}

	override := integration.CheckOrdersOverride{Start: body.Start, End: body.End}
	res, err := s.runner.RunWithResult(c.Request.Context(), integration.RunRequest{
		Domain:       runtimecfg.DomainDSCOLingXing,
		Job:          runtimecfg.JobCheckOrders,
		Trigger:      integration.TriggerManual,
		OnlyPONumber: po,
		Override:     override,
	})
	if err != nil {
		if errors.Is(err, integration.ErrJobRunning) {
			fail(c, http.StatusConflict, 409, "job running")
			return
		}
		fail(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if res == nil {
		fail(c, http.StatusInternalServerError, 500, "task returned empty result")
		return
	}
	ok(c, res)
}

func (s *Server) apiOrderDetail(c *gin.Context) {
	type orderDetail struct {
		ID                int64           `json:"id"`
		PONumber          string          `json:"po_number"`
		DSCOCreateTime    int64           `json:"dsco_create_time"`
		DSCOStatus        string          `json:"dsco_status"`
		Status            int16           `json:"status"`
		MSKUs             []string        `json:"mskus"`
		WarehouseID       string          `json:"warehouse_id"`
		Shipment          string          `json:"shipment"`
		DSCOREtailerID    string          `json:"dsco_retailer_id"`
		ShippedTrackingNo string          `json:"shipped_tracking_no"`
		DSCOInvoiceID     string          `json:"dsco_invoice_id"`
		Payload           json.RawMessage `json:"payload"`
		CreatedAt         int64           `json:"created_at"`
		UpdatedAt         int64           `json:"updated_at"`
	}

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
	ok(c, orderDetail{
		ID:                row.ID,
		PONumber:          row.PONumber,
		DSCOCreateTime:    row.DSCOCreateTime,
		DSCOStatus:        row.DSCOStatus,
		Status:            row.Status,
		MSKUs:             []string(row.MSKUs),
		WarehouseID:       row.WarehouseID,
		Shipment:          row.Shipment,
		DSCOREtailerID:    row.DSCOREtailerID,
		ShippedTrackingNo: row.ShippedTrackingNo,
		DSCOInvoiceID:     row.DSCOInvoiceID,
		Payload:           row.Payload,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	})
}

func (s *Server) apiManualPullOrders(c *gin.Context) {
	var body struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	if body.Start <= 0 || body.End <= 0 || body.End <= body.Start {
		fail(c, http.StatusBadRequest, 400, "invalid time range")
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
		Override: integration.PullDSCOOrdersOverride{Start: body.Start, End: body.End},
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
	type warehouseItem struct {
		ID                   int64  `json:"id"`
		SyncTime             int64  `json:"sync_time"`
		DSCOWarehouseID      string `json:"dsco_warehouse_id"`
		DSCOWarehouseSKU     string `json:"dsco_warehouse_sku"`
		DSCOWarehouseNum     int    `json:"dsco_warehouse_num"`
		LingXingWarehouseID  string `json:"lingxing_warehouse_id"`
		LingXingWarehouseSKU string `json:"lingxing_warehouse_sku"`
		LingXingWarehouseNum int    `json:"lingxing_warehouse_num"`
		Diff                 int    `json:"diff"`
		Status               int16  `json:"status"`
		Reason               string `json:"reason"`
	}

	filter := store.DSCOWarehouseSyncListFilter{
		Offset:               parseInt(c.Query("offset"), 0),
		Limit:                parseInt(c.Query("limit"), 50),
		DSCOWarehouseID:      c.Query("dsco_warehouse_id"),
		DSCOWarehouseSKU:     c.Query("dsco_warehouse_sku"),
		LingXingWarehouseID:  c.Query("lingxing_warehouse_id"),
		LingXingWarehouseSKU: c.Query("lingxing_warehouse_sku"),
	}
	applyWarehouseDiffRange(&filter, c.Query("diff_range"))
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
	out := make([]warehouseItem, 0, len(items))
	for _, it := range items {
		out = append(out, warehouseItem{
			ID:                   it.ID,
			SyncTime:             it.SyncTime,
			DSCOWarehouseID:      it.DSCOWarehouseID,
			DSCOWarehouseSKU:     it.DSCOWarehouseSKU,
			DSCOWarehouseNum:     it.DSCOWarehouseNum,
			LingXingWarehouseID:  it.LingXingWarehouseID,
			LingXingWarehouseSKU: it.LingXingWarehouseSKU,
			LingXingWarehouseNum: it.LingXingWarehouseNum,
			Diff:                 it.Diff,
			Status:               it.Status,
			Reason:               it.Reason,
		})
	}
	ok(c, map[string]any{"items": out, "total": total})
}

func (s *Server) apiWarehouseSyncOptions(c *gin.Context) {
	dscoIDs, err := s.warehouseStore.ListDistinctDSCOWarehouseIDs(c.Request.Context(), 200)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	lxIDs, err := s.warehouseStore.ListDistinctLingXingWarehouseIDs(c.Request.Context(), 200)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	ok(c, map[string]any{
		"dsco_warehouse_ids":     dscoIDs,
		"lingxing_warehouse_ids": lxIDs,
	})
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
	loc := loadDisplayLocation(s.env.Admin.DisplayTimezone)
	finalName, err := writeCSVExportFile(dir, "dsco_order_sync", func(w *csv.Writer) error {
		if err := w.Write([]string{"id", "po_number", "dsco_create_time", "created_at", "updated_at", "dsco_status", "dsco_retailer_id", "status", "warehouse_id", "shipment", "shipped_tracking_no", "dsco_invoice_id"}); err != nil {
			return err
		}
		offset := 0
		for {
			filter.Offset = offset
			items, total, err := s.orderStore.ListByCreateTimeAsc(c.Request.Context(), filter)
			if err != nil {
				return err
			}
			for _, it := range items {
				if err := w.Write([]string{
					strconv.FormatInt(it.ID, 10),
					it.PONumber,
					formatUnixSecToDisplay(it.DSCOCreateTime, loc),
					formatUnixSecToDisplay(it.CreatedAt, loc),
					formatUnixSecToDisplay(it.UpdatedAt, loc),
					it.DSCOStatus,
					it.DSCOREtailerID,
					strconv.FormatInt(int64(it.Status), 10),
					it.WarehouseID,
					it.Shipment,
					it.ShippedTrackingNo,
					it.DSCOInvoiceID,
				}); err != nil {
					return err
				}
			}
			offset += len(items)
			if offset >= int(total) || len(items) == 0 {
				break
			}
		}
		return nil
	})
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	c.FileAttachment(finalName, filepath.Base(finalName))
}

func (s *Server) apiExportWarehouseSync(c *gin.Context) {
	type exportWarehouseBody struct {
		store.DSCOWarehouseSyncListFilter
		DiffRange string `json:"diffRange"`
	}
	var body exportWarehouseBody
	_ = c.ShouldBindJSON(&body)
	filter := body.DSCOWarehouseSyncListFilter
	applyWarehouseDiffRange(&filter, body.DiffRange)
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
	finalName, err := writeCSVExportFile(dir, "dsco_warehouse_sync", func(w *csv.Writer) error {
		if err := w.Write([]string{"id", "sync_time", "dsco_warehouse_id", "dsco_warehouse_sku", "dsco_warehouse_num", "lingxing_warehouse_id", "lingxing_warehouse_sku", "lingxing_warehouse_num", "diff", "status", "reason"}); err != nil {
			return err
		}
		offset := 0
		for {
			filter.Offset = offset
			items, total, err := s.warehouseStore.List(c.Request.Context(), filter)
			if err != nil {
				return err
			}
			for _, it := range items {
				if err := w.Write([]string{
					strconv.FormatInt(it.ID, 10),
					strconv.FormatInt(it.SyncTime, 10),
					it.DSCOWarehouseID,
					it.DSCOWarehouseSKU,
					strconv.Itoa(it.DSCOWarehouseNum),
					it.LingXingWarehouseID,
					it.LingXingWarehouseSKU,
					strconv.Itoa(it.LingXingWarehouseNum),
					strconv.Itoa(it.Diff),
					strconv.FormatInt(int64(it.Status), 10),
					it.Reason,
				}); err != nil {
					return err
				}
			}
			offset += len(items)
			if offset >= int(total) || len(items) == 0 {
				break
			}
		}
		return nil
	})
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
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

func applyWarehouseDiffRange(f *store.DSCOWarehouseSyncListFilter, diffRange string) {
	diffRange = strings.TrimSpace(diffRange)
	if diffRange == "" || f == nil {
		return
	}
	switch diffRange {
	case "lt_-5":
		v := -6
		f.DiffMax = &v
	case "neg_5_1":
		min := -5
		max := -1
		f.DiffMin = &min
		f.DiffMax = &max
	case "eq_0":
		v := 0
		f.DiffEq = &v
	case "pos_1_5":
		min := 1
		max := 5
		f.DiffMin = &min
		f.DiffMax = &max
	case "gt_5":
		v := 6
		f.DiffMin = &v
	default:
		// ignore unknown
	}
}

// Keep json import used (payload debug).
var _ = json.RawMessage{}
