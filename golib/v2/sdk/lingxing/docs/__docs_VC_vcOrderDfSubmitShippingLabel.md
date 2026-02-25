VC订单-请求标签【DF】
请求标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/vcOrderDf/submitShippingLabel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
ids	订单ID，查询VC订单列表接口对应字段【id】	是	[array]	["107"]
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "46ff0566a6c449ab8b62e20e229398b8.93.16983154160003075",
    "response_time": "2023-10-26 18:16:56",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
VC订单-确认发货【DF】
下一章节 
VC订单-打印标签【DF】