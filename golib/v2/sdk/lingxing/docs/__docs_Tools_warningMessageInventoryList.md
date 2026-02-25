查询预警消息列表-库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/settings/warningMessage/inventoryList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	否	[int]	0
length	分页长度，默认50，上限200	否	[int]	50
model_id_list	预警模型：
4 本地库存预警
5 亚马逊库存预警
22 本地库龄预警
23 亚马逊库龄预警	否	[array]	[5,22]
product_type_list	产品类型：
2 MSKU
3 SKU+仓库+店铺+FNSKU	否	[array]	[2,3]
start_date	开始日期【提醒时间】，闭区间，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2024-06-05
end_date	结束日期【提醒时间】，闭区间，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2024-08-05
show_status	处理状态：
0 待处理
1 全部	是	[int]	1
请求示例
{
    "offset": 0,
    "length": 50,
    "mode_id_list": [5,22],
    "product_type_list": [2,3],
    "start_date": "2024-06-05",
    "end_date": "2024-08-05",
    "show_status": 1
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
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>message_id	用户消息id	是	[string]	906377395097195176
data>>model_id	预警模型id	是	[int]	4
data>>model_name	预警模型名称	是	[string]	本地库存预警
data>>product_type	产品类型：
2 MSKU
3 SKU+仓库+店铺+FNSKU	是	[int]	2
data>>product_type_str	产品类型说明	是	[string]	SKU+仓库+店铺+FNSKU
data>>monitor_type_str	监控指标	是	[string]	可用库存+待检量
data>>notify_way_str	提醒方式	是	[string]	站内消息
data>>notify_time	提醒时间	是	[string]	2023-11-13 09:45:29
data>>receive_uid	接收人id	是	[string]	1
data>>receiver	接收人名称	是	[string]	超级管理员
data>>handle_status	处理状态：
0 待处理
1 已处理	是	[int]	0
data>>handle_status_str	处理状态说明	是	[string]	待处理
data>>read_status	阅读状态：
0 未读
1 已读	是	[int]	0
data>>read_status_str	阅读状态说明	是	[string]	未读
data>>monitor_time	预警时间	是	[string]	2023-11-13
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "e0f84c1034024dcea0f031f17055bfd5.1699862034631",
    "response_time": "2023-11-13 15:53:54",
    "data": [
        {
            "message_id": "906377395097195176",
            "model_id": 4,
            "model_name": "本地库存预警",
            "product_type": 3,
            "product_type_str": "SKU+仓库+店铺+FNSKU",
            "monitor_type_str": "可用库存+待检量",
            "notify_way_str": "站内消息",
            "notify_time": "2023-11-13 09:45:29",
            "receive_uid": "1",
            "receiver": "0超级管理员01",
            "handle_status": 0,
            "handle_status_str": "待处理",
            "read_status": 0,
            "read_status_str": "未读",
            "monitor_time": "2023-11-13"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询预警消息列表-商品
下一章节 
查询销量、订单量、销售额