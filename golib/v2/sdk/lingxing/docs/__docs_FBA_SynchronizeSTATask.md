同步STA任务到ERP
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-plan/gatherInboundPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanIdList	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[array]	wf0a914e89-d126-4ed9-a093-2078289fed05
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-plan/gatherInboundPlan?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanIdList": ["wf0a914e89-d126-4ed9-a093-2078289fed05"],
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
data>>errorMsg	失败原因	是	[string]	
data>>failInboundPlanIds	同步失败STA任务编号数组	是	[array]	
data>>failNum	同步失败数量	是	[int]	
data>>successInboundPlanIds	同步成功的STA任务编号数组	是	[array]	
data>>successNum	同步成功数量	是	[int]	
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "errorDetails": [],
    "requestId": "3b3d867e7d014971a580549f107c8c5a.1732773886069",
    "responseTime": "2024-11-28T14:04:46.069",
    "data": {
        "errorMsg": "失败原因",
        "failInboundPlanIds": ["同步失败STA任务编号"],
        "failNum": 1,
        "successInboundPlanIds": ["同步成功的STA任务编号"],
        "successNum": 1
    }
}
复制
错误
复制成功
 上一章节
获取商品预处理信息
FBA
下一章节 
同步亚马逊货件到ERP
FBA