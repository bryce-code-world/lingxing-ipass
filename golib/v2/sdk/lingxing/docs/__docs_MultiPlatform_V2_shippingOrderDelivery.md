平台仓发货单发货
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/shippingList/delivery	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shippingIdList	发货单ID列表，对应查询平台仓发货单列表v2接口出参id	否	[array]	[1,2,36]
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/multiplatform/shippingList/delivery?app_key=&access_token=&timestamp=&sign=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "shippingIdList": [
        "675",
        "674",
        "672"
    ]
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	
data>>success	成功数量	是	[string]	0
data>>fail	失败数量	是	[string]	2
data>>total	总单数	是	[string]	2
data>>failList	失败明细	是	[array]	
data>>failList>>shippingListCode	失败单号	是	[string]	IB220927004
data>>failList>>failReason	失败原因	是	[string]	
返回结果示例
{
  "code": 0,
  "message": "success",
  "error_details": [],
  "request_id": "7b05b06f1575473c82f2261f8dc04e41.1747191247574",
  "response_time": "2025-05-14 10:54:07",
  "data": {
    "success": 1,
    "fail": 2,
    "total": 3,
    "errorList": [
      {
        "shippingListCode": "SP202407150030034",
        "failReason": "未选择物流方式"
      },
      {
        "shippingListCode": "SP202502270020076",
        "failReason": "未选择物流方式"
      }
    ]
  },
  "total": 3
}
复制
错误
复制成功
 上一章节
平台仓发货单拣货
下一章节 
查询Walmart Review列表