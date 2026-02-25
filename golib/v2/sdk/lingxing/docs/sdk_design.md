# 领星（LingXing）Go SDK 方案梳理（对齐 `golibv2/v2/sdk/dsco` 风格）

## 目标与边界

- 目标：提供一个**工程化、可维护、易用**的 Go SDK，用于调用 `https://openapi.lingxing.com` 下的开放接口。
- 风格对齐：整体结构参考 `golibv2/v2/sdk/dsco`（`Client + request.do + error + 分组 Service + models`）。
- 边界：本方案只做 **HTTP 请求封装 + 签名/鉴权 + 统一错误处理 + 结构化入参/出参**；不引入复杂的自动重试、自动续约、队列限流等“平台能力”，避免过度设计。

## 领星接口的关键特性（来自文档）

### 1) 域名与路径

- Domain：`https://openapi.lingxing.com`
- API Path：不同业务模块分散在不同 path 前缀下（如 `basicOpen/...`、`amzStaServer/...`、`erp/...` 等）。

### 2) 鉴权与公共参数

业务接口（非获取 token 接口）都要求携带公共参数：

- `access_token`：通过“获取接口令牌”接口获得
- `app_key`：AppId（文档中也称 AppID）
- `timestamp`：10 位 Unix 秒时间戳
- `sign`：签名（需要 URL 编码）

### 3) `sign` 生成规则（重点）

文档规则（原文语义整理）：

1. 参与签名的参数：**所有业务请求入参 + 3 个固定参数**（`access_token`、`app_key`、`timestamp`）
2. 按 **ASCII** 排序
3. 拼接为 `k1=v1&k2=v2&...`：
   - `value` 为空字符串：**不参与签名**
   - `value` 为 `null`：**参与签名**
4. 对拼接串进行 **MD5(32位)**，并转 **大写**
5. 再用 **AES/ECB/PKCS5PADDING** 加密该 MD5（密钥为 `appId`）
6. `sign` 作为参数传输时必须 **URL Encode**

补充注意点（来自常见问题）：

- POST 请求：URL 上**只能**带 `app_key/access_token/timestamp/sign` 四个公共参数；业务参数放 JSON body。
- body 中若嵌套集合（数组/List），**参与签名时需要转成 string**；签名完成后再用原类型回填，避免“参数不合法”。
- `sign` 有效期约 2 分钟；应使用实时 `timestamp` 生成，避免缓存 `sign`。

### 4) 返回格式不统一（必须在 SDK 层统一）

不同接口返回字段存在差异（示例）：

- `code`：可能是 `0`（成功）或 `"200"`（成功），且类型可能是数字/字符串
- 消息字段：`msg` 或 `message`
- 链路字段：`request_id` 或 `requestId`
- 错误详情：`error_details` 或 `errorDetails`

因此 SDK 需要一个“兼容多种字段名”的统一响应壳，避免业务侧重复兼容。

## SDK 包结构（建议落地形态）

目录建议：`golibv2/v2/sdk/lingxing`

```
golibv2/v2/sdk/lingxing/
  client.go           // New(Config) 与服务分组挂载（对齐 dsco）
  request.go          // do + buildURL + 编码 + 签名 + form/json 发送
  sign.go             // 签名生成（MD5 + AES/ECB/PKCS5PADDING）
  error.go            // APIError（含 http status、code、request_id、body）
  response.go         // 通用响应壳（兼容 msg/message 等）

  authorization.go        // 获取 token、续约 token（multipart/form-data）
  authorization_models.go // Token 结构等

  basicdata.go            // 基础数据分组
  basicdata_models.go

  fba.go                  // FBA / STA 相关接口分组
  fba_models.go

  ...（按文档分类继续扩展）
```

分组 Service 的风格建议完全对齐 dsco：

- `type Client struct { BasicData *BasicDataService; FBA *FBAService; Authorization *AuthorizationService; ... }`
- 每个 service 结构体只持有 `c *Client`
- 每个业务接口方法尽量只做 3 件事：组装入参 -> 调用 `c.doXXX` -> 返回结构化结果/错误

## 公共配置（Config）与 Client 设计（建议）

参考 dsco 的最小集合：

- `BaseURL`：默认 `https://openapi.lingxing.com`
- `HTTPClient`：可注入，默认超时 10s（或与现有工程一致）
- `AppID`（即 `app_key`）
- `AppSecret`（仅用于获取 token/续约，业务请求签名使用 `AppID` 作为 AES key）
- `AccessToken`：业务请求使用
- `UserAgent`：可选
- `Now func() time.Time`：用于 `timestamp`（方便测试与可控）

