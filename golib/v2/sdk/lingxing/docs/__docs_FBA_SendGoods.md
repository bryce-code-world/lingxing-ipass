FBA发货单发货

支持将已有的FBA发货单状态变更为“已发货”，并扣减发货仓库库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/storage/shipment/sendGoods	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_nos	发货单号列表	是	[array]	 
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2022-10-18 22:39:55
data	响应数据	是	[object]	 
 上一章节
发货单释放库存
FBA
下一章节 
编辑发货单
FBA