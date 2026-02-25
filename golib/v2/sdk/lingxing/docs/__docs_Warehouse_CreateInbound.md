创建待发货/待收货/已完成的备货单

支持创建“待发货/待收货/已完成”状态备货单，请求成功后，对应备货单在列表“待发货/待收货/已完成”状态。

若是已对接API的海外仓，提交到待发货状态为未申报状态。

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/createInbound	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
inbound_order_no	客户参考号（唯一单号）	是	[string]
custom_s_wid	自定义仓库id，custom_s_wid和s_wid其中一项必填，都填则优先custom_s_wid	否	[string]	
s_wid	发货仓库，仅限本地仓	是	[int]	
r_wid	收货仓库，仅限海外仓	是	[int]	
logistics_id	物流方式id，查询头程物流渠道列表接口对应字段【id】
（按计费重分摊时，需有传对应物流方式，以获取材积参数用于计算）	是	[int]	
status	订单状态：【默认60】
40 待发货
50 待收货
60 已完成
注：收货仓支持三方海外仓的备货单状态只会到待发货	否	[int]	
estimated_time	预计到货时间	否	[string]	
arrival_time	实际到货时间	否	[string]	
share_id	头程费分摊方式：【默认0】
0 按计费重
1 按实重
2 按体积重
3 按SKU数量
4 自定义
5 按箱子体积
注意：生成待发货状态备货单时，需要通过接口上传备货单装箱信息上传箱子信息；
待收货和已完成的订单不支持【上传备货单装箱信息】，无法按箱子体积分摊	否	[int]	
remark	备注	否	[string]	
file_id	附件id	否	[string]	
overseas_type	下单至第三方【默认2】： 1 否，2 是
注：当收货仓为API海外仓时可填，不填默认为是	否	[int]	
real_delivery_time	实际发货时间	否	[string]	 
logistics_list	物流信息	否	[array]	
logistics_list>>logistics_order_no	物流单号	否	[string]	
logistics_list>>logistics_money	预估物流费用	否	[string]	
logistics_list>>logistics_money_unit	预估物流费用币种	否	[string]	
logistics_list>>other_money	预估其他费用	否	[string]	
logistics_list>>other_money_unit	预估其他费用币种	否	[string]	
logistics_list>>track_order_no	追踪号	否	[string]	
logistics_list>>other_money_remark	预估费用备注	否	[string]	
logistics_list>>real_logistics_money	实际物流费用	否	[number]	
logistics_list>>real_logistics_money_unit	实际物流费用币种	否	[string]	
logistics_list>>real_other_money	实际其他费用	否	[number]	
logistics_list>>real_other_money_unit	实际其他费用币种	否	[string]	
logistics_list>>real_other_money_remark	实际其他费用备注	否	[string]	
logistics_list>>wool_id	物流记录id	否	[int]	
product_list	产品信息	是	[array]	
product_list>>product_id	本地商品id	是	[int]	
product_list>>fnsku	fnsku	否	[string]	
product_list>>tariffs	报关费用	否	[number]	
product_list>>tariffs_currency_unit	报关费用币种	否	[string]	
product_list>>cg_product_gross_weight	单品净重（G）	否	[number]	
product_list>>cg_package_length	包装规格-长（CM）	否	[number]	
product_list>>cg_package_width	包装规格-宽（CM）	否	[number]	
product_list>>cg_package_height	包装规格-高（CM）	否	[number]	
product_list>>stock_num	备货数量，整箱配对需要乘以配对数量	是	[int]	
product_list>>receive_num	收货数量（收货数量可以为0）	是	[int]	
product_list>>oversea_product_code	三方sku编码，查询系统产品与第三方海外仓产品映射列表接口对应字段【oversea_product_code】
【当仓库是有api海外仓必填】	否	[string]	
product_list>>sid	店铺id，有fnsku填 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[string]	 
logistics_list_type	物流信息版本：0或者不传：默认旧版物流信息
1：新版物流信息	是	[int]	0
head_logistics_list	新版头程物流信息（当logistics_list_type 为1时才有意义）	是	[object]	
head_logistics_list>>tax_fee_type	税费分摊方式：
0：产品-计费重
1：产品-实重
2：产品-体积重
3：产品-数量
4：自定义
5：箱子-体积)	是	[int]	1
head_logistics_list>>tracking_list	轨迹信息	是	[list]	
head_logistics_list>>tracking_list>>tracking_no	查询单号	是	[string]	cxdhxxxxxx123
head_logistics_list>>tracking_list>>transport_type	运输类型：
1：快递
2：海运
3：空运
4：其他
	是	[int]	1
