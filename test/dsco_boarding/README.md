# DSCO Onboarding（Boarding）库存测试

本目录提供 **boarding 专用的集成测试**（会真实调用 DSCO 接口）。

为避免默认 `go test ./...` 误触发，这些测试使用 build tag：`dsco_boarding`。

## 前置条件

- 已拿到 DSCO 的 bearer token
- 已确认要使用的仓库 `Code`（图 2 中的 `Code` 列，例如：`YQN-CA2` / `COOL-INTELOGICS` / `YQN-CA`）

## 配置方式

不使用环境变量；直接在测试文件里配置全局变量即可：

- `token`
- `warehouseCode`
- `skus`（3 个）

文件：`golibv2/v2/sdk/dsco/boarding_test/1.inventory_test.go`

> 注意：如果 DSCO 判断 SKU 为“新 SKU”，会要求 `upc/ean/gtin` 至少提供一个，否则会返回 400（VALIDATION_FAILED）。

## 测试方法

### 方法 1：创建 3 个 SKU 且库存=30

调用：`POST /inventory/singleItem`（同步）

PowerShell 示例：

```powershell
cd golibv2/v2
go test ./sdk/dsco/boarding_test -tags dsco_boarding -run TestBoarding_Method1 -v
```

### 方法 2：Small Batch 更新库存为 26/27/28

调用：`POST /inventory/batch/small`（异步，返回 `requestId`）

复用方法 1 使用的同一组 SKU：

```powershell
cd golibv2/v2
go test ./sdk/dsco/boarding_test -tags dsco_boarding -run TestBoarding_Method2 -v
```

> 说明：该接口为异步处理；本测试只校验返回 `requestId`，具体每条 SKU 是否成功需要后续通过 inventory change log / streams 追踪。

## 订单相关（boarding）

订单测试文件：`golibv2/v2/sdk/dsco/boarding_test/2.order_test.go`（同样使用 build tag：`dsco_boarding`）。