约束建议：

- `AppID` 为空：业务请求直接返回错误（缺少 app_key）
- `AccessToken` 为空：业务请求直接返回错误（缺少 access_token）
- AES key 长度必须满足 `16/24/32` 字节；否则签名函数返回明确错误，避免“悄悄 padding”导致线上不可控

## 请求封装（request.go）建议

对齐 dsco 的 `requestSpec` 思路，但领星需要额外支持：

1. **Signed Query（URL 上四个公共参数）**
2. **POST JSON body + 参与签名的 body 参数**
3. **multipart/form-data**（获取 token、续约 token）

建议提供三类内部方法（都返回 `error`）：

- `doSignedJSON(ctx, method, path string, query any, body any, out any)`
  - 自动填充 `access_token/app_key/timestamp/sign`
  - 仅将四个公共参数放入 URL Query
  - body 以 JSON 发送（`Content-Type: application/json`）
  - `sign` 的入参 = `query(业务) + body(业务) + 公共三参数`
- `doSignedGET(ctx, path string, query any, out any)`
  - GET 时业务参数直接拼 query（除四个公共参数外也在 URL 上）
  - `sign` 的入参 = `query(业务) + 公共三参数`
- `doMultipartForm(ctx, path string, form url.Values, out any)`
  - 用于 `/api/auth-server/oauth/access-token`、`/api/auth-server/oauth/refresh`

签名入参的构造建议（务实版本）：

- 以“参数键值对”视角构造 `map[string]any`
- 标量（string/number/bool）：直接转字符串
- `nil`：视为 `null`，参与签名（即 value 为 `"null"` 或空实现需与官方一致；这里必须在实现阶段通过最小接口联调确定）
- slice/array/map/struct：参与签名时用 JSON 字符串表示（满足“集合转 string”要求）

## 响应与错误处理（response.go / error.go）建议

### 统一响应壳

提供一个通用结构（泛型/非泛型都可，按现有工程版本选择）来兼容字段名差异：

- `code`：兼容 string/number
- `msg`/`message`
- `request_id`/`requestId`
- `error_details`/`errorDetails`
- `data`
- `total`（部分接口存在）

并提供统一判断：

- `code == 0` 或 `code == 200` 视为成功（兼容 token 接口与业务接口差异）

### 错误类型

建议错误类型包含：

- HTTP `StatusCode`
- 业务 `code`
- `message/msg`
- `request_id`
- `body`（原始响应，便于排查）

并提供 `Error()` 输出中包含 `method + url + code + request_id`，便于日志定位。

## 模块划分建议（按文档分类）

`golibv2/v2/sdk/lingxing/docs` 的文件命名已体现分类，可按以下 service 维度逐步落地：

- `AuthorizationService`：获取 token、续约 token
- `BasicDataService`：基础数据（如市场、店铺、币种、州省编码等）
- `FBAService`：FBA/STA（`amzStaServer/openapi/...` 等）
- `FBASugService`、`FBALimitService`：按文档分类拆分（如果接口量大，拆分更利于维护）
- `FinanceService`：财务/利润报表
- `WarehouseService`：仓库/出入库
- 其他：按文档目录继续添加（销售、产品、客服、多平台等）

## 使用方式（期望的调用体验）

示例（伪代码，仅表达 SDK 形态）：

```go
cfg := lingxing.Config{
	BaseURL:     lingxing.BaseURLProd,
	AppID:       "ak_xxx",
	AppSecret:   "xxx",
	AccessToken: "token_xxx",
}
cli, err := lingxing.New(cfg)
if err != nil { /* ... */ }

// 1) 基础数据：获取州/省编码（POST + JSON body）
resp, err := cli.BasicData.StateList(ctx, lingxing.StateListRequest{CountryCode: "AF"})
```

## 落地顺序（建议）

1. 先实现“基础设施”文件：`client.go/request.go/sign.go/response.go/error.go`
2. 实现 `AuthorizationService` 的两个接口（`access-token` + `refresh`），用最少联调校验签名规则与响应壳
3. 选 2~3 个代表性业务接口落地（一个 GET、一个 POST JSON、一个返回包含 `total` 的列表），把签名与响应兼容性跑通
4. 再按业务需要逐模块补齐接口与 models

