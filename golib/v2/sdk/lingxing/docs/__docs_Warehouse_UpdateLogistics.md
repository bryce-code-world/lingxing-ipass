更新备货单物流信息
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/updateLogistics	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	海外仓备货单号	是	[string]	OWS230700001
logistics_list	物流信息	是	[array]	
logistics_list>>logistics_order_no	物流单号	是	[string]	1233385552
logistics_list>>logistics_money	预估物流费用	是	[string]	100
logistics_list>>logistics_money_unit	预估物流费用币种	是	[string]	USD
logistics_list>>other_money	预估其他费用	是	[string]	200
logistics_list>>other_money_unit	预估其他费用币种	是	[string]	USD
logistics_list>>track_order_no	追踪号	是	[string]	526654523
logistics_list>>other_money_remark	预估费用备注	是	[string]	this is remark
logistics_list>>real_logistics_money	实际物流费用	是	[number]	100
logistics_list>>real_logistics_money_unit	实际物流费用币种	是	[string]	USD
logistics_list>>real_other_money	实际其他费用	是	[number]	100
logistics_list>>real_other_money_unit	实际其他费用币种	是	[string]	USD
logistics_list>>real_other_money_remark	实际其他费用备注	是	[string]	remark
logistics_list_type	物流信息版本：
0：旧版，即将下线
1：新版	是	[int]	0
head_logistics_list	新版头程物流信息（当logistics_list_type 为1时才有意义）	是	[object]	
head_logistics_list>>tax_fee_type	税费分摊方式：
0：产品-计费重
1：产品-实重
2：产品-体积重
3：产品-数量
4：自定义
5：箱子-体积)	是	[int]	1
head_logistics_list>>tracking_list	轨迹信息	是	[list]	
head_logistics_list>>tracking_list>>tracking_no	查询单号	是	[string]	cxdhxxxxxx123
head_logistics_list>>tracking_list>>transport_type	运输类型：
1：快递
2：海运
3：空运
4：其他
	是	[int]	1
head_logistics_list>>tracking_list>>order_type_code	单号类型：【注意：与运输类型联动关系】
1：订舱号
2：提单号
3：箱号
4：其他
5：跟踪单号
6：航班号
当transport_type=1时只能传5
当transport_type=2时只能传1、2、3、4
当transport_type=3时只能传2、6	是	[int]	2
head_logistics_list>>tracking_list>>shippers	承运商【运输类型为海运时才有意义】	是	[string]	coscoxxx
head_logistics_list>>tracking_list>>remark	备注	是	[string]	这是一个备注
head_logistics_list>>estimate_expenses_list	费用明细-预估费用	是	[object]	
head_logistics_list>>estimate_expenses_list>>chargeable_weight	计费重(单位KG)	否	[string]	111.00
head_logistics_list>>estimate_expenses_list>>price	单价	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>logistics_fee	物流费用	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>estimate_expenses_list>>remark	备注	否	[string]	预估费用备注
head_logistics_list>>estimate_expenses_list>>other_fee_arr	预估费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>fee_type_id	其他费id（20位）	是	[string]	241192037543649792
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	1.00
head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list	费用明细-实际费用	是	[object]	
head_logistics_list>>actual_expenses_list>>tax_fee	税费	是	[string]	1
head_logistics_list>>actual_expenses_list>>tax_fee_currency	税费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>chargeable_weight	计费重	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price	单价	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>price_currency	单价币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>logistics_fee	物流费用	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>remark	备注	否	[string]	这个一个实际费用的备注
head_logistics_list>>actual_expenses_list>>other_fee_arr	实际费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
head_logistics_list>>actual_expenses_list>>other_fee_arr>>fee_type_id	其他费id(20位)	是	[string]	241192037543649792
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
请求示例
{
    "overseas_order_no": "OWS230700001",
    "logistics_list": [
        {
            "logistics_order_no": "1233385552",
            "logistics_money": "100",
            "logistics_money_unit": "USD",
            "other_money": "200",
            "other_money_unit": "USD",
            "track_order_no": "526654523",
            "other_money_remark": "this is remark",
            "real_logistics_money": 100,
            "real_logistics_money_unit": "USD",
            "real_other_money": 100,
            "real_other_money_unit": "USD",
            "real_other_money_remark": "remark"
        }
    ],
    "logistics_list_type": 0,
    "head_logistics_list": {
        "tax_fee_type": 1,
        "tracking_list": [{
            "tracking_no": "cxdhxxxxxx123",
            "transport_type": "1",
            "order_type_code": "5",
            "shippers": "coscoxxx",
            "remark": "备注"
        }],
        "estimate_expenses_list": {
            "chargeable_weight": "111.00",
            "price": "1.00",
            "price_currency": "CNY",
            "logistics_fee": "1.00",
            "logistics_fee_currency": "CNY",
            "remark": "预估费用备注",
            "other_fee_arr": [{
                "fee_type_id": "241192037543649792",
                "other_amount": "1.00",
                "other_currency": "CNY"
            }]
        },
        "actual_expenses_list": {
            "tax_fee":"1",
            "tax_fee_currency":"CNY",
            "chargeable_weight": "111.00",
            "price": "111.00",
            "price_currency": "CNY",
            "logistics_fee": "111.00",
            "logistics_fee_currency": "CNY",
            "remark": "实际费用备注",
            "other_fee_arr": [{
                "fee_type_id": "241192037543649792",
                "other_amount": "111.00",
                "other_currency": "CNY"
            }]
        }
    }
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	7E80FC0B-55C1-BAC5-0E4B-7AE59706F051
response_time	响应时间	是	[string]	2023-07-24 15:12:10
data	响应数据	是	[array]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "7E80FC0B-55C1-BAC5-0E4B-7AE59706F051",
    "response_time": "2023-07-24 15:12:10",
    "error_details": [],
    "data": []
}
复制
错误
复制成功
 上一章节
更新备货单
下一章节 
查询备货单装箱信息