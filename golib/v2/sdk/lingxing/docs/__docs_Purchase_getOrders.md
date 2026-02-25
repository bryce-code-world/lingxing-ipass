查询委外订单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchaseOutsourceOrder/getOrders	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search_field_time	日期搜索类型
create_time:创建日期
expect_arrive_time:结束日期	否	[string]	create_time
start_date	开始日期（闭区间）	否	[string]	2024-08-02
end_date	结束日期（闭区间）	否	[string]	2024-08-02
offset	分页偏移量	是	[int]	0
length	分页长度，上限500	是	[int]	500
请求示例
{
    "offset": 0,
    "length": 500,
    "search_field_time": "create_time",
    "start_date": "2024-08-02",
    "end_date": "2024-08-02"
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
request_id	请求链路id	是	[string]	C1309FC7-1EC2-9B67-901A-2A0156CC7C25
response_time	响应时间	是	[string]	2022-05-20 11:48:25
data	响应数据	是	[object]	
data>>list	数据列表	是	[array]	
data>>list>>order_sn	单号	是	[string]	POP220518006
data>>list>>warehouse_name	仓库名	是	[string]	仓库1
data>>list>>outsource_warehouse_name	加工仓库名	是	[string]	刘东的本地仓
data>>list>>supplier_name	加工商名	是	[string]	供应商关联测试-刘东
data>>list>>create_time	创建时间	是	[string]	2022-05-18 17:31:08
data>>list>>status_text	单据状态	是	[string]	待提交
data>>list>>create_realname	创建人	是	[string]	刘东
data>>list>>ptp_sn	加工计划单号	是	[string]	
data>>list>>product_name	品名	是	[string]	liudong_fnsku_sid
data>>list>>sku	sku	是	[string]	liudong_fnsku_sid
data>>list>>fnsku	fnsku	是	[string]	
data>>list>>outsource_quantity	委外数量	是	[number]	1
data>>list>>receive_quantity	已收货量	是	[number]	
data>>list>>expect_arrive_time	预计到货时间	是	[string]	
data>>list>>msku	MSKU数组	是	[array]	
data>>list>>plan_sn	采购计划单号数组	是	[array]	["PP220415025"]
data>>list>>seller_name	店铺	是	[string]	
data>>list>>item	子项	是	[array]	
data>>list>>item>>purchase_order_sn	采购单号	是	[string]	PO220518023
data>>list>>item>>product_name	品名	是	[string]	liudong1
data>>list>>item>>sku	sku	是	[string]	liudong1
data>>list>>item>>quantity_require	需求量	是	[number]	1
data>>list>>item>>quantity_real	采购量	是	[number]	1
data>>list>>item>>quantity_entry	到货量	是	[number]	
data>>list>>item>>create_time	创建时间	是	[string]	2022-05-18 17:31:07
data>>list>>item>>create_realname	创建人	是	[string]	刘东
data>>list>>item>>warehouse_name	仓库名	是	[string]	刘东的本地仓
data>>list>>item>>supplier_name	供应商名	是	[string]	liudong3
data>>list>>item>>expect_arrive_time	预计到货时间	是	[string]	
data>>total	数量总计	是	[int]	20
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8E634F0B-2787-7C9B-457E-82D0BE09BD64",
    "response_time": "2024-08-02 17:50:29",
    "data": {
        "list": [
            {
                "order_sn": "POP240802001",
                "warehouse_name": "111",
                "outsource_warehouse_name": "testgong",
                "supplier_name": "testgong",
                "create_time": "2024-08-02 11:35:29",
                "status_text": "待下单",
                "create_realname": "叶盈莹",
                "ptp_sn": "PTP240802002",
                "product_name": "product",
                "sku": "0049",
                "fnsku": "",
                "outsource_quantity": 100,
                "receive_quantity": 0,
                "expect_arrive_time": "",
                "msku": [],
                "plan_sn": [
                    "PP240802007"
                ],
                "seller_name": "",
                "item": [
                    {
                        "purchase_order_sn": "PO240802006",
                        "product_name": "product 1",
                        "sku": "0047",
                        "quantity_require": 100,
                        "create_time": "2024-08-02 11:35:28",
                        "create_realname": "叶盈莹",
                        "expect_arrive_time": "",
                        "warehouse_name": "testgong",
                        "supplier_name": "心之钢联盟",
                        "quantity_real": 100,
                        "quantity_entry": 100
                    },
                    {
                        "purchase_order_sn": "PO240802007",
                        "product_name": "product 2",
                        "sku": "0048",
                        "quantity_require": 100,
                        "create_time": "2024-08-02 11:35:28",
                        "create_realname": "叶盈莹",
                        "expect_arrive_time": "",
                        "warehouse_name": "testgong",
                        "supplier_name": "wwj",
                        "quantity_real": 100,
                        "quantity_entry": 100
                    }
                ]
            }
        ],
        "total": 1
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
创建已完成的采购变更单
下一章节 
添加采购单物流信息