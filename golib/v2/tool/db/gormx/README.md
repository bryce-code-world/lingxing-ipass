# gormx

`gormx` 是 `golibv2/v2` 下的 GORM 工具层封装，路径为 `v2/tool/db/gormx`。

## 功能概览

- 支持 MySQL / PostgreSQL / SQLite 客户端初始化
- 支持 DSN 注入或结构化配置组装 DSN
- 支持连接池参数统一设置
- 支持启动自检（`SELECT 1`）
- 支持统一 GORM 日志适配（接入 `tool/logger`）

## 文件结构

- `mysql.go` / `mysql_test.go`：MySQL 连接初始化与测试
- `mysql_dsn.go` / `mysql_dsn_test.go`：MySQL DSN 构建与测试
- `postgres.go` / `postgres_test.go`：PostgreSQL 连接初始化与测试
- `postgres_dsn.go` / `postgres_dsn_test.go`：PostgreSQL DSN 构建与测试
- `sqlite.go` / `sqlite_test.go`：SQLite 连接初始化与测试
- `sqlite_dsn.go` / `sqlite_dsn_test.go`：SQLite DSN 构建与测试
- `config.go`：配置模型
- `client.go`：公共能力（Close、连接池、启动自检、日志）

## MySQL 示例

```go
ctx := context.Background()

gdb, err := gormx.OpenMySQL(ctx, gormx.Config{
    MySQL: gormx.MySQLConfig{
        Host:     "127.0.0.1",
        Port:     3306,
        User:     "root",
        Password: "123456",
        DBName:   "demo",
        Loc:      "UTC",
    },
    Pool: gormx.PoolConfig{
        MaxOpenConns: 20,
        MaxIdleConns: 10,
    },
})
if err != nil {
    panic(err)
}
defer gormx.Close(gdb)
```

## PostgreSQL 示例

```go
ctx := context.Background()

gdb, err := gormx.OpenPostgres(ctx, gormx.Config{
    Postgres: gormx.PostgresConfig{
        Host:     "127.0.0.1",
        Port:     5432,
        User:     "postgres",
        Password: "123456",
        DBName:   "demo",
        SSLMode:  "disable",
        TimeZone: "UTC",
    },
})
if err != nil {
    panic(err)
}
defer gormx.Close(gdb)
```

## SQLite 示例

```go
ctx := context.Background()

gdb, err := gormx.OpenSQLite(ctx, gormx.Config{
    SQLite: gormx.SQLiteConfig{
        DSN: "file::memory:?cache=shared",
        // 也可使用 Path: "./data.db"
    },
})
if err != nil {
    panic(err)
}
defer gormx.Close(gdb)
```

## 测试

在 `v2` 目录执行：

```bash
go test ./tool/db/gormx -run "SQLite|BuildSQLiteDSN"
```
