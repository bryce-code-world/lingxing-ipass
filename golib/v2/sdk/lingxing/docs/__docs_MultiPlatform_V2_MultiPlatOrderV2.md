查询订单管理订单列表

数据对应多平台管理系统中【订单】>【订单管理】的订单数据，涵盖所有平台自发货订单数据
该接口查询到的订单可修改，可能会与原始销售订单数据有差异，对于亚马逊原始销售订单数据可查询订单列表接口（含 FBM、FBA ）

API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/list	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	是	[int]	0
length	分页长度，上限500	是	[int]	20
date_type	时间类型：
更新时间 update_time
订购时间 global_purchase_time
发货时间 global_delivery_time
付款时间 global_payment_time
平台发货时间 delivery_time（附加说明.4）
当且仅当传入平台单号或平台单名称查询时可不必传	否	[string]	update_time
start_time	开始时间，时间戳格式【单位：秒】，双开区间
当且仅当传入平台单号或平台单名称查询时可不必传，查询时间跨度不能超过31天	否	[int]	1710925191
end_time	结束时间，时间戳格式【单位：秒】，双开区间
当且仅当传入平台单号或平台单名称查询时可不必传，查询时间跨度不能超过31天	否	[int]	1713430791
store_id	店铺id，取值等同于查询多平台店铺信息返回结果的store_id	否	[array]	["110418202566107648"]
platform_code	平台code：
10001 AMAZON
10002 Shopify
10003 eBay
10004 Wish
10005 AliExpress
10006 Shopee
10007 Lazada
10008 Walmart
10009 自定义平台
10010 Wayfair
10011 TikTok
10012 MERCADO
10013 CDISCOUNT
10014 NEWEGG
10015 RAKUTEN
10016 SHOPLINE
10017 TEAPPLIX
10018 SHOPLAZZA
10019 UEESHOP
10020 COUPANG
10021 SHEIN
10022 temu全托管
10024 temu半托管
10025 OTTO
10026 OZON
10027 SHEIN全托管
10028 SHEIN半托管
10029 AliExpress半托管
10030 AliExpress全托管
10033 Qoo10
10034 Mirakl
10025 OTTO
10026 OZON
10027 SHEIN全托管
10028 SHEIN半托管
10029 AliExpress半托管
10030 AliExpress全托管
10033 Qoo10
10034 Mirakl
10035 AMAZON VC
10036 Kaufland
10037 Allegro
10038 Line Shopping
10039 SPS Commerce
	否	[array]	["10001","10002"]
platform_order_nos	平台单号列表 ，元素不超过200个
以下平台不可用，需要用platform_order_names查询：
10003-ebay
10014-newegg
10020-coupang
10002-shopify
10012-美客多
10016-shopline	否	[array]	["119834203044003","120073601764065"]
platform_order_names	特定平台单号列表 ，元素不超过200个
10003-ebay
10014-newegg
10020-coupang
10002-shopify
10012-美客多
10016-shopline，使用该字段查询	否	[array]	["119834203044003","120073601764065"]
order_status	订单状态：
1 同步中
2 已同步
3 待付款
4 待审核
5 待发货
6 已发货
7 已取消/不发货
8 不显示
9 平台发货	否	[int]	3
platform_shipping_status	平台单发货状态
Shopify状态枚举值:
fulfilled：已发货，并且全部发货
null：未发货
partial：部分发货
restocked：已退货	否	[array]	["partial","fulfilled"]
platform_payment_status	平台单支付状态
Shopify状态枚举值:
pending：待支付
authorized：买家信用卡支付，并且已经确认授权，但是卖家并未收款
partially_paid：已完成部分款项支付
paid：已全部支付完成
partially_refunded：已部分退款
refunded：已全部退款
voided：支付被取消，或支付无效，或支付被撤销，通常出现在订单取消之后
	否	[array]	["pending","paid"]
