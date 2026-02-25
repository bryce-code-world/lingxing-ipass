查询WFS货件列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/warehouse/api/openApi/queryWFSCargoPage	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	否	[array]	["110418202566107648","11098710969430528"]
cargo_status_list	货件平台状态：
0 PENDING_SHIPMENT_DETAILS
1 AWAITING_DELIVERY
2 RECEIVING_IN_PROGRESS
3 CLOSED
4 CANCELLED	否	[array]	[1,2,3,4]
inbound_order_id	入库订单编号	否	[string]	lkxxxxxx555
start_time	开始时间【创建时间】，格式：Y-m-d，双闭区间	否	[string]	2023-07-15
end_time	结束时间【创建时间】，格式：Y-m-d，双闭区间	否	[string]	2023-07-16
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15，上限200	否	[int]	15
update_time_ge	货件更新开始时间，格式：yyyy-MM-dd HH:mm:SS	否	[string]	
update_time_le	货件更新结束时间，格式：yyyy-MM-dd HH:mm:SS	否	[string]	
cargo_update_time_ge	货件平台更新时间开始 格式:yyyy-MM-dd HH:mm:ss	否	[string]	
cargo_update_time_le	货件平台更新时间结束 格式:yyyy-MM-dd HH:mm:ss	否	[string]
请求示例
{
    "store_id": [
        "110418202566107648",
        "11098710969430528"
    ],
    "cargo_status_list": [1,2,3,4],
    "inbound_order_id": "lkxxxxxx555",
    "start_time": "2023-07-15",
    "end_time": "2023-07-16",
    "offset": 0,
    "length": 15
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
request_id	请求链路id	是	[string]	f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946
response_time	响应时间	是	[string]	2023-05-06 23:59:59
data	响应数据	是	[object]	
data>>total	总数	是	[string]	15
data>>records	列表数据	是	[array]	
data>>records>>id	WFS货件id	是	[string]	200498
data>>records>>cargo_code	货件单号	是	[string]	lry-shipment-31
data>>records>>cargo_create_date	货件创建时间	是	[string]	2023-05-06 23:59:59
data>>records>>cargo_good_list	商品信息	是	[array]	
data>>records>>cargo_good_list>>declare_num	货件初始申报量	是	[string]	10
data>>records>>cargo_good_list>>gtin	GTIN码	是	[string]	gtin
data>>records>>cargo_good_list>>msku	msku	是	[string]	mskuxxx
data>>records>>cargo_good_list>>product_name	品名	是	[string]	ueeshop222bt
data>>records>>cargo_good_list>>received_num	签收数量	是	[string]	10
data>>records>>cargo_good_list>>shipments_num	已发货数量	是	[string]	10
data>>records>>cargo_good_list>>sku	sku	是	[string]	sku
data>>records>>cargo_good_list>>dameged_qty	损坏数量	是	[string]	
data>>records>>cargo_status	平台货件状态	是	[string]	PENDING_SHIPMENT_DETAILS
data>>records>>cargo_sync_status	货件本地状态	是	[string]	已申报
data>>records>>country_name	国家	是	[string]	美国
data>>records>>creator	创建者	是	[string]	王xxx
data>>records>>distribution_addresses	配送地址(街道,城市，省/区，国家，国家编码,邮编)拼接	是	[string]	xxxxxxx
data>>records>>in_bound_order_id	入库订单编号	是	[string]	lry-inbound-31
data>>records>>logistics_code	物流编号	是	[string]	LRY03
data>>records>>return_addresses	退货地址（街道,城市，省/区，国家，国家编码)拼接）	是	[string]	xxxxxxx
data>>records>>store_id	店铺id	是	[string]	110000000020008001
data>>records>>store_name	店铺名称	是	[string]	测试店铺
data>>records>>shipping_list_codes	发货单号	是	[array]	["SP202307250010002","SP202307250010001"]
data>>records>>update_date	更新时间	是	[string]	2023-05-06 23:59:59
data>>records>>status	货件状态：
0 PENDING_SHIPMENT_DETAILS (等待发货详情) ， 表示系统正在等待发货的详细信息
1 AWAITING_DELIVERY (等待送达) ，表示货物已发出，正在等待送达
2 RECEIVING_IN_PROGRESS (接收中) ， 表示货物正在被接收处理中
3 CLOSED (已关闭) ， 表示物流过程已完成并关闭
4 CANCELLED (已取消) ， 表示物流过程已被取消	是	[int]	
data>>records>>status_name	货件状态说明	是	[string]	
data>>records>>to_pending_time	更新为Pending Shipment Details状态的时间戳	是	[string]	
data>>records>>to_await_time	更新为Awaiting Delivery状态的时间戳	是	[string]	
data>>records>>to_receive_time	更新Receiving in Progress状态的时间戳	是	[string]	
data>>records>>to_closed_time	更新为CLOSED状态的时间戳	是	[string]	
data>>records>>to_cancelled_time	更新为Cancelled状态的时间戳	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "c4db3691436d4acbbda0f8b1dc53fe97.1689845618385",
    "response_time": "2023-07-20 17:33:39",
    "data": {
        "total": "1",
        "records": [
            {
                "id": "200498",
                "in_bound_order_id": "test0508001-inbound-1",
                "store_id": "110000000018008001",
                "store_name": "测试店铺",
                "country_name": "美国",
                "creator": "",
                "shipping_list_codes": [
                    "SP202307270010196"
                ],
                "logistics_code": "LRY03",
                "cargo_create_date": "2023-03-08 03:02:30",
                "cargo_code": "test0508001-shipment-1",
                "update_date": "2023-05-08 11:20:52",
                "cargo_status": "PENDING_SHIPMENT_DETAILS",
                "cargo_sync_status": "已申报",
                "distribution_addresses": "3 Sorbello Way,Pedricktown1,NJ9,美国,US,08067",
                "return_addresses": "860 W California Ave,Sunnyvale,94086,加拿大,CA",
                "cargo_good_list": [
                    {
                        "msku": "test1_sku1CIlPbdpXO1",
                        "sku": "ueeshop1111",
                        "product_name": "ueeshop1111",
                        "gtin": "006344827368521",
                        "shipments_num": "0",
                        "received_num": "0",
                        "declare_num": "10",
                        "dameged_qty":"0"
                    }
                ]
            }
        ]
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "参数检验不通过",
    "error_details": ["length: length 最大值200"],
    "request_id": "f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946",
    "response_time": "2023-05-22 14:08:59",
    "data": null
}
复制
错误
复制成功
 上一章节
修改平台仓发货单备注
下一章节 
查询WFS库存列表