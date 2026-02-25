查询批次明细
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/getBatchDetailList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限400	否	[int]	20
show_zero_stock	是否显示0库存信息：0 不显示，1 显示	否	[int]	0
wids	仓库id，多个使用英文逗号分隔	否	[string]	1,5,18
stock_in_type_list	入库类型，多个使用英文逗号分隔：
19 其他入库
22 采购入库
24 调拨入库
23 委外入库
25 盘盈入库
16 换标入库
17 加工入库
18 拆分入库
26 退货入库
27 移除入库
45 赠品入库	否	[string]	19
search_field	搜索字段：
sku SKU
msku MSKU
fnsku FNSKU
order_sn 单据号
product_name 品名
batch_number 批次号
receipt_order 收货单
purchase_order 采购单
purchase_plan 采购计划
source_batch_number 源头批次号	否	[string]	product_name
search_value	搜索值	否	[string]	pan1575
请求示例
{
    "offset": 0,
    "length": 20,
    "show_zero_stock": 0,
    "wid": "1,5,18",
    "stock_in_type_list": "19",
    "search_field": "product_name",
    "search_value": "pan1575"
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
request_id	请求链路id	是	[string]	D76CFAC8-3385-0A42-8916-A168F3449996
response_time	响应时间	是	[string]	2022-03-23 09:17:03
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>batch_no	批次号	是	[string]	401352726865658370-1
data>>source_batch_no	源头批次号	是	[string]	["401351566731194390-1"]
data>>order_sn	入库单号	是	[string]	PT230901004
data>>type	入库类型	是	[int]	1701
data>>type_name	入库类型描述	是	[string]	加工入库
data>>product_id	本地产品id	是	[int]	81892
data>>product_name	品名	是	[string]	商品1
data>>sku	SKU	是	[string]	0710-1
data>>store_id	店铺id	是	[string]	0
data>>store_name	店铺名称	是	[string]	店铺1
data>>msku	MSKU	是	[string]	MSKU123
data>>fnsku	FNSKU	是	[string]	FNSKU412
data>>wid	仓库id	是	[int]	24
data>>wh_name	仓库名称	是	[string]	nickey111
data>>total	结存总数【在途结存 + 在库结存】	是	[int]	20
data>>transit_balance_num	在途结存	是	[int]	15
data>>balance_num	在库结存	是	[int]	5
data>>good_transit_num	可用在途量	是	[int]	2
data>>bad_transit_num	次品在途量	是	[int]	3
data>>qc_num	待检量	是	[int]	4
data>>good_num	可用量	是	[int]	1
data>>bad_num	次品量	是	[int]	10
data>>plan_sn	采购计划单号信息	是	[array]	["PP230901003", "PP230901002", "PP230901001"]
data>>purchase_order_sns	采购单单号信息	是	[array]	["PO230901001"]
data>>delivery_order_sns	收货单单号信息	是	[array]	["CR230901001"]
data>>supplier_ids	供应商id信息	是	[array]	["5221"]
data>>supplier_names	供应商名称列表	是	[array]	["供应商1"]
data>>amount	货值	是	[string]	2.5
data>>head_stock_cost	头程	是	[string]	3.1
data>>fee	费用	是	[string]	1.2
data>>stock_cost	库存成本	是	[string]	2.3
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "FE236155-8309-D871-F4B9-DC5D49F9EB36",
    "response_time": "2023-09-04 18:41:31",
    "data": [
        {
            "wid": 1,
            "wh_name": "默认仓库",
            "batch_no": "401286103900685831-1",
            "source_batch_no": [],
            "order_sn": "IB230228023",
            "type": 2202,
            "type_name": "采购入库",
            "product_id": 18019,
            "product_name": "A022401",
            "sku": "A022401",
            "msku": "",
            "fnsku": "",
            "store_id": "0",
            "store_name": "",
            "total": 1,
            "transit_balance_num": 0,
            "balance_num": 1,
            "good_transit_num": 0,
            "bad_transit_num": 0,
            "qc_num": 0,
            "good_num": 1,
            "bad_num": 0,
            "plan_sn": [],
            "purchase_order_sns": [
                "PO230228001"
            ],
            "delivery_order_sns": [
                "CR230228017"
            ],
            "supplier_ids": [
                "359"
            ],
            "supplier_names": [
                "xxx服装厂"
            ],
            "stock_cost": "5.23",
            "fee": "1.23",
            "head_stock_cost": "0.00",
            "amount": "4.00"
        }
    ],
    "total": 15038
}
复制
错误
复制成功
 上一章节
查询仓位库存明细
下一章节 
查询批次流水