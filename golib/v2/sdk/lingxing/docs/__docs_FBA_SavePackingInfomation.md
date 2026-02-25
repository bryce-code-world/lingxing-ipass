保存装箱信息
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-packing/setLocalPackingInformation	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
boxes	箱子信息	是	[array]	
boxes>>dimensions		是	[object]	
boxes>>dimensions>>height	高	是	[string]	
boxes>>dimensions>>length	长	是	[string]	
boxes>>dimensions>>unitOfMeasurement	长度单位：IN、CM	是	[string]	
boxes>>dimensions>>width	宽	是	[string]	
boxes>>items	商品信息	是	[array]	
boxes>>items>>expiration	有效期	否	[string]	
boxes>>items>>labelOwner	标签类型：AMAZON、SELLER、NONE	是	[string]	
boxes>>items>>msku	msku	是	[string]	
boxes>>items>>prepOwner	预处理提供方：AMAZON、SELLER、NONE	是	[string]	
boxes>>items>>quantity	申报量	是	[int]	
boxes>>weight	重量	是	[object]	
boxes>>weight>>unit	重量单位：LB、KG	是	[string]	
boxes>>weight>>value	重量单位值	是	[string]	
inboundPlanId	STA任务编号，对应创建STA任务接口对应字段【inboundPlanId】	是	[string]	
packingGroupId	包装组id：先装箱后分仓方式时必填；先分仓后装箱方式时无需填写	否	[string]	
shipmentId	货件id：先分仓后装箱方式时必填；分装箱后分仓方式时无需填写	否	[string]	
sid	店铺id，对应查询亚马逊店铺列表接口对应字段【sid】	是	[string]	
请求 cURL 示例
curl --location --request POST 'https://openapi.lingxing.com/openapi/inbound-packing/setLocalPackingInformation?app_key=value&access_token=value&timestamp=value&sign=value' \
--header 'Content-Type: application/json' \
--data-raw '{
  "boxes": [
    {
      "dimensions": {
        "height": "30",
        "length": "20",
        "unitOfMeasurement": "CM",
        "width": "10"
      },
      "items": [
        {
          "expiration": "2023-12-31",
          "labelOwner": "SELLER",
          "msku": "msku",
          "prepOwner": "SELLER",
          "quantity": 10
        }
      ],
      "weight": {
        "unit": "KG",
        "value": "5.5"
      }
    }
  ],
  "inboundPlanId": "亚马逊任务编号",
  "packingGroupId": "包装组id：先装箱后分仓方式时必填；先分仓后装箱方式时无需填写",
  "sid": "1"
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
errorDetails	错误信息	是	[array]	
requestId	请求链路id	是	[string]	
responseTime	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[string]	
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "errorDetails": [],
    "requestId": "3b3d867e7d014971a580549f107c8c5a.1732773886069",
    "responseTime": "2024-11-28T14:04:46.069",
    "data": null
}
复制
错误
复制成功
 上一章节
查询包装组
FBA
下一章节 
提交装箱信息
FBA