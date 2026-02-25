查询VC发货单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/getInvoice/page/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	偏移量(默认0)	否	[number]	0
length	每页条数(默认20）	否	[number]	20
sids	店铺id	否	[array]	["1","2"]
wid	国家id	否	[array]	[1,2]
shipmentType	出库类型 1:DF 2:PO 3:DI	是	[string]	1
status	订单状态 0: 全部 5:待配货 10:待出库 15:已完成 100:已作废 (默认0）	否	[number]	0
createTimeStartTime	创建日期-开始	否	[string]	2025-01-01
createTimeEndTime	创建日期-结束	否	[string]	2025-12-31
shipmentTimeStartTime	出库日期-开始	否	[string]	2025-01-01
shipmentTimeEndTime	出库日期-结束	否	[string]	2025-12-31
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	c94201a63d0c45ee9bb07cc6d6e4a260.1732604778665
response_time	响应时间	是	[string]	2024-11-26 15:06:20
data		是	[object]	
data>>count	总记录数	是	[number]	2
data>>list	发货单列表	是	[array]	
data>>list>>gmtCreate	创建时间	是	[string]	2025-01-01 10:00:00
data>>list>>gmtModified	修改时间	是	[string]	2025-01-01 10:00:00
data>>list>>id	主键id	是	[string]	123456
data>>list>>orderNo	发货单号	是	[string]	RO000000
data>>list>>purchaseOrderNumber	订单号/货件号	是	[string]	7876AAAAA
data>>list>>remark	备注	是	[string]	备注信息
data>>list>>shippingWid	发货仓库ID	是	[string]	11
data>>list>>shippingWarehouseName	发货仓库名称	是	[string]	仓库11
data>>list>>shipmentTime	发货时间	是	[string]	2025-01-01 15:30:00
data>>list>>shipmentUser	发货人	是	[string]	发货人
data>>list>>status	发货状态	是	[number]	5
data>>list>>createUser	创建人名称	是	[string]	创建人
data>>list>>createTime	创建时间	是	[string]	2025-01-01 10:00:00
data>>list>>shipmentType	发货类型	是	[string]	2
data>>list>>statusName	状态名称	是	[string]	待配货
data>>list>>totalNum	总发货量	是	[number]	10
data>>list>>estimatedPickupTime	预计到货时间	是	[string]	2025-01-02 10:00:00
data>>list>>shipmentTypeName	出库类型名称	是	[string]	PO
data>>list>>sourceType	来源类型（0：订单生成，1：货件生成）	是	[number]	0
data>>list>>invoiceModel	下单模式（0：手工下单 1：系统下单）	是	[number]	1
data>>list>>outboundDate	出库日期	是	[string]	2025-01-01
data>>list>>items	发货单明细列表	是	[array]	
data>>list>>items>>gmtCreate	创建时间	是	[string]	2025-01-01 10:00:00
data>>list>>items>>gmtModified	修改时间	是	[string]	2025-01-01 10:00:00
data>>list>>items>>id	发货单明细主键id	是	[string]	123456
data>>list>>items>>sku	SKU	是	[string]	SKU123456
data>>list>>items>>asin	ASIN	是	[string]	B091TMFQ28
data>>list>>items>>productId	产品ID	是	[string]	3211
data>>list>>items>>productName	产品名称	是	[string]	产品名称1
data>>list>>items>>num	发货数量	是	[number]	5
data>>list>>items>>storeId	店铺ID	是	[string]	13422222222222
data>>list>>items>>storeName	店铺名称	是	[string]	店铺名称
data>>list>>items>>warehouseStoreId	仓库店铺ID	是	[string]	0
data>>list>>items>>warehouseStoreName	仓库店铺名称	是	[string]	仓库店铺名称
data>>list>>items>>picUrl	图片链接	是	[string]	https://example.com/image1.jpg
data>>list>>items>>toShipAmount	待到货量	是	[number]	10
data>>list>>items>>purchasePrice	采购单价	是	[string]	10.50
data>>list>>items>>stockCost	库存成本	是	[string]	8.50
data>>list>>items>>orderNo	发货单号	是	[string]	RO255555555555
data>>list>>items>>purchaseOrderNumber	订单号	是	[string]	7T2KKKKKKK
data>>list>>items>>remark	备注	是	[string]	备注
data>>list>>items>>shippingWid	发货仓库ID	是	[string]	11
data>>list>>items>>shippingWarehouseName	发货仓库名称	是	[string]	仓库11
data>>list>>items>>shipmentTime	发货时间（时间戳）	是	[string]	2025-01-01 15:30:00
data>>list>>items>>shipmentUser	发货人	是	[string]	发货人
data>>list>>items>>isDeleted	是否删除（0：未删除）	是	[string]	0
data>>list>>items>>status	发货状态	是	[number]	5
data>>list>>items>>createUser	创建人名称	是	[string]	创建人
data>>list>>items>>createTime	创建时间（时间戳）	是	[string]	2025-01-01 10:00:00
data>>list>>items>>shipmentType	发货类型	是	[number]	2
data>>list>>items>>itemRemark	备注	是	[string]	商品备注1
data>>list>>items>>thirdPartyProductCode	海外仓产品编码	是	[string]	TPC123
data>>list>>items>>thirdPartyProductName	海外仓产品名称	是	[string]	第三方产品名称1
data>>list>>items>>vendorShipmentIdentifier	货件单号	是	[string]	VSI123456
data>>list>>items>>asn	ASN	是	[string]	ASN123456
data>>list>>items>>buyerReferenceNumber	ARN	是	[string]	BRN123456
data>>list>>items>>unitPrice	单位费用	是	[string]	10.50
data>>list>>items>>headStockUnitPrice	单位头程	是	[string]	8.50
data>>list>>items>>stockUnitPrice	单位库存成本	是	[string]	9.00
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "bc6f44535f4e4f208da1ae4e4ec73953555555555555555555",
    "response_time": "2025-09-25 18:44:44",
    "data": {
        "invoice": {
            "orderNo": "RO250000000",
            "purchaseOrderNumber": "7T2KKKKKK",
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
                "orderNo": "RO2500000000",
                "purchaseOrderNumber": "7T2Kkkkk",
                "remark": "备注",
                "shippingWid": "11",
                "shippingWarehouseName": "仓库11",
                "shipmentTime": "2025-01-01 15:30:00",
                "shipmentUser": "发货人",
                "isDeleted": "0",
                "status": 5,
                "createUid": "123456789",
                "createUser": "创建人",
                "createTime": 1704000000000,
                "shipmentType": "2",
                "itemRemark": "商品备注",
                "sku": "SKU123456",
                "asin": "B0000000000",
                "productId": "3211",
                "productName": "产品名称",
                "thirdPartyProductCode": "TTTTTTTTTT",
                "thirdPartyProductName": "第三方产品名称",
                "num": 5,
                "storeId": "1111111111111111",
                "storeName": "店铺名称",
                "warehouseStoreId": "0",
                "warehouseStoreName": "仓库店铺名称",
                "picUrl": "https://example.com/image.jpg",
                "toShipAmount": 10,
                "vendorShipmentIdentifier": "VSI1111111",
                "asn": "ASN111111",
                "buyerReferenceNumber": "BRN1111111",
                "unitPrice": "10.50",
                "headStockUnitPrice": "8.50",
                "purchasePrice": "10.50",
                "stockUnitPrice": "9.00",
                "stockCost": "8.50",
                "itemDimensionsList": [{
                    "id": "123456",
                    "purchaseOrderSkuId": "POSK123456",
                    "orderSn": "RO250000000",
                    "sourceOrderSn": "SO123456",
                    "containerSequenceNumber": "1",
                    "containerIdentificationNumber": "CIN123456",
                    "frontDimensionsId": "FD123456",
                    "boxNum": 2,
                    "packedAmount": 5,
                    "dimensionsName": "标准箱",
                    "weight": "2.5",
                    "weightUnit": "kg",
                    "length": "30.0",
                    "width": "20.0",
                    "height": "15.0",
                    "dimensionsUnit": "cm",
                    "thirdPartyOrderNo": "TPO123456",
                    "thirdPartyOrderStatus": 1,
                    "gmtCreate": "2025-01-01 10:00:00",
                    "gmtModified": "2025-01-01 10:00:00"
                }]
            }],
            "invoiceTrackingList": [{
                "orderSn": "RO25000000000",
                "sourceOrderSn": "SO123456",
                "containerSequenceNumber": 1,
                "containerIdentificationNumber": "CIN123456",
                "boxLabelUrl": "https://example.com/box-label.pdf",
                "cartonLabelUrl": "https://example.com/carton-label.pdf",
                "trackingNumber": "TR123456789"
            }]
        }
    }
}
复制
错误
复制成功
 上一章节
配对/批量配对
下一章节 
查询VC发货单详情