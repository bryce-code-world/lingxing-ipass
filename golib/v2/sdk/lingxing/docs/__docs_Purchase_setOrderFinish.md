采购单整单结束到货
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/purchase/setOrderFinish	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderSn	仅支持系统单号，不支持自定义采购单号	是	[array]	["PO210510004"]
请求curl示例
curl --location --globoff 'https://openapi.lingxing.com/basicOpen/purchase/setOrderFinish?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data '{
    "order_sn": [
        "PO250313001",
        "PO250313003",
        "PO250313002"
    ]
}'
复制
错误
复制成功
返回结果

JsonObject

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
request_id	请求链路id	是	[string]	550d352c-7a05-11ed-b0c7-0242ac1c0004
response_time	响应时间	是	[string]	2022-12-15 20:16:38
error_details	失败信息	是	[object]	orderSn=PO250103001, detail=存在未收货的收货单，暂不支持操作，请先收货或删除收货单)
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "550d352c-7a05-11ed-b0c7-0242ac1c0004",
    "response_time": "2022-12-15 20:16:38",
    "data": null
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "[(order_sn='PO250313001', detail='结束到货失败，采购单状态已变更'), (order_sn='PO250313003', detail='结束到货失败，采购单状态已变更'), (order_sn='PO250313002', detail='结束到货失败，采购单状态已变更')]"
    ],
    "request_id": "943f1193cba94c31b275c4155e532411.1741852509434",
    "response_time": "2025-03-13 15:55:11",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
编辑采购单备注
下一章节 
查询仓库列表