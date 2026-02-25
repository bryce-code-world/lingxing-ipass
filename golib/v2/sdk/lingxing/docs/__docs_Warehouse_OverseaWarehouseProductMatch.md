海外仓sku配对

支持海外仓商品进行sku配对

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overseaWarehouseSetting/productMatch	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
twId	三方仓id	是	[int]	7788
twpId	三方商品id	是	[int]	1393275
wpId	三方服务商id	是	[int]	1666
productId	商品id	是	[int]	2132
matchNum	整箱配对数量	是	[int]	1
matchAll	是否配对海外仓所有仓库，0否；1是，默认0	否	[int]	0
fnsku	fnsku	否	[string]	
sellerId	店铺id	否	[string]	
请求cURL示例
curl --location --request POST 'http://openapi.lingxing.com/basicOpen/overseaWarehouseSetting/productMatch?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data-raw '{
    "matchNum":1,
    "productId":"3331",
    "twId":"7788",
    "twpId":"1392222",
    "wpId":"1630"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[string]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8c9a1f1550c545eb88e863744f1fd43e.1740104222222
response_time	响应时间	是	[string]	2025-02-21 10:16:53
 上一章节
启用、禁用仓位
下一章节 
海外sku取消配对