创建订单

创建多平台店铺手工单，包含自定义平台订单，对应系统中订单类型为“手工订单”类型

API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/create	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
platform_code	平台code	是	[int]	10003
store_id	店铺id	是	[string]	11058263263899648
orders	订单列表	是	[array]	
orders>>platform_order_no	平台单号【同一店铺不支持重复】	是	[string]	10379511104466944
orders>>site_code	站点	否	[string]	10001-UK
orders>>buyer_note	买家留言	否	[string]	
orders>>receiver_country_code	国家/地区二字简码	是	[string]	CN
orders>>customer_shipping_amount	客付运费【币种默认为店铺币种】	否	[number]	10.22
orders>>customer_tax_amount	客付税费【币种默认为店铺币种】	否	[number]	1.22
orders>>order_total_amount	订单金额【币种默认为店铺币种】	否	[number]	20.20
orders>>wid	仓库id ，【当填写了该字段，则logistics_type_id 必填】	否	[string]	1
orders>>global_purchase_time	订购时间(秒级时间戳)	否	[number]	
orders>>global_payment_time	付款时间(秒级时间戳)	否	[number]	
orders>>remark	客服备注	否	[string]	
orders>>logistics_type_id	物流方式id， 【当填写了该字段，则wid必填】	否	[string]	1
orders>>receiver_name	收件人	是	[string]	1
orders>>address_type	地址类型
1:住宅地址”
2:“商业地址”	否	[int]	1
orders>>buyer_choose_express	客选物流	否	[string]	
orders>>receiver_company_name	公司名	否	[string]	
orders>>city	城市	是	[string]	1
orders>>address_line1	地址1	是	[string]	1
orders>>amount_currency	币种，默认与店铺的币种一致，ISO 4217标准码	否	[string]	USD
orders>>items	订单产品信息	是	[array]	
orders>>items>>sku	本地SKU（SKU和MSKU必填其中一个）	否	[string]	xhp004
orders>>items>>quantity	数量	是	[int]	10
orders>>items>>unit_price	单价【币种默认为店铺币种】	是	[number]	10
orders>>items>>msku	MSKU（SKU和MSKU必填其中一个）	否	[string]	
orders>>items>>stock_deduction_type	库存扣减类型 1表示“空”， 2表示“SKU+订单店铺”	否	[String]	1
请求示例
{
    "platform_code": 10009,
    "store_id": "110283406273996288",
    "orders": [
        {
            "buyer_note": "xxx",
            "customer_shipping_amount": "20.00",
            "customer_tax_amount": "2.00",
            "order_total_amount": "20.80",
            "platform_order_no": "10379511104466944",
            "receiver_country_code": "CN",
            "wid": "1",
            "logistics_type_id": "1",
            "items": [
                {
                    "sku": "xhp004",
                    "quantity": 1,
                    "unit_price": "10.00",
                    "stock_deduction_type": 1
                }
            ]
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
message	提示信息	是	[string]	
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
data	响应数据	是	[array]	
data>>error_details	错误信息	是	[array]	
data>>error_details>>error_message	失败原因	是	[string]	
data>>error_details>>platform_order_no	平台单号	是	[string]	
data>>success_details	成功详情	是	[array]	
data>>success_details>>global_order_no	系统单号	是	[int]	
data>>success_details>>platform_order_no	平台单号	是	[string]	10379511104466944
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "2ecb8a09-275e-4773-9399-ebfbb17c1804.1689906739215",
    "response_time": "2023-07-21 10:32:19",
    "data": {
        "error_details": [],
        "success_details": [
            {
                "platform_order_no": "263407812560417017",
                "global_order_no": "103336708973801472"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
审核发货
下一章节 
合并订单