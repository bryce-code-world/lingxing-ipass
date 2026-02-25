查询评价统计-Feedback列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/cs/feedbackReport/lists	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
start_date	开始时间【时间间隔不超过1年】，格式：Y-m-d	是	[string]	2024-01-01
end_date	结束时间【时间间隔不超过1年】，格式：Y-m-d	是	[string]	2024-08-05
请求示例
{
    "offset": 0,
    "length": 20,
    "start_date": "2024-01-01",
    "end_date": "2024-08-05"
}
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
data	响应数据	是	[array]	
data>>count_lifetime	feedback总数	是	[number]	137
data>>count_12	近1年feedback数	是	[number]	30
data>>count_30	近30天feedback数	是	[number]	7
data>>count_90	近90天feedback数	是	[number]	20
data>>feedback_num	feedback获取总数	是	[number]	2
data>>five_star	五星feedback获取数	是	[number]	2
data>>four_star	四星feedback获取数	是	[number]	
data>>three_star	三星feedback获取数	是	[number]	
data>>two_star	二星feedback获取数	是	[number]	
data>>one_star	一星feedback获取数	是	[number]	
data>>modified_num	feedback删评数	是	[number]	
data>>seller_name	店铺名称	是	[string]	店铺5
data>>country	国家	是	[string]	美国
data>>score	评分	是	[number]	5
data>>negative_lifetime	feedback累计差评率	是	[number]	0.01
data>>neutral_lifetime	feedback累计中评率	是	[number]	
data>>positive_lifetime	feedback累计好评率	是	[number]	0.99
total	总数	是	[int]	14
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "78390030-503E-2D07-4DD3-6505E102F3F2",
    "response_time": "2022-05-16 17:37:11",
    "data": [
        {
            "count_lifetime": 8904,
            "negative_lifetime": 0.01,
            "neutral_lifetime": 0,
            "positive_lifetime": 0.99,
            "count_12": 3804,
            "count_30": 255,
            "count_90": 779,
            "feedback_num": 285,
            "five_star": 264,
            "four_star": 21,
            "three_star": 0,
            "two_star": 0,
            "one_star": 0,
            "modified_num": 0,
            "seller_name": "店铺2",
            "country": "美国",
            "score": 4.9
        },
        {
            "count_lifetime": 223,
            "negative_lifetime": 0.06,
            "neutral_lifetime": 0.03,
            "positive_lifetime": 0.91,
            "count_12": 86,
            "count_30": 1,
            "count_90": 10,
            "feedback_num": 11,
            "five_star": 9,
            "four_star": 2,
            "three_star": 0,
            "two_star": 0,
            "one_star": 0,
            "modified_num": 0,
            "seller_name": "店铺7",
            "country": "日本",
            "score": 4.8
        },
        {
            "count_lifetime": 9303,
            "negative_lifetime": 0.01,
            "neutral_lifetime": 0.01,
            "positive_lifetime": 0.98,
            "count_12": 3844,
            "count_30": 311,
            "count_90": 783,
            "feedback_num": 323,
            "five_star": 299,
            "four_star": 24,
            "three_star": 0,
            "two_star": 0,
            "one_star": 0,
            "modified_num": 0,
            "seller_name": "店铺1",
            "country": "美国",
            "score": 4.9
        }
    ],
    "total": 14
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "start_date => start_date不能为空",
        "end_date => end_date不能为空"
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
查询评价统计-Review每日新增数
下一章节 
查询评价统计-Feedback每日新增数