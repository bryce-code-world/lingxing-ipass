修改货件实际状态
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/updateShipmentActualStatus	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
is_closed	货件状态：0 进行中，1 已完成	是	[int]	1
list	货件信息	是	[array]	
list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	12
list>>shipment_id	货件单号	是	[string]	FBA17124YWC9
请求cURL示例
curl --location 'https://openapi.lingxing.com/erp/sc/routing/storage/shipment/updateShipmentActualStatus?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "is_closed": 1,
    "list": [
        {
            "sid": 12,
            "shipment_id": "FBA17124YWC9"
        }
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
request_id	请求链路id	是	[string]	4A761C3F-D6BC-C81C-807F-4EC70B4DB516
response_time	响应时间	是	[string]	2023-02-23 10:27:32
data	响应数据	是	[array]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "4A761C3F-D6BC-C81C-807F-4EC70B4DB516",
    "response_time": "2023-02-23 10:27:32",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
修改货件装箱信息
FBA
下一章节 
上传货件跟踪号
FBA