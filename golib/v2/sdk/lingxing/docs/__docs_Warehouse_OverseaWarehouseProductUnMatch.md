海外仓sku取消配对

支持取消海外仓商品sku配对

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overseaWarehouseSetting/productUnMatch	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wpId	三方服务商id	是	[string]	1688
wpmId	配对id	是	[string]	1222
请求cURL示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/overseaWarehouseSetting/productUnMatch?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wpId":1666,
    "wpmId":1333
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[string]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8c9a1f1550c545eb88e863744f1fd43e.1740104555555
response_time	响应时间	是	[string]	2025-02-20 15:16:53
 上一章节
海外仓sku配对
下一章节 
查询产品仓位列表