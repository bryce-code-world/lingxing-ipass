查询客户列表（新）
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/crm/customer/index	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sort_field	结果按字段排序	否	[string]	order_items
sort_type	desc=倒序，asc=升序	否	[string]	desc
date_field	时间筛选查询类型，1：首次购买时间 ，2：最近购买时间	否	[string]	first_purchase_date
start_date	筛选开始时间	否	[string]	2021-01-31
end_date	筛选结束时间	否	[string]	2021-03-02
currency_type	币种，0=原币种，1=CNY，2=USD	否	[number]	0
search_field	支持搜索的字段 buyer_email、buyer_name	否	[string]	buyer_email
offset	偏移量	否	[number]	0
length	分页长度 ，默认20 ，上限200	否	[number]	20
search_value	搜索值	否	[string]	t83lnqt83k5317b@marketplace.amazon.com
sids	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[string]	17
请求示例
{
    "sort_field": "order_items",
    "sort_type": "desc",
    "date_field": "first_purchase_date",
    "start_date": "2021-01-31",
    "end_date": "2021-03-02",
    "currency_type": 0,
    "search_field": "buyer_email",
    "offset": 0,
    "length": 20,
    "search_value": "t83lnqt83k5317b@marketplace.amazon.com",
    "sids": "17"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	最近购买时间	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求id	是	[string]	842d7733ef5d4d968bc6877e06c34453.1697004896752
response_time	响应时间	是	[string]	2023-10-11 14:14:56
data		是	[object]	
data>>list		是	[array]	
data>>list>>buyer_email	买家邮箱	是	[string]	1f9rg94t5vnhqgm@marketplace.amazon.com
data>>list>>buyer_name	买家姓名	是	[string]	linda meckel
data>>list>>sid	店铺id	是	[array]	109
data>>list>>country	国家名称	是	[string]	美国
data>>list>>order_items	总订单	是	[number]	1
data>>list>>volume	总销量	是	[number]	1
data>>list>>amount	总销售额	是	[number]	0
data>>list>>per_customer_transaction	平均客单价	是	[number]	0
data>>list>>currency_icon	币种	是	[string]	$
data>>list>>refund_number	退款订单数	是	[number]	1
data>>list>>refund_sales_number	退款销售数	是	[number]	1
data>>list>>refund_rate	退款率	是	[number]	1
data>>list>>return_number	退货订单数	是	[number]	0
data>>list>>return_sales_number	退货销量数	是	[number]	0
data>>list>>return_rate	退货率	是	[number]	0
data>>list>>feedback_number	Feedback评论数	是	[number]	0
data>>list>>seller_name	店铺名称	是	[string]	篮球
data>>list>>feedback_bad_number	Feedback差评数	是	[number]	0
data>>list>>feedback_rate	Feedback留评率	是	[number]	0
data>>list>>feedback_bad_rate	Feedback差评率	是	[number]	0
data>>list>>first_purchase_date	首次购买时间	是	[string]	2021-02-23 14:24:46
data>>list>>last_purchase_date	最近购买时间	是	[string]	2021-02-23 14:24:46
data>>list>>remark	备注	是	[string]	
data>>list>>group	分组	是	[array]	["4.13验证","选择两个店铺","081111","重构后验证","自定义080801","12"]
data>>total	总数	是	[number]	
total	总数	是	[number]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "842d7733ef5d4d968bc6877e06c34453.1697004896752",
    "response_time": "2023-10-11 14:14:56",
    "data": {
        "list": [{
            "buyer_email": "1f9rg94t5vnhqgm@marketplace.amazon.com",
            "buyer_name": "linda meckel",
            "sid": 109,
            "country": "美国",
            "order_items": 1,
            "volume": 1,
            "amount": 0,
            "per_customer_transaction": 0,
            "currency_icon": "$",
            "refund_number": 1,
            "refund_sales_number": 1,
            "refund_rate": 1,
            "return_number": 0,
            "return_sales_number": 0,
            "return_rate": 0,
            "feedback_number": 0,
            "seller_name": "篮球",
            "feedback_bad_number": 0,
            "feedback_rate": 0,
            "feedback_bad_rate": 0,
            "first_purchase_date": "2021-02-23 14:24:46",
            "last_purchase_date": "2021-02-23 14:24:46",
            "remark": "",
            "group": ["4.13验证", "选择两个店铺", "081111", "重构后验证", "自定义080801", "12"]
        }],
        "total": 1
    },
    "total": 1
}
复制
错误
复制成功
 上一章节
查询RMA管理
下一章节 
查询业绩通知列表