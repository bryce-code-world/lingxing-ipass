上传货件跟踪号
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-shipment/updateShipmentTrack	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
sid	领星店铺ID，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
shipmentId	货件id，对应查询货件方案接口对应字段【shipmentId】	是	[string]	"shd10e38ca-45e7-4a97-8512-780acf343f4b3"
billOfLadingNumber	提货单号(LTL建议填写，非必填)	否	[object]	
freightBillNumber	LTL跟踪编号(LTL必填)	否	[string]	
trackBOList	SPD跟踪编号(SPD必填)	否	[array]	
trackBOList>>boxId	箱子id	否	[string]	
trackBOList>>trackingId	跟踪id	否	[string]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-shipment/updateShipmentTrack?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "sid": 1,
    "shipmentId": "货件id",
    "billOfLadingNumber": "提货单号",
    "freightBillNumber": "LTL001",
    "trackBOList": [{
        "boxId": "箱子id：通过查询货件装箱详情接口获取箱子ID，注意不是包裹ID",
        "trackingId": "跟踪id"
    }]
}'
复制
错误
复制成功
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
        "inboundPlanId": "亚马逊任务编号",
        "taskId": "任务id",
        "taskStatus": "任务状态"
    }
}
复制
错误
复制成功
 上一章节
提交货件配送服务
FBA
下一章节 
修改货件装箱信息
FBA