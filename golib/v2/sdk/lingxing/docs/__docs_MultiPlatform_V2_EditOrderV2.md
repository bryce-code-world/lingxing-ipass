编辑订单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/editOrder	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
order_list	订单列表	是	[array]	
order_list>>global_order_no	系统单号	是	[int]	103257522489451000
order_list>>logistics	物流信息	是	[object]	
order_list>>logistics>>logistics_type_id	物流方式id，对应 查询已启用的自发货物流方式 接口字段【type_id】	是	[int]	825
order_list>>logistics>>sys_wid	仓库id，对应 查询仓库列表 接口字段【wid】	是	[int]	50
请求示例
{
    "order_list": [{
        "global_order_no": 103257522489451000,
        "logistics": {
            "logistics_type_id": 825,
            "sys_wid": 50
        }
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码
0：订单更新成功
10000：订单更新失败
10001：订单更新部分成功	是	[int]	0
message	消息提示	是	[string]	订单更新成功
data	返回数据	是	[object]	
data>>error_details	错误信息	是	[array]	
data>>error_details>>global_order_no	系统单号	是	[string]	103257522489451008
data>>error_details>>error_message	更新失败错误信息	是	[string]	请求未知错误
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
返回成功示例
{
    "code": 0,
    "message": "订单更新成功",
    "data": {
        "error_details": []
    },
    "request_id": "ad9c1a8625c74f3ea0525e216374e31a.1676514819613",
    "response_time": "2023-02-16 10:33:40"
}
复制
错误
复制成功
返回失败示例
{
    "code": 10000,
    "message": "订单更新失败",
    "data": {
        "error_details": [
            {
                "global_order_no": "103257522489451008",
                "error_message": "请求未知错误"
            }
        ]
    },
    "request_id": "ad9c1a8625c74f3ea0525e216374e31a.1676514819613",
    "response_time": "2023-02-16 10:33:40"
}
复制
错误
复制成功
 上一章节
订单称重
下一章节 
编辑/更新自发货订单