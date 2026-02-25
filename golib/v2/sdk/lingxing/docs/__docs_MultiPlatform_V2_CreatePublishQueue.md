Temu修改库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/temu/createPublishQueue	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
platformCode	平台代码	是	[string]	10024
storeId	店铺id	是	[string]	110443249963082240
productUniqueId	商品uid	是	[string]	901458524464402944
quantity	数量	是	[string]	59
warehouseId	仓库ID	是	[string]	WH-03852452270231627
warehouseName	仓库名称	否	[string]	美东仓库
请求示例
{
    "platformCode": "10021",
    "storeId": "110000000018021002",
    "productUniqueId": "901457821424734720",
    "quantity": "30",
    "warehouseId": "PS1618418400",
    "warehouseName": "美国1"
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
request_id	请求链路id	是	[string]	5277ebc6b1654a96ad4ffc51729b191e.1749697431259
response_time	响应时间	是	[string]	2025-06-12 11:03:51
data	响应数据	是	[object]	true
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "5277ebc6b1654a96ad4ffc51729b191e.1749697431259",
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
    "request_id": "5277ebc6b1654a96ad4ffc51729b191e.1749697431259",
    "response_time": "2025-06-12 11:03:51",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
查询AliExpress在线商品列表
下一章节 
Shein修改库存