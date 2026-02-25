Shein修改库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/shein/createPublishQueue	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
platformCode	平台代码	否	[string]	10021
storeId	店铺id	否	[string]	110329283045806082
productUniqueId	商品uid	否	[string]	901461697188614656
quantity	数量	否	[string]	30
warehouseId	仓库ID	否	[string]	PS8682401841
warehouseName	仓库名称	否	[string]	法国
请求示例
{
    "platformCode": "10021",
    "storeId": "110329283045806082",
    "productUniqueId": "901461697188614656",
    "quantity": "30",
    "warehouseId": "PS8682401841",
    "warehouseName": "法国"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	[]
request_id	请求链路id	是	[string]	52c6b94a4af94dadb82873ee39db1558.1749697759505
response_time	响应时间	是	[string]	2025-06-12 11:09:19
data	响应数据	是	[object]	true
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "52c6b94a4af94dadb82873ee39db1558.1749697759505",
    "response_time": "2025-06-12 11:09:19",
    "data": true,
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "程序内部错误",
    "error_details": [
        "商品存在重复刊登"
    ],
    "request_id": "52c6b94a4af94dadb82873ee39db1558.1749697759505",
    "response_time": "2025-06-12 11:09:19",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
Temu修改库存
下一章节 
Walmart修改库存