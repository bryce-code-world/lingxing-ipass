创建AWD入库任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-plan/createInboundPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
awdDeliveredGoodsBOS	发货商品	是	[array]	
awdDeliveredGoodsBOS>>boxQuantity	箱数（只能是正整数）	是	[string]	5
awdDeliveredGoodsBOS>>expiration	有效期（yyyy-MM-dd）	否	[string]	2025-12-31
awdDeliveredGoodsBOS>>height	箱子高（2位小数）	否	[number]	12.34
awdDeliveredGoodsBOS>>labelOwner	标签类型（AMAZON,SELF）	否	[string]	SELF
awdDeliveredGoodsBOS>>length	箱子长（2位小数）	是	[number]	20.50
awdDeliveredGoodsBOS>>lengthUnit	长度单位(INCHES-英制,CENTIMETERS-公制)	否	[string]	INCHES
awdDeliveredGoodsBOS>>msku	msku	是	[string]	ABC123456
awdDeliveredGoodsBOS>>prepCategory	预处理类别：ADULT：成人；HANGER：悬挂在衣架上的服装；TEXTILE：服装、面料、毛绒玩具和纺织品；BABY：母婴用品；FRAGILE：易碎品；LIQUID：液体（未存放在玻璃容器中）；PERFORATED：打孔包装；GRANULAR：粉末、球状或颗粒状物品；SHARP：尖利物品；SMALL：小号；SET：套装销售；NONE：无需进行预处理；FC_Provided：亚马逊指定分类（创建时不可主动选择，创建AWD入库任务后若亚马逊接口有要求才可选择。该选项代表预处理分类由亚马逊指定，不可修改。您可前往后台，在创建STA任务页面该商品的包装详情处，查看具体分类名称。）	否	[string]	FRAGILE
awdDeliveredGoodsBOS>>prepOwner	预处理提供方（AMAZON,SELF）	否	[string]	SELF
awdDeliveredGoodsBOS>>quantityInBox	单箱数量（只能是正整数）	是	[string]	10
awdDeliveredGoodsBOS>>weight	箱子重量（2位小数）	是	[number]	15.75
awdDeliveredGoodsBOS>>weightUnit	重量单位（POUNDS-磅，KILOGRAMS-千克	否	[string]	POUNDS
awdDeliveredGoodsBOS>>width	箱子宽（2位小数）	是	[number]	15.25
awdShippingAddressBO	发货地址	是	[object]	
awdShippingAddressBO>>addressLine1	发货地址-详细街道地址1	是	[string]	123 Main St
awdShippingAddressBO>>addressLine2	发货地址-详细街道地址2	否	[string]	Suite 100
awdShippingAddressBO>>city	发货地址-城市	是	[string]	Seattle
awdShippingAddressBO>>countryCode	发货地址-国家(地区）	是	[string]	US
awdShippingAddressBO>>phoneNumber	发货地址-电话号码	是	[string]	13800138000
awdShippingAddressBO>>postalCode	发货地址-邮箱编码	是	[string]	98101
awdShippingAddressBO>>shipperName	发货地址-发货方名称	是	[string]	ABC Company
awdShippingAddressBO>>stateOrProvinceCode	发货地址-州/省/地区编码	是	[string]	WA
awdShippingAddressBO>>zone	发货地址-区	否	[string]	Downtown
destinationRegion	地区偏好：us-east：美国东海岸（马里兰州和宾夕法尼亚分拨中心）；us-west：美国西海岸（加利福尼亚州分拨中心）；us-southcentral：美国中南部（德克萨斯州分拨中心）；us-southeast：美国东南部（南卡罗来纳州分拨中心）；null：亚马逊分配（亚马逊将为您的货件分配最佳分拨中心）	否	[string]	us-west
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	1234567890
请求示例

json

{
    "awdDeliveredGoodsBOS": [
        {
            "boxQuantity": "5",
            "expiration": "2025-12-31",
            "height": 12.34,
            "labelOwner": "SELF",
            "length": 20.50,
            "lengthUnit": "INCHES",
            "msku": "ABC123456",
            "prepCategory": "FRAGILE",
            "prepOwner": "SELF",
            "quantityInBox": "10",
            "weight": 15.75,
            "weightUnit": "POUNDS",
            "width": 15.25
        }
    ],
    "awdShippingAddressBO": {
        "addressLine1": "123 Main St",
        "addressLine2": "Suite 100",
        "city": "Seattle",
        "countryCode": "US",
        "phoneNumber": "13800138000",
        "postalCode": "98101",
        "shipperName": "ABC Company",
        "stateOrProvinceCode": "WA",
        "zone": "Downtown"
    },
    "destinationRegion": "us-west",
    "sid": 1234567890
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码,0成功	否	[int]	0
data	数据	否	[object]	
data>>orderId	AWD任务编号	否	[string]	AWD123456789
error_details	错误信息	否	[array]	[]
message	提示信息	否	[string]	操作成功
request_id	请求链路id	否	[string]	a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8
response_time	响应时间	否	[string]	2024-01-01 12:00:00
返回成功示例

json

{
    "code": 0,
    "message": "操作成功",
    "data": {
        "orderId": "AWD123456789"
    },
    "error_details": [],
    "request_id": "a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8",
    "response_time": "2024-01-01 12:00:00"
}
复制
错误
复制成功
 上一章节
确认AWD入库任务
FBA
下一章节 
查询AWD入库任务详情
FBA