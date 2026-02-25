查询店铺绩效详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/storeTarget/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
pullDate	报表更新日期，必填，日期格式：yyyy-MM-dd	是	[string]	2024-11-01
sid	店铺ID，必填	是	[long]	123456789
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/customerService/storeTarget/detail?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "pullDate": "2024-11-01",
    "sid": 123456789
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	响应码	是	[int]	
data	返回数据	是	[object]	
data>>accountHealthRating	账户健康评级对象	是	[object]	
data>>accountHealthRating>>ahrScore	账户健康评分	是	[string]	200
data>>accountHealthRating>>ahrStatus	账户健康状态	是	[string]	Healthy
data>>accountHealthRating>>windowDayCount	窗口期的天数	是	[string]	30
data>>accountHealthRating>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>accountHealthRating>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>commodityPolicyCompliance	商品政策合规性对象	是	[object]	
data>>commodityPolicyCompliance>>count	总数	是	[string]	0
data>>commodityPolicyCompliance>>customerProductReviewsPolicyViolations	违反买家商品评论政策	是	[string]	0
data>>commodityPolicyCompliance>>foodAndProductSafetyIssues	食品和商品安全问题	是	[string]	0
data>>commodityPolicyCompliance>>intellectualPropertyData	知识产权投诉	是	[string]	0
data>>commodityPolicyCompliance>>listingPolicyData	上架政策违规	是	[string]	0
data>>commodityPolicyCompliance>>otherPolicyViolations	其他违反政策	是	[string]	0
data>>commodityPolicyCompliance>>policyViolationWarnings	违反政策警告	是	[string]	0
data>>commodityPolicyCompliance>>productAuthenticityData	商品真实性买家投诉	是	[string]	0
data>>commodityPolicyCompliance>>productConditionCustomerComplaints	商品状况买家投诉	是	[string]	0
data>>commodityPolicyCompliance>>productSafetyData	商品安全投诉	是	[string]	0
data>>commodityPolicyCompliance>>restrictedProductPolicyViolations	违反受限商品政策	是	[string]	0
data>>commodityPolicyCompliance>>suspectedIntellectualPropertyViolations	涉嫌侵犯知识产权	是	[string]	0
data>>commodityPolicyCompliance>>windowDayCount	窗口期的天数	是	[string]	30
data>>commodityPolicyCompliance>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>commodityPolicyCompliance>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>fbaOrderWithDefect	FBA订单缺陷率对象	是	[object]	
data>>fbaOrderWithDefect>>child	子项详情对象	是	[object]	
data>>fbaOrderWithDefect>>child>>aZClaim	亚马逊商城交易保障索赔对象	是	[object]	
data>>fbaOrderWithDefect>>child>>aZClaim>>count	比率对应的数量	是	[string]	4
data>>fbaOrderWithDefect>>child>>aZClaim>>rate	比率，百分比	是	[string]	0.3
data>>fbaOrderWithDefect>>child>>chargeback	信用卡拒付索赔对象	是	[object]	
data>>fbaOrderWithDefect>>child>>chargeback>>count	比率对应的数量	是	[string]	1
data>>fbaOrderWithDefect>>child>>chargeback>>rate	比率，百分比	是	[string]	0.1
data>>fbaOrderWithDefect>>child>>negativeFeedback	Feedback负面反馈对象	是	[object]	
data>>fbaOrderWithDefect>>child>>negativeFeedback>>count	比率对应的数量	是	[string]	5
data>>fbaOrderWithDefect>>child>>negativeFeedback>>rate	比率，百分比	是	[string]	0.4
data>>fbaOrderWithDefect>>count	FBA订单缺陷数	是	[string]	10
data>>fbaOrderWithDefect>>orderCount	FBA订单总数	是	[string]	1250
data>>fbaOrderWithDefect>>rate	FBA订单缺陷率，百分比	是	[string]	0.8
data>>fbaOrderWithDefect>>windowDayCount	窗口期的天数	是	[string]	30
data>>fbaOrderWithDefect>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>fbaOrderWithDefect>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>invoiceDefect	发票缺陷率对象	是	[object]	
data>>invoiceDefect>>count	比率对应的数量	是	[string]	5
data>>invoiceDefect>>invoiceDefectCount	发票缺失订单数	是	[string]	5
data>>invoiceDefect>>lateInvoiceCount	逾期发票订单数	是	[string]	3
data>>invoiceDefect>>missingInvoiceCount	遗失发票订单数	是	[string]	2
data>>invoiceDefect>>orderCount	订单总数	是	[string]	1250
data>>invoiceDefect>>rate	比率，百分比	是	[string]	0.4
data>>invoiceDefect>>windowDayCount	窗口期的天数	是	[string]	30
data>>invoiceDefect>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>invoiceDefect>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>lateShipment	迟发率对象	是	[object]	
data>>lateShipment>>count	比率对应的数量	是	[string]	31
data>>lateShipment>>orderCount	订单总数	是	[string]	1250
data>>lateShipment>>rate	比率，百分比	是	[string]	2.5
data>>lateShipment>>windowDayCount	窗口期的天数	是	[string]	30
data>>lateShipment>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>lateShipment>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>onTimeDelivery	准时交货率对象	是	[object]	
data>>onTimeDelivery>>count	比率对应的数量	是	[string]	1218
data>>onTimeDelivery>>orderCount	订单总数	是	[string]	1250
data>>onTimeDelivery>>rate	比率，百分比	是	[string]	97.5
data>>onTimeDelivery>>windowDayCount	窗口期的天数	是	[string]	30
data>>onTimeDelivery>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>onTimeDelivery>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>onTimeDeliverySource	准时交货率数据来源
1：取"onTimeDelivery"
2：取"unitOnTimeDelivery"	是	[int]	1
data>>orderWithDefect	FBM订单缺陷率对象	是	[object]	
data>>orderWithDefect>>child	子项详情对象	是	[object]	
data>>orderWithDefect>>child>>aZClaim	亚马逊商城交易保障索赔对象	是	[object]	
data>>orderWithDefect>>child>>aZClaim>>count	比率对应的数量	是	[string]	4
data>>orderWithDefect>>child>>aZClaim>>rate	比率，百分比	是	[string]	0.3
data>>orderWithDefect>>child>>chargeback	信用卡拒付索赔对象	是	[object]	
data>>orderWithDefect>>child>>chargeback>>count	比率对应的数量	是	[string]	1
data>>orderWithDefect>>child>>chargeback>>rate	比率，百分比	是	[string]	0.1
data>>orderWithDefect>>child>>negativeFeedback	Feedback负面反馈对象	是	[object]	
data>>orderWithDefect>>child>>negativeFeedback>>count	比率对应的数量	是	[string]	5
data>>orderWithDefect>>child>>negativeFeedback>>rate	比率，百分比	是	[string]	0.4
data>>orderWithDefect>>count	FBM订单缺陷数	是	[string]	10
data>>orderWithDefect>>orderCount	FBM订单总数	是	[string]	1250
data>>orderWithDefect>>rate	FBM订单缺陷率，百分比	是	[string]	0.8
data>>orderWithDefect>>windowDayCount	窗口期的天数	是	[string]	30
data>>orderWithDefect>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>orderWithDefect>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>preFulfillmentCancellation	预配送取消率对象	是	[object]	
data>>preFulfillmentCancellation>>count	比率对应的数量	是	[string]	6
data>>preFulfillmentCancellation>>orderCount	订单总数	是	[string]	1250
data>>preFulfillmentCancellation>>rate	比率，百分比	是	[string]	0.5
data>>preFulfillmentCancellation>>windowDayCount	窗口期的天数	是	[string]	30
data>>preFulfillmentCancellation>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>preFulfillmentCancellation>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>pullDate	报表获取时间日期，日期格式：yyyy-MM-dd	是	[string]	2024-11-01
data>>standard	评分标准对象	是	[object]	
data>>standard>>commodityPolicyCompliance	商品政策合规性评分标准，百分比	是	[string]	0
data>>standard>>fbaOrderWithDefect	FBA订单缺陷率评分标准，百分比	是	[string]	1
data>>standard>>invoiceDefect	发票缺陷评分标准，百分比	是	[string]	5
data>>standard>>lateShipment	迟发率评分标准，百分比	是	[string]	4
data>>standard>>onTimeDelivery	准时交货率评分标准，百分比	是	[string]	90
data>>standard>>orderWithDefect	订单缺陷率评分标准，百分比	是	[string]	1
data>>standard>>preFulfillmentCancellation	预配送取消率评分标准，百分比	是	[string]	2.5
data>>standard>>returnDissatisfaction	退货不满意度评分标准，百分比	是	[string]	10
data>>standard>>unitOnTimeDelivery	单位准时交货率评分标准，百分比	是	[string]	93
data>>standard>>validTracking	有效追踪率评分标准，百分比	是	[long]	95
data>>unitOnTimeDelivery	单位准时交货率对象	是	[object]	
data>>unitOnTimeDelivery>>count	比率对应的数量	是	[string]	1187
data>>unitOnTimeDelivery>>orderCount	订单总数	是	[string]	1250
data>>unitOnTimeDelivery>>rate	比率，百分比	是	[string]	95
data>>unitOnTimeDelivery>>windowDayCount	窗口期的天数	是	[string]	30
data>>unitOnTimeDelivery>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>unitOnTimeDelivery>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
data>>updateDate	报表数据更新时间，日期格式：yyyy-MM-dd	是	[string]	2024-11-12
data>>validTracking	有效追踪率对象	是	[object]	
data>>validTracking>>count	比率对应的数量	是	[string]	1243
data>>validTracking>>orderCount	货件总数	是	[string]	1250
data>>validTracking>>rate	比率，百分比	是	[string]	99.5
data>>validTracking>>windowDayCount	窗口期的天数	是	[string]	30
data>>validTracking>>windowTimeEnd	窗口期的结束日期	是	[string]	2024-11-01
data>>validTracking>>windowTimeStart	窗口期的开始日期	是	[string]	2024-10-01
error_details	错误详情	是	[array]	
message	提示信息	是	[string]	
request_id	请求id	是	[string]	
response_time	响应时间	是	[string]	
total	总记录数	是	[int]	
返回成功示例
{
  "code": 0,
  "data": {
    "accountHealthRating": {
      "ahrScore": "200",
      "ahrStatus": "Healthy",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "commodityPolicyCompliance": {
      "count": "0",
      "customerProductReviewsPolicyViolations": "0",
      "foodAndProductSafetyIssues": "0",
      "intellectualPropertyData": "0",
      "listingPolicyData": "0",
      "otherPolicyViolations": "0",
      "policyViolationWarnings": "0",
      "productAuthenticityData": "0",
      "productConditionCustomerComplaints": "0",
      "productSafetyData": "0",
      "restrictedProductPolicyViolations": "0",
      "suspectedIntellectualPropertyViolations": "0",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "fbaOrderWithDefect": {
      "child": {
        "aZClaim": {
          "count": "4",
          "rate": "0.3"
        },
        "chargeback": {
          "count": "1",
          "rate": "0.1"
        },
        "negativeFeedback": {
          "count": "5",
          "rate": "0.4"
        }
      },
      "count": "10",
      "orderCount": "1250",
      "rate": "0.8",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "invoiceDefect": {
      "count": "5",
      "invoiceDefectCount": "5",
      "lateInvoiceCount": "3",
      "missingInvoiceCount": "2",
      "orderCount": "1250",
      "rate": "0.4",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "lateShipment": {
      "count": "31",
      "orderCount": "1250",
      "rate": "2.5",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "onTimeDelivery": {
      "count": "1218",
      "orderCount": "1250",
      "rate": "97.5",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "onTimeDeliverySource": 1,
    "orderWithDefect": {
      "child": {
        "aZClaim": {
          "count": "4",
          "rate": "0.3"
        },
        "chargeback": {
          "count": "1",
          "rate": "0.1"
        },
        "negativeFeedback": {
          "count": "5",
          "rate": "0.4"
        }
      },
      "count": "10",
      "orderCount": "1250",
      "rate": "0.8",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "preFulfillmentCancellation": {
      "count": "6",
      "orderCount": "1250",
      "rate": "0.5",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "pullDate": "2024-11-01",
    "standard": {
      "commodityPolicyCompliance": "0",
      "fbaOrderWithDefect": "1",
      "invoiceDefect": "5",
      "lateShipment": "4",
      "onTimeDelivery": "90",
      "orderWithDefect": "1",
      "preFulfillmentCancellation": "2.5",
      "returnDissatisfaction": "10",
      "unitOnTimeDelivery": "93",
      "validTracking": 95
    },
    "unitOnTimeDelivery": {
      "count": "1187",
      "orderCount": "1250",
      "rate": "95",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    },
    "updateDate": "2024-11-12",
    "validTracking": {
      "count": "1243",
      "orderCount": "1250",
      "rate": "99.5",
      "windowDayCount": "30",
      "windowTimeEnd": "2024-11-01",
      "windowTimeStart": "2024-10-01"
    }
  },
  "error_details": [],
  "message": "Success",
  "request_id": "REQUEST_ID",
  "response_time": "2024-11-12 10:30:00",
  "total": 1
}
复制
错误
复制成功
 上一章节
查询店铺绩效列表
下一章节 
查询买家之声列表