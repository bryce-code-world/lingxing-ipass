查询结算利润（利润报表）-msku
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/profit/report/msku	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	是	[number]	0
length	分页长度，默认1000	是	[number]	100
platformCodeS	平台id：
10002 Shopify
10003 eBay
10005 AliExpress
10006 Shopee
10008 walmart
10011 Tiktok
10021 Shein平台模式
10022 Temu全托
10024 Temu半托
10028 Shein半托管	否	[array]	["10024"]
mids	国家id，多个使用英文逗号分隔	否	[string]	NA,MX,BR,US,CA
sids	店铺id，多个使用英文逗号分隔 ，对应查询多平台店铺信息接口对应字段【store_id】	否	[string]	110424575139430912
currencyCode	币种code：
原币种
USD
EUR
GBP
CNY	否	[string]	USD
startDate	开始时间【结算日期】，闭区间，格式：Y-m-d	是	[string]	2024-09-01
endDate	结束时间【结算日期】，闭区间，格式：Y-m-d	是	[string]	2024-09-30
searchField	搜索值类型：
msku MSKU
local_sku SKU
platform_order_no 平台单号	否	[string]	local_sku
searchValue	搜索值	否	[string]	123
developers	开发人	否	[array]	[128581]
cids	分类	否	[array]	[14]
bids	品牌	否	[array]	[2]
请求示例
{
    "offset": 0,
    "length": 100,
    "platformCodeS": ["10024"],
    "mids": "NA,MX,BR,US,CA",
    "sids": "110424575139430912",
    "currencyCode": "USD",
    "startDate": "2024-01-01",
    "endDate": "2024-06-30",
    "searchField": "local_sku",
    "searchValue": "test",
    "developers": [128581],
    "cids": [14],
    "bids": [2]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	6b6cfa6db9564a13afc284ec69cba07b.1728726118087
response_time	响应时间	是	[string]	2024-10-12 17:41:59
data	响应数据	是	[object]	
data>>totalSum	总计	是	[object]	
data>>totalSum>>currencyCode	币种	是	[string]	CNY
data>>totalSum>>currencyIcon	货币图标	是	[string]	￥
data>>totalSum>>salesNum	销量	是	[number]	281
data>>totalSum>>replacementNum	补货量	是	[number]	0
data>>totalSum>>salesAmount	销售额	是	[number]	557799.61
data>>totalSum>>promotionDiscountAmount	促销折扣	是	[number]	0
data>>totalSum>>buyerFreightAmount	买家运费	是	[number]	0
data>>totalSum>>platformOtherIncomeAmount	其他收入	是	[number]	0
data>>totalSum>>incomeRefundAmount	收入退款额	是	[number]	0
data>>totalSum>>feeRefundAmount	费用退款额	是	[number]	0
data>>totalSum>>refundAmount	退款金额	是	[number]	0
data>>totalSum>>refundNum	退款量	是	[string]	0.00
data>>totalSum>>refundRate	退款率	是	[number]	0
data>>totalSum>>returnedGoodsNum	退货量	是	[number]	0
data>>totalSum>>returnedGoodsRate	退货率	是	[number]	0
data>>totalSum>>promotionAmount	平台费	是	[number]	0
data>>totalSum>>platformLogisticsAmount	平台物流费	是	[number]	0
data>>totalSum>>promotionExtendAmount	推广费	是	[number]	0
data>>totalSum>>advertisementAmount	广告费	是	[number]	0
data>>totalSum>>adjustmentCostAmount	调整费	是	[number]	0
data>>totalSum>>platformStorageAmount	平台仓储费	是	[number]	0
data>>totalSum>>platformFineAmount	平台罚款	是	[number]	0
data>>totalSum>>platformOtherAmount	平台其他费	是	[number]	0
data>>totalSum>>taxAmount	销售税	是	[number]	0
data>>totalSum>>marketTaxAmount	市场税	是	[number]	0
data>>totalSum>>customOtherProductAmount	商品其他费	是	[number]	0
data>>totalSum>>customOtherSellerAmount	店铺其他费	是	[number]	0
data>>totalSum>>customOtherSalesOrderAmount	订单其他费	是	[number]	0
data>>totalSum>>purchaseAmount	采购成本	是	[number]	0
data>>totalSum>>transportationAmount	头程成本	是	[number]	0
data>>totalSum>>tailAmount	尾程成本	是	[number]	0
data>>totalSum>>otherAmount	其他成本	是	[number]	0
data>>totalSum>>grossProfit	毛利润	是	[number]	557799.61
data>>totalSum>>grossProfitRate	毛利率	是	[number]	1
data>>list	列表	是	[array]	
data>>list>>storeId	店铺id	是	[number]	110453560302608900
data>>list>>storeName	店铺	是	[string]	
data>>list>>platformCode	平台编码	是	[string]	10024
data>>list>>platformName	平台	是	[string]	temu
data>>list>>msku	MSKU	是	[string]	ES33HB5025-PG000
data>>list>>productName	品名	是	[string]	
data>>list>>localSku	SKU	是	[string]	
data>>list>>deliveryDate	下单时间（废弃字段，仅订单维度返回）	是	[string]	
data>>list>>countryName	国家	是	[null]	
data>>list>>bname	分类	是	[null]	
data>>list>>cname	品牌	是	[string]	
data>>list>>developer	开发人	是	[string]	
data>>list>>currencyCode	币种	是	[string]	USD
data>>list>>currencyIcon	货币图标	是	[string]	$
data>>list>>salesNum	销量	是	[number]	281
data>>list>>replacementNum	补货量	是	[number]	0
data>>list>>salesAmount	销售额	是	[number]	87669.88
data>>list>>promotionDiscountAmount	促销折扣	是	[number]	0
data>>list>>buyerFreightAmount	买家运费	是	[number]	0
data>>list>>platformOtherIncomeAmount	其他收入	是	[number]	0
data>>list>>incomeRefundAmount	收入退款额	是	[number]	0
data>>list>>feeRefundAmount	费用退款额	是	[number]	0
data>>list>>refundAmount	退款金额	是	[number]	0
data>>list>>refundNum	退款量	是	[string]	0.00
data>>list>>refundRate	退款率	是	[number]	0
data>>list>>returnedGoodsNum	退货量	是	[number]	0
data>>list>>returnedGoodsRate	退货率	是	[number]	0
data>>list>>promotionAmount	平台费	是	[number]	0
data>>list>>platformLogisticsAmount	平台物流费	是	[number]	0
data>>list>>promotionExtendAmount	推广费	是	[number]	0
data>>list>>advertisementAmount	广告费	是	[number]	0
data>>list>>adjustmentCostAmount	调整费	是	[number]	0
data>>list>>platformStorageAmount	平台仓储费	是	[number]	0
data>>list>>platformFineAmount	平台罚款	是	[number]	0
data>>list>>platformOtherAmount	平台其他费	是	[number]	0
data>>list>>taxAmount	销售税	是	[number]	0
data>>list>>marketTaxAmount	市场税	是	[number]	0
data>>list>>customOtherProductAmount	商品其他费	是	[number]	0
data>>list>>customOtherSellerAmount	店铺其他费	是	[number]	0
data>>list>>customOtherSalesOrderAmount	订单其他费	是	[number]	0
data>>list>>purchaseAmount	采购成本	是	[number]	0
data>>list>>transportationAmount	头程成本	是	[number]	0
data>>list>>tailAmount	尾程成本	是	[number]	0
data>>list>>otherAmount	其他成本	是	[number]	0
data>>list>>grossProfit	毛利润	是	[number]	87669.88
data>>list>>grossProfitRate	毛利率	是	[number]	1
data>>list>>otherTaxesFees	其他税费	否	[number]	
data>>list>>subsidyAmount	补贴金额	否	[number]	
沃尔玛平台字段补充				
data>>list>>platformWfsStorageAmount	WFS仓储费	否	[number]	
data>>list>>platformWfsRemoveAmount	WFS移除费	否	[number]	
data>>list>>wfsWarehousFee	WFS入仓费	否	[number]	
data>>list>>wfsPrepServiceFee	WFS上架前处理费	否	[number]	
data>>list>>wfsInventoryTransferFee	WFS存货转仓费	否	[number]	
data>>list>>wfsInventoryRTVFee	WFS RTV费	否	[number]	
data>>list>>productCommission	商品佣金	否	[number]	
data>>list>>shippingCommission	运费佣金	否	[number]	
data>>list>>wfsAdjustmentCostAmount	WFS调整费	否	[number]	
data>>list>>creditAdjustmentFee	信用额度调整费	否	[number]	
data>>list>>returnAdjustmentFee	退货调整费	否	[number]	
data>>list>>platformDetailOtherAmount	其他费	否	[number]	
data>>list>>commentAcceleratorFee	评论加速器费	否	[number]	
data>>list>>walmartSavingsBenefit	沃尔玛资助节省	否	[number]	
data>>list>>wfsLostInventoryFee	WFS库存丢失费	否	[number]	
data>>list>>wfsFoundInventoryFee	WFS库存找回费	否	[number]	
data>>list>>wfsDamageInWarehouseFee	WFS库存损坏费	否	[number]	
data>>list>>wfsReceivingErrorChargeBackFee	WFS接收错误收费	否	[number]	
data>>list>>wfsChargeFee	WFS调整收费	否	[number]	
data>>list>>wfsShipmentFee	WFS发货费	否	[number]	
data>>list>>walmartReturnServiceFee	沃尔玛退货服务费	否	[number]	
data>>list>>wfsReturnFee	WFS退货费	否	[number]	
data>>list>>platformAdvertisingFee	平台广告费	否	[number]	
data>>list>>semMarketingFee	自助搜索引擎营销费	否	[number]	
data>>list>>reserveCreditedBackAmount	释放储备金	否	[number]	
data>>list>>excessRefundAdjustmentAmount	超额退款调整	否	[number]	
data>>list>>walmartProductAdvertisingCreditsFee	沃尔玛产品广告积分	否	[number]	
data>>list>>walmartPromoCode	沃尔玛促销编码	否	[number]	
data>>list>>walmartExtraDiscount	沃尔玛额外折扣	否	[number]	
data>>list>>platformMultiChannelFulfillmentFee	多渠道配送费	否	[number]	
data>>list>>platformReserveFund	预留储备金	否	[number]	
Shopee平台字段补充				
data>>list>>sellerDiscount	卖家折扣	否	[number]	
data>>list>>shopeeDiscount	Shopee平台折扣	否	[number]	
data>>list>>discountFromCoin	Shopee币折扣	否	[number]	
data>>list>>discountFromVoucherShopee	Shopee优惠券折扣	否	[number]	
data>>list>>discountFromVoucherSeller	卖家优惠券折扣	否	[number]	
data>>list>>paymentPromotion	支付折扣	否	[number]	
data>>list>>sellerCoinCashBack	卖家Shopee币回扣	否	[number]	
data>>list>>shippingFeeDiscountFrom3pl	三方运费折扣	否	[number]	
data>>list>>sellerShippingDiscount	卖家运费折扣	否	[number]	
data>>list>>creditCardPromotion	信用卡折扣	否	[number]	
data>>list>>sellerLostCompensation	损失补偿	否	[number]	
data>>list>>shopeeShippingRebate	运费补偿	否	[number]	
data>>list>>sipSubsidy	SIP补贴	否	[number]	
data>>list>>sellerReturnRefund	退款金额	否	[number]	
data>>list>>drcAdjustableRefund	争议退款	否	[number]	
data>>list>>proratedCoinsValueOffsetReturnItems	Shopee币抵消退款	否	[number]	
data>>list>>proratedShopeeVoucherOffsetReturnItems	Shopee优惠券抵消退款	否	[number]	
data>>list>>proratedSellerVoucherOffsetReturnItems	卖家优惠券抵消退款	否	[number]	
data>>list>>proratedPaymentChannelPromoBankOffsetReturnItems	银行支付促销抵消退款	否	[number]	
data>>list>>proratedPaymentChannelPromoShopeeOffsetReturnItems	Shopee支付促销抵消退款	否	[number]	
data>>list>>sellerProtectionFeeClaimAmount	退货运费平台退款	否	[number]	
data>>list>>commissionFee	平台佣金	否	[number]	
data>>list>>amsCommissionFee	联盟营销佣金	否	[number]	
data>>list>>actualShippingFee	实际运费	否	[number]	
data>>list>>reverseShippingFee	退货运费	否	[number]	
data>>list>>finalReturnToSellerShippingFee	配送失败运费	否	[number]	
data>>list>>serviceFee	服务费	否	[number]	
data>>list>>buyerTransactionFee	买家交易手续费	否	[number]	
data>>list>>sellerTransactionFee	卖家交易手续费	否	[number]	
data>>list>>creditCardTransactionFee	信用卡交易手续费	否	[number]	
data>>list>>campaignFee	活动费	否	[number]	
data>>list>>shippingSellerProtectionFeeAmount	卖家保护费	否	[number]	
data>>list>>deliverySellerProtectionFeePremiumAmount	特别活动服务费	否	[number]	
data>>list>>overseasReturnServiceFee	海外退货服务费	否	[number]	
data>>list>>crossBorderTax	跨境税	否	[number]	
data>>list>>escrowTax	第三方托管税	否	[number]	
data>>list>>shippingFeeSst	运费税	否	[number]	
data>>list>>reverseShippingFeeSst	退货运费税	否	[number]	
data>>list>>salesTaxOnLvg	低值销售税	否	[number]	
data>>list>>finalProductVatTax	商品增值税	否	[number]	
data>>list>>finalShippingVatTax	运费增值税	否	[number]	
data>>list>>finalEscrowProductGst	GST商品增值税	否	[number]	
data>>list>>finalEscrowShippingGst	GST运费增值税	否	[number]	
data>>list>>vatOnImportedGoods	进口增值税	否	[number]	
data>>list>>selleroOrderProcessingFee	平台基础建设费	否	[number]	
data>>list>>buyerPaidPackagingFee	买家支付包装费	否	[number]	
data>>list>>withholdingPitTax	预扣个人所得税	否	[number]	
data>>list>>withholdingVatTax	预扣增值税	否	[number]	
data>>list>>withholdingTax	菲律宾市场税	否	[number]	
Temu全托管平台字段补充				
data>>list>>afterSalesDeduction	售后问题扣款	否	[number]	
data>>list>>stockingViolation	备货违规费	否	[number]	
data>>list>>qualityBreach	质量事故违规费	否	[number]	
data>>list>>ecoFeeForGood	商品环保费	否	[number]	
data>>list>>logisticsEcoPackagingFee	物流包装环保费（已扣费）	否	[number]	
data>>list>>withholdingServiceFee	代扣服务费	否	[number]	
data>>list>>promotionServiceFee	推广服务费	否	[number]	
data>>list>>eprAuthorizedAgencyFeeWithholding	EPR授权代理费代扣代缴	否	[number]	
data>>list>>platformOtherAdditionalAmount	平台其他费 - 其他费	否	[number]	
data>>list>>ecoFeeForRefundGood	商品环保费（退款）	否	[number]	
eBay平台字段补充				
data>>list>>ebaySubscriptionFee	eBay订阅费	否	[number]	
data>>list>>ebayPublicationFee	eBay刊登费	否	[number]	
data>>list>>regulatoryOperatingFee	监管运营费	否	[number]	
data>>list>>platformDetailOtherAmount	其它费	否	[number]	
Lazada平台字段补充				
data>>list>>fundedCommissionFromSellerVirtualCredit	SVC平台佣金	否	[number]
data>>list>>reversalPromotionalChargesVouchers	优惠券促销费用退款	否	[number]
data>>list>>reverseSellerVirtualCreditCoFundVoucher	SVC联合出资优惠券退款	否	[number]
data>>list>>sellerVirtualCreditCoFundVoucher	SVC联合出资优惠券	否	[number]
data>>list>>shippingFeeSubsidyBySeller	卖家运费补贴	否	[number]
data>>list>>promotionDiscountOtherAmount	促销折扣 - 其他费	否	[number]	
data>>list>>promotionOtherAmount	平台费 - 其他费	否	[number]	
data>>list>>feeRefundOtherAmount	费用退款额 - 其他费	否	[number]	
data>>list>>lazCoinsDiscount	LazCoinsDiscount金币折扣	否	[number]	
TikTok平台字段补充				
data>>list>>sellerDiscountAmount	卖家折扣	否	[number]	
data>>list>>shippingFeeDiscountAmount	运费折扣	否	[number]	
data>>list>>codServiceFeeAmount	COD费	否	[number]	
data>>list>>refundSubtotalBeforeDiscountAmount	销售额退款	否	[number]	
data>>list>>sellerDiscountRefundAmount	商家折扣退款	否	[number]	
data>>list>>refundCodServiceFeeAmount	COD费退款	否	[number]	
data>>list>>affiliateCommissionAmount	达人佣金	否	[number]	
data>>list>>affiliatePartnerCommissionAmount	达人伙伴佣金	否	[number]	
data>>list>>tspCommissionAmount	TSP佣金	否	[number]	
data>>list>>actualShippingFeeAmount	实际运费	否	[number]	
data>>list>>returnShippingFeeAmount	退货运费	否	[number]	
data>>list>>replacementShippingFeeAmount	替换运费	否	[number]	
data>>list>>exchangeShippingFeeAmount	换货运费	否	[number]	
data>>list>>signatureConfirmationFeeAmount	签名确认费	否	[number]	
data>>list>>shippingInsuranceFeeAmount	物流保险费	否	[number]	
data>>list>>transactionFeeAmount	交易费	否	[number]	
data>>list>>creditCardHandlingFeeAmount	信用卡处理费	否	[number]	
data>>list>>sfpServiceFeeAmount	SFP服务费	否	[number]	
data>>list>>liveSpecialsFeeAmount	直播特别项目费	否	[number]	
data>>list>>bonusCashbackServiceFeeAmount	奖金返现服务费	否	[number]	
data>>list>>mallServiceFeeAmount	TK商城服务费	否	[number]	
data>>list>>voucherXtraServiceFeeAmount	Voucher Xtra服务费	否	[number]	
data>>list>>flashSalesServiceFeeAmount	限时促销服务费	否	[number]	
data>>list>>cofundedPromotionServiceFeeAmount	共同出资促销服务费	否	[number]	
data>>list>>preOrderServiceFeeAmount	预售服务费	否	[number]	
data>>list>>sstAmount	马来西亚销售服务税	否	[number]	
data>>list>>gstAmount	新加坡商品服务税	否	[number]	
data>>list>>salesTaxRefundAmount	销售税退款额	否	[number]	
data>>list>>standardVatAmount	普通增值税	否	[number]	
data>>list>>importVatAmount	进口增值税	否	[number]	
data>>list>>ivaAmount	墨西哥增值税	否	[number]	
data>>list>>isrAmount	墨西哥联邦所得税	否	[number]	
data>>list>>antiDumpingDutyAmount	反倾销税	否	[number]	
data>>list>>customsDutyAmount	关税	否	[number]	
data>>list>>customsClearanceAmount	清关费	否	[number]	
data>>list>>chargeBack	争议退回	否	[number]	
data>>list>>customerServiceCompensation	客户服务补偿	否	[number]	
data>>list>>deductionsIncurredBySeller	卖家责任扣款	否	[number]	
data>>list>>gmvPaymentForAds	GMV广告扣费	否	[number]	
data>>list>>platformCommissionAdjustment	平台佣金调整	否	[number]	
data>>list>>platformCommissionCompensation	平台佣金补偿	否	[number]	
data>>list>>promotionAdjustment	优惠调整	否	[number]	
data>>list>>rebate	推荐费折扣	否	[number]	
data>>list>>platformCompensation	平台补偿	否	[number]	
data>>list>>platformReimbursement	平台退款报销	否	[number]	
data>>list>>cofundedCreatorRewards	共同出资创作者激励	否	[number]	
data>>list>>logisticsReimbursement	物流赔偿	否	[number]	
data>>list>>shippingFeeAdjustment	运费调整	否	[number]	
data>>list>>shippingFeeCompensation	运费补偿	否	[number]	
data>>list>>shippingFeeRebate	运费回扣	否	[number]	
data>>list>>sampleShippingFee	样品运费	否	[number]	
data>>list>>otherAdjustment	其他调整费	否	[number]	
data>>list>>fbtWarehouseServiceFee	FBT仓储服务费	否	[number]	
data>>list>>platformPenalty	平台罚款	否	[number]	
data>>list>>sellerPaylaterHandlingFeeAmount	先买后付手续费	否	[number]	
data>>list>>dtHandlingFeeAmount	DT配送手续费	否	[number]	
data>>list>>affiliateCommissionDeposit	联盟营销佣金收入	否	[number]	
data>>list>>affiliateCommissionRelease	联盟佣金释放	否	[number]	
data>>list>>tapShopAdsCommission	TapShop广告佣金	否	[number]	
data>>list>>shippingFeeGuaranteeServiceFee	物流保障服务费	否	[number]	
data>>list>>eprPobServiceFeeAmount	合规服务费	否	[number]	
data>>list>>feePerItemSoldAmount	单件商品销售服务费	否	[number]	
data>>list>>installationServiceFee	安装服务费	否	[number]	
data>>list>>cofundedCreatorBonusAmount	共同资助创作者奖金	否	[number]	
data>>list>>failedDeliverySubsidyAmount	配送失败补贴金额	否	[number]	
data>>list>>fbtFreeShippingFeeAmount	FBT免运费金额	否	[number]	
data>>list>>freeReturnSubsidyAmount	免费退货补贴金额	否	[number]	
data>>list>>returnShippingFeePaidBuyerAmount	卖家运费退回	否	[number]	
data>>list>>returnShippingLabelFeeAmount	退货标签费用	否	[number]	
data>>list>>shippingFeeGuaranteeReimbursement	物流保障服务费报销	否	[number]	
data>>list>>promoShippingIncentiveAmount	促销运费激励金额	否	[number]	
data>>list>>returnRefundSubsidyAmount	退货运费补偿款	否	[number]	
data>>list>>shippingFeeGuaranteeReimbursement	物流保障服务费报销	否	[number]	
data>>list>>promoShippingIncentiveAmount	促销运费激励金额	否	[number]	
data>>list>>returnRefundSubsidyAmount	退货运费补偿款	否	[number]	
data>>list>>fbtFulfillmentFeeReimbursementAmount	FBT发货费补偿款	否	[number]	
data>>list>>marketingBenefitsPackageFee	广告套餐包	否	[number]	
data>>list>>externalAffiliateMarketingFeeAmount	EAMS项目服务费	否	[number]	
data>>list>>fbtFulfillmentFeeReimbursementAmount	FBT发货费补偿款	否	[number]	
total	总数	是	[number]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.107.17289850958114669",
    "response_time": "2024-10-15 17:38:16",
    "data": {
        "totalSum": {
            "currencyCode": "CNY",
            "currencyIcon": "￥",
            "salesNum": 281,
            "replacementNum": 0,
            "salesAmount": 557799.61,
            "promotionDiscountAmount": 0.00,
            "buyerFreightAmount": 0.00,
            "platformOtherIncomeAmount": 0.00,
            "incomeRefundAmount": 0.00,
            "feeRefundAmount": 0.00,
            "refundAmount": 0.00,
            "refundNum": "0.00",
            "refundRate": 0.00,
            "returnedGoodsNum": 0,
            "returnedGoodsRate": 0.00,
            "promotionAmount": 0.00,
            "platformLogisticsAmount": 0.00,
            "promotionExtendAmount": 0.00,
            "advertisementAmount": 0.00,
            "adjustmentCostAmount": 0.00,
            "platformStorageAmount": 0.00,
            "platformFineAmount": 0.00,
            "platformOtherAmount": 0.00,
            "taxAmount": 0.00,
            "marketTaxAmount": 0.00,
            "customOtherProductAmount": 0.00,
            "customOtherSellerAmount": 0.00,
            "customOtherSalesOrderAmount": 0.00,
            "purchaseAmount": 0.00,
            "transportationAmount": 0.00,
            "tailAmount": 0.00,
            "otherAmount": 0.00,
            "grossProfit": 557799.61,
            "grossProfitRate": 1.00
        },
        "list": [
            {
                "storeId": 110453560302608896,
                "storeName": "",
                "platformCode": "10024",
                "platformName": "temu",
                "msku": "ES33HB5025-PG000",
                "productName": "",
                "localSku": "",
                "deliveryDate": "",
                "countryName": null,
                "bname": null,
                "cname": "",
                "developer": "",
                "currencyCode": "USD",
                "currencyIcon": "$",
                "salesNum": 281,
                "replacementNum": 0,
                "salesAmount": 87669.88,
                "promotionDiscountAmount": 0.00,
                "buyerFreightAmount": 0.00,
                "platformOtherIncomeAmount": 0.00,
                "incomeRefundAmount": 0.00,
                "feeRefundAmount": 0.00,
                "refundAmount": 0.00,
                "refundNum": "0.00",
                "refundRate": 0.00,
                "returnedGoodsNum": 0,
                "returnedGoodsRate": 0.00,
                "promotionAmount": 0.00,
                "platformLogisticsAmount": 0.00,
                "promotionExtendAmount": 0.00,
                "advertisementAmount": 0.00,
                "adjustmentCostAmount": 0.00,
                "platformStorageAmount": 0.00,
                "platformFineAmount": 0.00,
                "platformOtherAmount": 0.00,
                "taxAmount": 0.00,
                "marketTaxAmount": 0.00,
                "customOtherProductAmount": 0.00,
                "customOtherSellerAmount": 0.00,
                "customOtherSalesOrderAmount": 0.00,
                "purchaseAmount": 0.00,
                "transportationAmount": 0.00,
                "tailAmount": 0.00,
                "otherAmount": 0.00,
                "grossProfit": 87669.88,
                "grossProfitRate": 1.00
            }
        ]
    },
    "total": 1
}
复制
错误
复制成功
附加说明
斜体字段表示为可选的下拉明细费用项，如无需求可不获取。
 上一章节
查询利润报表-订单(旧版，将于04.30下线)
下一章节 
查询结算利润（利润报表）-sku