查询报告详情 - Walmart Payment
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/fms/openapi/walmartPayment/queryPage	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15，上限200	否	[int]	15
store_id	店铺id	是	[array]	[110000000018008000,110000000018008001]
report_id	报告id
【时间范围 与 报告id 二选一必填】	否	[string]	MI7cc13c6a4f7d8769629e807e8263fc87
transaction_posted_timestamp_start	开始时间，格式：Y-m-d
【时间范围 与 报告id 二选一必填】	否	[string]	2023-05-01
transaction_posted_timestamp_end	结束时间，格式：Y-m-d
【时间范围 与 报告id 二选一必填】	否	[string]	2023-05-20
请求示例
{
    "length": 15,
    "offset": 0,
    "report_id": "MI7cc13c6a4f7d8769629e807e8263fc87",
    "store_id": [
        110000000018008000,
        110000000018008001
    ],
    "transaction_posted_timestamp_start": "2023-05-01",
    "transaction_posted_timestamp_end": "2023-05-20"
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
response_time	响应时间	是	[string]	2022-05-22 14:08:59
data	响应数据	是	[object]	
data>>total	总数	是	[int]	10
data>>records	详细信息列表	是	[array]	
data>>records>>amount	金额	是	[string]	5
data>>records>>amount_type	金额类型	是	[string]	Product Price
data>>records>>commission_rate	佣金率	是	[string]	20
data>>records>>commission_rate_str	佣金率%	是	[string]	20%
data>>records>>commission_rule	佣金规则	是	[string]	0.00000
data>>records>>currency	币种	是	[string]	USD
data>>records>>customer_order	客户订单号	是	[string]	2023_05_01_oPX5au
data>>records>>fulfillment_type	发货类型	是	[string]	Walmart-fulfilled(WFS)
data>>records>>msku	msku	是	[string]	walm_wjc9_sku7xP4a
data>>records>>partner_gtin	Partner GTIN	是	[string]	850030773067
data>>records>>partner_item_name	商品标题	是	[string]	Casabrews Espresso Machine
data>>records>>period_end_date	结束日期	是	[string]	2023-05-15
data>>records>>period_start_date	开始日期	是	[string]	2023-05-20
data>>records>>product_taxCode	产品税号	是	[string]	2042275
data>>records>>product_type	产品类型	是	[string]	Espresso Machines
data>>records>>purchase_order	平台单号	是	[string]	walmart-test051702-XmIVJU
data>>records>>ship_qty	数量	是	[int]	1
data>>records>>ship_to_city	收货城市	是	[string]	Miami Lakes
data>>records>>ship_to_state	收货州	是	[string]	FL
data>>records>>ship_to_zipcode	收货邮编	是	[string]	33016
data>>records>>shipping_method	运输方式	是	[string]	Marketplace expedited
data>>records>>store_id	店铺id	是	[string]	110000000018008002
data>>records>>store_name	店铺名称	是	[string]	walmart测试店铺2号
data>>records>>transaction_description	交易描述	是	[string]	Purchase
data>>records>>transaction_key	交易key	是	[string]	2023_05_01_001
data>>records>>transaction_posted_timestamp	交易日期	是	[string]	2023-05-15
data>>records>>transaction_reason_description	交易原因描述	是	[string]	transaction_reason_description
data>>records>>transaction_type	交易类型	是	[string]	Sale
返回成功示例
{
     "code": 0,
     "message": "success",
     "error_details": [],
     "request_id": "45fd55efd92b4aad8c97f5c0371ec063.1684735677529",
     "response_time": "2023-05-22 14:07:58",
     "data": {
          "total": "7",
          "records": [
               {
                    "period_end_date": null,
                    "period_start_date": null,
                    "transaction_key": "2023_05_01_001",
                    "currency": "USD",
                    "transaction_posted_timestamp": "2023-05-15",
                    "transaction_description": "Purchase",
                    "transaction_type": "Sale",
                    "customer_order": "2023_05_01_oPX5au",
                    "purchase_order": "walmart-test051702-XmIVJU",
                    "amount_type": "Product Price",
                    "amount": 20,
                    "ship_qty": "11",
                    "transaction_reason_description": "No Longer Wanted",
                    "commission_rate": 2,
                    "commission_rate_str": "2.00%",
                    "msku": "walm_wjc2_sku7xP4a",
                    "partner_item_name": "Casabrews Espresso Machine",
                    "partner_gtin": "850030773067",
                    "product_taxCode": "2042275",
                    "ship_to_city": "Miami Lakes",
                    "ship_to_state": "FL",
                    "ship_to_zipcode": "33016",
                    "product_type": "Espresso Machines",
                    "commission_rule": "BSE-65-3.0",
                    "shipping_method": "Marketplace expedited",
                    "fulfillment_type": "Walmart-fulfilled(WFS)",
                    "store_id": "110000000018008002",
                    "store_name": "walmart测试店铺"
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
    "error_details": ["length: length 最小值为15"],
    "request_id": "f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946",
    "response_time": "2023-05-22 14:08:59",
    "data": null
}
复制
错误
复制成功
 上一章节
查询可用报告列表 - Walmart Payment
下一章节 
查询平台仓发货单列表