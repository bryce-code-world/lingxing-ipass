上传货件跟踪号
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-shipment/updateShipmentTrack	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
billOfLadingNumber	提货单号,LTL建议填写,非必填	否	[string]	
freightBillNumber	LTL跟踪编号(LTL必填)	否	[string]	
inboundPlanId	STA任务编号	否	[string]	
shipmentConfirmationId	货件单号	否	[string]	
shipmentId	货件id	否	[string]	
sid	领星店铺ID	否	[long]	
trackBOList	跟踪编号列表,SPD必填	否	[array]	
trackBOList>>boxId	箱子id	否	[string]	
trackBOList>>localBoxId	本地箱子id	否	[string]	
trackBOList>>trackingId	跟踪id	否	[string]
返回结果
参数名	说明	必填	类型	示例
code		否	[int]	
data		否	[object]	
data>>errorEnums	错误编码（让openapi的用户进行后续操作）,OpenApiTypeEnum 枚举值	否	[array]	
data>>errorMsg	错误信息	否	[string]	
data>>inboundPlanId	亚马逊任务编号	否	[string]	
data>>taskId	任务id	否	[string]	
data>>taskStatus	任务状态	否	[string]	
errorDetails		否	[array]	
message		否	[string]	
requestId		否	[string]	
responseTime		否	[string]
 上一章节
修改货件实际状态
FBA
下一章节 
取消STA任务
FBA