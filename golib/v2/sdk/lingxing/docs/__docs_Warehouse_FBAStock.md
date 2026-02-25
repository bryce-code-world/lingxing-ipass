查询FBA库存列表

支持查询FBA库存，对应系统【仓库】>【FBA库存明细】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/fba/fbaStock/fbaList	HTTPS	POST	3
请求参数
参数名	说明	必填	类型	示例
sid	店铺id，多个使用英文逗号分隔 ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[string]	1,136,139
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15	否	[int]	20
请求cURL示例
curl --location 'https://openapi.lingxing.com/erp/sc/routing/fba/fbaStock/fbaList?access_token=value&timestamp=value&sign=value&app_key=value'
--header 'Content-Type: application/json' 
--data '{
    "sid": "1,136,139",
    "offset": 0,
    "length": 20
}'
复制
错误
复制成功
请求示例
{
    "sid": "1,136,139",
    "offset": 0,
    "length": 20
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
request_id	请求链路id	是	[string]	80B5BD25-FD74-62D2-84C1-75A6D838A2CB
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[array]	
data>>total	总数	是	[int]	697
data>>list	列表数据	是	[array]	
data>>list>>wname	仓库名称	是	[string]	TJ-TK欧洲仓
data>>list>>name	仓库名称【同wname】	是	[string]	TJ-TK欧洲仓
data>>list>>sid	店铺id	是	[int]	13
data>>list>>asin	ASIN	是	[string]	B09XB4SDLK
data>>list>>product_name	品名	是	[string]	垒球-新品
data>>list>>product_image	图片	是	[string]	https://xxx/xxx.jPg
data>>list>>msku	MSKU	是	[string]	EU-FDH-843-22
data>>list>>fnsku	FNSKU	是	[string]	X001JRXEP9
data>>list>>sku	SKU	是	[string]	8911545789
data>>list>>category_name	分类名称	是	[string]	球
data>>list>>category_id	分类Id	是	[int]	7
data>>list>>brand_name	品牌名称	是	[string]	ABC
data>>list>>brand_id	品牌id	是	[int]	2
data>>list>>stock_cost_total	货值	是	[number]	
data>>list>>cost	库存成本	是	[number]	
data>>list>>share_type	共享类型：
0 非共享
1 北美共享
2 欧洲共享	是	[int]	2
data>>list>>afn_fulfillable_quantity	可售	是	[string]	0
data>>list>>reserved_fc_transfers	待调仓	是	[string]	0
data>>list>>reserved_fc_processing	调仓中	是	[string]	0
data>>list>>reserved_customerorders	待发货	是	[string]	0
data>>list>>total_fulfillable_quantity	总可用库存：可售+待调仓+调仓中
【非ERP页面对应总库存】	是	[int]	0
data>>list>>afn_unsellable_quantity	不可售	是	[int]	0
data>>list>>afn_inbound_working_quantity	计划入库	是	[int]	0
data>>list>>afn_inbound_shipped_quantity	在途	是	[int]	0
data>>list>>afn_inbound_receiving_quantity	入库中	是	[int]	0
data>>list>>afn_erp_real_shipped_quantity	实际在途	是	[int]	0
data>>list>>afn_researching_quantity	调查中数量	是	[int]	0
data>>list>>inv_age_0_to_30_days	0-1个月库龄	是	[string]	0
data>>list>>inv_age_31_to_60_days	1-2个月库龄	是	[string]	0
data>>list>>inv_age_61_to_90_days	2-3个月库龄	是	[string]	0
data>>list>>inv_age_0_to_90_days	0-3个月库龄	是	[string]	0
data>>list>>inv_age_91_to_180_days	3-6个月库龄	是	[string]	0
data>>list>>inv_age_181_to_270_days	6-9个月库龄	是	[string]	0
data>>list>>inv_age_271_to_330_days	9-11个月库龄	是	[string]	0
data>>list>>inv_age_331_to_365_days	11-12个月库龄	是	[string]	0
data>>list>>inv_age_271_to_365_days	9-12个月库存	是	[string]	0
data>>list>>inv_age_365_plus_days	12个月以上库龄	是	[string]	0
data>>list>>fulfillment_channel_name	配送方式	是	[string]	FBA
data>>list>>afn_fulfillable_quantity_multi	欧洲共享的多国店铺列表	是	[array]	0
data>>list>>afn_fulfillable_quantity_multi>>name	店铺名称	是	[string]	0
data>>list>>afn_fulfillable_quantity_multi>>quantity_for_local_fulfillment	FBA可售数量	是	[string]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "F95AD8DF-C7AB-050E-CE6D-EA69C3A3D6E9",
    "response_time": "2024-07-31 15:02:59",
    "data": {
        "total": 19,
        "list": [
            {
                "wname": "renxiao-US美国仓",
                "name": "renxiao-US美国仓",
                "asin": "B0CM3B75YP",
                "sid": 136,
                "msku": "CN0001",
                "fnsku": "X0040T7HJP",
                "sku": "01",
                "brand_id": 1021,
                "category_id": 1024,
                "brand_name": "蛋",
                "category_name": "王",
                "product_name": "烤红薯",
                "product_image": "https://image.distributetop.com/lingxing-erp/90128554873982976/20240523/d937bd94d48d4177ba7c1601e8ec8b54.jpg",
                "share_type": 0,
                "total_fulfillable_quantity": 0,
                "afn_fulfillable_quantity": 0,
                "afn_erp_real_shipped_quantity": 5,
                "reserved_fc_transfers": 0,
                "reserved_fc_processing": 0,
                "reserved_customerorders": 0,
                "afn_unsellable_quantity": 0,
                "afn_inbound_working_quantity": 334,
                "afn_inbound_shipped_quantity": 183,
                "afn_inbound_receiving_quantity": 0,
                "afn_researching_quantity": 0,
                "inv_age_0_to_30_days": 0,
                "inv_age_31_to_60_days": 0,
                "inv_age_61_to_90_days": 0,
                "inv_age_0_to_90_days": 0,
                "inv_age_91_to_180_days": 0,
                "inv_age_181_to_270_days": 0,
                "inv_age_271_to_330_days": 0,
                "inv_age_331_to_365_days": 0,
                "inv_age_271_to_365_days": 0,
                "inv_age_365_plus_days": 0,
                "fulfillment_channel_name": "FBA",
                "afn_fulfillable_quantity_multi": [],
                "cost": "184830.00",
                "stock_cost_total": 184830
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
附加说明
唯一键：sid+msku 组合
 上一章节
查询产品仓位列表
下一章节 
查询FBA库存列表v2