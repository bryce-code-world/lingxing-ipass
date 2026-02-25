查询VC订单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/vcOrder/pageList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
purchase_order_type	订单类型：
0 DF
1 PO	是	[array]	["0"]
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
vc_store_ids	vc店铺id，查询VC店铺列表 接口对应字段【vc_store_id】	否	[array]	["134225003201380860"]
search_field_time	查询时间类型：
1 订购时间
2 要求发货时间
3 订单更新时间	否	[string]	1
start_date	开始时间，开区间，格式：Y-m-d，时间间隔最长不超过90天	否	[string]	2023-10-16
end_date	结束时间，闭区间，格式：Y-m-d，时间间隔最长不超过90天	否	[string]	2023-10-17
search_field	搜索类型：
purchase_order_number 订单号
asin ASIN
local_name 品名
customer_order_number 客户订单号【DF类型订单】
vendor_product_id 商品编码	否	[string]	purchase_order_number
search_value	搜索值	否	[array]	["67HDLN3R"]
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
data>>gmt_create	订单创建时间	是	[string]	2022-09-20 17:06:05
data>>gmt_modified	订单更新时间	是	[string]	2023-09-29 21:20:24
data>>id	订单ID	是	[string]	107
data>>purchase_order_number	订单编号	是	[string]	XB95bX69r
data>>customer_order_number	客户订单号【DF类型订单】	是	[string]	305-6601334-7925136
data>>vc_store_id	vc店铺id	是	[string]	134225003201380864
data>>seller_name	店铺名称	是	[string]	VC01-美国
data>>purchase_order_type	订单类型：
0 DF
1 PO	是	[int]	0
data>>purchase_order_state	DF订单状态：
New 新的订单
SHIPPED 已发货
ACCEPTED 已确定
CANCELLED 已取消
PO订单状态：
Acknowledged 确认
Closed 关闭	是	[string]	New
data>>purchase_order_process_state	订单流转状态：
0 待处理
1 待发货
2 已完成
3 已取消	是	[int]	2
data>>purchase_order_date	订单下单时间	是	[string]	2022-08-24 23:44:25
data>>ack_status	ack状态：0：待确认
1：确认中
2：已确认
3：确认失败
4：平台已确认	是	[int]	3
data>>ack_status_desc	ack状态说明，0：待确认
1：确认中
2：已确认
3：确认失败
4：平台已确认	是	[string]	确认失败
data>>ack_update_time	ack更新时间	是	[string]	2022-12-13 16:32:24
data>>focus_party_id	仓库id	是	[string]	ENVW
data>>erp_warehouse_name	配对后的本地仓名称	是	[string]	仓库1
data>>erp_warehouse_id	配对后的本地仓id	是	[string]	
data>>ship_window_time	DF要求发货时间	是	[string]	2022-09-16 18:40:00
data>>ship_window_start	PO要求起始发货时间	是	[string]	2022-10-24 00:00:00
data>>ship_windows_end	PO要求截止发货时间	是	[string]	2022-10-27 00:00:00
data>>total_price	订单总金额	是	[string]	95.99
data>>currency_code	币种	是	[string]	USD
data>>currency_icon	币种符号	是	[string]	$
data>>item_amount	货物总数量	是	[int]	1
data>>local_po_number	本地po号	是	[string]	402229221848662528
data>>remark	订单备注	是	[string]	remark note
data>>shipment_confirm_status	确认发货状态：
1 未确认
2 确认中
3 确认失败
4 确认成功
5 平台已确认	是	[int]	1
data>>shipment_label_status	标签状态：
1 未请求
2 请求中
3 请求失败
4 请求成功	是	[int]	1
data>>print_num	标签打印次数	是	[int]	4
data>>purchase_order_sku_list	订单商品明细数据	是	[array]	
data>>purchase_order_sku_list>>id	ID	是	[string]	5069
data>>purchase_order_sku_list>>vc_store_id	店铺id	是	[string]	134225003xx1380864
data>>purchase_order_sku_list>>seller_name	店铺名称	是	[string]	VC01-美国
data>>purchase_order_sku_list>>asin	ASIN	是	[string]	B09WHZZ17G
data>>purchase_order_sku_list>>upc	UPC	是	[string]	upc_xxx
data>>purchase_order_sku_list>>ean	EAN	是	[string]	ean_xxx
data>>purchase_order_sku_list>>parent_asin	父ASIN	是	[string]	B0B1JC46X2
data>>purchase_order_sku_list>>item_name	标题	是	[string]	220GSM HDPE
data>>purchase_order_sku_list>>large_main_image_url	在线商品主图大图	是	[string]	https://xxx.com/1L.jpg
data>>purchase_order_sku_list>>medium_main_image_url	在线商品主图中尺寸	是	[string]	https://xxx.com/2L.jpg
data>>purchase_order_sku_list>>small_main_image_url	在线商品主图略缩图	是	[string]	https://xxx.com/3L.jpg
data>>purchase_order_sku_list>>has_principal	是否分配负责人：
0 否
1 是	是	[int]	0
data>>purchase_order_sku_list>>purchase_amount	采购量	是	[int]	1
data>>purchase_order_sku_list>>sequence_number	序列号	是	[string]	1
data>>purchase_order_sku_list>>vendor_product_id	商品编码	是	[string]	ZXCWXZJLZLDSEXSGUV0
data>>purchase_order_sku_list>>local_po_number	本地po编号	是	[string]	402229221848662528
data>>purchase_order_sku_list>>purchase_order_number	订单号	是	[string]	XB95bX69r
data>>purchase_order_sku_list>>unit_price	单价	是	[string]	79.99
data>>purchase_order_sku_list>>net_price	成本价	是	[string]	79.99
data>>purchase_order_sku_list>>net_price_currency_code	成本价货币类型	是	[string]	USD
data>>purchase_order_sku_list>>net_price_currency_icon	成本价货币符号	是	[string]	$
data>>purchase_order_sku_list>>tax_amount	税额	是	[string]	7.20
data>>purchase_order_sku_list>>tax_amount_currency_code	税额货币类型	是	[string]	USD
data>>purchase_order_sku_list>>tax_amount_currency_icon	税额货币符号	是	[string]	$
data>>purchase_order_sku_list>>tax_rate	税率	是	[string]	0.0750
data>>purchase_order_sku_list>>tax_rate_percent	税率百分比	是	[string]	7.50%
data>>purchase_order_sku_list>>deal_total_price	成交总价	是	[string]	95.99
data>>purchase_order_sku_list>>deal_unit_price	成交单价	是	[string]	95.99
data>>purchase_order_sku_list>>is_back_order_allowed	是否可以延期发货：
0 否
1 是	是	[string]	0
data>>purchase_order_sku_list>>shipped_amount	已发货数量	是	[int]	10
data>>purchase_order_sku_list>>to_ship_amount	待发货数量	是	[int]	-9
data>>purchase_order_sku_list>>local_name	品名	是	[string]	适用于iPad的手写笔
data>>purchase_order_sku_list>>local_sku	SKU	是	[string]	SKU1991FC9
data>>purchase_order_sku_list>>product_id	本地产品id	是	[string]	1
data>>purchase_order_sku_list>>available_amount	库存可用量	是	[int]	886
data>>purchase_order_sku_list>>asin_url	ASIN跳转地址	是	[string]	https://xxx.com/dp/B09WxxZ17G
data>>purchase_order_sku_list>>pic_url	图片地址	是	[string]	https://xxx.com/8f11.jpg
data>>purchase_order_sku_list>>accepted_quantity	接受量	是	[int]	10
data>>purchase_order_sku_list>>rejected_quantity	拒绝量	是	[int]	0
data>>purchase_order_sku_list>>received_quantity	签收量	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7d31f5ee09e54b8ea773e060c8a05498.1698311450835",
    "response_time": "2023-10-26 17:10:53",
    "data": [
        {
            "gmt_create": "2022-09-20 17:06:05",
            "gmt_modified": "2023-09-29 21:20:24",
            "id": "107",
            "purchase_order_number": "XB95bX69r",
            "customer_order_number": "305-6601334-7925136",
            "vc_store_id": "134225003201380864",
            "seller_name": "VC01-美国",
            "purchase_order_type": 0,
            "purchase_order_state": "New",
            "purchase_order_process_state": 2,
            "purchase_order_date": "2022-08-24 23:44:25",
            "ack_status": 3,
            "ack_status_desc": "确认失败",
            "ack_update_time": "2022-12-13 16:32:24",
            "focus_party_id": "ENVW",
            "erp_warehouse_name": "仓库1",
            "erp_warehouse_id": null,
            "ship_window_time": "2022-09-16 18:40:00",
            "ship_window_start": null,
            "ship_windows_end": null,
            "total_price": "95.99",
            "currency_code": "USD",
            "currency_icon": "$",
            "item_amount": 1,
            "local_po_number": "402229221848662528",
            "remark": null,
            "shipment_confirm_status": 4,
            "shipment_label_status": 4,
            "print_num": 4,
            "purchase_order_sku_list": [
                {
                    "id": "5069",
                    "vc_store_id": "134225003201380864",
                    "asin": "B09WHZZ17G",
                    "upc": "",
                    "ean": null,
                    "parent_asin": "B0B1JC46X2",
                    "item_name": "Also Indoor Windows Blackout Blinds 8'(W) x8'(H),Mocha",
                    "large_main_image_url": "https://xxx.com/images/I/719qWwKNVyL.jpg",
                    "medium_main_image_url": "https://xxx.com/kIX+YzeJL.jpg",
                    "small_main_image_url": "https://xxx.com/75.jpg",
                    "has_principal": 0,
                    "source_data_update_time": null,
                    "purchase_amount": 1,
                    "sequence_number": "1",
                    "vendor_product_id": "ZXCWXZJLZLDSEXSGUV0",
                    "local_po_number": "402229221848662528",
                    "purchase_order_number": "XB95bX69r",
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
                    "deal_unit_price": "95.99",
                    "is_back_order_allowed": null,
                    "shipped_amount": 10,
                    "to_ship_amount": -9,
                    "local_name": "适用于iPad的手写笔",
                    "local_sku": "SKU1991FC9",
                    "product_id": "1",
                    "available_amount": 886,
                    "asin_url": "https://www.xxx.com/dp/B09WxxxHZZ17G",
                    "seller_name": null,
                    "pic_url": "https://xxx.com/9664738f11.jpg",
                    "accepted_quantity": 10,
                    "rejected_quantity": 0,
                    "received_quantity": 0
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询VC-Listing列表
下一章节 
查询VC订单详情【PO】