include_delete	是否包含已删除订单
true 包含
false 不包含	否	[boolean]	true
请求示例
{
    "start_time": 1710925191,
    "end_time": 1713430791,
    "date_type": "update_time",
    "offset": 0,
    "length": 20,
    "order_status": 3,
    "store_id": ["110418202566107648"],
    "platform_code": [10001,10002],
    "platform_order_nos": ["119834203044003","120073601764065"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	操作成功
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
data	响应数据	是	[object]	
data>>total	总数	是	[int]	1
data>>list	详细列表	是	[array]	
data>>list>>store_id	店铺id	是	[string]	
data>>list>>wid	仓库id	是	[string]	1
data>>list>>reference_no	平台参考号	是	[string]	1
data>>list>>global_order_no	系统单号
系统本地唯一订单号	是	[string]	
data>>list>>original_global_order_no	补发订单原系统单号	是	[array]	
data>>list>>customer_shipping_list	客选物流列表	否	[array]	[""]
data>>list>>amount_currency	币种	是	[string]	
data>>list>>remark	客服备注	是	[string]	
data>>list>>order_from_name	订单来源	是	[string]	
data>>list>>status	系统订单状态：
1 同步中
2 已同步
3 未付款
4 待审核
5 待发货
6 已发货
7 已取消/不发货
8 不显示
9 平台发货	是	[int]	
data>>list>>split_type	拆分单类型：
1 原始单
2 合并单
3 拆分单	是	[string]	
data>>list>>delivery_type	发货方式：
1 混合方式【中转值，最终会转为2或3】
2 自发货
3 平台发货【指由平台仓库自动完成履约的订单，如Walmart的WFS订单】	是	[int]	2
data>>list>>supplier_id	Wayfair平台对订单指定的发货仓库 ID，店铺后台称为 Supplier ID	是	[string]	
data>>list>>update_time	订单更新时间	是	[string]	
data>>list>>global_purchase_time	订购时间	是	[int]	
data>>list>>global_review_time	审核时间	是	[int]	
data>>list>>global_cancel_time	取消时间	是	[int]	
data>>list>>global_delivery_time	发货时间	是	[int]	
data>>list>>global_distribution_time	配货时间	是	[int]	
data>>list>>global_latest_ship_time	发货时限	是	[string]	
data>>list>>global_payment_time	付款时间	是	[int]	
data>>list>>global_print_time	打单时间	是	[int]	
data>>list>>global_mark_time	标发时间	是	[int]	
data>>list>>item_info	商品信息	是	[array]	
data>>list>>item_info>>id	系统商品唯一id	是	[string]	
data>>list>>item_info>>item_from_name	商品来源
1 线上
2 本地
3 本地(自定义平台导入带MSKU商品，可换货)	是	[string]	
data>>list>>item_info>>customer_shipping_amount	客付运费，币种取amount_currency	是	[string]	
data>>list>>item_info>>customer_tip_amount	Shopify小费，币种取amount_currency	是	[string]	
data>>list>>item_info>>discount_amount	折扣，币种取amount_currency	是	[string]	
data>>list>>item_info>>item_price_amount	商品金额，币种取amount_currency	是	[string]	
data>>list>>item_info>>local_sku	SKU	是	[string]	
data>>list>>item_info>>msku	MSKU	是	[string]	
data>>list>>item_info>>local_product_name	品名	是	[string]	
data>>list>>item_info>>product_no	亚马逊返回asin，多平台返回平台商品id	是	[string]	
data>>list>>item_info>>order_item_no	订单明细单号	是	[string]	
data>>list>>item_info>>platform_order_no	平台单号	是	[string]	
data>>list>>item_info>>platform_status	平台订单商品状态	是	[string]	
data>>list>>item_info>>quantity	数量	是	[int]	
data>>list>>item_info>>remark	商品备注	是	[string]	
data>>list>>item_info>>customized_url	亚马逊定制商品文件下载链接	是	[string]	
data>>list>>item_info>>sales_revenue_amount	销售收益	是	[string]	
data>>list>>item_info>>stockCost	商品出库成本	是	[string]	
data>>list>>item_info>>tax_amount	商品税费，币种取amount_currency	是	[string]	
data>>list>>item_info>>transaction_fee_amount	商品交易费，币种取amount_currency	是	[string]	
data>>list>>item_info>>type	商品类型	否	[string]	
data>>list>>item_info>>unit_price_amount	单价，币种取amount_currency	是	[string]	
data>>list>>item_info>>variant_attr	变体属性	是	[string]	
data>>list>>item_info>>wms_outbound_cost_amount	实际出库成本，币种CNY，取自销售出库单出库成本	是	[string]	67.56
data>>list>>item_info>>stock_cost_amount	库存明细成本，币种CNY，取自订单创建时库存明细计算成本	是	[string]	67.56
data>>list>>item_info>>cg_price_amount	采购成本，币种CNY，取自产品管理固定值	是	[string]	65.00
data>>list>>item_info>>shipping_amount	预估运费，币种CNY	是	[string]	2984.42
data>>list>>item_info>>wms_shipping_price_amount	实际运费，取自销售出库单运费，币种取自data>>list>>logistics_info>>cost_currency_code	是	[string]	￥9.17
data>>list>>item_info>>is_bundled	是否为捆绑产品：
1 是
0 否	是	[int]	0
data>>list>>item_info>>sub_products	子产品	是	[object]	
data>>list>>item_info>>sub_products>>sku	子产品SKU	是	[String]	
data>>list>>item_info>>sub_products>>qty	子产品发货数量	是	[int]	
data>>list>>item_info>>title	商品标题	是	[string]	
data>>list>>item_info>>other_amount	平台其他费用，币种取amount_currency	是	[string]	0.00
data>>list>>item_info>>cod_amount	COD费用，币种取amount_currency	是	[string]	14.00
data>>list>>item_info>>gift_wrap_amount	礼品包装费，币种取amount_currency	是	[string]	300.00
data>>list>>item_info>>platform_tax_amount	销售税，币种取amount_currency	是	[string]	0.00
data>>list>>item_info>>points_granted_amount	积分成本，币种取amount_currency	是	[string]	0.00
data>>list>>item_info>>other_fee	其他费用，用户手动按系统订单导入，有正有负	是	[string]	4.00
data>>list>>item_info>>delivery_time	平台发货时间，当前仅支持Shopify平台，查询条件满足"platform_code":["10002"] && "date_type":"delivery_time"才返回此字段	否	[string]	
data>>list>>item_info>>platform_subsidy_amount	平台补贴(目前仅TikTok、Shopee支持返回）)	否	[string]	
data>>list>>item_info>>data_json	扩展字段(json)	否	[string]	
data>>list>>item_info>>stock_deduct_id	备货店铺ID	否	[string]	
data>>list>>item_info>>stock_deduct_name	备货店铺名称	否	[string]	
data>>list>>logistics_info	物流信息	是	[object]	
data>>list>>logistics_info>>actual_carrier	实际承运人	是	[string]	
data>>list>>logistics_info>>cost_amount	物流实际运费	是	[string]	
data>>list>>logistics_info>>cost_currency_code	物流实际运费币种code	是	[string]	
data>>list>>logistics_info>>logistics_provider_id	物流商id	是	[int]	
data>>list>>logistics_info>>logistics_provider_name	物流商名称	是	[string]	
data>>list>>logistics_info>>logistics_time	物流下单成功时间	是	[int]	
data>>list>>logistics_info>>logistics_type_id	物流方式id	是	[string]	
data>>list>>logistics_info>>logistics_type_name	物流方式名称	是	[string]	
data>>list>>logistics_info>>pkg_fee_weight	实际计费重量	是	[number]	
data>>list>>logistics_info>>pkg_fee_weight_unit	实际计费重单位	是	[string]	
data>>list>>logistics_info>>pkg_height	实际包裹高	是	[number]	
data>>list>>logistics_info>>pkg_length	实际包裹长	是	[number]	
data>>list>>logistics_info>>pkg_size_unit	包裹尺寸单位	是	[string]	
data>>list>>logistics_info>>pkg_width	实际包裹宽	是	[number]	
data>>list>>logistics_info>>pre_cost_amount	预估运费 负	是	[string]	
data>>list>>logistics_info>>pre_fee_weight	预估计费重量	是	[number]	
data>>list>>logistics_info>>pre_fee_weight_unit	预估计费重单位	是	[string]	
data>>list>>logistics_info>>pre_pkg_height	预估包裹高(cm)	是	[number]	
data>>list>>logistics_info>>pre_pkg_length	预估包裹长(cm)	是	[number]	
data>>list>>logistics_info>>pre_pkg_width	预估包裹宽(cm)	是	[number]	
data>>list>>logistics_info>>pre_weight	预估重量(g)	是	[number]	
data>>list>>logistics_info>>status	状态：
0 待物流下单
1 物流下单中
2 成功
3 失败
4 已取消	是	[int]	
data>>list>>logistics_info>>tracking_no	跟踪号	是	[string]	
data>>list>>logistics_info>>waybill_no	物流商返回的物流单号	是	[string]	
data>>list>>logistics_info>>weight	实际重量(g)	是	[number]	
data>>list>>logistics_info>>weight_unit	包裹实际重量单位	是	[string]	
data>>list>>order_tag	标签+处理类型	是	[array]	
data>>list>>order_tag>>tag_name	标签名称	是	[string]	
data>>list>>order_tag>>tag_no	标签编号	是	[string]	
data>>list>>order_tag>>tag_type	标签类型：
1 自定义处理类型
2 自定义订单标签
3 系统处理类型
4 系统订单标签	是	[string]	
data>>list>>exception_order_tag	异常类型标签	是	[array]	
data>>list>>pending_order_tag	待办类型标签	是	[array]	
data>>list>>platform_info	平台单信息	是	[array]	
data>>list>>platform_info>>cancel_time	取消单时间	是	[int]	
data>>list>>platform_info>>delivery_time	平台发货时间	是	[int]	
data>>list>>platform_info>>latest_ship_time	最后发货时间	是	[int]	
data>>list>>platform_info>>order_from	订单来源	是	[string]	
data>>list>>platform_info>>payment_status	平台单支付状态	是	[string]	
data>>list>>platform_info>>payment_time	支付时间	是	[int]	
data>>list>>platform_info>>platform_code	平台CODE	是	[int]	
data>>list>>platform_info>>platform_order_name	平台订单编号
【具体见附加说明】	是	[string]	
data>>list>>platform_info>>platform_order_no	平台订单号	是	[string]	
data>>list>>platform_info>>purchase_time	订购时间	是	[int]	
data>>list>>platform_info>>shipping_status	状态	是	[string]	
data>>list>>platform_info>>status	平台单状态	是	[string]	
data>>list>>platform_info>>store_Country_code	订单国家(二位iso_3166_1)	是	[string]	
data>>list>>transaction_info	交易信息	是	[array]	
data>>list>>transaction_info>>customer_shipping_amount	客付运费，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>customer_tax_amount_show	客付税费，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>customer_tip_amount	小费，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>discount_amount	折扣，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>order_item_amount	商品金额，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>order_total_amount	订单总额，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>outbound_cost_amount	预估出库成本
币种默认CNY	是	[string]	
data>>list>>transaction_info>>pre_cost_amount	预估运费
币种默认CNY	是	[string]	
data>>list>>transaction_info>>transaction_fee_amount	交易费，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>profit_amount	预估毛利润，币种取amount_currency	是	[string]	
data>>list>>transaction_info>>other_amount	平台其他费	是	[string]	￥0.00
data>>list>>transaction_info>>cg_price_amount	采购成本，币种CNY，取自产品管理固定值	是	[string]	-￥77.00
data>>list>>transaction_info>>stock_cost_amount	库存明细成本，币种CNY，取自订单创建时库存明细计算成本	是	[string]	-￥80.07
data>>list>>transaction_info>>wms_outbound_cost_amount	实际出库成本，币种CNY，取自销售出库单出库成本	是	[string]	-￥80.07
data>>list>>transaction_info>>wms_shipping_price_amount	实际运费，币种CNY，取自销售出库单出库运费	是	[string]	-￥11.00
data>>list>>transaction_info>>cod_amount	COD费用，币种取amount_currency	是	[string]	JP¥14.00
data>>list>>transaction_info>>gift_wrap_amount	礼品包装费，币种取amount_currency	是	[string]	JP¥601.00
data>>list>>transaction_info>>platform_tax_amount	销售税，币种取amount_currency	是	[string]	JP¥0.00
data>>list>>transaction_info>>points_granted_amount	积分成本，币种取amount_currency	是	[string]	-JP¥31.00
data>>list>>transaction_info>>other_fee	其他费用 用户手动按系统订单导入，有正有负	是	[string]	JP¥10.00
data>>list>>transaction_info>>platform_subsidy_amount	平台补贴(目前仅TikTok、Shopee支持返回）)	否	[string]
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "response_time": "2023-09-18 10:37:12",
    "request_id": "532e9175-a6d5-42ab-8744-b5d0d8eaea04.1695004631603",
    "data": {
        "total": "1",
        "list": [
            {
                "global_order_no": "103351413752115712",
                "store_id": "110268822038528000",
                "order_from_name": "线上订单",
                "delivery_type": 2,
                "split_type": "1",
                "status": 4,
                "global_purchase_time": 1693467680,
                "global_payment_time": 1693467680,
                "global_review_time": null,
                "global_distribution_time": null,
                "global_print_time": null,
                "global_delivery_time": 0,
                "amount_currency": "USD",
                "remark": "",
                "global_latest_ship_time": "1693497600",
                "global_cancel_time": 0,
                "update_time": "1693496773",
                "order_tag": [
                    {
                        "tag_type": "系统处理类型",
                        "tag_no": "3-8",
                        "tag_name": "商品未配对"
                    }
                ],
                "wid": "0",
                "item_info": [
                    {
                        "id": "8745387",
                        "platform_order_no": "CS483128864",
                        "order_item_no": "NBRC003BLKOS-MVQ",
                        "item_from_name": "线上",
                        "msku": "NBRC003BLKOS-MVQ",
                        "local_sku": "",
                        "product_no": "",
                        "local_product_name": "",
                        "variant_attr": "",
                        "unit_price_amount": "60.00",
                        "item_price_amount": "120.00",
                        "quantity": 2,
                        "remark": "",
                        "platform_status": "",
                        "type": "产品",
                        "stockCost": null,
                        "wms_outbound_cost_amount": "0.00",
                        "stock_cost_amount": "0.00",
                        "cg_price_amount": "0.00",
                        "shipping_amount": "0.00",
                        "wms_shipping_price_amount": "0.00",
                        "customer_shipping_amount": "0.00",
                        "discount_amount": "0.00",
                        "customer_tip_amount": "0.00",
                        "tax_amount": "0.00",
                        "sales_revenue_amount": "24.58",
                        "transaction_fee_amount": "0.49",
                        "other_amount": "0.00",
                        "customized_url": null,
                        "cod_amount": "0.00",
                        "gift_wrap_amount": "0.00",
                        "platform_tax_amount": "0.00",
                        "points_granted_amount": "0.00",
                        "other_fee": "0.00",
                        "platform_subsidy_amount":""
                    }
                ],
                "platform_info": [
                    {
                        "order_from": "线上订单",
                        "platform_order_no": "5501270851822",
                        "platform_order_name": "#1064",
                        "platform_code": "10002",
                        "store_Country_code": "US",
                        "status": "closed",
                        "payment_status": "paid",
                        "shipping_status": "fulfilled",
                        "purchase_time": 1696929350,
                        "payment_time": 1696929349,
                        "latest_ship_time": 0,
                        "cancel_time": 0,
                        "delivery_time": 0
                    }
                ],
                "logistics_info": {
                    "status": 0,
                    "logistics_type_id": "0",
                    "logistics_type_name": null,
                    "logistics_provider_id": "0",
                    "logistics_provider_name": null,
                    "actual_carrier": "",
                    "waybill_no": "",
                    "pre_weight": 0.00,
                    "pre_fee_weight": 0.00,
                    "pre_fee_weight_unit": "g",
                    "pre_pkg_length": 0.00,
                    "pre_pkg_height": 0.00,
                    "pre_pkg_width": 0.00,
                    "weight": 0.00,
                    "pkg_fee_weight": 0.00,
                    "pkg_fee_weight_unit": "",
                    "pkg_length": 0.00,
                    "pkg_width": 0.00,
                    "pkg_height": 0.00,
                    "weight_unit": "",
                    "pkg_size_unit": "",
                    "cost_currency_code": "",
                    "pre_cost_amount": "-￥0.00",
                    "cost_amount": "0",
                    "logistics_time": null,
                    "tracking_no": "",
                    "mark_no": ""
                },
                "transaction_info": [
                    {
                        "order_item_amount": "￥49.16",
                        "customer_tax_amount_show": "￥0.00",
                        "discount_amount": "￥0.00",
                        "customer_shipping_amount": "￥0.00",
                        "customer_tip_amount": "￥0.00",
                        "order_total_amount": "￥49.16",
                        "cg_price_amount": "￥0.00",
                        "stock_cost_amount": "￥0.00",
                        "outbound_cost_amount": "￥0.00",
                        "wms_outbound_cost_amount": "￥0.00",
                        "pre_cost_amount": "-￥0.00",
                        "wms_shipping_price_amount": "￥0.00",
                        "transaction_fee_amount": "￥-0.98",
                        "profit_amount": "￥48.18",
                        "other_amount": "￥0.00",
                        "cod_amount": "￥0.00",
                        "gift_wrap_amount": "￥0.00",
                        "platform_tax_amount": "￥0.00",
                        "points_granted_amount": "￥0.00",
                        "other_fee": "￥0.00",
                        "platform_subsidy_amount":""
                    }
                ],
                "original_global_order_no": null,
                "customer_shipping_list": null,
                "exception_order_tag": [],
                "pending_order_tag_order_tag": []
            }
        ]
    }
}
复制
错误
复制成功
附加说明
data>>list>>platform_info下 platform_order_name、platform_order_no区别：
Shopify平台下，二者返回值不同，核对平台单号使用 platform_order_name，platform_order_no为平台返回订单唯一标识
合并单下可能会存在一个系统单号对应多个平台单号
开始时间和结束时间的最大差距范围不超过30天
当date_type=delivery_time且platform_code=10002（Shopify平台）时，查询将筛选商品层级的data>>list>>item_info>>delivery_time字段，即订单中每个商品（Order Line）在Shopify店铺后台标记的实际发货时间1。
data>>list>>item_info>>delivery_time：表示商品维度的平台发货时间（当前仅支持Shopify平台）。当同一订单包含多个商品时，不同商品可能具有不同的发货时间（即“部分发货”场景），该字段反映对应商品的实际发货时间戳1。
 上一章节
查询多平台店铺信息
下一章节 
查询多平台配对列表