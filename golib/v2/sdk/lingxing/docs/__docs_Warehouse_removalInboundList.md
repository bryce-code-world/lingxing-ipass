查询移除入库单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/removalInbound/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
status	订单状态：
1 待提交-未提交
2 待提交-提交中
3 待提交-失败
4 待收货-未收货
5 待收货-异常
6 已完成
7 已作废	否	[int]	1
start_date	开始日期【发货日期，双闭区间】	否	[string]	2022-09-10
end_date	结束日期【发货日期，双闭区间】	否	[string]	2022-10-10
order_no	移除入库单号	否	[array]	["OWS231029001"]
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限1000	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "order_no": ["OWR220715000001"],
    "start_date": "2022-06-01",
    "end_date": "2024-07-30",
    "status": 7
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	0F63C7DB-D415-79A2-C086-1079A734C446
response_time	响应时间	是	[string]	2023-10-20 11:21:16
total	总数	是	[int]	3
data	响应数据	是	[array]	
data>>id	记录id	是	[int]	16
data>>sid	店铺id	是	[int]	1
data>>sid_name	店铺名称	是	[string]	店铺1 美国
data>>order_no	移除入库单号	是	[string]	OWR230725000004
data>>removal_order_no	移除订单号	是	[string]	20230601002
data>>wid	入库仓库id	是	[int]	2582
data>>wid_name	入库仓库名称	是	[string]	赛诺仓 美国仓
data>>estimated_arrival_time	预计到货时间	是	[string]	2023-07-27
data>>shippment_time	发货时间	是	[string]	2023-07-20T00:00:00-00:03
data>>shipper	承运商	是	[string]	XCET6
data>>delivery_no	跟踪号	是	[string]	XCET6_BHM1_071921_3
data>>order_status	订单状态：
1 待提交-未提交
2 待提交-提交中
3 待提交-失败
4 待收货-未收货
5 待收货-异常
6 已完成
7 已作废	是	[int]	6
data>>remark	备注	是	[string]	
data>>uid	提交人id	是	[int]	129057
data>>uid_name	提交人姓名【同submiter】	是	[string]	丁三
data>>submiter	提交人姓名	是	[string]	丁三
data>>submit	是否提交到三方海外仓：1 否，2 是	是	[int]	2
data>>inbound_order_sns	关联入库单号	是	[array]	["IB230725006"]
data>>address	仓库收货地址信息	是	[object]	
data>>address>>sid	店铺id	是	[int]	1
data>>address>>order_id	移除入库单id	是	[int]	16
data>>address>>order_no	移除入库单号	是	[string]	OWR230725000004
data>>address>>address_line1	详细地址1	是	[string]	1
data>>address>>address_line2	详细地址2	是	[string]	COLOSSEUM DRIVE HOUGHTON REGIS
data>>address>>address_line3	详细地址3	是	[string]	
data>>address>>country_code	国家代码	是	[string]	GB
data>>address>>state_or_province	省州	是	[string]	
data>>address>>city	城市	是	[string]	DUNSTABLE
data>>address>>county	县	是	[string]	Bedfordshire
data>>address>>district	区	是	[string]	
data>>address>>postal_code	邮编	是	[string]	LU5 6QD
data>>address>>name	收件人名称	是	[string]	Anita
data>>address>>address_str	格式化地址	是	[string]	Anita,LU5 6QD,xxx,xxx,xxx
data>>product	产品信息	是	[array]	
data>>product>>id	子项记录id	是	[int]	12
data>>product>>sid	店铺id	是	[int]	1
data>>product>>order_id	移除入库单id	是	[int]	16
data>>product>>order_no	移除入库单号	是	[string]	OWR230725000004
data>>product>>msku	MSKU	是	[string]	MSKU8C06F3C
data>>product>>fnsku	FNSKU	是	[string]	FNA9F5E33
data>>product>>sku	SKU	是	[string]	SKU1991FC9
data>>product>>product_id	本地产品id	是	[int]	1
data>>product>>product_name	产品名称	是	[string]	[演示数据]适用于iPad的手写笔带手掌柜绝
data>>product>>pic_url	产品图片url	是	[string]	
data>>product>>third_product_name	三方产品名称	是	[string]	苹果电脑
data>>product>>third_code	三方产品编码	是	[string]	1003
data>>product>>sellable_num	可售数量	是	[int]	100
data>>product>>unsellable_num	不可售数量	是	[int]	70
data>>product>>declare_num	申报量	是	[int]	170
data>>product>>avaliable_num	可用量	是	[int]	170
data>>product>>defective_num	次品量	是	[int]	0
data>>product>>recieve_num	收货数量	是	[int]	170
data>>product>>differences	待收货量	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "41A01120-0E4A-96D3-42DC-CE313EFE7808",
    "response_time": "2023-10-20 14:28:59",
    "total": 1
    "data": [
        {
            "id": 2,
            "sid": 110,
            "sid_name": "8p-CA-1",
            "order_no": "OWR220715000001",
            "removal_order_no": "R-22062101",
            "wid": 9,
            "wid_name": "小徐仓",
            "estimated_arrival_time": "2022-07-03",
            "shippment_time": "2022-06-26T10:28:29-05:00",
            "shipper": "USPS",
            "delivery_no": "9374889763012717921279",
            "order_status": 7,
            "remark": "备注",
            "uid": 58,
            "uid_name": "xx",
            "submit": 1,
            "submiter": "xx",
            "inbound_order_sns": [],
            "address": {
                "sid": 110,
                "order_id": 2,
                "order_no": "OWR220715000001",
                "address_line1": "5600 Harvey St",
                "address_line2": "",
                "address_line3": "",
                "country_code": "US",
                "state_or_province": "MI",
                "city": "Michigan",
                "county": "",
                "district": "",
                "postal_code": "49444",
                "name": "Michael Morton",
                "address_str": "Michael Morton,xxxx,xxxx"
            },
            "product": [
                {
                    "id": 2,
                    "sid": 110,
                    "order_id": 2,
                    "order_no": "OWR220715000001",
                    "msku": "Black_ Head_Rope",
                    "fnsku": "B09MT39000",
                    "sku": "AK-000002",
                    "product_id": 24136,
                    "product_name": "AK-000002",
                    "pic_url": "",
                    "sellable_num": 1,
                    "unsellable_num": 0,
                    "third_product_name": "",
                    "third_code": "",
                    "declare_num": 1,
                    "avaliable_num": 0,
                    "defective_num": 0,
                    "recieve_num": 0,
                    "differences": 1
                }
            ]
        }
    ]
}
复制
错误
复制成功
 上一章节
获取第三方箱唛
下一章节 
备货单分配库存