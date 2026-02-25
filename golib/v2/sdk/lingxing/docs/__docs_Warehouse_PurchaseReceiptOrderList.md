查询收货单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/deliveryReceipt/PurchaseReceiptOrder/getOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
date_type	查询时间类型：1 预计到货时间，2 收货时间，3 创建时间，4 更新时间	否	[int]	1
start_date	开始时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-29
end_date	结束时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-29
order_sns	收货单号，多个使用英文逗号分隔	否	[string]	CR240729025,CR240729013
status	状态：10 待收货，40 已完成	否	[int]	10
wid	仓库id，多个使用英文逗号分隔	否	[string]	1,2
order_type	收货类型：1 采购订单，2 委外订单	否	[int]	1
qc_status	质检状态，多个使用英文逗号分隔：0 未质检，1 部分质检，2 完成质检	否	[string]	0,1,2
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认200，上限500	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": "1,578,765",
    "status": 40,
    "order_type": "1",
    "qc_status": "0,1,2",
    "order_sns": "CR240729025,CR240729013",
    "date_type": "3",
    "start_date": "2024-07-29",
    "end_date": "2024-07-29"
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
request_id	请求链路id	是	[string]	C125A617-CDB0-DFDC-5D63-78BB84D27821
response_time	响应时间	是	[string]	2022-10-18 11:23:47
data	响应数据	是	[object]	
data>>total	总数	是	[int]	5404
data>>list	列表	是	[array]	
data>>list>>order_sn	收货单号	是	[string]	CR220927002
data>>list>>status	状态：10 待收货，40 已完成	是	[int]	40
data>>list>>create_time	创建时间	是	[string]	2022-09-27 10:53:59
data>>list>>create_uid	创建人id	是	[int]	1
data>>list>>create_realname	创建人	是	[string]	超级管理员
data>>list>>update_time	更新时间	是	[string]	2022-09-27 10:53:59
data>>list>>receive_time	收货时间	是	[string]	2022-09-27 10:53:59
data>>list>>receive_uid	收货人id	是	[int]	1
data>>list>>receive_realname	收货人	是	[string]	超级管理员
data>>list>>wid	仓库id	是	[int]	1
data>>list>>order_type	收货类型：1 采购订单，2 委外订单	是	[int]	1
data>>list>>qc_type	质检类型：1 仓库质检，2 预检，3 免检	是	[int]	1
data>>list>>business_order_sn	来源单号	是	[string]	PO220927002
data>>list>>supplier_id	供应商id	是	[int]	1
data>>list>>logistics_company	物流商	是	[string]	
data>>list>>logistics_order_no	物流单号	是	[string]	
data>>list>>expect_arrival_time	预计到货时间	是	[string]	
data>>list>>shipping_currency	运费币种	是	[string]	CNY
data>>list>>shipping_cost	运费	是	[string]	0
data>>list>>other_currency	其他费用币种	是	[string]	CNY
data>>list>>other_fee	其他费用	是	[string]	0
data>>list>>opt_uid	采购员id	是	[int]	1
data>>list>>opt_realname	采购员	是	[string]	超级管理员
data>>list>>inbound_order_sns	入库单号	是	[array]	["IB230828018"]
data>>list>>remark	单据备注	是	[string]	
data>>list>>item_list	产品列表	是	[array]	
data>>list>>item_list>>item_id	收货单子项id	是	[string]	19671
data>>list>>item_list>>order_item_id	采购单子项id	是	[string]	21904
data>>list>>item_list>>sku	SKU	是	[string]	1313355
data>>list>>item_list>>product_name	品名	是	[string]	[演示数据]USB壁式充电器2.1A/5V双端口充电器
data>>list>>item_list>>fnsku	FNSKU	是	[string]	
data>>list>>item_list>>seller_id	店铺id	是	[string]	0
data>>list>>item_list>>notice_num_total	通知收货量	是	[number]	164
data>>list>>item_list>>product_receive_num	收货量	是	[number]	164
data>>list>>item_list>>quantity_qc_prepare	待检量	是	[number]	0
data>>list>>item_list>>quantity_qc_already	已检量	是	[number]	164
data>>list>>item_list>>quality_examine_status	质检状态：0 未质检，1 部分质检，2 完成质检	是	[int]	2
data>>list>>item_list>>qc_sn	关联质检单号	是	[array]	
data>>list>>item_list>>remark	备注	是	[string]	备注
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D1DD0410-E19F-550C-EE2A-3BE957331240",
    "response_time": "2024-07-31 15:53:11",
    "data": {
        "total": 2,
        "list": [
            {
                "order_sn": "CR240729025",
                "qc_type": 1,
                "status": 40,
                "create_time": "2024-07-29 15:06:34",
                "create_uid": 14,
                "create_realname": "李立明",
                "update_time": "2024-07-29 15:06:34",
                "receive_time": "2024-07-29 15:06:34",
                "receive_uid": 14,
                "receive_realname": "李立明",
                "wid": 1,
                "order_type": "1",
                "business_order_sn": "PO240729019",
                "inbound_order_sns": [
                    "IB240729031"
                ],
                "supplier_id": 1784,
                "logistics_company": "",
                "logistics_order_no": "",
                "expect_arrival_time": "",
                "shipping_currency": "CNY",
                "shipping_cost": "0.00",
                "other_currency": "CNY",
                "other_fee": "0.00",
                "opt_uid": 14,
                "opt_realname": "李立明",
                "remark": "",
                "item_list": [
                    {
                        "item_id": "114258",
                        "order_item_id": "114404",
                        "sku": "abc001-Y",
                        "product_name": "abc001-Y",
                        "fnsku": "",
                        "seller_id": "0",
                        "notice_num_total": 1200,
                        "product_receive_num": 1200,
                        "quantity_qc_prepare": 0,
                        "quantity_qc_already": 1200,
                        "quality_examine_status": 2,
                        "remark": "",
                        "qc_sn": [
                            "QC240729032"
                        ]
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
查询仓位流水
下一章节 
创建待收货的收货单