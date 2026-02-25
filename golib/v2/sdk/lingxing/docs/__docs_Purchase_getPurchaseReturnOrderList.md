查询采购退货单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase_return_order/getPurchaseReturnOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search_field_time	时间搜索维度：
create_time 创建时间【默认值】
last_time 更新时间	否	[string]	create_time
start_date	开始时间，格式：Y-m-d，双闭区间
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-08-02
end_date	结束时间，格式：Y-m-d，双闭区间
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-08-02
status	状态：
121 待审批
122 已驳回
124 已作废（审批作废）
10 已处理
20 已作废（单据作废）
5 待退货	否	[array]	[10, 20]
offset	分页偏移量	是	[int]	0
length	分页长度，上限500	是	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "start_date": "2024-08-02",
    "end_date": "2024-08-02",
    "status": [10,20]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2022-05-20 16:57:03
data	响应数据	是	[object]	
data>>total	总数	是	[int]	1000
data>>list	列表	是	[array]	
data>>list>>wid	仓库id	是	[int]	1
data>>list>>order_sn	退货单号	是	[string]	PR220520004
data>>list>>create_uid	创建人id	是	[int]	125
data>>list>>create_realname	创建人名称	是	[string]	JSTAPI
data>>list>>create_time	创建时间	是	[string]	2022-05-20 11:48:23
data>>list>>last_time	更新时间	是	[string]	2022-05-20 11:48:23
data>>list>>buyer_uid	采购员id	是	[int]	175
data>>list>>buyer_realname	采购员名称	是	[string]	API
data>>list>>purchase_order_sn	采购单号	是	[string]	PO220513032
data>>list>>supplier_id	供应商id	是	[int]	12
data>>list>>supplier_name	供应商名称	是	[string]	供应商3
data>>list>>return_method	退货方式：1 退货扣款，2 退货补货	是	[int]	1
data>>list>>replenish_method	补货方式：1 源单补货	是	[int]	
data>>list>>receipt_funds_order_sn	收款单号	是	[string]	
data>>list>>status	状态：
121 待审批
122 已驳回
124 已作废（审批作废）
10 已处理
20 已作废（单据作废）	是	[int]	10
data>>list>>purchase_currency	采购币种	是	[string]	CNY
data>>list>>purchase_currency_icon	采购币种符号	是	[string]	￥
data>>list>>fee_part_type	费用分配方式：
0 不分配
1 按金额
2 按数量	是	[int]	
data>>list>>shipping_currency	运费币种	是	[string]	CNY
data>>list>>shipping_price	退货运费	是	[number]	0.00
data>>list>>other_currency	其他费用币种	是	[string]	CNY
data>>list>>other_fee	其他费用	是	[number]	0.00
data>>list>>return_reason	退货原因	是	[string]	单品0001破损
data>>list>>return_amount_total	退货总金额	是	[string]	
data>>list>>remark	单据备注	是	[string]	破损需要退货
data>>list>>item_list	子项列表	是	[array]	
data>>list>>item_list>>item_id	子项id	是	[int]	313
data>>list>>item_list>>spu_name	款名	是	[string]	SPU NAME
data>>list>>item_list>>spu	SPU	是	[string]	SPU1
data>>list>>item_list>>product_name	品名	是	[string]	单品0001
data>>list>>item_list>>sku	SKU	是	[string]	DP0001
data>>list>>item_list>>fnsku	FNSKU	是	[string]	
data>>list>>item_list>>msku	MSKU	是	[array]	
data>>list>>item_list>>attribute	属性	是	[array]	
data>>list>>item_list>>attribute>>attr_id	属性id	是	[string]	1
data>>list>>item_list>>attribute>>attr_name	属性名	是	[string]	颜色
data>>list>>item_list>>attribute>>attr_value	属性值	是	[string]	白
data>>list>>item_list>>price	含税单价	是	[number]	1.0000
data>>list>>item_list>>quantity_real	采购数量	是	[int]	20
data>>list>>item_list>>return_good_num	良品退货量	是	[int]	5
data>>list>>item_list>>return_bad_num	次品退货量	是	[int]	
data>>list>>item_list>>replenish_num	扣款数量	是	[int]	5
data>>list>>item_list>>deduction_amount	退货金额	是	[number]	5.00
data>>list>>item_list>>expect_arrive_time	预计到货时间	是	[string]	 
data>>list>>item_list>>remark	备注	是	[string]	破损
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "0CDB0A84-AE71-A25C-7432-041C78192C34",
    "response_time": "2024-08-02 17:46:27",
    "data": {
        "total": 2,
        "list": [
            {
                "order_sn": "PR240802002",
                "create_realname": "何承伟",
                "create_uid": 10548111,
                "buyer_realname": "甄新晴",
                "buyer_uid": 10401445,
                "create_time": "2024-08-02 10:56:57",
                "last_time": "2024-08-02 10:56:57",
                "purchase_order_sn": "PO240104006",
                "wid": 44,
                "supplier_name": "深圳市健爱户外运动用品有限公司（已删除）",
                "supplier_id": 1583,
                "return_method": 2,
                "replenish_method": 1,
                "receipt_funds_order_sn": "",
                "status": 10,
                "purchase_currency": "CNY",
                "purchase_currency_icon": "￥",
                "fee_part_type": null,
                "shipping_currency": "CNY",
                "shipping_price": "0.00",
                "other_currency": "CNY",
                "other_fee": "0.00",
                "return_amount_total": "0.00",
                "return_reason": "",
                "remark": "",
                "item_list": [
                    {
                        "item_id": 2414,
                        "spu_name": "",
                        "spu": "",
                        "product_name": "服装1111",
                        "sku": "fuzhuang1111",
                        "fnsku": "B0BBB14W7B",
                        "msku": [],
                        "attribute": [],
                        "seller_id": "109",
                        "price": "100.0000",
                        "quantity_real": 80,
                        "return_good_num": 0,
                        "return_bad_num": 5,
                        "replenish_num": 5,
                        "deduction_amount": "0.00",
                        "expect_arrive_time": "",
                        "remark": ""
                    }
                ]
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
创建待到货的采购单
下一章节 
创建已完成的采购退货单