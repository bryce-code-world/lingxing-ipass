删除发货单

支持操作删除状态为待配货，待审批，待发货的发货单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/fbaShipment/deleteShipmentList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_nos	发货单单号，对应查询FBA发货单列表接口字段【shipment_sn】	是	[array]	["SP240123007"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/fbaShipment/deleteShipmentList?access_token=value&app_key=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data '{
    "shipment_nos":["SP240123007"]
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	数据校验失败时的错误详情	是	[array]	["操作成功"]
request_id	请求链路id	是	[string]	ac8dcba41e2d402cb3853a48d3993376.1732090163746
response_time	响应时间	是	[string]	2024-11-20 16:09:24
data	响应数据	是	[string]
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [
        "操作成功"
    ],
    "request_id": "ac8dcba41e2d402cb3853a48d3993376.1732090163746",
    "response_time": "2024-11-20 16:09:24",
    "data": null,
}
复制
错误
复制成功
 上一章节
编辑发货单
FBA
下一章节 
作废发货单
FBA