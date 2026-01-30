package admin

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"lingxingipass/infra/config"
	"lingxingipass/infra/runtimecfg"
	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

type Server struct {
	env config.EnvConfig

	cfgMgr *runtimecfg.Manager
	reg    *integration.Registry
	runner *integration.Runner

	orderStore     *store.DSCOOrderSyncStore
	warehouseStore *store.DSCOWarehouseSyncStore

	engine *gin.Engine
}

func NewServer(env config.EnvConfig, cfgMgr *runtimecfg.Manager, reg *integration.Registry, runner *integration.Runner, orderStore *store.DSCOOrderSyncStore, warehouseStore *store.DSCOWarehouseSyncStore) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	tpl := template.Must(template.New("admin").ParseFS(assets, "templates/*.html"))
	engine.SetHTMLTemplate(tpl)
	if sub, err := fs.Sub(assets, "static"); err == nil {
		engine.StaticFS("/admin/static", http.FS(sub))
	}

	s := &Server{
		env:            env,
		cfgMgr:         cfgMgr,
		reg:            reg,
		runner:         runner,
		orderStore:     orderStore,
		warehouseStore: warehouseStore,
		engine:         engine,
	}
	s.registerRoutes()
	return s
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}

func (s *Server) registerRoutes() {
	admin := s.engine.Group("/admin")

	admin.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/admin/dashboard")
	})
	admin.GET("/login", s.uiLogin)
	admin.POST("/login", s.uiLoginPost)
	admin.GET("/logout", s.requireAuth(), s.uiLogout)
	admin.GET("/dashboard", s.requireAuth(), s.uiDashboard)
	admin.GET("/config", s.requireAuth(), s.uiConfig)
	admin.GET("/tasks", s.requireAuth(), s.uiTasks)
	admin.GET("/orders", s.requireAuth(), s.uiOrders)
	admin.GET("/warehouses", s.requireAuth(), s.uiWarehouses)

	api := s.engine.Group("/admin/api")
	api.POST("/auth/login", s.apiLogin)
	api.POST("/auth/logout", s.apiLogout)
	api.GET("/auth/me", s.apiMe)

	authed := api.Group("")
	authed.Use(s.requireAuth())

	authed.GET("/config/runtime", s.apiGetRuntimeConfig)
	authed.PUT("/config/runtime", s.apiUpdateRuntimeConfig)

	authed.POST("/jobs/run", s.apiRunJobs)
	authed.POST("/jobs/run_one", s.apiRunOneJob)

	authed.GET("/dsco_order_sync/list", s.apiListOrders)
	authed.GET("/dsco_order_sync/detail", s.apiOrderDetail)
	authed.POST("/dsco_order_sync/pull", s.apiManualPullOrders)

	authed.GET("/dsco_warehouse_sync/list", s.apiListWarehouseSync)

	authed.POST("/export/dsco_order_sync", s.apiExportOrders)
	authed.POST("/export/dsco_warehouse_sync", s.apiExportWarehouseSync)
}
