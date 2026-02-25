收货单快捷入库
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/deliveryReceipt/PurchaseReceiptOrder/fastReceive	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	收货单号	是	[string]	CR230520001
expect_arrival_time	预计收货时间，不传时默认取自收货单	否	[string]	2023-06-01
custom_receive_time	自定义收货时间， 自定义日期须早于请求当天日期	否	[string]	2023-05-30
logistics_company	物流商，不传时默认取自收货单	否	[string]	京东快递
logistics_order_no	物流单号，仅支持字母、数字、下划线、中横线，不传时默认取自收货单	否	[string]	JD100001
shipping_cost	运费，仅支持2位小数，不传时默认取自收货单	否	[number]	10
other_fee	其他费用，仅支持2位小数，不传时默认取自收货单	否	[number]	2
remark	备注，最大支持255个字符，不传时默认取自收货单	否	[string]	加急
item_list	收货明细	是	[array]	
item_list>>id	收货单子项id，查询收货单列表接口对应字段【item_id】	是	[int]	134624
item_list>>product_good_num	良品量，总良品量+总次品量必须大于0	是	[int]	10
item_list>>product_bad_num	次品量，总良品量+总次品量必须大于0	是	[int]	0
item_list>>remark	备注，最大支持255个字符，不传时默认取自收货单	否	[string]	易碎产品
请求示例
{
    "order_sn": "CR230919006",
    "expect_arrival_time": "2023-09-20",
    "custom_receive_time": "2023-09-18",
    "logistics_company": "京东快递",
    "logistics_order_no": "JD100001",
    "shipping_cost": 10,
    "other_fee": 2,
    "remark": "加急",
    "item_list": [
        {
            "id": 6652,
            "product_good_num": 1,
            "product_bad_num": 0,
            "remark": "易碎产品"
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
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	5295A11B-7151-0419-D407-A4C24F9EF1E7
response_time	响应时间	是	[string]	2023-09-19 10:45:58
data	响应数据	是	[array]	
total	总数	是	[int]	0
 上一章节
收货单到货
下一章节 
查询销售退货单列表