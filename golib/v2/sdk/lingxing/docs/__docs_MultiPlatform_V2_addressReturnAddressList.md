查询退件地址列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/address/returnAddressList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	是	[string]	110418202566107648
请求示例
{
    "store_id": "110418202566107648"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	77ac259a67d5462594c83b80669b6eae.1692331008758
response_time	响应时间	是	[string]	2023-08-18 11:56:49
total	总数	是	[int]	200
data	响应数据	是	[array]	
data>>address_alias	地址别名	是	[string]	客户退货地址111
data>>city	城市	是	[string]	Santa Fe springs
data>>address_id	地址id	是	[string]	20
data>>mobile	电话	是	[string]	177xxxxxxxx
data>>postal_code	邮政编码	是	[string]	906701
data>>province	州/省/地区	是	[string]	CA1
data>>receive_or_deliver_country	发货方/收货方所属国家(或地区)编码	是	[string]	TH
data>>receive_or_deliver_country_name	发货方/收货方所属国家(或地区)名称	是	[string]	泰国
data>>street_detail	街道详细地址	是	[string]	12319 Telegraph Rd
data>>store_country	店铺所属国家	是	[string]	美国
data>>store_id	店铺id	是	[string]	110000000020008001
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "d1b14a1f93c8473aaccefb96c94c332b.1694768691907",
    "response_time": "2023-09-15 17:04:53",
    "data": [
        {
            "addressAlias": "123",
            "city": "1",
            "address_id": "39",
            "mobile": "",
            "postal_code": "1",
            "province": "1",
            "receive_or_deliver_country": "AD",
            "receive_or_deliver_country_name": "安道尔",
            "store_country": "美国",
            "store_id": "110000000020008001",
            "street_detail": "1,"
        }
    ],
    "total": 3
}
复制
错误
复制成功
 上一章节
查询Temu货件
下一章节 
查询WFS货件可添加商品列表