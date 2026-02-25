备货单分批收货
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/batchesReceipt	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	OWS240724004
product_list	产品信息	是	[array]	
product_list>>product_id	本地产品id	是	[int]	33055
product_list>>current_receive_num	收货数量	是	[int]	10
product_list>>sid	店铺id【备货单中对应有值则必填】 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[int]	0
product_list>>fnsku	fnsku【备货单中对应有值则必填】	否	[string]	FN7DB7A81
product_list>>product_code	第三方仓sku【备货单中对应有值则必填】	否	[string]	00015214
请求示例
{
    "overseas_order_no": "OWS240724004",
    "product_list": [
        {
            "product_id": 33055,
            "current_receive_num": 10,
            "sid": "0",
            "fnsku": "FN7DB7A81",
            "product_code": "00015214"
        }
    ]
}
复制
错误
复制成功
返回信息

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	 
 上一章节
海外仓备货单发货
下一章节 
备货单结束到货