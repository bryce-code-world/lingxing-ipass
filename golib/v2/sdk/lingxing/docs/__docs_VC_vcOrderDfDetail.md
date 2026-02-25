查询VC订单详情【DF】
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/vcOrderDf/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
vc_store_id	vc店铺id，查询VC店铺列表 接口对应字段【vc_store_id】	是	[string]	134225003201380864
purchase_order_number	订单编号	是	[string]	XB95bX69r
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>vc_store_id	vc店铺id	是	[string]	134225003201380864
data>>seller_name	店铺名称	是	[string]	VC01-美国
data>>local_po_number	本地po号	是	[string]	402229221848662528
data>>purchase_order_number	订单编号	是	[string]	XB95bX69r
data>>purchase_order_date	下单时间	是	[string]	2022-08-24 23:44:25
data>>purchase_order_state	订单状态：
New 新的订单
SHIPPED 已发货
ACCEPTED 已确定
CANCELED 已取消	是	[string]	New
data>>purchase_order_type	订单类型：
0 DF
1 PO	是	[string]	0
data>>bill_to_party_id	结算方式	是	[string]	ASI
data>>ship_from_party_id	供货编码	是	[string]	ENVW
data>>related_warehouse_id	仓库id	是	[string]	1
data>>related_warehouse_name	仓库名称	是	[string]	美国-仓
data>>ship_method	运输方式	是	[string]	UPS_2ND
data>>ship_window_time	要求发货时间	是	[string]	2022-09-16 18:40:00
data>>promised_delivery_date	承诺送达时间	是	[string]	2022-08-31 23:59:59
data>>is_pslip_required	是否需要装箱清单：
0 否
1 是	是	[string]	0
data>>is_gift	是否包含礼物：
0 否
1 是	是	[string]	0
data>>is_scheduled_delivery_shipment	是否预定交付计划：
0 否
1 是	是	[string]	0
data>>is_priority_shipment	是否优先发货：
0 否
1 是	是	[string]	0
data>>ship_to_party_address	收货方地址	是	[object]	
data>>ship_to_party_address>>name	收件人	是	[string]	Kathy Allstot
data>>ship_to_party_address>>address_line1	地址1	是	[string]	10213 EVERGREEN RANCH CT
data>>ship_to_party_address>>address_line2	地址2	是	[string]	
data>>ship_to_party_address>>address_line3	地址3	是	[string]	
data>>ship_to_party_address>>city	城市	是	[string]	GRASS VALLEY
data>>ship_to_party_address>>county	国家	是	[string]	CA
data>>ship_to_party_address>>district	区域	是	[string]	
data>>ship_to_party_address>>state_or_region	州	是	[string]	CA
data>>ship_to_party_address>>postal_code	邮编	是	[string]	95949-8206
data>>ship_to_party_address>>country_code	国家编码	是	[string]	US
data>>ship_to_party_address>>phone	电话	是	[string]	532154241
data>>message_to_customer	交易赠言	是	[string]	
data>>total_price	订单总金额	是	[string]	95.99
data>>currency_code	币种	是	[string]	USD
data>>currency_icon	货币符号	是	[string]	$
data>>item_amount	货物数量	是	[string]	1
data>>remark	备注	是	[string]	this is a remark
data>>items	商品列表	是	[array]	
data>>items>>asin	ASIN	是	[string]	WHZZ17G
data>>items>>msku	MSKU	是	[string]	msku_sss
data>>items>>parent_asin	父ASIN	是	[string]	
data>>items>>item_name	标题	是	[string]	Amagenix Outdoor
data>>items>>large_main_image_url	在线商品主图大图	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>medium_main_image_url	在线商品主图中尺寸	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>small_main_image_url	在线商品主图缩略图	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>pic_url	图片地址	是	[string]	http://www.xxx.com/wd
data>>items>>local_sku	SKU	是	[string]	RC_遥控车
data>>items>>local_name	品名	是	[string]	sku_jjj
data>>items>>purchase_amount	采购量	是	[string]	1
data>>items>>shipped_amount	已发货量	是	[string]	10
data>>items>>waiting_shipped_amount	待发货量	是	[string]	-9
data>>items>>available_amount	可用量	是	[string]	886
data>>items>>sequence_number	序号	是	[string]	1
data>>items>>unit_price	单价	是	[string]	79.99
data>>items>>net_price	成本价	是	[string]	79.99
data>>items>>net_price_currency_code	币种	是	[string]	USD
data>>items>>net_price_currency_icon	货币符号	是	[string]	$
data>>items>>tax_amount	税额	是	[string]	7.2
data>>items>>tax_amount_currency_code	币种	是	[string]	USD
data>>items>>tax_amount_currency_icon	货币符号	是	[string]	$
data>>items>>tax_rate	税率	是	[string]	7.50
data>>items>>tax_rate_percent	税率的百分比	是	[string]	7.50%
data>>items>>deal_total_price	成交总价	是	[string]	95.99
data>>items>>deal_unit_price	成交单价	是	[string]	95.99
data>>tracking_number_List	箱号/跟踪号列表	是	[array]	
data>>tracking_number_List>>box_no	箱号	是	[string]	1ZB348D60337150055
data>>tracking_number_List>>tracking_number	跟踪号	是	[string]	2023052502634193
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "0e64165eefb246919b3feb07df06a766.1699265254827",
    "response_time": "2023-11-06 18:07:37",
    "data": {
        "vc_store_id": "134225003201380864",
        "seller_name": "VC01-美国",
        "local_po_number": "402229221848662528",
        "purchase_order_number": "XB95bX69r",
        "purchase_order_date": "2022-08-24 23:44:25",
        "purchase_order_state": "New",
        "purchase_order_type": "0",
        "bill_to_party_id": "ASI",
        "ship_from_party_id": "ENVW",
        "related_warehouse_id": "1",
        "related_warehouse_name": "仓库1",
        "ship_method": "UPS_2ND",
        "ship_window_time": "2022-09-16 18:40:00",
        "promised_delivery_date": "2022-08-31 23:59:59",
        "is_pslip_required": "0",
        "is_gift": "0",
        "is_scheduled_delivery_shipment": "0",
        "is_priority_shipment": "0",
        "ship_to_party_address": {
            "name": "Kathy Allstot",
            "address_line1": "10213 EVERGREEN RANCH CT",
            "address_line2": "",
            "address_line3": "",
            "city": "GRASS VALLEY",
            "county": "CA",
            "district": null,
            "state_or_region": "CA",
            "postal_code": "95949-8206",
            "country_code": "US",
            "phone": "5302106350"
        },
        "message_to_customer": null,
        "total_price": "95.99",
        "currency_code": "USD",
        "currency_icon": "$",
        "item_amount": "1",
        "remark": null,
        "items": [
            {
                "asin": "B09WHZZ17G",
                "msku": "RC0808-01Z",
                "parent_asin": "B0B1JC46X2",
                "item_name": "Amagenix Outdoor ",
                "large_main_image_url": "https://m.xxx.com/images/I/71ONL.jpg",
                "medium_main_image_url": "https://m.xxx.com/images/I/5wz11L.jpg",
                "small_main_image_url": "https://m.xxx.com/images/I/SL75_.jpg",
                "pic_url": "https://image.xxx.com/9664738f11.jpg",
                "local_sku": "SKU1991FC9",
                "local_name": "[演示数据]适用于iPad的手写笔带手掌柜绝",
                "purchase_amount": "1",
                "shipped_amount": "10",
                "waiting_shipped_amount": "-9",
                "available_amount": 886,
                "sequence_number": "1",
                "unit_price": "79.99",
                "net_price": "79.99",
                "net_price_currency_code": "USD",
                "net_price_currency_icon": "$",
                "tax_amount": "7.20",
                "tax_amount_currency_code": "USD",
                "tax_amount_currency_icon": "$",
                "tax_rate": "0.0750",
                "tax_rate_percent": "7.50%",
                "deal_total_price": "95.99",
                "deal_unit_price": "95.99"
            }
        ],
        "tracking_number_List": [
                {
                    "tracking_number": "1ZB348D60337150055",
                    "box_no": "2023052502634193"
                },
                {
                    "tracking_number": "1ZB348D60337150055-1",
                    "box_no": "202305250263411"
                }
            ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询VC订单详情【PO】
下一章节 
VC订单-确认发货【DF】