查询售后工单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/returns/workOrder/list	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
date_type	时间类型：
create_time 创建时间
complete_time 完成时间	是	[string]	create_time
start_time	开始时间，闭区间，格式：Y-m-d H:i:s	是	[string]	2023-06-15 00:00:00
end_time	结束时间，闭区间，格式：Y-m-d H:i:s	是	[string]	2023-06-16 23:59:59
offset	分页偏移量	是	[int]	0
length	分页长度，上限500	是	[int]	20
请求示例
{
    "date_type": "create_time",
    "start_time": "2023-06-15 00:00:00",
    "end_time": "2023-06-16 23:59:59",
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	操作成功
request_id	请求链路id	是	[string]	eb2bdeed-4842-42b0-8b5e-81345dda5078.1679967282563
response_time	响应时间	是	[string]	2023-03-28 09:34:45
data	响应数据	是	[object]	
data>>total	总数	是	[int]	46
data>>list	列表信息	是	[array]	
data>>list>>rma_id	售后单号	是	[string]	202306050003
data>>list>>rma_type	售后类型：
1 退货退款
2 仅退货
3 仅退款
4 退货补发
5 补发	是	[int]	1
data>>list>>status	状态【售后单状态/订单状态】	是	[string]	处理中/待收货
data>>list>>order_number	系统单号	是	[string]	103290951812762112
data>>list>>platform_code	平台code	是	[string]	10002
data>>list>>country	国家	是	[string]	美国
data>>list>>store_name	店铺名称	是	[string]	test-walmart
data>>list>>rma_from	单据来源：1 手工创建，2 线上同步	是	[int]	2
data>>list>>rma_reason	售后原因	是	[string]	不给换
data>>list>>return_warehouse_code	退货仓库id	是	[string]	11
data>>list>>return_warehouse_name	退货仓库名称	是	[string]	仓库11
data>>list>>return_type_code	退货类型：1 买家退货，2 卖家退货	是	[int]	1
data>>list>>return_logistice_type_code	物流类型：1 自选物流，2 三方物流	是	[int]	1
data>>list>>return_logistics	物流商	是	[string]	物流商xx
data>>list>>tracking_number	跟踪号	是	[string]	123
data>>list>>estimate_arrive_date	预计到货时间	是	[string]	168658600000
data>>list>>order_amount	订单金额	是	[string]	338.0
data>>list>>returen_method	退款方式：0 无选择，1 手动退款	是	[int]	1
data>>list>>refund_status	退款状态：
待退款
退款中
退款失败
退款完成	是	[string]	退款完成
data>>list>>comment	备注	是	[string]	备注xx
data>>list>>create_time	创建时间	是	[string]	2023-06-09 11:23:44
data>>list>>update_time	更新时间	是	[string]	2023-06-09 11:25:01
data>>list>>complete_time	完成时间	是	[string]	2023-06-09 11:25:01
data>>list>>tag_list	标签信息	是	[array]	
data>>list>>tag_list>>tag_id	标签id	是	[string]	907225000920047630
data>>list>>tag_list>>tag_name	标签名称	是	[string]	RMA标签1
data>>list>>refund	退款信息	是	[array]	
data>>list>>refund>>platform_order_number	退款平台单号	是	[string]	ama-test0330-mnkg8a
data>>list>>refund>>refund_amount	退款金额	是	[number]	2.0
data>>list>>rma_info	售后信息	是	[array]	
data>>list>>rma_info>>sku	SKU	是	[string]	danpin2
data>>list>>rma_info>>msku	MSKU	是	[string]	sHOLDER001
data>>list>>rma_info>>product_name	品名	是	[string]	单品2
data>>list>>rma_info>>product_image	商品图片	是	[string]	https//xxxx.jpg
data>>list>>rma_info>>product_amount	商品订购数量	是	[int]	1
data>>list>>rma_info>>product_rma_amount	商品售后数量	是	[int]	1
data>>list>>rma_info>>product_comment	商品售后说明	是	[string]	暂无说明
data>>list>>rma_info>>product_rma_image	商品售后图片	是	[string]	https://xxxx-xxxxx.com/xxxxxxxxxx.jpg
data>>list>>rma_info>>platform_order_no	平台单号	是	[string]	
data>>list>>related_info	关联单据	是	[array]	
data>>list>>related_info>>return_order_id	销售退货单单据号	是	[string]	RT901305646307553280
data>>list>>related_info>>return_order_status_code	销售退货单状态：
-1 异常
1 待提交
2 待审批
3 待收货
4 已作废
5 已完成
6 导入中	是	[int]	3
data>>list>>related_info>>issue_order_number	补发订单系统单号	是	[string]	103295742586946048
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "895f2f0493414a4184d28cffaff87a09.1686797615263",
    "response_time": "2023-06-15 10:53:36",
    "data": {
        "total": 101,
        "list": [
            {
                "status": "处理中/待收货",
                "rma_id": 202306050003,
                "rma_type": 2,
                "order_number": 103280851812762112,
                "platform_code": "10008",
                "country": "美国",
                "store_name": "test-walmart",
                "rma_from": 2,
                "rma_reason": "不给换",
                "return_warehouse_code": 11,
                "return_warehouse_name": "仓库11",
                "return_type_code": 1,
                "return_logistice_type_code": 1,
                "return_logistics": "",
                "tracking_number": "",
                "estimate_arrive_date": 0,
                "order_amount": null,
                "refund_method": 1,
                "refund_status": "",
                "comment": "",        
                "create_time": "2023-06-09 11:23:44",
                "update_time": "2023-06-09 11:25:01",
                "complete_time": "2023-06-09 11:25:01",
                "tag_list": [
                    {
                        "tag_id": "907225000920047630",
                        "tag_name": "RMA标签1"
                    }
                ],
                "refund": [],
                "rma_info": [
                    {
                        "sku": "SKUPIC",
                        "msku": "walm_test2_sku6kGpZ",
                        "product_name": "pic",
                        "product_image": "https://xxx/47889efe3d00df67cb39.jpg",
                        "product_amount": 5,
                        "product_rma_amount": 1,
                        "product_comment": "",
                        "product_rma_image": null,
                        "platform_order_no": "NINE00002"
                    }
                ],
                "related_info": [
                    {
                        "return_order_id": "RT901280870700859904",
                        "return_order_status_code": 5,
                        "issue_order_number": null
                    },
                    {
                        "return_order_id": "RT901294602998157824",
                        "return_order_status_code": 3,
                        "issue_order_number": null
                    },
                    {
                        "return_order_id": "RT901294603713667072",
                        "return_order_status_code": 3,
                        "issue_order_number": null
                    },
                    {
                        "return_order_id": "RT901320516537774080",
                        "return_order_status_code": 3,
                        "issue_order_number": null
                    }
                ]
            }
        ]
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "response_time": "2023-06-15 10:30:23",
    "message": "startTime时间格式不符合 yyyy-MM-dd HH:mm:ss",
    "request_id": "bbf9d362-1b69-4703-bbb4-16dc2465f553.1688351423595",
    "data": null
}
复制
错误
复制成功
 上一章节
查询客户列表（旧）
下一章节 
查询店铺绩效列表