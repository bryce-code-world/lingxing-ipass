查询海外仓sku配对列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overseaWarehouseSetting/matchList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wpId	三方服务商id	是	[int]	1666
twIds	三方仓id，多个之间用逗号隔开	否	[string]	123,124
offset	分页偏移量，默认0	否	[int]	
length	分页大小，默认20，上限200	否	[int]	
isMatched	是否配对，0否，1是	否	[int]	
keyword	关键词，搜索sku / 品名 / 第三方产品名 / 产品编码	否	[string]	
请求cURL示例

curl --location --request POST 'https://openapi.lingxing.com/basicOpen/overseaWarehouseSetting/matchList?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wpId":1630,
    "twids":"123,124"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[string]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	
total	总数	是	[string]	2
data		是	[array]	
data>>countryName	国家	是	[string]	
data>>fnsku	fnsku	是	[string]	
data>>isMatched	是否配对	是	[boolean]	0
data>>isMatchedText	是否配对	是	[string]	已配对
data>>localName	本地品名	是	[string]	aa
data>>localSku	本地sku	是	[string]	aa1246
data>>matchMsg	配对信息	是	[string]	
data>>matchNum	配对数量	是	[int]	1
data>>overseaProductCode	第三方仓sku	是	[string]	
data>>overseaProductName	第三方仓品名	是	[string]	
data>>overseaSpec	规格，winit专用	是	[string]	
data>>overseaUniqueCode	三方sku唯一编码	是	[string]	
data>>productId	产品id	是	[int]	4312
data>>sellerId	店铺id	是	[string]	123
data>>sellerName	店铺名称	是	[string]	localname
data>>sid	同sellerId	是	[int]	
data>>twId	三方仓id	是	[int]	123
data>>twpId	三方商品id	是	[int]	132
data>>warehouseCode	第三方仓仓库代码	是	[string]	15154
data>>warehouseName	第三方仓仓库	是	[string]	
data>>warehouseNameLocal	本地仓库名称	是	[string]	localhouse
data>>wid	仓库id	是	[int]	
data>>wpmId	配对id	是	[int]	
返回成功示例
"code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "f91edad8ebb242c29312ce9d6968adc4.114.17411626672521103",
    "response_time": "2025-03-05 16:17:47",
    "data": [
        {
            "countryName": "",
            "fnsku": "",
            "isMatched": true,
            "isMatchedText": "已配对",
            "localName": "HH",
            "localSku": "HH",
            "matchMsg": null,
            "matchNum": 1,
            "overseaProductCode": "123",
            "overseaProductName": "123",
            "overseaSpec": "",
            "overseaUniqueCode": "1659610245_9650000002",
            "productId": 3186,
            "sellerId": 0,
            "sellerName": "-",
            "sid": 0,
            "twId": 7808,
            "twpId": 1392732,
            "warehouseCode": "53",
            "warehouseName": "PH8806",
            "warehouseNameLocal": "测试仓[3] PH8806",
            "wid": 1190,
            "wpmId": 1343
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询系统产品与第三方海外仓产品映射列表
下一章节 
获取第三方SKU标签PDF文件