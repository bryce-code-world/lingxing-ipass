添加入库单

支持添加本地仓库入库单，并增加入库仓库库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/storage/orderAdd	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	自定义仓库id，wid和sys_wid其中一项必填，都填则优先wid	否	[string]	CUSTOM-id-1
sys_wid	系统仓库id，wid和sys_wid其中一项必填，都填则优先wid	是	[int]	1
type	单据类型：
1 其他入库
2 采购入库
26 退货入库
27 移除入库	是	[int]	1
supplier_id	自定义供应商id【supplier_id、sys_supplier_id 二选一必填，都填优先取supplier_id】	否	[string]	S1929
sys_supplier_id	系统供应商id【supplier_id、sys_supplier_id 二选一必填，都填优先取supplier_id】	否	[int]	125
order_sn	采购单号【对此采购单执行快捷入库】，不支持自定义采购单号	否	[string]	PO220118001
remark	单据备注	否	[string]	备注
ship_fee	运费	否	[string]	1.22
other_fee	其它费用	否	[string]	5
fee_part_type	费用分配方式:
0 不分摊
1 按金额
2 按数量	否	[int]	2
inbound_time	自定义入库时间，格式：Y-m-d	否	[string]	2022-12-02
inbound_idempotent_code	（入库单）客户参考号, 该字段校验唯一不可重复	否	[string]	
product_list	产品明细	是	[array]	
product_list>>sku	sku	是	[string]	MDR-ZX110
product_list>>good_num	良品数量	是	[int]	10
product_list>>bad_num	次品数量	是	[int]	
product_list>>price	单价	是	[string]	8.66
product_list>>seller_id	店铺id【没有店铺时传0】
亚马逊店铺对应查询亚马逊店铺信息字段【sid】
多平台店铺对应查询多平台店铺信息字段【store_id】	是	[int]	21
product_list>>fnsku	fnsku【存在fnsku时店铺id必填，没有时传空字符串】	是	[string]	FN206265A
product_list>>good_whb_code	可用仓位	否	[string]	A-1-1
product_list>>bad_whb_code	次品仓位	否	[string]	次品暂存
请求示例
{
    "wid": "CUSTOM-ID-1",
    "sys_wid": 1,
    "type": 1,
    "supplier_id": "S1929",
    "sys_supplier_id": 125,
    "order_sn": "PO220118001",
    "remark": "备注",
    "ship_fee": "1.22",
    "other_fee": "5",
    "fee_part_type": 2,
    "inbound_time": "2022-12-02",
    "product_list": [
        {
            "sku": "MDR-ZX110",
            "good_num": 10,
            "bad_num": 0,
            "price": "8.66",
            "seller_id": 21,
            "fnsku": "FN206265A",
            "good_whb_code":"A-1-1",
            "bad_whb_code":"次品暂存",
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
request_id	请求链路id	是	[string]	5A1EBE67-7793-9E94-7790-AA457B81B3F2
response_time	响应时间	是	[string]	2022-06-27 11:54:31
data	响应数据	是	[object]	
data>>order_sn_arr	入库单号数组【兼容多个单号情况】	是	[array]	[IB221207001,IB221207002]
data>>order_sn	入库单号【多个单号情况下值为order_sn_arr里的第一个】	是	[string]	IB221207001
附加说明

针对采购单快捷入库【即入参order_sn传入采购单号】，以下字段会直接引用传入的采购单号对应的相关数据：

type 单据类型【自动根据采购入库类型变更】；
remark 单据备注；
fee_part_type 费用分配方式；
price 单价【单价 = 采购单单价 × 汇率】；
 上一章节
订单退款
下一章节 
查询入库单列表