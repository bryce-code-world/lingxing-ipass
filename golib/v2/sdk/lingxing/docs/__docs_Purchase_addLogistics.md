添加采购单物流信息
接口信息

支持对“待到货”或“已完成”状态的采购单，添加采购单物流信息

API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchase/addLogistics	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	采购单号（待到货或已完成状态）	是	[string]	PO220418003
items	物流信息	是	[array]	
items>>logistics_company	物流商	是	[string]	中通快递
items>>logistics_order_no	物流单号（支持字母、数字、下划线、短划线）	是	[string]	ZT12456789
请求示例
{
    "order_sn": "POC240802001",
    "items": [
        {
            "logistics_company": "中通快递",
            "logistics_order_no": "ZT12456789"
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8DFEE1CA-EC9B-F401-5D41-9F95251D5D50
response_time	响应时间	是	[string]	2020-09-21 15:48:58
data	数据	是	[array]	
total		是	[int]	 
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
 上一章节
查询委外订单列表
下一章节 
作废采购计划