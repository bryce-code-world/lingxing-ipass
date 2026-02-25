# lingxing-ipass

DSCO 同步领星（LingXing）服务，单进程运行（`main.go`）。

- HTTP：Gin（Admin UI + Admin API + 其他路由）
- Admin 路由前缀：`/admin/*`
- Scheduler：`github.com/robfig/cron/v3`（启用 seconds，时区 UTC）
- DB：PostgreSQL + GORM
- 本地稳定依赖：`golib/v2`（项目内包，不依赖外部 golib 模块）

## 快速开始

### 1. 环境准备

- Go 1.23+
- PostgreSQL 14+

### 2. 初始化数据库

执行建表脚本：

```bash
psql -h 127.0.0.1 -U <user> -d <db_name> -f migrations/init.sql
```

### 3. 配置

复制配置模板并修改：

```bash
cp env.yaml.example env.yaml
```

至少确认以下配置：

- `db.dsn` 或 `db.postgres.*`
- `admin.password`
- `auth.dsco.*`
- `auth.lingxing.*`
- `integration.lingxing.platform_code`

### 4. 启动

```bash
go run -buildvcs=false ./main.go
```

默认 `:8080` 时：

- Admin UI: `http://127.0.0.1:8080/admin`
- Liveness: `GET /healthz`
- Readiness: `GET /readyz`

## 测试

```bash
go test ./...
```

说明：`test/*_boarding` 为真实接口联调测试，若网络或凭证不可用可能超时失败。

## Docker（可选）

```powershell
docker compose -p lingxing-ipass -f docker\docker-compose.lingxing-ipass.service.yml up -d
docker logs -f lingxingipass
```

## 项目结构（核心）

- `main.go`：入口
- `admin/`：Admin UI + Admin API
- `transport/`：HTTP Server + Scheduler（只负责触发）
- `integration/`：任务编排与任务实现
- `infra/`：配置、DB、锁、存储、运行时配置
- `golib/v2/`：项目内稳定 SDK/工具包
- `migrations/init.sql`：最新建表脚本
- `doc/DSCO同步领星设计文档/`：设计与口径文档

## 文档

- 设计文档目录：`doc/DSCO同步领星设计文档/`
- DSCO API 规格：`golib/v2/sdk/dsco/docs/dsco-api-spec.yaml`
- 领星接口文档：`golib/v2/sdk/lingxing/docs/*.md`
