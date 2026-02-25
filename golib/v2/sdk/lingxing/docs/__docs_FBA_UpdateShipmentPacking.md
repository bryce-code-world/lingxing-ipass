修改货件装箱信息
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-packing/updateShipmentPacking	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
boxes	装箱明细数据	是	[array]	
boxes>>contentInformationSource	提供方式
BOX_CONTENT_PROVIDED、MANUAL_PROCESS、BARCODE_2D	是	[string]	
boxes>>dimensions		否	[object]	
boxes>>dimensions>>height	高	是	[number]	
boxes>>dimensions>>length	长	是	[number]	
boxes>>dimensions>>unitOfMeasurement	长度单位：IN、CM	是	[string]	
boxes>>dimensions>>width	宽	是	[number]	
boxes>>items	商品信息	是	[array]	
boxes>>items>>expiration	有效期	否	[string]	
boxes>>items>>labelOwner	标签类型（AMAZON,SELLER,NONE）	是	[string]	
boxes>>items>>msku	msku	是	[string]	
boxes>>items>>prepOwner	预处理提供方（AMAZON,SELLER,NONE）	是	[string]	
boxes>>items>>quantity	申报量	是	[int]	
boxes>>packageId	包裹id	否	[string]	
boxes>>weight		否	[object]	
boxes>>weight>>unit	重量单位：LB、KG	是	[string]	
boxes>>weight>>value	重量	是	[number]	
inboundPlanId	任务编号	否	[string]	
items		否	[array]	
items>>expiration	有效期	否	[string]	
items>>labelOwner	标签类型（AMAZON,SELLER,NONE）	是	[string]	
items>>msku	msku	是	[string]	
items>>prepOwner	预处理提供方（AMAZON,SELLER,NONE）	是	[string]	
items>>quantity	申报量	是	[int]	
shipmentId	货件号	否	[string]	
sid	领星店铺ID，对应查询亚马逊店铺列表接口对应字段【sid】	否	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-packing/updateShipmentPacking?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
   "boxes": [{
        "contentInformationSource": "BOX_CONTENT_PROVIDED",
        "dimensions": {
            "height": 3,
            "length": 4,
            "width": 5,
            "unitOfMeasurement": "CM"
        },
        "items": [{
            "expiration": "2023-02-06",
            "labelOwner": "AMAZON",
            "msku": "msku",
            "prepOwner": "SELLER",
            "quantity": 1
        }],
        "packageId": "包裹id",
        "weight": {
            "value": 10
        }
    }],
    "inboundPlanId": "任务编号",
    "items": [{
        "expiration": "2023-02-06",
        "labelOwner": "SELLER",
        "msku": "msku",
        "prepOwner": "AMAZON",
        "quantity": 1
    }],
    "shipmentId": "货件id",
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
上传货件跟踪号
FBA
下一章节 
修改货件实际状态
FBA