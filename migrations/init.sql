-- PostgreSQL init schema for lingxingipass (DSCO ↔ 领星).
-- 时间统一使用 UTC 秒级时间戳（bigint）。

-- 运行时配置（Admin 保存即生效；任务执行读取快照）
CREATE TABLE IF NOT EXISTS runtime_config (
  id         bigserial PRIMARY KEY,
  domain     text   NOT NULL,
  config     jsonb  NOT NULL,
  updated_at bigint NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT uk_runtime_config_domain UNIQUE (domain)
);

COMMENT ON TABLE runtime_config IS '运行时配置（Admin 修改后保存即生效；任务执行按“配置快照”读取）';
COMMENT ON COLUMN runtime_config.domain IS '业务领域标识（例如 dsco_lingxing）；同一 domain 仅保留 1 份当前配置';
COMMENT ON COLUMN runtime_config.config IS '运行时配置 JSON（domain/jobs/mapping 等；见配置设计文档）';
COMMENT ON COLUMN runtime_config.updated_at IS '配置更新时间（UTC 秒级时间戳）';

-- DSCO 订单同步状态表
CREATE TABLE IF NOT EXISTS dsco_order_sync (
  id                    bigserial PRIMARY KEY,

  po_number             text     NOT NULL,
  dsco_order_id         text     NOT NULL DEFAULT '',
  consumer_order_number text     NOT NULL DEFAULT '',
  channel               text     NOT NULL DEFAULT '',

  dsco_create_time      bigint   NOT NULL,
  status                smallint NOT NULL,

  payload               jsonb    NOT NULL,
  mskus                 text[]   NOT NULL DEFAULT '{}'::text[],
  warehouse_id          text     NOT NULL DEFAULT '',
  shipment              text     NOT NULL DEFAULT '',
  dsco_retainer_id      text     NOT NULL DEFAULT '',

  shipped_tracking_no   text     NOT NULL DEFAULT '',
  dsco_invoice_id       text     NOT NULL DEFAULT '',

  created_at            bigint   NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),
  updated_at            bigint   NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT uk_dsco_order_sync_po_number UNIQUE (po_number),
  CONSTRAINT ck_dsco_order_sync_status CHECK (status BETWEEN 1 AND 5)
);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_create_time
  ON dsco_order_sync (dsco_create_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_status_create_time
  ON dsco_order_sync (status, dsco_create_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_dsco_order_id
  ON dsco_order_sync (dsco_order_id);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_channel
  ON dsco_order_sync (channel);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_warehouse_id
  ON dsco_order_sync (warehouse_id);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_shipped_tracking_no
  ON dsco_order_sync (shipped_tracking_no);

CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_dsco_invoice_id
  ON dsco_order_sync (dsco_invoice_id);

-- 支持按 SKU 筛选：WHERE 'sku123' = ANY(mskus)
CREATE INDEX IF NOT EXISTS idx_dsco_order_sync_mskus_gin
  ON dsco_order_sync USING GIN (mskus);

COMMENT ON TABLE dsco_order_sync IS 'DSCO 订单同步状态表（状态机 1~5 + 原始 payload；允许覆盖更新）';
COMMENT ON COLUMN dsco_order_sync.po_number IS 'DSCO 订单唯一键（po_number；同时也是领星 platform_order_no）';
COMMENT ON COLUMN dsco_order_sync.dsco_order_id IS 'DSCO 订单 ID（用于查询/对账）';
COMMENT ON COLUMN dsco_order_sync.consumer_order_number IS '参考信息：consumerOrderNumber（用于查询/对账）';
COMMENT ON COLUMN dsco_order_sync.channel IS '渠道/店铺标识（用于 mapping.shop 与查询筛选）';
COMMENT ON COLUMN dsco_order_sync.dsco_create_time IS 'DSCO 订单 created_at（UTC 秒级时间戳；用于增量拉单与筛选）';
COMMENT ON COLUMN dsco_order_sync.status IS '同步状态：1待同步（推单到领星）2待确认（回传 ack）3待发货回传（已确认）4待发票回传（已发货）5完成（已回传发票）';
COMMENT ON COLUMN dsco_order_sync.payload IS 'DSCO 原始订单 JSON（覆盖更新）';
COMMENT ON COLUMN dsco_order_sync.mskus IS '订单 SKU 列表（用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.warehouse_id IS 'DSCO 仓库编码（warehouseCode；用于 mapping.warehouse 与筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.shipment IS 'DSCO 物流方式编码（shipMethod；用于 mapping.shipment 与筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.dsco_retainer_id IS '订单关联的 dsco 标记销售渠道 ID 信息';
COMMENT ON COLUMN dsco_order_sync.shipped_tracking_no IS '已回传的物流运单号（trackingNumber；用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.dsco_invoice_id IS '已回传的发票 ID（invoiceId；用于筛选/导出）';
COMMENT ON COLUMN dsco_order_sync.created_at IS '创建时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_order_sync.updated_at IS '更新时间（UTC 秒级时间戳）';

-- DSCO 仓库库存同步日志表（用于查询与导出；不做 job_run 记录）
CREATE TABLE IF NOT EXISTS dsco_warehouse_sync (
  id                     bigserial PRIMARY KEY,

  sync_time              bigint   NOT NULL,

  dsco_warehouse_id      text     NOT NULL DEFAULT '',
  dsco_warehouse_sku     text     NOT NULL DEFAULT '',
  dsco_warehouse_num     integer  NOT NULL DEFAULT 0,

  lingxing_warehouse_id  text     NOT NULL DEFAULT '',
  lingxing_warehouse_sku text     NOT NULL DEFAULT '',
  lingxing_warehouse_num integer  NOT NULL DEFAULT 0,

  status                 smallint NOT NULL,
  reason                 text     NOT NULL DEFAULT '',

  created_at             bigint   NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),
  updated_at             bigint   NOT NULL DEFAULT (EXTRACT(EPOCH FROM now())::bigint),

  CONSTRAINT ck_dsco_warehouse_sync_status CHECK (status IN (1,2))
);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_time
  ON dsco_warehouse_sync (sync_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_dsco_wh_sku_time
  ON dsco_warehouse_sync (dsco_warehouse_id, dsco_warehouse_sku, sync_time DESC);

CREATE INDEX IF NOT EXISTS idx_dsco_warehouse_sync_status_time
  ON dsco_warehouse_sync (status, sync_time DESC);

COMMENT ON TABLE dsco_warehouse_sync IS '库存同步日志（领星库存 → DSCO 库存回写的记录，用于查询与导出）';
COMMENT ON COLUMN dsco_warehouse_sync.sync_time IS '同步时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_id IS 'DSCO 仓库编码（warehouseCode）';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_sku IS 'DSCO SKU（partnerSku 或 sku；用于 DSCO 入参）';
COMMENT ON COLUMN dsco_warehouse_sync.dsco_warehouse_num IS 'DSCO 数量（回写数量）';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_id IS '领星仓库 WID';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_sku IS '领星 SKU';
COMMENT ON COLUMN dsco_warehouse_sync.lingxing_warehouse_num IS '领星数量（可用库存）';
COMMENT ON COLUMN dsco_warehouse_sync.status IS '同步状态：1成功 2失败';
COMMENT ON COLUMN dsco_warehouse_sync.reason IS '失败原因（一期默认不写入，仅保留字段扩展位）';
COMMENT ON COLUMN dsco_warehouse_sync.created_at IS '创建时间（UTC 秒级时间戳）';
COMMENT ON COLUMN dsco_warehouse_sync.updated_at IS '更新时间（UTC 秒级时间戳）';

