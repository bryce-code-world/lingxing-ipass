查询评价统计-Review列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/v2/cs/reviewReport/lists	HTTPS	GET	1
请求参数

参数作为query参数拼接在url上，签名时参数用json格式参与生成签名

参数名	说明	必填	类型	示例
start_date	开始时间【时间间隔不超过1年】，格式：Y-m-d	是	[string]	2020-01-01
end_date	结束时间【时间间隔不超过1年】，格式：Y-m-d	是	[string]	2024-08-05
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[array]	[1,136,139]
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
curl --location --request GET 'https://openapi.lingxing.com/erp/sc/v2/cs/reviewReport/lists?app_key=&access_token=&timestamp=&sign=&start_date=2024-09-01&end_date=2025-05-01&offset=0&length=20'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	73A0718B-AA68-E120-D28B-F7175018175D
response_time	响应时间	是	[string]	2022-04-19 15:16:18
total	总数	是	[int]	701
data	响应数据	是	[array]	
data>>ratings	子rating总数	是	[number]	3112
data>>five_star	5星review新增数	是	[number]	5
data>>four_star	4星review新增数	是	[number]	
data>>three_star	3星review新增数	是	[number]	1
data>>two_star	2星review新增数	是	[number]	1
data>>one_star	1星review新增数	是	[number]	2
data>>review_num	review数	是	[number]	9
data>>good_num	review好评数	是	[number]	5
data>>negative_num	review中差评数	是	[number]	4
data>>good_rate	review好评率	是	[number]	0.5556
data>>negative_rate	review中差评率	是	[number]	0.4444
data>>modified_num	review改评数	是	[number]	
data>>remove_num	review删评数	是	[number]	
data>>asin	子asin	是	[string]	B0ABC7613019
data>>asin_url	asin链接	是	[string]	https://www.amazon.com/dp/B0ABC7613019
data>>image_url	图片链接	是	[string]	
data>>title	商品标题	是	[string]	标题38746
data>>country	国家	是	[string]	美国
data>>score	评分	是	[number]	4.6
data>>mark	仅评分数	是	[number]	
data>>seller_name	店铺名称	是	[array]	["account5","account2"]
data>>local_info		是	[array]	
data>>local_info>>local_sku	sku	是	[string]	SKU27CB788
data>>local_info>>local_name	品名	是	[string]	[演示数据]Android充电线
data>>parent_asin	父asin	是	[array]	["B0ABC3499"]
data>>seller_list	店铺列表	是	[array]	
data>>seller_list>>sid	sid	是	[int]	2
data>>seller_list>>seller_name	店铺名称	是	[string]	店铺2
成功请求示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D31C7F9E-675B-EFBD-4E64-FD5D909948C1",
    "response_time": "2022-09-28 11:03:26",
    "total": 1
    "data": [
        {
            "ratings": 64,
            "five_star": 1,
            "four_star": 0,
            "three_star": 0,
            "two_star": 0,
            "one_star": 0,
            "review_num": 1,
            "good_num": 1,
            "negative_num": 0,
            "good_rate": 1,
            "negative_rate": 0,
            "modified_num": 0,
            "remove_num": 0,
            "asin": "B0ABC50851",
            "asin_url": "https://www.amazon.com/dp/B0ABC50851",
            "image_url": "https://image-1254213275.cos.ap-guangzhou.myqcloud.com/dev/20210914/Sony%20ZX%E7%B3%BB%E5%88%97%E6%9C%89%E7%BA%BF%E5%A4%B4%E6%88%B4%E5%BC%8F%E8%80%B3%E6%9C%BA%20%E9%BB%91%E8%89%B2%20MDR-ZX110.jpg",
            "title": "标题38657",
            "country": "美国",
            "score": 4.6,
            "mark": 0,
            "seller_name": [
                "店铺2"
            ],
            "seller_list": [
                {
                    "sid": 2,
                    "seller_name": "店铺2"
                }
            ],
            "local_info": [
                {
                    "local_sku": "SKU09707D9",
                    "local_name": "[演示数据]Sony ZX系列有线头戴式耳机 黑色 MDR-ZX110"
                }
            ],
            "parent_asin": [
                "B0ABC86025"
            ]
        }
    ]
}
复制
错误
复制成功
失败请求示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "start_date => start_date不能为空"
    ],
    "request_id": "8671A9D1-BC61-9398-FBD9-C90A672DAD11",
    "response_time": "2022-03-31 15:55:03",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询评价管理-4-5星Feedback列表
下一章节 
查询评价统计-Review每日新增数