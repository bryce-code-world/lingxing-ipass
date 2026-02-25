创建待到货的采购单

支持创建采购单，状态为“待到货”

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase/createPurchaseOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	客户仓库id	否	[int]	
sys_wid	系统仓库id【与客户仓库id 二选一必填】	否	[int]	
supplier_id	客户供应商id	否	[int]	
sys_supplier_id	系统供应商id【与客户供应商id 二选一必填】	否	[int]	
custom_order_sn	自定义采购单号【不传此字段则系统自动生成采购单号】	否	[string]	
contact_person	联系人	否	[string]	
contact_number	联系电话	否	[string]	
settlement_method	结算方式：7 现结，8 月结	否	[int]	
settlement_description	结算描述	否	[string]	
payment_method	支付方式：1 网银转账，2 网上支付	否	[int]	
purchase_currency	采购币种	否	[string]	
rate	汇率	否	[number]	
shipping_currency	运费币种	否	[string]	
shipping_price	运费	否	[number]	
other_currency	其它费用币种	否	[string]	
other_fee	其它费用	否	[number]	
fee_part_type	费用分摊方式：0 不分摊，1 按金额，2 按数量	否	[int]	1
is_tax	是否含税：0 否，1 是【当含税为1时，tax_rate为必传字段】	否	[int]	
remark	备注	否	[string]	
opt_uid	采购员uid	是	[int]	
purchaser_id	采购方id，查询采购方列表 接口对应字段【purchaser_id】	是	[int]	1
product_list		是	[array]	
product_list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[int]	
product_list>>seller_id	亚马逊店铺id【建议使用sid替换】	否	[string]	
product_list>>marketplace_id	marketplace_id【建议使用sid替换】	否	[string]	
product_list>>sku	sku	是	[string]	
product_list>>fnsku	fnsku	否	[string]	
product_list>>price	单价	是	[string]	
product_list>>tax_rate	税率，范围为[0,100)【当含税为1时，tax_rate为必传字段】	否	[string]	
product_list>>cases_num	箱数	否	[int]	
product_list>>quantity_per_case	单箱数量	否	[int]	
product_list>>quantity_real	实际采购量	是	[string]	
product_list>>expect_arrive_time	预计到货时间，格式：Y-m-d	否	[string]	
product_list>>remark	备注	否	[string]	
product_list>>plan_sn	采购计划编号	否	[string]	
options	创建选项	否	[object]	
options>>is_auto_fill_store	是否自动填充店铺：【默认0】 0 否，1 是	否	[int]	0
options>>is_auto_fill_fnsku	是否自动填充fnsku：【默认0】0 否，1 是	否	[int]	0
请求示例
{
    "sys_wid": 1,
    "sys_supplier_id": 2,
    "settlement_method": 8,
    "purchase_currency": "CNY",
    "fee_part_type": 0,
    "remark": "测试123",
    "opt_uid": 1,
    "purchaser_id": 1,
    "is_tax": 0,
    "product_list": [
        {
            "sku": "SKU-1",
            "price": "1.00",
            "quantity_real": 50,
            "plan_sn": "PP231211001",
            "expect_arrive_time": "2023-12-21",
            "tax_rate":"10"
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
message	响应消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8DFEE1CA-EC9B-F401-5D41-9F95251D5D50
response_time	响应时间	是	[string]	2020-09-21 15:48:58
data	响应数据	是	[object]	
data>>order_sn	采购单号	是	[string]	PO211112002
data>>custom_order_sn	自定义采购单号	是	[string]	PO211112002
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D41373ED-AF54-CDF1-1658-B8F86286F209",
    "response_time": "2021-11-12 12:37:55",
    "data": {
        "order_sn": "PO211112002",
        "custom_order_sn":"PO211112002"
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "仓库参数错误 [请求码:3CF77C]",
    "request_id": "31667FCC-24E5-0CBE-2FC8-23E6908D9473",
    "response_time": "2021-11-12 12:34:26",
    "data": [],
    "total": 0
}
复制
错误
复制成功
附加说明
客户供应商id，指客户通过领星开放接口创建的供应商，在传值时向领星传入的供应商id。如果供应商是在领星ERP操作创建的，该供应商没有客户供应商id；
系统供应商id，指领星ERP自动生成的供应商id，唯一标识，每个供应商均有自己的系统id；
客户仓库id，指客户通过领星开放接口创建的仓库，在传值时向领星传入的仓库id。如果仓库是在领星ERP操作创建的，该仓库没有客户仓库id；
系统仓库id，指领星ERP自动生成的仓库id，唯一标识，每个仓库均有自己的系统id；
 上一章节
作废采购单
下一章节 
查询采购退货单列表