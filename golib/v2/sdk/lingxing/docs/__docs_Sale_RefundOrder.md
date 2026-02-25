订单退款
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/salesOrder/refundOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id	是	[number]	17
amazonOrderId	亚马逊订单ID	是	[string]	111-1111745-7011114
purchaseDateLocal	订购时间	是	[string]	2025-10-24 20:17:13
data		否	[array]	
data>>orderItemId	商品行id	是	[string]	54278911110546
data>>asin	asin	是	[string]	B01111RQCF
data>>sellerSku	msku	是	[string]	171111023
data>>title	商品名称	是	[string]	xxxx
data>>quantityOrdered	下单数量	是	[number]	1
data>>quantityShipped	到货数量	是	[number]	1
data>>reason	退款原因
CustomerReturn
GeneralAdjustment	是	[string]	GeneralAdjustment
data>>asinUrl	asinUrl	否	[string]	https://www.amazon.com/dp/B111111
data>>smallImageUrl	商品图片	否	[string]	
data>>unitPriceIcon	商品单价货币符号	否	[string]	$
data>>unitPriceAmount	单价	否	[number]	192.78
data>>itemList	退款费用项目列表	否	[array]	
data>>itemList>>type	费用类型	是	[string]	Principal
data>>itemList>>name	费用名称	否	[string]	商品
data>>itemList>>currencyCode	货币编码	是	[string]	USD
data>>itemList>>icon	货币符号	否	[string]	$
data>>itemList>>amount	金额	否	[number]	192.78
data>>itemList>>refundedPrice	已申请退款金额	否	[string]	0.00
data>>itemList>>returnAmount	本次申请退款金额	是	[string]	1999
data>>itemRefundTotal		是	[string]	208.68
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/salesOrder/refundOrder?access_token=value&sign=value&timestamp=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "sid": 17,
    "amazonOrderId": "111-4279745-7077814",
    "purchaseDateLocal": "2025-10-24 20:17:13",
    "data": [{
        "orderItemId": "542711111546",
        "asin": "B081111CF",
        "sellerSku": "1704111123",
        "title": "xxxxxxxxxxx",
        "quantityOrdered": 1,
        "quantityShipped": 1,
        "reason": "Other",
        "asinUrl": "https://www.amazon.com/dp/B11111",
        "smallImageUrl": "",
        "unitPriceIcon": "$",
        "unitPriceAmount": 192.78,
        "itemList": [{
            "type": "Principal",
            "name": "商品",
            "currencyCode": "USD",
            "icon": "$",
            "amount": 192.78,
            "refundedPrice": "0.00",
            "returnAmount": "1999"
        }],
        "itemRefundTotal": "208.68"
    }]
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[number]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	1ad812bb5273462c9d2077107ff7cd99.1757412110374
response_time	响应时间	是	[string]	2025-09-09 18:01:52
data		否	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "1ad812bb5273462c9d2077107ff7cd99.1757412110374",
    "response_time": "2025-09-09 18:01:52",
    "data": null
}
复制
错误
复制成功
附加说明

1）：亚马逊API限制：如同一订单存在多个商品，退款无法仅退某个商品，如需按比例退款，请自行分配每个商品退款比例。 2）：type枚举类如下： Principal 商品 Tax 商品税 Shipping 配送费 ShippingTax 配送税 GiftWrap 包装费 GiftWrapTax 包装税 COD 货到付款 ReturnShipping 退货优惠 RestockingFee 重新入库费 RestockingFeeTax 重新入库税

 上一章节
待收货退货单快捷入库
下一章节 
添加入库单