-- 创建运行时配置表
CREATE TABLE IF NOT EXISTS runtime_config (
  id              bigserial PRIMARY KEY,
  domain          text NOT NULL,
  config          jsonb NOT NULL,
  updated_at      bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT uk_runtime_config_domain UNIQUE (domain)
);

COMMENT ON TABLE runtime_config IS '运行时配置（Admin 修改，保存即生效；任务执行使用快照）';
COMMENT ON COLUMN runtime_config.domain IS '业务领域标识（例如 dsco_lingxing；同一 domain 仅保留 1 份当前配置）';
COMMENT ON COLUMN runtime_config.config IS '运行时配置 JSON（外部系统连接/任务开关周期批大小/映射/展示时区/导出目录等）';
COMMENT ON COLUMN runtime_config.updated_at IS '配置更新时间（UTC 秒级时间戳）';

-- 创建 DSCO 订单同步状态表
CREATE TABLE IF NOT EXISTS dsco_order_sync (
  id                bigserial PRIMARY KEY,

  po_number         text NOT NULL,
  dsco_create_time  bigint NOT NULL,
  status            smallint NOT NULL,

  payload           jsonb NOT NULL,
  mskus             text[] NOT NULL DEFAULT '{}'::text[],
  warehouse_id      text NOT NULL DEFAULT '',
  shipment          text NOT NULL DEFAULT '',

  shipped_tracking_no text NOT NULL DEFAULT '',
  dsco_invoice_id     text NOT NULL DEFAULT '',

  created_at        bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),
  updated_at        bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT uk_dsco_order_sync_po_number UNIQUE (po_number),
  CONSTRAINT ck_dsco_order_sync_status CHECK (status BETWEEN 1 AND 5)
);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_create_time
  ON dsco_order_sync (dsco_create_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_status_create_time
  ON dsco_order_sync (status, dsco_create_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_warehouse_id
  ON dsco_order_sync (warehouse_id);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_shipped_tracking_no
  ON dsco_order_sync (shipped_tracking_no);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_dsco_invoice_id
  ON dsco_order_sync (dsco_invoice_id);

-- 支持按 SKU 筛选：WHERE 'sku123' = ANY(mskus)
CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_mskus_gin
  ON dsco_order_sync USING GIN (mskus);

COMMENT ON TABLE dsco_order_sync IS 'DSCO 订单同步状态表（状态机 1~5 + 原始 payload）';
COMMENT ON COLUMN dsco_order_sync.po_number IS 'DSCO 订单唯一键';
COMMENT ON COLUMN dsco_order_sync.dsco_create_time IS 'DSCO 订单 created_at（UTC 秒级时间戳，用于增量拉单）';
COMMENT ON COLUMN dsco_order_sync.status IS '同步状态：1待同步 2待确认(ack) 3已确认 4已发货 5已创建发票(完成态)';
COMMENT ON COLUMN dsco_order_sync.payload IS 'DSCO 原始订单 JSON（覆盖更新）';
COMMENT ON COLUMN dsco_order_sync.mskus IS '订单 SKU 列表（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.warehouse_id IS 'DSCO 发货仓库 ID（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.shipment IS 'DSCO 指定物流方式（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.shipped_tracking_no IS '已回传的物流跟踪号（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.dsco_invoice_id IS '已回传的发票 ID（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.created_at IS '创建时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_order_sync.updated_at IS '更新时间（UTC 秒级时间戳）';

-- 创建 DSCO 仓库同步日志表
CREATE TABLE IF NOT EXISTS dsco_warehouse_sync (
  id                    bigserial PRIMARY KEY,

  sync_time             bigint NOT NULL,

  dsco_warehouse_id     text NOT NULL DEFAULT '',
  dsco_warehouse_sku    text NOT NULL DEFAULT '',
  dsco_warehouse_num    integer NOT NULL DEFAULT 0,

  lingxing_warehouse_id text NOT NULL DEFAULT '',
  lingxing_warehouse_sku text NOT NULL DEFAULT '',
  lingxing_warehouse_num integer NOT NULL DEFAULT 0,

  status                smallint NOT NULL,
  reason                text NOT NULL DEFAULT '',

  created_at            bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),
  updated_at            bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT ck_dsco_warehouse_sync_status CHECK (status IN (1,2))
);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_time
  ON dsco_warehouse_sync (sync_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_dsco_wh_sku_time
  ON dsco_warehouse_sync (dsco_warehouse_id, dsco_warehouse_sku, sync_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_status_time
  ON dsco_warehouse_sync (status, sync_time DESC);

COMMENT ON TABLE dsco_warehouse_sync IS '库存同步日志（领星与 DSCO 对账/回写结果）';
COMMENT ON COLUMN dsco_warehouse_sync.sync_time IS '同步时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_id IS 'DSCO 仓库 ID';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_sku IS 'DSCO 仓库 SKU';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_num IS 'DSCO 数量';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_id IS '领星仓库 ID';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_sku IS '领星仓库 SKU';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_num IS '领星数量';
COMMENT ON COLUMN dsco_warehouse_sync.status IS '同步状态：1成功 2失败';
COMMENT ON COLUMN dsco_warehouse_sync.reason IS '失败原因';
COMMENT ON COLUMN dsco_warehouse_sync.created_at IS '创建时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_warehouse_sync.updated_at IS '更新时间（UTC 秒级时间戳）';