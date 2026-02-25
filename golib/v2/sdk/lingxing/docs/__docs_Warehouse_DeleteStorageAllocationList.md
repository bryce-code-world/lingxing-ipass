删除调拨单

支持操作删除状态为待审批，待配货，待调拨的调拨单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/storageAllocationList/delete	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSn	调拨单单号，对应查询调拨单列表接口字段【order_sn】	是	[array]	["TF211231001","TF210916001"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/storageAllocationList/delete?access_token=value&app_key=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data '{
    "orderSn":["TF221130003"]
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	e846890330694325a1d04895b4a9bf2d.1732089924626
response_time	响应时间	是	[string]	2024-11-20 16:05:25
data	响应数据	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "ac8dcba41e2d402cb3853a48d3993376.1732090163746",
    "response_time": "2024-11-20 16:09:24",
    "data": null,
}
复制
错误
复制成功
 上一章节
调拨单结束到货
下一章节 
撤销调拨单