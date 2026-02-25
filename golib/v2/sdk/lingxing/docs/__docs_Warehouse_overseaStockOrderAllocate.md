备货单分配库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/overSeaWarehouse/stockOrder/allocate	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
orderNo	备货单号	是	[string]	OWS241231002
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/overSeaWarehouse/stockOrder/allocate?app_key=&access_token=&timestamp=&sign=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderNo":"OWS250303002"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	错误说明	是	[string]	success
request_id	请求id	是	[string]	69a2ddc2-f3bb-11ed-aed5-0242ac130003
response_time	请求时间	是	[string]	2023-05-16 15:29:41
error_details	错误信息	是	[array]	
data	响应数据	是	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "47525526bfd3484cbb1aba1870a1b2c3.1744366551396",
    "response_time": "2025-04-11 18:15:52",
    "data": null
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "程序内部错误",
    "error_details": [
        "锁定库存出现异常:sku:'testwpw4447' fnsku:'' 店铺：'' 仓库可用库存不足 [请求码:B113C8]"
    ],
    "request_id": "7a751c1662b645dd8576a71dabf362f0.1747188035250",
    "response_time": "2025-05-14 10:00:38",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
查询移除入库单列表
下一章节 
查询销售出库单列表