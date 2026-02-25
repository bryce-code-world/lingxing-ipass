查询AWD入库货件列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-shipment/page	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
page	分页页码	是	[long]	1
dateType	时间类型 1:创建 2更新	是	[int]	1
endDateTime	结束时间，格式：YYYY-MM-DD 双闭区间	是	[datetime]	2023-07-31
shipmentId	货件单号	否	[string]	AWD20230715001
sidList	店铺id列表	否	[array]	[123456789, 987654321]
length	分页大小，上限	是	[long]	10
startDateTime	开始时间，格式：YYYY-MM-DD 双闭区间	是	[datetime]	2023-07-01
statusList	任务状态：CREATED：已创建；SHIPPED：已发货；IN_TRANSIT：运输中；RECEIVING：接收中；DELIVERED：已送达；CLOSED：已关闭；CANCELLED：已取消	否	[array]	["CREATED", "IN_TRANSIT"]
请求示例

json

{
  "page": 1,
  "dateType": 1,
  "endDateTime": "2023-07-31T23:59:59Z",
  "shipmentId": "AWD20230715001",
  "sidList": [123456789, 987654321],
  "length": 10,
  "startDateTime": "2023-07-01T00:00:00Z",
  "statusList": ["CREATED", "IN_TRANSIT"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码,0成功	否	[int]	
data		否	[object]	
data>>countId		否	[string]	
data>>current		否	[long]	
data>>maxLimit		否	[long]	
data>>optimizeCountSql		否	[boolean]	
data>>orders		否	[array]	
data>>orders>>asc		否	[boolean]	
data>>orders>>column		否	[string]	
data>>pages		否	[long]	
data>>records		否	[array]	
data>>records>>awdDeliveredGoodsBO	配送地址	否	[object]	
data>>records>>awdDeliveredGoodsBO>>destinationAddressLine1	配送地址-详细街道地址1	否	[string]	
data>>records>>awdDeliveredGoodsBO>>destinationAddressLine2	配送地址-详细街道地址2	否	[string]	
data>>records>>awdDeliveredGoodsBO>>destinationCity	配送地址-城市	否	[string]	
data>>records>>awdDeliveredGoodsBO>>destinationCountryCode	配送地址-国家（地区）	否	[string]	
data>>records>>awdDeliveredGoodsBO>>destinationPostalCode	配送地址-邮政编码	否	[string]	
data>>records>>awdDeliveredGoodsBO>>destinationStateOrRegion	配送地址-州/省/地区编码	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS	发货商品	否	[array]	
data>>records>>awdDeliveredGoodsItemBOS>>asin	asin	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>boxQuantity	箱数	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>expiration	有效期	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>fnsku	fnsku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>height	箱子高	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>records>>awdShippingAddressBO	发货地址	否	[object]	
data>>records>>awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	否	[string]	
data>>records>>awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	
data>>records>>awdShippingAddressBO>>city	发货地址-城市	否	[string]	
data>>records>>awdShippingAddressBO>>countryCode	发货地址-国家(地区）	否	[string]	
data>>records>>awdShippingAddressBO>>phoneNumber	发货地址-电话号码	否	[string]	
data>>records>>awdShippingAddressBO>>postalCode	发货地址-邮箱编码	否	[string]	
data>>records>>awdShippingAddressBO>>shipperName	发货地址-发货方名称	否	[string]	
data>>records>>awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	否	[string]	
data>>records>>awdShippingAddressBO>>zone	发货地址-区	否	[string]	
data>>records>>carrierCode	承运方式	否	[string]	
data>>records>>createByName	创建人名称	否	[string]	
data>>records>>gmtCreate	创建时间	否	[string]	
data>>records>>gmtModified	更新时间	否	[string]	
data>>records>>orderId	AWD入库任务号	否	[string]	
data>>records>>remark	备注	否	[string]	
data>>records>>shipBy	发货时间	否	[string]	
data>>records>>sid	领星店铺ID	否	[long]	
data>>records>>status	货件状态：CREATED：已创建；SHIPPED：已发货；IN_TRANSIT：运输中；RECEIVING：接收中；DELIVERED：已送达；CLOSED：已关闭；CANCELLED：已取消	否	[string]	
data>>records>>trackingId	跟踪编码	否	[string]	
data>>records>>warehouseReferenceId	物流中心编码	否	[string]	
data>>searchCount		否	[boolean]	
data>>size		否	[long]	
data>>total		否	[long]	
message	提示信息	否	[string]	
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[array]	
response_time	响应时间	否	[string]	
返回成功示例

json

{
  "code": 0,
  "data": {
    "countId": null,
    "current": 1,
    "maxLimit": 1000,
    "optimizeCountSql": true,
    "orders": [
      {
        "asc": false,
        "column": "gmtCreate"
      }
    ],
    "pages": 1,
    "records": [
      {
        "awdDeliveredGoodsBO": {
          "destinationAddressLine1": "456 Distribution Ave",
          "destinationAddressLine2": "Building 3",
          "destinationCity": "Phoenix",
          "destinationCountryCode": "US",
          "destinationPostalCode": "85001",
          "destinationStateOrRegion": "AZ"
        },
        "awdDeliveredGoodsItemBOS": [
          {
            "asin": "B0ABCD1234",
            "boxQuantity": "5",
            "expiration": "2025-12-31",
            "fnsku": "X000123ABC",
            "height": 10.5,
            "labelOwner": "SELF",
            "length": 15.75,
            "lengthUnit": "INCHES",
            "msku": "ABC123",
            "parentAsin": "B0PARENT12",
            "prepCategory": "NONE",
            "prepOwner": "SELF",
            "productName": "Example Product",
            "quantityInBox": "20",
            "sku": "SKU12345",
            "title": "Example Product Title",
            "url": "https://example.com/product.jpg",
            "weight": 12.3,
            "weightUnit": "POUNDS",
            "width": 8.25
          }
        ],
        "awdShippingAddressBO": {
          "addressLine1": "123 Main St",
          "addressLine2": "Suite 100",
          "city": "Seattle",
          "countryCode": "US",
          "phoneNumber": "1234567890",
          "postalCode": "98101",
          "shipperName": "Example Shipper",
          "stateOrProvinceCode": "WA",
          "zone": "Downtown"
        },
        "carrierCode": "UPS",
        "createByName": "API User",
        "gmtCreate": "2023-07-15T10:00:00Z",
        "gmtModified": "2023-07-15T14:30:00Z",
        "orderId": "AWD20230715001",
        "remark": "Fragile items, handle with care",
        "shipBy": "2023-07-20",
        "sid": 123456789,
        "status": "SHIPPED",
        "trackingId": "1Z1234567890123456",
        "warehouseReferenceId": "PHX5"
      }
    ],
    "searchCount": true,
    "size": 10,
    "total": 1
  },
  "message": "Success",
  "error_details": [],
  "request_id": "c3d4e5f6-g7h8-i9j0-k1l2-m3n4o5p6q7r8",
  "response_time": "2023-07-15T16:30:45Z"
}
复制
错误
复制成功
 上一章节
查询AWD入库货件详情
FBA
下一章节 
更新AWD货件跟踪编号
FBA