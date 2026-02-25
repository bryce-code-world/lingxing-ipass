package httpserver

import (
	"context"
	"net/http"
	"time"

	"lingxingipass/golib/v2/tool/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lingxingipass/infra/runtimecfg"
)

type Server struct {
	srv *http.Server
}

func New(listenAddr string, engine *gin.Engine, cfgMgr *runtimecfg.Manager, gdb *gorm.DB) *Server {
	// Minimal liveness.
	engine.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// Readiness: runtime config loaded; optional DB ping.
	engine.GET("/readyz", func(c *gin.Context) {
		if _, ok := cfgMgr.Snapshot(runtimecfg.DomainDSCOLingXing); !ok {
			c.String(http.StatusServiceUnavailable, "runtime config not loaded")
			return
		}
		sqlDB, err := gdb.DB()
		if err == nil {
			if pingErr := sqlDB.PingContext(c.Request.Context()); pingErr != nil {
				c.String(http.StatusServiceUnavailable, "db ping failed")
				return
			}
		}
		c.String(http.StatusOK, "ok")
	})

	s := &http.Server{
		Addr:              listenAddr,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      5 * time.Minute,
		IdleTimeout:       60 * time.Second,
	}
	logger.Info(context.Background(), "http listen", "addr", listenAddr)
	return &Server{srv: s}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
