创建待收货的收货单

支持创建【仓库】>【收货管理】中的收货单，状态为“待收货”

API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/deliveryReceipt/PurchaseReceiptOrder/createReceiptOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
list	收货数据，支持批量	是	[array]	
list>>business_order_sn	业务单号	是	[string]	PO211105076
list>>wid	仓库id	是	[int]	1
list>>order_type	订单类型：1 采购单，2 委外订单	是	[int]	1
list>>expect_arrival_time	期望收货时间	是	[string]	2021-11-12
list>>logistics_company	物流商	是	[string]	物流商
list>>logistics_order_no	物流单号	是	[string]	物流单号
list>>shipping_cost	物流费用	是	[number]	0
list>>other_fee	其他费用	是	[number]	0
list>>remark	备注	是	[string]	备注一下
list>>item_list	收货明细	是	[array]	
list>>item_list>>order_item_id	采购单子项id，查询采购单列表接口对应字段【id】	是	[int]	14062
list>>item_list>>notice_num_total	通知收货量	是	[int]	10
list>>item_list>>remark	备注	是	[string]	我是备注
请求示例
{
    "list": [
        {
            "business_order_sn": "PO230919004",
            "wid": 1,
            "order_type": 1,
            "expect_arrival_time": "2023-09-21",
            "logistics_company": "物流商1",
            "logistics_order_no": "JD-0002",
            "shipping_cost": 10,
            "other_fee": 2,
            "remark": "备注一下",
            "item_list": [
                {
                    "order_item_id": 8823,
                    "notice_num_total": 1,
                    "remark": "我是备注"
                }
            ]
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
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8D680B8B-3C6F-3E16-F055-73DC70EAD977
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[array]	 
 上一章节
查询收货单列表
下一章节 
收货单到货