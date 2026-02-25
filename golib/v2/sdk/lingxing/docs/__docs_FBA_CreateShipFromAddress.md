地址簿-发货地址创建
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/fba/shipment/createShipFromAddress	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	18
alias_name	地址簿别名，店铺内唯一	是	[string]	xxxxxx
country_name	发货国家/地区	是	[string]	安道尔
sender_name	发货方名称	是	[string]	xxx
street_detail1	街道地址1	是	[string]	xxxxxx
street_detail2	街道地址2	否	[string]	xxxxxx
city	城市	是	[string]	xxxx
region	区	否	[string]	xxxx
province	省/州/地区，美国发货地址限制长度为2位	是	[string]	xxxxx
zip_code	邮政编码	是	[string]	417000
phone	电话号码	否	[string]	13390876339
请求cURL示例
curl --location 'https://openapi.lingxing.com/erp/sc/routing/fba/shipment/createShipFromAddress?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "sid": 18,
    "alias_name": "xxxxxx",
    "country_name": "安道尔",
    "sender_name": "xxx",
    "street_detail1": "xxxxxx",
    "street_detail2": "xxxxxx",
    "city": "xxxx",
    "region": "xxxx",
    "province": "xxxxx",
    "zip_code": "417000",
    "phone": "13390876339"
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
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
data	响应数据	是	[array]	
data>>id	创建成功时返回地址ID	是	[int]	256
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D77D57BD-2C54-6004-3149-46BA3CA8A315",
    "response_time": "2022-12-08 18:32:16",
    "data": {
        "id": "319"
    }
}
复制
错误
复制成功
 上一章节
地址簿-发货地址列表
FBA
下一章节 
地址簿-发货地址修改
FBA