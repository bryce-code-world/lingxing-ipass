调拨单全部收货

支持将系统“待收货“状态的调拨单全部收货，单据状态转为“已完成“，并增加入库仓的库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/receiveAllocationOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSnMany	调拨单号，支持多个，英文逗号分隔	是	[string]	TF220425001,TF220419001
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误消息	是	[array]	
request_id	请求链路id	是	[string]	5A1EBE67-7793-9E94-7790-AA457B81B3F2
response_time	响应时间	是	[string]	2022-06-27 11:54:31
data	返回数据	是	[array]	
total	返回总数	是	[int]	 
 上一章节
创建待收货/已完成的调拨单
下一章节 
调拨单分批收货