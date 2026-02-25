订单称重
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/wms/order/setOrderWeighed	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_number	系统单号 与销售出库单二选一	否	[string]	
wo_number	销售出库单 与系统单号二选一	否	[string]	
pkg_real_weight	重量	是	[string]	
pkg_real_weight_unit	单位 支持 g,kg,oz,lb	是	[string]	
sync_product_gross_weight	一单一件同步重量到产品模块 0:否,1:是 默认否	否	[string]
请求示例
curl --location 'https://openapi.lingxing.com/erp/sc/routing/wms/order/setOrderWeighed?access_token=&app_key=&sign=&timestamp=' \
--header 'Content-Type: application/json' \
--data '{
    "order_number": "103536668312138123",
    "pkg_real_weight": "123",
    "pkg_real_weight_unit": "g",
    "sync_product_gross_weight": "0"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	99C3E97E-821D-911A-CEA5-C9212020038F
response_time	响应时间	是	[string]	2025-02-18 14:15:38
data	响应数据(空)	是	[array]	
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "",
    "error_details": "",
    "request_id": "1D30B180-86C2-B50C-C26E-804FD2B3EE0F",
    "response_time": "2025-02-19 15:57:16",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
订单发货
下一章节 
编辑订单（新版）