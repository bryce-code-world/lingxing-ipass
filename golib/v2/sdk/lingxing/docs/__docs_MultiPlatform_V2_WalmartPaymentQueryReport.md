查询可用报告列表 - Walmart Payment
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/fms/openapi/walmartPayment/queryReport	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
new_report	是否查询最新的报告：1 是，2 否	是	[int]	1
store_id	店铺id	是	[array]	[110000000018008000]
请求示例
{
    "new_report": 1,
    "store_id": [
        110000000018008000
    ]
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
request_id	请求链路id	是	[string]	f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946
response_time	响应时间	是	[string]	2023-05-22 15:59:59
data	数据	是	[array]	
data>>list	报告详情	是	[array]	
data>>list>>report_date	报告时间	是	[string]	2023-05-16
data>>list>>report_id	报告id	是	[string]	MI7cc13c6a4f7d8769629e807e8263fc87
data>>store_id	店铺id	是	[string]	110000000018008000
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8ab5b8fe5bc749c9bdd9e75f8ec3972b.1684737404581",
    "response_time": "2023-05-22 14:36:44",
    "data": [
        {
            "list": [
                {
                    "report_date": "2023-05-16",
                    "report_id": "MIc29691afa98fd0f97ea7513d569c2280"
                }
            ],
            "store_id": "110000000018008001"
        },
        {
            "list": [
                {
                    "report_date": "2023-05-17",
                    "report_id": "MI7cc13c6a4f7d8769629e807e8263fc87"
                }
            ],
            "store_id": "110000000018008002"
        }
    ]
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "参数检验不通过"
    "error_details": ["newReport: 标识的值只能为1或者2"],
    "request_id": "6f81180cc2f542669f9173a456d7f92d.1684737459164",
    "response_time": "2023-05-22 14:37:39",
    "data": null
}
复制
错误
复制成功
 上一章节
查询销量统计列表V2
下一章节 
查询报告详情 - Walmart Payment