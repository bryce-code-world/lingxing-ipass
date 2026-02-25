SC订单-设置订单备注
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/scOrder/setRemark	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	12
amazonOrderId	订单id	是	[string]	249-6838153-6838402
remark	备注	是	[string]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/platformOrder/scOrder/setRemark?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "sid": 12,
    "amazonOrderId": "249-6838153-6838402",
    "remark": ""
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	720e6a6bac5b4c729d4ecea5690c4e80.1734939490648
response_time	响应时间	是	[string]	2024-12-23 15:38:10
data	null	是	[null]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "",
    "response_time": "",
    "data": null
}
复制
错误
复制成功
 上一章节
查询亚马逊订单详情
销售
下一章节 
查询亚马逊多渠道订单列表v2
销售