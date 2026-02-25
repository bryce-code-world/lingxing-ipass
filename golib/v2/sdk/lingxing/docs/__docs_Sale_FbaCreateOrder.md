创建亚马逊多渠道订单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/order/amzod/api/createOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_name	店铺名	是	[string]	8speed-US
country	店铺国家	是	[string]	美国
order_id	订单号	是	[string]	1000796
is_blank_box	是否使用无品牌包装箱（“是”/“否”，默认为“否”）	否	[string]	否
is_block_amzl	是否阻止亚马逊物流（“是”/“否”，默认为“否”）	否	[string]	否
receiver	收件人	是	[string]	jack
country_code	收货地址国家/地区（输入国家/地区简码）	是	[string]	US
region	地区	是	[string]	AL-Alabama
city	城市（日本市场非必填，其他市场必填）	否	[string]	Autaugaville
address1	地址1	是	[string]	Autauga
address2	地址2	否	[string]	Autauga
postcode	邮编	是	[string]	36003
phone_number	电话号码	否	[string]	123456
buyers_mailbox	买家邮箱	是	[string]	123456@gmail.com
order_id_for_packing	装箱单-订单号	是	[string]	1802
date_for_packing	装箱单-订单日期	是	[string]	2022-12-09或2025-04-09T14:30:45
remark_for_packing	装箱单-装箱单备注	否	[string]	Thank you for your order
delivery_operation	配送操作（“立即配送”/“保留订单”）	否	[string]	保留订单
delivery_service	配送服务（“标准配送”/“加急配送”/“优先配送”）	否	[string]	标准配送
remark	订单备注	否	[string]	备注
item_list	商品列表	是	[array]	
item_list>>msku	MSKU	是	[string]	msku123
item_list>>quantity_shipped	发货量	是	[number]	1
item_list>>declared_value	申报价值	否	[number]	1.23
item_list>>declared_currency	申报货币	否	[string]	USD
请求示例
{
    "store_name": "8speed-US1",
    "country": "美国",
    "order_id": "1000796",
    "is_blank_box": "否",
    "is_block_amzl": "否",
    "receiver": "jack",
    "country_code": "US",
    "region": "AL-Alabama",
    "city": "Autaugaville",
    "address1": "Autauga",
    "address2": "Autauga",
    "postcode": "36003",
    "phone_number": "123456",
    "buyers_mailbox": "123456@gmail.com",
    "order_id_for_packing": "1802",
    "date_for_packing": "2022-12-09",
    "remark_for_packing": "Thank you for your order",
    "delivery_operation": "保留订单",
    "delivery_service": "标准配送",
    "remark": "备注",
    "item_list": [{
        "msku": "msku123",
        "quantity_shipped": 1,
        "declared_value": 1.23,
        "declared_currency": "USD"
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	错误信息	是	[array]	
response_time	响应时间	是	[string]	2023-04-27 09:28:29
request_id	请求链路id	是	[string]	212149f55784429aaa69fe7f0cab12e7.100.16825589067040171
total	总数	是	[int]	0
data	响应数据	是	[array]	 
返回成功示例
 {
    "code": 0,
    "data": [],
    "error_details": [],
    "message": "操作成功",
    "response_time": "2023-04-27 09:28:29",
    "request_id": "212149f55784429aaa69fe7f0cab12e7.100.16825589067040171",
    "total": 0
}
复制
错误
复制成功
返回失败示例
 {
    "code": 500,
    "data": null,
    "error_details": [
        "订单号已经存在"
    ],
    "message": "失败",
    "response_time": "2023-04-27 09:36:32",
    "request_id": "212149f55784429aaa69fe7f0cab12e7.109.16825593920660185",
    "total": 0
}
复制
错误
复制成功
 上一章节
查询亚马逊多渠道订单详情-退货换货信息
销售
下一章节 
查询售后订单列表
销售