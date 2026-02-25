查询AWD入库货件详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-shipment/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipmentId	AWD入库货件单号	是	[string]	AWD20230715001
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	123456789
请求示例

json

{
  "shipmentId": "AWD20230715001",
  "sid": 123456789
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码,0成功	否	[int]	
data		否	[object]	
data>>awdDeliveredGoodsBO	配送地址	否	[object]	
data>>awdDeliveredGoodsBO>>destinationAddressLine1	配送地址-详细街道地址1	否	[string]	
data>>awdDeliveredGoodsBO>>destinationAddressLine2	配送地址-详细街道地址2	否	[string]	
data>>awdDeliveredGoodsBO>>destinationCity	配送地址-城市	否	[string]	
data>>awdDeliveredGoodsBO>>destinationCountryCode	配送地址-国家（地区）	否	[string]	
data>>awdDeliveredGoodsBO>>destinationPostalCode	配送地址-邮政编码	否	[string]	
data>>awdDeliveredGoodsBO>>destinationStateOrRegion	配送地址-州/省/地区编码	否	[string]	
data>>awdDeliveredGoodsItemBOS	发货商品	否	[array]	
data>>awdDeliveredGoodsItemBOS>>asin	asin	否	[string]	
data>>awdDeliveredGoodsItemBOS>>boxQuantity	箱数	否	[string]	
data>>awdDeliveredGoodsItemBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>fnsku	fnsku	否	[string]	
data>>awdDeliveredGoodsItemBOS>>height	箱子高	否	[number]	
data>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF,NULL）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF,NULL）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>awdShippingAddressBO	发货地址	否	[object]	
data>>awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	否	[string]	
data>>awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	
data>>awdShippingAddressBO>>city	发货地址-城市	否	[string]	
data>>awdShippingAddressBO>>countryCode	发货地址-国家(地区）	否	[string]	
data>>awdShippingAddressBO>>phoneNumber	发货地址-电话号码	否	[string]	
data>>awdShippingAddressBO>>postalCode	发货地址-邮箱编码	否	[string]	
data>>awdShippingAddressBO>>shipperName	发货地址-发货方名称	否	[string]	
data>>awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	否	[string]	
data>>awdShippingAddressBO>>zone	发货地址-区	否	[string]	
data>>carrierCode	承运方式	否	[string]	
data>>createByName	创建人名称	否	[string]	
data>>gmtCreate	创建时间	否	[string]	
data>>gmtModified	更新时间	否	[string]	
data>>orderId	AWD入库任务号	否	[string]	
data>>remark	备注	否	[string]	
data>>shipBy	发货时间	否	[string]	
data>>sid	领星店铺ID	否	[long]	
data>>status	货件状态：CREATED：已创建；SHIPPED：已发货；IN_TRANSIT：运输中；RECEIVING：接收中；DELIVERED：已送达；CLOSED：已关闭；CANCELLED：已取消	否	[string]	
data>>trackingId	跟踪编码	否	[string]	
data>>warehouseReferenceId	物流中心编码	否	[string]	
message	提示信息	否	[string]	
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[array]	
response_time	响应时间	否	[string]	
返回成功示例

json

{
  "code": 0,
  "data": {
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
        "height": 10.50,
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
        "weight": 12.30,
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
    "status": "IN_TRANSIT",
    "trackingId": "1Z1234567890123456",
    "warehouseReferenceId": "PHX5"
  },
  "message": "Success",
  "error_details": [],
  "request_id": "b2c3d4e5-f6g7-h8i9-j0k1-l2m3n4o5p6q7",
  "response_time": "2023-07-15T15:45:30Z"
}
复制
错误
复制成功
 上一章节
更新AWD入库任务
FBA
下一章节 
查询AWD入库货件列表
FBA