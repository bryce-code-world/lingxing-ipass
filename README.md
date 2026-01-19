## 项目简介

本项目是 DSCO 与领星之间的数据自动化通道（一期），以领星作为业务数据中心：
- 从 DSCO 增量拉取订单并同步到领星；
- 跟踪领星订单状态，向 DSCO 回传 ACK、发货、发票；
- 从领星拉取可用库存并同步到 DSCO。

一期原则：
- 最小存储：不做“全量业务翻译”，只落最小状态/水位/人工队列；同时保留 DSCO 原始订单快照用于回溯审计；
- 幂等：以 `dscoOrderId` 为订单幂等主键；
- 失败可追踪：全链路 `trace_id` + JSON Lines 落盘日志；
- 可运维：提供最小 HTTP 管理端用于改水位、查人工任务、手动触发 job。
- 可视化：提供 Admin 管理后台（页面 + API），用于长期运维排障。

相关设计文档：
- `doc/一期系统设计文档/1. 一期业务需求（DSCO-领星自动化）.md`
- `doc/一期系统设计文档/2. 一期系统设计（DSCO-领星自动化）.md`
- `doc/一期系统设计文档/3. 一期开发设计（DSCO-领星自动化）.md`

## 依赖与准备

- Go（建议 1.20+）
- MySQL 8+（或兼容的 MySQL）

## 数据库初始化

执行建表脚本：
- `migrations/0001_init.sql`

表说明（一期）：
- `sync_order_state`：订单闭环状态（推单/ACK/发货/发票）与幂等控制
- `job_watermark`：各任务水位（增量起点）
- `manual_task`：人工处理队列（多包裹/缺映射/坏数据等）
- `dsco_order_raw`：DSCO 原始订单快照（仅保留“最新一份”，用于回溯审计/排查）

## 配置

程序启动时会自动加载项目根目录 `.env`（不覆盖已存在的环境变量）。参考示例：
- `.env.example`

关键配置项（节选）：

系统配置（基础设施/可靠性/外部系统）：
- `IPASS_DB_DSN`：MySQL DSN
- `IPASS_LOG_DIR`：日志目录（默认 `logs/`）
- `IPASS_HTTP_ENABLE` / `IPASS_HTTP_ADDR`：HTTP 管理端开关与监听地址
- `IPASS_ADMIN_PASSWORD`：Admin 后台访问密码（一期最小：单密码；用于 UI 与管理 API）
- `IPASS_DSCO_BASE_URL` / `IPASS_DSCO_TOKEN`：DSCO API 配置（一期使用请求头 bearer token）
- `IPASS_LINGXING_BASE_URL` / `IPASS_LINGXING_APP_ID` / `IPASS_LINGXING_ACCESS_TOKEN`：领星 OpenAPI 配置
- `IPASS_LINGXING_PLATFORM_CODE` / `IPASS_LINGXING_STORE_ID`：推单固定参数（一期写死在配置）
- `IPASS_LINGXING_SID`：WMS 出库单查询参数（`sid_arr`）
- `IPASS_MAX_RETRY_PER_ORDER`：同一 dscoOrderId 单环节最大重试次数（默认 5，达到上限转人工）

业务配置（口径/映射/可随业务调整）：
- `IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON`：库存同步必填映射（领星 WID -> DSCO warehouseCode）
- `IPASS_STOCK_SKU_TO_DSCO_SKU_JSON`：库存同步可选映射（领星 SKU -> DSCO SKU）
- `IPASS_SHIP_DATE_SOURCE`：发货回传 shipDate 取值来源（`delivered_at`/`stock_delivered_at`/`none`）

## 业务映射如何注入（保姆级）

一期的“业务映射”不做复杂后台配置，统一通过 `.env` 注入（环境变量），优点是简单、可审计、可复现。

### 1) 库存：领星仓库 WID -> DSCO warehouseCode（必填）

配置项：`IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON`

说明：
- Key：领星仓库 WID（注意：这里按字符串处理，例如 `"26"`）
- Value：DSCO 仓库编码 `warehouseCode`（例如 `"WH1"`）
- 该映射仅用于 `sync_stock` 任务

示例（`.env`）：

```dotenv
IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON="{\"26\":\"WH1\",\"27\":\"WH2\"}"
```

### 2) 库存：领星 SKU -> DSCO SKU（可选）

配置项：`IPASS_STOCK_SKU_TO_DSCO_SKU_JSON`

说明：
- Key：领星侧 SKU（字符串）
- Value：DSCO 侧 SKU（字符串）
- 不配置时默认“同名直传”（领星 SKU 原样作为 DSCO SKU）

示例（`.env`）：

```dotenv
IPASS_STOCK_SKU_TO_DSCO_SKU_JSON="{\"LXSKU-1\":\"DSCOSKU-1\"}"
```

### 3) 发货回传 shipDate 口径（可选但建议明确）

配置项：`IPASS_SHIP_DATE_SOURCE`

可选值：
- `delivered_at`：使用领星 WMS 的 `delivered_at` 作为 DSCO `shipDate`（默认）
- `stock_delivered_at`：使用领星 WMS 的 `stock_delivered_at`
- `none`：不回传 `shipDate`

示例（`.env`）：

```dotenv
IPASS_SHIP_DATE_SOURCE="delivered_at"
```

