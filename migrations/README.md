# 数据库迁移（一期）

一期只提供最小的建表 SQL：`migrations/0001_init.sql`。

执行方式（示例）：

1. 确保已创建数据库并具备权限
2. 执行 `0001_init.sql`

注意：
- 一期不内置迁移工具，避免引入额外复杂度；由部署脚本或 DBA 执行即可。

## 初始水位（默认从 0 开始，可随时修改）

一期默认“起始水位从 0 开始”：
- `pull_dsco_orders`：默认 `since=1970-01-01T00:00:00Z`
- `ack_to_dsco/ship_to_dsco`：默认 `since=0`（首次运行会按领星 30 天限制裁剪查询窗口）

上线后如需纠偏/回溯，通过 HTTP 管理端口调整水位即可：
- `POST /admin/watermark/set?job=pull_dsco_orders`（Body：`{"mode":"updatedSince","since":"2025-01-01T00:00:00Z"}`）
- `POST /admin/watermark/set?job=ack_to_dsco`（Body：`{"mode":"update_time","since":1710925191}`）
- `POST /admin/watermark/set?job=ship_to_dsco`（Body：`{"mode":"update_time","since":1710925191}`）
