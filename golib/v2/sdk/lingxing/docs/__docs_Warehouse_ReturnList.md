查询销售退货单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/returns/v2/list	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
offset	页码，非偏移量
offset传1，则返回第一页数据	是	[int]	0
length	每页记录数	是	[int]	20
time_type	搜索时间类型：updateTime 更新时间【不传默认为创建时间】	否	[string]	updateTime
start_time	开始时间，格式：Y-m-d H:i:s	是	[string]	2023-01-01 00:00:00
end_time	结束时间，格式：Y-m-d H:i:s	是	[string]	2023-01-01 00:00:00
platform_code	平台code	否	[array]	[10007,10008]
sales_type	退货类型：1 买家退货，2 物流商退货	否	[int]	1
status	订单状态：
-1 异常
1 待提交
2 待审批
3 待收货
4 已作废
5 已完成
6 导入中	否	[array]	[1,4,5]
store_id	店铺id	否	[array]	["127190049090011648","127190049090011654"]
wid	系统仓库id	否	[array]	[1,578,765]
请求示例
{
    "offset": 0,
    "length": 20,
    "time_type": "updateTime",
    "start_time": "2023-01-01 00:00:00",
    "end_time": "2024-07-29 20:50:00",
    "sales_type": 1,
    "status": [1,4,5],
    "store_id": ["127190049090011648","127190049090011654"],
    "wid": [1,578,765]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	操作成功	是	[string]	操作成功
request_id	请求链路id	是	[string]	052e20647c9c4622adb7576ae335211d.1687946868074
response_time	响应时间	是	[string]	2023-06-28 18:07:50
data	响应数据	是	[object]	
data>>total	总条数	是	[int]	41
data>>list	详细列表	是	[array]	
data>>list>>complete_time	完成时间	是	[string]	2023-06-05 16:25:08
data>>list>>global_order_no	系统单号	是	[long]	103280851812762112
data>>list>>gmt_create	创建时间	是	[string]	2023-06-05 16:25:08
data>>list>>gmt_modified	更新时间	是	[string]	2023-06-05 16:25:08
data>>list>>has_prediction	是否三方仓预报 0 否 1 是	是	[int]	0
data>>list>>logistics_provider_name	物流商	是	[string]	xx物流
data>>list>>platform	平台	是	[string]	walmart
data>>list>>pre_arrival_time	预计到货时间	是	[string]	2023-06-05 16:25:08
data>>list>>reason	退货原因	是	[string]	不给换
data>>list>>remark	备注	是	[string]	xxxx备注
data>>list>>rma_order_no	退货单号	是	[string]	RT901320516537774080
data>>list>>sales_type	退货类型	是	[string]	买家退货
data>>list>>site	站点	是	[string]	美国
data>>list>>status	订单状态	是	[string]	待收货
data>>list>>store_id	多平台店铺id	是	[int]	110000000017008003
data>>list>>sid	亚马逊店铺id	是	[int]	17
data>>list>>store_name	店铺名称	是	[string]	xxx店铺
data>>list>>tracking_no	跟踪号	是	[string]	212
data>>list>>uid_name	创建人名称	是	[string]	张三
data>>list>>sys_wid	退货仓库id	是	[int]	11
data>>list>>w_name	退货仓库名称	是	[string]	仓库11
data>>list>>relation_order_info	关联单据	是	[object]	
data>>list>>relation_order_info>>global_order_no	系统单号	是	[int]	103296739036561920
data>>list>>relation_order_info>>storage_in_order_no	入库单号	是	[string]	IB230520001
data>>list>>items	商品信息	是	[array]	
data>>list>>items>>available_quantity	可用量	是	[int]	1
data>>list>>items>>defective_quantity	次品量	是	[int]	1
data>>list>>items>>destroyed_quantity	销毁量	是	[int]	1
data>>list>>items>>quantity	数量	是	[int]	1
data>>list>>items>>receiving_quantity	收货数量	是	[int]	1
data>>list>>items>>return_quantity	退货数量	是	[int]	1
data>>list>>items>>handler_type	处理方式：
默认不处理
重新上架
由WMS决定
暂存区
次品
销毁	是	[string]	默认不处理
data>>list>>items>>pic_infos	售后图片信息	是	[array]	
data>>list>>items>>remark	售后说明	是	[string]	实例xxxx
data>>list>>items>>pic_infos>>access_url	访问地址	是	[string]	https://xxx/xxx
data>>list>>items>>pic_infos>>file_name	图片名	是	[string]	file
data>>list>>items>>platform_order_no	平台单号	是	[string]	ama-test0330-ZEjHkL
data>>list>>items>>msku	msku	是	[string]	HOLDER001
data>>list>>items>>sku	本地产品sku	是	[string]	danpin2
data>>list>>items>>product_name	品名	是	[string]	单品2
data>>list>>items>>id	id	是	[long]	id
data>>list>>items>>defective_whb_code	次品仓位编码	是	[string]	
data>>list>>items>>available_whb_code	可用仓位编码	是	[string]
返回成功示例
{
    "code": 0,
    "data": {
        "total": "10",
        "list": [
            {
                "reason": "产品质量问题",
                "status": "已完成",
                "platform": "amazon",
                "site": null,
                "remark": "",
                "sid": "101",
                "items": [
                    {
                        "sku": "mini-blue2",
                        "msku": "",
                        "quantity": 3,
                        "remark": "",
                        "stockDeductId": "0",
                        "platform_order_no": "1258762477858",
                        "product_name": "干花*2",
                        "return_quantity": 3,
                        "receiving_quantity": 3,
                        "handler_type": "默认不处理",
                        "available_quantity": 3,
                        "defective_quantity": null,
                        "destroyed_quantity": 0,
                        "stock_cost": null,
                        "pic_infos": null
                    }
                ],
                "rma_order_no": "RT901406878224720384",
                "sales_type": "买家退货",
                "sys_wid": "1",
                "w_name": "wcn测试仓库测试改名",
                "delivery_w_name": null,
                "delivery_w_nid": null,
                "global_order_no": "103345679826100736",
                "logistics_provider_name": "",
                "tracking_no": "132",
                "uid_name": "陈旭鹏",
                "gmt_create": "2024-02-04 17:11:44",
                "store_name": null,
                "has_prediction": 0,
                "pre_arrival_time": "1970-01-01 08:00:00",
                "complete_time": "2024-02-04 17:12:19",
                "gmt_modified": "2024-02-04 17:12:19",
                "store_id": "127190049090011648",
                "relation_order_info": {
                    "global_order_no": "103345679826100736",
                    "storage_in_order_no": "IB240204014"
                }
            },
            {
                "reason": "包裹损坏",
                "status": "已完成",
                "platform": "amazon",
                "site": null,
                "remark": "",
                "sid": "101",
                "items": [
                    {
                        "sku": "BAN",
                        "msku": "",
                        "quantity": 2,
                        "remark": "",
                        "stockDeductId": "0",
                        "platform_order_no": "test1203-110",
                        "product_name": "BAN",
                        "return_quantity": 1,
                        "receiving_quantity": 1,
                        "handler_type": "默认不处理",
                        "available_quantity": 0,
                        "defective_quantity": null,
                        "destroyed_quantity": 0,
                        "stock_cost": null,
                        "pic_infos": null
                    }
                ],
                "rma_order_no": "RT901261790905177088",
                "sales_type": "买家退货",
                "sys_wid": "1",
                "w_name": "wcn测试仓库测试改名",
                "delivery_w_name": null,
                "delivery_w_nid": null,
                "global_order_no": "103255373488840704",
                "logistics_provider_name": "",
                "tracking_no": "",
                "uid_name": "金巧玲",
                "gmt_create": "2022-12-21 17:49:55",
                "store_name": null,
                "has_prediction": 0,
                "pre_arrival_time": "1970-01-01 08:00:00",
                "complete_time": "2023-11-13 15:01:23",
                "gmt_modified": "2023-11-13 15:01:23",
                "store_id": "127190049090011648",
                "relation_order_info": {
                    "global_order_no": "103255373488840704",
                    "storage_in_order_no": "IB231113004"
                }
            }
        ]
    },
    "response_time": "2024-07-31 16:00:40",
    "message": "操作成功",
    "request_id": "aa4794a7b2c0437483966bed3e7bee75.114.17224128397332303"
}
复制
错误
复制成功
 上一章节
收货单快捷入库
下一章节 
查询质检单列表