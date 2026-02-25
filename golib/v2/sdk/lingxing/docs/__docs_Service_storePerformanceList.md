查询店铺绩效列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/storeTarget/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
search_field_time	搜索时间类型：
pull_date 报表获取时间
update_date 更新时间	否	[string]	pull_date
search_time	搜索时间，格式：Y-m-d	否	[string]	2024-08-05
sids	店铺id，多个使用英文逗号分隔 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[string]	135,141,145,148
anomaly_indicator	异常指标：
commodity_policy_compliance 商品政策合规性
on_time_delivery 准时交货率
valid_tracking 有效追踪率
pre_fulfillment_cancellation 预配送取消率
late_shipment 迟发率
invoice_defect 发票缺陷率
fba_order_with_defect FBA订单缺陷率
order_with_defect FBM订单缺陷率	否	[array]	["commodity_policy_compliance"]
请求示例
{
    "offset": 0,
    "length": 20,
    "search_field_time": "pull_date",
    "search_time": "2024-08-05",
    "sids": "135,141,145,148",
    "anomaly_indicator": ["commodity_policy_compliance"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	xxxx.xxxx
response_time	响应时间	是	[string]	2023-08-16 09:41:39
total	总数	是	[int]	100
data	响应数据	是	[array]	
data>>sid	店铺id	是	[int]	31
data>>pull_date	报表获取日期	是	[string]	2022-05-09
data>>update_date	报表数据更新时间	是	[string]	2022-05-09 13:01:38
data>>order_with_defect	fbm订单缺陷率，百分比	是	[string]	1.20
data>>return_dissatisfaction	退货不满意度，百分比	是	[string]	-
data>>late_shipment	迟发率，百分比	是	[string]	0.00
data>>pre_fulfillment_cancellation	预配送取消率，百分比	是	[string]	10.00
data>>valid_tracking	有效追踪率，百分比	是	[string]	-
data>>on_time_delivery	准时交货率，百分比	是	[string]	97.00
data>>commodity_policy_compliance	商品政策合规性	是	[string]	0.00
data>>fba_order_with_defect	fba订单缺陷率，百分比	是	[string]	1.00
data>>invoice_defect	发票缺陷率，百分比	是	[string]	5.00
data>>ahr_score	账户状况分数	是	[string]	1000.00
data>>ahr_status	账户状况评级	是	[string]	GREAT
返回示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "48a3dba39bd94518b9109e1965780875.1692949438126",
    "response_time": "2023-08-25 15:43:58",
    "total": 7,
    "data": [
        {
            "sid": 31,
            "pull_date": "2022-05-09",
            "update_date": "2022-05-09 13:01:38",
            "order_with_defect": "1.00",
            "return_dissatisfaction": "-",
            "late_shipment": "5.00",
            "pre_fulfillment_cancellation": "97.00",
            "valid_tracking": "-",
            "on_time_delivery": "-",
            "commodity_policy_compliance": "0",
            "fba_order_with_defect": "88.00",
            "invoice_defect": "55.00",
            "ahr_score": "1000.00",
            "ahr_status": "GREAT"
        }
    ]
}
复制
错误
复制成功
 上一章节
查询售后工单列表
下一章节 
查询店铺绩效详情