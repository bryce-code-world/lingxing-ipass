查询备货单详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overSeaWarehouse/stockOrder/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	OWS241231002
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/overSeaWarehouse/stockOrder/detail?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "overseas_order_no":"OWS241231002"
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
request_id	请求链路id	是	[string]	582616552e104529b0e0d2d216f97c1f.1736410942873
response_time	响应时间	是	[string]	2025-01-09 16:22:25
data	响应数据	是	[object]	
data>>overseas_order_no	备货单号	是	[string]	OWS250918001
data>>inbound_order_no	第三方单号	是	[string]	ASNTEST_38888888001
data>>s_wid	发货仓库id	是	[number]	144
data>>s_wname	发货仓库名称	是	[string]	fq本地仓
data>>r_wid	收货仓库id	是	[number]	862
data>>r_wname	收货仓库名称	是	[string]	败欧洲[4] 主仓
data>>transportation_name	运输方式	是	[string]	海派
data>>logistics_name	物流商名称	是	[string]	加州(432432)
data>>logistics_provider_id	物流商id	是	[number]	1
data>>logistics_provider_name	物流商名称	是	[string]	432432
data>>logistics_way_id	物流渠道id	是	[number]	241401213931552770
data>>logistics_way_name	物流渠道名称	是	[string]	加州
data>>share_text	分摊方式	是	[string]	按计费重
data>>estimated_time	预计到货时间	是	[string]	2025-01-29
data>>real_delivery_time	实际发货时间	否	[string]	2025-01-29
data>>status	备货单状态
枚举类，详情见附加说明	是	[string]	待发货
data>>status_text	备货单状态名称
枚举类，详情见附加说明	是	[string]	待发货
data>>remark	备注	是	[string]	1111备注旧版本
data>>total	总计信息	是	[object]	
data>>total>>product_count	产品总数	是	[number]	2
data>>total>>package_num	装箱数量	是	[string]	3
data>>total>>stock_num	备货数量	是	[string]	3
data>>products	备货产品列表	是	[array]	
data>>products>>product_code	第三方sku	是	[string]	LTHM145AB
data>>products>>twp_name	第三方产品名	是	[string]	头灯
data>>products>>product_id	产品id	是	[number]	3186
data>>products>>sku	系统sku	是	[string]	HH
data>>products>>product_name	产品名称	是	[string]	HH
data>>products>>fnsku	fnsku	是	[string]	
data>>products>>seller_id	店铺id	是	[string]	0
data>>products>>seller_name	店铺名称	是	[string]	-
data>>products>>match_num	配对数量	是	[number]	1
data>>products>>stock_num	备货数量	是	[number]	1
data>>products>>package_num	装箱数量	是	[number]	1
data>>products>>tariffs_currency_unit	预估税费单位	是	[string]	JPY
data>>products>>tariffs	预估税费	是	[number]	2
data>>products>>spec_name	箱规名称	是	[string]	默认箱规
data>>products>>pic_url	图片链接	是	[string]	
data>>products>>country_name	国家名称	是	[string]	
data>>products>>spec_info	箱规信息	是	[object]	
data>>products>>spec_info>>product_id	产品id	是	[number]	3186
data>>products>>spec_info>>cg_product_spec_unit	产品规格单位	是	[string]	cm
data>>products>>spec_info>>cg_product_net_weight_unit	产品重量单位	是	[string]	g
data>>products>>spec_info>>cg_product_length	产品长度	是	[string]	10.00
data>>products>>spec_info>>cg_product_width	产品宽度	是	[string]	10.00
data>>products>>spec_info>>cg_product_height	产品高度	是	[string]	10.00
data>>products>>spec_info>>cg_product_net_weight	产品净重	是	[string]	10000.00
data>>products>>spec_info>>default_spec_pack	默认箱规信息	是	[object]	
data>>products>>spec_info>>default_spec_pack>>cg_box_spec_unit	箱子尺寸单位	是	[string]	cm
data>>products>>spec_info>>default_spec_pack>>cg_box_weight_unit	箱子重量单位	是	[string]	kg
data>>products>>spec_info>>default_spec_pack>>cg_package_spec_unit	包装尺寸单位	是	[string]	cm
data>>products>>spec_info>>default_spec_pack>>cg_product_gross_weight_unit	产品毛重单位	是	[string]	g
data>>products>>spec_info>>default_spec_pack>>spec_title	箱规名称	是	[string]	默认箱规
data>>products>>spec_info>>default_spec_pack>>cg_box_length	箱子尺寸长度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_box_width	箱子尺寸宽度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_box_height	箱子尺寸高度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_box_weight	箱子重量	是	[string]	10.0000
data>>products>>spec_info>>default_spec_pack>>cg_box_pcs	装箱规格数量	是	[number]	1
data>>products>>spec_info>>default_spec_pack>>cg_package_length	包装规格长度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_package_width	包装规格宽度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_package_height	包装规格高度	是	[string]	10.00
data>>products>>spec_info>>default_spec_pack>>cg_product_gross_weight	产品毛重	是	[string]	100.0000
data>>products>>cg_package_length	包装规格长度	是	[number]	10
data>>products>>cg_package_width	包装规格宽度	是	[number]	10
data>>products>>cg_package_height	包装规格高度	是	[number]	10
data>>products>>cg_product_gross_weight	单品毛重	是	[number]	100
data>>products>>remark	备注	是	[string]	HH备注
data>>purchaseInfo	采购信息列表	是	[array]	
data>>purchaseInfo>>product_id	产品ID	是	[number]	3186
data>>purchaseInfo>>product_name	产品名称	是	[string]	HH
data>>purchaseInfo>>sku	SKU	是	[string]	HH
data>>purchaseInfo>>fnsku	FNSKU	是	[string]	
data>>purchaseInfo>>seller_id	店铺ID	是	[string]	0
data>>purchaseInfo>>seller_name	店铺名称	是	[string]	
data>>purchaseInfo>>stock_num	可用总出库量	是	[number]	1
data>>purchaseInfo>>batch_record_list	批次记录列表	是	[array]	
data>>purchaseInfo>>batch_record_list>>batch_no	批次号	是	[string]	2412110005-1
data>>purchaseInfo>>batch_record_list>>storage_good_num	存储良品数量	是	[number]	43
data>>purchaseInfo>>batch_record_list>>good_num	良品数量	是	[number]	1
data>>purchaseInfo>>batch_record_list>>batch_order_sn	批次入库单据号	是	[string]	IB241211005
data>>purchaseInfo>>batch_record_list>>purchase_order_sns	采购订单编号列表	是	[array]	
data>>purchaseInfo>>batch_record_list>>supplier_names	供应商名称列表	是	[array]	
data>>logisticsInfo	物流信息	是	[object]	
data>>logisticsInfo>>head_logistics_info	物流信息	是	[object]	
data>>logisticsInfo>>head_logistics_info>>order_no	物流单号	是	[string]	FL241231002
data>>logisticsInfo>>head_logistics_info>>method_name	运输方式	是	[string]	海派
data>>logisticsInfo>>head_logistics_info>>logistics_provider_id	物流商id	是	[string]	1
data>>logisticsInfo>>head_logistics_info>>logistics_provider_name	物流商名称	是	[string]	432432
data>>logisticsInfo>>head_logistics_info>>logistics_channel_id	物流渠道id	是	[string]	241401213931552768
data>>logisticsInfo>>head_logistics_info>>logistics_channel_name	物流渠道名称	是	[string]	加州
data>>logisticsInfo>>head_logistics_info>>volume_parameter	材积参数	是	[number]	5000
data>>logisticsInfo>>head_logistics_info>>receiving_type	收件方式, 1上门揽收	是	[number]	1
data>>logisticsInfo>>head_logistics_info>>is_customs_or_clearance	一般贸易，0 否， 1 是	是	[number]	0
data>>logisticsInfo>>head_logistics_info>>is_tax_farming	是否包税，0 否， 1 是	是	[number]	1
data>>logisticsInfo>>head_logistics_info>>is_insurance	是否保险，0 否， 1 是	是	[number]	1
data>>logisticsInfo>>head_logistics_info>>is_points_behind	是否分抛计算，0 否， 1 是	是	[number]	0
data>>logisticsInfo>>head_logistics_info>>collection_time	揽收时间	是	[string]	2024-12-30 00:00
data>>logisticsInfo>>head_logistics_info>>eori	EORI	是	[string]	1234
data>>logisticsInfo>>head_logistics_info>>vat_code	VAT税号	是	[string]	1122
data>>logisticsInfo>>head_logistics_info>>logistics_tracking_number	物流商单号	是	[string]	1234567890
data>>logisticsInfo>>head_logistics_info>>custom_order_no	自定义单号	是	[string]	14136347685
data>>logisticsInfo>>head_logistics_info>>head_fee_type_name	分摊方式	是	[string]	产品-计费重
data>>logisticsInfo>>head_logistics_info>>logistics_status_type	物流状态 1 运输中	是	[number]	3
data>>logisticsInfo>>head_logistics_info>>logistics_status	物流状态详情	是	[string]	
data>>logisticsInfo>>head_logistics_info>>order_time	下单时间	是	[string]	2025-01-14
data>>logisticsInfo>>head_logistics_info>>expected_shipment_time	计划发货时间	是	[string]	
data>>logisticsInfo>>head_logistics_info>>shipment_time	发货时间	是	[string]	
data>>logisticsInfo>>head_logistics_info>>transportation_start_time	运输开始时间	是	[string]	2024-12-31
data>>logisticsInfo>>head_logistics_info>>transportation_arrival_time	运输到达时间	是	[string]	2025-01-30
data>>logisticsInfo>>head_logistics_info>>pick_cabinet_time	提柜时间	是	[string]	2025-01-20
data>>logisticsInfo>>head_logistics_info>>send_time	派送时间	是	[string]	2025-01-22
data>>logisticsInfo>>head_logistics_info>>delivery_date	实际妥投时间	是	[string]	2025-01-26
data>>logisticsInfo>>head_logistics_info>>order_remark	下单备注	是	[string]	下单备注
data>>logisticsInfo>>head_logistics_info>>other_remark	其他备注	是	[string]	其他备注
data>>logisticsInfo>>head_logistics_fee_info	新版费用明细列表	是	[array]	
data>>logisticsInfo>>head_logistics_fee_info>>fee_type	费用类型 1 预估 、2 实际	是	[number]	1
data>>logisticsInfo>>head_logistics_fee_info>>chargeable_weight	实重	是	[string]	12.00
data>>logisticsInfo>>head_logistics_fee_info>>volume	体积	是	[string]	0.02
data>>logisticsInfo>>head_logistics_fee_info>>weight	计费重	是	[string]	12.00
data>>logisticsInfo>>head_logistics_fee_info>>price_currency	单价货币	是	[string]	CNY
data>>logisticsInfo>>head_logistics_fee_info>>price	单价	是	[number]	0
data>>logisticsInfo>>head_logistics_fee_info>>logistics_fee_currency	物流费用货币	是	[string]	CNY
data>>logisticsInfo>>head_logistics_fee_info>>logistics_fee	物流费用	是	[string]	20.00
data>>logisticsInfo>>head_logistics_fee_info>>tax_fee_currency	税费货币	是	[string]	JPY
data>>logisticsInfo>>head_logistics_fee_info>>tax_fee	税费	是	[string]	9.00
data>>logisticsInfo>>head_logistics_fee_info>>0_other_currency	其他费用货币	是	[string]	CAD
data>>logisticsInfo>>head_logistics_fee_info>>0_other_amount	其他费用	是	[string]	3.00
data>>logisticsInfo>>head_logistics_fee_info>>total_fee_currency	费用合计货币	是	[string]	CNY
data>>logisticsInfo>>head_logistics_fee_info>>total_fee	费用合计	是	[string]	36.71
data>>logisticsInfo>>head_logistics_fee_info>>remark	备注	是	[string]	6
data>>logisticsInfo>>logistics_list	旧版费用明细列表	是	[array]	
data>>logisticsInfo>>logistics_list>>tracking_number	物流商单号	是	[string]	0987654321
data>>logisticsInfo>>logistics_list>>replace_tracking_number	跟踪号	是	[string]	112233445566
data>>logisticsInfo>>logistics_list>>predicted_transportation_currency	预估物流费用货币	是	[string]	CNY
data>>logisticsInfo>>logistics_list>>predicted_transportation_cost	预估物流费用	是	[string]	20.00
data>>logisticsInfo>>logistics_list>>predicted_other_currency	预估其他费用货币	是	[string]	CAD
data>>logisticsInfo>>logistics_list>>predicted_other_cost	预估其他费用	是	[string]	3.00
data>>logisticsInfo>>logistics_list>>transportation_currency	实际物流费用货币	是	[string]	MXN
data>>logisticsInfo>>logistics_list>>transportation_cost	实际物流费用	是	[string]	4.00
data>>logisticsInfo>>logistics_list>>other_currency	实际其他费用货币	是	[string]	GBP
data>>logisticsInfo>>logistics_list>>other_cost	实际其他费用	是	[string]	5.00
data>>logisticsInfo>>logistics_list>>other_cost_remark	其他费用备注	是	[string]	6
data>>logisticsInfo>>head_logistics_tracking_info	轨迹信息	是	[array]	
data>>logisticsInfo>>head_logistics_tracking_info>>tracking_no	查询单号	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>transport_type_name	运输类型	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>order_type_code_name	单号类型	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>shippers_name	承运商	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>third_status	运输状态	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>last_stayed_at	状态更新时间	是	[String]	
data>>logisticsInfo>>head_logistics_tracking_info>>last_location_name	状态发生地	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info	多物流商列表	否	[array]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>logistics_provider_id	物流商id	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>logistics_provider_name	物流商名	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>fee_type_id	费用类型id	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>fee_type_name	费用类型	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>actual_cost	实际费用	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>actual_currency	实际费用币种	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>predicted_cost	预估费用	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>predicted_currency	预估费用币种	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>gmt_create	创建时间	是	[String]	
data>>logisticsInfo>>head_logistics_fee_provider_info>>gmt_modified	修改时间	是	[String]	
data>>box_data	装箱信息	是	[object]	
data>>box_data>>box_type	装箱方式	是	[number]	1
data>>box_data>>box_content	装箱内容列表	是	[array]	
data>>box_data>>box_content>>pic_url	图片url	是	[string]	
data>>box_data>>box_content>>product_name	品名	是	[string]	HH
data>>box_data>>box_content>>sku	SKU	是	[string]	HH
data>>box_data>>box_content>>third_party_product_name	第三方产品名	是	[string]	头灯
data>>box_data>>box_content>>third_party_product_code	第三方产品编码	是	[string]	LTHM145AB
data>>box_data>>box_content>>fnsku	FNSKU	是	[string]	
data>>box_data>>box_content>>seller_id	店铺id	是	[string]	0
data>>box_data>>box_content>>seller_name	店铺名称	是	[string]	
data>>box_data>>box_content>>boxEnchaseNum	备货量(已装箱)	是	[number]	1
data>>box_data>>box_content>>num	箱数	是	[number]	1
data>>box_data>>box_content>>boxInfo	箱子信息	是	[array]	
data>>box_data>>box_content>>boxInfo>>boxRange	装箱区间	是	[string]	1
data>>box_data>>box_content>>boxInfo>>boxNumber	箱数	是	[int]	
data>>box_data>>box_content>>boxInfo>>cg_box_weight	箱子毛重	是	[number]	10
data>>box_data>>box_content>>boxInfo>>cg_box_length	箱子尺寸长（cm）	是	[number]	10
data>>box_data>>box_content>>boxInfo>>cg_box_width	箱子尺寸宽（cm）	是	[number]	10
data>>box_data>>box_content>>boxInfo>>cg_box_height	箱子尺寸高（cm）	是	[number]	10
data>>box_data>>box_content>>boxInfo>>quantity_in_case	单箱数量	是	[number]	1
data>>box_data>>box_content>>boxInfo>>box_cbm	单箱体积（m³）	是	[number]	0
data>>box_data>>box_content>>boxInfo>>total_box_volume	总体积（m³）	是	[number]	0
data>>box_data>>box_content>>boxInfo>>total_box_weight	总重量（kg）	是	[number]	10
data>>box_data>>box_content>>boxInfo>>total_box_volume_weight	总体积重（kg）	是	[number]	0.2
data>>box_data>>box_remark	装箱备注	是	[string]	
data>>box_data>>total_box_num	总箱数	是	[number]	3
data>>box_data>>total_box_weight	总重量（kg）	是	[number]	12
data>>box_data>>total_box_volume	总体积（m³）	是	[number]	0.02
data>>box_data>>total_box_volume_weight	总体积重（kg）	是	[number]	3.47
data>>custom_fields	自定义字段，返回格式见附加说明	否	[object]		
data>>head_logistics_list	头程物流列表	是	[object]	
data>>head_logistics_list>>actual_expenses_list	实际费用列表	是	[object]	
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr	其他费用数组	是	[array]	
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>fee_type_id	费用类型ID	是	[string]	1
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>name	名称	是	[string]	报关费
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>logistics_provider_id	物流服务商ID	是	[string]	1
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_amount	其他费用金额	是	[string]	100.00
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_currency	其他费用币种	是	[string]	CNY
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "d75cf0dd29454de399b4bcc56ca7f8b8.162.17364169668305063",
    "response_time": "2025-01-09 18:02:48",
    "data": {
        "overseas_order_no": "OWS250918001",
        "inbound_order_no": "ASNTEST_38888888001",
        "s_wid": 144,
        "s_wname": "本地仓",
        "r_wid": 862,
        "r_wname": "主仓",
        "transportation_name": "海派",
        "logistics_name": "2432",
        "logistics_provider_id": 1,
        "logistics_provider_name": "432432",
        "logistics_way_id": 241401213931552768,
        "logistics_way_name": "加州",
        "share_text": "按计费重",
        "estimated_time": "2025-01-29",
        "real_delivery_time": "2025-01-29",
        "status_text": "待发货",
        "remark": "11",
        "total": {
            "product_count": 2,
            "package_num": "3",
            "stock_num": "3"
        },
        "products": [
            {
                "product_code": "LT15AB",
                "twp_name": "头灯",
                "product_id": 3186,
                "sku": "HH",
                "product_name": "HH",
                "fnsku": "",
                "seller_id": "0",
                "seller_name": "-",
                "match_num": 1,
                "stock_num": 1,
                "package_num": 1,
                "tariffs_currency_unit": "JPY",
                "tariffs": 2.0,
                "spec_name": "默认箱规",
                "pic_url": "",
                "country_name": "",
                "spec_info": {
                    "product_id": 3186,
                    "cg_product_spec_unit": "cm",
                    "cg_product_net_weight_unit": "g",
                    "cg_product_length": "10.00",
                    "cg_product_width": "10.00",
                    "cg_product_height": "10.00",
                    "cg_product_net_weight": "10000.00",
                    "default_spec_pack": {
                        "cg_box_spec_unit": "cm",
                        "cg_box_weight_unit": "kg",
                        "cg_package_spec_unit": "cm",
                        "cg_product_gross_weight_unit": "g",
                        "spec_title": "默认箱规",
                        "cg_box_length": "10.00",
                        "cg_box_width": "10.00",
                        "cg_box_height": "10.00",
                        "cg_box_weight": "10.0000",
                        "cg_box_pcs": 1,
                        "cg_package_length": "10.00",
                        "cg_package_width": "10.00",
                        "cg_package_height": "10.00",
                        "cg_product_gross_weight": "100.0000"
                    }
                },
                "cg_package_length": 10.0,
                "cg_package_width": 10.0,
                "cg_package_height": 10.0,
                "cg_product_gross_weight": 100.0,
                "remark": "HH备注"
            },
            {
                "product_code": "0000100000647",
                "twp_name": "",
                "product_id": 3696,
                "sku": "0004",
                "product_name": "wuliu中111",
                "fnsku": "X001JH7Z13",
                "seller_id": "35",
                "seller_name": "",
                "match_num": 1,
                "stock_num": 2,
                "package_num": 2,
                "tariffs_currency_unit": "JPY",
                "tariffs": 7.0,
                "spec_name": "默认箱规",
                "pic_url": "",
                "country_name": "法国",
                "spec_info": {
                    "product_id": 3696,
                    "cg_product_spec_unit": "cm",
                    "cg_product_net_weight_unit": "g",
                    "cg_product_length": "22.00",
                    "cg_product_width": "11.99",
                    "cg_product_height": "30.99",
                    "cg_product_net_weight": "1000.00",
                    "default_spec_pack": {
                        "cg_box_spec_unit": "cm",
                        "cg_box_weight_unit": "kg",
                        "cg_package_spec_unit": "cm",
                        "cg_product_gross_weight_unit": "g",
                        "spec_title": "默认箱规",
                        "cg_box_length": "22.00",
                        "cg_box_width": "11.99",
                        "cg_box_height": "30.99",
                        "cg_box_weight": "1.0000",
                        "cg_box_pcs": 1,
                        "cg_package_length": "33.00",
                        "cg_package_width": "11.99",
                        "cg_package_height": "30.99",
                        "cg_product_gross_weight": "1000.0000"
                    }
                },
                "cg_package_length": 33.0,
                "cg_package_width": 11.99,
                "cg_package_height": 30.99,
                "cg_product_gross_weight": 1000.0,
                "remark": "0004备注"
            }
        ],
        "purchaseInfo": [
            {
                "product_id": 3186,
                "product_name": "HH",
                "sku": "HH",
                "fnsku": "",
                "seller_id": "0",
                "seller_name": "",
                "stock_num": 1,
                "batch_record_list": [
                    {
                        "batch_no": "2412110005-1",
                        "storage_good_num": 43,
                        "good_num": 1,
                        "batch_order_sn": "IB241211005",
                        "purchase_order_sns": [],
                        "supplier_names": []
                    }
                ]
            },
            {
                "product_id": 3696,
                "product_name": "wuliu中111",
                "sku": "0004",
                "fnsku": "X001JH7Z13",
                "seller_id": "35",
                "seller_name": "",
                "stock_num": 2,
                "batch_record_list": [
                    {
                        "batch_no": "2409090002-1",
                        "storage_good_num": 3,
                        "good_num": 1,
                        "batch_order_sn": "IB240909002",
                        "purchase_order_sns": [],
                        "supplier_names": []
                    },
                    {
                        "batch_no": "2412300048-1",
                        "storage_good_num": 19,
                        "good_num": 1,
                        "batch_order_sn": "IB241230019",
                        "purchase_order_sns": [],
                        "supplier_names": []
                    }
                ]
            }
        ],
        "logisticsInfo": {
            "head_logistics_info": {
                "order_no": "FL241231002",
                "method_name": "海派",
                "logistics_provider_id": "1",
                "logistics_provider_name": "432432",
                "logistics_channel_id": "241401213931552768",
                "logistics_channel_name": "加州",
                "volume_parameter": 5000,
                "receiving_type": 1,
                "is_customs_or_clearance": 0,
                "is_tax_farming": 1,
                "is_insurance": 1,
                "is_points_behind": 0,
                "collection_time": "2024-12-30 00:00",
                "eori": "1234",
                "vat_code": "1122",
                "logistics_tracking_number": "1234567890",
                "custom_order_no": "14136347685",
                "head_fee_type_name": "产品-计费重",
                "logistics_status_type": 3,
                "logistics_status": "",
                "order_time": "2025-01-14",
                "expected_shipment_time": "",
                "shipment_time": "",
                "transportation_start_time": "2024-12-31",
                "transportation_arrival_time": "2025-01-30",
                "pick_cabinet_time": "2025-01-20",
                "send_time": "2025-01-22",
                "delivery_date": "2025-01-26",
                "order_remark": "下单备注",
                "other_remark": "其他备注"
            },
            "head_logistics_fee_info": [
                {
                    "fee_type": 1,
                    "chargeable_weight": "12.00",
                    "volume": "0.02",
                    "weight": "12.00",
                    "price_currency": "CNY",
                    "price": 0.0,
                    "logistics_fee_currency": "CNY",
                    "logistics_fee": "20.00",
                    "tax_fee_currency": "JPY",
                    "tax_fee": "9.00",
                    "0_other_currency": "CAD",
                    "0_other_amount": "3.00",
                    "total_fee_currency": "CNY",
                    "total_fee": "36.71",
                    "remark": "6"
                },
                {
                    "fee_type": 2,
                    "chargeable_weight": "0",
                    "volume": "0",
                    "weight": "0",
                    "price_currency": "CNY",
                    "price": 0.0,
                    "logistics_fee_currency": "MXN",
                    "logistics_fee": "4.00",
                    "tax_fee_currency": "",
                    "tax_fee": "",
                    "0_other_currency": "GBP",
                    "0_other_amount": "5.00",
                    "total_fee_currency": "CNY",
                    "total_fee": "44.94",
                    "remark": "6"
                }
            ],
            "logistics_list": [
                {
                    "tracking_number": "0987654321",
                    "replace_tracking_number": "112233445566",
                    "predicted_transportation_currency": "CNY",
                    "predicted_transportation_cost": "20.00",
                    "predicted_other_currency": "CAD",
                    "predicted_other_cost": "3.00",
                    "transportation_currency": "MXN",
                    "transportation_cost": "4.00",
                    "other_currency": "GBP",
                    "other_cost": "5.00",
                    "other_cost_remark": "6"
                }
            ]
        },
        "box_data": {
            "box_type": 1,
            "box_content": [
                {
                    "pic_url": "",
                    "product_name": "HH",
                    "sku": "HH",
                    "third_party_product_name": "头灯",
                    "third_party_product_code": "LTHM145AB",
                    "fnsku": "",
                    "seller_id": "0",
                    "seller_name": "",
                    "boxEnchaseNum": 1,
                    "num": 1,
                    "boxInfo": [
                        {
                            "boxRange": "1",
                            "cg_box_weight": 10.0,
                            "cg_box_length": 10.0,
                            "cg_box_width": 10.0,
                            "cg_box_height": 10.0,
                            "quantity_in_case": 1,
                            "box_cbm": 0.0,
                            "total_box_volume": 0.0,
                            "total_box_weight": 10.0,
                            "total_box_volume_weight": 0.2
                        }
                    ]
                },
                {
                    "pic_url": "",
                    "product_name": "wuliu中111",
                    "sku": "0004",
                    "third_party_product_name": "",
                    "third_party_product_code": "0000100000647",
                    "fnsku": "X001JH7Z13",
                    "seller_id": "35",
                    "seller_name": "非凡-FR",
                    "boxEnchaseNum": 2,
                    "num": 2,
                    "boxInfo": [
                        {
                            "boxRange": "2-3",
                            "cg_box_weight": 1.0,
                            "cg_box_length": 22.0,
                            "cg_box_width": 11.99,
                            "cg_box_height": 30.99,
                            "quantity_in_case": 1,
                            "box_cbm": 0.01,
                            "total_box_volume": 0.02,
                            "total_box_weight": 2.0,
                            "total_box_volume_weight": 3.27
                        }
                    ]
                }
            ],
            "box_remark": "",
            "total_box_num": 3,
            "total_box_weight": 12.0,
            "total_box_volume": 0.02,
            "total_box_volume_weight": 3.47
        },
        "head_logistics_list": {
            "actual_expenses_list": {
                "other_fee_arr": [
                    {
                        "fee_type_id": "1",
                        "name": "报关费",
                        "logistics_provider_id": "1",
                        "other_amount": "100.00",
                        "other_currency": "CNY"
                    }
                ]
            }
        }
    }
}
复制
错误
复制成功
附加说明：

status枚举类如下：
10、待审批
30、待配货
40、待发货
50、待收货
60、已完成

data>>custom_fields返回字段如下：

参数名	说明	类型	示例
data>>custom_fields	自定义字段，键为自定义字段标识id，值为对象	[object]	
data>>custom_fields>>id	字段Id	[string]	
data>>custom_fields>>name	字段名称	[string]	
data>>custom_fields>>val	字段值	[array]	
data>>custom_fields>>val_text	字段值文本描述	[string]	
data>>custom_fields>>character	字段属性	[string]
 上一章节
查询海外仓备货单列表
下一章节 
获取备货单号