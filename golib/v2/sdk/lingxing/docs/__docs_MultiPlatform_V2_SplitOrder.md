拆分订单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/splitOrder	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
split_mod	拆分模式：1 按商品拆分，2 按捆绑商品拆分	是	[int]	1
global_order_no	系统单号	是	[string]	103313000188121600
order_item	订单数据	是	[array]	
order_item>>item_id	商品行主键id【拆分模式split_mod = 1 时必填】	是	[string]	6454374
order_item>>quantity	数量【拆分模式split_mod = 1 时必填】	是	[string]	1
order_item>>pid	本地商品id【拆分模式split_mod = 2 时必填，没有值传空字符串】	是	[string]	18522
order_item>>order_item_no	订单明细单号【拆分模式split_mod = 2 时必填，没有值传空字符串】	是	[string]	1
order_item>>msku	平台msku【拆分模式split_mod = 2 时必填，没有值传空字符串】	是	[string]	xxxx
order_item>>platform_order_no	平台单号【拆分模式split_mod = 2 时必填，没有值传空字符串】	是	[string]	123
请求示例
{
    "split_mod": 1,
    "order_item": [
        [
            {
                "item_id": 6454374,
                "quantity": 1,
                "pid": 18522,
                "order_item_no": "",
                "msku": "",
                "platform_order_no": 123
            },
            {
                "item_id": 6454375,
                "quantity": 1,
                "pid": 18525,
                "order_item_no": "",
                "msku": "",
                "platform_order_no": 123
            }
        ],
        [
            {
                "item_id": 6454389,
                "quantity": 1,
                "pid": 18518,
                "order_item_no": "",
                "msku": "",
                "platform_order_no": 123
            }
        ]
    ],
    "global_order_no": "103314147463472128"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	操作成功
request_id	请求链路id	是	[string]	8d990d35-8e0a-4103-a670-2d0220235a57.1684316043488
response_time	响应时间	是	[string]	2023-05-17 17:34:06
data	响应数据	是	[object]	
data>>num	成功数量	是	[int]	1
data>>global_order_no	拆分后新的系统单号	是	[array]	["103313809484075008"]
data>>result	拆分后形成的所有系统订单号	是	[array]	
data>>result>>global_order_no	系统单号	是	[string]	103313692201529856
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "937d31c3-1ee6-4b55-aae4-de1c51dea956.1686042075741",
    "response_time": "2023-06-06 17:01:18",
    "data": {
        "num": 1,
        "global_order_no": [
            "103320879312134144"
        ],
        "result": [
            {
                "global_order_no": "103314147463472128"
            },
            {
                "global_order_no": "103320879312134144"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
合并订单
下一章节 
标记订单不发货