更新发货单物流信息

用于更新发货单内的【物流信息】

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/shipment/updateListLogistics	HTTPS	POST	2
请求参数
参数名	说明	必填	类型	示例
data	参数数组	是	[array]	
data>>order_sn	发货单号	是	[string]	SP220120057
data>>expected_arrival_date	到货时间，格式：Y-m-d	否	[string]	2022-05-06
data>>etd_date	开船时间，格式：Y-m-d	否	[string]	2022-05-06
data>>eta_date	预计到港时间，格式：Y-m-d	否	[string]	2022-05-06
data>>delivery_date	实际妥投时间，格式：Y-m-d	否	[string]	2022-05-06
data>>actual_shipment_time	实际发货时间，格式：Y-m-d	否	[string]	2022-05-06
data>>tax_fee_type	实际税费分配方式：【默认0】
0 产品-计费重
1 产品-实重
2 产品-体积重
3 产品-数量
4 自定义
5 箱子-体积/体积重
6 箱子-计费重
7 箱子-实重
8 产品-根据清关单价*税率占比	否	[int]	0
data>>logistics_channel_id	物流渠道id：按计费重分摊时必填，以获取材积参数用于计算
查询头程物流渠道列表接口对应字段【id】	否	[int]	
data>>logistics_list_type	物流信息版本：
1 新版	否	[int]	1
data>>head_logistics_list	新版头程物流信息
【对应 logistics_list_type = 1】
【注意：新版头程物流数据为覆盖式更新，包括tracking_list、estimate_expenses_list、actual_expenses_list，不传或者传空也会置空】	否	[object]	
data>>head_logistics_list>>tracking_list	轨迹信息数组	是	[array]	
data>>head_logistics_list>>tracking_list>>tracking_no	查询单号	否	[string]	cxdhxxxxxx123
data>>head_logistics_list>>tracking_list>>transport_type	运输类型：
1 快递
2 海运
3 空运
4 其他	是	[int]	1
data>>head_logistics_list>>tracking_list>>order_type_code	单号类型：【注意：与运输类型联动关系】
1 订舱号
2 提单号
3 箱号
4 其他
5 跟踪单号
6航班号
当transport_type=1时只能传5
当transport_type=2时只能传1、2、3、4
当transport_type=3时只能传2、6
当transport_type=4只能传空	是	[int]	5
data>>head_logistics_list>>tracking_list>>shippers	承运商，运输类型为海运时才有意义：
获取发货单头程物流-承运商信息接口获取	否	[string]	coscoxxx
data>>head_logistics_list>>tracking_list>>remark	备注	否	[string]	这是一个备注
data>>head_logistics_list>>estimate_expenses_list	费用明细-预估费用	是	[object]	
data>>head_logistics_list>>estimate_expenses_list>>chargeable_weight	计费重(单位KG)【废弃字段】	否	[string]	111.00
data>>head_logistics_list>>estimate_expenses_list>>price	单价	是	[string]	1.00
data>>head_logistics_list>>estimate_expenses_list>>price_currency	单价币种	是	[string]	CNY
data>>head_logistics_list>>estimate_expenses_list>>logistics_fee	物流费用	是	[string]	1.00
data>>head_logistics_list>>estimate_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
data>>head_logistics_list>>estimate_expenses_list>>remark	备注	否	[string]	预估费用备注
data>>head_logistics_list>>estimate_expenses_list>>other_fee_arr	预估费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
data>>head_logistics_list>>estimate_expenses_list>>other_fee_arr>>fee_type_id	其他费id（20位）	是	[string]	241192037543649792
data>>head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	1.00
data>>head_logistics_list>>estimate_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
data>>head_logistics_list>>actual_expenses_list	费用明细-实际费用	是	[object]	
data>>head_logistics_list>>actual_expenses_list>>tax_fee	税费	是	[string]	1
data>>head_logistics_list>>actual_expenses_list>>tax_fee_currency	税费币种	是	[string]	CNY
head_logistics_list>>actual_expenses_list>>chargeable_weight	计费重【废弃字段】	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>weight	实重（单位：KG）	是	[string]	111.00
head_logistics_list>>actual_expenses_list>>volume	体积（单位：m³）	是	[string]	111.00
data>>head_logistics_list>>actual_expenses_list>>price	单价	是	[string]	111.00
data>>head_logistics_list>>actual_expenses_list>>price_currency	单价币种	是	[string]	CNY
data>>head_logistics_list>>actual_expenses_list>>logistics_fee	物流费用	是	[string]	111.00
data>>head_logistics_list>>actual_expenses_list>>logistics_fee_currency	物流费用币种	是	[string]	CNY
data>>head_logistics_list>>actual_expenses_list>>remark	备注	否	[string]	这个一个实际费用的备注
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr	实际费用-其他费：
获取发货单头程物流-其他费类型接口获取	是	[array]	
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>fee_type_id	其他费id	是	[string]	241192037543649792
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_amount	其他费金额	是	[string]	111.00
data>>head_logistics_list>>actual_expenses_list>>other_fee_arr>>other_currency	其他费币种	是	[string]	CNY
data>>logistics_list	旧版物流信息，即将下线	是	[array]	
data>>logistics_list>>tracking_number	物流商单号	是	[string]	22
data>>logistics_list>>replace_tracking_number	跟踪号	是	[string]	20215486
data>>logistics_list>>transportation_cost	实际物流费用，精度是小数点后2位	否	[number]	
data>>logistics_list>>transportation_currency	实际物流费用币种，费用填写时必填	否	[string]	CNY
data>>logistics_list>>other_cost	实际其他费用，精度是小数点后2位	否	[number]	
data>>logistics_list>>other_currency	实际其他费用币种，费用填写时必填	否	[string]	CNY
data>>logistics_list>>other_cost_remark	其他费用备注	否	[string]	其他费用备注
data>>logistics_list>>predicted_transportation_cost	预估物流费用，精度是小数点后2位	否	[number]	44
data>>logistics_list>>predicted_transportation_currency	预估物流费用币种，费用填写时必填	否	[string]	CNY
data>>logistics_list>>predicted_other_cost	预估其他费用，精度是小数点后2位	否	[number]	12
data>>logistics_list>>predicted_other_currency	预估其他费用币种，费用填写时必填	否	[string]	CNY
请求示例
{
    "data": [{
        "order_sn": "SP220120057",
        "expected_arrival_date": "2022-05-06",
        "etd_date": "2022-05-06",
        "eta_date": "2022-05-06",
        "delivery_date": "2022-05-06",
        "actual_shipment_time": "2022-05-06",
        "tax_fee_type":"0",
        "logistics_list": [{
            "tracking_number": "22",
            "replace_tracking_number": "20215486",
            "transportation_cost": 0,
            "transportation_currency": "CNY",
            "other_cost": 0,
            "other_currency": "CNY",
            "other_cost_remark": "其他费用备注",
            "predicted_transportation_cost": 44,
            "predicted_transportation_currency": "CNY",
            "predicted_other_cost": 12,
            "predicted_other_currency": "CNY"
        }],
        "logistics_list_type": 0,
        "head_logistics_list": {
            "tracking_list": [{
                "tracking_no": "cxdhxxxxxx123",
                "transport_type": "1",
                "order_type_code": "5",
                "shippers": "coscoxxx",
                "remark": "备注"
            }],
            "estimate_expenses_list": {
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
                "weight": "111.00",
                "volume": "111.00",
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
    }]
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
request_id	请求链路id	是	[string]	426951E6-1B62-1B12-90AF-07648AACC512
response_time	响应时间	是	[string]	2022-02-24 20:04:36
total	总数	是	[int]	0
data	响应数据	是	[array]	 
 上一章节
作废发货单
FBA
下一章节 
更新发货单自定义成本
FBA