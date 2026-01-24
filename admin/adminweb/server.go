package adminweb

import (
	"errors"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/admin/store"
)

type Options struct {
	AdminPassword string

	OpsBaseURL  string
	OpsPassword string

	Watermark *store.WatermarkStore
	Manual    *store.ManualTaskStore
	Order     *store.OrderStateStore

	Now func() time.Time
}

type Server struct {
	adminPassword string
	opsBaseURL    string
	opsPassword   string
	now           func() time.Time

	watermark *store.WatermarkStore
	manual    *store.ManualTaskStore
	order     *store.OrderStateStore

	engine *gin.Engine
}

func NewServer(opt Options) http.Handler {
	if opt.Watermark == nil {
		panic("Watermark 不能为 nil")
	}
	if opt.Manual == nil {
		panic("Manual 不能为 nil")
	}
	if opt.Order == nil {
		panic("Order 不能为 nil")
	}
	if opt.Now == nil {
		opt.Now = time.Now
	}

	s := &Server{
		adminPassword: opt.AdminPassword,
		opsBaseURL:    opt.OpsBaseURL,
		opsPassword:   opt.OpsPassword,
		now:           opt.Now,
		watermark:     opt.Watermark,
		manual:        opt.Manual,
		order:         opt.Order,
	}

	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.Use(s.traceMiddleware())

	tpl, err := template.New("admin").Funcs(template.FuncMap{
		"safeJSON": func(b []byte) template.HTML {
			// 仅用于把 JSON 原样展示在页面上（先做 HTML 转义，避免 XSS）。
			return template.HTML(template.HTMLEscapeString(string(b)))
		},
	}).ParseFS(assetsFS, "templates/*.html")
	if err != nil {
		panic(err)
	}
	s.engine.SetHTMLTemplate(tpl)

	s.registerRoutes()
	return s.engine
}

func (s *Server) traceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-Id")
		if traceID == "" {
			traceID = logger.NewTraceID()
		}
		ctx := logger.WithTraceID(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Request.Header.Set("X-Trace-Id", traceID)
		c.Header("X-Trace-Id", traceID)
		c.Next()
	}
}

func (s *Server) registerRoutes() {
	// 健康检查不需要鉴权。
	s.engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	// 静态资源（CSS/JS）。
	s.engine.GET("/admin/ui/static/:name", s.handleStatic)

	// 登录页本身不需要鉴权。
	s.engine.GET("/admin/ui/login", s.uiLoginGet)
	s.engine.POST("/admin/ui/login", s.uiLoginPost)
	s.engine.POST("/admin/ui/logout", s.uiLogoutPost)

	// 需要鉴权的 UI。
	ui := s.engine.Group("/admin/ui", s.requireAdminForUI())
	ui.GET("/", s.uiDashboard)
	ui.GET("/orders", s.uiOrders)
	ui.GET("/order", s.uiOrderDetail)
	ui.GET("/manual_tasks", s.uiManualTasks)
	ui.GET("/watermarks", s.uiWatermarks)
	ui.GET("/jobs", s.uiJobs)

	// 需要鉴权的 API（保留原路径，兼容脚本/CI）。
	api := s.engine.Group("/admin", s.requireAdminForAPI())
	api.POST("/run", s.apiRunJob)
	api.GET("/watermark/get", s.apiWatermarkGet)
	api.POST("/watermark/set", s.apiWatermarkSet)
	api.GET("/manual_tasks", s.apiManualTasks)
	api.GET("/order_state/get", s.apiOrderStateGet)
	api.GET("/order_states", s.apiOrderStates)

	// 兼容建议的 `/admin/api` 前缀（可选别名）。
	api2 := s.engine.Group("/admin/api", s.requireAdminForAPI())
	api2.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })
	api2.POST("/run", s.apiRunJob)
	api2.GET("/watermark/get", s.apiWatermarkGet)
	api2.POST("/watermark/set", s.apiWatermarkSet)
	api2.GET("/manual_tasks", s.apiManualTasks)
	api2.GET("/order_state/get", s.apiOrderStateGet)
	api2.GET("/order_states", s.apiOrderStates)
}

func (s *Server) getAdminPassword() string {
	return s.adminPassword
}

func mustJSONBody(c *gin.Context, maxBytes int64) ([]byte, error) {
	if maxBytes <= 0 {
		maxBytes = 1 << 20
	}
	body, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	if int64(len(body)) > maxBytes {
		return nil, errors.New("body 太大")
	}
	if len(body) == 0 {
		return nil, errors.New("empty body")
	}
	if !jsonValid(body) {
		return nil, errors.New("invalid json")
	}
	return body, nil
}
