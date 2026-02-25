亚马逊订单提交标发

当前接口为提交亚马逊订单标发任务，单个订单标发结果非当前接口同步返回
亚马逊订单标发为异步任务，请在提交标发后使用【task_id】通过接口查询亚马逊标发结果获取已提交任务的标发结果

接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/submitFulfillment	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
region	店铺注册所属区域：仅支持 NA、EU、FE
【对应区域值支持国家见附加说明】	是	[string]	NA
seller_id	亚马逊店铺id ,对应查询亚马逊店铺列表接口对应字段【seller_id】	是	[string]	A1MQMW3JWPNCBX
marketplace_id	市场id	是	[string]	xxxxxxxxxxxxxxxx
order_list	提交标发数据列表	是	[array]	
order_list>>id	标发记录id【20位以内的正整数，对应标发结果中的message_id】	是	[int]	2725
order_list>>platform_order_no	平台单号	是	[string]	xhp_1121_0111
order_list>>tracking_no	物流追踪号	是	[string]	12344
order_list>>carrier_code	承运商code	是	[string]	360lion
order_list>>carrier_name	承运商名称	是	[string]	360lion
order_list>>shipping_service	配送服务，没有可以传空字符串	是	[string]	usps
order_list>>order_item	商品信息	是	[array]	
order_list>>order_item>>order_item_no	亚马逊商品行id	是	[string]	763432323
order_list>>order_item>>quantity	商品数量	是	[int]	1
请求示例
{
    "region": "fe",
    "seller_id": "A4IBC1X76HVPJ",
    "marketplace_id": "A39IBJ37TRP1C6",
    "order_list": [
        {
            "id": "103313699178382336",
            "platform_order_no": "250-8182459-1815841",
            "tracking_no": "GMQ6012738",
            "carrier_code": "",
            "carrier_name": "Australia Post",
            "shipping_service": "",
            "order_item": [
                {
                    "order_item_no": "33051859",
                    "quantity": 1
                }
            ]
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	操作成功
request_id	请求链路id	是	[string]	0a618db6bfe24f25ae13f5b217b8a340.1676370043019
response_time	响应时间	是	[string]	2023-02-15 13:01:08
data	响应数据	是	[object]	
data>>task_id	任务id	是	[string]	1003117
data>>task_status	任务状态	是	[string]	in_process
data>>failure_reason	提交失败原因	是	[string]	null
返回成功示例

{
    "code": 0,
    "message": "操作成功",
    "request_id": "fe51fee0ff7e4e8797abfb549d3a63a8.159.16843756058584983",
    "response_time": "2023-02-14 18:20:44",    
    "data": {
        "task_id": "1003094",
        "task_status": "in_process",
        "failure_reason": null
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "orderList[0].carrierCode承运商code不能为空",
    "request_id": "9ce3945f8a1f4ca4907b5d3413c11bab.157.16758503506230001",
    "response_time": "2023-02-14 18:10:04",
    "data": null
}
复制
错误
复制成功
附加说明

region 说明如下：

NA【North America】：包括国家为 CA、US、MX、BR
EU【Europe】：包括国家为 ES、UK、FR、BE、NL、DE、IT、SE、ZA、PL、EG、TR、SA、AE、IN
FE【Far East】：包括国家为 SG、AU、JP

task_status 说明如下：

pending：feed任务刚创建时的状态，此时任务在排队等待发起状态；
in_process：feed任务已经被消费者调度执行，处于该状态时将开始与亚马逊进行交互，此时会出现如下三个阶段：
submit_feed【提交feed阶段】、listen_notification【监听阶段】、get_feed_result【获取结果阶段】
fatal：feed任务最终执行失败；
done：feed任务最终执行成功；
cancelled：feed服务的feed任务被取消；
 上一章节
查询亚马逊自发货订单详情
销售
下一章节 
查询亚马逊标发结果
销售