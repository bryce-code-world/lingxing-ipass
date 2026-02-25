查询VC发货单详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/getInvoice/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderNo	订单号	是	[string]	“11”
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	c94201a63d0c45ee9bb07cc6d6e4a260.1732604778665
response_time	响应时间	是	[string]	2024-11-26 15:06:20
data		是	[object]	
data>>invoice	发货单信息	是	[object]	
data>>invoice>>orderNo	发货单号	是	[string]	RO2500000
data>>invoice>>purchaseOrderNumber	订单号/货件号	是	[string]	KKKKKKKKKKK
data>>invoice>>remark	备注	是	[string]	备注信息
data>>invoice>>shippingWid	发货仓库id	是	[string]	11
data>>invoice>>shippingWarehouseName	发货仓库名称	是	[string]	仓库11
data>>invoice>>shipmentTime	发货时间	是	[string]	2025-01-01 15:30:00
data>>invoice>>shipmentUser	发货人	是	[string]	发货人
data>>invoice>>status	发货状态	是	[number]	5
data>>invoice>>createUser	创建人名称	是	[string]	创建人
data>>invoice>>createTime	创建时间	是	[string]	2025-01-01 10:00:00
data>>invoice>>shipmentType	发货类型	是	[string]	2
data>>invoice>>statusName	状态名称	是	[string]	待配货
data>>invoice>>totalNum	总发货量	是	[number]	10
data>>invoice>>estimatedPickupTime	预计到货时间	是	[string]	2025-01-02 10:00:00
data>>invoice>>shipmentTypeName	出库类型	是	[string]	PO
data>>invoice>>sourceType	来源类型（0：订单生成，1：货件生成）	是	[number]	0
data>>invoice>>invoiceModel	下单模式（0：手工下单 1：系统下单）	是	[number]	1
data>>invoice>>outboundDate	出库日期	是	[string]	2025-01-01
data>>invoice>>items	发货单明细列表	是	[array]	
data>>invoice>>items>>orderNo	发货单号	是	[string]	RO25000
data>>invoice>>items>>purchaseOrderNumber	订单号	是	[string]	7LLKKKKKK
data>>invoice>>items>>remark	备注	是	[string]	备注
data>>invoice>>items>>shippingWid	发货仓库id	是	[string]	11
data>>invoice>>items>>shippingWarehouseName	发货仓库名称	是	[string]	仓库11
data>>invoice>>items>>shipmentTime	发货时间	是	[string]	2025-01-01 15:30:00
data>>invoice>>items>>shipmentUser	发货人	是	[string]	发货人
data>>invoice>>items>>isDeleted	是否删除	是	[string]	0
data>>invoice>>items>>status	发货状态	是	[number]	5
data>>invoice>>items>>createUser	创建人名称	是	[string]	创建人
data>>invoice>>items>>createTime	创建时间	是	[number]	1704067200000
data>>invoice>>items>>shipmentType	发货类型	是	[string]	2
data>>invoice>>items>>itemRemark	备注	是	[string]	商品备注
data>>invoice>>items>>sku	sku	是	[string]	SKU123456
data>>invoice>>items>>asin	asin	是	[string]	B00000000000
data>>invoice>>items>>productId	产品id	是	[string]	3211
data>>invoice>>items>>productName	产品名称	是	[string]	产品名称
data>>invoice>>items>>thirdPartyProductCode	海外仓产品编码	是	[string]	TPC123
data>>invoice>>items>>thirdPartyProductName	海外仓产品名称	是	[string]	第三方产品名称
data>>invoice>>items>>num	发货数量	是	[number]	5
data>>invoice>>items>>storeId	店铺id	是	[string]	1342250000000000000
data>>invoice>>items>>storeName	店铺名称	是	[string]	店铺名称
data>>invoice>>items>>warehouseStoreId	仓库店铺id	是	[string]	0
data>>invoice>>items>>warehouseStoreName	仓库店铺名称	是	[string]	仓库店铺名称
data>>invoice>>items>>picUrl	图片链接	是	[string]	https://example.com/image.jpg
data>>invoice>>items>>toShipAmount	待到货量	是	[number]	10
data>>invoice>>items>>vendorShipmentIdentifier	货件单号	是	[string]	VSI123456
data>>invoice>>items>>asn	asn	是	[string]	ASN123456
data>>invoice>>items>>buyerReferenceNumber	arn	是	[string]	BRN123456
data>>invoice>>items>>unitPrice	单位费用	是	[string]	10.50
data>>invoice>>items>>headStockUnitPrice	单位头程	是	[string]	8.50
data>>invoice>>items>>purchasePrice	采购单价	是	[string]	10.50
data>>invoice>>items>>stockUnitPrice	单位库存成本	是	[string]	9.00
data>>invoice>>items>>stockCost	库存成本	是	[string]	8.50
data>>invoice>>items>>itemDimensionsList	商品箱规列表	是	[array]	
data>>invoice>>items>>itemDimensionsList>>id	主键ID	是	[string]	123456
data>>invoice>>items>>itemDimensionsList>>purchaseOrderSkuId	订单商品主键i	是	[string]	POSK123456
data>>invoice>>items>>itemDimensionsList>>orderSn	发货单单号	是	[string]	RO250000000000000
data>>invoice>>items>>itemDimensionsList>>sourceOrderSn	来源单据	是	[string]	SO123456
data>>invoice>>items>>itemDimensionsList>>frontDimensionsId	页面展示的箱规Id	是	[string]	FD123456
data>>invoice>>items>>itemDimensionsList>>boxNum	箱数	是	[number]	2
data>>invoice>>items>>itemDimensionsList>>packedAmount	单箱数量	是	[number]	5
data>>invoice>>items>>itemDimensionsList>>dimensionsName	箱规名称	是	[string]	标准箱
data>>invoice>>items>>itemDimensionsList>>weight	箱子重量	是	[string]	2.5
data>>invoice>>items>>itemDimensionsList>>weightUnit	箱子重量单位	是	[string]	kg
data>>invoice>>items>>itemDimensionsList>>length	箱规长	是	[string]	30.0
data>>invoice>>items>>itemDimensionsList>>width	箱规宽	是	[string]	20.0
data>>invoice>>items>>itemDimensionsList>>height	箱规高	是	[string]	15.0
data>>invoice>>items>>itemDimensionsList>>dimensionsUnit	箱规单位	是	[string]	cm
data>>invoice>>items>>itemDimensionsList>>thirdPartyOrderNo	三方仓订单号	是	[string]	TPO123456
data>>invoice>>items>>itemDimensionsList>>thirdPartyOrderStatus	三方仓下单状态	是	[number]	1
data>>invoice>>items>>itemDimensionsList>>gmtCreate	创建时间	是	[string]	2025-01-01 10:00:00
data>>invoice>>items>>itemDimensionsList>>gmtModified	修改时间	是	[string]	2025-01-01 10:00:00
data>>invoice>>invoiceTrackingList	物流信息	是	[array]	
data>>invoice>>invoiceTrackingList>>orderSn	发货单单号	是	[string]	RO250101001
data>>invoice>>invoiceTrackingList>>sourceOrderSn	来源单据（货件or订单）	是	[string]	SO123456
data>>invoice>>invoiceTrackingList>>boxLabelUrl	货件标签文件地址	是	[string]	https://example.com/box-label.pdf
data>>invoice>>invoiceTrackingList>>cartonLabelUrl	箱唛文件地址	是	[string]	https://example.com/carton-label.pdf
data>>invoice>>invoiceTrackingList>>trackingNumber	跟踪编号	是	[string]	TR123456789
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "bc6f44535f4e4f208da1ae4e4ec73953.1758797083737",
    "response_time": "2025-09-25 18:44:44",
    "data": {
        "count": 2,
        "list": [{
            "gmtCreate": "2025-01-01 10:00:00",
            "gmtModified": "2025-01-01 10:00:00",
            "id": "123456",
            "orderNo": "RO250101001",
            "purchaseOrderNumber": "7T2KKKKK",
            "remark": "备注信息",
            "shippingWid": "11",
            "shippingWarehouseName": "仓库11",
            "shipmentTime": "2025-01-01 15:30:00",
            "shipmentUser": "发货人",
            "status": 5,
            "createUser": "创建人",
            "createTime": "2025-01-01 10:00:00",
            "shipmentType": "2",
            "statusName": "待配货",
            "totalNum": 10,
            "estimatedPickupTime": "2025-01-02 10:00:00",
            "shipmentTypeName": "PO",
            "sourceType": 0,
            "invoiceModel": 1,
            "outboundDate": "2025-01-01",
            "items": [{
                "gmtCreate": "2025-01-01 10:00:00",
                "gmtModified": "2025-01-01 10:00:00",
                "id": "123456",
                "sku": "SKU123456",
                "asin": "B0999999999",
                "productId": "3211",
                "productName": "产品名称1",
                "num": 5,
                "storeId": "134444444444444444444",
                "storeName": "店铺名称",
                "warehouseStoreId": "0",
                "warehouseStoreName": "仓库店铺名称",
                "picUrl": "https://example.com/image1.jpg",
                "toShipAmount": 10,
                "purchasePrice": "10.50",
                "stockCost": "8.50",
                "orderNo": "RO2500000000",
                "purchaseOrderNumber": "7T2KKKKK",
                "remark": "备注",
                "shippingWid": "11",
                "shippingWarehouseName": "仓库11",
                "shipmentTime": "2025-01-01 15:30:00",
                "shipmentUser": "发货人",
                "isDeleted": "0",
                "status": 5,
                "createUid": "123456789",
                "createUser": "创建人",
                "createTime": "2025-01-01 10:00:00",
                "shipmentType": 2,
                "itemRemark": "商品备注1",
                "thirdPartyProductCode": "TPC123",
                "thirdPartyProductName": "第三方产品名称1",
                "vendorShipmentIdentifier": "VSI123456",
                "asn": "ASN123456",
                "buyerReferenceNumber": "BRN123456",
                "unitPrice": "10.50",
                "headStockUnitPrice": "8.50",
                "stockUnitPrice": "9.00"
            }]
        }]
    }
}
复制
错误
复制成功
 上一章节
查询VC发货单列表
下一章节 
查询亚马逊源报表-所有订单