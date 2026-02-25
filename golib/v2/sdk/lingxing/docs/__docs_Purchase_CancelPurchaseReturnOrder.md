作废采购/委外退货单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/purchase/cancelPurchaseReturnOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	采购/委外退货单号	是	[array]	["PR240605003"]
cancel_reason	作废原因	是	[string]	"33"
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/purchase/cancelPurchaseReturnOrder?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "order_sn": [
        "PR241113001"
    ],
    "cancel_reason": "12"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	41be37b6468d434d90064b2da1a0d7b8.1733724738645
response_time	响应时间	是	[string]	2024-12-09 14:12:22
data	该接口data返回为null	是	[null]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "",
    "response_time": "",
    "data": null
}
复制
错误
复制成功
 上一章节
创建已完成的采购退货单
下一章节 
查询采购变更单列表