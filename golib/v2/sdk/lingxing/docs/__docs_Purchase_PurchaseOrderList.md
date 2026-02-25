查询采购单列表

支持查询采购单列表，对应系统【采购】>【采购单】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/purchaseOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
start_date	开始时间，格式：Y-m-d，双闭区间
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	是	[string]	2024-05-30
end_date	结束时间，格式：Y-m-d，双闭区间
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	是	[string]	2024-08-02
search_field_time	时间搜索维度：
create_time 创建时间【默认值】
expect_arrive_time 预计到货时间
update_time 更新时间	否	[string]	create_time
order_sn	采购单号，上限500	否	[array]	["PO240802012","PO240802011"]
custom_order_sn	自定义采购单号，上限500	否	[array]	["PO240802012","PO240802011"]
purchase_type	采购类型，1：普通采购，2:1688采购	否	[int]	1
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认500，上限500	否	[int]	500
请求示例
{
    "search_field_time": "create_time",
    "start_date": "2024-05-30",
    "end_date": "2024-08-02",
    "order_sn": [
        "PO240802012",
        "PO240802011"
    ],
    "custom_order_sn": [
        "PO240802012",
        "PO240802011"
    ],
    "purchase_type":1,
    "offset": 0,
    "length": 20
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
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2022-05-20 16:57:03
data	响应数据	是	[array]	
data>>order_sn	采购单号	是	[string]	PO201019001
data>>custom_order_sn	自定义单号	是	[string]	
data>>supplier_id	供应商id	是	[int]	291
data>>supplier_name	供应商	是	[string]	深圳xxx有限公司
data>>opt_uid	采购员id	是	[int]	297
data>>principal_uids	单据负责人信息	是	[array]	
data>>principal_uids>>id	单据负责人UID	是	[string]	1221
data>>principal_uids>>name	单据负责人名称	是	[string]	用户A
data>>auditor_realname	审核人姓名	是	[string]	
data>>opt_realname	操作人姓名	是	[string]	测试人员
data>>last_realname	最后操作人姓名	是	[string]	测试人员
data>>create_time	创建时间	是	[string]	2020-10-19 10:43:01
data>>order_time	下单时间	是	[string]	2020-10-19 10:43:01
data>>payment	应付货款（手工）	是	[string]	0.00
data>>auditor_uid	审核人员id	是	[int]	10
data>>auditor_time	审核时间	是	[string]	2020-10-19 10:43:01
data>>last_uid	最后操作人员id	是	[int]	297
data>>last_time	最后操作时间	是	[string]	2020-10-19 10:43:01
data>>reason	作废原因	是	[string]	采购数量有误
data>>is_tax	是否含税：0 否，1 是	是	[int]	
data>>status	采购单状态：
-1 作废
3 待提交
1 待下单 - 已审核
2 待签收(待到货) - 已下单
9 完成
121 (审批流)待审核
122 (审批流)驳回
124 (审批流)作废	是	[int]	1
data>>status_text	状态说明	是	[string]	待审核
data>>pay_status_text	支付状态说明	是	[string]	未申请
data>>status_shipped	到货状态：
1 未到货
2 部分到货
3 全部到货	是	[int]	1
data>>status_shipped_text	到货状态说明	是	[string]	未到货
data>>amount_total	货物总价	是	[number]	200.00
data>>total_price	总金额	是	[number]	200.00
data>>icon	币种符号	是	[string]	￥
data>>pay_status	付款状态：
0 未申请
1 已申请
2 部分付款
3 已付款	是	[int]	1
data>>remark	备注	是	[string]	
data>>other_fee	其他费用	是	[number]	0.00
data>>other_currency	其他费用币种	是	[string]	CNY
data>>fee_part_type	费用分摊方式：
0 不分摊
1 按金额
2 按数量	是	[int]	1
data>>shipping_price	运费	是	[number]	0.00
data>>shipping_currency	运费币种	是	[string]	CNY
data>>purchase_currency	采购币种	是	[string]	CNY
data>>purchase_rate	采购汇率	是	[number]	1.0000
data>>quantity_total	采购总量	是	[number]	200
data>>wid	仓库id	是	[int]	8
data>>ware_house_name	仓库名	是	[string]	测试仓库
data>>ware_house_bak_name	仓库名(备份)	是	[string]	测试仓库
data>>quantity_entry	入库量	是	[int]	10
data>>quantity_real	实际采购量	是	[int]	200
data>>quantity_receive	待到货量	是	[int]	20
data>>update_time	采购单更新时间	是	[string]	2020-10-19 10:43:01
data>>purchaser_id	采购方id	是	[int]	
data>>contact_person	联系人	是	[string]	
data>>contact_number	联系方式	是	[string]	
data>>settlement_method	结算方式：
7 现结
8 月结	是	[int]	
data>>settlement_description	结算描述	是	[string]	
data>>purchase_type	采购类型 1:普通采购；2:1688采购	是	[string]	
data>>purchase_type_text	采购类型文本	是	[string]	
data>>alibaba_order_sn	1688订单号	是	[string]	
data>>sub_status	1688订单状态，1：待1688下单，2：等待买家付款，3：等待买家	是	[string]	
data>>sub_status_text	1688订单状态文本	是	[string]	
data>>custom_fields	自定义字段	是	[array]	
data>>payment_method	支付方式	是	[int]	0
data>>item_list	采购单子项	是	[array]	
data>>item_list>>id	采购单子项id	是	[int]	112
data>>item_list>>wid	仓库id	是	[int]	
data>>item_list>>ware_house_name	仓库名称	是	[string]	
data>>item_list>>relation_purchase_plan	更多采购计划号	是	[array]	
data>>item_list>>plan_sn	采购计划号	是	[string]	PP171225003
data>>item_list>>product_id	本地产品id	是	[int]	11528
data>>item_list>>product_name	品名	是	[string]	铅笔
data>>item_list>>sku	SKU	是	[string]	WYLS003
data>>item_list>>fnsku	FNSKU	是	[string]	ES3001
data>>item_list>>sid	店铺id	是	[string]	6
data>>item_list>>model	型号	是	[string]	
data>>item_list>>price	含税单价	是	[number]	1.0000
data>>item_list>>amount	价税合计	是	[number]	100.00
data>>item_list>>quantity_plan	计划采购量	是	[int]	10
data>>item_list>>quantity_real	实际采购量	是	[int]	100
data>>item_list>>quantity_entry	到货入库量	是	[int]	10
data>>item_list>>quantity_receive	待到货量	是	[int]	2
data>>item_list>>quantity_return	退货数	是	[int]	
data>>item_list>>quantity_exchange	换货量	是	[int]	120
data>>item_list>>quantity_qc	质检量	是	[int]	10
data>>item_list>>quantity_qc_prepare	待质检量	是	[int]	10
data>>item_list>>expect_arrive_time	期待到货时间	是	[string]	2020-10-01
data>>item_list>>remark	备注	是	[string]	
data>>item_list>>cases_num	箱数	是	[int]	
data>>item_list>>quantity_per_case	单箱数量	是	[int]	20
data>>item_list>>is_delete	是否删除：0 否，1 是	是	[int]	1
data>>item_list>>msku	MSKU	是	[array]	
data>>item_list>>attribute	属性	是	[string]	
data>>item_list>>tax_rate	税率	是	[string]	
data>>item_list>>spu	spu	是	[string]	
data>>item_list>>spu_name	款名	是	[string]	
data>>item_list>>custom_fields	自定义字段	是	[array]	
data>>logistics_info	物流信息	是	[array]	
data>>logistics_info>>logistics_company	物流公司	是	[string]	
data>>logistics_info>>logistics_order_no	物流单号	是	[string]	
data>>logistics_info>>pol_id	物流信息记录id	是	[string]	
data>>logistics_info>>purchase_order_id	采购订单唯一id	是	[string]	
data>>logistics_info>>purchase_order_sn	采购订单号（order_sn）	是	[string]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "46CDFE68-C43F-03A8-0A58-093B8AB18ED8",
    "response_time": "2023-10-10 11:38:50",
    "data": [
        {
            "order_sn": "PO230914108",
            "custom_order_sn": "PO230914108",
            "supplier_id": 925,
            "supplier_name": "xxx",
            "opt_uid": 10317908,
            "opt_realname": "账户2",
            "create_time": "2023-09-14 15:34:19",
            "order_time": "2023-09-14 15:34:28",
            "purchase_currency": "CNY",
            "shipping_currency": "CNY",
            "purchase_rate": "1.0000",
            "status_shipped": 3,
            "status_shipped_text": "全部到货",
            "quantity_total": 10,
            "payment": "0.00",
            "pay_status": 0,
            "pay_status_text": "未申请",
            "principal_uids": [
                {
                    "id": 10317908,
                    "name": "账户2"
                }
            ],
            "auditor_uid": 0,
            "auditor_time": "",
            "auditor_realname": "",
            "last_uid": 10317908,
            "last_time": "2023-09-14 17:40:34",
            "last_realname": "账户2",
            "purchaser_id": 129,
            "contact_person": "",
            "contact_number": "",
            "settlement_method": 7,
            "settlement_description": "",
            "payment_method": 0,
            "reason": "",
            "wid": 1,
            "is_tax": 0,
            "status": 9,
            "status_text": "已完成",
            "ware_house_name": "默认仓库测试",
            "ware_house_bak_name": "默认仓库测试",
            "shipping_price": "0.00",
            "amount_total": "660.00",
            "other_fee": "0.00",
            "other_currency": "CNY",
            "fee_part_type": 0,
            "total_price": "660.00",
            "icon": "￥",
            "remark": "",
            "quantity_entry": 10,
            "quantity_real": 10,
            "quantity_receive": 0,
            "update_time": "2023-09-14 17:40:34",
            "logistics_info": [],
            "item_list": [
                {
                    "id": 8727,
                    "plan_sn": "",
                    "product_id": 22500,
                    "product_name": "规格测试-1",
                    "model": "",
                    "wid": 1,
                    "ware_house_name": "默认仓库测试",
                    "sku": "guigeceshi-1",
                    "fnsku": "",
                    "sid": 0,
                    "price": "66.0000",
                    "amount": "660.00",
                    "quantity_plan": 0,
                    "quantity_real": 10,
                    "quantity_per_case": 0,
                    "quantity_return": 0,
                    "quantity_exchange": 0,
                    "cases_num": 0,
                    "msku": [],
                    "attribute": [
                        {
                            "attr_id": 770,
                            "attr_name": "规格22",
                            "attr_value": "10-10"
                        }
                    ],
                    "tax_rate": "0.00",
                    "spu": "guigeceshi",
                    "spu_name": "规格测试",
                    "quantity_entry": 10,
                    "quantity_receive": 0,
                    "quantity_qc": 9,
                    "quantity_qc_prepare": 9,
                    "expect_arrive_time": "",
                    "remark": "",
                    "is_delete": 0
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
创建待采购的采购计划
下一章节 
采购单下单