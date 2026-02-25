查询Shopify在线商品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/shopify/variantList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_ids	店铺Id	否	[array]	["110000000018002001","110477634932417024"]
status	状态
1、Active
2、Draft
3、Archived
4、Deleted	否	[array]	[1,2,3,4]
inventory_policy	库存策略
1、不跟踪库存
2、缺货停止销售
3、缺货继续销售	否	[array]	[1,2,3]
type_id	分类Id	否	[array]	["901235959162748928"]
offset	分页偏移量	否	[int]	0
length	分页长度，上限1000	否	[int]	20
search_field	搜索维度	否	[int]	1
search_single_value	模糊搜索值	否	[string]	B0106008
search_values	精确搜索列表，上限200个	否	[array]	["B0106008", "B0106007"]
quantity	库存数量	否	[string]	123
quantity_condition	库存数量大于或小于
1、大于
2、小于	否	[int]	2
price	售价	否	[long]	123000000
price_condition	售价大于或小于
1、大于
2、小于	否	[int]	2
listing_time_field	时间维度	否	[int]	1
listing_start_time	开始时间	否	[string]	2011-09-08
listing_end_time	结束时间	否	[string]	2024-09-09
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/multiplatform/shopify/variantList?access_token=value&sign=value&timestamp=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "store_ids": ["value"],
    "status": ["value"],
    "inventory_policy": ["value"],
    "offset": value,
    "length": value,
    "search_field": value,
    "search_single_value": "value",
    "quantity": "value",
    "quantity_condition": value,
    "price": "value",
    "price_condition": value,
    "listing_time_field": value,
    "listing_start_time": "value",
    "listing_end_time": "value"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	2024-09-10 10:19:07
data	响应数据	是	[object]	
data>>count	总数	是	[number]	13
data>>list		是	[array]	
data>>list>>weight	重量	是	[string]	0
data>>list>>item_id	商品Id	是	[string]	6700083839169
data>>list>>item_url	商品链接	是	[string]	
data>>list>>product_unique_id	产品uid	是	[string]	901286166546538496
data>>list>>variant_unique_id	子体uid	是	[string]	901286161295692288
data>>list>>company_id	企业Id	是	[string]	901157321452879872
data>>list>>parent_picture_url	父图链接	是	[string]	
data>>list>>picture_url	图片链接	是	[string]	
data>>list>>parent_msku	父MSKU	是	[string]	
data>>list>>msku	MSKU	是	[string]	custom003
data>>list>>parent_attribute	母体属性	是	[string]	Size/Color
data>>list>>attribute	变体属性	是	[string]	42 / red
data>>list>>title	标题	是	[string]	定制产品
data>>list>>local_sku	SKU	是	[string]	test061302sku
data>>list>>pid	本地商品id	是	[string]	
data>>list>>local_name	品名	是	[string]	test061302pm
data>>list>>store_id	店铺Id	是	[string]	110000000018002001
data>>list>>store_name	店铺名称	是	[string]	
data>>list>>price	售价	是	[string]	900.00
data>>list>>quantity	库存	是	[string]	111
data>>list>>currency_icon	货币图标	是	[string]	￥
data>>list>>has_variant	是否有子体	是	[boolean]	false
data>>list>>listing_start_time	开始时间	是	[string]	2021-05-10 11:07:44
data>>list>>listing_end_time	结束时间	是	[string]	2021-10-26 17:26:08
data>>list>>status	商品状态：
1 active
2 draft
3 archived
4 deleted	是	[number]	1
data>>list>>status_name	商品状态：
1 active
2 draft
3 archived
4 deleted	是	[string]	active
data>>list>>inventory_policy	库存策略：
1不跟踪库存
2缺货停止销售
3缺货继续销售	是	[int]	3
data>>list>>inventory_policy_name	库存策略：
1不跟踪库存
2缺货停止销售
3缺货继续销售	是	[string]	缺货继续销售
data>>list>>type_id	商品分类Id	是	[string]	0
data>>list>>type_name	商品分类名	是	[string]	
data>>list>>requires_shipping	是否需要运输
0、否
1、是	是	[boolean]	true
data>>list>>taxable	出售产品变体时是否要征税
0、否
1、是	是	[boolean]	true
data>>list>>min_weight	最小重量	是	[long]	
data>>list>>max_weight	最大重量	是	[long]	
data>>list>>weight_unit	重量单位	是	[string]	kg
data>>list>>location_infos	商品所在地详细信息	是	[array]	
data>>list>>location_infos>>quantity_adjust_after	库存调整后的值	是	[string]	
data>>list>>location_infos>>quantity_adjust_before	库存调整前的值	是	[string]	
data>>list>>location_infos>>quantity_user_id	库存调整操作人id	是	[string]	
data>>list>>location_infos>>quantity_user_id_name	库存调整操作人名字	是	[string]	
data>>list>>location_infos>>quantity_queue_type	库存队列类型	是	[int]	0
data>>list>>location_infos>>quantity_queue_status	库存队列状态	是	[string]	0
data>>list>>location_infos>>location_name	商品所在地名称	是	[string]	
data>>list>>location_infos>>location_quantity	location库存	是	[string]	111
data>>list>>location_infos>>location_id	商品所在地id	是	[string]	64418349246
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.148.17261287335766689",
    "response_time": "2024-09-12 16:12:14",
    "data": {
        "count": 1,
        "list": [
            {
                "weight": "2",
                "item_id": "6768632430804",
                "item_url": "",
                "product_unique_id": "9012162617856",
                "variant_unique_id": "9019162671616",
                "company_id": "9011572879872",
                "parent_picture_url": "",
                "picture_url": "",
                "parent_msku": null,
                "msku": "B0106008",
                "parent_attribute": "SHIP TO/Color",
                "attribute": "EU / Blue",
                "title": "ERYONE PETG Filament, 1.75mm ±0.03mm Filament for 3D Printer, 1KG(2.2LBS)/ Spool",
                "local_sku": "test061302sku",
                "pid": null,
                "local_name": "test061302pm",
                "store_id": "110000000018002001",
                "store_name": "shopify测试店铺",
                "price": "27.99",
                "quantity": "17",
                "currency_icon": "￥",
                "has_variant": false,
                "listing_start_time": "2021-09-10 15:40:51",
                "listing_end_time": "2022-04-12 11:46:56",
                "status": 1,
                "status_name": "active",
                "inventory_policy": 2,
                "inventory_policy_name": "缺货停止销售",
                "type_id": "901235959162748928",
                "type_name": "PETG",
                "requires_shipping": true,
                "taxable": false,
                "min_weight": null,
                "max_weight": null,
                "weight_unit": "lb",
                "location_infos": []
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询结算利润（利润报表）-订单
下一章节 
查询Shein在线商品