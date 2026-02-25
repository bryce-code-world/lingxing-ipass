查询库存流水（新）
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryLog/WareHouseInventory/wareHouseCenterStatement	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wids	仓库id，多个使用英文逗号分隔	否	[string]	1,578,765
types	流水类型，多个使用英文逗号分隔：【不填默认全部类型】
19 其他入库
22 采购入库
24 调拨入库
23 委外入库
25 盘盈入库
15 FBM退货
16 换标入库
17 加工入库
18 拆分入库
26 退货入库
27 移除入库
28 采购质检
29 委外质检
71 采购上架
72 委外上架
42 其他出库
41 调拨出库
32 委外出库
33 盘亏出库
34 换标出库
35 加工出库
36 拆分出库
37 FBA出库
38 FBM出库
39 退货出库
65 WFS出库
100 锁定流水
51 销毁出库	否	[string]	19
sub_types	子类流水类型，多个使用英文逗号分隔：【不填默认全部类型】
1901 其他入库 手工其他入库
1902 其他入库 用户初始化
1903 其他入库 系统初始化
2201 采购入库 手工采购入库
2202 采购入库 采购单创建入库单
2801 采购质检 质检
7101 采购上架 PDA上架入库
7201 委外上架 PDA委外上架
2401 调拨入库 调拨单入在途
2402 调拨入库 调拨单收货
2403 调拨入库 备货单入在途
2404 调拨入库 备货单收货
2405 调拨入库 备货单入库结束到货
2301 委外入库 委外订单完成加工后入库
2901 委外质检 委外订单质检
2501 盘盈入库 盘点单入库
2502 盘盈入库 数量调整单正向
1501 FBM退货 退货入库
1502 FBM退货 退货入库质检
1601 换标入库 换标调整入库
1701 加工入库 加工单入库
1702 加工入库 委外订单加工入库
1801 拆分入库 拆分单入库
2601 自动退货入库
2602 手动退货入库
2701 移除入库
4201 其他出库 手工其他出库
4101 调拨出库 调拨单出库
4102 调拨出库 备货单出库
3201 委外出库 委外订单完成加工后出库
3301 盘亏出库 盘点单出库
3302 盘亏出库 数量调整单负向
3401 换标出库 换标调整出库
3501 加工出库 加工单出库
3502 加工出库 委外订单加工出库
3601 拆分出库 拆分单出库
3701 FBA出库 发货单出库
3702 FBA出库 手工FBA出库
3801 FBM出库 销售出库单
3901 退货出库 手工退货出库
3902 退货出库 采购单生成的退货出库单
10001 库存锁定-出库
10002 库存锁定-调拨
10003 库存锁定-调整
10004 库存锁定-加工
10005 库存锁定-加工计划
10006 库存锁定-拆分
10007 库存锁定-海外备货
10008 库存锁定-发货
10009 库存锁定-自发货
10010 库存锁定-主动释放
10012 库存锁定-发货拣货
10013 库存锁定-发货计划
10014 库存锁定-WFS库存调整
10011 仓位转移和一键上架	否	[string]	1901
start_date	操作开始时间，格式：Y-m-d，闭区间，联合结束时间使用	否	[string]	2024-06-29
end_date	操作结束时间，格式：Y-m-d，开区间，联合开始时间使用	否	[string]	2024-07-29
offset	分页偏移量，默认0	是	[int]	0
length	分页长度，默认20	是	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wids": "1,578,765",
    "types": "19",
    "sub_types": "1901",
    "start_date": "2024-06-29",
    "end_date": "2024-07-29"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	6C4B7038-6671-73B1-BF83-752314DE7AFE
