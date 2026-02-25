查询利润报表-订单(旧版，将于04.30下线)

请尽快切换到：查询结算利润（利润报表）-订单

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatformFinance/profitReportPageList/order	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
platform_codes	平台code【目前支持的平台如下】
10005-速卖通
10008-沃尔玛
10003-ebay	否	[array]	["10005","10008"]
store_ids	店铺id	否	[array]	["110000000018003003", "110000000018003004"]
line_type_list	交易类型：
0 销售
2 退货
3 AE放款差异
4 退款
5 补发	否	[array]	[2,3]
currency_code	币种：
0 原币种
1 USD
2 EUR
3 DBP
4 CNY	是	[string]	0
time_type	时间查询类型：
0 结算日期
1 发货日期	是	[int]	0
start_time	开始日期【近三个月】，闭区间，格式：Y-m-d	是	[string]	2024-06-01
end_time	结束日期【近三个月】，闭区间，格式：Y-m-d	是	[string]	2024-06-30
search_field	搜索类型：
1 MSKU
2 SKU
3 品名
4 平台单号	否	[int]	1
search_value	搜索值	否	[string]	sku_aXArb
请求示例
{
    "offset": 0,
    "length": 20,
    "platform_codes": [
        "10005",
        "10008"
    ],
    "store_ids": [
        "110000000018003003",
        "110000000018003004"
    ],
    "line_type_list": [2,3],
    "currency_code": "0",
    "time_type": "0",
    "start_time": "2024-06-01",
    "end_time": "2024-06-30",
    "search_field": 1,
    "search_value": "sku_aXArb"
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
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>parent_id	父项报表id	是	[string]	1718952728827199481
data>>id	报表id	是	[string]	1718952728827199489
data>>pid	产品id	是	[string]	5752
data>>goods_url	图片地址	是	[string]	https://www.image.com/xxx
data>>order_id	平台单号	是	[string]	walmart-test0714-7GNnGJ
data>>global_order_no	系统单号	是	[string]	103334301947408384
data>>develop_user_id	产品开发人id	是	[int]	10325715
data>>cid	类型id	是	[int]	2
data>>bid	品牌id	是	[int]	1
data>>classification	类型	是	[string]	分类2
data>>brand	品牌	是	[string]	联想
data>>currency	币种	是	[string]	USD
data>>currency_coin	币种符号	是	[string]	$
data>>delivery_date	发货日期	是	[string]	2023-07-08
data>>sales_num	销量	是	[string]	1
data>>replenishment_num	补货量	是	[string]	2
data>>sales_amount	销售额	是	[string]	35.00
data>>buyer_freight	买家运费	是	[string]	12.00
data>>other_income	其他收入	是	[string]	0.00
data>>income_return	收入退款额	是	[string]	11.00
data>>cost_refund	费用退款额	是	[string]	10.00
data>>refund_num	退款量	是	[string]	2
data>>refund_rate	退款率	是	[string]	0.3143
data>>sales_return_num	退货量	是	[string]	1
data>>sales_return_rate	退货率	是	[string]	1
data>>platform_fee	平台费	是	[string]	10.00
data>>discount_fee	促销折扣费	是	[string]	0.00
data>>ad_fee	广告费	是	[string]	0.49
data>>adjustment_fee	调整费	是	[string]	1.48
data>>wfs_adjustment_fee	wfs调整费	是	[string]	1.48
data>>ebay_adjustment_fee	ebay调整费	是	[string]	0.00
data>>aliexpress_adjustment_fee	aliexpress调整费	是	[string]	0.00
data>>platform_transfer_fee	平台物流费	是	[string]	-22.00
data>>wfs_shipment_fee	wfs发货费	是	[string]	6.00
data>>wfs_return_transfer_fee	wfs退货运费	是	[string]	-12.00
data>>walmart_return_service_fee	walmart退货服务费	是	[string]	-16.00
data>>platform_storage_fee	平台仓储费	是	[string]	1.48
data>>wfs_storage_fee	wfs仓储费	是	[string]	0.49
data>>wfs_remove_fee	wfs移除费	是	[string]	0.99
data>>platform_other_fee	平台其他费	是	[string]	6.00
data>>other_fee	其他费	是	[string]	6.00
data>>ebay_publish_fee	ebay刊登费	是	[string]	0.00
data>>ebay_subscription_fee	ebay订阅费	是	[string]	0.00
data>>sales_tax	销售税	是	[string]	48.00
data>>goods_amount_after_tax	商品税后金额	是	[string]	0.00
data>>market_tax	市场税	是	[string]	32.00
data>>goods_other_fee	商品其他费用	是	[string]	0.00
data>>store_other_fee	店铺其他费用	是	[string]	0.00
data>>order_other_fee	订单其他费用	是	[string]	0.00
data>>purchase_cost	采购成本	是	[string]	-6.00
data>>sales_order_purchase_cost	售出订单采购成本	是	[string]	-9.00
data>>return_order_purchase_cost	退货订单采购成本	是	[string]	3.00
data>>firstlet_cost	头程成本	是	[string]	0.00
data>>sales_order_firstlet_cost	售出订单头程成本	是	[string]	-9.00
data>>return_order_firstlet_cost	退货订单头程成本	是	[string]	3.00
data>>tail_cost	尾程成本	是	[string]	0.00
data>>other_cost	其他成本	是	[string]	0.00
data>>sales_order_other_cost	售出订单其他成本	是	[string]	0.00
data>>return_order_other_cost	退货订单其他成本	是	[string]	0.00
data>>gross_profit	毛利润	是	[string]	139.45
data>>gross_profit_rate	毛利率	是	[string]	3.9843
data>>store_id_list	店铺id	是	[array]	["110000000018008003"]
data>>platform_code_list	平台code	是	[array]	["10008"]
data>>line_type_list	交易类型：
0 销售
2 退货
3 AE放款差异
4 退款
5 补发	是	[array]	[2,4]
data>>msku_list	MSKU	是	[array]	["walm_wjc2_skuaA"]
data>>local_name_list	本地SKU信息	是	[array]	
data>>local_name_list>>sku	SKU	是	[string]	"wjc2_sku"
data>>local_name_list>>product_name	品名	是	[string]	"wjc2_pm"
data>>create_date_list	结算日期	是	[array]	["2023-07-18"]
data>>child_item_list	报表子项数据【返回字段同父级】	是	[array]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7d1ba452a68c4724a9129175fa980028.118.16989107423739539",
    "response_time": "2023-11-02 15:39:02",
    "data": [
        {
            "parent_id": null,
            "id": "1719982434977710080",
            "pid": "5750",
            "goods_url": "https://image.xxx.com/d97094f5.jpeg",
            "order_id": "walmart-test0714-7GNnGJ",
            "global_order_no": "103334301947408384",
            "develop_user_id": 10062027,
            "cid": 1,
            "bid": 12,
            "classification": "分类1",
            "brand": "ANT",
            "currency": "USD",
            "currency_coin": "$",
            "delivery_date": "2023-07-08",
            "sales_num": "8",
            "replenishment_num": "0",
            "sales_amount": "86.00",
            "buyer_freight": "0.00",
            "other_income": "0.00",
            "income_return": "25.00",
            "cost_refund": "35.00",
            "refund_num": "12",
            "refund_rate": "0.2907",
            "sales_return_num": "11",
            "sales_return_rate": "1.375",
            "platform_fee": "35.00",
            "discount_fee": "0.00",
            "ad_fee": "33.96",
            "adjustment_fee": "101.89",
            "wfs_adjustment_fee": "101.89",
            "ebay_adjustment_fee": "0.00",
            "aliexpress_adjustment_fee": "0.00",
            "platform_transfer_fee": "-74.00",
            "wfs_shipment_fee": "-23.00",
            "wfs_return_transfer_fee": "-25.00",
            "walmart_return_service_fee": "-26.00",
            "platform_storage_fee": "101.89",
            "wfs_storage_fee": "33.96",
            "wfs_remove_fee": "67.93",
            "platform_other_fee": "1.00",
            "other_fee": "1.00",
            "ebay_publish_fee": "0.00",
            "ebay_subscription_fee": "0.00",
            "sales_tax": "118.00",
            "goods_amount_after_tax": "0.00",
            "market_tax": "84.00",
            "goods_other_fee": "0.00",
            "store_other_fee": "0.00",
            "order_other_fee": "0.00",
            "purchase_cost": "3.50",
            "sales_order_purchase_cost": "-30.00",
            "return_order_purchase_cost": "33.50",
            "firstlet_cost": "30.75",
            "sales_order_firstlet_cost": "-30.00",
            "return_order_firstlet_cost": "60.75",
            "tail_cost": "0.00",
            "other_cost": "0.25",
            "sales_order_other_cost": "0.00",
            "return_order_other_cost": "0.25",
            "gross_profit": "581.24",
            "gross_profit_rate": "6.7586",
            "store_id_list": [
                "110000000018008006"
            ],
            "platform_code_list": [
                "10008"
            ],
            "line_type_list": [
                0,
                2,
                4
            ],
            "msku_list": [
                "walm_wjc0_skuLZg2u",
                "walm_wjc8_skuLZg2u"
            ],
            "local_name_list": [
                {
                    "sku": "wjc0_sku",
                    "product_name": "wjc0_pm"
                },
                {
                    "sku": "wjc8_sku",
                    "product_name": "wjc8_pm"
                }
            ],
            "create_date_list": [
                "2023-07-08",
                "2023-07-15"
            ],
            "child_item_list": [
                {
                    "parent_id": "1719982434977710080",
                    "id": "1719982434990292998",
                    "pid": "5758",
                    "goods_url": "",
                    "order_id": "walmart-test0714-7GNnGJ",
                    "global_order_no": "103334301947408384",
                    "develop_user_id": 10062027,
                    "cid": 8,
                    "bid": 7,
                    "classification": "分类8",
                    "brand": "品牌7",
                    "currency": "USD",
                    "currency_coin": "$",
                    "delivery_date": "2023-07-08",
                    "sales_num": "0",
                    "replenishment_num": "0",
                    "sales_amount": "0.00",
                    "buyer_freight": "0.00",
                    "other_income": "0.00",
                    "income_return": "0.00",
                    "cost_refund": "0.00",
                    "refund_num": "0",
                    "refund_rate": "0.0",
                    "sales_return_num": "3",
                    "sales_return_rate": "0.0",
                    "platform_fee": "0.00",
                    "discount_fee": "0.00",
                    "ad_fee": "0.00",
                    "adjustment_fee": "0.00",
                    "wfs_adjustment_fee": "0.00",
                    "ebay_adjustment_fee": "0.00",
                    "aliexpress_adjustment_fee": "0.00",
                    "platform_transfer_fee": "0.00",
                    "wfs_shipment_fee": "0.00",
                    "wfs_return_transfer_fee": "0.00",
                    "walmart_return_service_fee": "0.00",
                    "platform_storage_fee": "0.00",
                    "wfs_storage_fee": "0.00",
                    "wfs_remove_fee": "0.00",
                    "platform_other_fee": "0.00",
                    "other_fee": "0.00",
                    "ebay_publish_fee": "0.00",
                    "ebay_subscription_fee": "0.00",
                    "sales_tax": "0.00",
                    "goods_amount_after_tax": "0.00",
                    "market_tax": "0.00",
                    "goods_other_fee": "0.00",
                    "store_other_fee": "0.00",
                    "order_other_fee": "0.00",
                    "purchase_cost": "13.50",
                    "sales_order_purchase_cost": "0.00",
                    "return_order_purchase_cost": "13.50",
                    "firstlet_cost": "0.45",
                    "sales_order_firstlet_cost": "0.00",
                    "return_order_firstlet_cost": "0.45",
                    "tail_cost": "0.00",
                    "other_cost": "0.15",
                    "sales_order_other_cost": "0.00",
                    "return_order_other_cost": "0.15",
                    "gross_profit": "14.10",
                    "gross_profit_rate": "0.0",
                    "store_id_list": [
                        "110000000018008006"
                    ],
                    "platform_code_list": [
                        "10008"
                    ],
                    "line_type_list": [
                        2
                    ],
                    "msku_list": [
                        "walm_wjc8_skuLZg2u"
                    ],
                    "local_name_list": [
                        {
                            "sku": "wjc8_sku",
                            "product_name": "wjc8_pm"
                        }
                    ],
                    "create_date_list": [
                        "2023-07-15"
                    ]
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询利润报表-店铺(旧版，将于04.30下线)
下一章节 
查询结算利润（利润报表）-msku