package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lingxingipass/golib/v2/tool/logger"

	"lingxingipass/admin"
	"lingxingipass/api"
	"lingxingipass/infra/config"
	"lingxingipass/infra/db"
	"lingxingipass/infra/runtimecfg"
	"lingxingipass/infra/store"
	"lingxingipass/integration"
	"lingxingipass/integration/dsco_lingxing"
	"lingxingipass/integration/ops"
	"lingxingipass/transport/httpserver"
	"lingxingipass/transport/scheduler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	envCfg, err := config.LoadEnvYAML("env.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "load env.yaml failed:", err)
		os.Exit(1)
	}
	if err := config.ValidateEnv(envCfg); err != nil {
		fmt.Fprintln(os.Stderr, "validate env.yaml failed:", err)
		os.Exit(1)
	}

	if err := logger.Init(envCfg.Log); err != nil {
		fmt.Fprintln(os.Stderr, "init logger failed:", err)
		os.Exit(1)
	}
	defer logger.Sync()

	gdb, err := db.OpenPostgres(ctx, envCfg.DB)
	if err != nil {
		logger.Error(ctx, "open postgres failed", "err", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(gdb); err != nil {
			logger.Warn(context.Background(), "close db failed", "err", err)
		}
	}()

	runtimeStore := store.NewRuntimeConfigStore(gdb)
	orderStore := store.NewDSCOOrderSyncStore(gdb)
	warehouseStore := store.NewDSCOWarehouseSyncStore(gdb)

	cfgMgr := runtimecfg.NewManager(runtimeStore)
	if err := cfgMgr.LoadOrInit(ctx, runtimecfg.DomainDSCOLingXing); err != nil {
		logger.Error(ctx, "load runtime config failed", "err", err)
		os.Exit(1)
	}

	reg := integration.NewRegistry()
	runner := integration.NewRunner(cfgMgr, reg, gdb)

	// Register domain tasks.
	dscoLingxing := dsco_lingxing.NewDomain(envCfg, orderStore, warehouseStore)
	dsco_lingxing.Register(reg, dscoLingxing)
	opsDomain := ops.NewDomain(envCfg)
	ops.Register(reg, opsDomain)

	adminServer := admin.NewServer(envCfg, cfgMgr, reg, runner, orderStore, warehouseStore)
	api.Register(adminServer.Engine(), envCfg, runner, warehouseStore)

	httpSrv := httpserver.New(envCfg.Base.ListenAddr, adminServer.Engine(), cfgMgr, gdb)

	sched := scheduler.New(cfgMgr, reg, runner)

	// Start scheduler before HTTP (both will stop on ctx).
	if err := sched.Start(ctx); err != nil {
		logger.Error(ctx, "start scheduler failed", "err", err)
		os.Exit(1)
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		_ = sched.Stop(shutdownCtx)
		_ = httpSrv.Shutdown(shutdownCtx)
	}()

	if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(ctx, "http server error", "err", err)
		os.Exit(1)
	}
}
