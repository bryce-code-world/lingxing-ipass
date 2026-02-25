查询入库单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/inbound/getOrders	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
wid	系统仓库id	否	[int]	1
search_field_time	日期筛选类型：
创建时间 create_time
入库时间 opt_time
更新时间 increment_time	否	[string]	create_time
start_date	日期查询开始时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-30
end_date	日期查询结束时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-30
order_sn	入库单单号，多个使用英文逗号分隔	否	[string]	IB240730005
inbound_idempotent_code	客户参考单号，多个使用英文逗号分隔	否	[string]	IB240730005
status	入库单状态：
10 待提交
20 待入库
40 已完成
50 已撤销
121 待审批
122 已驳回	否	[int]	40
type	入库类型：
1 其他入库
2 采购入库
3 调拨入库
4 赠品入库
26 退货入库
27 移除入库	否	[int]	2
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": "1",
    "search_field_time": "create_time",
    "start_date": "2024-07-30",
    "end_date": "2024-07-30",
    "order_sn": "IB240730005",
    "inbound_idempotent_code": "IB240730005",
    "status": 40,
    "type": 2
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
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
total	总数	是	[int]	100
data	响应数据	是	[array]	
data>>increment_time	单据数据更新时间	是	[string]	2021-08-22 10:14:46
data>>custom_fields	自定义字段	是	[array]	
data>>custom_fields>>id	字段ID	是	[string]	
data>>custom_fields>>name	字段名	是	[string]	
data>>custom_fields>>val_text	字段值	是	[string]	
data>>opt_realname	入库人姓名	是	[string]	
data>>opt_time	入库时间	是	[string]	2021-08-22 10:14:46
data>>opt_uid	操作人id	是	[int]	475
data>>inbound_time	自定义入库时间	是	[string]	2022-12-02 12:14:46
data>>commit_realname	提交人名称	是	[string]	
data>>commit_uid	提交人id	是	[int]	475
data>>commit_time	提交时间	是	[string]	2021-08-22 10:14:46
data>>order_sn	入库单号	是	[string]	IB210822001
data>>status	入库单状态	是	[int]	
data>>status_text	入库单状态名称	是	[string]	
data>>create_time	创建时间	是	[string]	2021-08-22 10:14:46
data>>create_uid	创建人id	是	[int]	
data>>create_realname	创建人名称	是	[string]	
data>>purchase_order_sn	采购单号	是	[string]	
data>>receipt_order_sn	收货单号	是	[string]	CR230825001
data>>revoke_realname	撤销人名称	是	[string]	
data>>revoke_uid	撤销人id	是	[int]	
data>>revoke_time	撤销时间	是	[string]	
data>>supplier_id	供应商id	是	[string]	
data>>supplier_name	供应商名称	是	[string]	1
data>>source_sn	关联单据号	是	[string]	
data>>order_amount	单据入库成本	是	[string]	100.00
data>>cg_uid	采购员id	是	[int]	
data>>return_price	运费	是	[string]	0.00
data>>currency	运费币种	是	[string]	
data>>other_fee	其他费用	是	[string]	0.00
data>>fee_part_type	费用分摊方式	是	[string]	不分配
data>>fee_part_type_text	费用分摊方式名称	是	[string]	
data>>type	入库类型	是	[int]	
data>>type_text	入库类型名称	是	[string]	其他入库
data>>custom_type_id	自定义类型ID	是	[Long]	0
data>>custom_type_name	自定义类型名称	是	[string]	""
data>>cg_realname	采购员姓名	是	[string]	
data>>wid	仓库id	是	[string]	
data>>ware_house_name	仓库名称	是	[string]	仓库2
data>>remark	单据备注	是	[string]	
data>>inbound_idempotent_code	（入库单）客户参考号, 该字段校验唯一不可重复	是	[string]	
data>>item_list	产品明细	是	[array]	
data>>item_list>>product_name	品名	是	[string]	痛仰
data>>item_list>>sku	sku	是	[string]	sku-1
data>>item_list>>fnsku	fnsku	是	[string]	
data>>item_list>>seller_id	系统店铺id	是	[string]	
data>>item_list>>purchase_item_id	采购单子项id	是	[int]	45
data>>item_list>>price	采购单价	是	[string]	10.0000
data>>item_list>>amount	入库成本	是	[string]	100.00
data>>item_list>>fee_cost	费用	是	[string]	0.00
data>>item_list>>product_good_num	良品量	是	[int]	10
data>>item_list>>product_bad_num	次品量	是	[int]	
data>>item_list>>product_qc_num	待检量	是	[int]	
data>>item_list>>product_total	入库量	是	[int]	10
data>>item_list>>product_amounts	货值	是	[string]	100.00
data>>item_list>>single_fee	单位费用	是	[string]	0.0000
data>>item_list>>single_stock_cost	单位入库成本	是	[string]	10.0000
data>>item_list>>product_remark	产品备注	是	[string]	产品备注
data>>item_list>>custom_fields	自定义字段	是	[array]	
data>>item_list>>custom_fields>>id	字段ID	是	[string]	
data>>item_list>>custom_fields>>name	字段名	是	[string]	
data>>item_list>>custom_fields>>val_text	字段值	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "26BEDD0E-4DC9-F663-7E29-4C0C66142D50",
    "response_time": "2024-07-31 16:16:52",
    "data": [
        {
            "order_sn": "IB240730005",
            "wid": 1,
            "purchase_order_sn": "PO240729025",
            "opt_uid": 61,
            "opt_time": "2024-07-30 10:20:42",
            "inbound_time": "2024-07-30 10:20:42",
            "create_uid": 61,
            "create_time": "2024-07-30 10:20:42",
            "commit_uid": 61,
            "commit_time": "2024-07-30 10:20:42",
            "revoke_uid": 0,
            "revoke_time": "",
            "supplier_id": 1788,
            "supplier_name": "中山市羽翼照明科技有限公司",
            "order_amount": "0.000000",
            "cg_uid": 116,
            "return_price": "0.00",
            "other_fee": "0.00",
            "fee_part_type": 0,
            "fee_part_type_text": "不分配",
            "opt_realname": "刘可怡",
            "type": 2,
            "type_text": "采购入库",
            "custom_type_id": 0,
            "custom_type_name": "",
            "status": 40,
            "status_text": "已完成",
            "source_sn": "PO240729025",
            "receipt_order_sn": "CR240730005",
            "currency": "￥",
            "cg_realname": "陈志杰",
            "create_realname": "刘可怡",
            "commit_realname": "刘可怡",
            "revoke_realname": "",
            "ware_house_name": "wcn测试仓库测试改名",
            "remark": "采购单号：PO240729025，预检",
            "increment_time": "2024-07-30 10:20:43",
            "custom_fields": [],
            "item_list": [
                {
                    "product_name": "单品B-new",
                    "sku": "SINGLE-B",
                    "fnsku": "",
                    "seller_id": "120",
                    "product_total": 50,
                    "product_good_num": 0,
                    "product_bad_num": 2,
                    "product_qc_num": 0,
                    "product_onshelf_num": 0,
                    "product_shelf_num": 48,
                    "price": "0.0000",
                    "amount": "0.00",
                    "fee_cost": "0.00",
                    "head_way_fee": "0.00",
                    "product_amounts": "0.00",
                    "single_fee": "0.0000",
                    "unit_head_fee": "0.0000",
                    "single_stock_cost": "0.0000",
                    "purchase_item_id": 114411,
                    "product_remark": "",
                    "custom_fields": []
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
添加入库单
下一章节 
撤销入库单