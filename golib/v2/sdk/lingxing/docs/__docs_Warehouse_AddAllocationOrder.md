创建待收货/已完成的调拨单

支持创建“待收货/已完成”状态调拨单，请求成功后，对应调拨单在列表“待收货/已完成”状态

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/addAllocationOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	客户出库仓库id（与系统仓库出库id任一必填，优先取客户出库仓库id）	否	[int]	
sys_wid	系统仓库出库id（与客户仓库出库id任一必填，优先取客户出库仓库id）	否	[int]	
to_wid	客户入库仓库id（与系统仓库入库id任一必填，优先取客户入库仓库id）	否	[int]	
sys_to_wid	系统仓库入库id（与客户仓库入库id任一必填，优先取客户入库仓库id）	否	[int]	
freight_fee	运费	否	[string]	
other_fee	其他费用	否	[string]	
fee_part_type	费用分摊方式：【默认0】
0 不分摊
2 按sku数量分摊
3 按重量
4 按体积
5 按自定义	否	[int]	0
remark	备注	否	[string]	
type	调拨类型：【默认1】
1 简易调拨【创建已完成状态的单据】
2 完整调拨【创建待收货状态的单据】	否	[int]	1
predict_time	预计到货时间，格式：Y-m-d	否	[string]	2022-08-01
product_list		是	[array]	
product_list>>sku	sku	是	[string]	
product_list>>seller_id	店铺id，不传默认0【对应查询亚马逊店铺信息接口字段sid】	否	[int]	
product_list>>fnsku	fnsku，不传默认为空	否	[string]	
product_list>>good_num	可用调拨量，和次品调拨量其中一项必填	否	[int]	
product_list>>bad_num	次品调拨量，和可用调拨量其中一项必填	否	[int]	
product_list>>cg_package_length	包装规格-长（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_package_width	包装规格-宽（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_package_height	包装规格-高（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_product_gross_weight	单品净重，不传或者传空则取产品管理的值	否	[string]	
product_list>>freight_fee_unit	单位运费，若费用分摊方式为5-自定义则必填，其他则该字段无效	否	[string]	
product_list>>other_fee_unit	单位其他费用，若费用分摊方式为5-自定义则必填，其他则该字段无效	否	[string]	 
product_list>>product_remark	明细备注，最大长度为255个字符	否	[string]	备注
out_available_bin	出库可用仓位列表	否	[array]	
out_available_bin>>whb_code	出库可用仓位编码	否	[string]	A-1-1
out_available_bin>>num	出库可用仓位调拨量	否	[string]	1
out_inferior_bin	出库次品仓位列表	否	[array]	
out_inferior_bin>>whb_code	出库次品仓位编码	否	[string]	A-1-1
out_inferior_bin>>num	出库次品仓位数量	否	[string]	1
to_available_bin	入库可用仓位列表	否	[array]	
to_available_bin>>whb_code	入库可用仓位编码	否	[string]	A-1-1
to_available_bin>>num	入库可用仓位数量	否	[string]	1
to_inferior_bin	入库次品仓位列表	否	[array]	
to_inferior_bin>>whb_code	入库次品仓位编码	否	[string]	A-1-1
to_inferior_bin>>num	入库次品仓位数量	否	[string]	1
out_bin_type	0 默认
1 出库仓位不为空时，必传	否	[string]	1
请求示例
{
    "wid": "",
    "sys_wid": 1,
    "to_wid": "",
    "sys_to_wid": 2,
    "freight_fee": "",
    "other_fee": "",
    "fee_part_type": 0,
    "remark": "备注xx",
    "type": 1,
    "predict_time": "2022-08-01",
    "product_list": [
        {
            "sku": "sku-123",
            "seller_id": "",
            "fnsku": "",
            "good_num": 1,
            "bad_num": 0,
            "cg_package_length": "",
            "cg_package_width": "",
            "cg_package_height": "",
            "cg_product_gross_weight": "",
            "freight_fee_unit": "",
            "other_fee_unit": "",
            "product_remark": "产品备注"
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
message	返回信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C70FD2A7-4E7D-6B32-7D72-FC42587CE97E
response_time	响应时间	是	[string]	2022-07-29 14:33:42
data	响应数据	是	[object]	
data>>order_sn	调拨单号	是	[string]	TF220729002
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C70FD2A7-4E7D-6B32-7D72-FC42587CE97E",
    "response_time": "2022-07-29 14:33:42",
    "data": {
        "order_sn": "TF220729002"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询调拨单列表
下一章节 
调拨单全部收货