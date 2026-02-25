获取备货单号
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/listOrderNos	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inbound_order_no	客户参考号 数组	否	[array]	 
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	7E80FC0B-55C1-BAC5-0E4B-7AE59706F051
response_time	响应时间	是	[string]	2021-06-15 20:16:38
data	响应数据	是	[array]	
data>>overseas_order_no	备货单号	是	[string]	OWS211108003
data>>inbound_order_no	客户参考号	是	[string]	inbound-521-1636352128
data>>create_time	创建时间	是	[string]	2021-11-08 14:48:49
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "30B862A0-8974-34D5-9FD9-1C31B2E2EB10",
    "response_time": "2024-08-01 16:07:17",
    "data": [
        {
            "overseas_order_no": "OWS210303001",
            "inbound_order_no": "",
            "create_time": "2021-03-03 10:23:08"
        },
        {
            "overseas_order_no": "OWS240801001",
            "inbound_order_no": "",
            "create_time": "2024-08-01 11:34:41"
        }
    ],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询海外仓备货单详情
下一章节 
删除备货单