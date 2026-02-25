作废采购单

允许作废处于 “待审核”、“待下单”、“待到货”、“已完成” 状态下的采购单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase/cancel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	采购单系统单号	是	[string]	PO211112002
reason	作废原因，长度不超过80	是	[string]	不需要采购了
is_cancel_relation	是否取消关联采购计划：0 否【默认】，1 是	是	[int]	1
请求示例
{
    "order_sn": "PO211112002",
    "reason": "不需要采购了",
    "is_cancel_relation": 1
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	81CA6A9A-A5C4-8B5E-5FED-E1C2EC04753C
response_time	响应时间	是	[string]	2020-12-01 12:11:12
data	响应数据	是	[array]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8DFEE1CA-EC9B-F401-5D41-9F95251D5D50",
    "response_time": "2023-06-25 12:41:04",
    "data": [],
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "到过货的采购单不允许作废 [请求码:457CD6]",
    "request_id": "545B63DD-7945-067B-B861-0EDEB7723773",
    "response_time": "2023-06-25 12:51:14",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
采购单下单
下一章节 
创建待到货的采购单