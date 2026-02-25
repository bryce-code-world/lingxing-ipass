地址簿-发货地址修改
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/fba/shipment/updateShipFromAddress	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	18
alias_name	地址簿别名，店铺内唯一	是	[string]	xxx
country_name	发货国家/地区	是	[string]	安道尔
sender_name	发货方名称	是	[string]	xxxxx
street_detail1	街道地址1	是	[string]	xxx
street_detail2	街道地址2	否	[string]	xxx
city	城市	是	[string]	xxx
region	区	否	[string]	xxx
province	省/州/地区，美国发货地址限制长度为2位	是	[string]	xxx
zip_code	邮政编码	是	[string]	417000
phone	电话号码	否	[string]	133xxxx6339
id	地址簿-发货地址列表接口返回id	是	[int]	256
请求cURL示例
curl --location 'https://openapi.lingxing.com/erp/sc/routing/fba/shipment/updateShipFromAddress?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "sid": 18,
    "alias_name": "xxx",
    "country_name": "安道尔",
    "sender_name": "xxxxx",
    "street_detail1": "xxx",
    "street_detail2": "xxx",
    "city": "xxx",
    "region": "xxx",
    "province": "xxx",
    "zip_code": "417000",
    "phone": "133xxxx6339",
    "id": 256
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	3E62C9E8-B059-CC31-3586-5DCDD31269B5
response_time	响应时间	是	[string]	2022-12-08 18:35:05
data	响应数据	是	[string]	null
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "3E62C9E8-B059-CC31-3586-5DCDD31269B5",
    "response_time": "2022-12-08 18:35:05",
    "data": null
}
复制
错误
复制成功
 上一章节
地址簿-发货地址创建
FBA
下一章节 
地址簿-配送地址详情
FBA