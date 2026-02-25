更新发货单自定义成本
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/updateCustomCost	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_sn	发货单号	是	[string]	SP220120057
is_custom_cost	是否自定义成本	是	[int]	1
list	自定义成本信息数组	否	[array]	
list>>msku	msku	否	[string]	MSKU2A13471
list>>sku	sku	否	[string]	SKUC5A788F
list>>shipment_id	货件单号	否	[string]	FBA16SP419P4
list>>custom_purchase_price_unit	采购单价（自定义成本)	否	[number]	1.0001
list>>custom_outbound_cost_unit	单位出库费用（自定义成本)	否	[number]	1.001
list>>custom_aux_cost	单位辅料费用（自定义成本)	否	[number]	1.01
list>>custom_outbound_head_cost_unit	单位出库头程（自定义成本)	否	[number]	1
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2022-10-18 22:39:55
data	响应数据	是	[array]	[]
 上一章节
更新发货单物流信息
FBA
下一章节 
获取发货单头程物流-承运商信息
FBA