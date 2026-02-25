生成已发货的发货单

支持推送已发货的FBA发货单到领星ERP，并扣减发货仓库库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/storage/shipment/createSendedOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	自定义仓库id，wid和sys_wid其中一项必填，都填则优先wid	否	[int]	
sys_wid	系统仓库id，wid和sys_wid其中一项必填，都填则优先wid	是	[int]	1
expected_arrival_date	预计到达时间：Y-m-d	否	[string]	2022-05-06
etd_date	开船时间，格式：Y-m-d	否	[string]	2022-05-06
eta_date	预计到港时间，格式：Y-m-d	否	[string]	2022-05-06
delivery_date	实际妥投时间，格式：Y-m-d	否	[string]	2022-05-06
actual_shipment_time	实际发货时间，格式：Y-m-d	否	[string]	2022-05-06
head_fee_type	头程费分配方式：【默认0】
0 按计费重
1 按实重
2 按体积重
3 按SKU数量
4 自定义
5 按箱子体积	否	[int]	0
tax_fee_type	实际税费分配方式：【默认0】
0 产品-计费重
1 产品-实重
2 产品-体积重
3 产品-数量
4 自定义
5 箱子-体积	否	[int]	0
is_points_behind	是否分抛计算：0 否，1 是，头程分摊方式为按计费重时用	否	[int]	
points_behind_coeffient	分抛系数：0~100，分抛计算选是时必填	否	[int]	
logistics_channel_id	物流渠道id：按计费重分摊时必填，以获取材积参数用于计算
查询头程物流渠道列表接口对应字段【id】	否	[int]	
is_related	组合商品扣减库存时是否自动拆分成单品进行扣减：
0 否
1 是【会拆分组合商品】	否	[int]	0
request_flag	自定义请求标识，本次请求超时后可根据此标识查询此次请求的结果，由请求方保持标识唯一性	否	[string]	
ship_mode	发货方式：1-默认，2-工厂直发	否	[int]	
hand_pick_purchase	工厂直发时手动选择出库批次：1-否，2-是	否	[int]	
remark	备注	否	[string]	
list		是	[array]	
list>>seller_id	亚马逊店铺id ,对应查询亚马逊店铺列表接口对应字段【seller_id】	是	[string]	
list>>warehouse_seller_id	仓库店铺id	否	[int]	
list>>marketplace_id	市场id	是	[string]	
list>>shipment_id	货件单号	是	[string]	
list>>fulfillment_network_sku	货件fnsku	是	[string]	
list>>fnsku	本地发货的fnsku，默认空	否	[string]	
list>>num	发货数量	是	[int]	
list>>box_num	箱数	否	[int]	
list>>sku	sku	是	[string]	
list>>tax_currency	报关税费币种	否	[string]	
list>>tax_amount	税费值	否	[string]	
list>>cg_product_gross_weight	采购：单品净重（G）；不传值则自动取【商品管理】模块维护的SKU“单品毛重”；按计费重分摊时需用到该字段	否	[string]	
list>>cg_package_length	采购：包装规格-长（CM）；不传值则自动取【商品管理】模块维护的SKU“包装规格-长”；按计费重分摊时需用到该字段	否	[string]	
list>>cg_package_width	采购：包装规格-宽（CM）；不传值则自动取【商品管理】模块维护的SKU“包装规格-宽”；按计费重分摊时需用到该字段	否	[string]	
list>>cg_package_height	采购：包装规格-高（CM）；不传值则自动取【商品管理】模块维护的SKU“包装规格-高”；按计费重分摊时需用到该字段	否	[string]	
list>>cg_box_weight	采购：单箱重量（KG）；不传值则自动取【商品管理】模块维护的SKU“单箱重量”；按计费重分摊时需用到该字段	否	[string]	
list>>transport_cost	自定义头程费，head_fee_type为4才生效	否	[string]	
list>>purchase_items	工厂直发手动选择出库批次	否	[array]	
list>>purchase_items>>relation_num	关联量	是	[int]	
list>>purchase_items>>purchase_sn	系统采购单号	否	[String]	
list>>purchase_items>>custom_purchase_sn	自定义采购单号(无系统采购单号时必填)	否	[String]	
box_type	装箱类型：
SINGLE 每箱只允许一款SKU
MULTIPLE 每箱允许多款SKU	否	[string]	
box_remark	装箱备注	否	[string]	
box_list	箱规列表，每个子项代表一个箱规，在装箱类型为MULTIPLE时必填	否	[array]	
box_list>>box_num	箱子数	是	[string]	3
box_list>>cg_box_length	箱子长	是	[number]	20
box_list>>cg_box_width	箱子宽	是	[number]	30
box_list>>cg_box_height	箱子高	是	[number]	40
box_list>>cg_box_weight	箱子重	是	[number]	1.33
box_list>>box_skus	箱子内包含的SKU信息，需与list里的数据保持一致	是	[array]	
box_list>>box_skus>>seller_id	亚马逊店铺id ,对应查询亚马逊店铺列表接口对应字段【seller_id】	是	[string]	A1MQMW3JWPNCBX
box_list>>box_skus>>marketplace_id	市场id	是	[string]	xxxxxxxxxxxxxxxx
box_list>>box_skus>>shipment_id	货件单号	是	[string]	FBA16XNMN3N4
box_list>>box_skus>>fulfillment_network_sku	货件fnsku	是	[string]	B09MT3989Q
box_list>>box_skus>>quantity_in_case	单箱数量	是	[int]	32
box_list>>box_skus>>sku	SKU	是	[string]	0530-1
logistics_list_type	物流信息版本：
0 旧版
1 新版	否	[int]	
head_logistics_list	新版头程物流信息	是	[object]	
head_logistics_list>>tracking_list	轨迹信息数组	否	[array]	
head_logistics_list>>tracking_list>>tracking_no	查询单号	否	[string]	cxdhxxxxxx123
head_logistics_list>>tracking_list>>transport_type	运输类型：
1 快递
2 海运
3 空运
4 其他	否	[int]	1
head_logistics_list>>tracking_list>>order_type_code	单号类型：【注意：与运输类型联动关系】
1 订舱号
2 提单号
3 箱号
4 其他
5 跟踪单号
6航班号
当transport_type=1时只能传5
当transport_type=2时只能传1、2、3、4
当transport_type=3时只能传2、6
当transport_type=4只能传空	否	[int]	5
head_logistics_list>>tracking_list>>shippers	承运商，运输类型为海运时才有意义：
获取发货单头程物流-承运商信息接口获取	否	[string]	coscoxxx
head_logistics_list>>tracking_list>>remark	备注	否	[string]	这是一个备注
head_logistics_list>>estimate_expenses_list	费用明细-预估费用	否	[object]	
head_logistics_list>>estimate_expenses_list>>chargeable_weight	计费重(单位KG)	否	[string]	111.00
head_logistics_list>>estimate_expenses_list>>price	单价	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>logistics_fee	物流费用	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>remark	备注	否	[string]	这个一个预估费用的备注
head_logistics_list>>estimate_expenses_list>>other_fee_arr	预估费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>fee_type_id	其他费id	是	[string]	241192037543649792
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list	费用明细-实际费用	否	[object]	
head_logistics_list>>actual_expenses_list>>tax_fee	税费	是	[string]	1
head_logistics_list>>actual_expenses_list>>tax_fee_currency	税费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>chargeable_weight	计费重【废弃字段】	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>weight	实重（单位：KG）	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>volume	体积（单位：m³）	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price	单价	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>logistics_fee	物流费用	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>remark	备注	否	[string]	这个一个实际费用的备注
head_logistics_list>>actual_expenses_list>>other_fee_arr	实际费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>actual_expenses_list>>other_fee_arr>>fee_type_id	其他费id	是	[string]	241192037543649792
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
logistics_list	旧版物流信息，即将下线	否	[array]	
logistics_list>>tracking_number	物流商单号	否	[string]	
logistics_list>>replace_tracking_number	跟踪号	否	[string]	
logistics_list>>transportation_cost	实际物流费用	否	[number]	
logistics_list>>transportation_currency	实际物流费用币种	否	[string]	
logistics_list>>other_cost	实际其他费用	否	[number]	
logistics_list>>other_currency	实际其他费用币种	否	[string]	
logistics_list>>other_cost_remark	其他费用备注	否	[string]	
logistics_list>>predicted_transportation_cost	预估物流费用	否	[number]	
logistics_list>>predicted_transportation_currency	预估物流费用币种	否	[string]	
logistics_list>>predicted_other_cost	预估其他费用	否	[number]	
logistics_list>>predicted_other_currency	预估其他费用币种	否	[string]
请求示例
{
    "sys_wid": 1,
    "expected_arrival_date": "2022-05-06",
    "etd_date": "2022-05-06",
    "eta_date": "2022-05-06",
    "delivery_date": "2022-05-06",
    "actual_shipment_time": "2022-05-06",
    "head_fee_type": 1,
    "tax_fee_type": 0,
    "is_points_behind": 1,
    "points_behind_coeffient": 1,
    "logistics_channel_id": 1,
    "remark": "发货单备注",
    "request_flag": "1123123aaaaaa",
    "logistics_list": [{
        "tracking_number": "aaaaaa",
        "replace_tracking_number": "bbbbbb",
        "transportation_cost": 1,
        "transportation_currency": "CNY",
        "other_cost": 1,
        "other_currency": "1.00",
        "other_cost_remark": "其他费备注",
        "predicted_transportation_cost": 1,
        "predicted_transportation_currency": "CNY",
        "predicted_other_cost": 1,
        "predicted_other_currency": "CNY"
    }],
    "list": [{
        "seller_id": "",
        "warehouse_seller_id": 0,
        "marketplace_id": "",
        "shipment_id": "",
        "fulfillment_network_sku": "",
        "fnsku": "",
        "num": 0,
        "box_num": 0,
        "sku": "",
        "tax_currency": "",
        "tax_amount": "",
        "cg_product_gross_weight": "",
        "cg_package_length": "",
        "cg_package_width": "",
        "cg_package_height": "",
        "cg_box_weight": "",
        "transport_cost": ""
    }],
    "is_related": "",
    "box_type": "",
    "box_remark": "",
    "box_list": [{
        "box_num": "3",
        "cg_box_length": 20,
        "cg_box_width": 30,
        "cg_box_height": 40,
        "cg_box_weight": 1.33,
        "box_skus": [{
            "seller_id": "A1MQMW3JWPNCBX",
            "marketplace_id": "xxxxxxxxxxxxxxxx",
            "shipment_id": "FBA16XNMN3N4",
            "fulfillment_network_sku": "B09MT3989Q",
            "quantity_in_case": 32,
            "sku": "0530-1"
        }]
    }],
    "logistics_list_type": 0,
    "head_logistics_list": {
        "tracking_list": [{
            "tracking_no": "cxdhxxxxxx123",
            "transport_type": "1",
            "order_type_code": "5",
            "shippers": "coscoxxx",
            "remark": "备注"
        }],
        "estimate_expenses_list": {
            "chargeable_weight": "111.00",
            "price": "1.00",
            "price_currency": "CNY",
            "logistics_fee": "1.00",
            "logistics_fee_currency": "CNY",
            "remark": "预估费用备注",
            "other_fee_arr": [{
                "fee_type_id": "241192037543649792",
                "other_amount": "1.00",
                "other_currency": "CNY"
            }]
        },
        "actual_expenses_list": {
            "tax_fee":"1",
            "tax_fee_currency":"CNY",
            "weight": "111.00",
            "volume": "111.00",
            "price": "111.00",
            "price_currency": "CNY",
            "logistics_fee": "111.00",
            "logistics_fee_currency": "CNY",
            "remark": "实际费用备注",
            "other_fee_arr": [{
                "fee_type_id": "241192037543649792",
                "other_amount": "111.00",
                "other_currency": "CNY"
            }]
        }
    }
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[string]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
data	响应数据	是	[object]	
data>>order_sn	发货单号	是	[string]	 
 上一章节
生成待发货的发货单
FBA
下一章节 
发货单创建接口结果查询
FBA