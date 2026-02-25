快速出库

创建一个"快速出库"请求的异步任务，如需获取快速出库结果，请在发送请求额外调用获取快速出库结果接口

接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/fastOutbound	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
package	出库包裹信息，最多1000个订单	是	[array]	
package>>global_order_no	系统单号	是	[string]	103256450214985728
package>>wid	出库仓库ID，可通过查询本地仓库接口获得	是	[number]	123
package>>logistics_type_id	物流商ID-物流方式ID，
物流商ID对应 查询已启用的自发货物流方式 接口字段【logistics_provider_id】
物流方式ID对应 查询已启用的自发货物流方式 接口字段【type_id】	是	[string]	123-123
package>>waybill_no	运单号	是	[string]	123
package>>tracking_no	跟踪号	否	[string]	123
package>>weight_unit	重量单位，可选g、kg，默认g	否	[string]	123
package>>real_weight	包裹重量	否	[string]	123
package>>size_unit	尺寸单位，可选mm、cm，默认cm	否	[string]	cm
package>>length	包裹尺寸长	否	[string]	123
package>>width	包裹尺寸宽	否	[string]	123
package>>height	包裹尺寸高	否	[string]	123
package>>fee_weight	包裹计费重	否	[string]	123
package>>logistics_freight	物流运费	否	[string]	123
package>>logistics_freight_currency_code	物流运费币种代码，默认 CNY	否	[string]	CNY
请求示例
{
    "package": [
        {
            "global_order_no": "103256450214985728",
            "wid": 123,
            "logistics_type_id": "123",
            "waybill_no": "123",
            "tracking_no": "123",
            "weight_unit": "g",
            "real_weight": "123",
            "size_unit": "cm",
            "length": "123",
            "width": "123",
            "height": "123",
            "fee_weight": "123",
            "logistics_freight": "123",
            "logistics_freight_currency_code": "CNY"
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
message	消息提示	是	[string]	
data	响应数据	是	[boolean]	true
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
success	是否成功	是	[string]	true
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "data": true,
    "request_id": "609a88ea800d419d9771389b5e751bf0.158.16704888930330001",
    "response_time": "2023-03-07 15:19:49",
    "success": true
}
复制
错误
复制成功
 上一章节
更新订单客服备注
下一章节 
获取快速出库结果