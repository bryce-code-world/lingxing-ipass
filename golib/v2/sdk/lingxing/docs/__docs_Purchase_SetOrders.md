采购单下单

支持将 “待下单” 状态采购单变更为 “待到货” 状态

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase/setOrders	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	采购单，对应查询采购单列表接口字段data>>order_sn	是	[array]	["PO210705007"]
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2022-05-20 16:57:03
data	响应数据	是	[array]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8DFEE1CA-EC9B-F401-5D41-9F95251D5D50",
    "response_time": "2021-11-12 12:41:04",
    "data": [],
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "success",
    "error_details": [],
    "request_id": "1DE78F69-1370-4C80-F97A-D9DDF0710F9E",
    "response_time": "2021-11-12 12:39:50",
    "data": [
        {
            "order_sn": "PO211112002",
            "detail": "错误：采购单状态已变更"
        }
    ],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询采购单列表
下一章节 
作废采购单