package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) uiLogin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<!doctype html>
<html><head><meta charset="utf-8"><title>Admin Login</title></head>
<body style="font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial; margin:40px;">
<h1>Admin Login</h1>
<form method="post" action="/admin/api/auth/login">
  <input type="password" name="password" placeholder="Password" style="padding:10px; width:260px;" />
  <button type="submit" style="padding:10px 16px;">Login</button>
</form>
</body></html>`)
}

func (s *Server) uiDashboard(c *gin.Context)  { s.simplePage(c, "Dashboard") }
func (s *Server) uiConfig(c *gin.Context)     { s.simplePage(c, "Config") }
func (s *Server) uiTasks(c *gin.Context)      { s.simplePage(c, "Tasks") }
func (s *Server) uiOrders(c *gin.Context)     { s.simplePage(c, "Orders") }
func (s *Server) uiWarehouses(c *gin.Context) { s.simplePage(c, "Warehouses") }

func (s *Server) simplePage(c *gin.Context, title string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<!doctype html>
<html><head><meta charset="utf-8"><title>`+title+`</title></head>
<body style="font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial; margin:40px;">
<h1>`+title+`</h1>
<p>API-first MVP. Use <code>/admin/api/*</code>.</p>
<p><a href="/admin/dashboard">Dashboard</a> | <a href="/admin/config">Config</a> | <a href="/admin/tasks">Tasks</a> | <a href="/admin/orders">Orders</a> | <a href="/admin/warehouses">Warehouses</a></p>
</body></html>`)
}
