取消多渠道订单
API Path	请求协议	请求方式	令牌桶容量
/order/amzod/api/cancelOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
seller_fulfillment_order_id	卖家订单号	是	[string]	332565
请求示例
{
    "sid":6,
    "seller_fulfillment_order_id":"332565"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	3EE1982B-EDC3-0531-83DF-2BEA6F080DDA
response_time	响应时间	是	[string]	2022-12-0714:54:56
total	总记录数	是	[int]	1
data	响应数据	是	[array]	
返回成功示例
{
    "code": 0,
    "data": [],
    "response_time": "2023-04-27 09:28:29",
    "total": 0,
    "error_details": [],
    "message": "操作成功",
    "request_id": "212149f55784429aaa69fe7f0cab12e7.100.16825589067040171"
}
复制
错误
复制成功
 上一章节
查询售后订单列表
销售
下一章节 
查询多渠道订单-交易明细
销售