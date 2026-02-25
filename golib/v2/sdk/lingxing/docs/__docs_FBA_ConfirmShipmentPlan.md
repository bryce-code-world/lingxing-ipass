确认货件方案
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-shipment/confirmPlacementOption	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
placementOptionId	货件方案id	是	[string]	
shipmentIds	货件列表：传入对应货件方案id下的所有货件id	是	[array]	["shd10e38ca-45e7-4a97-8512-780acf343f4b3"]
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-shipment/confirmPlacementOption?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "placementOptionId": "货件方案id",
    "shipmentIds": ["货件id"],
    "sid": 1
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
errorDetails	错误信息	是	[array]	
requestId	请求链路id	是	[string]	
responseTime	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[object]	
data>>errorMsg	错误信息	是	[string]	
data>>inboundPlanId	亚马逊任务编号	是	[string]	
data>>taskId	任务id	是	[string]	
data>>taskStatus	任务状态
process
success
failure
local_failure	是	[string]	
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "errorDetails": [],
    "requestId": "3b3d867e7d014971a580549f107c8c5a.1732773886069",
    "responseTime": "2024-11-28T14:04:46.069",
    "data": {
        "errorMsg": "错误信息",
        "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
        "taskId": "任务id",
        "taskStatus": "任务状态"
    }
}
复制
错误
复制成功
 上一章节
查询货件方案的装箱信息
FBA
下一章节 
生成承运方式
FBA