查询Temu货件
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/temu/cargo	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
endTime	yyyy-MM-dd	是	[date]	
length	每页条数	否	[long]	20
offset	偏移量	否	[long]	0
startTime	yyyy-MM-dd	是	[date]	
statusList	待发货：0 ；待收货：1 ；已收货：2 ；已入库：3 ；已退货：4 ；已取消：5 ；部分收货：6 ;待申报（本地状态）7	是	[array]	
storeIdList		是	[array]	
timeType	1:创建时间 2：发货时间 3：收货时间 4：入库时间	是	[int]	
请求示例
{
    "endTime": "",
    "length": 20,
    "offset": 0,
    "startTime": "",
    "statusList": [],
    "storeIdList": [],
    "timeType": 0
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	消息提示	是	[int]	
data		是	[object]	
data>>count	总条数（不再使用，用于兼容其他系统）	是	[long]	
data>>length	每页条数	是	[long]	
data>>list	详细信息列表	是	[array]	
data>>list>>appointmentEndTime	预约结束时间	是	[string]	
data>>list>>appointmentStartTime	预约开始时间	是	[string]	
data>>list>>cargoCode	货件单号	是	[string]	
data>>list>>cargoGoodListResponses		是	[array]	
data>>list>>cargoGoodListResponses>>boxMsgInfo	包裹信息{{序号,件数}}	是	[object]	
data>>list>>cargoGoodListResponses>>deliverVariance	申报差异 申报量-发货量	是	[string]	
data>>list>>cargoGoodListResponses>>id	货件商品id	是	[long]	
data>>list>>cargoGoodListResponses>>msku	MSKU	是	[string]	
data>>list>>cargoGoodListResponses>>mskuCargo	MSKU货号	是	[string]	
data>>list>>cargoGoodListResponses>>picUrl	图片	是	[string]	
data>>list>>cargoGoodListResponses>>platformCode	平台编号	是	[string]	
data>>list>>cargoGoodListResponses>>productId	产品编号	是	[long]	
data>>list>>cargoGoodListResponses>>productName	品名	是	[string]	
data>>list>>cargoGoodListResponses>>purchaseOrderSn	备货单号	是	[string]	
data>>list>>cargoGoodListResponses>>purchaseQuantity	已送货	是	[int]	
data>>list>>cargoGoodListResponses>>shipmentsNum	已发货	是	[long]	
data>>list>>cargoGoodListResponses>>skc	skc	是	[string]	
data>>list>>cargoGoodListResponses>>skcCargo	skc货号	是	[string]	
data>>list>>cargoGoodListResponses>>sku	sku	是	[string]	
data>>list>>cargoGoodListResponses>>skuNum	申报量	是	[int]	
data>>list>>cargoGoodListResponses>>storeId	storeId	是	[long]	
data>>list>>cargoGoodListResponses>>temuCargoCode	temu货件单号	是	[string]	
data>>list>>deliverMethod	配送方式	是	[string]	
data>>list>>deliverPackageNum	已发包裹数	是	[int]	
data>>list>>deliverReceivePackageSkcNum	已发/实收包裹件数	是	[string]	
data>>list>>deliverReceiveSkcNum	已发/实收SKC件数	是	[string]	
data>>list>>deliverSkcNum	已发SKC件数	是	[int]	
data>>list>>deliverTime	发货时间	是	[string]	
data>>list>>deliveryAddress	发货地址	是	[string]	
data>>list>>driverName	司机姓名	是	[string]	
data>>list>>inboundTime	入库时间	是	[string]	
data>>list>>licenseNumber	车牌号码	是	[string]	
data>>list>>logisticsCompany	物流公司	是	[string]	
data>>list>>logisticsCompanyOffer	物流公司(报价)	是	[string]	
data>>list>>logisticsNumber	物流单号	是	[string]	
data>>list>>packageSn	包裹号(多个用分号；分隔)	是	[string]	
data>>list>>phoneNumber	手机号码	是	[string]	
data>>list>>platformCode	平台编号	是	[string]	
data>>list>>purchaseOrderSn	备货单号	是	[string]	
data>>list>>purchaseTime	创建时间	是	[string]	
data>>list>>receiveAddress	收货地址	是	[string]	
data>>list>>receivePackageNum	实收包裹数	是	[int]	
data>>list>>receiveSkcNum	实收SKC件数	是	[int]	
data>>list>>receiveTime	收货时间	是	[string]	
data>>list>>remark	备注	是	[string]	
data>>list>>shippingInfo	发货单信息	是	[array]	
data>>list>>shippingInfo>>shippingListCode	发货单号	是	[string]	
data>>list>>shippingInfo>>shippingNum	发货数量	是	[string]	
data>>list>>shippingInfo>>shippingStatus	发货状态	是	[int]	
data>>list>>shippingInfo>>shippingStatusName	发货状态名称	是	[string]	
data>>list>>shippingListCodeStr	发货单号	是	[string]	
data>>list>>status	状态	是	[int]	
data>>list>>statusStr	状态文本	是	[string]	
data>>list>>storeId	storeId	是	[long]	
data>>list>>storeName	店铺	是	[string]	
data>>list>>subWarehouseId	收货仓库id	是	[string]	
data>>list>>subWarehouseName	实际收货仓	是	[string]	
data>>list>>updateDate	更新时间	是	[string]	
data>>list>>urgencyType	是否是紧急发货单，0-普通 1-急采	是	[int]	
data>>list>>weightDeliveredCargo	发货重量	是	[string]	
data>>total	总条数	是	[long]	
message	消息提示	是	[string]	
request_id	请求链路id	是	[string]	
error_details	消息提示	是	[array]	
返回成功示例
{
    "code": 0,
    "data": {
        "count": 0,
        "length": 0,
        "list": [{
            "appointmentEndTime": "",
            "appointmentStartTime": "",
            "cargoCode": "",
            "cargoGoodListResponses": [{
                "boxMsgInfo": {},
                "deliverVariance": "",
                "id": 0,
                "msku": "",
                "mskuCargo": "",
                "picUrl": "",
                "platformCode": "",
                "productId": 0,
                "productName": "",
                "purchaseOrderSn": "",
                "purchaseQuantity": 0,
                "shipmentsNum": 0,
                "skc": "",
                "skcCargo": "",
                "sku": "",
                "skuNum": 0,
                "storeId": 0,
                "temuCargoCode": ""
            }],
            "deliverMethod": "",
            "deliverPackageNum": 0,
            "deliverReceivePackageSkcNum": "",
            "deliverReceiveSkcNum": "",
            "deliverSkcNum": 0,
            "deliverTime": "",
            "deliveryAddress": "",
            "driverName": "",
            "inboundTime": "",
            "licenseNumber": "",
            "logisticsCompany": "",
            "logisticsCompanyOffer": "",
            "logisticsNumber": "",
            "packageSn": "",
            "phoneNumber": "",
            "platformCode": "",
            "purchaseOrderSn": "",
            "purchaseTime": "",
            "receiveAddress": "",
            "receivePackageNum": 0,
            "receiveSkcNum": 0,
            "receiveTime": "",
            "remark": "",
            "shippingInfo": [{
                "shippingListCode": "",
                "shippingNum": "",
                "shippingStatus": 0,
                "shippingStatusName": ""
            }],
            "shippingListCodeStr": "",
            "status": 0,
            "statusStr": "",
            "storeId": 0,
            "storeName": "",
            "subWarehouseId": "",
            "subWarehouseName": "",
            "updateDate": "",
            "urgencyType": 0,
            "weightDeliveredCargo": ""
        }],
        "total": 0
    },
    "message": "",
    "request_id": "",
    "error_details": []
}
复制
错误
复制成功
 上一章节
查询Temu库存
下一章节 
查询退件地址列表