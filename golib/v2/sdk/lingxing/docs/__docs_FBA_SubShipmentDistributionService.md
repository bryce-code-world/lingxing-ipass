提交货件配送服务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-shipment/setDeliveryService	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
shipmentDistributionInfo	货件配送信息	是	[array]	
shipmentDistributionInfo>>alphaCode	承运方式编码	是	[string]	
shipmentDistributionInfo>>alphaName	承运方式名称	是	[string]	
shipmentDistributionInfo>>declaredAmount	申报价值	否	[number]	
shipmentDistributionInfo>>declaredCode	申报价值货币	否	[string]	
shipmentDistributionInfo>>deliveryWindowOptionId	送达时间ID	是	[string]	
shipmentDistributionInfo>>endDate	送达时段-结束时间:(北京时间)
备注:格式：YYYY-MM-DD	是	[string]	
shipmentDistributionInfo>>freightClass	货物等级	否	[string]	
shipmentDistributionInfo>>palletList	帕托信息	否	[array]	
shipmentDistributionInfo>>palletList>>height	高	否	[number]	
shipmentDistributionInfo>>palletList>>length	长	否	[number]	
shipmentDistributionInfo>>palletList>>lengthUnit	长度单位(IN-英制,CM-公制)	否	[string]	
shipmentDistributionInfo>>palletList>>quantity	数量	否	[int]	
shipmentDistributionInfo>>palletList>>stackability	是否可堆叠	否	[string]	
shipmentDistributionInfo>>palletList>>weight	单个重量	否	[number]	
shipmentDistributionInfo>>palletList>>weightUnit	重量单位（LB-磅，KG-千克）	否	[string]	
shipmentDistributionInfo>>palletList>>width	宽	否	[number]	
shipmentDistributionInfo>>shipingTime	发货时间:(北京时间)
备注:格式：YYYY-MM-DD	是	[string]	
shipmentDistributionInfo>>shipmentId	货件单号	是	[string]	
shipmentDistributionInfo>>shippingMode	货件类型（GROUND_SMALL_PARCEL代表小包裹快递（SPD）、FREIGHT_LTL代表汽运零担（LTL））	是	[string]	
shipmentDistributionInfo>>shippingSolution	承运人(USE_YOUR_OWN_CARRIER代表其他承运人、AMAZON_PARTNERED_CARRIER代表亚马逊合作承运人)	是	[string]	
shipmentDistributionInfo>>startDate	送达时段-开始时间:(北京时间)
备注:格式：YYYY-MM-DD	是	[string]	
shipmentDistributionInfo>>transportationOptionId	承运方式ID	是	[string]	
sid	领星店铺ID，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-shipment/setDeliveryService?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "shipmentDistributionInfo": [{
        "alphaCode": "CY0001",
        "alphaName": "CY0001-00",
        "declaredAmount": 288,
        "declaredCode": "申报价值货币",
        "deliveryWindowOptionId": "",
        "endDate": "2022-12-22",
        "freightClass": "货物等级",
        "palletList": [{
            "height": 5,
            "length": 6,
            "lengthUnit": "CM",
            "quantity": 10,
            "stackability": "",
            "weight": 20,
            "weightUnit": "LB",
            "width": "宽"
        }],
        "shipingTime": "2022-11-10",
        "shipmentId": "货件单号",
        "shippingMode": "SPD",
        "shippingSolution": "USE_YOUR_OWN_CARRIER",
        "startDate": "2022-11-22",
        "transportationOptionId": "CY0003"
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
data>>operationId	操作id	是	[string]	
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "errorDetails": [],
    "requestId": "3b3d867e7d014971a580549f107c8c5a.1732773886069",
    "responseTime": "2024-11-28T14:04:46.069",
    "data": {
        "operationId": "操作id"
    }
}
复制
错误
复制成功
 上一章节
提交送达时间
FBA
下一章节 
上传货件跟踪号
FBA