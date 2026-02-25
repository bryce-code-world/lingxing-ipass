取消STA任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-plan/cancelInboundPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
sid	亚马逊店铺sid，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-plan/cancelInboundPlan?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "sid": 104
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
    "errorDetails": null,
    "requestId": "7004576e64e04a489f18186c9ec92659.1732858868684",
    "responseTime": "2024-11-29T13:41:08.684",
    "data": {
        "taskId": "231512640894445056",
        "taskStatus": "process",
        "inboundPlanId": "wf05d62111-7ae6-4e31-8e69-1ffd5ac09fd2",
        "errorMsg": ""
    }
}
复制
错误
复制成功
 上一章节
上传货件跟踪号
FBA
下一章节 
查询异步任务状态
FBA