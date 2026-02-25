查询平台仓发货单列表

该接口仅支持walmart，现已提供v2接口查询平台仓发货单列表v2，支持多平台下全部平台的平台仓发货单列表数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/warehouse/api/openApi/queryShippingListPage	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	否	[array]	[110000000020008001]
cargo_code	货件单号	否	[string]	test0508001-shipment-2
shipping_list_codes	发货单编号，上限100	否	[array]	["SP202307310010190"]
shipping_list_status	发货单状态：
0 待配货
1 待发货
2 已发货
3 已作废	否	[int]	1
start_time	开始时间【创建时间】，格式：Y-m-d，双闭区间	否	[string]	2023-07-08
end_time	结束时间【创建时间】，格式：Y-m-d，双闭区间	否	[string]	2023-07-08
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15，上限200	否	[int]	15
请求示例
{
    "store_id": [
        "110000000020008001"
    ],
    "cargo_code": "test0508001-shipment-2",
    "shipping_list_codes": [
        "SP202307310010190"
    ],
    "shipping_list_status": 1,
    "start_time": "2023-07-08",
    "end_time": "2023-08-08",
    "offset": 0,
    "length": 15
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
error_details	错误信息	是	[array]	
message	消息提示	是	[string]	success
request_id	请求链路id	是	[string]	110e3649dabd4f14a660de04b1755ee7.1690443265873
response_time	响应时间	是	[string]	2023-07-01 08:11:11
data	响应数据	是	[object]	
data>>total	总数	是	[string]	200
data>>records	详细信息列表	是	[array]	
data>>records>>actual_delivery_time	实际发货时间	是	[string]	2023-08-13
data>>records>>actual_due_ime	实际妥投时间	是	[string]	2023-08-13
data>>records>>arrival_time	到货时间	是	[string]	2023-08-13
data>>records>>creator	创建人	是	[string]	管理员
data>>records>>delivery_time	发货时间	是	[string]	2023-08-13
data>>records>>expected_arrival_time	预计到港时间	是	[string]	2023-08-13
data>>records>>gmt_create	发货单创建时间	是	[string]	2023-07-31 15:33:05
data>>records>>head_fee_type	分摊方式	是	[string]	按SKU数量
data>>records>>logistics_code	物流中心编码	是	[string]	LRY03
data>>records>>logistics_provider_id	物流商id	是	[string]	1
data>>records>>order_logistics_status	订单物流状态	是	[string]	物流状态
data>>records>>principal_user_name	负责人名称	是	[string]	吴
data>>records>>sail_time	开船时间	是	[string]	2023-08-13
data>>records>>shipping_code	发货单号	是	[string]	SP202307310010190
data>>records>>shipping_status	状态：
待配货
待发货
已发货
已作废	是	[string]	已发货
data>>records>>shipping_accessories	平台发货单关联辅料信息	是	[array]	
data>>records>>shipping_accessories>>accessories_price	单位成本	是	[string]	0.00
data>>records>>shipping_accessories>>accessories_total_price	辅料总成本	是	[string]	0.00
data>>records>>shipping_accessories>>aux_name	辅料品名	是	[string]	lry辅料009
data>>records>>shipping_accessories>>aux_num	辅料发货量	是	[string]	2
data>>records>>shipping_accessories>>aux_sku	辅料sku	是	[string]	lryfl009
data>>records>>shipping_accessories>>relation_typeName	费用关联方式	是	[string]	关联商品
data>>records>>shipping_first_lets	平台发货单头程分摊信息	是	[array]	
data>>records>>shipping_first_lets>>aux_stock_price	单位辅料费用	是	[string]	0.00
data>>records>>shipping_first_lets>>delivery_num	发货量	是	[string]	1
data>>records>>shipping_first_lets>>good_name	品名	是	[string]	test050502pm
data>>records>>shipping_first_lets>>gtin	平台标签	是	[string]	006344827368521
data>>records>>shipping_first_lets>>head_stock_price	单位出库头程	是	[string]	0.00
data>>records>>shipping_first_lets>>msku	MSKU	是	[string]	msku
data>>records>>shipping_first_lets>>per_first_let_cost	单位头程费用	是	[string]	0.00
data>>records>>shipping_first_lets>>per_tax	单位税费	是	[string]	0.00
data>>records>>shipping_first_lets>>price	单位出库费用	是	[string]	0.00
data>>records>>shipping_first_lets>>purchase_price	采购单价	是	[string]	0.00
data>>records>>shipping_first_lets>>sku	sku	是	[string]	sku
data>>records>>shipping_first_lets>>value_source	头程费用计算取值类型	是	[string]	实际费用
data>>records>>shipping_first_lets>>wfs_stock_price	单位WFS库存成本	是	[string]	0.00
data>>records>>shipping_goods	平台发货单商品信息	是	[array]	
data>>records>>shipping_goods>>apply_num	申报量	是	[string]	10
data>>records>>shipping_goods>>box_gross_weight	单箱毛重（kg）	是	[string]	10.00
data>>records>>shipping_goods>>box_num	箱数	是	[string]	10
data>>records>>shipping_goods>>box_wet_weight	单箱净重（kg）	是	[string]	10.00
data>>records>>shipping_goods>>cargo_code	货件单号	是	[string]	test0508001-shipment-2
data>>records>>shipping_goods>>cbm	CMB(m³)	是	[string]	0.01
data>>records>>shipping_goods>>cg_product_gross_weight	单品毛重(g)	是	[string]	1.01
data>>records>>shipping_goods>>cg_product_net_weight	单品净重(g)	是	[string]	1.01
data>>records>>shipping_goods>>country_name	国家名称	是	[string]	中国
data>>records>>shipping_goods>>goods_url	商品图片	是	[string]	https://image.disibxxx/03a8b.jpeg
data>>records>>shipping_goods>>msku	msku	是	[string]	msku
data>>records>>shipping_goods>>platform_name	平台名称	是	[string]	沃尔玛
data>>records>>shipping_goods>>product_name	商品名称	是	[string]	商品
data>>records>>shipping_goods>>quantity_in_case	单箱数量	是	[long]	10
data>>records>>shipping_goods>>remark	备注	是	[string]	备注
data>>records>>shipping_goods>>shipments_num	发货量	是	[string]	10
data>>records>>shipping_goods>>sku	sku	是	[string]	ku
data>>records>>shipping_goods>>spec_info	产品箱规总信息	是	[object]	
data>>records>>shipping_goods>>spec_info>>cg_product_net_weight	单品净重	是	[string]	2.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list	箱规信息	是	[array]	
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_box_height	箱子高度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_box_length	箱子长度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_box_pcs	单箱数量	是	[string]	1
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_box_weight	单箱重量	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_box_width	箱子宽度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_package_height	包装规格高度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_package_length	包装规格长度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_package_width	包装规格宽度	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>cg_product_gross_weight	单品毛重	是	[string]	1.00
data>>records>>shipping_goods>>spec_info>>spec_pack_list>>pps_id	规格id	是	[string]	210261409165525663
data>>records>>shipping_goods>>store_name	店铺名称	是	[string]	测试店铺
data>>records>>shipping_goods>>tax_amount	税费	是	[string]	1.00
data>>records>>shipping_goods>>tax_currency	税费币种	是	[string]	CNY
data>>records>>shipping_goods>>total_gw	总毛重（kg）	是	[string]	1.00
data>>records>>shipping_goods>>total_nw	总净重（kg）	是	[string]	1.00
data>>records>>shipping_goods>>volume_weight	体积重（kg）	是	[string]	.00
data>>records>>shipping_goods>>whbCodeInfo	仓位信息	是	[array]	
data>>records>>shipping_goods>>whbCodeInfo>>product_id	产品编号	是	[long]	5280
data>>records>>shipping_goods>>whbCodeInfo>>stock_num	仓位对应库存数量	是	[long]	1988
data>>records>>shipping_goods>>whbCodeInfo>>whb_code	仓位code	是	[string]	ts_valid
data>>records>>shipping_goods>>whbCodeInfo>>whb_name	仓位名称	是	[string]	可用
data>>records>>shipping_goods>>whbCodeInfo>>whb_num	手动填写的仓位数量	是	[long]	10
data>>records>>shipping_goods>>box_length	箱子长度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>box_width	箱子宽度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>box_height	箱子高度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>package_length	包装长度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>package_width	包装宽度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>package_height	包装高度(cm)	是	[string]	1.00
data>>records>>shipping_goods>>box_weight	单箱重量kg）	是	[string]	1.00
data>>records>>shipping_goods>>spec_id	规格id	是	[string]	210261409165525663
data>>records>>shipping_goods>>store_id	店铺id	是	[string]	110000000017008001
data>>records>>shipping_list_files	附件信息	是	[array]	
data>>records>>shipping_list_files>>shipping_list_file_name	文件名称	是	[string]	已发货订单模板20230729.xlsx
data>>records>>shipping_list_files>>shipping_list_file_url	附件url	是	[string]	https://xxxx/A1%E6%9D%BF20230729.xlsx
data>>records>>shipping_logistics	平台发货单物流信息	是	[array]	
data>>records>>shipping_logistics>>expected_other_cost	预估其他费用	是	[string]	1.00
data>>records>>shipping_logistics>>expected_other_cost_remark	预估其他费用备注	是	[string]	备注
data>>records>>shipping_logistics>>expected_other_currency	预估其他费用币种	是	[string]	CNY
data>>records>>shipping_logistics>>expected_transportation_cost	预估物流费用	是	[string]	1.00
data>>records>>shipping_logistics>>expected_transportation_currency	预估物流费用币种	是	[string]	CNY
data>>records>>shipping_logistics>>logistics_number	物流商单号	是	[string]	6674390726051
data>>records>>shipping_logistics>>other_cost	实际其他费用	是	[string]	1.00
data>>records>>shipping_logistics>>other_cost_remark	实际其他费用备注	是	[string]	实际其他费用备注
data>>records>>shipping_logistics>>other_currency	实际其他费用币种	是	[string]	CNY
data>>records>>shipping_logistics>>tracking_number	物流商跟踪号	是	[string]	667439072605
data>>records>>shipping_logistics>>transportation_cost	实际费用	是	[string]	1.00
data>>records>>shipping_logistics>>transportation_currency	实际费用币种	是	[string]	CNY
data>>records>>shipping_remark	备注	是	[string]	备注
data>>records>>warehouse_id	发货仓库id	是	[int]	70
data>>records>>logistics_channel_id	物流渠道id	是	[string]	3
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "1d466c647e624c7cb6c686e332ade41f.117.16909686672353653",
    "response_time": "2023-08-02 17:31:07"
    "data": {
        "total": "29",
        "records": [
            {
                "shipping_code": "SP202302080010030",
                "shipping_status": "待发货",
                "delivery_time": "1970-01-01",
                "warehouse_id": 4244,
                "logistics_provider_id": "145",
                "logistics_channel_id": "3",
                "creator": "吴x",
                "gmt_create": "2023-02-08 10:37:29",
                "actual_delivery_time": null,
                "arrival_time": "2023-02-08",
                "sail_time": null,
                "expected_arrival_time": null,
                "actual_due_ime": null,
                "logistics_code": "LRY03",
                "principal_user_name": null,
                "shipping_remark": "",
                "head_fee_type": "按计费重",
                "order_logistics_status": "",
                "shipping_list_files": [
                    {
                        "shipping_list_file_url": "https://xxx/xxx30731.xlsx",
                        "shipping_list_file_name": "已发货订单模板20230731.xlsx"
                   }
                ],
                "shipping_logistics": [
                    {
                        "logistics_number": "",
                        "tracking_number": "667439072605",
                        "expected_transportation_cost": "0.00",
                        "expected_transportation_currency": "CNY",
                        "expected_other_cost_remark": "",
                        "expected_other_cost": "0.00",
                        "expected_other_currency": "CNY",
                        "other_cost_remark": "",
                        "transportation_currency": "CNY",
                        "transportation_cost": "0.00",
                        "other_cost": "0.00",
                        "other_currency": "CNY"
                    },
                    {
                        "logistics_number": "",
                        "tracking_number": "667439072606",
                        "expected_transportation_cost": "0.00",
                        "expected_transportation_currency": "CNY",
                        "expected_other_cost_remark": "",
                        "expected_other_cost": "0.00",
                        "expected_other_currency": "CNY",
                        "other_cost_remark": "",
                        "transportation_currency": "CNY",
                        "transportation_cost": "0.00",
                        "other_cost": "0.00",
                        "other_currency": "CNY"
                    }
                ],
                "shipping_goods": [
                    {
                        "whbCodeInfo": null,
                        "platform_name": "Walmart",
                        "country_name": "美国",
                        "store_name": "test6自创walmart测试店铺3号",
                        "cargo_code": "lry-shipment-47",
                        "msku": "lry-sku3-47",
                        "product_name": "ww-0825-01",
                        "sku": "ww-0825-01",
                        "goods_url": "",
                        "quantity_in_case": "0",
                        "box_num": "0",
                        "cbm": "0.00",
                        "cg_product_net_weight": "0.00",
                        "cg_product_gross_weight": "0.00",
                        "box_wet_weight": "0.00",
                        "box_gross_weight": "0.00",
                        "volume_weight": "0.00",
                        "apply_num": "10",
                        "shipments_num": "10",
                        "total_nw": "0.00",
                        "total_gw": "0.00",
                        "tax_amount": "0.00",
                        "tax_currency": "USD",
                        "remark": "",
                        "spec_info": {
                            "cg_product_net_weight": "0.00",
                            "spec_pack_list": [
                                {
                                    "pps_id": "210324392619860481",
                                    "spec_title": "默认箱规",
                                    "is_default": "1",
                                    "cg_box_pcs": "0",
                                    "cg_box_length": "0.00",
                                    "cg_box_width": "0.00",
                                    "cg_box_height": "0.00",
                                    "cg_package_length": "0.00",
                                    "cg_package_width": "0.00",
                                    "cg_package_height": "0.00",
                                    "cg_product_gross_weight": "0.00",
                                    "cg_box_weight": "0.00"
                                }
                            ]
                        },
                        "box_length": "0.00",
                        "box_width": "0.00",
                        "box_height": "0.00",
                        "package_length": "0.00",
                        "package_width": "0.00",
                        "package_height": "0.00",
                        "box_weight": "0.00",
                        "spec_id": "210258979648459236",
                        "store_id":"110000000017008001",
                    }
                ],
                "shipping_accessories": null,
                "shipping_first_lets": null
            }
        ]
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "参数检验不通过",
    "error_details": ["length: length 最小值为1"],
    "request_id": "f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946",
    "response_time": "2023-05-22 14:08:59",
    "data": null
}
复制
错误
复制成功
 上一章节
查询报告详情 - Walmart Payment
下一章节 
查询平台仓发货单列表v2