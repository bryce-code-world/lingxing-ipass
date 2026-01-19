package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/admin/adminweb"
	"lingxingipass/admin/platform/config"
	"lingxingipass/admin/platform/db"
	"lingxingipass/admin/store"
)

func main() {
	// admin 独立进程：优先从项目根目录 `.env.admin` 加载（不覆盖已有环境变量）。
	if err := config.LoadDotEnv(".env.admin"); err != nil {
		log.Fatalf("加载 .env.admin 失败: %v", err)
	}

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("加载 admin 配置失败: %v", err)
	}

	gdb, err := db.OpenMySQL(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	if rawDB, err := gdb.DB(); err == nil {
		defer rawDB.Close()
	}

	// admin 进程也使用统一 logger（便于统一格式与追踪）。
	if err := logger.Init(logger.Config{Dir: "logs", Stdout: true}); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
	defer logger.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	watermark, err := store.NewWatermarkStore(gdb)
	if err != nil {
		log.Fatalf("初始化 WatermarkStore 失败: %v", err)
	}
	manual, err := store.NewManualTaskStore(gdb)
	if err != nil {
		log.Fatalf("初始化 ManualTaskStore 失败: %v", err)
	}
	order, err := store.NewOrderStateStore(gdb)
	if err != nil {
		log.Fatalf("初始化 OrderStateStore 失败: %v", err)
	}

	h := adminweb.NewServer(adminweb.Options{
		AdminPassword: cfg.Auth.Password,
		OpsBaseURL:    cfg.Ops.BaseURL,
		OpsPassword:   cfg.Ops.Password,
		Watermark:     watermark,
		Manual:        manual,
		Order:         order,
		Now:           time.Now,
	})

	if !cfg.HTTP.Enable {
		log.Printf("ADMIN_HTTP_ENABLE=false，admin 不启动 HTTP 服务")
		return
	}

	srv := &http.Server{Addr: cfg.HTTP.Addr, Handler: h}
	go func() {
		logger.Info(ctx, "admin_http_start", "addr", cfg.HTTP.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, "admin_http_error", "err", err.Error())
			stop()
		}
	}()
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	<-ctx.Done()
}
