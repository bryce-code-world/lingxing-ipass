发货单创建接口结果查询
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/searchProcessResult	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
request_flag	生成单据时传的请求标识	是	[string]	3caabbccddefdaf
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2022-10-18 22:39:55
data	响应数据，按照创建时间倒序返回	是	[array]	
data>>request_flag	请求标识	否	[string]	3caabbccddefdaf
data>>request_url	请求接口	否	[string]	/storage/shipment/createSendedOrder
data>>process_result	请求结果	否	[object]	
data>>process_result>>error_details	失败时返回错误信息数组	否	[array]	
data>>process_result>>code	失败时返回错误代码	否	[string]	
data>>process_result>>data	成功时返回	否	[object]	
data>>process_result>>data>>order_sn	成功创建的单号	否	[string]	SP221018082
data>>process_msg		否	[string]	创建发货单成功，生成拣货流水失败
data>>process_status	请求状态说明：
0 处理中
1 已成功处理
2 已失败处理	否	[int]	1
data>>order_sn	成功时返回单号，失败时为空字符串	否	[string]	SP221018082
data>>gmt_create	请求创建时间	否	[string]	2022-10-18 22:39:46
data>>gmt_modified	请求修改时间	否	[string]	2022-10-18 22:39:53
 上一章节
生成已发货的发货单
FBA
下一章节 
发货单分配库存
FBA