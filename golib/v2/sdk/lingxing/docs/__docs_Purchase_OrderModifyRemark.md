编辑采购单备注

支持快速编辑采购单的单据备注

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/purchase/orderModifyRemark	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sns	采购单号	是	[array]	["PO240925001"]
value	备注内容	是	[string]	你好
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/purchase/orderModifyRemark?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "order_sns": [
        "PO240925001"
    ],
    "value": "你好"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	dcdde6267f8b42b4be559fb94e835daf.1728716051245
response_time	响应时间	是	[string]	2024-10-12 14:54:11
data	该接口data返回为null	是	[object]	
total	总数	是	[int]	0
返回成功示例
{
"code": 0,
"message": "success",
"error_details": [],
"request_id": "0511bfa5f9b8492d81ad01eb06aa4ce4.1728716772752",
"response_time": "2024-10-12 15:06:12",
"data": null,
"total": 0
}
复制
错误
复制成功
 上一章节
作废采购计划
下一章节 
采购单整单结束到货