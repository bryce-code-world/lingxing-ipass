调拨单分批收货

支持将系统待收货状态的调拨单部分收货

基本信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/partlyReceiveAllocationOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	调拨单单号	是	[string]	TF240727001
product_list	本次收货的调拨单单据明细	是	[array]	
product_list>>product_id	产品id	是	[int]	10089
product_list>>seller_id	店铺id，不传或者传空则默认为0 ,对应查询亚马逊店铺列表接口对应字段【seller_id】	否	[string]	0
product_list>>fnsku	fnsku，不传则默认为空	否	[string]	abcdefg
product_list>>received_good_num	本次收货可用收货量，和本次次品收货次品收货量其中一项必填且大于0	否	[int]	0
product_list>>received_bad_num	本次收货次品收货量，和本次可用收货次品收货量其中一项必填且大于0	否	[int]	0
请求示例
{
    "order_sn": "TF240727001",
    "product_list": [
        {
            "product_id": "10089",
            "seller_id": "0",
            "fnsku": "abcdefg",
            "received_good_num": 0,
            "received_bad_num": 0
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
message	提示消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C70FD2A7-4E7D-6B32-7D72-FC42587CE97E
response_time	响应时间	是	[string]	2022-07-29 14:33:42
data	返回数据	是	[array]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C70FD2A7-4E7D-6B32-7D72-FC42587CE97E",
    "response_time": "2022-07-29 14:33:42",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
调拨单全部收货
下一章节 
调拨单结束到货