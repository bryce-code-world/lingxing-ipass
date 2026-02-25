查询平台订单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepfPlatformOrder/open-api/newPlatformOrder/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
dateType	时间类型 0.平台数据变动时间 1.订购时间 2.订购时间-北京 3.支付时间 4.支付时间-北京 5.发货时间 6.发货时间-北京	是	[int]	
deliveryTypeList	发货类型: 0-自发货 1-平台发货 2-部分自发货	否	[array]	
pageNum	查询起始位置	否	[int]	
pageSize	分页大小	否	[int]	
platformCodeList	平台CODE，目前仅支持 TikTok、TEMU 半托管、Line Shopping、Lazada、Shopee、Shopify、Walmart、Wayfair 平台	否	[array]	
searchMultiValue	多个精确搜索查询值	否	[array]	
searchSingleValue	单个模糊搜索查询值	否	[string]	
searchType	搜索查询类型：1：sku，2：品名，3：msku 4.商品id 5.平台单号 6.参考号 7.商品标题	否	[int]	
siteCodeList	站点列表	否	[array]	
sortField	排序字段列表字段支持：purchaseTime，paymentTime，platformOrderModifiedTime,deliveryTime	否	[string]	
sortType	升降序 asc desc	否	[string]	
startDate	开始时间，闭区间 格式：2025-10-22 00:00:01	是	[string]	
endDate	结束时间，闭区间 格式：2025-10-22 20:00:01	是	[string]	
statusList	平台单状态的编码 平台订单状态枚举	否	[array]	
storeIdList	店铺唯一标识	否	[array]
请求示例
{
  "dateType": 1,
  "deliveryTypeList": [0, 1],
  "pageNum": 1,
  "pageSize": 50,
  "platformCodeList": ["10011", "10024"],
  "searchMultiValue": ["SKU123", "SKU456"],
  "searchSingleValue": "商品关键词",
  "searchType": 1,
  "siteCodeList": ["10008-US"],
  "sortField": "purchaseTime",
  "sortType": "desc",
  "startDate": "2025-01-01 01:01:01",
  "endDate": "2025-01-31 01:01:01",
  "statusList": ["10002#202"],
  "storeIdList": ["110000000018008001"]
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
request_id	请求链路id	是	[string]	1979122130ba45e984cbc9b2e571d86f.1690513512438
response_time	响应时间	是	[string]	2023-07-28 11:05:12
data	响应数据	是	[object]	
data>>current	当前分页	否	[long]	
data>>list	订单列表	否	[array]	
data>>list>>address	地址对象	否	[object]	
data>>list>>address>>addressLine1	详细地址1	否	[string]	
data>>list>>address>>addressLine2	详细地址2	否	[string]	
data>>list>>address>>addressLine3	详细地址3	否	[string]	
data>>list>>address>>addressType	地址类型 1商业地址、2住宅	否	[int]	
data>>list>>address>>addressTypeName	地址类型 商业地址、住宅	否	[string]	
data>>list>>address>>companyName	公司名	否	[string]	
data>>list>>address>>detailAddress	详细地址	否	[string]	
data>>list>>address>>houseNumber	门牌号	否	[string]	
data>>list>>address>>recipientCity	收件人城市	否	[string]	
data>>list>>address>>recipientCountry	收件人国家	否	[string]	
data>>list>>address>>recipientDistrict	收件人区县	否	[string]	
data>>list>>address>>recipientName	收件人姓名	否	[string]	
data>>list>>address>>recipientPhone	收件人电话	否	[string]	
data>>list>>address>>recipientPostalCode	收件邮编	否	[string]	
data>>list>>address>>recipientStateOrRegion	收件人省、州	否	[string]	
data>>list>>erpInfo	系统信息	否	[array]	
data>>list>>erpInfo>>customerShippingName	发货物流	否	[string]	
data>>list>>erpInfo>>globalDeliveryTime	出库时间也是发货时间	否	[string]	
data>>list>>erpInfo>>globalOrderNo	系统单号	否	[string]	
data>>list>>erpInfo>>globalReissueTime	系统创建时间	否	[string]	
data>>list>>erpInfo>>globalReviewTime	审核时间	否	[string]	
data>>list>>erpInfo>>itemInfos	商品子行信息	否	[array]	
data>>list>>erpInfo>>itemInfos>>localProductName	本地产品品名	否	[string]	
data>>list>>erpInfo>>itemInfos>>localSku	sku	否	[string]	
data>>list>>erpInfo>>itemInfos>>msku	msku	否	[string]	
data>>list>>erpInfo>>itemInfos>>picUrl	sku图片url	否	[string]	
data>>list>>erpInfo>>itemInfos>>productNo	产品编号	否	[string]	
data>>list>>erpInfo>>itemInfos>>quantity	下单数量	否	[string]	
data>>list>>erpInfo>>itemInfos>>title	title	否	[string]	
data>>list>>erpInfo>>logisticsFreight	物流运费(直接展示，自带币种)	否	[string]	
data>>list>>erpInfo>>logisticsTime	物流下单时间	否	[string]	
data>>list>>erpInfo>>outboundCost	出库成本(直接展示，自带币种)	否	[string]	
data>>list>>erpInfo>>statusName	状态	否	[string]	
data>>list>>erpInfo>>warehouseName	发货仓库名称	否	[string]	
data>>list>>itemList	订单子行	否	[array]	
data>>list>>itemList>>bid	品牌ID	否	[long]	
data>>list>>itemList>>cid	分类ID	否	[long]	
data>>list>>itemList>>currency	币种	否	[string]	USD
data>>list>>itemList>>customerShipping	客选物流	否	[string]	
data>>list>>itemList>>customizedText	定制文本	否	[string]	
data>>list>>itemList>>deliveryType	发货类型（0-自发货 1-平台发货）	否	[int]	
data>>list>>itemList>>id	唯一键分布式ID	否	[long]	
data>>list>>itemList>>imgUrl	平台图片	否	[string]	
data>>list>>itemList>>localSku	本地SKU	否	[string]	
data>>list>>itemList>>model	产品型号	否	[string]	
data>>list>>itemList>>msku	平台msku	否	[string]	
data>>list>>itemList>>mskuId	平台mskuId	否	[string]	
data>>list>>itemList>>mskuRelationKey	配对相关：平台SKU关系键	否	[string]	
data>>list>>itemList>>orderItemNo	平台商品行ID	否	[string]	
data>>list>>itemList>>pid	本地产品ID（SKU级别ID）	否	[long]	
data>>list>>itemList>>pname	品名	否	[string]	
data>>list>>itemList>>productDeveloperUid	开发人UID	否	[long]	
data>>list>>itemList>>productNo	平台商品ID	否	[string]	
data>>list>>itemList>>productStatusName	平台商品状态名称	否	[string]	
data>>list>>itemList>>productUrl	商品链接	否	[string]	
data>>list>>itemList>>psId	本地SPU ID	否	[long]	
data>>list>>itemList>>quantity	数量	否	[long]	1
data>>list>>itemList>>remark	商品备注	否	[string]	
data>>list>>itemList>>spu	本地SPU	否	[string]	
data>>list>>itemList>>spuName	款名	否	[string]	
data>>list>>itemList>>title	标题	否	[string]	无标题
data>>list>>itemList>>transaction	交易信息汇总	否	[object]	
data>>list>>itemList>>transaction>>giftWrappingFee	礼品包装费	否	[object]	
data>>list>>itemList>>transaction>>giftWrappingFee>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>giftWrappingFee>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>giftWrappingFee>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>grossMargin	毛利率 返回0.01，前端展示1%	否	[number]	
data>>list>>itemList>>transaction>>grossProfit	毛利润	否	[object]	
data>>list>>itemList>>transaction>>grossProfit>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>grossProfit>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>grossProfit>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>logisticsFreight	物流运费	否	[object]	
data>>list>>itemList>>transaction>>logisticsFreight>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>logisticsFreight>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>logisticsFreight>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>outboundCost	出库成本	否	[object]	
data>>list>>itemList>>transaction>>outboundCost>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>outboundCost>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>outboundCost>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>platformLogisticsFree	平台物流费	否	[object]	
data>>list>>itemList>>transaction>>platformLogisticsFree>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>platformLogisticsFree>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>platformLogisticsFree>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>platformOtherFee	平台其他费	否	[object]	
data>>list>>itemList>>transaction>>platformOtherFee>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>platformOtherFee>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>platformOtherFee>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>platformSubsidy	平台补贴	否	[object]	
data>>list>>itemList>>transaction>>platformSubsidy>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>platformSubsidy>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>platformSubsidy>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>productDiscount	商品折扣	否	[object]	
data>>list>>itemList>>transaction>>productDiscount>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>productDiscount>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>productDiscount>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>productTax	商品税	否	[object]	
data>>list>>itemList>>transaction>>productTax>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>productTax>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>productTax>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>productTotalAmount	商品总额	否	[object]	
data>>list>>itemList>>transaction>>productTotalAmount>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>productTotalAmount>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>productTotalAmount>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>productTotalAmountTaxIncluded	商品总额是否含税	否	[boolean]	
data>>list>>itemList>>transaction>>realPayAmount	实付金额	否	[object]	
data>>list>>itemList>>transaction>>realPayAmount>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>realPayAmount>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>realPayAmount>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>salesRevenue	销售收益	否	[object]	
data>>list>>itemList>>transaction>>salesRevenue>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>salesRevenue>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>salesRevenue>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>shippingFee	客付运费	否	[object]	
data>>list>>itemList>>transaction>>shippingFee>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>shippingFee>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>shippingFee>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>shippingFeeDiscount	运费折扣	否	[object]	
data>>list>>itemList>>transaction>>shippingFeeDiscount>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>shippingFeeDiscount>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>shippingFeeDiscount>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>shippingFeeTaxIncluded	客付运费是否含税	否	[boolean]	
data>>list>>itemList>>transaction>>shippingTax	运费税	否	[object]	
data>>list>>itemList>>transaction>>shippingTax>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>shippingTax>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>shippingTax>>value	金额	否	[number]	
data>>list>>itemList>>transaction>>tax	税费	否	[object]	
data>>list>>itemList>>transaction>>tax>>currency	币种	否	[string]	
data>>list>>itemList>>transaction>>tax>>icon	币种符号	否	[string]	
data>>list>>itemList>>transaction>>tax>>value	金额	否	[number]	
data>>list>>itemList>>unitPriceAmount	单价	否	[number]	
data>>list>>itemList>>variantAttr	变体属性	否	[string]	
data>>list>>order	订单基本信息	否	[object]	
data>>list>>order>>buyerEmail	买家邮箱	否	[string]	
data>>list>>order>>buyerMessage	买家留言	否	[string]	
data>>list>>order>>buyerName	买家姓名/ID	否	[string]	
data>>list>>order>>companyId	企业唯一标识	否	[string]	
data>>list>>order>>deliveryStatusName	平台单发货状态状态	否	[string]	
data>>list>>order>>deliveryTime	发货时间	否	[string]	
data>>list>>order>>deliveryTimeBj	发货时间北京	否	[string]	
data>>list>>order>>deliveryType	发货类型（0-自发货 1-平台发货 2-部分自发货）	否	[int]	
data>>list>>order>>expectedDeliveryTime	指定配送时间	否	[string]	
data>>list>>order>>id	分布式ID	否	[long]	
data>>list>>order>>latestDeliveryTime	发货时限	否	[string]	
data>>list>>order>>latestShipTime	最晚配送时间	否	[string]	
data>>list>>order>>orderStatusName	平台单订单状态	否	[string]	
data>>list>>order>>paymentStatusName	平台单付款状态	否	[string]	
data>>list>>order>>paymentTime	支付时间	否	[string]	
data>>list>>order>>paymentTimeBj	支付时间（北京时间）	否	[string]	
data>>list>>order>>platformCode	平台CODE	否	[string]	
data>>list>>order>>platformName	平台名称	否	[string]	
data>>list>>order>>platformOrderModifiedTime	平台订单更新时间（平台订单有变动才变动）	否	[datetime]	
data>>list>>order>>platformOrderName	平台订单号(展示字段)	否	[string]	
data>>list>>order>>platformOrderNo	平台订单号(唯一键)	否	[string]	
data>>list>>order>>purchaseTime	订购时间	否	[string]	
data>>list>>order>>purchaseTimeBj	订购时间（北京时间）	否	[string]	
data>>list>>order>>referenceNo	参考号	否	[string]	
data>>list>>order>>remark	备注（商家自定义备注）	否	[string]	
data>>list>>order>>siteCode	平台+站点的唯一标识	否	[string]	
data>>list>>order>>storeId	店铺唯一标识	否	[string]	
data>>list>>order>>storeName	店铺名称	否	[string]	
data>>list>>order>>uniqueOrderNo	唯一订单(store_id#订单id)	否	[string]	
data>>list>>order>>sourceName	shopify订单的销售渠道	否	[string]	
data>>list>>packages	包裹信息	否	[array]	
data>>list>>packages>>deliveryTimeBj	发货时间（北京时间）	否	[string]	
data>>list>>packages>>itemList	商品信息	否	[array]	
data>>list>>packages>>itemList>>imgUrl	平台图片	否	[string]	
data>>list>>packages>>itemList>>msku	平台msku	否	[string]	
data>>list>>packages>>itemList>>mskuId	平台mskuId	否	[string]	
data>>list>>packages>>itemList>>orderItemNo	平台商品行ID	否	[string]	
data>>list>>packages>>itemList>>productNo	平台商品ID	否	[string]	
data>>list>>packages>>itemList>>quantity	数量	否	[int]	1
data>>list>>packages>>itemList>>title	标题	否	[string]	无标题
data>>list>>packages>>itemList>>variantAttr	变体属性	否	[string]	
data>>list>>packages>>packageId	包裹id	否	[string]	
data>>list>>packages>>shippingProviderName	发货物流	否	[string]	
data>>list>>packages>>trackingNumber	追踪单号/物流单号	否	[string]	
data>>list>>tagInfos	标签信息说明	否	[array]	
data>>list>>tagInfos>>id	标签ID	否	[long]	
data>>list>>tagInfos>>tagColor	标签颜色	否	[string]	
data>>list>>tagInfos>>tagName	标签名称	否	[string]	
data>>list>>tagInfos>>tagType	标签类型	否	[string]	
data>>list>>tagInfos>>tagTypeId	标签唯一键	否	[string]	
data>>list>>transaction	交易信息汇总	否	[object]	
data>>list>>transaction>>giftWrappingFee	礼品包装费	否	[object]	
data>>list>>transaction>>giftWrappingFee>>currency	币种	否	[string]	
data>>list>>transaction>>giftWrappingFee>>icon	币种符号	否	[string]	
data>>list>>transaction>>giftWrappingFee>>value	金额	否	[number]	
data>>list>>transaction>>grossMargin	毛利率 返回0.01，前端展示1%	否	[number]	
data>>list>>transaction>>grossProfit	毛利润	否	[object]	
data>>list>>transaction>>grossProfit>>currency	币种	否	[string]	
data>>list>>transaction>>grossProfit>>icon	币种符号	否	[string]	
data>>list>>transaction>>grossProfit>>value	金额	否	[number]	
data>>list>>transaction>>logisticsFreight	物流运费	否	[object]	
data>>list>>transaction>>logisticsFreight>>currency	币种	否	[string]	
data>>list>>transaction>>logisticsFreight>>icon	币种符号	否	[string]	
data>>list>>transaction>>logisticsFreight>>value	金额	否	[number]	
data>>list>>transaction>>outboundCost	出库成本	否	[object]	
data>>list>>transaction>>outboundCost>>currency	币种	否	[string]	
data>>list>>transaction>>outboundCost>>icon	币种符号	否	[string]	
data>>list>>transaction>>outboundCost>>value	金额	否	[number]	
data>>list>>transaction>>platformLogisticsFree	平台物流费	否	[object]	
data>>list>>transaction>>platformLogisticsFree>>currency	币种	否	[string]	
data>>list>>transaction>>platformLogisticsFree>>icon	币种符号	否	[string]	
data>>list>>transaction>>platformLogisticsFree>>value	金额	否	[number]	
data>>list>>transaction>>platformOtherFee	平台其他费	否	[object]	
data>>list>>transaction>>platformOtherFee>>currency	币种	否	[string]	
data>>list>>transaction>>platformOtherFee>>icon	币种符号	否	[string]	
data>>list>>transaction>>platformOtherFee>>value	金额	否	[number]	
data>>list>>transaction>>platformSubsidy	平台补贴	否	[object]	
data>>list>>transaction>>platformSubsidy>>currency	币种	否	[string]	
data>>list>>transaction>>platformSubsidy>>icon	币种符号	否	[string]	
data>>list>>transaction>>platformSubsidy>>value	金额	否	[number]	
data>>list>>transaction>>productDiscount	商品折扣	否	[object]	
data>>list>>transaction>>productDiscount>>currency	币种	否	[string]	
data>>list>>transaction>>productDiscount>>icon	币种符号	否	[string]	
data>>list>>transaction>>productDiscount>>value	金额	否	[number]	
data>>list>>transaction>>productTax	商品税	否	[object]	
data>>list>>transaction>>productTax>>currency	币种	否	[string]	
data>>list>>transaction>>productTax>>icon	币种符号	否	[string]	
data>>list>>transaction>>productTax>>value	金额	否	[number]	
data>>list>>transaction>>productTotalAmount	商品总额	否	[object]	
data>>list>>transaction>>productTotalAmount>>currency	币种	否	[string]	
data>>list>>transaction>>productTotalAmount>>icon	币种符号	否	[string]	
data>>list>>transaction>>productTotalAmount>>value	金额	否	[number]	
data>>list>>transaction>>productTotalAmountTaxIncluded	商品总额是否含税	否	[boolean]	
data>>list>>transaction>>realPayAmount	实付金额	否	[object]	
data>>list>>transaction>>realPayAmount>>currency	币种	否	[string]	
data>>list>>transaction>>realPayAmount>>icon	币种符号	否	[string]	
data>>list>>transaction>>realPayAmount>>value	金额	否	[number]	
data>>list>>transaction>>salesRevenue	销售收益	否	[object]	
data>>list>>transaction>>salesRevenue>>currency	币种	否	[string]	
data>>list>>transaction>>salesRevenue>>icon	币种符号	否	[string]	
data>>list>>transaction>>salesRevenue>>value	金额	否	[number]	
data>>list>>transaction>>shippingFee	客付运费	否	[object]	
data>>list>>transaction>>shippingFee>>currency	币种	否	[string]	
data>>list>>transaction>>shippingFee>>icon	币种符号	否	[string]	
data>>list>>transaction>>shippingFee>>value	金额	否	[number]	
data>>list>>transaction>>shippingFeeDiscount	运费折扣	否	[object]	
data>>list>>transaction>>shippingFeeDiscount>>currency	币种	否	[string]	
data>>list>>transaction>>shippingFeeDiscount>>icon	币种符号	否	[string]	
data>>list>>transaction>>shippingFeeDiscount>>value	金额	否	[number]	
data>>list>>transaction>>shippingFeeTaxIncluded	客付运费是否含税	否	[boolean]	
data>>list>>transaction>>shippingTax	运费税	否	[object]	
data>>list>>transaction>>shippingTax>>currency	币种	否	[string]	
data>>list>>transaction>>shippingTax>>icon	币种符号	否	[string]	
data>>list>>transaction>>shippingTax>>value	金额	否	[number]	
data>>list>>transaction>>tax	税费	否	[object]	
data>>list>>transaction>>tax>>currency	币种	否	[string]	
data>>list>>transaction>>tax>>icon	币种符号	否	[string]	
data>>list>>transaction>>tax>>value	金额	否	[number]	
data>>pages	当前页码	否	[long]	
data>>size	当前页数量	否	[long]	
data>>total	总数	否	[long]
返回请求示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "2939b59f09bd49359c8cc328980a12f1.288.17534161636756829",
    "response_time": "2025-07-25 12:02:43",
    "data": {
        "total": 3,
        "size": 100,
        "current": 1,
        "pages": 1,
        "list": [
            {
                "order": {
                    "storeId": "110595848248180224",
                    "storeName": "真实店铺勿动!!",
                    "platformOrderNo": "129020638291281",
                    "platformOrderName": "129020638291281",
                    "referenceNo": "200013306906980",
                    "uniqueOrderNo": "110595848248180224#129020638291281",
                    "buyerName": null,
                    "buyerEmail": "32818112E02443689A0744011F043312@relay.walmart.com",
                    "buyerMessage": null,
                    "sourceName": null,
                    "platformCode": "10008",
                    "platformName": "Walmart",
                    "deliveryType": 1,
                    "orderStatusName": "Shipped",
                    "deliveryStatusName": null,
                    "paymentStatusName": null,
                    "siteCode": "10008-US",
                    "purchaseTime": "2025-07-24 08:34:14 -7:00",
                    "purchaseTimeBj": "2025-07-24 23:34:14 +8:00",
                    "paymentTime": "2025-07-24 08:34:14 -7:00",
                    "paymentTimeBj": "2025-07-24 23:34:14 +8:00",
                    "deliveryTime": "2025-07-24 14:27:10 -7:00",
                    "deliveryTimeBj": "2025-07-25 05:27:10 +8:00",
                    "latestDeliveryTime": "2025-07-24 14:31:00 -7:00",
                    "latestShipTime": null,
                    "expectedDeliveryTime": null,
                    "platformOrderModifiedTime": "2025-07-24T21:33:36.000+00:00",
                    "remark": null
                },
                "itemList": [
                    {
                        "orderItemNo": "752",
                        "productNo": null,
                        "deliveryType": 1,
                        "msku": "MCS_Pencils_Test",
                        "mskuId": "MCS_Pencils_Test",
                        "title": null,
                        "variantAttr": null,
                        "quantity": 1,
                        "unitPriceAmount": "0.00",
                        "productStatusName": "Shipped",
                        "currency": "",
                        "imgUrl": null,
                        "customizedText": null,
                        "remark": null,
                        "productUrl": null,
                        "mskuRelationKey": "MCS_Pencils_Test",
                        "pid": null,
                        "localSku": null,
                        "pname": null,
                        "model": null,
                        "psId": null,
                        "spu": null,
                        "spuName": null,
                        "bid": null,
                        "cid": null,
                        "productDeveloperUid": null,
                        "customerShipping": "Express",
                        "transaction": {
                            "grossProfit": null,
                            "grossMargin": null,
                            "salesRevenue": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "productTotalAmount": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "productTotalAmountTaxIncluded": false,
                            "productTax": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "productDiscount": null,
                            "shippingFeeDiscount": null,
                            "shippingFee": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "shippingFeeTaxIncluded": false,
                            "shippingTax": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "platformOtherFee": null,
                            "giftWrappingFee": null,
                            "realPayAmount": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "platformSubsidy": null,
                            "tax": {
                                "value": "0.00",
                                "currency": ""
                            },
                            "outboundCost": null,
                            "logisticsFreight": null,
                            "platformLogisticsFree": null
                        }
                    }
                ],
                "transaction": {
                    "grossProfit": null,
                    "grossMargin": null,
                    "salesRevenue": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "productTotalAmount": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "productTotalAmountTaxIncluded": false,
                    "productTax": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "productDiscount": null,
                    "shippingFeeDiscount": null,
                    "shippingFee": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "shippingFeeTaxIncluded": false,
                    "shippingTax": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "platformOtherFee": null,
                    "giftWrappingFee": null,
                    "realPayAmount": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "platformSubsidy": null,
                    "tax": {
                        "value": "0.00",
                        "currency": ""
                    },
                    "outboundCost": null,
                    "logisticsFreight": null,
                    "platformLogisticsFree": null
                },
                "address": {
                    "recipientName": "Seung Ro",
                    "recipientPhone": "0000000000",
                    "recipientCountry": "US",
                    "recipientPostalCode": "60137",
                    "detailAddress": "160 Cranston Ct.,Glen Ellyn,IL,60137,US,",
                    "recipientStateOrRegion": "IL",
                    "recipientCity": "Glen Ellyn",
                    "recipientDistrict": null,
                    "houseNumber": null,
                    "addressType": 2,
                    "addressTypeName": "住宅",
                    "companyName": null,
                    "addressLine1": "160 Cranston Ct.",
                    "addressLine2": null,
                    "addressLine3": null
                },
                "tagInfos": [],
                "erpInfo": null,
                "packages": [
                    {
                        "packageId": "1LSCXLNA006748496",
                        "shippingProviderName": null,
                        "trackingNumber": "1",
                        "deliveryTimeBj": "2025-07-25 05:27:10",
                        "itemList": [
                            {
                                "orderItemNo": "752",
                                "productNo": null,
                                "msku": "MCS_Pencils_Test",
                                "mskuId": "MCS_Pencils_Test",
                                "title": null,
                                "variantAttr": null,
                                "quantity": 1,
                                "imgUrl": null
                            }
                        ]
                    }
                ]
            },
            {
                "order": {
                    "storeId": "11091275586175488",
                    "storeName": "监控测试03",
                    "platformOrderNo": "6594028994753",
                    "platformOrderName": "#1259",
                    "referenceNo": "6594028994753",
                    "uniqueOrderNo": "11091275586175488#6594028994753",
                    "buyerName": "Yaowei Chen",
                    "buyerEmail": "chenyaowei@lingxing.com",
                    "buyerMessage": "这个是留言",
                    "sourceName": "shopify_draft_order",
                    "platformCode": "10002",
                    "platformName": "Shopify",
                    "deliveryType": 0,
                    "orderStatusName": null,
                    "deliveryStatusName": "Partial Fulfilled",
                    "paymentStatusName": "Paid",
                    "siteCode": "",
                    "purchaseTime": "2025-07-21 05:41:40 -4:00",
                    "purchaseTimeBj": "2025-07-21 17:41:40 +8:00",
                    "paymentTime": "",
                    "paymentTimeBj": null,
                    "deliveryTime": "2025-07-24 01:57:48 -4:00",
                    "deliveryTimeBj": "2025-07-24 13:57:48 +8:00",
                    "latestDeliveryTime": null,
                    "latestShipTime": null,
                    "expectedDeliveryTime": null,
                    "platformOrderModifiedTime": "2025-07-24T22:01:49.000+00:00",
                    "remark": null
                },
                "itemList": [
                    {
                        "orderItemNo": "16510851186881",
                        "productNo": "6700083839169",
                        "deliveryType": 0,
                        "msku": "custom001",
                        "mskuId": "custom001",
                        "title": "测试定制产品",
                        "variantAttr": "41 / red",
                        "quantity": 3,
                        "unitPriceAmount": "1000.00",
                        "productStatusName": null,
                        "currency": "CNY",
                        "imgUrl": null,
                        "customizedText": null,
                        "remark": null,
                        "productUrl": null,
                        "mskuRelationKey": "custom001",
                        "pid": null,
                        "localSku": null,
                        "pname": null,
                        "model": null,
                        "psId": null,
                        "spu": null,
                        "spuName": null,
                        "bid": null,
                        "cid": null,
                        "productDeveloperUid": null,
                        "customerShipping": null,
                        "transaction": {
                            "grossProfit": null,
                            "grossMargin": null,
                            "salesRevenue": {
                                "value": "3150.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmount": {
                                "value": "3000.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmountTaxIncluded": false,
                            "productTax": {
                                "value": "150.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeTaxIncluded": false,
                            "shippingTax": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformOtherFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "giftWrappingFee": null,
                            "realPayAmount": {
                                "value": "3150.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformSubsidy": null,
                            "tax": {
                                "value": "-150.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "outboundCost": null,
                            "logisticsFreight": null,
                            "platformLogisticsFree": null
                        }
                    },
                    {
                        "orderItemNo": "16510851219649",
                        "productNo": "9943760863425",
                        "deliveryType": 0,
                        "msku": "CJJJJFMJ00034-Gray-Set-Cm",
                        "mskuId": "CJJJJFMJ00034-Gray-Set-Cm",
                        "title": "Household Pure Cotton Towel Towel Bath Towel",
                        "variantAttr": "Gray / Set / Cm",
                        "quantity": 4,
                        "unitPriceAmount": "100.00",
                        "productStatusName": null,
                        "currency": "CNY",
                        "imgUrl": "",
                        "customizedText": null,
                        "remark": null,
                        "productUrl": null,
                        "mskuRelationKey": "CJJJJFMJ00034-Gray-Set-Cm",
                        "pid": 29855,
                        "localSku": "MMM-GGG",
                        "pname": "MMM-GGG",
                        "model": "",
                        "psId": 0,
                        "spu": "",
                        "spuName": "",
                        "bid": 0,
                        "cid": 0,
                        "productDeveloperUid": 0,
                        "customerShipping": null,
                        "transaction": {
                            "grossProfit": null,
                            "grossMargin": null,
                            "salesRevenue": {
                                "value": "400.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmount": {
                                "value": "400.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmountTaxIncluded": false,
                            "productTax": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeTaxIncluded": false,
                            "shippingTax": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformOtherFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "giftWrappingFee": null,
                            "realPayAmount": {
                                "value": "400.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformSubsidy": null,
                            "tax": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "outboundCost": null,
                            "logisticsFreight": null,
                            "platformLogisticsFree": null
                        }
                    },
                    {
                        "orderItemNo": "16510851252417",
                        "productNo": "7605912830145",
                        "deliveryType": 0,
                        "msku": "SKU001",
                        "mskuId": "SKU001",
                        "title": "测试不需发货",
                        "variantAttr": null,
                        "quantity": 4,
                        "unitPriceAmount": "19.00",
                        "productStatusName": null,
                        "currency": "CNY",
                        "imgUrl": "https://image.umaicloud.com/lingxing/20210521/7e9fa40621ee33609692b315b7900c61.jpg",
                        "customizedText": null,
                        "remark": null,
                        "productUrl": null,
                        "mskuRelationKey": "SKU001",
                        "pid": 10154,
                        "localSku": "SKU001",
                        "pname": "电脑配件套装",
                        "model": "",
                        "psId": 0,
                        "spu": "",
                        "spuName": "",
                        "bid": 0,
                        "cid": 0,
                        "productDeveloperUid": 0,
                        "customerShipping": null,
                        "transaction": {
                            "grossProfit": null,
                            "grossMargin": null,
                            "salesRevenue": {
                                "value": "79.80",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmount": {
                                "value": "76.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productTotalAmountTaxIncluded": false,
                            "productTax": {
                                "value": "3.80",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "productDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeDiscount": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "shippingFeeTaxIncluded": false,
                            "shippingTax": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformOtherFee": {
                                "value": "0.00",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "giftWrappingFee": null,
                            "realPayAmount": {
                                "value": "79.80",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "platformSubsidy": null,
                            "tax": {
                                "value": "-3.80",
                                "currency": "CNY",
                                "icon": "￥"
                            },
                            "outboundCost": null,
                            "logisticsFreight": null,
                            "platformLogisticsFree": null
                        }
                    }
                ],
                "transaction": {
                    "grossProfit": null,
                    "grossMargin": 0.0000,
                    "salesRevenue": {
                        "value": "3629.80",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "productTotalAmount": {
                        "value": "3476.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "productTotalAmountTaxIncluded": false,
                    "productTax": {
                        "value": "153.80",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "productDiscount": {
                        "value": "0.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "shippingFeeDiscount": {
                        "value": "0.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "shippingFee": {
                        "value": "0.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "shippingFeeTaxIncluded": false,
                    "shippingTax": {
                        "value": "0.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "platformOtherFee": {
                        "value": "0.00",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "giftWrappingFee": null,
                    "realPayAmount": {
                        "value": "3629.80",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "platformSubsidy": null,
                    "tax": {
                        "value": "-153.80",
                        "currency": "CNY",
                        "icon": "￥"
                    },
                    "outboundCost": null,
                    "logisticsFreight": null,
                    "platformLogisticsFree": null
                },
                "address": {
                    "recipientName": "Amanda Chen",
                    "recipientPhone": "+11934222463",
                    "recipientCountry": "United States",
                    "recipientPostalCode": "72211",
                    "detailAddress": "Rua Abelardo Pires de Ávila 1176,Little Rock,Arkansas,72211,United States,",
                    "recipientStateOrRegion": "Arkansas",
                    "recipientCity": "Little Rock",
                    "recipientDistrict": null,
                    "houseNumber": null,
                    "addressType": 0,
                    "addressTypeName": "住宅",
                    "companyName": null,
                    "addressLine1": "Rua Abelardo Pires de Ávila 1176",
                    "addressLine2": null,
                    "addressLine3": null
                },
                "tagInfos": [],
                "erpInfo": null,
                "packages": [
                    {
                        "packageId": "5860931633345",
                        "shippingProviderName": "FedEx",
                        "trackingNumber": "765456788888",
                        "deliveryTimeBj": "2025-07-24 13:54:44",
                        "itemList": [
                            {
                                "orderItemNo": "16510851186881",
                                "productNo": "6700083839169",
                                "msku": "custom001",
                                "mskuId": "custom001",
                                "title": "测试定制产品",
                                "variantAttr": "41 / red",
                                "quantity": 1,
                                "imgUrl": null
                            }
                        ]
                    },
                    {
                        "packageId": "5860938612929",
                        "shippingProviderName": "FedEx",
                        "trackingNumber": "765456788888",
                        "deliveryTimeBj": "2025-07-24 13:57:48",
                        "itemList": [
                            {
                                "orderItemNo": "16510851252417",
                                "productNo": "7605912830145",
                                "msku": "SKU001",
                                "mskuId": "SKU001",
                                "title": "测试不需发货",
                                "variantAttr": null,
                                "quantity": 1,
                                "imgUrl": "https://image.umaicloud.com/lingxing/20210521/7e9fa40621ee33609692b315b7900c61.jpg"
                            }
                        ]
                    }
                ]
            },
            {
                "order": {
                    "storeId": "110547389043513344",
                    "storeName": "测试店铺-MY",
                    "platformOrderNo": "480041002944310",
                    "platformOrderName": "480041002944310",
                    "referenceNo": "480041002944310",
                    "uniqueOrderNo": "110547389043513344#480041002944310",
                    "buyerName": "26b21d0e ",
                    "buyerEmail": null,
                    "buyerMessage": "",
                    "sourceName": null,
                    "platformCode": "10007",
                    "platformName": "Lazada",
                    "deliveryType": 0,
                    "orderStatusName": "Pending",
                    "deliveryStatusName": null,
                    "paymentStatusName": null,
                    "siteCode": "10007-MY",
                    "purchaseTime": "2025-07-21 13:35:19 +8:00",
                    "purchaseTimeBj": "2025-07-21 13:35:19 +8:00",
                    "paymentTime": "",
                    "paymentTimeBj": null,
                    "deliveryTime": null,
                    "deliveryTimeBj": null,
                    "latestDeliveryTime": "",
                    "latestShipTime": null,
                    "expectedDeliveryTime": null,
                    "platformOrderModifiedTime": "2025-07-22T05:50:27.000+00:00",
                    "remark": null
                },
                "itemList": [
                    {
                        "orderItemNo": "480041003044310",
                        "productNo": "3871789773",
                        "deliveryType": 0,
                        "msku": "51165165",
                        "mskuId": "51165165",
                        "title": "NIIMBOT index Smart Printer Thermal Label Stickers Colorful Label Sticker Waterproof Paper For B21 B3S B1 B203 Self-adhesive Tag",
                        "variantAttr": null,
                        "quantity": 1,
                        "unitPriceAmount": "95.20",
                        "productStatusName": "Pending",
                        "currency": "MYR",
                        "imgUrl": "https://my-live.slatic.net/p/a338b271e73052d7c44c2023c9bb4b14.jpg",
                        "customizedText": null,
                        "remark": null,
                        "productUrl": null,
                        "mskuRelationKey": "51165165",
                        "pid": null,
                        "localSku": null,
                        "pname": null,
                        "model": null,
                        "psId": null,
                        "spu": null,
                        "spuName": null,
                        "bid": null,
                        "cid": null,
                        "productDeveloperUid": null,
                        "customerShipping": null,
                        "transaction": {
                            "grossProfit": null,
                            "grossMargin": null,
                            "salesRevenue": {
                                "value": "81.87",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "productTotalAmount": {
                                "value": "95.20",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "productTotalAmountTaxIncluded": false,
                            "productTax": {
                                "value": "0.00",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "productDiscount": {
                                "value": "-15.23",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "shippingFeeDiscount": {
                                "value": "-4.90",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "shippingFee": {
                                "value": "4.90",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "shippingFeeTaxIncluded": false,
                            "shippingTax": null,
                            "platformOtherFee": null,
                            "giftWrappingFee": null,
                            "realPayAmount": {
                                "value": "79.97",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "platformSubsidy": {
                                "value": "1.90",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "tax": {
                                "value": "0.00",
                                "currency": "MYR",
                                "icon": "RM"
                            },
                            "outboundCost": null,
                            "logisticsFreight": null,
                            "platformLogisticsFree": {
                                "value": "0.00",
                                "currency": "MYR",
                                "icon": "RM"
                            }
                        }
                    }
                ],
                "transaction": {
                    "grossProfit": null,
                    "grossMargin": null,
                    "salesRevenue": {
                        "value": "81.87",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "productTotalAmount": {
                        "value": "95.20",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "productTotalAmountTaxIncluded": false,
                    "productTax": {
                        "value": "0.00",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "productDiscount": {
                        "value": "-15.23",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "shippingFeeDiscount": {
                        "value": "-4.90",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "shippingFee": {
                        "value": "4.90",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "shippingFeeTaxIncluded": false,
                    "shippingTax": null,
                    "platformOtherFee": null,
                    "giftWrappingFee": null,
                    "realPayAmount": {
                        "value": "79.97",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "platformSubsidy": {
                        "value": "1.90",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "tax": {
                        "value": "0.00",
                        "currency": "MYR",
                        "icon": "RM"
                    },
                    "outboundCost": null,
                    "logisticsFreight": null,
                    "platformLogisticsFree": {
                        "value": "0.00",
                        "currency": "MYR",
                        "icon": "RM"
                    }
                },
                "address": {
                    "recipientName": "test u ",
                    "recipientPhone": "600149052502",
                    "recipientCountry": "Malaysia",
                    "recipientPostalCode": "82100",
                    "detailAddress": "深圳市 南山区, 阿里中心T2 test,Ayer Baloi,Johor,82100,Malaysia,",
                    "recipientStateOrRegion": "Johor",
                    "recipientCity": "Ayer Baloi",
                    "recipientDistrict": null,
                    "houseNumber": null,
                    "addressType": null,
                    "addressTypeName": null,
                    "companyName": null,
                    "addressLine1": "深圳市 南山区, 阿里中心T2 test",
                    "addressLine2": "",
                    "addressLine3": null
                },
                "tagInfos": [],
                "erpInfo": null,
                "packages": [
                    {
                        "packageId": "",
                        "shippingProviderName": "",
                        "trackingNumber": "",
                        "deliveryTimeBj": null,
                        "itemList": [
                            {
                                "orderItemNo": "480041003044310",
                                "productNo": "3871789773",
                                "msku": "51165165",
                                "mskuId": "51165165",
                                "title": "NIIMBOT index Smart Printer Thermal Label Stickers Colorful Label Sticker Waterproof Paper For B21 B3S B1 B203 Self-adhesive Tag",
                                "variantAttr": null,
                                "quantity": 1,
                                "imgUrl": "https://my-live.slatic.net/p/a338b271e73052d7c44c2023c9bb4b14.jpg"
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
查询Walmart Review列表
下一章节 
查询VC店铺列表