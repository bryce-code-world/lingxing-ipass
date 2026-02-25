发货单分配库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/lockStock	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_nos	发货单单号，对应查询FBA发货单列表接口字段【shipment_sn】	是	[array]	["SP240123007"]
is_auto_batch	是否锁定至批次，1：是，0：否，默认为否，否：只锁定库存数量，发货时按先进先出规则匹配出库批次；是：按先进先锁规则自动指定批次并锁定，发货时按锁定批次出库；分配库存后，可在【查询发货单详情】接口的采购信息中查看锁定的批次	否	[int]	0
请求示例
{
    "shipment_nos": ["SP240123007"],
    "is_auto_batch": 0
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
request_id	请求链路id	是	[string]	21EE37F0-5C84-0A70-9A7D-C9695538A59A
response_time	响应时间	是	[string]	2024-10-03 14:51:38
data	响应数据	是	[array]	
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "21EE37F0-5C84-0A70-9A7D-C9695538A59A",
    "response_time": "2024-10-03 14:51:38",
    "data": [],
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 1000,
    "message": "业务处理失败",
    "error_details": [
        "发货单SP2308150021不存在或已删除"
    ],
    "request_id": "5126F5E4-D31D-7BE2-E570-C0828731B37A",
    "response_time": "2024-10-03 14:59:05",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
发货单创建接口结果查询
FBA
下一章节 
发货单释放库存
FBA