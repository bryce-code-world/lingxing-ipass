备货单结束到货
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/completeReceipt	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	OWS220828005
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	 
 上一章节
备货单分批收货
下一章节 
创建待发货/待收货/已完成的备货单