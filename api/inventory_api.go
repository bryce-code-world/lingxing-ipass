package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gitee.com/lsy007/golibv2/v2/tool/logger"
	"github.com/gin-gonic/gin"

	"lingxingipass/infra/runtimecfg"
	"lingxingipass/infra/store"
	"lingxingipass/integration"
	"lingxingipass/integration/dsco_lingxing"
)

type inventoryAPI struct {
	runner         *integration.Runner
	warehouseStore *store.DSCOWarehouseSyncStore
}

func parseBool(s string, def bool) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return def
	}
	return v
}

func parseInt(s string, def int) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func parseUTCDate(s string, def time.Time) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		t := def.UTC()
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), true
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, false
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), true
}

func (a *inventoryAPI) getDiff(c *gin.Context) {
	if a.warehouseStore == nil {
		fail(c, http.StatusServiceUnavailable, 503, "warehouse store not ready")
		return
	}

	dt, okDate := parseUTCDate(c.Query("date"), time.Now())
	if !okDate {
		fail(c, http.StatusBadRequest, 400, "invalid date (expected YYYY-MM-DD)")
		return
	}
	start := dt.Unix()
	end := dt.Add(24 * time.Hour).Unix()

	page := parseInt(c.Query("page"), 1)
	size := parseInt(c.Query("size"), 50)
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 500 {
		size = 50
	}
	offset := (page - 1) * size

	diffOnly := parseBool(c.Query("diff_only"), false)

	filter := store.DSCOWarehouseSyncListFilter{
		StartTime: &start,
		EndTime:   &end,

		DSCOWarehouseID:      strings.TrimSpace(c.Query("dsco_wid")),
		DSCOWarehouseSKU:     strings.TrimSpace(c.Query("dsco_sku")),
		LingXingWarehouseID:  strings.TrimSpace(c.Query("lingxing_wid")),
		LingXingWarehouseSKU: strings.TrimSpace(c.Query("lingxing_sku")),

		DiffNotZero: diffOnly,

		Offset: offset,
		Limit:  size,
	}

	logger.Info(c.Request.Context(), "api inventory diff request",
		"ip", c.ClientIP(),
		"date", dt.Format("2006-01-02"),
		"diff_only", diffOnly,
		"page", page,
		"size", size,
		"dsco_wid", filter.DSCOWarehouseID,
		"dsco_sku", filter.DSCOWarehouseSKU,
		"lingxing_wid", filter.LingXingWarehouseID,
		"lingxing_sku", filter.LingXingWarehouseSKU,
	)

	items, total, err := a.warehouseStore.ListLatestByFullKey(c.Request.Context(), filter)
	if err != nil {
		logger.Warn(c.Request.Context(), "api inventory diff failed", "err", err)
		fail(c, http.StatusInternalServerError, 500, err.Error())
		return
	}

	type item struct {
		SyncTime string `json:"sync_time"`
		DSCOWID  string `json:"dsco_warehouse_id"`
		DSCOSKU  string `json:"dsco_sku"`
		DSCONum  int    `json:"dsco_num"`

		LingXingWID string `json:"lingxing_wid"`
		LingXingSKU string `json:"lingxing_sku"`
		LingXingNum int    `json:"lingxing_num"`

		Diff   int    `json:"diff"`
		Status int16  `json:"status"`
		Reason string `json:"reason,omitempty"`
	}

	out := make([]item, 0, len(items))
	for _, it := range items {
		st := ""
		if it.SyncTime > 0 {
			st = time.Unix(it.SyncTime, 0).UTC().Format("2006-01-02 15:04:05")
		}
		out = append(out, item{
			SyncTime: st,
			DSCOWID:  it.DSCOWarehouseID,
			DSCOSKU:  it.DSCOWarehouseSKU,
			DSCONum:  it.DSCOWarehouseNum,

			LingXingWID: it.LingXingWarehouseID,
			LingXingSKU: it.LingXingWarehouseSKU,
			LingXingNum: it.LingXingWarehouseNum,

			Diff:   it.Diff,
			Status: it.Status,
			Reason: it.Reason,
		})
	}

	logger.Info(c.Request.Context(), "api inventory diff ok",
		"ip", c.ClientIP(),
		"date", dt.Format("2006-01-02"),
		"returned", len(out),
		"total", total,
	)

	ok(c, map[string]any{
		"date":  dt.Format("2006-01-02"),
		"page":  page,
		"size":  size,
		"total": total,
		"items": out,
	})
}

type syncBody struct {
	DryRun *bool `json:"dry_run"`

	// daily_pulled mode
	Date     string `json:"date"`
	DiffOnly *bool  `json:"diff_only"`

	DSCOWID      string   `json:"dsco_wid"`
	LingXingWID  string   `json:"lingxing_wid"`
	DSCOSKUList  []string `json:"dsco_sku_list"`
	LingXingSKUs []string `json:"lingxing_sku_list"`

	// manual_items mode
	Items []struct {
		DSCOWID string `json:"dsco_wid"`
		DSCOSKU string `json:"dsco_sku"`
		Qty     int    `json:"qty"`
	} `json:"items"`
}

