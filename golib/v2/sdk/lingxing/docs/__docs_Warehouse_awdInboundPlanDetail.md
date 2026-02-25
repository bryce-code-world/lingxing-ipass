查询AWD入库任务详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-plan/detail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderId	STA任务编号	是	[string]	STA202400001
sid	领星店铺ID 列表，对应查询亚马逊店铺列表	是	[long]	1234567890
请求示例
{
    "orderId": "STA202400001",
    "sid": 1234567890
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
data>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF）	否	[string]	
data>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（（POUNDS-磅，KILOGRAMS-千克））	否	[string]	
data>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>awdShipmentVOS	AWD货件	否	[array]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO	配送地址	否	[object]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationAddressLine1	配送地址-详细街道地址1	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationAddressLine2	配送地址-详细街道地址2	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationCity	配送地址-城市	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationCountryCode	配送地址-国家（地区）	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationPostalCode	配送地址-邮政编码	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsBO>>destinationStateOrRegion	配送地址-州/省/地区编码	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS	发货商品	否	[array]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>asin	asin	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>boxQuantity	箱数	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>fnsku	fnsku	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>height	箱子高	否	[number]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>labelOwner	标签类型（AMAZON,SELF）	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>length	箱子长	否	[number]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>lengthUnit	长度单位(INCHES-英制, CENTIMETERS-公制)	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>msku	msku	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>parentAsin	父asin	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>prepOwner	预处理提供方（AMAZON,SELF）	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>productName	品名	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>quantityInBox	单箱数量	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>sku	sku	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>title	标题	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>url	图片url	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>weight	箱子重量	否	[number]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>weightUnit	重量单位（（POUNDS-磅，KILOGRAMS-千克））	否	[string]	
data>>awdShipmentVOS>>awdDeliveredGoodsItemBOS>>width	箱子宽	否	[number]	
data>>awdShipmentVOS>>awdShippingAddressBO	发货地址	否	[object]	
data>>awdShipmentVOS>>awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>city	发货地址-城市	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>countryCode	发货地址-国家(地区）	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>phoneNumber	发货地址-电话号码	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>postalCode	发货地址-邮箱编码	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>shipperName	发货地址-发货方名称	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	否	[string]	
data>>awdShipmentVOS>>awdShippingAddressBO>>zone	发货地址-区	否	[string]	
data>>awdShipmentVOS>>carrierCode	承运方式	否	[string]	
data>>awdShipmentVOS>>createByName	创建人名称	否	[string]	
data>>awdShipmentVOS>>gmtCreate	创建时间	否	[string]	
data>>awdShipmentVOS>>gmtModified	更新时间	否	[string]	
data>>awdShipmentVOS>>orderId	AWD入库任务号	否	[string]	
data>>awdShipmentVOS>>remark	备注	否	[string]	
data>>awdShipmentVOS>>shipBy	发货时间	否	[string]	
data>>awdShipmentVOS>>sid	领星店铺ID	否	[long]	
data>>awdShipmentVOS>>status	货件状态：CREATED：已创建；SHIPPED：已发货；IN_TRANSIT：运输中；RECEIVING：接收中；DELIVERED：已送达；CLOSED：已关闭；CANCELLED：已取消	否	[string]	
data>>awdShipmentVOS>>trackingId	跟踪编码	否	[string]	
data>>awdShipmentVOS>>warehouseReferenceId	物流中心编码	否	[string]	
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
data>>createByName	创建人名称	否	[string]	
data>>destinationRegion	地区偏好：us-east：美国东海岸（马里兰州和宾夕法尼亚分拨中心）；us-west：美国西海岸（加利福尼亚州分拨中心）；us-southcentral：美国中南部（德克萨斯州分拨中心）；us-southeast：美国东南部（南卡罗来纳州分拨中心）；null：亚马逊分配（亚马逊将为您的货件分配最佳分拨中心）	否	[string]	
data>>gmtCreate	创建时间（yyyy-MM-dd HH:mm:ss）	否	[string]	
data>>gmtModified	更新时间（yyyy-MM-dd HH:mm:ss）	否	[string]	
data>>orderId	AWD入库任务号	否	[string]	
data>>remark	备注	否	[string]	
data>>sid	店铺id	否	[long]	
data>>status	任务状态：LOCALDRAFT：草稿；DRAFT：待确认；VALIDATING：更新中；CONFIRMED：已确认；CLOSED： 已关闭；EXPIRED：已过期；CANCELLED：已取消	否	[string]	
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
    },
    "error_details": [],
    "request_id": "req_1234567890abcdef",
    "response_time": "2024-01-01T10:10:00Z"
}
复制
错误
复制成功
 上一章节
创建AWD入库任务
FBA
下一章节 
查询AWD入库任务列表
FBA