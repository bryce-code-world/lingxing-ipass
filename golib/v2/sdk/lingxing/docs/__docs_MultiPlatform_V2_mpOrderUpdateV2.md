更新订单客服备注
接口信息

支持更新订单内的【客服备注】

API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/setRemark	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
orders	系统订单列表	是	[array]	
orders>>global_order_no	系统单号	是	[string]	103430576541848064
orders>>remark	客服备注文本	是	[string]	这是订单客服备注
orders>>remark_is_append	是否追加原有备注：【默认true】
true 追加客服备注
false 替换客服备注	否	[string]	true
请求示例
{
    "orders": [
        {
            "global_order_no": "103430576541848064",
            "remark": "这是订单客服备注",
            "remark_is_append": true
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	10002：订单更新成功
10000：订单更新失败
10001：订单更新部分成功	是	[int]	10001
message	消息提示	是	[string]	订单更新部分成功
data>>error_details	错误信息	是	[array]	[{global_order_no:xx,error_message:xxx},..]
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
success	是否成功	是	[string]	true
返回成功示例
{
    "code": 10002,
    "data": {
        "error_details": []
    },
    "response_time": "2024-07-23 15:14:59",
    "message": "订单更新成功",
    "request_id": "87e5340c55b44347bdc53d8ce4f8f0a8.152716.17217188992811689"
}
复制
错误
复制成功
 上一章节
编辑/更新自发货订单
下一章节 
快速出库