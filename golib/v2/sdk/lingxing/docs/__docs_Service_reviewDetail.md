查询评价统计-Review每日新增数
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/cs/reviewReport/detail	HTTPS	POST	3
请求参数
参数名	说明	必填	类型	示例
mid	国家id	是	[int]	1
asin	asin	是	[string]	B085NQDDXS
start_date	开始时间【时间间隔不超过1年】	是	[string]	2024-01-01 00:00:00
end_date	结束时间【时间间隔不超过1年】	是	[string]	2024-08-05 00:00:00
请求示例
{
    "mid": 1,
    "asin": "B085NQDDXS",
    "start_date": "2024-01-01 00:00:00",
    "end_date": "2024-08-05 00:00:00"
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
data>>report_date	日期	是	[string]	2021-04-05
data>>review_num	review新增数	是	[number]	10
data>>five_star	5星review新增数	是	[number]	10
data>>four_star	4星review新增数	是	[number]	
data>>three_star	3星review新增数	是	[number]	
data>>two_star	2星review新增数	是	[number]	
data>>one_star	1星review新增数	是	[number]	
data>>ratings	rating总数	是	[number]	300
data>>ratings_inc	rating新增数
备注：可以为负值，代表减少数	是	[number]	
total	总数	是	[int]	29
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "49E093D9-3200-6178-0069-7CDF3A40ECD7",
    "response_time": "2022-09-27 17:49:30",
    "data": [
        {
            "report_date": "2021-08-01",
            "review_num": 5,
            "five_star": 3,
            "four_star": 1,
            "three_star": 1,
            "two_star": 0,
            "one_star": 0,
            "ratings": 123,
            "score": 4.6,
            "ratings_inc": 0
        },
        {
            "report_date": "2021-08-02",
            "review_num": 6,
            "five_star": 5,
            "four_star": 0,
            "three_star": 0,
            "two_star": 0,
            "one_star": 1,
            "ratings": 136,
            "score": 4.6,
            "ratings_inc": 13
        }
    ],
    "total": 2
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "最大日期范围不能超过1年"
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
查询评价统计-Review列表
下一章节 
查询评价统计-Feedback列表