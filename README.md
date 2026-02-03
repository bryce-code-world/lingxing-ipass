# lingxing-ipass

DSCO ↔ 领星（LingXing）的同步服务（一期）。单进程运行：一个 `main.go` 同时启动 HTTP（Admin UI + Admin API + healthz/readyz）与定时调度（cron，启用 seconds，固定 UTC）。

本 README 只覆盖“从 0 到 1 跑起来 / 运维使用”的信息；详细业务口径与字段映射请查看 `doc/DSCO同步领星设计文档/`。

## 你会得到什么

- Admin UI + Admin API：统一挂载在 `/admin/*`
- 健康检查：`GET /healthz`（存活）与 `GET /readyz`（runtime config + DB ping）
- 运行时配置：存储在 PostgreSQL 的 `runtime_config` 表（Admin 保存即生效；任务执行读取配置快照）

## 快速开始（本地运行）

### 0) 前置条件

- Go：本项目 `go.mod` 为 `go 1.23`（建议安装 Go 1.23+；或设置 `GOTOOLCHAIN=auto`）
- PostgreSQL：建议 14+（本项目使用 GORM + Postgres driver）
- `golibv2`：本仓库通过 `go.mod replace` 指向 `./golibv2/v2`，而 `./golibv2` 在当前环境是 **Junction**（不在 git 内）。你需要自行准备 `golibv2` 源码并在本仓库下创建同名链接目录。

### 1) 拉取仓库

```bash
git clone <your-repo-url> lingxing-ipass
cd lingxing-ipass
```

### 2) 准备 `golibv2`（必须）

确保你的本机已有 `golibv2` 代码目录（例如 `D:\src\golibv2`），然后在本仓库根目录创建链接：

- Windows（PowerShell，管理员权限通常不需要）：

```powershell
New-Item -ItemType Junction -Path .\golibv2 -Target D:\src\golibv2
```

- Windows（cmd）：

```bat
mklink /J golibv2 D:\src\golibv2
```

- Linux/macOS（符号链接）：
```bash
ln -s /path/to/golibv2 ./golibv2
```

完成后应满足：`./golibv2/v2` 存在，否则 `go mod` 会报 `replacement directory ./golibv2/v2 does not exist`。

### 3) 初始化数据库（PostgreSQL）

本项目不引入迁移工具，只维护最新建表脚本：`migrations/init.sql`。

1) 创建数据库（示例名：`lingxing_ipass`）：

```sql
CREATE DATABASE lingxing_ipass;
```

2) 执行建表脚本：

```bash
psql -h 127.0.0.1 -U root -d lingxing_ipass -f migrations/init.sql
```

### 4) 配置 `env.yaml`

程序启动时读取仓库根目录 `env.yaml`，并支持用环境变量覆盖部分字段（见 `infra/config/envyaml.go` 的 `IPASS_*` 覆盖项）。

推荐做法：从模板开始复制一份并填写敏感信息：

```bash
cp env.yaml.example env.yaml
```

最少需要关注：
- `db.dsn` 或 `db.postgres.*`（连接到上一步的 Postgres）
- `db.postgres.dbName`（需与你创建的数据库一致；仓库内常用示例为 `lingxing_ipass`）
- `admin.password`（Admin 登录密码）
- `auth.dsco.token`、`auth.lingxing.app_id/app_secret`
- `integration.lingxing.platform_code`（必填）

### 5) 启动

```bash
go test ./...
go run -buildvcs=false ./main.go
```

启动后访问：
- `base.listen_addr=":8080"` 时：Admin UI `http://127.0.0.1:8080/admin`
- 健康检查：`/healthz`、`/readyz`（同端口）

## 运行（Docker：本仓库自带的开发编排）

本仓库提供了用于 Windows 开发机的示例编排：
- 服务编排：`docker/docker-compose.lingxing-ipass.service.yml`
- 一键重启脚本：`restart.ps1`
- 容器内配置示例：`docker/lingxingipass/env.yaml`

说明：这些 compose 文件包含与本机目录、外部网络（例如 `infra_default`）相关的配置；在新环境使用时需要按实际情况调整 `volumes` 与 `networks`。

示例（在本仓库根目录）：

```powershell
# 启动（或重启）服务
docker compose -p lingxing-ipass -f docker\\docker-compose.lingxing-ipass.service.yml up -d

# 查看日志
docker logs -f lingxingipass
```

## 运维说明（面向线上/长期运行）

### Admin 与安全

- Admin UI + Admin API 统一在 `/admin/*`，目前是“单密码 + session cookie”的最小实现，建议只在内网使用或置于反向代理/鉴权之后。
- `admin.password` 来自 `env.yaml`（或环境变量覆盖）。

### 定时任务与手动触发

- 调度器使用 `github.com/robfig/cron/v3`，启用 seconds，时区固定 UTC。
- 定时是否触发由 runtime config 中的 `jobs.<job>.enable` 控制；但 Admin 的手动触发不受该开关限制（用于排障/补偿）。

### 配置变更

- 运行时配置落在 `runtime_config` 表：Admin 保存即生效。
- 任务执行时读取一次“配置快照”，本次执行全程使用同一快照；下次执行再读取新快照。

### 日志与导出文件

- 日志：统一使用 `example.com/lingxing/golib/v2/tool/logger`，输出到 `log.dir`（并可同时 stdout）。
- 导出：Admin 的 CSV 导出落在 `admin.export.dir`，并提供清理任务（见 runtime config jobs）。

### 健康检查

- `GET /healthz`：仅表示 HTTP 进程存活。
- `GET /readyz`：检查 runtime config 已加载，且尽量对 DB 做一次 ping（用于容器编排 readiness）。

## 进一步阅读

- 业务/字段口径/任务定义与数据库设计：`doc/DSCO同步领星设计文档/`
