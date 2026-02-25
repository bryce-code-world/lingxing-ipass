删除Listing标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/globalTag/listing/removeTag	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
tag_ids	标签id，上限200	是	[array]	["907365081622413334"]
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	77ac259a67d5462594c83b80669b6eae.1692331008758
response_time	响应时间	是	[string]	2023-08-18 11:56:49
total	总数	是	[int]	0
data	响应数据	是	[object]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7db166f311594a9986e882af1bd69b7c.1696852339529",
    "response_time": "2023-10-09 19:52:20",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
添加Listing标签
销售
下一章节 
FBA费差异-异常订单-订单
销售