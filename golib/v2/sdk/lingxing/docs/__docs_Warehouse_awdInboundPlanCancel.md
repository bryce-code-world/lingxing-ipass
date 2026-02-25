取消AWD入库任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-plan/cancel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderId	AWD任务编号	是	[string]	AWD123456789
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	109
请求示例
{
    "orderId": "AWD123456789",
    "sid": 109
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码,0成功	否	[int]	0
data	数据	否	[object]	
data>>orderId	AWD任务编号	否	[string]	AWD123456789
error_details	错误信息	否	[array]	[]
message	提示信息	否	[string]	操作成功
request_id	请求链路id	否	[string]	a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8
response_time	响应时间	否	[string]	2024-01-01 12:00:00
返回成功示例

json

{
    "code": 0,
    "message": "success",
    "data": {
        "orderId": "AWD123456789"
    },
    "error_details": [],
    "request_id": "a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8",
    "response_time": "2024-01-01 12:00:00"
}
复制
错误
复制成功
 上一章节
地址簿-配送地址详情
FBA
下一章节 
确认AWD入库任务
FBA