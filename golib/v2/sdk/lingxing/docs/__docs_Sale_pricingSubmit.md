批量修改Listing价格

支持批量修改Listing价格，价格幅度不超50%

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/listing/ProductPricing/pricingSubmit	HTTPS	POST	1
请求参数
参数名	类型	必填	说明	示例
pricing_params	[array]	是	参数数组，支持多个listing批量调价	
pricing_params>>sid	[int]	是	店铺sid，对应查询亚马逊店铺列表接口对应字段【sid】	1
pricing_params>>msku	[string]	是	seller_sku，卖家sku	MSKU679E6BF
pricing_params>>standard_price	[number]	是	价格原价	16.99
pricing_params>>sale_price	[number]	否	价格优惠价	11.99
pricing_params>>start_date	[string]	否	优惠价开始日期，格式：Y-m-d	2023-02-10
pricing_params>>end_date	[string]	否	优惠价结束日期，格式：Y-m-d	2023-02-11
返回结果

Json Object

参数名	类型	必填	说明	示例
code	[int]	是	状态码，0 成功	0
message	[string]	是	消息提示	success
error_details	[array]	是	错误信息	
request_id	[string]	是	请求链路id	38B2A580-8170-A242-CE77-1F2001E4E1A2
response_time	[string]	是	响应时间	2022-07-21 14:21:39
data	[object]	是	响应数据	
data>>success_num	[int]	是	调价提交成功listing数量	0
data>>failure_num	[int]	是	调价提交失败listing数量	2
data>>failure_detail	[array]	是	调价提交失败详情	
data>>failure_detail>>sid	[int]	是	店铺sid	109
data>>failure_detail>>msku	[string]	是	seller_sku，卖家sku	Hair-Ties-P0360
data>>failure_detail>>asin	[string]	是	asin	B0BB389BKQ
data>>failure_detail>>msg	[string]	是	调价提交失败原因	店铺状态异常不可调价
total	[int]	是		0
 上一章节
批量添加/编辑Listing配对
销售
下一章节 
批量获取Listing费用
销售