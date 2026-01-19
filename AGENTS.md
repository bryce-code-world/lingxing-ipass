# AGENTS.md（项目级）

> 通用规范：请同时遵循本仓库根目录的 `AGENTS.base.md`（通用版 · v2.3）。  
> 项目规范：本文件补充本项目的技术事实、工程边界与交付门槛。  
> 若两者冲突，以本文件 `AGENTS.md` 为准。  
> 输出前必须检查是否存在乱码/编码问题（尤其是中文与 Markdown）。

---

## 1. 项目事实（AI 不得偏离）

- Go 后端项目，模块名：`lingxingipass`（见 `go.mod`）
- 本项目为“两进程分离”：
  - 业务进程：`cmd/ipass`（闭环调度 + ops HTTP）
  - admin 进程：`cmd/admin`（Admin UI + Admin API）
- admin 进程的 HTTP 服务：Gin（用于 Admin UI + Admin API），入口：`cmd/admin/main.go`
- 数据库：MySQL（一期），数据库访问：GORM（`gorm.io/gorm` + `gorm.io/driver/mysql`）
- 外部系统 SDK：
  - DSCO：`example.com/lingxing/golib/v2/sdk/dsco`
  - 领星：`example.com/lingxing/golib/v2/sdk/lingxing`
- 日志：统一使用 `example.com/lingxing/golib/v2/tool/logger`（本项目不再引入/复制其他日志实现）

---

## 2. 目录结构与分层边界（强约束）

### 2.1 目录约定（事实）

```
.
├── admin/
│   └── adminweb/                 # Admin 管理后台（Gin + template + 单密码鉴权）
│       ├── server.go             # 路由组装（UI + API）+ 模板初始化
│       ├── auth.go               # 单密码鉴权 + session cookie
│       ├── handlers_ui.go        # UI 页面 handler
│       ├── handlers_api.go       # JSON API handler
│       ├── handlers_static.go    # 静态资源 handler（CSS/JS）
│       ├── assets.go             # embed 资源（templates/static）
│       ├── util.go               # 小工具方法（仅限 adminweb 内部）
│       ├── static/
│       │   ├── admin.css
│       │   └── admin.js
│       └── templates/
│           ├── login.html
│           ├── dashboard.html
│           ├── orders.html
│           ├── order_detail.html
│           ├── manual_tasks.html
│           ├── watermarks.html
│           ├── jobs.html
│           └── error.html
├── cmd/
│   └── ipass/                    # 程序入口
│       └── main.go
│   └── admin/                    # admin 入口（独立进程）
│       └── main.go
├── internal/
│   ├── platform/                 # 通用能力：config/db/scheduler/retry/...
│   │   ├── config/
│   │   │   ├── config.go         # 配置结构/加载/校验
│   │   │   └── dotenv.go         # .env 加载（不覆盖已有 env）
│   │   ├── db/
│   │   │   └── db.go             # GORM MySQL 初始化
│   │   ├── scheduler/
│   │   │   └── scheduler.go      # ticker 调度器
│   │   ├── retry/
│   │   │   └── retry.go          # 重试/退避
│   │   └── adminhttp/            # 兼容旧管理接口（net/http）
│   │       └── handler.go
│   ├── store/                    # 数据访问层：最小状态库（GORM）
│   │   ├── order_state_store.go
│   │   ├── watermark_store.go
│   │   ├── manual_task_store.go
│   │   └── dsco_order_raw_store.go
│   ├── lingxing/                 # 领域层：以“领星”为中心的业务处理（映射/规则/校验）
│   │   └── order/
│   │       └── mapper_from_dsco.go
│   └── sync/                     # 编排层：把 store + SDK 串成闭环 pipeline
│       ├── order_pipeline.go
│       └── stock_pipeline.go
├── migrations/                   # 仅提供 SQL 迁移脚本（不内置迁移工具）
│   └── 0001_init.sql
└── golib/
    ├── library/                  # 旧版公共库（存量参考）
    └── v2/                       # 新版公共库（项目应优先引用）
```

### 2.2 分层调用规则（必须遵守）

允许的调用方向：

```
cmd -> internal/sync -> (internal/lingxing, internal/store, golib/v2/sdk/*) -> 外部系统
cmd/admin -> admin/adminweb -> admin/store -> MySQL
internal/platform ->（可被其他层调用，但不得反向依赖业务层）
```