response_time	响应时间	是	[string]	2023-02-22 10:29:32
data	响应数据	是	[array]	
data>>wid	仓库id	是	[int]	1
data>>ware_house_name	仓库名称	是	[string]	波兰仓库
data>>order_sn	操作单据号	是	[string]	WO103283976471200256
data>>product_id	产品id	是	[int]	90
data>>product_name	品名	是	[string]	产品1
data>>sku	sku	是	[string]	sku1
data>>seller_id	店铺id	是	[int]	店铺ID
data>>fnsku	fnsku	是	[string]	FNCF8427E
data>>product_good_num	可用量	是	[int]	10
data>>product_bad_num	次品量	是	[int]	121
data>>product_qc_num	待检量	是	[int]	22
data>>product_lock_good_num	可用锁定量	是	[int]	45
data>>product_lock_bad_num	次品锁定量	是	[int]	412
data>>good_transit_num	良品在途	是	[int]	33
data>>bad_transit_num	次品在途	是	[int]	231
data>>type	流水类型	是	[int]	100
data>>type_text	流水类型文本	是	[string]	库存调整
data>>sub_type	子类型	是	[string]	10009
data>>sub_type_text	子类型文本	是	[string]	库存锁定
data>>fee_cost	总费用成本	是	[string]	542
data>>single_cg_price	采购单价	是	[string]	412
data>>single_fee_cost	单位费用	是	[string]	33
data>>single_stock_price	单位库存成本	是	[string]	12
data>>stock_cost	库存成本	是	[string]	412
data>>product_amounts	货值	是	[string]	123
data>>head_stock_price	单位头程	是	[string]	44
data>>head_stock_cost	头程	是	[string]	523
data>>opt_uid	操作人员ID	是	[int]	56
data>>opt_time	操作时间	是	[string]	2023-02-22 10:23
data>>opt_real_name	操作人员姓名	是	[string]	张三
data>>remark	备注	是	[string]	备注
data>>bid	品牌id	是	[int]	1
data>>brand_name	品牌名称	是	[string]	p1
data>>ref_order_sn	关联单据号	是	[string]	WO103283976471200256
data>>product_total	总量	是	[int]	4342
data>>good_balance_num	可用结存量	是	[int]	452
data>>bad_balance_num	次品结存量	是	[int]	234
data>>good_lock_balance_num	可用锁定结存量	是	[int]	98
data>>bad_lock_balance_num	次品锁定结存量	是	[int]	687
data>>qc_balance_num	质检结存量	是	[int]	354
data>>good_transit_balance_num	可用在途结存量	是	[int]	23
data>>statement_id	流水ID	是	[string]	401283976472981506
data>>bad_transit_balance_num	次品在途结存量	是	[int]	66
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "6C4B7038-6671-73B1-BF83-752314DE7AFE",
    "response_time": "2023-02-22 10:43:57",
    "data": [{
        "wid": 1,
        "ware_house_name": "仓库1",
        "order_sn": "IB230110003",
        "product_id": 17180,
        "product_name": "2023011002",
        "sku": "2023011002",
        "seller_id": 0,
        "fnsku": "",
        "product_good_num": -22,
        "product_bad_num": 0,
        "product_qc_num": 0,
        "product_lock_good_num": 0,
        "product_lock_bad_num": 0,
        "good_transit_num": 0,
        "bad_transit_num": 0,
        "type": 19,
        "type_text": "其他入库",
        "sub_type": "1901",
        "sub_type_text": "其他入库-手工其他入库",
        "fee_cost": "",
        "single_cg_price": "",
        "single_fee_cost": "",
        "single_stock_price": "",
        "stock_cost": "",
        "product_amounts": "",
        "head_stock_price": "",
        "head_stock_cost": "",
        "opt_uid": 128643,
        "opt_time": "2023-02-22 10:43",
        "opt_real_name": "姚xx",
        "remark": "入库单撤销",
        "bid": 0,
        "brand_name": "",
        "ref_order_sn": "",
        "product_total": -22,
        "good_balance_num": 0,
        "bad_balance_num": 0,
        "good_lock_balance_num": 0,
        "bad_lock_balance_num": 0,
        "qc_balance_num": 0,
        "good_transit_balance_num": 0,
        "bad_transit_balance_num": 0,
        "statement_id": "401283981476127232"
    }],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询库存流水（旧）
下一章节 
查询仓位流水