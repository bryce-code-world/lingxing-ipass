编辑发货单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/updateInboundShipmentListMws	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipment_sn	发货单号	是	[string]	SP241016009
remark	备注	否	[string]	test001
items	发货商品	否	[array]	
items>>id	商品明细id，查询发货单详情接口对应字段【data>>items>>id】	否	[int]	test001
items>>num	发货量，发货量不允许大于计划发货量	否	[string]	10
请求示例
{
    "shipment_sn": "SP241016009",
    "remark": "test001"
    "items":[
        {
            "id":1,
            "num":"10"
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	107DEE19-E3DD-E6C6-F63D-EB8FF2D92327
response_time	响应时间	是	[string]	2024-10-17 14:51:53
data	响应数据	是	[array]	
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "107DEE19-E3DD-E6C6-F63D-EB8FF2D92327",
    "response_time": "2024-10-17 14:51:53",
    "data": [],
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 1000,
    "message": "业务处理失败",
    "error_details": [
        "发货单SP24101600不存在或已删除"
    ],
    "request_id": "4ECEB7CF-5B32-6CEA-E6F6-32B34EF38E47",
    "response_time": "2024-10-17 14:57:43",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
FBA发货单发货
FBA
下一章节 
删除发货单
FBA