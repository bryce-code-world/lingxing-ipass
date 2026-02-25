作废采购计划

允许作废处于 “待审批”、“待采购”状态下的采购计划

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/purchase/planCancel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
plan_sn	计划编号	是	[array]	["PP2203103920","PP2203103918"]
reason	作废原因	是	[string]	12
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/storage/awdWarehouseDetail?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "plan_sn": [
        "PP2203103923"
    ],
    "reason": "12"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	3b41b5ac4348416b8b1e8210d5d8f39f.1728612378392
response_time	响应时间	是	[string]	2024-10-11 10:06:18
data		是	[null]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [
        "操作成功"
    ],
    "request_id": "41989fd5eadf468dbd5a74ee679c24a6.1728628880708",
    "response_time": "2024-10-11 14:41:20",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
添加采购单物流信息
下一章节 
编辑采购单备注