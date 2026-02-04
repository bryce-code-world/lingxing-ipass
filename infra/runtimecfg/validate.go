package runtimecfg

import (
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func DefaultConfig(domain string) Config {
	return Config{
		Domain: domain,
		Jobs: map[JobName]JobConfig{
			JobPullDSCOOrders: {Enable: false, Cron: "0 5 * * * *", Size: 50},
			JobPushToLingXing: {Enable: false, Cron: "0 10 * * * *", Size: 50},
			JobAckToDSCO:      {Enable: false, Cron: "0 15 * * * *", Size: 50, MultiBan: false},
			JobShipToDSCO:     {Enable: false, Cron: "0 20 * * * *", Size: 50, MultiBan: false},
			JobInvoiceToDSCO:  {Enable: false, Cron: "0 25 * * * *", Size: 50, MultiBan: false},
			JobSyncStock:      {Enable: false, Cron: "0 30 1 * * *", Size: 50, Sync: false},
			JobPullSKUPair:    {Enable: false, Cron: "0 0 19 * * *", Size: 200}, // UTC 每晚 19 点，对应北京时间次日 3 点
			JobCleanupExports: {Enable: true, Cron: "0 0 1 * * *", Size: 1},
		},
		Mapping: Mapping{
			// 初始默认值：用于快速跑通闭环（可在 Admin 后台随时修改并热更新）。
			Shop: map[string]string{
				"1000011436": "110658143132021760",
			},
			Warehouse: map[string]string{
				"YQN-CA":  "33328",
				"YQN-CA2": "28393",
			},
			SKU: map[string]string{},
			Shipment: map[string]string{
				"YQN-CA-FEHD":  "203662740883276800",
				"YQN-CA2-FEHD": "203656345743220736",
				"YQN-CA2-USCG": "203656345742652928",
			},
		},
	}
}

func Validate(cfg Config, supportedJobs map[JobName]struct{}) error {
	if cfg.Domain != DomainDSCOLingXing {
		return fmt.Errorf("domain 不支持：%s", cfg.Domain)
	}
	if len(cfg.Jobs) == 0 {
		return errors.New("jobs 不能为空")
	}
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	for name, jc := range cfg.Jobs {
		if _, ok := supportedJobs[name]; !ok {
			return fmt.Errorf("jobs 包含未知任务：%s", name)
		}
		if jc.Cron == "" {
			return fmt.Errorf("jobs.%s.cron 不能为空", name)
		}
		_, err := parser.Parse(jc.Cron)
		if err != nil {
			return fmt.Errorf("jobs.%s.cron 非法：%w", name, err)
		}
		if jc.Size <= 0 {
			return fmt.Errorf("jobs.%s.size 必须为正整数", name)
		}
	}
	// Basic mapping direction sanity: DSCO->LingXing (strings only).
	// Ensure no empty key/value.
	checkMap := func(name string, m map[string]string) error {
		for k, v := range m {
			if k == "" || v == "" {
				return fmt.Errorf("mapping.%s 不允许空 key/value", name)
			}
		}
		return nil
	}
	if err := checkMap("shop", cfg.Mapping.Shop); err != nil {
		return err
	}
	if err := checkMap("warehouse", cfg.Mapping.Warehouse); err != nil {
		return err
	}
	if err := checkMap("sku", cfg.Mapping.SKU); err != nil {
		return err
	}
	if err := checkMap("shipment", cfg.Mapping.Shipment); err != nil {
		return err
	}
	// If stock job enabled, warehouse mapping must be non-empty.
	if jc, ok := cfg.Jobs[JobSyncStock]; ok && jc.Enable {
		if len(cfg.Mapping.Warehouse) == 0 {
			return errors.New("启用 sync_stock 前：mapping.warehouse 必须非空")
		}
	}
	// If push job enabled, shop mapping must be non-empty to determine store_id.
	if jc, ok := cfg.Jobs[JobPushToLingXing]; ok && jc.Enable {
		if len(cfg.Mapping.Shop) == 0 {
			return errors.New("启用 push_to_lingxing 前：mapping.shop 必须非空")
		}
		// If warehouse mapping exists, shipment mapping should exist as well because push may set wid.
		// When wid is filled, lingxing CreateOrdersV2 requires logistics_type_id.
		if len(cfg.Mapping.Warehouse) > 0 && len(cfg.Mapping.Shipment) == 0 {
			return errors.New("启用 push_to_lingxing 且配置了 mapping.warehouse 时：mapping.shipment 必须非空（用于 logistics_type_id）")
		}
	}
	// If ship/invoice jobs enabled, shop mapping must be non-empty to obtain lingxing SID for WmsOrderList.
	if jc, ok := cfg.Jobs[JobShipToDSCO]; ok && jc.Enable {
		if len(cfg.Mapping.Shop) == 0 {
			return errors.New("启用 ship_to_dsco 前：mapping.shop 必须非空（用于 WmsOrderList.sid_arr）")
		}
	}
	if jc, ok := cfg.Jobs[JobInvoiceToDSCO]; ok && jc.Enable {
		if len(cfg.Mapping.Shop) == 0 {
			return errors.New("启用 invoice_to_dsco 前：mapping.shop 必须非空（用于 WmsOrderList.sid_arr）")
		}
	}

	_ = time.UTC // doc: cron is UTC; enforcement in scheduler.
	return nil
}
