查询备货单收货记录
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/getReceiveGoodRecords	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单单号【不支持批量】	否	[string]	OWS230314002
start_date	收货开始时间，闭区间，格式：Y-m-d	否	[string]	2023-10-12
end_date	收货结束时间，开区间，格式：Y-m-d	否	[string]	2023-12-11
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认500	否	[int]	500
请求示例
{
    "overseas_order_no": "OWS240724004",
    "start_date": "2021-07-01",
    "end_date": "2024-07-30",
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	结果	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	821FA057-C388-70AC-3235-2EC7C818F342
response_time	响应时间	是	[string]	2023-03-28 15:38:12
data	响应数据	是	[array]	
data>>woop_id	订单商品id	是	[int]	26206
data>>overseas_order_no	备货单单号	是	[string]	OWS230314002
data>>inbound_order_no	入库单号	是	[string]	IB230314008
data>>current_receive_num	本次收货数量	是	[int]	1
data>>receive_num	应收数量	是	[int]	1
data>>update_user	操作人	是	[string]	林少@testSalMan
data>>update_time	操作时间	是	[string]	2023-03-14 11:08:22
data>>sid	店铺id	是	[int]	28
data>>seller_name	店铺名称	是	[string]	jzy-jp
data>>product_id	产品id	是	[int]	2124018
data>>fnsku	FNSKU	是	[string]	X000X5138P
data>>sku	本地商品sku	是	[string]	组合1
data>>in_storage_bins	入库仓位	是	[string]	
data>>real_receive_at	实际收货时间	是	[string]	2023-03-14 00:00:00
data>>product_name	本地商品名称	是	[string]	组合1
data>>twp_sku	第三方海外仓商品sku	是	[string]	
data>>twp_name	第三方海外仓商品名称	是	[string]	
data>>match_num	第三方海外仓商品对应的本地商品配对数量	是	[int]	1
total	总数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "821FA057-C388-70AC-3235-2EC7C818F342",
    "response_time": "2023-03-28 15:38:12",
    "data": [
        {
            "woop_id": 26206,
            "overseas_order_no": "OWS230314002",
            "inbound_order_no": "IB230314008",
            "current_receive_num": 1,
            "receive_num": 1,
            "update_user": "林少@testSalMan",
            "update_time": "2023-03-14 11:08:22",
            "sid": 28,
            "seller_name": "jzy-jp",
            "product_id": 2124018,
            "fnsku": "X000X5138P",
            "sku": "组合1",
            "in_storage_bins": "",
            "real_receive_at": "2023-03-14 00:00:00",
            "product_name": "组合1",
            "twp_sku": "",
            "twp_name": "",
            "match_num": 1,
        }
    ],
    "total": 1
}
复制
错误
复制成功
附加说明

1.当备货单单号、收货开始时间、收货结束时间均有值，则按照单号+时间查询
2.当备货单单号、收货开始时间、收货结束时间均无值，则按照查询时间前一个月查询
3.当备货单单号无值，收货开始时间、收货结束时间均有值，则按照时间查询
4.当备货单单号、收货开始时间无值，收货结束时间有值，则按照收货结束时间前一个月查询
5.当备货单单号、收货结束时间无值，收货开始时间有值，则按照收货开始时间后一个月查询
6.当备货单单号有值，收货开始时间、收货结束时间均无值，则按照单号查询
7.当备货单单号有值，收货开始时间、收货结束时间任一无值，则按照单号+时间（前后）一个月查询

 上一章节
上传备货单装箱信息
下一章节 
海外仓备货单发货