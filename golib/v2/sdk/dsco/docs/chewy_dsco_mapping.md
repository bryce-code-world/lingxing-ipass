````markdown
# DSCO / Rithum × Chewy  
## 后端枚举与映射定义（Go）

本文档用于 **后端工程师** 实现 DSCO（Rithum）API 对接时的
**枚举定义、映射关系与使用规范**。

目标原则：

- 后端 **只使用内部强类型枚举**
- **不在业务代码中硬编码 DSCO 字符串**
- 所有对外值统一通过映射层转换

---

## 一、取消原因（取消码）

### 业务背景
Chewy 在取消订单时，只接受固定的取消原因码（CX**）。
DSCO 负责将你们的内部原因映射为 Chewy 可识别的值。

---

### Go 枚举定义
```go
// CancelReason 表示后端内部使用的订单取消原因枚举
type CancelReason string

const (
	// CancelOutOfStock 表示商品缺货导致无法履约
	CancelOutOfStock CancelReason = "OUT_OF_STOCK"

	// CancelDiscontinued 表示商品已停产 / 下架
	CancelDiscontinued CancelReason = "DISCONTINUED"

	// CancelByRetailer 表示零售商（Chewy）主动取消订单
	CancelByRetailer CancelReason = "CANCELLED_BY_RETAILER"

	// CancelInsufficientStock 表示库存不足（区别于完全无库存）
	CancelInsufficientStock CancelReason = "INSUFFICIENT_STOCK"
)
````

---

### CancelReason → DSCO 取消码映射

```go
// CancelReasonToDSCO 用于将内部取消原因映射为 DSCO / Chewy 接受的取消码
var CancelReasonToDSCO = map[CancelReason]string{
	// CXSS：缺货
	CancelOutOfStock: "CXSS",

	// CXSD：停产/下架
	CancelDiscontinued: "CXSD",

	// CXSC：零售商请求取消
	CancelByRetailer: "CXSC",

	// CXSN：库存不足
	CancelInsufficientStock: "CXSN",
}
```

---

## 二、承运商（物流商）

### 业务背景

DSCO 要求发货时传 **规范化的承运商名称**，
而不是你们系统内部的随意字符串。

---

### Go 枚举定义

```go
// Carrier 表示后端支持的物流承运商
type Carrier string

const (
	// CarrierFedEx 表示联邦快递
	CarrierFedEx Carrier = "FEDEX"

	// CarrierUSPS 表示美国邮政
	CarrierUSPS Carrier = "USPS"
)
```

---

### Carrier → DSCO 值映射

```go
// CarrierToDSCO 将内部 Carrier 映射为 DSCO 识别的承运商字符串
var CarrierToDSCO = map[Carrier]string{
	// DSCO 期望的 FedEx 表示形式
	CarrierFedEx: "FedEx",

	// DSCO 期望的 USPS 表示形式
	CarrierUSPS: "USPS",
}
```

---

## 三、物流服务等级（服务等级）

### 业务背景

Chewy 会指定物流服务等级（如 2 Day（两日达）、Home Delivery（住宅配送））。
DSCO 要求使用其标准化的服务码（service code）。

---

### Go 枚举定义

```go
// ServiceLevel 表示后端内部使用的物流服务等级
type ServiceLevel string

