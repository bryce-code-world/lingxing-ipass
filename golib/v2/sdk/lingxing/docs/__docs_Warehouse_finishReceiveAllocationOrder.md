调拨单结束到货
基本信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/finishReceiveAllocationOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	调拨单单号	是	[string]	TF230309001
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C70FD2A7-4E7D-6B32-7D72-FC42587CE97E
response_time	响应时间	是	[string]	2022-07-29 14:33:42
data	返回数据	是	[array]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C70FD2A7-4E7D-6B32-7D72-FC42587CE97E",
    "response_time": "2022-07-29 14:33:42",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
调拨单分批收货
下一章节 
删除调拨单