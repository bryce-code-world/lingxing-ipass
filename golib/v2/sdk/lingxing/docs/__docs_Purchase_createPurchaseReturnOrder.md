创建已完成的采购退货单

支持创建采购退货单，状态为“已完成”

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase_return_order/createPurchaseReturnOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
purchase_order_sn	采购单号	是	[string]	
return_method	退货方式，1：退货扣款 2：退货补货	是	[int]	
replenish_method	补货方式，1：源单补货【退货方式为2时必填】	否	[int]	
fee_part_type	分摊方式，0：不分摊 1：按金额 2：按数量	是	[int]	
shipping_currency	退货运费币种，支持CNY、USD，当源单币种为CNY时，运费币种只能为CNY	是	[string]	
shipping_price	退货运费	否	[number]	
other_currency	其他费用币种，支持CNY、USD，当源单币种为CNY时，其他费用币种只能为CNY	是	[string]	
other_fee	其他费用	否	[number]	
return_reason	退货原因	否	[string]	
remark	单据备注	否	[string]	
item_list	退货产品	是	[array]	
item_list>>purchase_order_item_id	采购单子项id	是	[int]	
item_list>>return_good_num	良品退货量，良品退货量和次品退货量不可同时为空	否	[int]	
item_list>>return_bad_num	次品退货量，良品退货量和次品退货量不可同时为空	否	[int]	
item_list>>expect_arrive_time	预计到货时间，退货补货可设置该字段	否	[string]	
item_list>>deduction_amount	退货金额，退货方式为退货扣款时必填，填写值不可大于扣款数量*含税单价	否	[number]	
item_list>>remark	备注	否	[string]	 
请求示例
{
    "purchase_order_sn": "PO220513032",
    "return_method": 1,
    "replenish_method": 0,
    "fee_part_type": 0,
    "shipping_currency": "CNY",
    "shipping_price": 1,
    "other_currency": "CNY",
    "other_fee": 1,
    "return_reason": "产品颜色差异大",
    "remark": "DP0001退货",
    "item_list": [
        {
            "purchase_order_item_id": 19302,
            "return_good_num": 5,
            "return_bad_num": 5,
            "expect_arrive_time": "",
            "deduction_amount": 10,
            "remark": "颜色差异"
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
message	消息提示	是	[string]	success
error_details	错误信息	是	[string]	[]
request_id	请求链路id	是	[string]	C1309FC7-1EC2-9B67-901A-2A0156CC7C25
response_time	响应时间	是	[string]	2022-05-20 11:48:25
data	响应数据	是	[object]	
data>>order_sn	采购退货单号	是	[string]	PR220520003
 上一章节
查询采购退货单列表
下一章节 
作废采购/委外退货单