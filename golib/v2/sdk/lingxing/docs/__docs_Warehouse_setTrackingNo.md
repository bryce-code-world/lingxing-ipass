物流下单 - 编辑运单号/跟踪号
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/logisticsOrdering/setTrackingNo	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
waybill_no	运单号	是	[string]	DXSXS412566
wo_number	销售出库单号	是	[string]	WO103247958556684288
tracking_no	跟踪号	否	[string]	S1245224552
logistics_freight	物流运费	否	[string]	100
logistics_freight_currency_code	物流运费币种：
CNY
USD
EUR
JPY
AUD
CAD
MXN
GBP
INR
AED
SGD
SAR
BRL
SEK
PLN
TRY
HKD	否	[string]	CNY
pkg_fee_weight	计费重	否	[string]	10
pkg_fee_weight_unit	计费重单位：
g
kg	否	[string]	g
请求示例
{
   "waybill_no":"xsdsdsds",
   "wo_number":"WO10324795xx93872384",
   "tracking_no":"xsxsxsxs",
   "logistics_freight":"100",
   "logistics_freight_currency_code":"CNY",
   "pkg_fee_weight":"10",
   "pkg_fee_weight_unit":"g"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "c7f5ded0c503494980bf9c39f138c594.1703299828218",
    "response_time": "2023-12-23 10:50:29",
    "data": null,
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "程序内部错误",
    "error_details": [当前订单不存在或物流状态不是待导入 [请求码:008C66]],
    "request_id": "c7f5ded0c503494980bf9c39f138c594.1703299828218",
    "response_time": "2023-12-23 10:50:29",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
查询销售出库单列表
下一章节 
查询销售出库单物流面单