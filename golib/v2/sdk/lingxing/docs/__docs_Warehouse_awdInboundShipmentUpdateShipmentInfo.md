更新AWD货件跟踪编号
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-shipment/updateShipmentInfo	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderId	STA任务编号	是	[string]	STA20230715001
shipmentId	货件号	是	[string]	AWD20230715001
sid	领星店铺ID 对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	123456789
trackingId	跟踪编号	是	[string]	1Z1234567890123456
请求示例

json

{
  "orderId": "STA20230715001",
  "shipmentId": "AWD20230715001",
  "sid": 123456789,
  "trackingId": "1Z1234567890123456"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code		否	[int]	
data		否	[boolean]	
message	提示信息	否	[string]	
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[array]	
response_time	响应时间	否	[string]	
返回成功示例

json

{
  "code": 0,
  "data": true,
  "message": "Success",
  "error_details": [],
  "request_id": "d4e5f6g7-h8i9-j0k1-l2m3-n4o5p6q7r8s9",
  "response_time": "2023-07-15T17:15:30Z"
}
复制
错误
复制成功
 上一章节
查询AWD入库货件列表
FBA
下一章节 
打印AWD入库货件箱子标签
FBA