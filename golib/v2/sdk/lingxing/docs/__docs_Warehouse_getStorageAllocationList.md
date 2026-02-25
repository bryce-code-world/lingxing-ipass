查询调拨单列表

支持查询本地仓库调拨单列表

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/getStorageAllocationList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	出库仓库id，多个以英文逗号分隔	否	[string]	432
to_wid	入库仓库id，多个以英文逗号分隔	否	[string]	581
search_date_type	时间类型：【不传或传空则默认为 1】
1 创建时间
2 调拨时间
3 完成时间
4 更新时间	否	[int]	1
start_date	开始日期，格式：Y-m-d，只有和结束日期同时有值才会生效	否	[string]	2024-07-25
end_date	结束日期，格式：Y-m-d，只有和开始日期同时有值才会生效	否	[string]	2024-07-30
page	当前页码，默认1	否	[int]	1
page_size	分页条数，默认15	否	[int]	15
请求示例
{
    "wid": "432",
    "to_wid": "581",
    "search_date_type": 1,
    "start_date": "2024-07-25",
    "end_date": "2024-07-30",
    "page": 1,
    "page_size": 15
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
request_id	请求链路id	是	[string]	86C7B3F5-C488-9BA4-EBD9-5E0525E51699
response_time	响应时间	是	[string]	2022-06-25 11:24:44
data	响应数据	是	[array]	
data>>order_sn	单据号	是	[string]	TF220624002
data>>wid	出库仓库id	是	[int]	11
data>>to_wid	入库仓库id	是	[int]	1
data>>opt_uid	操作人（用户id）	是	[int]	175
data>>opt_time	操作时间	是	[string]	2022-06-24 18:05
data>>order_amount	订单金额	是	[number]	0.000000
data>>out_bin_type	出仓类型：
0 不指定仓位
1 指定仓位	是	[int]	
data>>ware_house_bak_name	出库仓库名称	是	[string]	仓库11
data>>to_ware_house_bak_name	入库仓库名称	是	[string]	仓库1
data>>status	单据状态：
5-待提交
10 待调拨
19 待收货
20 已完成
30 已删除
121 待审批
122 已驳回	是	[int]	20
data>>status_text	单据状态文本	是	[string]	已完成
data>>create_time	创建时间	是	[string]	2022-06-24 18:05:46
data>>create_uid	创建人（用户id）	是	[int]	175
data>>transfer_time	操作调拨时间	是	[string]	2022-06-24 18:05:46
data>>transfer_uid	操作调拨人（用户id）	是	[int]	175
data>>remark	单据备注	是	[string]	
data>>freight_fee	运费	是	[number]	0.00
data>>other_fee	其他费用	是	[number]	0.00
data>>fee_part_type	费用分摊方式：
0 不分摊
1 按金额分摊
2 按sku数量分摊
3 按重量
4 按体积
5 按自定义	是	[int]	1
data>>type	调拨类型：1-简易调拨；2-完整调拨	是	[int]	1
data>>receive_uid	确认收货操作人（用户id）	是	[int]	
data>>receive_time	确认收货收货时间	是	[string]	
data>>finish_time	单据完成时间	是	[string]	2022-06-24 18:05:46
data>>finish_uid	单据完成操作人（用户id）	是	[int]	175
data>>increment_time	单据更新时间	是	[string]	"2025-02-11 10:49:37"
data>>outbound_order_sn	关联的出库单单号	是	[string]	OB220624004
data>>inbound_order_sn	关联的入库单单号	是	[string]	IB220624016
data>>opt_realname	操作人姓名	是	[string]	JSTAPI
data>>create_realname	创建人姓名	是	[string]	JSTAPI
data>>transfer_realname	确认调拨操作人姓名	是	[string]	JSTAPI
data>>receive_realname	确认收货操作人姓名	是	[string]	
data>>finish_realname	单据完成操作人姓名	是	[string]	JSTAPI
data>>fee_part_type_text	费用分摊方式文本	是	[string]	不分配
data>>out_bin_type_text	出仓方式文本	是	[string]	不指定仓位
data>>item_list	产品明细列表	是	[array]	
data>>item_list>>product_id	产品id	是	[int]	2117165
data>>item_list>>sku	SKU	是	[string]	2022052702
data>>item_list>>fnsku	FNSKU	是	[string]	
data>>item_list>>product_total	调拨总量	是	[int]	1
data>>item_list>>product_good_num	调拨可用量	是	[int]	1
data>>item_list>>product_bad_num	调拨次品量	是	[int]	
data>>item_list>>price	单价	是	[number]	0.000000
data>>item_list>>amount	货值	是	[number]	0.000000
data>>item_list>>seller_id	店铺id	是	[int]	
data>>item_list>>to_available_bin	入库可用仓位编码	是	[string]	ts_valid
data>>item_list>>to_inferior_bin	入库次品仓位编码	是	[string]	ts_bad
data>>item_list>>seller_name	店铺名称	是	[string]	
data>>item_list>>country_name	店铺所属国家名称	是	[string]	
data>>item_list>>product_name	品名	是	[string]	2022052702
data>>item_list>>in_available_storage_bin_code	入库可用仓位名称	是	[string]	可用暂存
data>>item_list>>in_inferior_storage_bin_code	入库次品仓位名称	是	[string]	
data>>item_list>>outboundList	出库仓位列表	是	[array]	
data>>item_list>>outboundList>>whb_code	仓位编码	是	[string]	ts_valid
data>>item_list>>outboundList>>whb_code_name	仓位名称	是	[string]	可用暂存
data>>item_list>>outboundList>>product_num	出仓数量	是	[int]	1
data>>item_list>>outboundList>>whb_type	仓位类型	是	[int]	2
data>>item_list>>outboundList>>whb_type_name	仓位类型名称	是	[string]	可用暂存
data>>item_list>>out_stock_cost	出库单位费用	是	[string]	0.0000
data>>item_list>>pic_url	产品图片链接	是	[string]	
data>>item_list>>cg_package_length	包装规格-长（CM）	是	[number]	0.00
data>>item_list>>cg_package_width	包装规格-宽（CM）	是	[number]	0.00
data>>item_list>>cg_package_height	包装规格-高（CM）	是	[number]	0.00
data>>item_list>>cg_product_gross_weight	单品净重	是	[number]	0.00
data>>item_list>>freight_fee_unit	单位运费	是	[number]	1.0000
data>>item_list>>other_fee_unit	单位其他费用	是	[number]	1.0000
data>>item_list>>product_remark	产品备注	是	[string]	产品备注
data>>item_list>>out_available_bin_list	出库可用仓位列表	是	[array]	
data>>item_list>>out_available_bin_list>>whb_code	仓位编码	是	[string]	ts_valid
data>>item_list>>out_available_bin_list>>whb_type	仓位类型	是	[int]	2
data>>item_list>>out_available_bin_list>>whb_code_name	仓位名称	是	[string]	可用暂存
data>>item_list>>out_available_bin_list>>whb_type_name	仓位类型文本	是	[string]	可用暂存
data>>item_list>>out_available_bin_list>>product_num	出库可用量	是	[int]	1
data>>item_list>>out_inferior_bin_list	出库次品仓位列表	是	[array]	
data>>item_list>>out_inferior_bin_list>>whb_code	仓位编码	是	[string]	RB-001
data>>item_list>>out_inferior_bin_list>>whb_type	仓位类型	是	[int]	6
data>>item_list>>out_inferior_bin_list>>whb_code_name	仓位名称	是	[string]	RB-001
data>>item_list>>out_inferior_bin_list>>whb_type_name	仓位类型文本	是	[string]	次品
data>>item_list>>out_inferior_bin_list>>product_num	出库次品量	是	[int]	2
data>>item_list>>in_available_bin_list	入库可用仓位列表	是	[array]	
data>>item_list>>in_available_bin_list>>whb_code	仓位编码	是	[string]	ts_valid
data>>item_list>>in_available_bin_list>>whb_type	仓位类型	是	[int]	2
data>>item_list>>in_available_bin_list>>whb_code_name	仓位名称	是	[string]	可用暂存
data>>item_list>>in_available_bin_list>>whb_type_name	仓位类型文本	是	[string]	可用暂存
data>>item_list>>in_available_bin_list>>product_num	入库可用量	是	[int]	1
data>>item_list>>in_inferior_bin_list	入库次品仓位列表	是	[array]	
data>>item_list>>in_inferior_bin_list>>whb_code	仓位编码	是	[string]	ts_bad
data>>item_list>>in_inferior_bin_list>>whb_type	仓位类型	是	[int]	3
data>>item_list>>in_inferior_bin_list>>whb_code_name	仓位名称	是	[string]	次品暂存
data>>item_list>>in_inferior_bin_list>>whb_type_name	仓位类型文本	是	[string]	次品暂存
data>>item_list>>in_inferior_bin_list>>product_num	入库次品量	是	[int]	2
data>>goodTotalNum	订单可用调拨总量	是	[int]	1
data>>badTotalNum	订单次品调拨总量	是	[int]	
total	查询总数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "E9DE57E4-C265-7434-874B-C890A10BF515",
    "response_time": "2022-12-14 17:54:53",
    "data": [
        {
            "order_sn": "TF221214002",
            "wid": 1,
            "to_wid": 11,
            "opt_uid": 129130,
            "opt_time": "2022-12-14 17:49",
            "order_amount": "0.000000",
            "out_bin_type": 1,
            "ware_house_bak_name": "仓库1",
            "to_ware_house_bak_name": "仓库11",
            "status": 5,
            "create_time": "2022-12-14 15:36:02",
            "create_uid": 129130,
            "transfer_time": "",
            "transfer_uid": 0,
            "remark": "",
            "freight_fee": "0.00",
            "other_fee": "0.00",
            "fee_part_type": 0,
            "type": 2,
            "receive_uid": 0,
            "receive_time": "",
            "finish_time": "",
            "finish_uid": 0,
            "increment_time": "2022-12-14 10:49:37",
            "outbound_order_sn": null,
            "inbound_order_sn": null,
            "opt_realname": "欧树权",
            "create_realname": "欧树权",
            "transfer_realname": "",
            "receive_realname": "",
            "finish_realname": "",
            "fee_part_type_text": "不分配",
            "out_bin_type_text": "指定仓位",
            "status_text": "待提交",
            "item_list": [
                {
                    "product_id": 2124879,
                    "sku": "rwb-chanpin2",
                    "fnsku": "",
                    "product_total": 3,
                    "product_good_num": 1,
                    "product_bad_num": 2,
                    "price": "0.000000",
                    "amount": "0.000000",
                    "seller_id": 0,
                    "cg_package_length": "0.00",
                    "cg_package_width": "0.00",
                    "cg_package_height": "0.00",
                    "cg_product_gross_weight": "0.00",
                    "freight_fee_unit": "0.0000",
                    "other_fee_unit": "0.0000",
                    "to_available_bin": "",
                    "to_inferior_bin": "",
                    "product_remark": "",
                    "seller_name": "",
                    "country_name": "",
                    "product_name": "rwb-chanpin2",
                    "out_available_bin_list": [
                        {
                            "whb_code": "ts_valid",
                            "whb_type": 2,
                            "whb_code_name": "可用暂存",
                            "whb_type_name": "可用暂存",
                            "product_num": 1,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        }
                    ],
                    "out_inferior_bin_list": [
                        {
                            "whb_code": "RB-001",
                            "whb_type": 6,
                            "whb_code_name": "RB-001",
                            "whb_type_name": "次品",
                            "product_num": 2,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        }
                    ],
                    "in_available_bin_list": [
                        {
                            "whb_code": "ts_valid",
                            "whb_type": 2,
                            "whb_code_name": "可用暂存",
                            "whb_type_name": "可用暂存",
                            "product_num": 1,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        }
                    ],
                    "in_inferior_bin_list": [
                        {
                            "whb_code": "ts_bad",
                            "whb_type": 3,
                            "whb_code_name": "次品暂存",
                            "whb_type_name": "次品暂存",
                            "product_num": 2,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        }
                    ],
                    "in_available_storage_bin_code": "",
                    "in_inferior_storage_bin_code": "",
                    "outboundList": [
                        {
                            "whb_code": "ts_valid",
                            "whb_type": 2,
                            "whb_code_name": "可用暂存",
                            "whb_type_name": "可用暂存",
                            "product_num": 1,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        },
                        {
                            "whb_code": "RB-001",
                            "whb_type": 6,
                            "whb_code_name": "RB-001",
                            "whb_type_name": "次品",
                            "product_num": 2,
                            "product_id": 2124879,
                            "sku": "rwb-chanpin2",
                            "seller_id": 0,
                            "fnsku": ""
                        }
                    ],
                    "out_stock_cost": "0.0000",
                    "pic_url": "https://image.distributetop.com/erp-vue/90124993550340096/20221206/f1bb7ed736e249cc969ad70c8a5d148c.jpeg"
                }
            ],
            "goodTotalNum": 1,
            "badTotalNum": 2
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
加工单列表
下一章节 
创建待收货/已完成的调拨单