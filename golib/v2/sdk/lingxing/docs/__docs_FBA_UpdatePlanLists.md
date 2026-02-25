编辑FBA发货计划
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/updateShipmentPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	发货计划单号	是	[string]	R221215003
shipment_time	发货时间，格式：Y-m-d	否	[string]	2023-01-03
packing_type	包装类型： 1 混装，2 原厂	否	[int]	2
logistics_provider_id	物流商id	否	[int]	1
logistics_channel_id	物流渠道id	否	[int]	12
shipment_plan_quantity	计划发货量	否	[int]	6
quantity_in_case	单箱数量（PCS）	否	[int]	3
box_num	箱数	否	[int]	2
sys_wid	系统仓库id【发货仓库】	否	[int]	1
cg_package_length	包装规格长（cm）【保留两位小数】	否	[number]	10.56
cg_package_width	包装规格宽（cm）【保留两位小数】	否	[number]	10.56
cg_package_height	包装规格高（cm）【保留两位小数】	否	[number]	10.56
cg_box_length	箱规长（cm）【保留两位小数】	否	[number]	10.56
cg_box_width	箱规宽（cm）【保留两位小数】	否	[number]	10.56
cg_box_height	箱规高（cm）【保留两位小数】	否	[number]	10.56
nw	单品净重（g）【保留两位小数】	否	[number]	10.56
gw	单品毛重（g）【保留两位小数】	否	[number]	10.56
cg_box_weight	单箱重量（kg）【保留两位小数】	否	[number]	10.56
remark	备注	否	[string]	备注1
请求示例
{
    "order_sn": "R221215003",
    "shipment_time": "2023-01-03",
    "packing_type": 2,
    "logistics_provider_id": 1,
    "logistics_channel_id": 12,
    "shipment_plan_quantity": 6,
    "quantity_in_case": 3,
    "box_num": 2,
    "sys_wid": 1,
    "cg_package_length": 10.56,
    "cg_package_width": 10.56,
    "cg_package_height": 10.56,
    "cg_box_length": 10.56,
    "cg_box_width": 10.56,
    "cg_box_height": 10.56,
    "nw": 10.56,
    "gw": 10.56,
    "cg_box_weight": 10.56,
    "remark": "备注示例"
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
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2022-10-18 22:39:55
total	总数	是	[int]	0
data	响应数据	是	[array]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "951E55ED-1038-A7DF-9060-FBCDDDB06FB0",
    "response_time": "2023-02-14 16:09:22",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
创建FBA发货计划
FBA
下一章节 
查询FBA发货单列表
FBA