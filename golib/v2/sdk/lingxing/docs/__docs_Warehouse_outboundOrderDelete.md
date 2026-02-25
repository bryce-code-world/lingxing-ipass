删除出库单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/outboundOrder/outbound/delete	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSn	出库单单号	否	[array]	["OB250227001"]
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/outboundOrder/outbound/delete?app_key=&access_token=&timestamp=&sign=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderSn":["OB250227001","OB250218001"]
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
data	响应数据	是	[object]	
data>>success	成功数量	是	[string]	0
data>>fail	失败数量	是	[string]	2
data>>total	总单数	是	[string]	2
data>>failList	失败明细	是	[array]	
data>>failList>>orderSn	失败单号	是	[string]	IB220927004
data>>failList>>detail	失败原因	是	[string]	
返回示例

Json Object

{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "69dbe8b2b0024c329c302bae86b5af91.1747188526466",
    "response_time": "2025-05-14 10:08:46",
    "data": {
        "success": 1,
        "fail": 2,
        "total": 3,
        "failList": [
            {
                "orderSn": "OB250227001",
                "detail": "订单不存在或已删除"
            },
            {
                "orderSn": "OB250218001",
                "detail": "订单不存在或已删除"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询调整单确认调整结果
下一章节 
创建已完成的成本补录单