head_logistics_list>>tracking_list>>order_type_code	单号类型：【注意：与运输类型联动关系】
1：订舱号
2：提单号
3：箱号
4：其他
5：跟踪单号
6：航班号
当transport_type=1时只能传5
当transport_type=2时只能传1、2、3、4
当transport_type=3时只能传2、6	是	[int]	2
head_logistics_list>>tracking_list>>shippers	承运商【运输类型为海运时才有意义】	是	[string]	coscoxxx
head_logistics_list>>tracking_list>>remark	备注	是	[string]	这是一个备注
head_logistics_list>>estimate_expenses_list	费用明细-预估费用	是	[object]	
head_logistics_list>>estimate_expenses_list>>chargeable_weight	计费重(单位KG)	否	[string]	111.00
head_logistics_list>>estimate_expenses_list>>price	单价	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>logistics_fee	物流费用	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>remark	备注	否	[string]	预估费用备注
head_logistics_list>>estimate_expenses_list>>other_fee_arr	预估费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>fee_type_id	其他费id（20位）	是	[string]	241192037543649792
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list	费用明细-实际费用	是	[object]	
head_logistics_list>>actual_expenses_list>>tax_fee	税费	是	[string]	1
head_logistics_list>>actual_expenses_list>>tax_fee_currency	税费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>chargeable_weight	计费重	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price	单价	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>logistics_fee	物流费用	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>remark	备注	否	[string]	这个一个实际费用的备注
head_logistics_list>>actual_expenses_list>>other_fee_arr	实际费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>actual_expenses_list>>other_fee_arr>>fee_type_id	其他费id(20位)	是	[string]	241192037543649792
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
method_id	运输方式 查询运输方式列表接口对应字段【method_id】	否	[string]	241261034153378816
custom_fields	自定义字段	否	[object]	
请求示例
{
    "inbound_order_no": "20221116161351",
    "s_wid": "1",
    "r_wid": "3835",
    "status": 60,
    "estimated_time": "2022-05-11",
    "arrival_time": "2022-05-13",
    "share_id": "0",
    "remark": "备注33",
    "file_id": "",
    "logistics_id": "35",
    "method_id": "241261034153378816",
    "logistics_list": [
        {
            "logistics_order_no": "YT21316212661013364",
            "logistics_money": "10",
            "logistics_money_unit": "TRY",
            "other_money": "9.6",
            "other_money_unit": "CNY",
            "track_order_no": "H1018660368132401072",
            "other_money_remark": "预估备注",
            "real_logistics_money": "4",
            "real_logistics_money_unit": "CNY",
            "real_other_money": "4",
            "real_other_money_unit": "CNY",
            "real_other_money_remark": "实际备注"
        }
    ],
    "product_list": [
        {
            "product_id": "33664",
            "fnsku": "",
            "oversea_product_code": "LINGXING-YF",
            "tariffs": 1,
            "tariffs_currency_unit": "CNY",
            "stock_num": "2",
            "receive_num": "1",
            "cg_product_gross_weight": 1,
            "cg_package_length": 1,
            "cg_package_width": 2,
            "cg_package_height": 4
        }
    ],
    "logistics_list_type": 0,
    "head_logistics_list": {
        "tax_fee_type": 1,
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
            "chargeable_weight": "111.00",
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
message	信息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	D17B0688-C52D-1BA9-1BCC-50AA00058A3F
response_time	响应时间	是	[string]	2022-08-23 16:56:31
data	响应数据	是	[object]	
data>>overseas_order_no	系统备货单号	是	[string]	
total	数据总量	是	[int]	 
 上一章节
备货单结束到货
下一章节 
查询系统产品与第三方海外仓产品映射列表