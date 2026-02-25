查询Temu库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/fbt/stockSearch	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
length	每页条数	否	[long]	20
offset	偏移量	否	[long]	0
storeIdList	店铺Id集合	是	[array]	[]
请求示例
{
    "length": 20,
    "offset": 0,
    "storeIdList": []
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	
data		是	[object]	
data>>list		是	[array]	
data>>list>>adviceQuantity	建议备货量	是	[int]	
data>>list>>availableSaleDays	可售天数	是	[int]	
data>>list>>childList	子行数据	是	[array]	
data>>list>>companyId	企业ID	是	[long]	
data>>list>>id	id	是	[long]	
data>>list>>inventorySaleDays	库存可售天数	是	[int]	
data>>list>>lackQuantity	缺货量	是	[int]	
data>>list>>lastSevenDaysSaleVolume	最近七天销量	是	[int]	
data>>list>>lastThirtyDaysSaleVolume	最近三十天天销量	是	[int]	
data>>list>>mskuList	msku	是	[array]	
data>>list>>mskuList>>msku	msku	是	[string]	
data>>list>>mskuList>>mskuCode	msku货号	是	[string]	
data>>list>>parentId	父行id	是	[long]	
data>>list>>pic	图片	是	[string]	
data>>list>>propertyValue	属性	是	[array]	
data>>list>>quantity	总库存	是	[int]	
data>>list>>skc	skc	是	[string]	
data>>list>>skcCode	skc货号	是	[string]	
data>>list>>skuList	sku	是	[array]	
data>>list>>skuList>>pid	pid	是	[long]	
data>>list>>skuList>>productName	品名	是	[string]	
data>>list>>skuList>>sku	sku	是	[string]	
data>>list>>spu	spu	是	[string]	
data>>list>>storeId	店铺ID	是	[long]	
data>>list>>storeName	店铺名称	是	[string]	
data>>list>>todaySaleVolume	今日销量	是	[int]	
data>>list>>waitApproveInventoryNum	待审核备货	是	[int]	
data>>list>>waitDeliveryInventoryNum	待发货	是	[int]	
data>>list>>waitOnShelfNum	仓内暂不可用	是	[int]	
data>>list>>waitReceiveNum	已发货	是	[int]	
data>>list>>warehouseInventoryNum	仓内可用	是	[int]	
data>>list>>warehouseSaleDays	仓内库存可售天数	是	[int]	
data>>total		是	[long]	
message	消息提示	是	[string]	
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	
返回成功示例
{
    "code": 0,
    "data": {
        "list": [{
            "adviceQuantity": 0,
            "availableSaleDays": 0,
            "childList": [],
            "companyId": 0,
            "id": 0,
            "inventorySaleDays": 0,
            "lackQuantity": 0,
            "lastSevenDaysSaleVolume": 0,
            "lastThirtyDaysSaleVolume": 0,
            "mskuList": [{
                "msku": "",
                "mskuCode": ""
            }],
            "parentId": 0,
            "pic": "",
            "propertyValue": [],
            "quantity": 0,
            "skc": "",
            "skcCode": "",
            "skuList": [{
                "pid": 0,
                "productName": "",
                "sku": ""
            }],
            "spu": "",
            "storeId": 0,
            "storeName": "",
            "todaySaleVolume": 0,
            "waitApproveInventoryNum": 0,
            "waitDeliveryInventoryNum": 0,
            "waitOnShelfNum": 0,
            "waitReceiveNum": 0,
            "warehouseInventoryNum": 0,
            "warehouseSaleDays": 0
        }],
        "total": 0
    },
    "message": "",
    "request_id": "",
    "response_time": ""
}
复制
错误
复制成功
 上一章节
查询WFS库存列表
下一章节 
查询Temu货件