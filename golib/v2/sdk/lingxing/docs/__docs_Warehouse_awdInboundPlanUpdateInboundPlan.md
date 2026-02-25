更新AWD入库任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-plan/updateInboundPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
awdDeliveredGoodsBOS	发货商品	是	[array]	
awdDeliveredGoodsBOS>>boxQuantity	箱数（只能是正整数）	是	[string]	
awdDeliveredGoodsBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	
awdDeliveredGoodsBOS>>height	箱子高（2位小数）	是	[number]	
awdDeliveredGoodsBOS>>labelOwner	标签类型（AMAZON,SELF）	否	[string]	
awdDeliveredGoodsBOS>>length	箱子长（2位小数）	是	[number]	
awdDeliveredGoodsBOS>>lengthUnit	长度单位(INCHES-英制,CENTIMETERS-公制)	是	[string]	
awdDeliveredGoodsBOS>>msku	msku	是	[string]	
awdDeliveredGoodsBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	
awdDeliveredGoodsBOS>>prepOwner	预处理提供方（AMAZON,SELF）	否	[string]	
awdDeliveredGoodsBOS>>quantityInBox	单箱数量（只能是正整数）	是	[string]	
awdDeliveredGoodsBOS>>weight	箱子重量（2位小数）	是	[number]	
awdDeliveredGoodsBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克）	是	[string]	
awdDeliveredGoodsBOS>>width	箱子宽（2位小数）	是	[number]	
awdShippingAddressBO	发货地址	是	[object]	
awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	是	[string]	
awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	
awdShippingAddressBO>>city	发货地址-城市	是	[string]	
awdShippingAddressBO>>countryCode	发货地址-国家(地区）	是	[string]	
awdShippingAddressBO>>phoneNumber	发货地址-电话号码	是	[string]	
awdShippingAddressBO>>postalCode	发货地址-邮箱编码	是	[string]	
awdShippingAddressBO>>shipperName	发货地址-发货方名称	是	[string]	
awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	是	[string]	
awdShippingAddressBO>>zone	发货地址-区	否	[string]	
createBy	创建人id，默认API账号id	否	[string]	
destinationRegion	地区偏好：us-east：美国东海岸（马里兰州和宾夕法尼亚分拨中心）；us-west：美国西海岸（加利福尼亚州分拨中心）；us-southcentral：美国中南部（德克萨斯州分拨中心）；us-southeast：美国东南部（南卡罗来纳州分拨中心）；null：亚马逊分配（亚马逊将为您的货件分配最佳分拨中心）	否	[string]	
orderId	STA任务编号	是	[string]	
remark	备注	否	[string]	
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
请求示例

json

{
  "awdDeliveredGoodsBOS": [
    {
      "boxQuantity": "5",
      "expiration": "2025-12-31",
      "height": 10.50,
      "labelOwner": "SELF",
      "length": 15.75,
      "lengthUnit": "INCHES",
      "msku": "ABC123",
      "prepCategory": "NONE",
      "prepOwner": "SELF",
      "quantityInBox": "20",
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
  "createBy": "api_user_123",
  "destinationRegion": "us-west",
  "orderId": "STA20230715001",
  "remark": "Fragile items, handle with care",
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
data>>orderId	AWD任务编号	否	[string]	
message	提示信息	否	[string]	
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[array]	
response_time	响应时间	否	[string]	
返回成功示例

json

{
  "code": 0,
  "data": {
    "orderId": "AWD20230715001"
  },
  "message": "Success",
  "request_id": "a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6",
  "response_time": "2023-07-15T14:30:45Z"
}
复制
错误
复制成功
 上一章节
查询AWD入库任务列表
FBA
下一章节 
查询AWD入库货件详情
FBA