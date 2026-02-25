查询库存流水（旧）

支持查询本地仓库库存流水，对应系统【仓库】>【库存流水】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/wareHouseStatement	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库ID，多个仓库ID用英文逗号分隔，不填默认所有仓库	否	[string]	1,5,18
type	流水类型：【多个流水类型用英文逗号分隔，不填默认全部类型】
1 其他入库
2 采购入库
3 调拨入库
10 其它入库（已撤销）
11 其他出库
12 FBA出库
13 调拨出库
14 退货出库
15 FBM退货
16 换标入库
17 加工入库
18 拆分入库
20 采购入库（已撤销）
21 库存调整
23 委外入库
25 盘盈入库
32 委外出库
33 盘亏出库
34 换标出库
35 加工出库
36 拆分出库
43 FBM出库
50 成本补录
110 其它出库（已撤销）
120 FBA出库（已撤销）
130 调拨出库（已撤销）
140 退货出库（已撤销）
210 库存调整（已撤销）
500 成本补录（已撤销）	否	[string]	100
start_date	操作开始时间，格式：Y-m-d，闭区间，联合结束时间使用	否	[string]	2024-07-29
end_date	操作结束时间，格式：Y-m-d，开区间，联合开始时间使用	否	[string]	2024-07-29
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": "1,5,18",
    "start_date": "2024-07-29",
    "end_date": "2024-07-29",
    "type": "100"
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
request_id	请求链路id	是	[string]	CF670F06-BD65-E580-7018-47219C66AC26
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[array]	
data>>statement_id	流水ID	是	[string]	11036
data>>wid	仓库ID	是	[int]	8
data>>ware_house_name	仓库名	是	[string]	测试仓库
data>>bid	商品品牌ID	是	[int]	18
data>>order_sn	单据号	是	[string]	OB220527028
data>>product_id	商品ID	是	[int]	18232
data>>product_name	品名	是	[string]	AA-组合3
data>>sku	SKU	是	[string]	AA-组合3
data>>fnsku	FNSKU	是	[string]	
data>>product_total	商品总量	是	[number]	
data>>product_good_num	良品量	是	[number]	
data>>product_bad_num	次品量	是	[number]	
data>>product_qc_num	质检量	是	[number]	
data>>product_lock_num	锁定量	是	[number]	
data>>price	单价	是	[number]	1.00
data>>amount	金额	是	[number]	0.00
data>>type	流水类型：
1 其他入库
2 采购入库
3 调拨入库
10 其它入库（已撤销）
11 其他出库
12 FBA出库
13 调拨出库
14 退货出库
15 FBM退货
16 换标入库
17 加工入库
18 拆分入库
20 采购入库（已撤销）
21 库存调整
23 委外入库
25 盘盈入库
32 委外出库
33 盘亏出库
34 换标出库
35 加工出库
36 拆分出库
43 FBM出库
50 成本补录
110 其它出库（已撤销）
120 FBA出库（已撤销）
130 调拨出库（已撤销）
140 退货出库（已撤销）
210 库存调整（已撤销）
500 成本补录（已撤销）	是	[string]	1
data>>remark	备注	是	[string]	库存初始化
data>>opt_uid	操作人员ID	是	[int]	230
data>>opt_time	操作时间	是	[string]	2020-09-05 16:23
data>>cancel_time	撤销时间	是	[string]	2020-09-06 10:23
data>>fee_cost	费用成本	是	[string]	0.00
data>>brand_name	品牌	是	[string]	New desk
data>>single_fee_cost	单位费用成本	是	[number]	0.00
data>>single_cg_price	采购单价	是	[number]	1.00
data>>product_amounts	货值	是	[number]	0.00
data>>type_text	流水类型文本	是	[string]	其他入库
data>>opt_realname	操作人员姓名	是	[string]	李X
total	总数目	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "596A9701-F2B8-1291-3EFC-9099C295A9C4",
    "response_time": "2024-07-31 15:39:39",
    "data": [
        {
            "wid": 18,
            "ware_house_name": "深圳仓",
            "order_sn": "PT240729003",
            "product_id": 33043,
            "product_name": "电脑cyt",
            "sku": "diannaocyt",
            "seller_id": "0",
            "fnsku": "",
            "product_good_num": -10,
            "product_bad_num": 0,
            "product_qc_num": 0,
            "product_lock_num": 10,
            "price": "",
            "amount": "",
            "product_amounts": "",
            "type": 100,
            "type_text": "库存调整",
            "fee_cost": "",
            "single_cg_price": "",
            "single_fee_cost": "",
            "opt_uid": 10603434,
            "opt_time": "2024-07-29 20:12",
            "opt_realname": "马旭哲",
            "cancel_time": "",
            "remark": "加工单锁定库存",
            "bid": 0,
            "brand_name": "",
            "ref_order_sn": "",
            "product_total": 0,
            "statement_id": "401469208181637651"
        }
    ],
    "total": 32
}
复制
错误
复制成功
 上一章节
查询批次流水
下一章节 
查询库存流水（新）