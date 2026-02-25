导入面单

支持自发货订单导入面单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/selfShipmentOrder/importLabel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
fileName	面单文件名	是	[string]	aa
base64File	PDF/PNG/JPG/JPEG格式文件 Base64编码	是	[string]	
trackingNo	运单号	是	[string]	B09JK94H12
waybillNo	跟踪号	是	[string]	15165615
woId	出库单id，对应查询销售出库单列表	是	[int]	1234
请求示例
curl --location --request POST 'http://openapi.lingxing.com/basicOpen/selfShipmentOrder/importLabel?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fileName":"bear_big.jpg",
    "trackingNo":"123456",
    "waybillNo":false,
    "woId":6178,
    "base64File":""
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	错误说明	是	[string]	success
request_id	请求id	是	[string]	69a2ddc2-f3bb-11ed-aed5-0242ac130003
response_time	请求时间	是	[string]	2023-05-16 15:29:41
error_details	错误信息	是	[array]	
data	响应数据	是	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "47525526bfd3484cbb1aba1870a1b2c3.1744366551396",
    "response_time": "2025-04-11 18:15:52",
    "data": null
}
复制
错误
复制成功
 上一章节
查询亚马逊标发结果
销售
下一章节 
查询促销活动列表-优惠券
销售