# AGENTS.md（项目级）

> 通用规范：请同时遵循本仓库根目录的 `AGENTS.base.md`（通用版 · v2.3）。  
> 项目规范：本文件补充本项目的技术事实、工程边界与交付门槛。  
> 若两者冲突，以本文件 `AGENTS.md` 为准。  
> 输出前必须检查是否存在乱码/编码问题（尤其是中文与 Markdown/SQL 注释）。

---

## 1. 项目事实（AI 不得偏离）

- Go 后端项目，模块名：`lingxingipass`（见 `go.mod`）
- 单服务/单进程：唯一入口 `main.go`（仓库根目录）
- HTTP 框架：Gin（Admin UI + Admin API + 其他 HTTP 路由统一通过一个 Gin Engine 提供）
  - Admin UI + Admin API：统一挂载 `/admin/*`
- Scheduler：`github.com/robfig/cron/v3`（启用 seconds，时区固定 UTC）
- 数据库：PostgreSQL，一期使用 GORM（`gorm.io/gorm`）
  - DB 客户端封装：`infra/db`（内部调用 `golibv2/v2/tool/db/gormx`）
- 迁移/建表：仅维护最新 `migrations/init.sql`（不引入迁移工具；禁止运行时 `AutoMigrate`）
- 外部系统 SDK（通过 `go.mod` 直接依赖线上模块 `gitee.com/lsy007/golibv2/v2`，版本以 `go.mod` 为准）：
  - DSCO：`gitee.com/lsy007/golibv2/v2/sdk/dsco`
  - 领星：`gitee.com/lsy007/golibv2/v2/sdk/lingxing`
- 日志：统一使用 `gitee.com/lsy007/golibv2/v2/tool/logger`（禁止引入/复制其他日志实现）

---

## 2. 目录结构与分层边界（强约束）

### 2.1 目录约定（事实，以当前代码为准）

```
.
├── main.go                         # 服务入口：加载 env.yaml；初始化 logger/DB/runtimecfg；启动 HTTP + scheduler
├── env.yaml                        # 基础启动配置（可被环境变量覆盖）
├── admin/                          # Admin：资源与代码统一放在此目录；对外统一 /admin 前缀
│   ├── assets.go                   # embed templates/static
│   ├── server.go                   # Gin 路由组装：/admin UI + /admin/api
│   ├── auth.go                     # 单密码鉴权 + session cookie
│   ├── ui.go                       # Admin UI 页面（dashboard/config/tasks/orders/warehouses）
│   ├── api.go                      # Admin API（runtime_config/jobs/orders/warehouses/export）
│   ├── static/
│   │   ├── admin.css
│   │   └── admin.js
│   └── templates/
│       ├── base.html
│       ├── login.html
│       ├── dashboard.html
│       ├── config.html
│       ├── tasks.html
│       ├── orders.html
│       └── warehouses.html
├── transport/                      # 协议层：HTTP Server + Scheduler（只负责触发，不写业务）
│   ├── httpserver/                 # httpserver.New(listenAddr, ginEngine, cfgMgr, gdb)
│   └── scheduler/                  # cron 调度器：读取 runtimecfg.jobs，触发 runner
├── integration/                    # 任务模块层：按领域组织任务；只允许被 runner 调用
│   ├── registry.go                 # 任务注册器
│   ├── runner.go                   # 统一编排样板：配置快照、互斥锁、日志、调用任务
│   ├── types.go                    # TaskContext/RunRequest 等
│   ├── dsco_lingxing/              # 领域：DSCO ↔ 领星
│   │   ├── register.go             # 注册 pull/push/ack/ship/invoice/stock 等任务
│   │   └── task_*.go               # 每个文件一个具体任务（禁止任务互调）
│   └── ops/                        # 运维/健康检查等任务域（如有）
├── infra/                          # 基础设施层：配置/DB/锁/运行时配置/存储
│   ├── config/                     # env.yaml 结构、加载、校验（支持 env 覆盖）
│   ├── db/                         # DB 打开/关闭（Postgres + gormx）
│   ├── lock/                       # PostgreSQL advisory lock（任务互斥）
│   ├── runtimecfg/                 # runtime_config：默认配置/校验/快照/热更新
│   ├── store/                      # 数据访问层（GORM）：runtime_config/dsco_order_sync/dsco_warehouse_sync
│   └── fileutil/                   # 文件与目录小工具（如导出目录清理）
├── migrations/
│   ├── init.sql                    # 最新建表脚本（仅此一个为准）
│   └── README.md
├── doc/DSCO同步领星设计文档/       # 设计文档（以实现为准持续对齐）
└── test/                           # 手动/boarding 测试（默认跳过真实接口调用）
```

### 2.2 分层调用规则（必须遵守）

允许的调用方向：

```
main -> (infra/*, admin/*, transport/*, integration/*)
transport/* -> integration.Runner
admin/* -> (integration.Runner, infra/runtimecfg, infra/store)
integration/* -> (infra/store, infra/runtimecfg types, golibv2 sdk/*) -> 外部系统
infra/* ->（可被其他层调用，但不得反向依赖 integration/admin/transport）
```

