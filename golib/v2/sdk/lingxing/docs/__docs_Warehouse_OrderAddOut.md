添加出库单

支持添加本地仓库出库单，推送成功后扣减出库仓库库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/storage/orderAddOut	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	自定义仓库ID，wid和sys_wid其中一项必填，都填则优先wid	否	[string]	custom-A1
sys_wid	系统仓库ID，sys_wid和wid其中一项必填，都填则优先wid	是	[int]	1
type	单据类型：
11 其他出库
12 FBA出库
14 退货出库	是	[int]	11
status	新建单据状态：10：待提交，40：已完成【默认值】	否	[int]	40
sys_supplier_id	系统客户供应商ID（退货出库：客户供应商ID, sys_supplier_id和supplier_id其中一个必填，都填则取supplier_id）	否	[int]	321
supplier_id	客户供应商ID（退货出库：客户供应商ID, sys_supplier_id和supplier_id其中一个必填，都填则取supplier_id）	否	[string]	AVS2132
idempotent_code	客户参考号, 该字段校验唯一不可重复	否	[string]	R0011231
remark	单据备注	否	[string]	
return_price	退货费（退货出库）	否	[number]	6.66
other_fee	其它费用（退货出库）	否	[number]	8.88
sys_to_wid	系统客户目的仓库ID（非退货出库）	否	[int]	56
to_wid	客户目的仓库ID（非退货出库）	否	[string]	WA323
outbound_time	自定义出库时间，格式：Y-m-d	否	[string]	2022-12-30
product_list	产品明细	是	[array]	
product_list>>sku	sku	是	[string]	kk-1232
product_list>>good_num	良品数量	是	[int]	3
product_list>>bad_num	次品数量	是	[int]	1
product_list>>seller_id	店铺id【没有店铺时传0】
亚马逊店铺对应查询亚马逊店铺信息字段【sid】
多平台店铺对应查询多平台店铺信息字段【store_id】	是	[int]	21
product_list>>fnsku	fnsku【存在fnsku时店铺id必填，没有时传空字符串】	是	[string]	FN206265A
请求示例
{
    "wid": "custom-A1",
    "sys_wid": 1,
    "type": 11,
    "sys_supplier_id": 321,
    "supplier_id": "AVS2132",
    "idempotent_code":"R0011231",
    "remark": "",
    "return_price": 6.66,
    "other_fee": 8.88,
    "sys_to_wid": 56,
    "to_wid": "WA323",
    "outbound_time": "2022-12-30",
    "product_list": [{
        "sku": "kk-1232",
        "good_num": 3,
        "bad_num": 1,
        "seller_id": 249,
        "fnsku": "FN4658CDA"
    }]
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
data>>order_sn	出库单号	是	[string]	IB221207153
 上一章节
撤销入库单
下一章节 
查询出库单列表