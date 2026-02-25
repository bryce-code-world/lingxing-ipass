预发货

订单管理预发货，对应系统中订单管理“待审核发货”类型的订单进行预发货

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/multiplatform/order/preShipment	HTTPS	POST	1
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
data>>failure_info>>message	失败原因	[string]	订单预发货中
data>>failure_info>>global_order_no	系统单号	[string]	103579112899859456
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "f50f8534a7c84d5da3631be8d35ddb7b.1750728380658",
    "response_time": "2025-06-24 09:26:21",
    "data": {
        "success_num": 1,
        "fail_num": 1,
        "failure_info": [{
            "message": "订单预发货中，如需操作订单请稍候",
            "global_order_no": "103579112899859456"
        }]
    }
}
复制
错误
复制成功
 上一章节
批量添加/编辑多平台配对关系
下一章节 
审核发货