查询买家之声列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/voiceOfBuyer/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
fulfillment_channel	配送方式：
FBA FBA
MFN FBM	否	[string]	FBA
sids	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[array]	25
pxc_health	满意度状况：
-1 反馈不足
0 极差
1 不合格
2 一般
3 良好
4 极好	否	[array]	["0"]
search_field	搜索类型：
asin ASIN
msku MSKU
sku SKU	否	[string]	msku
search_value	搜索值	否	[array]	MSKU4C1DA7B
请求示例
{
    "offset": 0,
    "length": 20,
    "fulfillment_channel": "FBA",
    "sids": 25,
    "pxc_health": ["0"],
    "search_field": "msku",
    "search_value": ["MSKU4C1DA7B"]
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
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>sid	店铺id	是	[string]	1
data>>seller_name	店铺名称	是	[string]	美国仓
data>>country	国家名称	是	[string]	美国
data>>image_url	图片地址	是	[string]	图片地址
data>>asin	ASIN	是	[string]	B0ABCxxx50521
data>>asin_url	ASIN地址	是	[string]	https://www.xxx.com/dp/B0ABCxxx50521
data>>title	标题	是	[string]	amazon product title 07E1A4C731
data>>msku	MSKU	是	[string]	MSKU52FB8CE
data>>fnsku	FNSKU	是	[string]	X003IMYVZX
data>>fulfillment_channel	配送方式：
FBA FBA
MFN FBM	是	[string]	FBA
data>>ncx_rate	不满意率	是	[string]	0.9000
data>>ncx_count	不满意订单数量	是	[int]	3
data>>order_count	订单总数	是	[int]	5
data>>most_common_return_reason_bucket	主要退货原因	是	[string]	ui_too_small
data>>last_action_date	最近停售日期	是	[string]	-
data>>event_date	上次更新日期	是	[string]	2023-05-28
data>>pcx_health_text	满意度状况说明	是	[string]	一般
data>>product_name	品名	是	[string]	RC_遥控车
data>>sku	SKU	是	[string]	SKU_jjj
data>>listing_exists	是否删除：
true 删除
false 不删除	是	[boolean]	false
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "219a9c09c8814245be6417c82e1617d5.1699342655448",
    "response_time": "2023-11-07 15:37:37",
    "data": [
        {
            "sid": "1",
            "seller_name": "店铺1",
            "country": "美国",
            "image_url": "https://xxx.com/E6%9C%BA.jpg",
            "asin": "B0ABC50521",
            "asin_url": "https://www.xxx.com/dp/B0ABCxxx50521",
            "title": "amazon product title 07E1A4C731",
            "msku": "MSKU52FB8CE",
            "fnsku": "X003IMYVZX",
            "fulfillment_channel": "FBA",
            "ncx_rate": "0.9000",
            "ncx_count": 3,
            "order_count": 5,
            "most_common_return_reason_bucket": "ui_too_small",
            "last_action_date": "-",
            "event_date": "2023-05-28",
            "pcx_health_text": "良好",
            "product_name": "MSK测试",
            "sku": "MSK",
            "listing_exists": false
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询店铺绩效详情
下一章节 
查询RMA管理