const (
	// ---------- FedEx ----------

	// FedEx2Day 表示 FedEx 两日达
	FedEx2Day ServiceLevel = "FEDEX_2_DAY"

	// FedExHomeDelivery 表示 FedEx 住宅配送
	FedExHomeDelivery ServiceLevel = "FEDEX_HOME_DELIVERY"

	// FedExGround 表示 FedEx Ground（陆运）
	FedExGround ServiceLevel = "FEDEX_GROUND"

	// FedExExpressSaver 表示 FedEx Express Saver（经济快递）
	FedExExpressSaver ServiceLevel = "FEDEX_EXPRESS_SAVER"

	// FedExFirstOvernight 表示 FedEx First Overnight（隔夜优先）
	FedExFirstOvernight ServiceLevel = "FEDEX_FIRST_OVERNIGHT"

	// FedExPriorityOvernight 表示 FedEx Priority Overnight（隔夜加急）
	FedExPriorityOvernight ServiceLevel = "FEDEX_PRIORITY_OVERNIGHT"

	// FedExStandardOvernight 表示 FedEx Standard Overnight（标准隔夜）
	FedExStandardOvernight ServiceLevel = "FEDEX_STANDARD_OVERNIGHT"

	// ---------- USPS ----------

	// USPSFirstClassMail 表示 USPS First-Class Mail（平信）
	USPSFirstClassMail ServiceLevel = "USPS_FIRST_CLASS_MAIL"

	// USPSPriorityMail 表示 USPS Priority Mail（优先邮件）
	USPSPriorityMail ServiceLevel = "USPS_PRIORITY_MAIL"

	// USPSPriorityMailExpress 表示 USPS Priority Mail Express（特快专递）
	USPSPriorityMailExpress ServiceLevel = "USPS_PRIORITY_MAIL_EXPRESS"

	// USPSGround 表示 USPS Ground（陆运）
	USPSGround ServiceLevel = "USPS_GROUND"
)
```

---

### ServiceLevel → DSCO Service Code 映射

```go
// ServiceLevelToDSCO 将内部物流服务等级映射为 DSCO 的 service code
var ServiceLevelToDSCO = map[ServiceLevel]string{
	// ---------- FedEx ----------
	FedEx2Day:              "FE2D", // FedEx 两日达
	FedExHomeDelivery:      "FEHD", // FedEx 住宅配送
	FedExGround:            "FECG", // FedEx Ground（陆运）
	FedExExpressSaver:      "FEES", // FedEx Express Saver（经济快递）
	FedExFirstOvernight:    "FEFO", // FedEx First Overnight（隔夜优先）
	FedExPriorityOvernight: "FEPO", // FedEx Priority Overnight（隔夜加急）
	FedExStandardOvernight: "FESO", // FedEx Standard Overnight（标准隔夜）

	// ---------- USPS ----------
	USPSFirstClassMail:      "USFC", // USPS First-Class Mail（平信）
	USPSPriorityMail:        "USPM", // USPS Priority Mail（优先邮件）
	USPSPriorityMailExpress: "USPE", // USPS Priority Mail Express（特快专递）
	USPSGround:              "USGG", // USPS Ground（陆运）
}
```

---

## 四、发货请求结构（后端推荐模型）

```go
// ShipmentRequest 表示后端内部的发货请求模型
type ShipmentRequest struct {
	// Carrier 承运商（FedEx / USPS）
	Carrier Carrier `json:"carrier"`

	// ServiceLevel 物流服务等级（强类型）
	ServiceLevel ServiceLevel `json:"service_level"`

	// TrackingNumber 运单号
	TrackingNumber string `json:"tracking_number"`
}
```

---

## 五、构造 DSCO 发货参数（核心函数）

```go
// BuildDSCOShipment 将内部发货请求转换为 DSCO 可接受的字段
func BuildDSCOShipment(req ShipmentRequest) (carrier string, serviceCode string, err error) {
	// 校验并映射承运商
	c, ok := CarrierToDSCO[req.Carrier]
	if !ok {
		return "", "", fmt.Errorf("无效的承运商: %s", req.Carrier)
	}

	// 校验并映射物流服务等级
	s, ok := ServiceLevelToDSCO[req.ServiceLevel]
	if !ok {
		// 与 DSCO “未知发货服务映射”策略保持一致：
		// 测试期允许未知服务等级透传（不填 serviceCode）
		return c, "", nil
	}

	return c, s, nil
}
```

---

## 六、取消订单映射函数

```go
// GetDSCOCancelCode 根据内部取消原因获取 DSCO 取消码
func GetDSCOCancelCode(reason CancelReason) (string, error) {
	code, ok := CancelReasonToDSCO[reason]
	if !ok {
		return "", fmt.Errorf("无效的取消原因: %s", reason)
	}
	return code, nil
}
```

---

## 七、工程建议（强烈）

* 所有 API 入参 **只接受枚举**
* 枚举转换失败 **在服务内拦截**
* DSCO 层永远只拿映射后的值

---

## 总结

> **DSCO 是翻译层，Go 枚举是唯一事实源。
> 只要枚举不乱，Chewy API 就不会乱。**

```
