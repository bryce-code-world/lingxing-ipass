查询客户列表（旧）

本接口即将下线，请尽快切换到：查询客户列表（新）

接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/crm/open/api/customer/list	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
sids	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[array]	[1,136,115]
time_search_type	时间筛选查询类型：
1 首次购买时间
2 最近购买时间	是	[int]	1
start_date	开始时间	是	[string]	2024-01-01
end_date	结束时间	是	[string]	2024-08-05
offset	页码，默认1	否	[int]	1
length	每页条数，默认100	否	[int]	100
请求示例
{
    "sids": [1,136,115],
    "time_search_type": 1,
    "start_date": "2024-01-01",
    "end_date": "2024-08-05",
    "offset": 1,
    "length": 100
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
request_id	请求链路id	是	[string]	4679b04770724a69a15a9d67415c101d.189.16733356099605979
response_time	响应时间	是	[string]	2023-01-10 16:47:05
data	响应数据	是	[array]	
data>>store_name	店铺名称	是	[string]	店铺3
data>>country_name	国家名称	是	[string]	美国
data>>group	分组	是	[array]	["分组1","分组二"]
data>>order_items	总订单	是	[number]	2
data>>volume	总销量	是	[number]	3
data>>amount	总销售额	是	[number]	32.25
data>>per_customer_transaction	平均客单价	是	[number]	16.49
data>>currency_icon	币种符号	是	[string]	$
data>>refund_number	退款订单数	是	[number]	1
data>>refund_sales_number	退款销售数	是	[object]	1
data>>refund_rate	退款率	是	[object]	0.20
data>>return_number	退货订单数	是	[number]	
data>>return_sales_number	退货销量数	是	[number]	1
data>>return_rate	退货率	是	[number]	0.30
data>>feedback_number	Feedback评论数	是	[string]	2
data>>feedback_bad_number	Feedback差评数	是	[string]	1
data>>feedback_rate	留平率	是	[string]	0.6
data>>first_purchase_date	首次购买时间	是	[string]	2021-07-01 19:07:34
data>>last_purchase_date	最近购买时间	是	[string]	2021-08-15 08:08:24
total	总数	是	[int]	20
 上一章节
查询评价统计-Feedback每日新增数
下一章节 
查询售后工单列表