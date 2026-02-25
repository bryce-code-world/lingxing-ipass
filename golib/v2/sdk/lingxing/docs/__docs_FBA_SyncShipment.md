同步亚马逊货件到ERP
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/fba/shipment/syncShipment	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	16
shipment_ids	货件编号	是	[array]	["FBA16ZK1LLX5"]
sync_anyway	报错是否继续：0 否【默认】，1 是
当系统检测到货件归属国家与店铺不符时，会提示报错，此时传1则按照店铺进行同步	否	[int]	0
请求示例
{
    "sid": 16,
    "shipment_ids": ["FBA16ZK1LLX5"],
    "sync_anyway": 0
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[string]	
request_id	请求链路id	是	[string]	EE6DEBE6-4CA9-A727-5211-1E5A6B81109F
response_time	响应时间	是	[string]	2023-05-31 19:54:54
data	响应数据	是	[object]	
data>>succ_num	同步成功数量	是	[int]	1
data>>fail_num	同步事变数量	是	[int]	2
data>>error	失败具体原因	是	[array]	
data>>error>>shipment_id	失败货件编号	是	[string]	xxx
data>>error>>detail	原因	是	[string]	货件单号不存在
data>>error>>is_error_seller	是否是店铺不匹配：0 否，1 是	是	[int]	0
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": null,
    "request_id": "F9FEA9AF-8CFC-7368-B98E-5B6910ABCAD7",
    "response_time": "2023-06-07 18:10:12",
    "data": {
        "succ_num": 1,
        "fail_num": 0,
        "error": []
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
同步STA任务到ERP
FBA
下一章节 
查询FBA货件箱子、卡板标签
FBA