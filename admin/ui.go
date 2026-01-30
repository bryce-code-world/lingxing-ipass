package admin

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"lingxingipass/infra/runtimecfg"
)

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
		"Jobs":            rc.Config.Jobs,
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
