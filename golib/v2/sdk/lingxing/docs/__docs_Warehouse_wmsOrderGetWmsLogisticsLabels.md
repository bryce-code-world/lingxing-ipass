查询销售出库单物流面单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/wms/order/getWmsLogisticsLabels	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wo_number_arr	销售出库单号,上限50【销售出库单号与系统单号二选一必填】	否	[array]	["WO103132593465409536"]
order_number_arr	系统单号,上限50【销售出库单号与系统单号二选一必填】	否	[array]	["103130837064323072"]
请求示例
{
    "wo_number_arr": [
        "WO103470130495287296",
        "WO103470130483106304"
    ]
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	DCF829F2-1054-08BC-7501-FBCF8BC614CA
response_time	响应时间	是	[string]	2023-04-24 17:29:54
data	响应数据	是	[array]	
data>>wo_number	销售出库单号	是	[string]	WO103321499775717888
data>>order_number	系统单号	是	[string]	103319017000570880
data>>logistics_provider_id	物流服务商id	是	[number]	3
data>>logistics_type_id	物流方式id	是	[number]	919
data>>file_id	文件id	是	[number]	765
data>>file_type	文件类型	是	[string]	pdf
data>>file_size	文件尺寸	是	[string]	100*100
data>>file_b64	面单base64	是	[string]	JVBERi0xLjcKJ
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7755005B-2D32-D55A-780B-44A32A00790E",
    "response_time": "2024-04-10 14:36:16",
    "data": [
        {
            "wo_number": "WO103429905417286144",
            "order_number": "103425935326859776",
            "logistics_provider_id": 1666,
            "logistics_type_id": 41925,
            "file_id": 0,
            "file_type": "",
            "file_size": "",
            "file_b64": ""
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
物流下单 - 编辑运单号/跟踪号
下一章节 
销售出库单截单