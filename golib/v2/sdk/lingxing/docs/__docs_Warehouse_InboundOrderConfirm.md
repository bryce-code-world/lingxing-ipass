入库单确认入库
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/inboundOrder/inbound/setInbound	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSn	入库单单号	否	[array]	IB240730005
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/inboundOrder/inbound/setInbound?app_key=ak_SEpKDE0QPc9Md&access_token=73b3b68e-1692-46fd-86b5-0d236cdf5a30&timestamp=1743660157&sign=%2BpopkuJ4Dcxx%2BTadNm6uoBoXvX6gecBhfzqfP4%2FY4f7LhFQTjRgX%2FS%2Bgzn9j6L85' \
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
data>>success	成功数量	是	[string]	0
data>>fail	失败数量	是	[string]	2
data>>total	总单数	是	[string]	2
data>>failList	失败明细	是	[array]	
data>>failList>>orderSn	单号	是	[string]	IB220927004
data>>failList>>detail	失败原因	是	[string]	
返回结果示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [],
    "request_id": "d58586799b514e83ab3a8c74dbd4af01.110.17442563966110149",
    "response_time": "2025-04-10 11:40:03",
    "data": {
        "success": 0,
        "fail": 2,
        "total": 2,
        "failList": [
            {
                "orderSn": "IB220927004",
                "detail": "单据状态不是待入库，不可执行入库操作"
            },
            {
                "orderSn": "888",
                "detail": "错误：入库单已删除"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
创建待调拨的调拨单
下一章节 
出库单确认出库