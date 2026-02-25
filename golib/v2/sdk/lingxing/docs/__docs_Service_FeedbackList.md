查询评价管理 4-5星Feedback列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/cs/feedback/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
start_date	评论开始日期，格式：Y-m-d	是	[string]	2020-01-01
end_date	评论结束日期，格式：Y-m-d	是	[string]	2024-08-05
offset	分页偏移量，默认0	是	[int]	0
length	分页长度，默认20	是	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "sid": 1,
    "start_date": "2020-01-01",
    "end_date": "2024-08-05"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8D680B8B-3C6F-3E16-F055-73DC70EAD977
response_time	响应时间	是	[string]	2022-11-23 19:52:36
data	响应数据	是	[array]	
data>>sid	店铺id	是	[int]	2
data>>seller_name	店铺名称	是	[string]	店铺2
data>>country	国家	是	[string]	加拿大
data>>star	星级	是	[number]	5
data>>amazon_order_id	订单号	是	[string]	978-3423-3793796
data>>feedback_date	评论时间	是	[string]	2021-08-22
data>>feedback_content	评论内容	是	[string]	It is what I was I seeking.
data>>update_time	更新时间	是	[string]	2021-08-23 06:11:21
data>>operation_time	操作时间	是	[string]	2022-09-05 10:55:57
data>>remark	备注	是	[string]	22
data>>status	feedback处理状态：
0 待处理
1 处理中
2 已完成	是	[int]	1
data>>productList	商品信息	是	[array]	
data>>productList>>title	商品标题	是	[string]	product title 4C89BF4F06
data>>productList>>asin	asin	是	[string]	B3Mlf32uJL-U
data>>productList>>seller_sku	msku	是	[string]	M-6Qi32fu2L-U
total	总数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "9262E27D-B08D-4B53-DA9D-855A2FA3B001",
    "response_time": "2022-11-22 11:38:57",
    "data": [
        {
            "sid": 2,
            "seller_name": "店铺2",
            "country": "加拿大",
            "star": 5,
            "amazon_order_id": "978-3423-3793796",
            "feedback_date": "2021-08-22",
            "feedback_content": "It is what I was I seeking.",
            "update_time": "2021-08-23 06:11:21",
            "operation_time": "2022-09-05 10:55:57",
            "remark": "22",
            "status": 0,
            "productList": [
                {
                    "title": "product title 4C89BF4F06",
                    "asin": "B3Mlf32uJL-U",
                    "seller_sku": "M-6Qi32fu2L-U"
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询评价管理-1-3星Feedback列表
下一章节 
查询评价统计-Review列表