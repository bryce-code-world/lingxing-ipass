更新备货单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/updateInbound	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	海外仓备货单号	是	[string]	
logistics_id	物流方式id【按计费重分摊时，需传对应物流方式，以获取材积参数用于计算】	否	[int]	
product_list	产品信息	否	[array]	
product_list>>product_id	本地商品id	否	[int]	
product_list>>fnsku	fnsku	否	[string]	
product_list>>tariffs	报关费用	否	[number]	
product_list>>tariffs_currency_unit	报关费用币种	否	[string]	
product_list>>cg_product_gross_weight	单品净重（G）	否	[number]	
product_list>>cg_package_length	包装规格-长（CM）	否	[number]	
product_list>>cg_package_width	包装规格-宽（CM）	否	[number]	
product_list>>cg_package_height	包装规格-高（CM）	否	[number]	
product_list>>stock_num	备货数量,整箱配对需要乘以配对数量	否	[int]	
product_list>>sid	店铺id，有fnsku填	否	[string]	
product_list>>product_code	三方产品编码	否	[string]	
product_list>>fba_cost	单位头程费用	否	[number]	
product_list>>fba_cost_currency_unit	单位头程费用币种单位	否	[string]	
product_list>>remark	商品备注	否	[string]	
estimated_time	预计到货时间	否	[string]	
arrival_time	实际到货时间	否	[string]	
share_id	头程费分配方式：
0 按计费重【默认值】
1 按实重
2 按体积重
3 按SKU数量
4自定义	否	[int]	
remark	备注	否	[string]	
file_id	附件id	否	[string]	
overseas_type	下单至第三方【当收货仓为API海外仓时可填，不填默认为是】：1 否，2 是【默认】	否	[int]	
real_delivery_time	实际发货时间，格式：Y-m-d H:i:s	否	[string]	2022-6-24 12:00:25
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
logistics_list	旧版物流信息，即将下线	否	[array]	
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
logistics_list>>operation_type	物流费用操作类型:=新增2=修改3=删除	否	[int]
请求示例
{
    "overseas_order_no": "OWS220624005",
    "logistics_id": 35,
    "share_id": 1,
    "overseas_type": 1,
    "arrival_time": "2022-06-24 12:10:25",
    "estimated_time": "2022-06-25 12:10:25",
    "real_delivery_time": "2022-6-24 12:00:25",
    "remark": "999",
    "file_id": "",
    "logistics_list": [
        {
            "wool_id": 2949,
            "operation_type": 2,
            "logistics_money": "77.00",
            "logistics_money_unit": "CNY",
            "other_money": "5",
            "other_money_unit": "CNY",
            "other_money_remark": "",
            "real_logistics_money": "300.00",
            "real_logistics_money_unit": "CNY",
            "real_other_money": "10.00",
            "real_other_money_unit": "CNY",
            "real_other_money_remark": "",
            "logistics_order_no": "99",
            "track_order_no": "2333"
        }
    ],
    "product_list": [
        {
            "product_id": 2117299,
            "product_code": "TESTSKU1",
            "cg_package_height": "50.00",
            "cg_package_length": "20.00",
            "cg_package_width": "10.02",
            "cg_product_gross_weight": "20.00",
            "fba_cost": "9",
            "fba_cost_currency_unit": "CNY",
            "tariffs": "8",
            "tariffs_currency_unit": "CNY",
            "fnsku": "",
            "sid": 0,
            "stock_num": 7,
            "remark": "商品备注999999999999999"
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
data	响应数据	是	[array]	 
 上一章节
删除备货单
下一章节 
更新备货单物流信息