发货单释放库存

待发货状态的发货单，释放库存后，将流转为待配货状态。若已拣货/部分拣货，请确认线下已将对应产品上架至原仓位，以防库存错乱。

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/releaseStock	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_nos	发货单号	是	[array]	["SP240123007"]
请求示例
{
    "shipment_nos": ["SP241016009"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	107DEE19-E3DD-E6C6-F63D-EB8FF2D92327
response_time	响应时间	是	[string]	2024-10-17 14:51:53
data	响应数据	是	[array]	
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "107DEE19-E3DD-E6C6-F63D-EB8FF2D92327",
    "response_time": "2024-10-17 14:51:53",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
发货单分配库存
FBA
下一章节 
FBA发货单发货
FBA