查询质检单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/deliveryReceipt/ReceiptOrderQc/getOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
date_type	查询时间类型：1 质检时间，2 收货时间，3 创建时间	否	[int]	3
start_date	开始时间	否	[string]	2024-07-30
end_date	结束时间	否	[string]	2024-07-30
qc_sns	质检单号，多个使用英文逗号分隔	否	[string]	QC240730019
status	状态，多个使用英文逗号分隔：
0 待质检
1 已质检
2 已免检
10 已质检（撤销）
20 已免检（撤销）	否	[string]	0
wid	仓库id，多个用英文逗号分隔	否	[string]	1643
offset	分页偏移量，默认为0	否	[int]	0
length	分页长度，默认为200，上限500	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": 1643,
    "date_type": 3,
    "status": 0,
    "start_date": "2024-07-30",
    "end_date": "2024-07-30",
    "qc_sns": "QC240730019"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[string]	
request_id	请求链路id	是	[string]	C125A617-CDB0-DFDC-5D63-78BB84D27821
response_time	响应时间	是	[string]	2022-10-18 11:23:47
data	响应数据	是	[object]	
data>>total	总数	是	[int]	15243
data>>list	列表	是	[array]	
data>>list>>qc_sn	质检单号	是	[string]	QC220927002
data>>list>>order_type	订单类型：1 采购订单，2 委外订单	是	[int]	1
data>>list>>qc_type	质检类型：1 仓库质检，2 预检，3 免检	是	[int]	1
data>>list>>qc_method	质检方式：1 抽检，2 全检	是	[int]	1
data>>list>>create_time	创建时间	是	[string]	2022-09-27 10:54:09
data>>list>>status	状态：
0 待质检
1 已质检
2 已免检
10 已质检（撤销）
20 已免检（撤销）	是	[int]	1
data>>list>>order_sn	来源单号	是	[string]	PO220927002
data>>list>>opt_uid	采购员id	是	[int]	1
data>>list>>opt_realname	采购员	是	[string]	xx
data>>list>>receive_uid	收货人id	是	[int]	1
data>>list>>receive_realname	收货人	是	[string]	xx
data>>list>>receive_time	到货时间	是	[string]2022-09-27 10:53:59
data>>list>>qc_uid	质检人id	是	[int]	1
data>>list>>qc_realname	质检人	是	[string]	xx
data>>list>>qc_time	质检时间	是	[string]	2022-09-27 10:54:09
data>>list>>wid	仓库id	是	[int]	1
data>>list>>supplier_id	供应商id	是	[int]	1
data>>list>>order_item_id	采购单子项id	是	[int]	21904
data>>list>>delivery_order_sn	收货单号	是	[string]	CR230824001
data>>list>>delivery_item_id	收货单子项id	是	[int]	19671
data>>list>>sku	SKU	是	[string]	1313355
data>>list>>product_name	品名	是	[string]	[演示数据]USB壁式充电器2.1A/5V双端口充电器
data>>list>>fnsku	FNSKU	是	[string]	
data>>list>>seller_id	店铺id	是	[int]	0
data>>list>>product_receive_num	质检量	是	[int]	164
data>>list>>qc_num	抽检量	是	[int]	164
data>>list>>qc_bad_num	抽检次品量	是	[int]	0
data>>list>>qc_rate_pass	抽检合格率	是	[string]	100%
data>>list>>qc_rate	抽检比例	是	[string]	100%
data>>list>>product_good_num	总良品量	是	[int]	164
data>>list>>product_bad_num	总次品量	是	[int]	0
data>>list>>whb_code_good	可用仓位	是	[string]	可用暂存
data>>list>>whb_code_bad	次品仓位	是	[string]	次品暂存
data>>list>>qc_remark	备注	是	[string]	xxx
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "796041B5-8B02-8238-AB93-0790DDB59522",
    "response_time": "2024-07-31 16:09:35",
    "data": {
        "total": 1,
        "list": [
            {
                "qc_sn": "QC240730019",
                "order_type": 1,
                "qc_type": 1,
                "qc_method": 1,
                "create_time": "2024-07-30 21:24:30",
                "status": 0,
                "order_sn": "PO240717004",
                "opt_uid": 10580661,
                "opt_realname": "黄伊",
                "receive_uid": 10580661,
                "receive_realname": "黄伊",
                "receive_time": "2024-07-17 14:12:19",
                "qc_uid": 0,
                "qc_realname": "",
                "qc_time": "",
                "wid": 1643,
                "supplier_id": 1766,
                "delivery_order_sn": "CR240717010",
                "delivery_item_id": "113900",
                "order_item_id": "114038",
                "sku": "serion-001",
                "product_name": "Serion",
                "fnsku": "",
                "seller_id": "114",
                "product_receive_num": 100,
                "qc_num": 0,
                "qc_bad_num": 0,
                "qc_rate_pass": "0%",
                "qc_rate": "0%",
                "product_good_num": 0,
                "product_bad_num": 0,
                "whb_code_good": "",
                "whb_code_bad": "",
                "qc_remark": ""
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询销售退货单列表
下一章节 
查询质检单详情