获取快速出库结果
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/getFastOutboundResult	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
global_order_no	系统单号数组，最大1000单	是	[array]	["103256450214985728","103256450209420800","1003","1004"]
请求示例
{
    "global_order_no": [
        "103432007211410432",
        "103256450209420800",
        "1003",
        "1004"
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	
data	响应数据	是	[array]	
data>>success	快速出库成功	是	[array]	
data>>success>>global_order_no	订单号	是	[string]	No1002
data>>success>>wo_number	销售出库单号	是	[string]	Wo1002
data>>failure	快速出库失败	是	[array]	
data>>failure>>global_order_no	订单号	是	[string]	No1001
data>>failure>>error_message	失败原因	是	[string]	正在处理
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
success		是	[string]	true
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "data": {
        "success": [
            {
                "global_order_no": "No1002",
                "wo_number": "Wo1002"
            }
        ],
        "failure": [
            {
                "global_order_no": "No1001",
                "error_message": "正在处理"
            },
            {
                "global_order_no": "No1003",
                "error_message": "库存不足"
            },
            {
                "global_order_no": "No1006",
                "error_message": "订单没提交快速出库"
            }
        ]
    },
    "request_id": "49245bf4-454a-4057-b2d4-6e566a2554f8.1669714091555",
    "response_time": "2023-03-07 15:19:49",
    "success": true
}
复制
错误
复制成功
 上一章节
快速出库
下一章节 
查询销量统计列表(将于04.30下线)