调整单确认调整

异步任务，查询结果请调用：查询确认调整异步任务结果

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/adjustOrder/adjust/setAdjust	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSn	调整单单号	否	[array]	IB240730005
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/outboundOrder/outbound/setOutbound?app_key=ak_SEpKDE0QPc9Md&access_token=73b3b68e-1692-46fd-86b5-0d236cdf5a30&timestamp=1743660157&sign=%2BpopkuJ4Dcxx%2BTadNm6uoBoXvX6gecBhfzqfP4%2FY4f7LhFQTjRgX%2FS%2Bgzn9j6L85' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderSn":["IB220927004"]
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	
data>>taskNo	异步任务编号	是	[string]	292559361475654656
data>>type		是	[string]	50
data>>actionType		是	[string]	5003
返回结果示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "d58586799b514e83ab3a8c74dbd4af01.141.17442652599960609",
    "response_time": "2025-04-10 14:07:40",
    "data": {
        "taskNo": "292559361475654656",
        "type": 50,
        "actionType": 5003
    }
}
复制
错误
复制成功
 上一章节
出库单确认出库
下一章节 
查询调整单确认调整结果