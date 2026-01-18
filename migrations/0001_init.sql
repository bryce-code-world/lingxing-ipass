-- 一期初始化建表（来自 doc/一期系统设计（DSCO-领星自动化）.md）

CREATE TABLE IF NOT EXISTS sync_order_state (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  dsco_order_id VARCHAR(64) NOT NULL COMMENT 'DSCO dscoOrderId（幂等主键）',
  lingxing_global_order_no VARCHAR(64) NULL COMMENT '领星系统单号（global_order_no），创建成功后回写',

  pushed_to_lx_status TINYINT NOT NULL DEFAULT 0 COMMENT '推单到领星状态：0未处理 1成功 2失败 3人工 9处理中',
  pushed_to_lx_at DATETIME(3) NULL,

  acked_to_dsco_status TINYINT NOT NULL DEFAULT 0 COMMENT '回ACK到DSCO：0未处理 1成功 2失败 3人工 9处理中',
  acked_to_dsco_at DATETIME(3) NULL,

  shipped_to_dsco_status TINYINT NOT NULL DEFAULT 0 COMMENT '回传发货到DSCO：0未处理 1成功 2失败 3人工 9处理中',
  shipped_to_dsco_at DATETIME(3) NULL,
  shipped_tracking_no VARCHAR(64) NULL COMMENT '已回传的跟踪号（一期只回传一个）',

  invoiced_to_dsco_status TINYINT NOT NULL DEFAULT 0 COMMENT '回传发票到DSCO：0未处理 1成功 2失败 3人工 9处理中',
  invoiced_to_dsco_at DATETIME(3) NULL,
  dsco_invoice_id VARCHAR(64) NULL COMMENT '已回传的发票号（invoiceId）',

  retry_count INT NOT NULL DEFAULT 0,
  last_error TEXT NULL,
  last_attempt_at DATETIME(3) NULL,

  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

  PRIMARY KEY (id),
  UNIQUE KEY uk_dsco_order_id (dsco_order_id),
  KEY idx_push_status (pushed_to_lx_status, updated_at),
  KEY idx_ack_status (acked_to_dsco_status, updated_at),
  KEY idx_ship_status (shipped_to_dsco_status, updated_at),
  KEY idx_invoice_status (invoiced_to_dsco_status, updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS job_watermark (
  job_name VARCHAR(64) NOT NULL COMMENT '任务名，例如 pull_dsco_orders/sync_stock',
  watermark JSON NOT NULL COMMENT '任务水位（按任务自定义结构）',
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (job_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- DSCO 订单原始快照（一期用于回溯/审计/排查：仅存最新一份，避免无限膨胀）
CREATE TABLE IF NOT EXISTS dsco_order_raw (
  dsco_order_id VARCHAR(64) NOT NULL COMMENT 'DSCO dscoOrderId',
  payload JSON NOT NULL COMMENT 'DSCO Order 原始 JSON（尽量不修改）',
  payload_sha256 CHAR(64) NOT NULL COMMENT 'payload 的 sha256（用于判断是否变化）',
  fetched_at DATETIME(3) NOT NULL COMMENT '抓取时间（UTC）',
  PRIMARY KEY (dsco_order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS manual_task (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  task_type VARCHAR(32) NOT NULL COMMENT '类型：multi_shipment/missing_mapping/bad_payload 等',
  dsco_order_id VARCHAR(64) NULL,
  payload JSON NULL COMMENT '用于排查/补处理的上下文（脱敏后）',
  status TINYINT NOT NULL DEFAULT 0 COMMENT '0待处理 1处理中 2已完成 3已忽略',
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_task_status (task_type, status, updated_at),
  KEY idx_task_order (dsco_order_id, updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
