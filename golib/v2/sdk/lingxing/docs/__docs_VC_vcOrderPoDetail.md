查询VC订单详情【PO】
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/vcOrderPo/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
local_po_number	本地po号，查询VC订单列表 接口字段【local_po_number】	是	[string]	402242689523401371
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
data>>vc_store_id	店铺id	是	[string]	134225003201380864
data>>seller_name	店铺名称	是	[string]	VC01-美国
data>>purchase_order_number	订单编号	是	[string]	2VDWJFMP
data>>local_po_number	本地订单编号	是	[string]	402242689523401357
data>>purchase_order_date	下单时间	是	[string]	2022-10-16 23:57:37
data>>purchase_order_state	订单状态：
Acknowledged 确认
Closed 关闭	是	[string]	Acknowledged
data>>purchase_order_process_state	订单流转状态：
0 待处理
1 确认中
2 确认成功
3 确认失败	是	[int]	1
data>>payment_method	支付类型	是	[string]	"Invoice"
data>>purchase_order_type	订单类型：
0 DF
1 PO	是	[string]	1
data>>remark	备注	是	[string]	this is a remark
data>>related_warehouse_id	仓库id	是	[string]	1
data>>related_warehouse_name	仓库名称	是	[string]	北京仓
data>>ship_to_party_id	收件人	是	[string]	TEB9
data>>total_price	订单总金额	是	[string]	2699.0
data>>currency_code	币种	是	[string]	USD
data>>currency_icon	货币符号	是	[string]	$
data>>item_amount	货物数量	是	[string]	100
data>>ship_window_start	发货窗口开始时间	是	[string]	2022-10-17 00:00:00
data>>ship_window_end	发货窗口结束时间	是	[string]	2022-10-20 00:00:00
data>>delivery_window_start	交货窗口开始时间	是	[string]	2023-10-17 10:55:00
data>>delivery_window_end	交货窗口结束时间	是	[string]	2023-10-27 00:00:00
data>>items	商品数据	是	[array]	
data>>items>>sequence_number	序号	是	[string]	1
data>>items>>asin	ASIN	是	[string]	B08QZ2ZDDM
data>>items>>asin_url	ASIN地址	是	[string]	https://www.amazon.com/xxx
data>>items>>msku	MSKU	是	[string]	msku_jjss
data>>items>>item_name	标题	是	[string]	it_computer
data>>items>>large_main_image_url	在线商品主图大图	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>medium_main_image_url	在线商品主图中尺寸	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>small_main_image_url	在线商品主图缩略图	是	[string]	https://m.xxx.com/images/UONL.jpg
data>>items>>purchase_amount	数量	是	[string]	10
data>>items>>local_po_number	本地po号	是	[string]	402242689523401371
data>>items>>unit_price	单价	是	[string]	10
data>>items>>net_price	成本价	是	[string]	100
data>>items>>net_price_currency_code	币种	是	[string]	USD
data>>items>>net_price_currency_icon	货币符号	是	[string]	$
data>>items>>tax_amount	税额	是	[string]	10
data>>items>>tax_amount_currency_code	币种	是	[string]	USD
data>>items>>tax_amount_currency_icon	货币符号	是	[string]	$
data>>items>>tax_rate	税率	是	[string]	15
data>>items>>tax_rate_percent	税率的百分比	是	[string]	26.99
data>>items>>deal_total_price	成交总价	是	[string]	153.20
data>>items>>deal_unit_price	成交单价	是	[string]	10
data>>items>>is_back_order_allowed	是否可以延期发货：
0 否
1 是	是	[string]	1
data>>items>>shipped_amount	已发货量	是	[string]	0
data>>items>>to_ship_amount	待发货量	是	[string]	10
data>>items>>local_name	品名	是	[string]	RC_遥控车
data>>items>>local_sku	SKU	是	[string]	sku_jjjs
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "b5b12e0109ad4ef288578b325a516020.1698992970537",
    "response_time": "2023-11-03 14:29:33",
    "data": {
        "vc_store_id": "134225003201380864",
        "seller_name": "VC01-美国",
        "purchase_order_number": "2VDWJFMP",
        "local_po_number": "402242689523401357",
        "purchase_order_date": "2022-10-16 23:57:37",
        "purchase_order_state": "Acknowledged",
        "purchase_order_process_state": "1",
        "payment_method": "Invoice",
        "purchase_order_type": "1",
        "remark": null,
        "related_warehouse_id": "1",
        "related_warehouse_name": "仓库1",
        "ship_to_party_id": "TEB9",
        "total_price": "2699.0",
        "currency_code": "USD",
        "currency_icon": "$",
        "item_amount": "100",
        "ship_window_start": "2022-10-17 00:00:00",
        "ship_window_end": "2022-10-20 00:00:00",
        "delivery_window_start": "2023-12-19 10:55:25",
        "delivery_window_end": "2023-12-26 10:55:31",
        "items": [
            {
                "sequence_number": "2",
                "asin": "B09STGM7QT",
                "msku": "KWST121290",
                "item_name": "Patio Garden Backyard",
                "large_main_image_url": "https://m.xxx.com/images/I/71ONL.jpg",
                "medium_main_image_url": "https://m.xxx.com/images/I/5wz11L.jpg",
                "small_main_image_url": "https://m.xxx.com/images/I/SL75_.jpg",
                "purchase_amount": "50",
                "local_po_number": "402242689523401357",
                "unit_price": "25.64",
                "net_price": "1282.00",
                "net_price_currency_code": "USD",
                "net_price_currency_icon": "$",
                "tax_amount": null,
                "tax_amount_currency_code": null,
                "tax_amount_currency_icon": null,
                "tax_rate": null,
                "tax_rate_percent": null,
                "deal_total_price": "1349.50",
                "deal_unit_price": "26.99",
                "is_back_order_allowed": "1",
                "shipped_amount": "0",
                "to_ship_amount": "10",
                "local_name": null,
                "local_sku": null,
                "asin_url": null
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询VC订单列表
下一章节 
查询VC订单详情【DF】