禁止：
- `internal/store` 反向依赖 `internal/sync`/`internal/lingxing`
- 在业务层直接引入“新的一层”或随意新增分层（无真实需求不得加）
- 为“解耦/扩展性”引入只有一个实现的 interface
- admin 代码 import `internal/*`（必须保持 admin 与业务代码解耦）

---

## 3. 数据库与持久化（强约束）

### 3.1 GORM 使用规则

- 业务代码（含 `internal/*`、`admin/*`）禁止直接使用 `database/sql` 做 CRUD
- DB 初始化统一通过 `internal/platform/db`（返回 `*gorm.DB`）
- DB 操作统一放在 `internal/store/*`
- `internal/store` 内允许使用 GORM 的 `Exec/Raw/Transaction` 执行明确 SQL（为保持语义可控与迁移风险最低）

例外（允许）：
- 在入口处通过 `gdb.DB()` 获取底层连接用于 `Close()`（仅用于关闭资源）

### 3.2 迁移与建表

- 一期不引入迁移工具，迁移脚本位于 `migrations/0001_init.sql`
- 不得在运行时做 `AutoMigrate` 或隐式改表

---

## 4. 配置与“业务映射”注入（强约束）

### 4.1 配置加载事实

- 本项目从根目录 `.env` 加载环境变量（不覆盖进程已存在的 env）
- 配置结构与校验：`internal/platform/config`

### 4.2 业务映射注入规则（一期统一用 env JSON）

映射类配置统一通过环境变量注入 JSON 字符串，做到可复制、可审计、可回滚：

- `IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON`（必填）：领星 WID（字符串）-> DSCO `warehouseCode`
- `IPASS_STOCK_SKU_TO_DSCO_SKU_JSON`（可选）：领星 SKU -> DSCO SKU（不配默认同名直传）
- `IPASS_SHIP_DATE_SOURCE`：发货回传 `shipDate` 口径（`delivered_at`/`stock_delivered_at`/`none`）

说明：
- JSON 不合法必须在启动时直接报错（禁止带病启动）
- key/value 都按字符串处理（WID 也按字符串，例如 `"26"`）

---

## 5. 公共方法/工具方法提炼流程（强约束）

当你需要新增或复用“工具方法/公共能力”（例如：日志、时间、加解密、HTTP、并发、小工具函数等）时，必须按以下流程执行：

1. 先查旧版公共库是否已有实现：`golib/library`
2. 如果存在且可复用：
   - 先提炼/迁移到新版公共库：`golib/v2`（补齐中文注释与必要测试）
   - 本项目代码只允许引用 `golib/v2`（禁止直接从 `golib/library` 引用旧实现进入业务代码）
3. 如果旧版不存在或不适合复用：
   - 在 `golib/v2` 中新增公共方法（同样补齐中文注释与必要测试）
   - 再由本项目引用

额外约束：
- 提炼时优先“最小改动、保持语义一致”，避免顺带重构扩大风险
- 严禁在项目内复制一份“类似工具函数”造成重复与分叉

---

## 6. 测试与交付门槛（项目级）

### 6.1 测试要求

- 默认使用 Go 标准库 `testing`
- 测试与业务代码放在同一 package（避免为了测试引入多余抽象）
- 优先表驱动测试（table-driven）
- 错误判断使用 `errors.Is` / `errors.As`

### 6.2 sqlmock + GORM 绑定规则（本项目事实）

当测试需要 mock DB 时：
- 使用 `sqlmock.New()` 创建连接
- 用 `gorm.io/driver/mysql` 绑定该连接，并设置 `SkipInitializeWithVersion=true`，避免额外探测查询干扰期望

### 6.3 交付前必须通过的命令

- `go test ./...`
- `gofmt -w (git ls-files '*.go')`（PowerShell）

并发/共享状态改动（如有）额外建议：
- `go test -race ./...`

---

## 7. 文档与验收（项目级）

- 设计与验收文档目录：`doc/一期系统设计文档/`
- 任何“接口字段口径”争议以 SDK 文档为准：
  - DSCO：`golib/v2/sdk/dsco/docs/dsco-api-spec.yaml`
  - 领星：`golib/v2/sdk/lingxing/docs/*.md`
- 修改文档时必须避免乱码，且保持可复制执行的命令示例（包含 PowerShell 版本）

---

## 8. 项目级禁止事项（补充）

- 禁止引入新的 Web 框架/ORM/日志框架（除非明确需求并先确认）
- 禁止为了“未来扩展”提前引入抽象层或模式堆叠
- 禁止在业务代码里直接使用 `database/sql` 做 CRUD（必须走 GORM + store）
