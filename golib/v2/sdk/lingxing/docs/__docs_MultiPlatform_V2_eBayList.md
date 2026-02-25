查询eBay在线商品列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/ebay/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	否	[int]	0
length	分页长度，默认20，最大上限200	否	[int]	20
store_ids	店铺id	否	[array]	["11084135313207808"]
site_code	站点code	否	[array]	["10003-AU"]
listing_status	销售状态	否	[array]	[1]
auto_restocks	是否自动补货：
0 无补货规则
1 启用
2 停用	否	[array]	[1]
listing_type	销售类型：
1 拍卖
2 固价
3 多属性	否	[array]	[1]
search_field	查询字段类型：
1 msku
2 商品ID
3 sku
4 标题
5 品名
6 walmart gtin码	否	[int]	1
search_single_value	搜索值(字符串,单个模糊搜索)	否	[string]	111
listing_time_field	查询时间类型：
1 创建时间
2 结束时间	否	[int]	1
listing_start_time	开始时间(站点时间)，Y-m-d，闭区间【开始结束时间不超过31天】	否	[string]	2024-09-03
listing_end_time	结束时间(站点时间)，Y-m-d，闭区间【开始结束时间不超过31天】	否	[string]	2024-09-05
请求示例
{
    "store_ids": [
        "110000000018003001"
    ],
    "site_code": [
        "10003-AU",
        "10003-AT"
    ],
    "listing_status": [
        1
    ],
    "auto_restocks": [
        1,
        2
    ],
    "listing_type": [
        1,
        2
    ],
    "offset": 0,
    "length": 50,
    "search_field": 1,
    "search_single_value": "test",
    "listing_time_field": 1,
    "listing_start_time": "2024-08-14",
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
request_id	请求链路id	是	[string]	a7cd109d697e45babbf4a9673ea2189f.1726022450730
response_time	响应时间	是	[string]	2024-09-11 10:40:54
data		是	[array]	
data>>item_url	商品跳转链接	是	[string]	https://www.ebay.com/itm/Cotton-Short-Sleeve-T-shirt-Boys-/353943065214
data>>picture_url	图片链接	是	[string]	"https://www.ebay.com/asr/3498e623-4de1-4d4e-8780-d1571dd2af14.5596ac930b26139b73c49a8145501e83.jpeg"
data>>item_id	商品ID	是	[string]	123943065211
data>>msku	MSKU	是	[string]	
data>>attribute	变体属性	是	[string]	Black/Regular/S
data>>title	标题	是	[string]	Cotton Short Sleeve T shirt for Boys
data>>local_name	品名	是	[null]	
data>>local_sku	SKU	是	[null]	
data>>listing_status	状态值	是	[int]	1
data>>listing_status_name	状态	是	[string]	在售
data>>store_id	店铺id	是	[string]	110000000018003001
data>>store_name	店铺	是	[string]	test18&自创ebay测试店铺1号11
data>>site_code	国家id	是	[string]	10003-US
data>>site_name	国家	是	[string]	美国
data>>listing_type	销售类型id	是	[int]	3
data>>listing_type_name	销售类型	是	[string]	多属性
data>>price	价格	是	[string]	0.99
data>>start_price	起拍价	是	[string]	0.00
data>>accept_price	一口价	是	[string]	0.00
data>>quantity	库存	是	[int]	1
data>>auto_restock	是否自动补货	是	[string]	
data>>product_auto_restock_response	自动补货响应	是	[object]	
data>>product_auto_restock_response>>variantUniqueId	商品变体唯一ID(雪花ID)	是	[string]	901172147307683840
data>>product_auto_restock_response>>type	补货类型:
1 固定值
2 指定仓库	是	[int]	1
data>>product_auto_restock_response>>typeName	补货类型:
1 固定值
2 指定仓库	是	[string]	固定值
data>>product_auto_restock_response>>syncIntervalTime	同步间隔时间:
1 每30分钟
2 每1小时
3 每6小时
4 每12小时
5 每24小时	是	[int]	1
data>>product_auto_restock_response>>syncIntervalTimeName	同步间隔时间:
1 每30分钟
2 每1小时
3 每6小时
4 每12小时
5 每24小时	是	[string]	30分钟
data>>product_auto_restock_response>>autoRestock	是否自动补货:
0 未设置补货规则
1 启用
2 停用	是	[int]	2
data>>product_auto_restock_response>>autoRestockName	是否自动补货:
0 未设置补货规则
1 启用
2 停用	是	[string]	停用
data>>product_auto_restock_response>>disableReason	停用原因:
1 手动停用
2 自动停用	是	[int]	1
data>>product_auto_restock_response>>disableReasonName	停用原因:
1 手动停用
2 自动停用	是	[string]	手动停用
data>>product_auto_restock_response>>autoDisableReason	自动停用具体原因	是	[string]	
data>>product_auto_restock_response>>restockQuantity	补货数量	是	[string]	10
data>>product_auto_restock_response>>restockWarehouseId	补货仓库ID	是	[string]	0
data>>location	物品所在地	是	[string]	Shenzhen
data>>dispatch_time_max	处理时间	是	[int]	20
data>>listing_start_time	创建时间	是	[string]	2022-03-07 14:41:01
data>>listing_end_time	结束时间	是	[string]	2022-04-03 14:41:01
total	总数	是	[int]	141
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.181.17262202316017601",
    "response_time": "2024-09-13 17:37:12",
    "data": [
        {
            "item_url": "https://www.ebay.com/itm/Cotton-Short-Sleeve-T-shirt-Boys-/353943065214",
            "item_id": "123943065211",
            "msku": "",
            "attribute": "Black/Regular/S",
            "title": "Cotton Short Sleeve T shirt for Boys",
            "local_name": null,
            "local_sku": null,
            "listing_status": 1,
            "listing_status_name": "在售",
            "store_id": "110000000018003001",
            "store_name": "test18&自创ebay测试店铺1号11",
            "site_code": "10003-US",
            "site_name": "美国",
            "listing_type": 3,
            "listing_type_name": "多属性",
            "price": "0.99",
            "start_price": "0.00",
            "accept_price": "0.00",
            "quantity": 1,
            "auto_restock": null,
            "product_auto_restock_response": {
                "variantUniqueId": "901172147307683840",
                "type": 1,
                "typeName": "固定值",
                "syncIntervalTime": 1,
                "syncIntervalTimeName": "30分钟",
                "autoRestock": 2,
                "autoRestockName": "停用",
                "disableReason": 1,
                "disableReasonName": "手动停用",
                "autoDisableReason": "",
                "restockQuantity": "10",
                "restockWarehouseId": "0"
            },
            "location": "Shenzhen",
            "dispatch_time_max": 20,
            "listing_start_time": "2022-03-07 14:41:01",
            "listing_end_time": "2022-04-03 14:41:01"
        }
    ],
    "total": 141
}
复制
错误
复制成功
 上一章节
查询Temu在线商品
下一章节 
查询Walmart在线商品列表