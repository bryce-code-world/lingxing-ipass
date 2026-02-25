创建已完成的采购变更单

支持创建已完成状态的采购变更单，即创建成功后采购单立即变更

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchaseChangeOrder/createPurchaseChangeOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	系统仓库id	是	[int]	1
supplier_id	系统供应商id	是	[int]	1
order_sn	采购单号	是	[string]	PO220714001
contact_person	联系人	否	[string]	张三
contact_number	联系方式	否	[string]	1345467576
settlement_method	结算方式：7 现结，8 月结	是	[int]	7
settlement_description	结算描述	否	[string]	现金结算
shipping_price	运费	否	[number]	1.00
payment_method	支付方式：1 网银转账，2 网上支付	否	[int]	1
purchase_currency	采购币种	是	[string]	CNY
shipping_currency	运费币种	是	[string]	CNY
other_currency	其他费用币种	是	[string]	CNY
rate	汇率	是	[number]	1
other_fee	其他费用	否	[number]	20
fee_part_type	费用分配方式：0 不分配，1 按金额，2 按数量	是	[int]	1
remark	变更单备注	否	[string]	
prepay_percent	预付比例	否	[number]	1
is_tax	是否含税：0 否，1 是	否	[int]	1
opt_uid	采购员U	是	[int]	21
product_list	采购单子项	是	[array]	
product_list>>id	采购单子项id	是	[int]	12
product_list>>quantity_real	实际采购量	是	[string]	100
product_list>>remark	备注	否	[string]	
product_list>>fnsku	FNSKU	否	[string]	FN001
product_list>>sid	店铺 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[int]	10
product_list>>quantity_per_case	单箱数量	否	[int]	1
product_list>>cases_num	箱数	否	[int]	1
product_list>>price	含税单价	是	[number]	100
product_list>>tax_rate	税率	否	[number]	1
product_list>>product_id	产品	是	[int]	123
product_list>>expect_arrive_time	预计到货时间，格式："Y-m-d"	否	[string]	
new_product_list	新增采购单子项	否	[array]	
new_product_list>>quantity_real	实际采购量	是	[int]	100
new_product_list>>remark	备注	否	[string]	
new_product_list>>fnsku	FNSKU	否	[string]	FN002
new_product_list>>sid	店铺 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[int]	2
new_product_list>>quantity_per_case	单箱数量	否	[int]	1
new_product_list>>cases_num	箱数	否	[int]	1
new_product_list>>price	含税单价	是	[number]	200
new_product_list>>tax_rate	税率	否	[number]	2
new_product_list>>product_id	本地产品id	是	[int]	334
请求示例
{
    "wid": 1,
    "supplier_id": 1,
    "order_sn": "PO220714001",
    "contact_person": "张三",
    "contact_number": "1345467576",
    "settlement_method": 7,
    "settlement_description": "现金结算",
    "shipping_price": 1,
    "payment_method": 1,
    "purchase_currency": "CNY",
    "shipping_currency": "CNY",
    "other_currency": "CNY",
    "rate": 1,
    "other_fee": 20,
    "fee_part_type": 1,
    "remark": "",
    "prepay_percent": 1,
    "is_tax": 1,
    "opt_uid": 21,
    "product_list": [
        {
            "id": 12,
            "quantity_real": "100",
            "remark": "",
            "fnsku": "FN001",
            "sid": 10,
            "quantity_per_case": 1,
            "cases_num": 1,
            "price": 100,
            "tax_rate": 1,
            "product_id": 123
        }
    ],
    "new_product_list": [
        {
            "quantity_real": 100,
            "remark": "",
            "fnsku": "FN002",
            "sid": 2,
            "quantity_per_case": 1,
            "cases_num": 1,
            "price": 200,
            "tax_rate": 2,
            "product_id": 334
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C1309FC7-1EC2-9B67-901A-2A0156CC7C25
response_time	响应时间	是	[string]	2022-05-20 11:48:25
data	响应数据	是	[object]	
data>>order_sn	采购变更单号	是	[string]	PR220520003
 上一章节
查询采购变更单列表
下一章节 
查询委外订单列表