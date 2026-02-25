合并订单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/mergeOrder	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
platform_code	平台code【不支持10007 Lazada、10011 TikTok、10012 MERCADO】	是	[string]	10002
order_list	系统单号	是	[array]	["103318651460264448","103318391748825088"]
请求示例
{
    "platform_code": "10002",
    "order_list": [
        "103318651460264448",
        "103318391748825088"
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	成功
response_time	响应时间	是	[string]	2023-06-01 16:39:20
request_id	请求链路id	是	[string]	1573eb10-f465-11ed-87da-0242ac150002
data	响应数据	是	[object]	
data>>fail_num	失败数量	是	[int]	1
data>>failure_info	失败明细	是	[array]	
data>>failure_info>>global_order_no	系统单号	是	[array]	["103313004478456320"]
data>>failure_info>>message	失败原因	是	[string]	操作的勾选项无可合并的订单!
data>>success_num	成功数量	是	[int]	0
data>>success_info	成功明细	是	[array]	 
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "47ce90d6-e0c4-4ae6-ade2-b3986cb5c80a.1686041522668",
    "response_time": "2023-06-06 16:52:04",
    "data": {
        "success_num": 2,
        "fail_num": 0,
        "failure_info": [],
        "success_info": [
            {
                "global_order_no": [
                    "103320875956794368",
                    "103320876940682240"
                ],
                "new_order_no": "103320875956794368"
            }
        ]
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "15a8b242-b0a6-4643-8e5b-b976a8c8cd65.1686043720608",
    "response_time": "2023-06-06 17:28:40",
    "data": {
        "success_num": 0,
        "fail_num": 1,
        "failure_info": [
            {
                "global_order_no": [
                    "103320875956794369"
                ],
                "message": "操作的勾选项无可合并的订单!"
            }
        ],
        "success_info": []
    }
}
复制
错误
复制成功
 上一章节
创建订单
下一章节 
拆分订单