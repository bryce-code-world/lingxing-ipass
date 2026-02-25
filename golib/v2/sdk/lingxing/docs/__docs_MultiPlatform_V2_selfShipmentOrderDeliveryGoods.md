订单发货

对应订单管理-》发货 操作，对系统订单标记出库，扣减库存

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/selfShipmentOrder/deliveryGoods	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_number_list	系统单号列表，多个使用英文逗号分隔，上限100	是	[string]	"103248664914830848,103157081579986952"
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>success_list	发货成功列表	是	[array]	
data>>success_list>>order_number	系统单号	是	[string]	103248664914830848
data>>success_list>>status_name	订单状态说明	是	[string]	已发货
data>>fail_list	发货失败列表	是	[array]	
data>>fail_list>>order_number	系统单号	是	[string]	103248664914830848
data>>fail_list>>err_msg	错误信息	是	[string]	订单不在发货中
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "57dd5343b56f41409a5cbb9d8260adf9.1699496736534",
    "response_time": "2023-11-09 10:25:37",
    "data": {
        "success_list": [],
        "fail_list": [
            {
                "order_number": "103157081579986952",
                "err_msg": "订单不在发货中"
            },
            {
                "order_number": "103157081579986953",
                "err_msg": "订单不在发货中"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
标记订单不发货
下一章节 
订单称重