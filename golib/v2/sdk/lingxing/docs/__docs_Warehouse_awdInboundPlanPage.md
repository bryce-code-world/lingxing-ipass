查询AWD入库任务列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-plan/page	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
page	分页页码	是	[long]	1
dateType	时间类型 1:创建 2更新	是	[int]	1
endDateTime	结束时间，格式：YYYY-MM-DD 双闭区间	是	[datetime]	2025-07-15
orderId	awd入库任务编号	否	[string]	12345
shipmentId	awd货件单号	否	[string]	ABC123
sidList	店铺id列表	否	[array]	[1, 2, 3]
length	分页大小，上限	是	[long]	10
startDateTime	开始时间，格式：YYYY-MM-DD 双闭区间	是	[datetime]	2025-07-01
statusList	任务状态：LOCALDRAFT：草稿；DRAFT：待确认；VALIDATING：更新中；CONFIRMED：已确认；CLOSED： 已关闭；EXPIRED：已过期；CANCELLED：已取消	否	[array]	["DRAFT"]
请求示例
{
  "page": 1,
  "dateType": 1,
  "endDateTime": "2025-07-15",
  "orderId": "12345",
  "shipmentId": "ABC123",
  "sidList": [1, 2, 3],
  "length": 10,
  "sortField": "date",
  "sortType": "asc",
  "startDateTime": "2025-07-01",
  "statusList": ["DRAFT"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码,0成功	否	[int]	
data	数据	否	[object]	
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
data>>records>>awdDeliveredGoodsItemBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>fnsku	fnsku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>height	箱子高	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF,NULL）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF,NULL）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>records>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克）	否	[string]	
data>>records>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>records>>awdShipmentVOS	AWD货件	否	[array]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO	配送地址	否	[object]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationAddressLine1	配送地址-详细街道地址1	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationAddressLine2	配送地址-详细街道地址2	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationCity	配送地址-城市	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationCountryCode	配送地址-国家（地区）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationPostalCode	配送地址-邮政编码	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationStateOrRegion	配送地址-州/省/地区编码	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS	发货商品	否	[array]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>asin	asin	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>boxQuantity	箱数	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>fnsku	fnsku	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>height	箱子高	否	[number]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF,NULL）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF,NULL）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克）	否	[string]	
data>>records>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO	发货地址	否	[object]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>city	发货地址-城市	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>countryCode	发货地址-国家(地区）	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>phoneNumber	发货地址-电话号码	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>postalCode	发货地址-邮箱编码	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>shipperName	发货地址-发货方名称	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	否	[string]	
data>>records>>awdShipmentVOS>>awdShippingAddressBO>>zone	发货地址-区	否	[string]	
data>>records>>awdShipmentVOS>>carrierCode	承运方式	否	[string]	
data>>records>>awdShipmentVOS>>createByName	创建人名称	否	[string]	
data>>records>>awdShipmentVOS>>gmtCreate	创建时间	否	[string]	
data>>records>>awdShipmentVOS>>gmtModified	更新时间	否	[string]	
data>>records>>awdShipmentVOS>>orderId	AWD入库任务号	否	[string]	
data>>records>>awdShipmentVOS>>remark	备注	否	[string]	
data>>records>>awdShipmentVOS>>shipBy	发货时间	否	[string]	
data>>records>>awdShipmentVOS>>sid	领星店铺ID	否	[long]	
data>>records>>awdShipmentVOS>>status	货件状态：CREATED：已创建；SHIPPED：已发货；IN_TRANSIT：运输中；RECEIVING：接收中；DELIVERED：已送达；CLOSED：已关闭；CANCELLED：已取消	否	[string]	
data>>records>>awdShipmentVOS>>trackingId	跟踪编码	否	[string]	
data>>records>>awdShipmentVOS>>warehouseReferenceId	物流中心编码	否	[string]	
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
data>>records>>createByName	创建人名称	否	[string]	
data>>records>>destinationRegion	地区偏好：us-east：美国东海岸（马里兰州和宾夕法尼亚分拨中心）；us-west：美国西海岸（加利福尼亚州分拨中心）；us-southcentral：美国中南部（德克萨斯州分拨中心）；us-southeast：美国东南部（南卡罗来纳州分拨中心）；null：亚马逊分配（亚马逊将为您的货件分配最佳分拨中心）	否	[string]	
data>>records>>gmtCreate	创建时间	否	[string]	
data>>records>>gmtModified	更新时间	否	[string]	
data>>records>>orderId	AWD入库任务号	否	[string]	
data>>records>>remark	备注	否	[string]	
data>>records>>sid	店铺id	否	[long]	
data>>records>>status	任务状态：LOCALDRAFT：草稿；DRAFT：待确认；VALIDATING：更新中；CONFIRMED：已确认；CLOSED： 已关闭；EXPIRED：已过期；CANCELLED：已取消	否	[string]	
data>>searchCount		否	[boolean]	
data>>size		否	[long]	
data>>total	总数	否	[long]	
message	提示信息	否	[string]	
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[array]	
response_time	响应时间	否	[string]	
返回成功示例

json

{
    "code": 0,
    "message": "success",
    "data": {
        "records": [
            {
                "orderId": "AWD202400001",
                "status": "CONFIRMED",
                "createByName": "张三",
                "gmtCreate": "2024-01-01 10:00:00",
                "gmtModified": "2024-01-01 10:05:00",
                "remark": "标准发货",
                "destinationRegion": "us-west",
                "awdShippingAddressBO": {
                    "addressLine1": "123 Main St",
                    "city": "Seattle",
                    "countryCode": "US",
                    "phoneNumber": "13800138000",
                    "postalCode": "98101",
                    "shipperName": "ABC Company",
                    "stateOrProvinceCode": "WA"
                },
                "awdDeliveredGoodsBO": {
                    "destinationAddressLine1": "789 FBA Way",
                    "destinationCity": "San Francisco",
                    "destinationCountryCode": "US",
                    "destinationPostalCode": "94103",
                    "destinationStateOrRegion": "CA"
                },
                "awdDeliveredGoodsItemBOS": [
                    {
                        "msku": "ABC123456",
                        "fnsku": "X00012345",
                        "asin": "B08N5KWB9H",
                        "productName": "Wireless Headphones",
                        "boxQuantity": "5",
                        "quantityInBox": "10",
                        "length": 20.5,
                        "width": 15.25,
                        "height": 12.34,
                        "weight": 15.75,
                        "lengthUnit": "INCHES",
                        "weightUnit": "POUNDS",
                        "prepCategory": "FRAGILE",
                        "prepOwner": "SELF",
                        "labelOwner": "SELF"
                    }
                ],
                "awdShipmentVOS": [
                    {
                        "orderId": "AWD202400001",
                        "status": "SHIPPED",
                        "trackingId": "FBA123456789US",
                        "warehouseReferenceId": "SFO1",
                        "carrierCode": "UPS",
                        "shipBy": "2024-01-05",
                        "awdShippingAddressBO": {
                            "addressLine1": "456 Ship St",
                            "city": "Portland",
                            "countryCode": "US",
                            "phoneNumber": "13800138001",
                            "postalCode": "97201",
                            "shipperName": "XYZ Logistics",
                            "stateOrProvinceCode": "OR"
                        },
                        "awdDeliveredGoodsBO": {
                            "destinationAddressLine1": "789 FBA Way",
                            "destinationCity": "San Francisco",
                            "destinationCountryCode": "US",
                            "destinationPostalCode": "94103",
                            "destinationStateOrRegion": "CA"
                        },
                        "awdDeliveredGoodsItemBOS": [
                            {
                                "msku": "ABC123456",
                                "fnsku": "X00012345",
                                "asin": "B08N5KWB9H",
                                "productName": "Wireless Headphones",
                                "boxQuantity": "5",
                                "quantityInBox": "10",
                                "length": 20.5,
                                "width": 15.25,
                                "height": 12.34,
                                "weight": 15.75,
                                "lengthUnit": "INCHES",
                                "weightUnit": "POUNDS",
                                "prepCategory": "FRAGILE",
                                "prepOwner": "SELF",
                                "labelOwner": "SELF"
                            }
                        ]
                    }
                ]
            }
        ],
        "total": 1,
        "size": 10,
        "current": 1,
        "pages": 1,
        "searchCount": true,
        "optimizeCountSql": true,
        "countId": null,
        "orders": [
            {
                "column": "gmtCreate",
                "asc": false
            }
        ]
    },
    "error_details": [],
    "request_id": "req_1234567890abcdef",
    "response_time": "2024-01-01T10:10:00Z"
}
复制
错误
复制成功
 上一章节
查询AWD入库任务详情
FBA
下一章节 
更新AWD入库任务
FBA