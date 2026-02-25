查询WFS库存列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/warehouse/api/openApi/queryWFSInventionPage	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	是	[array]	["110418202566107648"]
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15，上限200	否	[int]	15
请求示例
{
    "store_id": [
        "110418202566107648"
    ],
    "offset": 0,
    "length": 15
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	1979122130ba45e984cbc9b2e571d86f.1690513512438
response_time	响应时间	是	[string]	2023-07-28 11:05:12
data	响应数据	是	[object]	
data>>records	列表数据	是	[array]	
data>>records>>ats03_months	3个月内库龄	是	[string]	2
data>>records>>ats1_years	12个月以上库龄	是	[string]	2
data>>records>>ats36_months	3-6个月库龄	是	[string]	2
data>>records>>ats69_months	6-9个月库龄	是	[string]	3
data>>records>>ats912_months	9-12个月库龄	是	[string]	2
data>>records>>available_quantity	WFS可售	是	[string]	4
data>>records>>gtin	GTIN码	是	[string]	00462301613707
data>>records>>inbound_quantity	标发在途	是	[string]	13
data>>records>>item_id	平台商品id	是	[string]	2301613700
data>>records>>last30_days_po_units	近30天计划入库	是	[string]	3
data>>records>>last30_days_units_received	近30天入库	是	[string]	3
data>>records>>platform_product_status	商品状态：停售、在售	是	[string]	在售
data>>records>>msku	msku	是	[string]	ast6-purple-xl-du00
data>>records>>pid	本地商品id	是	[string]	5764
data>>records>>product_name	品名	是	[string]	pm
data>>records>>sku	sku	是	[string]	sku
data>>records>>quantity	总库存	是	[string]	15
data>>records>>store_id	店铺id	是	[string]	110000000018008005
data>>records>>store_name	店铺名称	是	[string]	walmart测试店铺
data>>records>>unabled_warehousing_quantity	无法入库数量	是	[string]	1
data>>records>>warehouse_name	仓库名称	是	[string]	walmart测试店铺05号加拿大仓
data>>records>>warehouse_unique_id	仓库id【雪花id】	是	[string]	901337792229298176
data>>records>>pic_url	图片url	是	[string]	https://imagxxx/xxxx/5dc111.jpeg
data>>total	总数	是	[string]	21
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "1979122130ba45e984cbc9b2e571d86f.1690513512438",
    "response_time": "2023-07-28 11:05:12"
    "data": {
        "total": "21",
        "records": [
            {
                "msku": "ast6-purple-xl-du00",
                "gtin": "00462301111",
                "pid": "5764",
                "sku": "wjc14_sku",
                "quantity": "15",
                "store_id": "110000000018008005",
                "store_name": "walmart测试店铺",
                "warehouse_unique_id": "90133779222911",
                "warehouse_name": "walmart测试店铺05号加拿大仓",
                "item_id": "230161371",
                "product_name": "wjc14_pm",
                "platform_product_status": "在售",
                "available_quantity": "13",
                "unabled_warehousing_quantity": "1",
                "inbound_quantity": "2",
                "ats03_months": "3",
                "ats36_months": "4",
                "ats69_months": "5",
                "ats912_months": "6",
                "ats1_years": "7",
                "last30_days_units_received": "8",
                "last30_days_po_units": "9",
                "pic_url": "https://imagxxx/xxxx/5dc111.jpeg"
            }
        ]
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "参数检验不通过",
    "error_details": ["length: length 最大值200"],
    "request_id": "f4044649b1be46398c7eb4e6eb6ab6d5.1684735739946",
    "response_time": "2023-05-22 14:08:59",
    "data": null
}
复制
错误
复制成功
 上一章节
查询WFS货件列表
下一章节 
查询Temu库存