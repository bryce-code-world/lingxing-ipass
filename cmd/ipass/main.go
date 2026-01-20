package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/internal/platform/adminhttp"
	"lingxingipass/internal/platform/config"
	"lingxingipass/internal/platform/db"
	"lingxingipass/internal/platform/opshttp"
	"lingxingipass/internal/platform/scheduler"
	"lingxingipass/internal/store"
	"lingxingipass/internal/sync"
)

func main() {
	// 约定：优先从项目根目录 `.env` 加载配置（不覆盖已有环境变量），便于本地开发与部署一致。
	if err := config.LoadDotEnv(".env"); err != nil {
		log.Fatalf("加载 .env 失败: %v", err)
	}

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	gdb, err := db.OpenMySQL(cfg.System.DB.DSN)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	if rawDB, err := gdb.DB(); err == nil {
		defer rawDB.Close()
	}

	if err := logger.Init(logger.Config{Dir: cfg.System.Log.Dir, Stdout: true}); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
	defer logger.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := scheduler.New()
	// 一期先跑通骨架：具体 jobs 按开发步骤逐个接入。
	s.Add("heartbeat", time.Duration(cfg.System.Jobs.HeartbeatIntervalSec)*time.Second, func(ctx context.Context) error {
		logger.Info(ctx, "heartbeat", "time", time.Now().UTC().Format(time.RFC3339))
		return nil
	})

	// store 始终初始化：HTTP 管理端需要读写水位、查看人工任务。
	orderState, err := store.NewOrderStateStore(gdb)
	if err != nil {
		log.Fatalf("初始化 OrderStateStore 失败: %v", err)
	}
	watermark, err := store.NewWatermarkStore(gdb)
	if err != nil {
		log.Fatalf("初始化 WatermarkStore 失败: %v", err)
	}
	manual, err := store.NewManualTaskStore(gdb)
	if err != nil {
		log.Fatalf("初始化 ManualTaskStore 失败: %v", err)
	}
	orderRaw, err := store.NewDscoOrderRawStore(gdb)
	if err != nil {
		log.Fatalf("初始化 DscoOrderRawStore 失败: %v", err)
	}

	runners := map[string]adminhttp.JobRunner{}

	hasDSCO := strings.TrimSpace(cfg.System.DSCO.Token) != ""
	hasLingXing := strings.TrimSpace(cfg.System.LingXing.AppID) != "" && strings.TrimSpace(cfg.System.LingXing.AccessToken) != ""

	var dscoCli *dsco.Client
	if hasDSCO {
		dscoCli, err = dsco.New(dsco.Config{
			BaseURL: cfg.System.DSCO.BaseURL,
			Token:   cfg.System.DSCO.Token,
		})
		if err != nil {
			log.Fatalf("初始化 DSCO SDK 失败: %v", err)
		}
	}

	var lxCli *lingxing.Client
	if hasLingXing {
		lxCli, err = lingxing.New(lingxing.Config{
			BaseURL:     cfg.System.LingXing.BaseURL,
			AppID:       cfg.System.LingXing.AppID,
			AccessToken: cfg.System.LingXing.AccessToken,
		})
		if err != nil {
			log.Fatalf("初始化 领星 SDK 失败: %v", err)
		}
	}

	var p *sync.OrderPipeline
	if dscoCli != nil {
		p, err = sync.NewOrderPipeline(dscoCli, lxCli, orderState, watermark, manual, orderRaw, time.Now, cfg.System.Reliability.MaxRetryPerOrder, cfg.Biz.Shipment.ShipDateSource)
		if err != nil {
			log.Fatalf("初始化 OrderPipeline 失败: %v", err)
		}
	}



	var sp *sync.StockPipeline
	if dscoCli != nil && lxCli != nil && len(cfg.Biz.Stock.LingXingWIDToDSCOWarehouseCode) > 0 {
		sp, err = sync.NewStockPipeline(dscoCli, lxCli, manual, cfg.Biz.Stock.LingXingWIDToDSCOWarehouseCode, cfg.Biz.Stock.LingXingSKUToDSCOSKU)
		if err != nil {
			log.Fatalf("初始化 StockPipeline 失败: %v", err)
		}
	}

	// 调度任务接入
	// ops 手动触发：无论定时任务开关是否启用，都必须注册标准 job 名称，避免返回 unknown job。
	registerStandardOpsRunners(runners, opsRunnerDeps{
		PullDSCOOrders: func(ctx context.Context) error {
			if p == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 pull_dsco_orders")
			}
			return p.PullOrders(ctx)
		},
		PushOrdersToLingXing: func(ctx context.Context) error {
			if p == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 push_orders_to_lingxing")
			}
			if lxCli == nil {
				return errors.New("未配置 IPASS_LINGXING_APP_ID / IPASS_LINGXING_ACCESS_TOKEN，无法执行 push_orders_to_lingxing")
			}
			return p.PushOrdersToLingXing(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID, cfg.System.Jobs.PushOrdersToLingXingBatchSize)
		},
		AckToDSCO: func(ctx context.Context) error {
			if p == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 ack_to_dsco")
			}
			if lxCli == nil {
				return errors.New("未配置 IPASS_LINGXING_APP_ID / IPASS_LINGXING_ACCESS_TOKEN，无法执行 ack_to_dsco")
			}
			return p.AckToDSCO(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID)
		},
		ShipToDSCO: func(ctx context.Context) error {
			if p == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 ship_to_dsco")
			}
			if lxCli == nil {
				return errors.New("未配置 IPASS_LINGXING_APP_ID / IPASS_LINGXING_ACCESS_TOKEN，无法执行 ship_to_dsco")
			}
			if cfg.System.LingXing.SID <= 0 {
				return errors.New("未配置 IPASS_LINGXING_SID，无法执行 ship_to_dsco")
			}
			return p.ShipToDSCO(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID, cfg.System.LingXing.SID)
		},
		InvoiceToDSCO: func(ctx context.Context) error {
			if p == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 invoice_to_dsco")
			}
			return p.InvoiceToDSCO(ctx, cfg.System.Jobs.InvoiceToDSCOBatchSize)
		},
		SyncStock: func(ctx context.Context) error {
			if dscoCli == nil {
				return errors.New("未配置 IPASS_DSCO_TOKEN，无法执行 sync_stock")
			}
			if lxCli == nil {
				return errors.New("未配置 IPASS_LINGXING_APP_ID / IPASS_LINGXING_ACCESS_TOKEN，无法执行 sync_stock")
			}
			if len(cfg.Biz.Stock.LingXingWIDToDSCOWarehouseCode) == 0 {
				return errors.New("未配置 IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON，无法执行 sync_stock")
			}
			if sp == nil {
				return errors.New("sync_stock 初始化失败，请检查库存映射与外部系统配置")
			}
			return sp.SyncStock(ctx, cfg.System.Jobs.SyncStockBatchSize)
		},
	})

	if cfg.System.Jobs.PullDSCOOrdersEnable && p != nil {
		s.Add("pull_dsco_orders", time.Duration(cfg.System.Jobs.PullDSCOOrdersIntervalSec)*time.Second, func(ctx context.Context) error {
			return p.PullOrders(ctx)
		})
	}
	if cfg.System.Jobs.PushOrdersToLingXingEnable && p != nil {
		s.Add("push_orders_to_lingxing", time.Duration(cfg.System.Jobs.PushOrdersToLingXingIntervalSec)*time.Second, func(ctx context.Context) error {
			return p.PushOrdersToLingXing(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID, cfg.System.Jobs.PushOrdersToLingXingBatchSize)
		})
	}
	if cfg.System.Jobs.AckToDSCOEnable && p != nil {
		s.Add("ack_to_dsco", time.Duration(cfg.System.Jobs.AckToDSCOIntervalSec)*time.Second, func(ctx context.Context) error {
			return p.AckToDSCO(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID)
		})
	}
	if cfg.System.Jobs.ShipToDSCOEnable && p != nil {
		s.Add("ship_to_dsco", time.Duration(cfg.System.Jobs.ShipToDSCOIntervalSec)*time.Second, func(ctx context.Context) error {
			return p.ShipToDSCO(ctx, cfg.System.LingXing.PlatformCode, cfg.System.LingXing.StoreID, cfg.System.LingXing.SID)
		})
	}
	if cfg.System.Jobs.InvoiceToDSCOEnable && p != nil {
		s.Add("invoice_to_dsco", time.Duration(cfg.System.Jobs.InvoiceToDSCOIntervalSec)*time.Second, func(ctx context.Context) error {
			return p.InvoiceToDSCO(ctx, cfg.System.Jobs.InvoiceToDSCOBatchSize)
		})
	}
	if cfg.System.Jobs.SyncStockEnable && sp != nil {
		s.Add("sync_stock", time.Duration(cfg.System.Jobs.SyncStockIntervalSec)*time.Second, func(ctx context.Context) error {
			return sp.SyncStock(ctx, cfg.System.Jobs.SyncStockBatchSize)
		})
	}

	// HTTP 管理端：始终可用（用于改水位/查人工任务/手动跑一次任务）
	if cfg.System.HTTP.Enable {
		h, err := adminhttp.NewHandler(watermark, manual, orderState, runners)
		if err != nil {
			log.Fatalf("初始化 ops handler 失败: %v", err)
		}
		authed := opshttp.Wrap(h, cfg.System.Ops.Password)

		srv := &http.Server{Addr: cfg.System.HTTP.Addr, Handler: authed}
		go func() {
			logger.Info(ctx, "ops_http_start", "addr", cfg.System.HTTP.Addr)
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error(ctx, "ops_http_error", "err", err.Error())
				stop()
			}
		}()
		go func() {
			<-ctx.Done()
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = srv.Shutdown(shutdownCtx)
		}()
	}

	s.Start(ctx)
	<-ctx.Done()
	s.Stop()
}
