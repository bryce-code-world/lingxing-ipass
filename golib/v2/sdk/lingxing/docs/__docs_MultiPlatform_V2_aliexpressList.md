查询AliExpress在线商品

支持查询 “自运营”->“子体” 的数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/aliExpress/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	50
store_ids	店铺id	否	[array]	["110000000018005005"]
status	状态：
1 正在销售
2 已下架
3 审核中
4 审核不通过	否	[array]	[1,2]
listing_time_field	查询时间类型：
1 创建时间
2 结束时间	否	[int]	1
listing_start_time	开始日期，Y-m-d，闭区间【开始结束时间不超过31天】	否	[string]	2024-08-14
listing_end_time	结束日期，Y-m-d，闭区间【开始结束时间不超过31天】	否	[string]	2024-09-13
search_field	搜索字段类型：
1 MSKU
2 商品ID
3 SKU
4 标题	否	[int]	1
search_single_value	搜索值(字符串,单个模糊搜索）	否	[string]	114
请求示例
{
    "store_ids": [
        "110000000018005005",
        "110000000018005006"
    ],
    "status": [
        1,
        2
    ],
    "offset": 0,
    "length": 50,
    "search_field": 1,
    "search_single_value": "test",
    "listing_time_field": 1,
    "listing_start_time": "2024-09-01",
    "listing_end_time": "2024-09-13"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	d8a6ed6210e3413db4d5ce63921eb7f8.1726195350647
response_time	响应时间	是	[string]	2024-09-13 10:42:31
data	响应数据	是	[array]	
data>>item_url	图片	是	[string]	https://www.aliexpress.com/item/1005002355553844.html
data>>item_id	商品ID	是	[string]	1005002355553844
data>>msku	MSKU	是	[string]	ICE0001
data>>title	标题	是	[string]	With cover ball ice grid plastic ice block mold refrigerator ice hockey mold ice box round ice mold ice box box123--xhp0721
data>>local_name	品名	是	[string]	xhp测试产品004
data>>local_sku	SKU	是	[string]	xhp004
data>>status	状态id	是	[int]	2
data>>status_name	状态	是	[string]	已下架
data>>store_id	店铺id	是	[string]	110000000018005005
data>>store_name	店铺	是	[string]	test18自创AliExpress测试店铺405号
data>>currency_icon	货币符号	是	[string]	$
data>>price	价格	是	[string]	12.12
data>>quantity	库存	是	[string]	5
data>>delivery_time	备货时间	是	[int]	7
data>>listing_start_time	创建时间	是	[string]	2021-03-25 09:58:03
data>>listing_end_time	结束时间	是	[string]	2022-07-22 20:00:30
total	总数	是	[int]	10
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.181.17262207959337581",
    "response_time": "2024-09-13 17:46:36",
    "data": [
        {
            "item_url": "https://www.aliexpress.com/item/1005002355553844.html",
            "item_id": "1005002355553844",
            "msku": "ICE0001",
            "title": "With cover ball ice grid plastic ice block mold refrigerator ice hockey mold ice box round ice mold ice box box123--xhp0721",
            "local_name": null,
            "local_sku": null,
            "status": 2,
            "status_name": "已下架",
            "store_id": "110000000018005005",
            "store_name": "test18自创AliExpress测试店铺405号",
            "currency_icon": "$",
            "price": "12.12",
            "quantity": "5",
            "delivery_time": 7,
            "listing_start_time": "2021-03-25 09:58:03",
            "listing_end_time": "2022-07-22 20:00:30"
        }
    ],
    "total": 10
}
复制
错误
复制成功
 上一章节
查询Walmart在线商品列表
下一章节 
Temu修改库存