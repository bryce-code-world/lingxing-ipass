删除备货单

支持操作删除状态为待审批，待配货，待发货，已驳回的海外仓备货单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overSeaWarehouse/stockOrder/delete	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_nos	备货单单号，对应获取备货单号接口字段【overseas_order_no】	是	[array]	["OWS210923003","OWS210923008"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/overSeaWarehouse/stockOrder/delete?access_token=value&app_key=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data '{
    "overseas_order_nos":["OWS231016001"]
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	6a08e9dd9e2c4f2eb7730da71bcad8f6.1732089599447
response_time	响应时间	是	[string]	2024-11-20 16:00:03
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
获取备货单号
下一章节 
更新备货单