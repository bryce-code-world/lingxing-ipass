FBA-作废发货单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/fbaShipment/shipmentSn/invalid	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipmentNos	发货单号	是	[array]	["SP241217006"]
isReturnStock	产品库存是否恢复 1恢复 0不恢复	是	[int]	1
isReturnStockAux	辅料库存是否恢复 1恢复 0不恢复	是	[int]	0
cancelReason	作废原因	否	[string]	SP241217006
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/fbaShipment/shipmentSn/invalid?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "shipmentNos":["SP241217006"],
    "isReturnStock":1,
    "isReturnStockAux":0,
    "cancelReason":"SP241217006"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	值可能性	限制	示例
code	状态码，0：成功	是	[int]	;		0
message	消息提示	是	[string]	;		success
error_details	数据校验失败时的错误详情	是	[array]			
request_id	请求链路id	是	[string]	;		720e6a6bac5b4c729d4ecea5690c4e80.1734939490648
response_time	响应时间	是	[string]	;		2024-12-23 15:38:10
data	null	是	[null]			
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
删除发货单
FBA
下一章节 
更新发货单物流信息
FBA