package admin

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"

	"lingxingipass/infra/runtimecfg"
)

type dashboardJobRow struct {
	Job    runtimecfg.JobName
	Desc   string
	Enable bool
	Cron   string
	Size   int
}

func jobDesc(job runtimecfg.JobName) string {
	switch job {
	case runtimecfg.JobPullDSCOOrders:
		return "拉取 DSCO 订单入库"
	case runtimecfg.JobPushToLingXing:
		return "推送订单到领星（生成出库单）"
	case runtimecfg.JobAckToDSCO:
		return "回传确认（ACK）到 DSCO"
	case runtimecfg.JobShipToDSCO:
		return "回传发货（运单号）到 DSCO"
	case runtimecfg.JobInvoiceToDSCO:
		return "回传发票到 DSCO"
	case runtimecfg.JobSyncStock:
		return "同步库存到领星"
	case runtimecfg.JobPullSKUPair:
		return "拉取 SKU 配对数据"
	case runtimecfg.JobCleanupExports:
		return "清理导出文件"
	default:
		return ""
	}
}

func sortJobsForDashboard(jobs map[runtimecfg.JobName]runtimecfg.JobConfig) []dashboardJobRow {
	if len(jobs) == 0 {
		return nil
	}

	order := []runtimecfg.JobName{
		runtimecfg.JobPullDSCOOrders,
		runtimecfg.JobPushToLingXing,
		runtimecfg.JobAckToDSCO,
		runtimecfg.JobShipToDSCO,
		runtimecfg.JobInvoiceToDSCO,
		runtimecfg.JobSyncStock,
		runtimecfg.JobCleanupExports,
	}
	orderIndex := make(map[runtimecfg.JobName]int, len(order))
	for i, j := range order {
		orderIndex[j] = i
	}

	names := make([]runtimecfg.JobName, 0, len(jobs))
	for j := range jobs {
		names = append(names, j)
	}
	sort.Slice(names, func(i, j int) bool {
		ai, aok := orderIndex[names[i]]
		bi, bok := orderIndex[names[j]]
		if aok && bok {
			return ai < bi
		}
		if aok != bok {
			return aok
		}
		return string(names[i]) < string(names[j])
	})

	out := make([]dashboardJobRow, 0, len(names))
	for _, name := range names {
		cfg := jobs[name]
		out = append(out, dashboardJobRow{
			Job:    name,
			Desc:   jobDesc(name),
			Enable: cfg.Enable,
			Cron:   cfg.Cron,
			Size:   cfg.Size,
		})
	}
	return out
}

func (s *Server) uiLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"Error": "",
	})
}

func (s *Server) uiLoginPost(c *gin.Context) {
	pass := strings.TrimSpace(c.PostForm("password"))
	if pass == "" || pass != s.env.Admin.Password {
		c.HTML(http.StatusUnauthorized, "login", gin.H{
			"Error": "invalid password",
		})
		return
	}
	s.setSession(c)
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

func (s *Server) uiLogout(c *gin.Context) {
	s.clearSession(c)
	c.Redirect(http.StatusFound, "/admin/login")
}

func (s *Server) uiDashboard(c *gin.Context) {
	rc, _ := s.cfgMgr.Snapshot(runtimecfg.DomainDSCOLingXing)
	c.HTML(http.StatusOK, "dashboard", gin.H{
		"Title":           "Dashboard",
		"DisplayTimezone": s.env.Admin.DisplayTimezone,
		"Domain":          runtimecfg.DomainDSCOLingXing,
		"UpdatedAt":       rc.UpdatedAt,
		"Jobs":            sortJobsForDashboard(rc.Config.Jobs),
	})
}

func (s *Server) uiConfig(c *gin.Context) {
	rc, ok := s.cfgMgr.Snapshot(runtimecfg.DomainDSCOLingXing)
	var cfgJSON string
	if ok {
		if b, err := json.MarshalIndent(rc.Config, "", "  "); err == nil {
			cfgJSON = string(b)
		}
	}
	c.HTML(http.StatusOK, "config", gin.H{
		"Title":           "Runtime Config",
		"DisplayTimezone": s.env.Admin.DisplayTimezone,
		"Domain":          runtimecfg.DomainDSCOLingXing,
		"ConfigJSON":      cfgJSON,
	})
}

func (s *Server) uiTasks(c *gin.Context) {
	c.HTML(http.StatusOK, "tasks", gin.H{
		"Title":           "Tasks",
		"DisplayTimezone": s.env.Admin.DisplayTimezone,
	})
}

func (s *Server) uiOrders(c *gin.Context) {
	c.HTML(http.StatusOK, "orders", gin.H{
		"Title":           "Orders",
		"DisplayTimezone": s.env.Admin.DisplayTimezone,
	})
}

func (s *Server) uiWarehouses(c *gin.Context) {
	c.HTML(http.StatusOK, "warehouses", gin.H{
		"Title":           "Warehouses",
		"DisplayTimezone": s.env.Admin.DisplayTimezone,
	})
}
