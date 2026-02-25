查询海外仓备货单列表

支持查询海外仓备货单列表，对应系统【仓库】>【海外仓备货单】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/listInbound	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
status	状态：
10 待审核
20 已驳回
30 待配货
40 待发货
50 待收货
51 已撤销
60 已完成	否	[int]	60
sub_status	子状态：【仅在待收货状态下生效】
0 全部
1 未收货
2 部分收货	否	[int]	
s_wid	发货仓库id	否	[array]	53
r_wid	收货仓库id	否	[array]	9
overseas_order_no	备货单号	否	[string]	OWS240730001
create_time_from	查询开始日期，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-25
create_time_to	查询结束日期，格式：Y-m-d
当筛选更新时间时，支持Y-m-d或Y-m-d H:i:s	否	[string]	2024-07-31
page_size	分页数量，最大50，默认20	否	[int]	20
page	当前页码，默认1	否	[int]	1
date_type	备货单时间查询类型：【默认create_time】
delivery_time 发货时间
create_time 创建时间
receive_time 收货时间
update_time 更新时间	否	[string]	create_time
is_delete	订单是否删除：
0 未删除【默认】
1 已删除
2 全部	否	[int]	0
请求示例
{
    "s_wid": 53,
    "r_wid": 9,
    "status": 60,
    "overseas_order_no": "OWS240730001",
    "create_time_from": "2024-07-25",
    "create_time_to": "2024-07-31",
    "page_size": 20,
    "page": 1,
    "date_type": "create_time",
    "is_delete": 0
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
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	
data>>overseas_order_no	备货单号	是	[string]	OWS210413018
data>>inbound_order_no	三方入库单号	是	[string]	IC90999921041305
data>>customer_reference_no	客户提交参考号	是	[string]	
data>>s_wid	发货仓id	是	[int]	1785
data>>s_wname	发货仓名称	是	[string]	谷仓三期本地仓
data>>r_wid	收货仓id	是	[int]	1787
data>>r_wname	收货仓名称	是	[string]	4PX 沙田快件二仓
data>>logistics_id	物流方式id	是	[int]	
data>>logistics_name	物流方式	是	[string]	
data>>remark	备注	是	[string]	
data>>status	状态：
10 待审核
20 已驳回
30 待配货
40 待发货
50 待收货
51 已撤销
60 已完成	是	[int]	60
data>>rollback_remark	驳回备注	是	[string]	
data>>is_delete	是否已删除：
0 正常
1 已删除	是	[int]	0
data>>transportation_mode	运输方式ID	是	[string]	241243642281943553
data>>transportation_name	运输方式名称	是	[string]	海卡
data>>uid	创建用户id	是	[int]	252
data>>create_user	创建人	是	[string]	申少田
data>>update_user	最后更新人	是	[string]	
data>>create_time	创建时间	是	[string]	2021-04-13 11:58:56
data>>estimated_time	预计到货时间	是	[string]	
data>>audit_handle_time	审核时间	是	[object]	
data>>send_good_handle_time	发货时间	是	[string]	
data>>receive_good_handle_time	收货时间	是	[string]	
data>>real_delivery_time	实际发货时间	否	[string]	2025-01-29
data>>update_time	最后更新时间	是	[string]	2021-06-01 22:16:27
data>>products	商品信息	是	[array]	
data>>products>>product_id	产品id	是	[int]	1
data>>products>>sku	sku	是	[string]	视觉111
data>>products>>product_name	产品名	是	[string]	儿童玩具A
data>>products>>fnsku	fnsku	是	[string]	B07DC
data>>products>>pic_url	商品图片	是	[string]	https://image.umaicloud.com/xxxxx/11aa08f123b5.jpg
data>>products>>seller_arr	店铺信息【已废弃】	是	[array]	
data>>products>>seller_arr>>sid	店铺id	是	[string]	22
data>>products>>seller_arr>>seller_name	店铺名称	是	[string]	UAC-US1美国
data>>products>>stock_num	备货数量	是	[int]	1
data>>products>>receive_num	收货数量	是	[int]	1
data>>products>>product_valid_num	可用库存	是	[int]	694
data>>products>>sid	店铺id	是	[string]	
data>>products>>product_code	三方产品编码	是	[string]	
data>>products>>remark	商品备注	是	[string]	
data>>products>>batch_record_list	采购信息	否	[array]	
data>>products>>batch_record_list>>seller_id	店铺id	否	[string]	
data>>products>>batch_record_list>>wid	发货仓库ID	否	[number]	
data>>products>>batch_record_list>>fnsku	FNSKU	否	[string]	
data>>products>>batch_record_list>>product_id	产品ID	否	[number]	
data>>products>>batch_record_list>>batch_no	出库批次号	否	[string]	
data>>products>>batch_record_list>>good_num	可用总出库量	否	[number]	
data>>products>>batch_record_list>>batch_order_sn	批次入库单据号	否	[string]	
data>>products>>batch_record_list>>purchase_order_sns	采购单号	否	[array]	
data>>products>>batch_record_list>>supplier_names	供应商	否	[array]	
data>>products>>batch_record_list>>unit_storage_cost	单位出库成本	否	[string]	
data>>products>>batch_record_list>>unit_cost	单位费用	否	[string]	
data>>products>>batch_record_list>>unit_head_range_cost	单位出库头程	否	[string]	
data>>products>>batch_record_list>>unit_purchase_price	单位采购成本	否	[string]	
data>>products>>batch_record_list>>storage_good_num	批次可用出库量	否	[number]	
data>>logistics	物流数据	是	[array]	
data>>logistics>>logistics_order_no	物流单号	否	[string]	
data>>logistics>>logistics_money	预估物流费用	否	[string]	
data>>logistics>>logistics_money_unit	预估物流费用币种	否	[string]	
data>>logistics>>other_money	预估其他费用	否	[string]	
data>>logistics>>other_money_unit	预估其他费用币种	否	[string]	
data>>logistics>>track_order_no	追踪号	否	[string]	
data>>logistics>>other_money_remark	预估费用备注	否	[string]	
data>>logistics>>real_logistics_money	实际物流费用	否	[number]	
data>>logistics>>real_logistics_money_unit	实际物流费用币种	否	[string]	
data>>logistics>>real_other_money	实际其他费用	否	[number]	
data>>logistics>>real_other_money_unit	实际其他费用币种	否	[string]	
data>>logistics>>real_other_money_remark	实际其他费用备注	否	[string]	
data>>logistics>>wool_id	物流记录id	否	[string]	
data>>logistics_list_type	物流信息版本：【默认0】
0 旧版
1 新版	否	[int]	0
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
data>>head_logistics_list>>estimate_expenses_list>>chargeable_weight	计费重(单位KG)	否	[string]	111.00
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
data>>head_logistics_list>>actual_expenses_list>>chargeable_weight	计费重	是	[string]	111.00
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
total	总数	是	[int]	92
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "0867EB04-4D5F-4F0B-BA55-6075DFC32B94",
    "response_time": "2024-08-01 16:05:00",
    "data": [
        {
            "overseas_order_no": "OWS240730001",
            "inbound_order_no": "",
            "customer_reference_no": "",
            "s_wid": 53,
            "s_wname": "义乌",
            "r_wid": 9,
            "r_wname": "小徐仓",
            "logistics_id": 502,
            "logistics_name": "765432123456789(顺丰物流3456788)",
            "remark": "",
            "status": 60,
            "rollback_remark": "",
            "transportation_mode":"241243642281943553",
            "transportation_name":"海卡",
            "is_delete": 0,
            "uid": 33,
            "update_user": "朱蒙蒙",
            "create_user": "朱蒙蒙",
            "audit_handle_time": "2024-07-30 11:05:20",
            "send_good_handle_time": "2024-07-30 11:05:28",
            "receive_good_handle_time": "2024-07-30 11:05:40",
            "create_time": "2024-07-30 11:05:17",
            "real_delivery_time": "2024-07-30",
            "estimated_time": "2024-07-30",
            "products": [
                {
                    "product_id": 33080,
                    "sku": "abc003-BWY",
                    "product_code": "",
                    "product_name": "abc003-BWY-xin",
                    "fnsku": "",
                    "pic_url": "",
                    "sid": "0",
                    "seller_arr": [],
                    "stock_num": 10,
                    "receive_num": 10,
                    "product_valid_num": null,
                    "remark": "",
                    "s_wid": 53,
                    "s_wname": ""
                }
            ],
            "logistics": [
                {
                    "logistics_order_no": "",
                    "track_order_no": "",
                    "logistics_money": null,
                    "logistics_money_unit": "CNY",
                    "other_money": null,
                    "other_money_unit": "CNY",
                    "other_money_remark": "",
                    "real_logistics_money": null,
                    "real_logistics_money_unit": "CNY",
                    "real_other_money": null,
                    "real_other_money_unit": "CNY",
                    "real_other_money_remark": "",
                    "wool_id": "242469427479240704"
                }
            ],
            "logistics_list_type": 0,
            "head_logistics_list": {
                "tracking_list": [],
                "estimate_expenses_list": null,
                "actual_expenses_list": null
            },
            "update_time": "2024-07-30 11:05:40"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
创建已完成的成本补录单
下一章节 
查询海外仓备货单详情