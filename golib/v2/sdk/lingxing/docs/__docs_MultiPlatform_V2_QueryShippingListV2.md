查询平台仓发货单列表v2

原接口接口仅支持walmart，现已提供v2接口，支持多平台下全部平台的平台仓发货单列表数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/query/shippingList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
platformCodes	平台代码
Walmart 10008
TikTok 10011
Temu 10022
Shein 10027	是	[array]	["10008","10027"]
offset	分页偏移量	否	[int]	0
length	分页长度	否	[int]	50
timeField	时间维度
1 创建时间
2 发货时间
3 开船时间
4 预计到港时间
5 实际妥投时间
6 实际发货时间	否	[int]	1
startTime	开始时间	否	[string]	2024-05-26
endTime	结束时间	否	[string]	2024-05-31
pickingStatus	拣货状态
1 已拣货
0 待拣货	否	[string]	
shippingListStatus	发货单状态
0 待配货
1 待发货
2 已发货
3 已作废	否	[int]	
searchField	搜索维度
1 MSKU
2 发货单号
7 货件单号
8 商品条码	否	[int]	8
searchSingleValue	模糊搜索值	否	[string]	sMM24042281198971
storeIds	店铺id列表，对应查询多平台店铺信息接口对应字段【store_id】	否	[array]	["110000000018027003"]
updateStartTime	修改开始时间	否	[string]	2024-01-01 12:00:00
updateEndTime	修改结束时间	否	[string]	2024-01-01 12:00:00
isDelete	是否删除 0 未删除（默认） 1 已删除	否	[int]	0
请求cURL示例
curl --location --globoff 'https://openapi.lingxing.com/basicOpen/multiplatform/query/shippingList?access_token=value&app_key=value&sign=value&timestamp=value' \
--header 'Content-Type: application/json' \
--data '{
    "searchSingleValue": "sMM24042281198971",
    "searchField": 8,
    "pickingStatus": "",
    "timeField": 1,
    "startTime": "2024-05-26",
    "endTime": "2024-05-31",
    "platformCodes": [
        "10008",
        "10027"
    ],
    "offset": 0,
    "length": 50,
    "storeIds": [
        "110000000018027003"
    ],
    "isDelete": 0
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	31a28511511441f9bac7c0d9e1e4e40f.1730454362823
response_time	响应时间	是	[string]	2024-11-01 17:46:06
data	响应数据	是	[object]	
data>>total	总数	是	[string]	13
data>>list		是	[array]	
data>>list>>id	id	是	[string]	609
data>>list>>shippingListCode	货件单号	是	[string]	SP202405310030028
data>>list>>warehouseId	仓库ID	是	[int]	1
data>>list>>platformCode	平台代码	是	[string]	10027
data>>list>>warehouseName	仓库名	是	[string]	仓库1
data>>list>>logisticsProviderId	物流商id	是	[string]	1
data>>list>>logisticsChannelId	物流渠道ID	是	[string]	1
data>>list>>logisticsProviderName	物流商名	是	[string]	未定义
data>>list>>logisticsChannelName	物流渠道名	是	[string]	物流商1
data>>list>>logisticsTypeName	物流方式	是	[string]	未定义物流商1
data>>list>>createUserId	创建者ID	是	[string]	1
data>>list>>creator	创建者名	是	[string]	0超级管理员01
data>>list>>gmtCreate	创建时间	是	[string]	2024-05-31 15:32
data>>list>>deliveryTime	发货时间	是	[string]	2024-05-31 15:34
data>>list>>arrivalTime	到货时间	是	[string]	2024-06-03 00:00
data>>list>>sailTime	开船时间	是	[string]	
data>>list>>expectedArrivalTime	预计到港时间	是	[string]	
data>>list>>actualDueTime	实际妥投时间	是	[string]	
data>>list>>orderLogisticsStatus	订单物流状态	是	[string]	
data>>list>>actualDeliveryTime	实际发货时间	是	[string]	
data>>list>>logisticsCode	物流中心编码	是	[string]	
data>>list>>shippingListRemark	备注	是	[string]	娟娟0531
data>>list>>shippingListStatus	物流状态
0、待配货
1、待发货
2、已发货
3、已作废	是	[int]	3
data>>list>>shippingListStatusDesc	物流状态名字	是	[string]	已作废
data>>list>>pickingStatus	拣货状态码
-1、未分配状态
0、待拣货
1、已拣货	是	[int]	-1
data>>list>>pickingStatusDesc	拣货状态名
-1、未分配状态
0、待拣货
1、已拣货	是	[string]	
data>>list>>goodExtDetails	商品详细列表	是	[array]	
data>>list>>goodExtDetails>>goodsUrl	图片链接	是	[string]	
data>>list>>goodExtDetails>>platformCode	平台代码	是	[string]	
data>>list>>goodExtDetails>>platformName	平台名	是	[string]	
data>>list>>goodExtDetails>>countryCode	国家代码	是	[string]	
data>>list>>goodExtDetails>>countryName	国家名	是	[string]	
data>>list>>goodExtDetails>>storeId	店铺ID	是	[string]	
data>>list>>goodExtDetails>>storeName	店铺名	是	[string]	
data>>list>>goodExtDetails>>cargoId	货件ID	是	[string]	
data>>list>>goodExtDetails>>cargoCode	货件单号	是	[string]	
data>>list>>goodExtDetails>>msku	MSKU	是	[string]	
data>>list>>goodExtDetails>>mskuId	MSKU ID	是	[string]	
data>>list>>goodExtDetails>>articleNumbering	GTIN	是	[string]	
data>>list>>goodExtDetails>>productId	产品ID	是	[string]	
data>>list>>goodExtDetails>>productName	产品名	是	[string]	
data>>list>>goodExtDetails>>sku	SKU	是	[string]	
data>>list>>goodExtDetails>>applyNum	申报量	是	[string]	
data>>list>>goodExtDetails>>shipmentsNum	发货量	是	[string]	
data>>list>>goodExtDetails>>hasPair	是否配对
ture、配对
false、未配对	是	[boolean]	
data>>list>>goodExtDetails>>productType	产品类型	是	[int]	
data>>list>>logisticsDetails	物流详细列表	是	[array]	
data>>list>>logisticsDetails>>id	ID	是	[string]	978
data>>list>>logisticsDetails>>shippingListId	物流信息列表ID	是	[string]	609
data>>list>>logisticsDetails>>logisticsNumber	物流商单号	是	[string]	
data>>list>>logisticsDetails>>trackingNumber	跟踪号	是	[string]	LRJ0531001
data>>list>>logisticsDetails>>transportationCost	物流费用	是	[string]	3.00
data>>list>>logisticsDetails>>otherCost	其他费用	是	[string]	4.00
data>>list>>logisticsDetails>>transportationCurrency	物流费用货币单位	是	[string]	CNY
data>>list>>logisticsDetails>>otherCurrency	其他费用货币单位	是	[string]	CNY
data>>list>>logisticsDetails>>transportationIcon	物流费用货币单位图标	是	[string]	￥
data>>list>>logisticsDetails>>otherIcon	其他费用货币单位图标	是	[string]	￥
data>>list>>logisticsDetails>>otherCostRemark	其他费用备注	是	[string]	
data>>list>>logisticsDetails>>expectedTransportationCost	预估物流费用	是	[string]	1.00
data>>list>>logisticsDetails>>expectedTransportationCurrency	预估物流费用货币	是	[string]	CNY
data>>list>>logisticsDetails>>expectedTransportationIcon	预估物流费用货币图标	是	[string]	￥
data>>list>>logisticsDetails>>expectedOtherCost	预估其他费用	是	[string]	2.00
data>>list>>logisticsDetails>>expectedOtherCurrency	预估其他费用货币	是	[string]	CNY
data>>list>>logisticsDetails>>expectedOtherIcon	预估其他费用货币图标	是	[string]	￥
data>>list>>logisticsDetails>>expectedOtherCostRemark	预估其他费用备注	是	[string]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "31a28511511441f9bac7c0d9e1e4e40f.1730454362823",
    "response_time": "2024-11-01 17:46:06",
    "data": {
        "total": "1",
        "list": [
            {
                "id": "609",
                "shippingListCode": "SP202405310030028",
                "warehouseId": 1,
                "platformCode": "10027",
                "warehouseName": "仓库1",
                "logisticsProviderId": "1",
                "logisticsChannelId": "1",
                "logisticsProviderName": "未定义",
                "logisticsChannelName": "物流商1",
                "logisticsTypeName": "未定义物流商1",
                "createUserId": "1",
                "creator": "0超级管理员01",
                "gmtCreate": "2024-05-31 15:32",
                "deliveryTime": "2024-05-31 15:34",
                "arrivalTime": "2024-06-03 00:00",
                "sailTime": "",
                "expectedArrivalTime": "",
                "actualDueTime": "",
                "orderLogisticsStatus": "",
                "actualDeliveryTime": "",
                "logisticsCode": "",
                "shippingListRemark": "娟娟0531",
                "shippingListStatus": 3,
                "shippingListStatusDesc": "已作废",
                "pickingStatus": -1,
                "pickingStatusDesc": "",
                "storeId": null,
                "goodExtDetails": [],
                "logisticsDetails": [
                    {
                        "id": "978",
                        "shippingListId": "609",
                        "logisticsNumber": "",
                        "trackingNumber": "LRJ0531001",
                        "transportationCost": "3.00",
                        "otherCost": "4.00",
                        "transportationCurrency": "CNY",
                        "otherCurrency": "CNY",
                        "transportationIcon": "￥",
                        "otherIcon": "￥",
                        "otherCostRemark": "",
                        "expectedTransportationCost": "1.00",
                        "expectedTransportationCurrency": "CNY",
                        "expectedTransportationIcon": "￥",
                        "expectedOtherCost": "2.00",
                        "expectedOtherCurrency": "CNY",
                        "expectedOtherIcon": "￥",
                        "expectedOtherCostRemark": ""
                    }
                ]
            }       
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询平台仓发货单列表
下一章节 
修改平台仓发货单备注