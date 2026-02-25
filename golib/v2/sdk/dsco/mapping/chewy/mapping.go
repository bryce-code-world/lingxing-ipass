package chewy

import "fmt"

// CancelReason 表示后端内部使用的订单取消原因枚举（Chewy 场景）。
type CancelReason string

const (
	// CancelOutOfStock 表示商品缺货导致无法履约。
	CancelOutOfStock CancelReason = "OUT_OF_STOCK"

	// CancelDiscontinued 表示商品已停产 / 下架。
	CancelDiscontinued CancelReason = "DISCONTINUED"

	// CancelByRetailer 表示零售商（Chewy）主动取消订单。
	CancelByRetailer CancelReason = "CANCELLED_BY_RETAILER"

	// CancelInsufficientStock 表示库存不足（区别于完全无库存）。
	CancelInsufficientStock CancelReason = "INSUFFICIENT_STOCK"
)

// CancelReasonToDSCO 用于将内部取消原因映射为 DSCO / Chewy 接受的取消码。
var CancelReasonToDSCO = map[CancelReason]string{
	// CXSS：Out of Stock
	CancelOutOfStock: "CXSS",

	// CXSD：Discontinued
	CancelDiscontinued: "CXSD",

	// CXSC：Cancelled at Retailer Request
	CancelByRetailer: "CXSC",

	// CXSN：Not Enough Stock
	CancelInsufficientStock: "CXSN",
}

// Carrier 表示后端支持的物流承运商（Chewy 场景）。
type Carrier string

const (
	// CarrierFedEx 表示联邦快递。
	CarrierFedEx Carrier = "FEDEX"

	// CarrierUSPS 表示美国邮政。
	CarrierUSPS Carrier = "USPS"
)

// CarrierToDSCO 将内部 Carrier 映射为 DSCO 识别的承运商字符串。
var CarrierToDSCO = map[Carrier]string{
	// DSCO 期望的 FedEx 表示形式
	CarrierFedEx: "FedEx",

	// DSCO 期望的 USPS 表示形式
	CarrierUSPS: "USPS",
}

// ServiceLevel 表示后端内部使用的物流服务等级（Chewy 场景）。
type ServiceLevel string

const (
	// ---------- FedEx ----------

	// FedEx2Day 表示 FedEx 两日达。
	FedEx2Day ServiceLevel = "FEDEX_2_DAY"

	// FedExHomeDelivery 表示 FedEx 住宅配送。
	FedExHomeDelivery ServiceLevel = "FEDEX_HOME_DELIVERY"

	// FedExGround 表示 FedEx Ground（陆运）。
	FedExGround ServiceLevel = "FEDEX_GROUND"

	// FedExExpressSaver 表示 FedEx Express Saver。
	FedExExpressSaver ServiceLevel = "FEDEX_EXPRESS_SAVER"

	// FedExFirstOvernight 表示 FedEx First Overnight。
	FedExFirstOvernight ServiceLevel = "FEDEX_FIRST_OVERNIGHT"

	// FedExPriorityOvernight 表示 FedEx Priority Overnight。
	FedExPriorityOvernight ServiceLevel = "FEDEX_PRIORITY_OVERNIGHT"

	// FedExStandardOvernight 表示 FedEx Standard Overnight。
	FedExStandardOvernight ServiceLevel = "FEDEX_STANDARD_OVERNIGHT"

	// ---------- USPS ----------

	// USPSFirstClassMail 表示 USPS First-Class Mail。
	USPSFirstClassMail ServiceLevel = "USPS_FIRST_CLASS_MAIL"

	// USPSPriorityMail 表示 USPS Priority Mail。
	USPSPriorityMail ServiceLevel = "USPS_PRIORITY_MAIL"

	// USPSPriorityMailExpress 表示 USPS Priority Mail Express。
	USPSPriorityMailExpress ServiceLevel = "USPS_PRIORITY_MAIL_EXPRESS"

	// USPSGround 表示 USPS Ground。
	USPSGround ServiceLevel = "USPS_GROUND"
)

// ServiceLevelToDSCO 将内部物流服务等级映射为 DSCO 的 service code。
var ServiceLevelToDSCO = map[ServiceLevel]string{
	// ---------- FedEx ----------
	FedEx2Day:              "FE2D", // FedEx 2 Day
	FedExHomeDelivery:      "FEHD", // FedEx Home Delivery
	FedExGround:            "FECG", // FedEx Ground
	FedExExpressSaver:      "FEES", // FedEx Express Saver
	FedExFirstOvernight:    "FEFO", // FedEx First Overnight
	FedExPriorityOvernight: "FEPO", // FedEx Priority Overnight
	FedExStandardOvernight: "FESO", // FedEx Standard Overnight

	// ---------- USPS ----------
	USPSFirstClassMail:      "USFC", // USPS First-Class Mail
	USPSPriorityMail:        "USPM", // USPS Priority Mail
	USPSPriorityMailExpress: "USPE", // USPS Priority Mail Express
	USPSGround:              "USGG", // USPS Ground
}

// ShipmentRequest 表示后端内部的发货请求模型（Chewy 场景）。
type ShipmentRequest struct {
	Carrier        Carrier
	ServiceLevel   ServiceLevel
	TrackingNumber string
}

// BuildDSCOShipment 将内部发货请求转换为 DSCO 可接受的字段。
//
// 约定：service level 未知时，按文档策略透传（返回空的 serviceCode），便于测试期逐步补齐映射。
func BuildDSCOShipment(req ShipmentRequest) (carrier string, serviceCode string, err error) {
	c, ok := CarrierToDSCO[req.Carrier]
	if !ok {
		return "", "", fmt.Errorf("invalid carrier: %s", req.Carrier)
	}

	s, ok := ServiceLevelToDSCO[req.ServiceLevel]
	if !ok {
		return c, "", nil
	}
	return c, s, nil
}

// GetDSCOCancelCode 根据内部取消原因获取 DSCO 取消码。
func GetDSCOCancelCode(reason CancelReason) (string, error) {
	code, ok := CancelReasonToDSCO[reason]
	if !ok {
		return "", fmt.Errorf("invalid cancel reason: %s", reason)
	}
	return code, nil
}
