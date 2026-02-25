审核发货

订单管理审核发货，对应系统中订单管理“待审核发货”类型的订单进行审核发货

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/multiplatform/order/review	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
global_order_no	系统单号列表	是	[array]	"global_order_no": ["103584428644401152"]
返回结果

Json Object

参数名称	参数说明	类型	示例
code	状态码，0 成功	[int]	0
message	返回消息	[String]	
request_id	请求id	[String]	
response_time	请求时间	[String]	
data	响应数据	[object]	
data>>success_num	成功数量	[int]	1
data>>fail_num	失败数量	[int]	1
data>>failure_info	失败详情信息	[array]	
data>>failure_info>>message	失败原因	[string]	未设置物流
data>>failure_info>>global_order_no	系统单号	[string]	103579112899859456
data>>success_info	成功详情信息	[array]	
data>>success_info>>global_order_no	系统单号	[string]	103579112899859456
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "faacdb1d437445f1bad7079d03cce8f5.1750727281588",
    "response_time": "2025-06-24 09:08:10",
    "data": {
        "success_num": 1,
        "fail_num": 1,
        "failure_info": [{
            "message": "未设置仓库;未设置物流",
            "global_order_no": "103584428644401152"
        }],
        "success_info": [{
            "global_order_no": "103584102204547072"
        }]
    }
}
复制
错误
复制成功
 上一章节
预发货
下一章节 
创建订单