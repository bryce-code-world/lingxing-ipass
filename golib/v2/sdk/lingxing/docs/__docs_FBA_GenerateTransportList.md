生成承运方式
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-shipment/generateTransportList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	wf0a914e89-d126-4ed9-a093-2078289fed05
shipmentIdList	发货信息：array，注意需提供所有货件单号的发货时间，生成所有货件的承运方式	是	[array]	["shd10e38ca-45e7-4a97-8512-780acf343f4b3"]
shipmentIdList>>palletList	帕托信息：如果修改了帕托信息需要重新生成承运方式	否	[array]	
shipmentIdList>>palletList>>height	高	否	[number]	
shipmentIdList>>palletList>>length	长	否	[number]	
shipmentIdList>>palletList>>lengthUnit	长度单位(IN-英制,CM-公制)	否	[string]	
shipmentIdList>>palletList>>quantity	数量	否	[int]	
shipmentIdList>>palletList>>stackability	是否可堆叠（STACKABLE、NON_STACKABLE）	否	[string]	
shipmentIdList>>palletList>>weight	单个重量	否	[number]	
shipmentIdList>>palletList>>weightUnit	重量单位（LB-磅，KG-千克）	否	[string]	
shipmentIdList>>palletList>>width	宽	否	[number]	
shipmentIdList>>shipingTime	发货时间：如果修改了发货时间需要重新生成承运方式
格式：yyyy-MM-dd	是	[datetime]	
shipmentIdList>>shipmentId	货件id	是	[string]	
sid	亚马逊店铺sid，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-shipment/generateTransportList?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "inboundPlanId": "wf0a914e89-d126-4ed9-a093-2078289fed05",
    "shipmentIdList": [{
        "palletList": [{
            "height": "高",
            "length": "长",
            "lengthUnit": "长度单位(IN-英制,CM-公制)",
            "quantity": "数量",
            "stackability": "是否可堆叠（STACKABLE、NON_STACKABLE）",
            "weight": "单个重量",
            "weightUnit": "重量单位（LB-磅，KG-千克）",
            "width": "宽"
        }],
        "shipingTime": "发货时间：如果修改了发货时间需要重新生成承运方式",
        "shipmentId": "货件id"
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
确认货件方案
FBA
下一章节 
生成可选送达时间
FBA