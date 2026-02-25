标记订单不发货
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/cancelOrder	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
order_list	系统单号列表	是	[array]	["103430576541848064","103432007211410432"]
请求示例
{
    "order_list": [
        "103430576541848064",
        "103432007211410432"
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[object]	操作成功
request_id	请求链路id	是	[string]	f8697c1c-d740-488a-84bb-fc3a512fc295.1684293345104
response_time	返回时间	是	[string]	2023-05-17 11:15:45
data	响应数据	是	[object]	
data>>success_num	成功数量	是	[int]	1
data>>fail_num	失败数量	是	[int]	0
data>>process_code	处理状态code：
10000 全部失败
10001 部分成功
0002 全部成功	是	[int]	10002
data>>failure_info	失败明细，全部成功返回空数组	是	[array]	
data>>failure_info>>global_order_no	系统单号	是	[string]	1033205360858531841
data>>failure_info>>message	失败原因	是	[string]	订单不存在
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "65b8587be2f545938d083d63280e1f81.98.16856086344797323",
    "response_time": "2023-06-01 16:37:14",
    "data": {
        "success_num": 1,
        "fail_num": 0,
        "process_code": 10002,
        "failure_info": []
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "da25c3b6-86ba-491c-b119-3cd4fcd088a2.1686043568992",
    "response_time": "2023-06-06 17:26:09",
    "data": {
        "success_num": 0,
        "fail_num": 1,
        "process_code": 10001,
        "failure_info": [
            {
                "global_order_no": "1033205360858531841",
                "message": "订单不存在！"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
拆分订单
下一章节 
订单发货