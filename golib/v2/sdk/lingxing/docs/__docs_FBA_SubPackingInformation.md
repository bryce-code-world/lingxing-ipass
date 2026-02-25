提交装箱信息

如需获取提交装箱信息结果，请再额外调用查询异步任务状态接口,获取提交结果

接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-packing/setPackingInformation	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
packageGroupings	分组装箱数据	是	[array]	
packageGroupings>>boxes	箱子信息	是	[array]	
packageGroupings>>boxes>>dimensions	维度信息	否	[object]	
packageGroupings>>boxes>>dimensions>>height	高	是	[number]	
packageGroupings>>boxes>>dimensions>>length	长	是	[number]	
packageGroupings>>boxes>>dimensions>>unitOfMeasurement	长度单位：IN、CM	是	[string]	
packageGroupings>>boxes>>dimensions>>width	宽	是	[number]	
packageGroupings>>boxes>>items	商品信息	是	[array]	
packageGroupings>>boxes>>items>>expiration	有效期	否	[string]	
packageGroupings>>boxes>>items>>labelOwner	标签类型
AMAZON
SELLER
NONE	是	[string]	
packageGroupings>>boxes>>items>>msku	msku	是	[string]	
packageGroupings>>boxes>>items>>prepOwner	预处理提供方
AMAZON
SELLER
NONE	是	[string]	
packageGroupings>>boxes>>items>>quantity	申报量	是	[int]	
packageGroupings>>boxes>>weight	重量	否	[object]	
packageGroupings>>boxes>>weight>>unit	重量单位：LB、KG	是	[string]	
packageGroupings>>boxes>>weight>>value	重量单位值	是	[number]	
packageGroupings>>packingGroupId	包装组id：先装箱后分仓方式时必填；先分仓后装箱方式时无需填写	否	[string]	
packageGroupings>>shipmentId	货件id：先分仓后装箱方式时必填；分装箱后分仓方式时无需填写	否	[string]	
sid	亚马逊店铺sid，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-packing/setPackingInformation?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "packageGroupings": [{
        "boxes": [{
            "dimensions": {
                "height": 3,
                "length": 5,
                "width": 5,              
                "unitOfMeasurement": "cm"                
            },
            "items": [{
                "expiration": "有效期",
                "labelOwner": "AMAZON",
                "msku": "msku",
                "prepOwner": "AMAZON",
                "quantity": 1
            }],
            "weight": {
                "value": 1
            }
        }],
        "packingGroupId": "包装组id：先装箱后分仓方式时必填；先分仓后装箱方式时无需填写",
        "shipmentId": "货件id：先分仓后装箱方式时必填；分装箱后分仓方式时无需填写"
    }],
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
        "inboundPlanId": "亚马逊任务编号",
        "taskId": "任务id",
        "taskStatus": "任务状态"
    }
}
复制
错误
复制成功
 上一章节
保存装箱信息
FBA
下一章节 
生成货件方案
FBA