海外仓备货单发货

支持将系统已有的海外仓备货单状态调整为“待收货”，并扣减发货仓库库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/sendInbound	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	 
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	信息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路	是	[string]	7E80FC0B-55C1-BAC5-0E4B-7AE59706F051
response_time	响应时间	是	[string]	2021-06-15 20:16:38
data	数据	是	[object]	
total	总数	是	[int]	 
 上一章节
查询备货单收货记录
下一章节 
备货单分批收货