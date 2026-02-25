创建FBA发货计划
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/createShipmentPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
remark	批次信息备注	否	[string]	这个是批次备注111
product_list	商品信息	是	[array]	
product_list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	16
product_list>>packing_type	包装类型： 1 混装，2 原厂	是	[int]	2
product_list>>shipment_time	发货时间，格式：Y-m-d	否	[string]	2022-08-31
product_list>>msku	MSKU	是	[string]	KJGEK4395
product_list>>fnsku	FNSKU	是	[string]	XXX40UMLRR
product_list>>shipment_plan_quantity	计划发货量	是	[int]	2
product_list>>quantity_in_case	单箱数量【PCS】	否	[int]	1
product_list>>box_num	箱数	否	[int]	2
product_list>>logistics_provider_id	物流商id，查询头程物流渠道列表 接口对应字段【data>>provider>>id】	否	[int]	
product_list>>logistics_channel_id	物流渠道id，查询头程物流渠道列表 接口对应字段【data>>id】	否	[int]	1
product_list>>wid	系统仓库id	否	[int]	1
product_list>>remark	商品信息备注	否	[string]	aa1
请求示例
{
    "remark":"备注",
    "product_list": [
        {
            "sid":104,
            "packing_type": 1,
            "shipment_time":"2022-12-14",
            "msku": "XGJ0008",
            "fnsku": "B08D3N9BK3",
            "shipment_plan_quantity": 2,
            "quantity_in_case":1,
            "box_num":1,
            "logistics_provider_id":1,
            "logistics_channel_id":10000405,
            "wid":1,
            "remark":"备注xx"

        }
    ]
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
request_id	请求链路id	是	[string]	D8344695-90C7-B7CF-86E2-25697AFFC527
response_time	响应时间	是	[string]	2022-12-08 16:27:10
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>seq	批次号	是	[string]	RP231220003
data>>order_sn	计划编号	是	[array]	["R231220003"]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D8344695-90C7-B7CF-86E2-25697AFFC527",
    "response_time": "2022-12-08 16:27:10",
    "data": {
        "seq": "RP231220003",
        "order_sn": [
            "R231220003"
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询FBA发货计划
FBA
下一章节 
编辑FBA发货计划
FBA