func (a *inventoryAPI) postSync(c *gin.Context) {
	if a.runner == nil {
		fail(c, http.StatusServiceUnavailable, 503, "runner not ready")
		return
	}

	var body syncBody
	if err := c.ShouldBindJSON(&body); err != nil {
		// Empty body is allowed; invalid json should fail.
		if c.Request.ContentLength > 0 {
			fail(c, http.StatusBadRequest, 400, "invalid json")
			return
		}
	}

	dryRun := true
	if body.DryRun != nil {
		dryRun = *body.DryRun
	}
	forceDoSync := !dryRun
	override := dsco_lingxing.SyncStockOverride{
		ForceDoSync: &forceDoSync,
		MaxKeys:     5000,
	}

	var source string
	var reqDate string
	var reqDiffOnly *bool
	var reqItems int
	var reqItemsSample []map[string]any

	if len(body.Items) > 0 {
		reqItems = len(body.Items)
		if len(body.Items) > 1000 {
			fail(c, http.StatusBadRequest, 400, "items too large (max=1000)")
			return
		}
		override.Source = dsco_lingxing.SyncStockSourceManualItems
		override.ManualItems = make([]dsco_lingxing.SyncStockManualItem, 0, len(body.Items))
		for _, it := range body.Items {
			if len(reqItemsSample) >= 10 {
				break
			}
			reqItemsSample = append(reqItemsSample, map[string]any{
				"dsco_wid": strings.TrimSpace(it.DSCOWID),
				"dsco_sku": strings.TrimSpace(it.DSCOSKU),
				"qty":      it.Qty,
			})
		}
		for i, it := range body.Items {
			if strings.TrimSpace(it.DSCOWID) == "" || strings.TrimSpace(it.DSCOSKU) == "" {
				fail(c, http.StatusBadRequest, 400, "items["+strconv.Itoa(i)+"] missing dsco_wid/dsco_sku")
				return
			}
			if it.Qty < 0 {
				fail(c, http.StatusBadRequest, 400, "items["+strconv.Itoa(i)+"] qty must be >= 0")
				return
			}
			override.ManualItems = append(override.ManualItems, dsco_lingxing.SyncStockManualItem{
				DSCOWarehouseID: strings.TrimSpace(it.DSCOWID),
				DSCOSKU:         strings.TrimSpace(it.DSCOSKU),
				Qty:             it.Qty,
			})
		}
		source = "manual_items"
	} else {
		dt, okDate := parseUTCDate(body.Date, time.Now())
		if !okDate {
			fail(c, http.StatusBadRequest, 400, "invalid date (expected YYYY-MM-DD)")
			return
		}
		start := dt.Unix()
		end := dt.Add(24 * time.Hour).Unix()

		diffOnly := true
		if body.DiffOnly != nil {
			diffOnly = *body.DiffOnly
		}

		reqDate = dt.Format("2006-01-02")
		reqDiffOnly = &diffOnly

		override.Source = dsco_lingxing.SyncStockSourceDailyPulled
		override.DailyPulled = &dsco_lingxing.SyncStockDailyPulledOptions{
			StartTime: start,
			EndTime:   end,
			DiffOnly:  diffOnly,

			DSCOWarehouseID:     strings.TrimSpace(body.DSCOWID),
			LingXingWarehouseID: strings.TrimSpace(body.LingXingWID),

			DSCOSKUList:     body.DSCOSKUList,
			LingXingSKUList: body.LingXingSKUs,
		}
		source = "daily_pulled"
	}

	logger.Info(c.Request.Context(), "api inventory sync request",
		"ip", c.ClientIP(),
		"dry_run", dryRun,
		"source", source,
		"date", reqDate,
		"diff_only", reqDiffOnly,
		"items", reqItems,
		"items_sample", reqItemsSample,
		"dsco_wid", strings.TrimSpace(body.DSCOWID),
		"lingxing_wid", strings.TrimSpace(body.LingXingWID),
		"dsco_sku_list", len(body.DSCOSKUList),
		"lingxing_sku_list", len(body.LingXingSKUs),
	)

	res, err := a.runner.RunWithResult(c.Request.Context(), integration.RunRequest{
		Domain:   runtimecfg.DomainDSCOLingXing,
		Job:      runtimecfg.JobSyncStock,
		Trigger:  integration.TriggerManual,
		Override: &override,
	})
	if err != nil {
		logger.Warn(c.Request.Context(), "api inventory sync failed", "ip", c.ClientIP(), "source", source, "dry_run", dryRun, "err", err)
		switch {
		case errors.Is(err, dsco_lingxing.ErrSyncStockTooManyKeys):
			fail(c, http.StatusBadRequest, 400, err.Error())
		case errors.Is(err, integration.ErrJobRunning):
			fail(c, http.StatusConflict, 409, "job running")
		case errors.Is(err, integration.ErrJobNotFound):
			fail(c, http.StatusNotFound, 404, "job not found")
		case errors.Is(err, integration.ErrConfigMissing):
			fail(c, http.StatusServiceUnavailable, 503, "runtime config not loaded")
		default:
			fail(c, http.StatusInternalServerError, 500, err.Error())
		}
		return
	}

	logger.Info(c.Request.Context(), "api inventory sync ok", "ip", c.ClientIP(), "source", source, "dry_run", dryRun, "result", res)

	if res == nil {
		ok(c, map[string]any{
			"ok":      true,
			"job":     "sync_stock",
			"dry_run": dryRun,
			"source":  source,
		})
		return
	}

	ok(c, res)
}
