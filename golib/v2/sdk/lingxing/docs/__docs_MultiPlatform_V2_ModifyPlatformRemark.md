修改平台仓发货单备注
接口信息
API Path	请求协议	请求方式	令牌桶容量
/cepf/warehouse/api/openApi/editPlatfromShippingRemark	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
remarkContent	备注内容	是	[string]	"this is test remark and that is the goal2"
shippingListCode	平台单号	是	[string]	"SP202504020010076"
请求示例
{
    "remarkContent": "this is test remark and that is the goal2",  
    "shippingListCode": "SP202504020010076" 
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
data	操作是否成功	是	[boolean]	true
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	dc0b4baf-46b3-4cd4-b4b5-6a522c3072d7.1744085059269
response_time	响应时间	是	[string]	2025-04-08 12:04:19
返回成功示例
{
    "code": 0,
    "message": "success",
    "data": true,
    "error_details": [],
    "request_id": "dc0b4baf-46b3-4cd4-b4b5-6a522c3072d7.1744085059269",
    "response_time": "2025-04-08 12:04:19"
}
复制
错误
复制成功
 上一章节
查询平台仓发货单列表v2
下一章节 
查询WFS货件列表