### 4) 如何验证映射生效

- 仓库映射：运行 `POST /admin/run?job=sync_stock` 后，DSCO 库存会落到对应 `warehouseCode`（若缺映射会写入 `manual_task(task_type=missing_mapping)`）。
- SKU 映射：同上；若 DSCO 库存更新的 SKU 与预期一致则表示生效（不配置则应与领星 SKU 一致）。
- shipDate 口径：运行 `POST /admin/run?job=ship_to_dsco` 后，在 DSCO 侧查看发货记录的 `shipDate` 是否符合预期。

任务开关与间隔（示例）：
- `IPASS_JOB_PULL_DSCO_ORDERS_ENABLE` / `IPASS_JOB_PULL_DSCO_ORDERS_INTERVAL_SEC`
- `IPASS_JOB_PUSH_ORDERS_TO_LINGXING_ENABLE` / `IPASS_JOB_PUSH_ORDERS_TO_LINGXING_INTERVAL_SEC` / `IPASS_JOB_PUSH_ORDERS_TO_LINGXING_BATCH_SIZE`
- `IPASS_JOB_ACK_TO_DSCO_ENABLE` / `IPASS_JOB_ACK_TO_DSCO_INTERVAL_SEC`
- `IPASS_JOB_SHIP_TO_DSCO_ENABLE` / `IPASS_JOB_SHIP_TO_DSCO_INTERVAL_SEC`
- `IPASS_JOB_INVOICE_TO_DSCO_ENABLE` / `IPASS_JOB_INVOICE_TO_DSCO_INTERVAL_SEC` / `IPASS_JOB_INVOICE_TO_DSCO_BATCH_SIZE`
- `IPASS_JOB_SYNC_STOCK_ENABLE` / `IPASS_JOB_SYNC_STOCK_INTERVAL_SEC` / `IPASS_JOB_SYNC_STOCK_BATCH_SIZE`

## 启动

在项目根目录执行：
- `go run ./cmd/ipass/main.go`

## HTTP 管理端（一期最小运维面）

启用条件：`IPASS_HTTP_ENABLE=true`，默认监听 `IPASS_HTTP_ADDR=":8080"`。

接口清单：
- `GET /healthz`：健康检查
- `POST /admin/run?job=...`：手动触发一次指定 job（同步执行）
- `GET /admin/order_state/get?dsco_order_id=...`：按 dscoOrderId 查询订单闭环状态（含 last_error、retry_count 等）
- `GET /admin/order_states?push_status=&ack_status=&ship_status=&invoice_status=&limit=&offset=`：按状态筛选列表（最小分页）
- `GET /admin/watermark/get?job=...`：获取某个 job 的水位（返回 JSON）
- `POST /admin/watermark/set?job=...`：设置某个 job 的水位（请求体为 JSON，直接写入 `job_watermark.watermark`）
- `GET /admin/manual_tasks?status=0&limit=50&offset=0`：查看人工任务队列

认证：
- 如果配置了 `IPASS_ADMIN_PASSWORD`，访问上述 `/admin/*` 接口需要提供认证：
  - 浏览器：先访问 `GET /admin/ui/login` 登录（session cookie）
  - 脚本/命令行：请求头携带 `X-Admin-Password: <IPASS_ADMIN_PASSWORD>`

## Admin 管理后台（可视化）

定位：用于运维与排障，不承载业务主流程；页面只读取本系统数据库并调用本系统管理 API。

入口：
- 登录页：`GET /admin/ui/login`
- 总览：`GET /admin/ui/`
- 订单：`GET /admin/ui/orders`、`GET /admin/ui/order?dsco_order_id=...`
- 人工任务：`GET /admin/ui/manual_tasks`
- 水位：`GET /admin/ui/watermarks`
- 手动触发：`GET /admin/ui/jobs`

认证口径：
- “每次打开浏览器/新会话登录一次”：登录成功后写入 session cookie（浏览器关闭即失效）
- 命令行/脚本访问 JSON API：请求头 `X-Admin-Password`（便于 curl/自动化）

代码位置（与业务代码分离）：
- `admin/adminweb/`：Admin UI + Admin API + 鉴权与静态资源（Gin + `html/template`）

一期 job 名称：
- `pull_dsco_orders`
- `push_orders_to_lingxing`
- `ack_to_dsco`
- `ship_to_dsco`
- `invoice_to_dsco`
- `sync_stock`

说明：
- 水位默认从 `0` 开始；当领星订单列表水位为 `0` 时，会按“最近 30 天”裁剪查询窗口（领星接口时间跨度限制）。
- HTTP 请求可通过请求头 `X-Trace-Id` 传入 trace id；不传则自动生成。
- 当某个环节对同一 `dscoOrderId` 的处理失败次数达到 `IPASS_MAX_RETRY_PER_ORDER`，系统会创建 `manual_task(task_type=max_retry_exceeded)` 并将该环节状态置为人工（3）。

## 日志

默认写入 `logs/` 目录，按 JSON Lines 落盘，文件包括：
- `logs/info.log`
- `logs/debug.log`
- `logs/warn.log`
- `logs/error.log`

核心字段：
- `trace_id`：贯穿一次 job 执行链路
- `job`：任务名
- `dsco_order_id`：涉及的订单（如有）
- `err`：错误信息（如有）
