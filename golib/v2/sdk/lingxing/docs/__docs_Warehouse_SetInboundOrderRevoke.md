撤销入库单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/inboundOrder/inbound/setOrderRevoke	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	入库单号
对应查询入库单列表data>>order_sn字段	是	[string]	"IB241122002"
delete_receipt_order	是否同步删除收货单
删除则传值 1，否则不传值	否	[int]	1
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/inboundOrder/inbound/setOrderRevoke?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "order_sn": "TF241119002" ,
    "delete_receipt_order": 1 
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	"success"
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	"e846890330694325a1d04895b4a9bf2d.1732089924626"
response_time	响应时间	是	[string]	"2024-11-20 16:05:25"
data	响应数据	是	[string]
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
查询入库单列表
下一章节 
添加出库单