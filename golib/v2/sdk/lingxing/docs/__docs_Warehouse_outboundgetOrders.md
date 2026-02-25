查询出库单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/outbound/getOrders	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
wid	系统仓库id	否	[string]	53
search_field_time	日期筛选类型：
创建时间 create_time
出库时间 opt_time
更新时间 increment_time	否	[string]	create_time
start_date	日期查询开始时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-30
end_date	日期查询结束时间，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-30
order_sn	出库单单号，多个使用英文逗号分隔	否	[string]	OB240730003
idempotent_code	客户参考号，多个使用英文逗号分隔	否	[string]	OB240730003
status	出库单状态：
10 待提交
30 待出库
40 已完成
50 已撤销
121 待审批
122 已驳回	否	[int]	40
type	出库类型：
11 其他出库
12 FBA出库
14 退货出库
15 调拨出库
16 WFS出库
17 Temu出库	否	[int]	15
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": 53,
    "search_field_time": "create_time",
    "start_date": "2024-07-30",
    "end_date": "2024-07-30",
    "order_sn": "OB240730003",
    "idempotent_code": "OB240730003",
    "status": 40,
    "type": 15
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误消息	是	[array]	
request_id	请求链路id	是	[string]	5A1EBE67-7793-9E94-7790-AA457B81B3F2
response_time	响应时间	是	[string]	2022-06-27 11:54:31
total	总数	是	[int]	456
data	响应数据	是	[array]	
data>>increment_time	单据数据更新时间	是	[string]	2021-08-22 10:14:46
data>>custom_fields	自定义字段	是	[array]	
data>>custom_fields>>id	字段ID	是	[string]	
data>>custom_fields>>name	字段名	是	[string]	
data>>custom_fields>>val_text	字段值	是	[string]	
data>>opt_realname	出库人姓名	是	[string]	
data>>opt_time	出库时间	是	[string]	2021-08-22 10:14:46
data>>opt_uid	操作人id	是	[int]	475
data>>outbound_time	自定义出库时间	是	[string]	2021-12-02 10:14:46
data>>commit_realname	提交人名称	是	[string]	
data>>commit_uid	提交人id	是	[int]	475
data>>commit_time	提交时间	是	[string]	2021-08-22 10:14:46
data>>order_sn	出库单号	是	[string]	IB210822001
data>>status	出库单状态	是	[int]	
data>>status_text	出库单状态名称	是	[string]	
data>>create_time	创建时间	是	[string]	2021-08-22 10:14:46
data>>create_uid	创建人id	是	[int]	
data>>create_realname	创建人名称	是	[string]	
data>>purchase_order_sn	采购单号	是	[string]	
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
data>>type	出库类型	是	[int]	
data>>type_text	出库类型名称	是	[string]	其他入库
data>>custom_type_id	自定义类型ID	是	[Long]	0
data>>custom_type_name	自定义类型名称	是	[string]	""
data>>cg_realname	采购员姓名	是	[string]	
data>>wid	仓库id	是	[string]	
data>>ware_house_name	仓库名称	是	[string]	仓库2
data>>to_wid	目的仓库id	是	[string]	
data>>to_ware_house_name	目的仓库名称	是	[string]	
data>>remark	单据备注	是	[string]	
data>>idempotent_code	客户参考号, 该字段校验唯一不可重复	是	[string]	
data>>item_list		是	[array]	
data>>item_list>>product_name	品名	是	[string]	痛仰
data>>item_list>>sku	sku	是	[string]	痛仰
data>>item_list>>fnsku	fnsku	是	[string]	
data>>item_list>>seller_id	系统店铺id	是	[string]	
data>>item_list>>price	采购单价	是	[string]	10.0000
data>>item_list>>amount	出库成本	是	[string]	100.00
data>>item_list>>fee_cost	费用	是	[string]	0.00
data>>item_list>>product_good_num	良品量	是	[int]	10
data>>item_list>>product_bad_num	次品量	是	[int]	
data>>item_list>>product_qc_num	待检量	是	[int]	
data>>item_list>>product_total	出库量	是	[int]	10
data>>item_list>>product_amounts	货值	是	[string]	100.00
data>>item_list>>single_fee	单位费用	是	[string]	0.0000
data>>item_list>>single_stock_cost	单位出库成本	是	[string]	10.0000
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
    "request_id": "92F6872A-B022-8FA6-8152-EAB2F401659F",
    "response_time": "2024-07-31 16:23:44",
    "data": [
        {
            "opt_realname": "朱蒙蒙",
            "opt_time": "2024-07-30 11:05:29",
            "outbound_time": "2024-07-30 11:05:29",
            "opt_uid": 33,
            "commit_realname": "朱蒙蒙",
            "commit_uid": 33,
            "commit_time": "2024-07-30 11:05:29",
            "order_sn": "OB240730003",
            "status": 40,
            "status_text": "已完成",
            "create_time": "2024-07-30 11:05:29",
            "create_uid": 33,
            "create_realname": "朱蒙蒙",
            "purchase_order_sn": "",
            "revoke_realname": "",
            "revoke_uid": 0,
            "revoke_time": "",
            "supplier_id": 0,
            "supplier_name": "",
            "source_sn": "OWS240730001",
            "order_amount": "0.000000",
            "cg_uid": 0,
            "return_price": "0.000000",
            "currency": "￥",
            "other_fee": "0.000000",
            "fee_part_type": 0,
            "fee_part_type_text": "不分配",
            "type": 15,
            "type_text": "调拨出库",
            "custom_type_id": 0,
            "custom_type_name": "",
            "cg_realname": "",
            "wid": 53,
            "ware_house_name": "义乌",
            "to_wid": 9,
            "to_ware_house_name": "小徐仓",
            "remark": "备货单号：OWS240730001",
            "increment_time": "2024-07-30 11:05:29",
            "custom_fields": [],
            "item_list": [
                {
                    "product_name": "abc003-BWY-xin",
                    "sku": "abc003-BWY",
                    "fnsku": "",
                    "seller_id": "0",
                    "product_total": 10,
                    "product_good_num": 10,
                    "product_bad_num": 0,
                    "product_qc_num": 0,
                    "amount": "0.00",
                    "fee_cost": "0.00",
                    "head_way_fee": "0.00",
                    "product_amounts": "0.00",
                    "single_fee": "0.0000",
                    "unit_head_fee": "0.0000",
                    "single_stock_cost": "0.0000",
                    "price": "0.0000",
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
添加出库单
下一章节 
撤销出库单