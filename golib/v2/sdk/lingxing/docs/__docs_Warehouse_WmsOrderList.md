查询销售出库单列表

支持查询ERP中【仓库】>【销售出库单】数据，即自发货订单销售出库单【默认返回一个月内审核出库的数据，超过一个月请加上时间入参】

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/wms/order/wmsOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
page	分页页码，默认1	否	[int]	1
page_size	分页长度，默认20，上限200	否	[int]	20
sid_arr	店铺id	否	[array]	[26]
status_arr	状态：
1 物流下单
2 发货中
3 已发货
4 已删除	否	[array]	[1,2]
logistics_status_arr	物流状态：
1 待导入
2 物流待下单
3 物流下单中
4 下单异常
5 下单完成
6 待海外仓下单
7 海外仓下单中
11 待导入国内物流
41 物流取消中
42 物流取消异常
43 物流取消完成	否	[array]	[1,2]
platform_order_no_arr	平台单号	否	[array]	["test123465021"]
order_number_arr	系统单号	否	[array]	["103130837064323072"]
wo_number_arr	销售出库单号	否	[array]	["WO103132593465409536"]
time_type	时间类型：
创建时间 create_at
出库时间【单据操作】 delivered_at
流水出库时间 stock_delivered_at
变更时间 update_at	否	[string]	create_at
start_date	开始日期，格式：Y-m-d，默认为最近1个月	否	[string]	2021-11-23
end_date	结束日期，格式：Y-m-d，默认为最近1个月	否	[string]	2021-12-20
请求示例
{
    "page": 1,
    "page_size": 20,
    "sid_arr": [
        26
    ],
    "status_arr": [
        1
    ],
    "logistics_status_arr": [
        2
    ],
    "platform_order_no_arr": [
        "test123465021"
    ],
    "wo_number_arr": [
        "WO103132593465409536"
    ],
    "order_number_arr": [
        "103130837064323072"
    ],
    "time_type": "create_at",
    "start_date": "2021-11-23",
    "end_date": "2021-12-20"
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
request_id	请求链路id	是	[string]	82D9A20A-FF87-4BFF-E6D9-22D7A6D3433C
response_time	响应时间	是	[string]	2020-05-18 11:23:47
total	总数	是	[int]	83
data	响应数据	是	[array]	
data>>wo_id	出库单id	是	[int]	4398
data>>wo_number	销售出库单号	是	[string]	WO103132593465409536
data>>sid	店铺id	是	[int]	16
data>>seller_name	店铺名称	是	[string]	8P-SLITE-US
data>>site_text	站点名称	是	[string]	[Amazon].美国
data>>wid	仓库id	是	[int]	26
data>>warehouse_name	仓库名	是	[string]	仓库1
data>>warehouse_type	仓库类型：
1 本地仓库
2 FBA仓
3 第三方海外仓	是	[int]	1
data>>batch_no	批次号	是	[string]	
data>>reference_no	参考号	是	[string]	
data>>waybill_no	运单号	是	[string]	
data>>tracking_no	跟踪号	是	[string]	
data>>picker	拣货人	是	[string]	
data>>platform_name	平台名称	是	[string]	Amazon
data>>platform_order_no	平台单号	是	[array]	["test123465021"]
data>>order_number	系统单号	是	[string]	103130837064323072
data>>order_from	订单来源	是	[string]	手工订单
data>>order_type	订单类型：
1 一单一件
2 多品多件【原一单多件】
3 单品多件	是	[int]	1
data>>order_origin_amount	订单金额	是	[string]	10.00
data>>order_currency_code	订单币种	是	[string]	USD
data>>order_customer_service_notes	客服备注	是	[string]	
data>>order_buyer_notes	买家留言	是	[string]	买家留言1
data>>status	状态：
1 物流下单
2 发货中
3 已发货
4 已删除	是	[int]	4
data>>status_name	状态名称	是	[string]	物流下单
data>>logistics_status	物流下单状态：
1 待导入
2 物流待下单
3 物流下单中
4 下单异常
5 下单完成
6 待海外仓下单
7 海外仓下单中
11 待导入国内物流
41 物流取消中
42 物流取消异常
43 物流取消完成	是	[int]	43
data>>logistics_status_name	物流下单状态名称	是	[string]	待导入
data>>logistics_message	物流下单消息	是	[string]	
data>>logistics_provider_id	物流服务商id	是	[int]	1
data>>logistics_provider_name	物流服务商名称	是	[string]	云途
data>>logistics_type_id	物流方式id	是	[int]	101
data>>logistics_type_name	物流方式名称	是	[string]	Wish邮平邮(DG)
data>>logistics_freight	物流运费	是	[string]	0.00
data>>logistics_freight_currency_code	物流运费币种	是	[string]	
data>>logistics_estimated_freight	预估运费	是	[string]	0.00
data>>logistics_estimated_freight_currency_code	预估运费币种	是	[string]	CNY
data>>is_check	是否验货：0 否，1 是	是	[int]	
data>>is_weigh	是否称重：0 否，1 是	是	[int]	
data>>is_surface_print	面单是否打印：0 否，1 是	是	[int]	
data>>is_order_print	订单是否打印：0 否，1 是	是	[int]	
data>>process_sn	加工单号	是	[string]	
data>>target_country	收货国家	是	[string]	美国
data>>tag_names	标签	是	[array]	
data>>pkg_volume	包裹体积	是	[string]	0.00
data>>pkg_length	包裹尺寸长	是	[string]	0.00
data>>pkg_width	包裹尺寸宽	是	[string]	0.00
data>>pkg_height	包裹尺寸高	是	[string]	0.00
data>>pkg_weight	估算重量	是	[string]	0.000
data>>pkg_real_weight	包裹实重	是	[string]	0.000
data>>pkg_fee_weight	包裹计费重	是	[string]	0.000
data>>pkg_weight_unit	预估重量单位	是	[string]	g
data>>pkg_real_weight_unit	包裹实重单位	是	[string]	g
data>>pkg_fee_weight_unit	包裹计费重单位	是	[string]	g
data>>pkg_size_unit	包裹尺寸单位	是	[string]	cm
data>>recipient_tax_no	收件人税号	是	[string]	
data>>sender_tax_no	发件人税号	是	[string]	
data>>deliverer	发货人	是	[string]	
data>>deliver_deadline	发货时限	是	[string]	
data>>delivered_at	出库时间	是	[string]	
data>>stock_delivered_at	库存流水出库时间	是	[string]	
data>>create_at	创建时间	是	[string]	2021-12-16 20:02:19
data>>update_at	变更时间	是	[string]	2021-12-16 20:02:19
data>>purchase_time	下单时间	是	[string]	2023-08-01 15:09:28
data>>payment_time	付款时间	是	[string]	2023-08-01 15:09:28
data>>surface_print_time	面单打印时间	是	[string]	0000-00-00 00:00:00
data>>order_print_time	订单打印时间	是	[string]	0000-00-00 00:00:00
data>>platform_payment_time	平台结算时间	是	[string]	
data>>package_no	小包号(用于组包)	是	[string]	
data>>package_delivered_data	包裹出库信息	是	[array]	
data>>transfer_logistics_company_code	国内中转物流公司代码	是	[string]	
data>>transfer_logistics_company_id	国内中转物流公司id	是	[string]	
data>>transfer_tracking_no	国内中转跟踪号	是	[string]	
data>>is_lock_storage	是否已锁定库存：0 否，1 是	是	[int]	0
data>>is_advance_delivery	是否预发货：0 否，1 是	是	[int]	1
data>>apportion_status	费用分摊状态：
1 未分摊
2 分摊失败
3 分摊成功	是	[int]	0
data>>apportion_message	费用分摊消息	是	[string]	
data>>remark_attachment	客服备注附件json	是	[string]	[]
data>>consignee	收件人	是	[string]	Rue
data>>consignee_phone	收件人电话	是	[string]	15233330000
data>>consignee_postcode	收件人邮编	是	[string]	40391
data>>consignee_address	收件人地址	是	[string]	3824 Rosewood Court
data>>consignee_full_address	收件地址	是	[string]	美国 KY Winchester 3824 Rosewood Court
data>>surface_file_type	面单文件类型	是	[string]	
data>>surface_file	面单文件	是	[object]	
data>>surface_file>>uri	链接	是	[string]	
data>>surface_file>>ext	文件后缀	是	[string]	
data>>surface_file>>size	文件尺寸	是	[string]	
data>>product_info	商品信息	是	[array]	
data>>product_info>>wod_id	出库单明细id	是	[int]	
data>>product_info>>product_id	商品id	是	[int]	6127
data>>product_info>>sku	SKU	是	[string]	zifahuo
data>>product_info>>count	数量	是	[int]	2
data>>product_info>>bundle_type	捆绑类型：
0 普通商品 - 含组合产品、子产品
1 捆绑产品
2 捆绑产品拆分子产品	是	[int]	0
data>>product_info>>bundle_wod_id	捆绑产品wod_id【只有拆分子产品才有】	是	[int]	
data>>product_info>>product_name	商品名	是	[string]	自发货接口测试
data>>product_info>>seller_sku	MSKU	是	[string]	
data>>product_info>>customization	商品备注	是	[string]	DDAV
data>>product_info>>cn_name	中文申报名	是	[string]	
data>>product_info>>en_name	英文申报名	是	[string]	
data>>product_info>>third_product_name	三方仓品名	是	[string]	
data>>product_info>>third_product_code	三方仓SKU	是	[string]	
data>>product_info>>unit_price	商品单价	是	[string]	
data>>product_info>>currency_code	币种	是	[string]	$
data>>product_info>>apportion_freight	分摊运费 (总计)	是	[string]	0.000000
data>>product_info>>apportion_freight_single	分摊运费 (单个)	是	[string]	0.000000
data>>product_info>>logistics_freight_currency_code	物流运费币种	是	[string]	
data>>product_info>>stock_cost	库存成本 (总计)	是	[string]	6.00
data>>product_info>>stock_sid	库存颗粒度店铺id	是	[string]	0
data>>product_info>>stock_seller_name	库存颗粒度店铺名称	是	[string]	
data>>product_info>>item_unit_price	销售单价	是	[string]	100.00
data>>product_info>>item_total_price	销售总价	是	[string]	100.00
data>>product_info>>real_weight_total	费用分摊-总实重	是	[string]	0.00
data>>product_info>>fee_weight_total	费用分摊-总计费重	是	[string]	0.00
data>>product_info>>volume_weight_total	费用分摊-总体集中	是	[string]	0.00
data>>product_info>>declared_currency_icon	申报比重图标	是	[string]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "42AAD00B-C91D-7E1A-7FF2-EC3169A9BCEF",
    "response_time": "2023-08-03 10:47:20",
    "data": [
        {
            "wo_id": 6879,
            "wo_number": "WO103341008688546816",
            "sid": 104,
            "wid": 25,
            "warehouse_type": 1,
            "order_number": "103340669924180992",
            "batch_no": "",
            "order_from": "补发订单",
            "reference_no": "",
            "deliver_deadline": "",
            "purchase_time": "2023-08-01 15:09:28",
            "payment_time": "2023-08-01 15:09:28",
            "surface_print_time": "0000-00-00 00:00:00",
            "order_print_time": "0000-00-00 00:00:00",
            "waybill_no": "",
            "tracking_no": "",
            "package_no": "",
            "transfer_logistics_company_code": "",
            "transfer_logistics_company_id": "",
            "transfer_tracking_no": "",
            "picker": "",
            "order_type": 2,
            "deliverer": "",
            "status": 4,
            "logistics_status": 5,
            "logistics_message": "",
            "order_customer_service_notes": "",
            "order_buyer_notes": "",
            "is_check": 0,
            "is_weigh": 0,
            "is_surface_print": 0,
            "is_order_print": 0,
            "order_origin_amount": "100.00",
            "order_currency_code": "USD",
            "delivered_at": "",
            "stock_delivered_at": "",
            "process_sn": "",
            "logistics_provider_id": 6,
            "logistics_type_id": 265,
            "surface_file_type": "",
            "is_lock_storage": 0,
            "is_advance_delivery": 1,
            "apportion_status": 0,
            "apportion_message": "",
            "remark_attachment": "[]",
            "platform_payment_time": "",
            "tag_names": [],
            "seller_name": "uboss 1-US",
            "site_text": "[Amazon].美国",
            "product_info": [
                {
                    "wod_id": 6945,
                    "product_id": 22162,
                    "sku": "tianjia",
                    "stock_sid": "0",
                    "stock_seller_name": "",
                    "bundle_type": 0,
                    "bundle_wod_id": 0,
                    "product_name": "添加关联辅料",
                    "count": 1,
                    "seller_sku": "",
                    "customization": "DDAV",
                    "item_unit_price": "100.00",
                    "item_total_price": "100.00",
                    "real_weight_total": "0.00",
                    "fee_weight_total": "0.00",
                    "volume_weight_total": "0.00",
                    "apportion_freight": "0.000000",
                    "apportion_freight_single": "0.000000",
                    "cn_name": "",
                    "en_name": "",
                    "unit_price": "",
                    "currency_code": "$",
                    "declared_currency_icon": "",
                    "logistics_freight_currency_code": "",
                    "stock_cost": "0.000000",
                    "third_product_name": "",
                    "third_product_code": ""
                }
            ],
            "status_name": "已删除",
            "logistics_status_name": "已下单",
            "target_country": "美国",
            "surface_file": {
                "uri": "",
                "ext": "",
                "size": ""
            },
            "platform_name": "Amazon",
            "platform_order_no": [
                "testverppp0117"
            ],
            "package_delivered_data": [],
            "logistics_provider_name": "shipout",
            "logistics_type_name": "未命名101010",
            "warehouse_name": "ruec本地仓",
            "pkg_volume": "0.00",
            "pkg_length": "0.00",
            "pkg_width": "0.00",
            "pkg_height": "0.00",
            "pkg_weight": "0.000",
            "pkg_real_weight": "0.000",
            "pkg_fee_weight": "0.000",
            "pkg_weight_unit": "g",
            "pkg_real_weight_unit": "g",
            "pkg_fee_weight_unit": "g",
            "pkg_size_unit": "cm",
            "logistics_estimated_freight": "0.00",
            "logistics_freight": "0.00",
            "logistics_estimated_freight_currency_code": "CNY",
            "logistics_freight_currency_code": "",
            "consignee_full_address": "美国 KY Winchester  3824 Rosewood Court",
            "consignee": "Rue",
            "consignee_phone": "15233330000",
            "consignee_postcode": "40391",
            "consignee_address": "3824 Rosewood Court",
            "recipient_tax_no": "",
            "sender_tax_no": "",
            "create_at": "2023-08-02 14:07:54"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
备货单分配库存
下一章节 
物流下单 - 编辑运单号/跟踪号