禁止：
- 任务模块互相调用（禁止 A 任务直接调用 B 任务），任务只允许被 `integration/runner.go` 调用
- 协议层（transport/admin）绕过 runner 直接调用任务函数
- `infra/*` 反向依赖 `integration/*` 或 `admin/*` 或 `transport/*`
- 为“未来扩展”引入只有一个实现的 interface（无真实需求不得加）

---

## 3. 数据库与持久化（强约束）

### 3.1 GORM 使用规则

- 全项目禁止直接使用 `database/sql` 做 CRUD（必须走 `infra/store` + GORM）
- 允许的例外：
  - 在入口处通过 `gdb.DB()` 取得底层连接用于 `Close()` 或 advisory lock（见 `integration/runner.go`）
- 时间类字段：DB 内统一用 UTC 秒级时间戳（`bigint`）
- 禁止运行时建表/改表：不允许 `AutoMigrate` 或隐式改表

### 3.2 迁移与建表

- 一期不引入迁移工具，仅维护最新 `migrations/init.sql`
- 任何表结构变更：同步更新 `migrations/init.sql` 与 `doc/DSCO同步领星设计文档/3.数据库设计.md`

---

## 4. 配置与运行时配置（强约束）

### 4.1 配置加载事实

- 基础启动配置：根目录 `env.yaml`（并支持环境变量覆盖同名字段）
- 运行时配置：DB 表 `runtime_config`（Admin 保存即生效）
  - 任务执行一致性：任务启动时读取一次“配置快照”，本次执行全程使用同一快照；下次执行再读取新快照
  - Scheduler 的 `jobs.<job>.enable` 仅控制“定时触发”；Admin 手动触发不受该开关限制

### 4.2 mapping 方向（写死）

- `runtime_config.config.mapping` 统一方向：**DSCO -> 领星**（key 为 DSCO 标识/编码，value 为 领星标识/编码）
- 如需反向映射（领星 -> DSCO），只允许在任务执行时在内存构建反向 map；若冲突（一对多/多对一）应视为配置错误

---

## 5. 公共方法/工具方法复用（强约束）

- 优先复用 `golibv2/v2` 的实现（本仓库通过 `go.mod` 依赖 `gitee.com/lsy007/golibv2/v2`，版本以 `go.mod` 为准）
- 禁止在项目内复制一份“类似工具函数”造成重复与分叉
- 禁止引入新的日志框架/DB 框架/HTTP 框架（除非明确需求并先确认）

---

## 6. 测试与交付门槛（项目级）

交付前必须通过：
- `go test ./...`
- `gofmt -w (git ls-files '*.go')`（PowerShell）

说明：
- `test/*_boarding` 属于“真实接口手动测试”，默认跳过，避免 `go test ./...` 误触发外部请求。

---

## 7. 文档与验收（项目级）

- 设计文档目录：`doc/DSCO同步领星设计文档/`
- 任何“接口字段口径”争议以 SDK 文档为准：
  - DSCO：`golibv2/v2/sdk/dsco/docs/dsco-api-spec.yaml`
  - 领星：`golibv2/v2/sdk/lingxing/docs/*.md`

---

## 8. 项目级禁止事项（补充）

- 禁止在业务代码里直接使用 `database/sql` 做 CRUD（必须走 GORM + store）
- 禁止引入新的 Web 框架/ORM/日志框架（除非明确需求并先确认）
- 禁止为了“未来扩展”提前引入抽象层或模式堆叠

---

## 9. `golibv2` 仓库协作规范（新增）

本项目依赖 `gitee.com/lsy007/golibv2/v2`，当出现“SDK 字段/接口缺失、口径不一致”等问题时，优先在 `golibv2` 仓库协作修复，而不是在本项目里做临时绕过。

### 9.1 字段/接口变更的判定依据（必须遵守）

- **不得凭空新增字段**：是否存在/字段名/类型/可空性，必须以对应 SDK 文档为准：
  - DSCO：`golibv2/v2/sdk/dsco/docs/dsco-api-spec.yaml`
  - 领星：`golibv2/v2/sdk/lingxing/docs/*.md`

### 9.2 标准协作流程（必须遵守）

- 在 `golibv2` 仓库中完成：模型字段/请求响应结构/接口封装的变更，并提交/合并
- **在本仓库只做版本升级**：通过更新 `go.mod` 的 `require gitee.com/lsy007/golibv2/v2 <version>` 引入变更（tag 或 pseudo-version 均可）
- 仓库内如存在 `./golibv2/` 目录，仅用于阅读文档/参与 `golibv2` 开发协作，不作为本仓库依赖来源
- 升级后必须执行：`go mod tidy`、`go test ./...`
- PR/变更说明必须写清：对应的 `golibv2` 版本号（或 pseudo-version）与关键变更点（便于回溯）

### 9.3 禁止事项（强约束）

- 禁止提交 `go.mod` 的本地替换依赖（例如 `replace gitee.com/lsy007/golibv2/v2 => ./golibv2/v2` 或任何本地路径）
  - 如确需本地调试，可在个人环境临时使用，但**提交前必须移除**，并确保 `go mod tidy` 后不残留
- 禁止在本仓库复制/粘贴一份 `golibv2` SDK 源码作为“临时修复”（避免分叉）
