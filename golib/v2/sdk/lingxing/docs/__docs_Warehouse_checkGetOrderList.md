查询盘点单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/InventoryCheck/getOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	盘点仓库id，多个使用英文逗号分隔	否	[string]	42
check_type	盘点类型，多个盘点类型用英文逗号分隔：
1 整仓盘点
2 SKU盘点
3 仓位盘点
4 SKU+仓位盘点	否	[string]	1,2
date_field	搜索时间类型：
create_date 创建时间【默认值】
check_date 盘点时间	否	[string]	create_date
start_date	开始日期，格式：Y-m-d	否	[string]	2024-07-20
end_date	结束日期，格式：Y-m-d	否	[string]	2024-07-31
search_field	搜索字段：
order_sn 盘点单号
create_user 创建人
check_user 盘点人
remark 备注	否	[string]	order_sn
search_value	搜索值	否	[string]	IC240726001
status	盘点状态：
10 待盘点
20 预锁
30 盘点中
40 已盘点
121 待审核
122 已驳回
123 通过
124 作废	否	[int]	
page	分页页码，默认1	否	[int]	1
page_size	分页长度，默认20	否	[int]	20
请求示例
{
    "wid": 42,
    "check_type": "1,2",
    "date_field": "create_date",
    "start_date": "2024-07-20",
    "end_date": "2024-07-31",
    "search_field": "order_sn",
    "search_value": "IC240726001",
    "page": 1,
    "page_size": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	信息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	D17B0688-C52D-1BA9-1BCC-50AA00058A3F
response_time	响应时间	是	[string]	2022-08-23 16:56:31
total	总数	是	[int]	231
data	响应数据	是	[object]	
data>>order_sn	盘点单号	是	[string]	IC220810001
data>>status	盘点状态：
10 待盘点
20 预锁
30 盘点中
40 已盘点
121 待审核
122 已驳回
123 通过
124 作废	是	[int]	40
data>>status_text	状态文本	是	[string]	已盘点
data>>wid	盘点仓库id	是	[int]	1
data>>ware_house_name	盘点仓库名称	是	[string]	仓库1
data>>check_type	盘点类型：
1 整仓盘点
2 SKU盘点
3 仓位盘点
4 SKU+仓位盘点	是	[int]	2
data>>check_type_text	盘点类型文本	是	[string]	SKU盘点
data>>is_display_check	是否明盘：0 否，1 是	是	[int]	1
data>>display_check_name	是否明盘文本	是	[string]	明盘
data>>is_zero	是否零库存参与盘点：0 否，1 是	是	[int]	1
data>>product_type	产品种类	是	[int]	1
data>>create_uid	创建人id	是	[int]	10002789
data>>create_user	创建人姓名	是	[string]	张三
data>>create_time	创建时间	是	[string]	2022-08-10 18:02:37
data>>check_uid	盘点人id	是	[int]	10002789
data>>check_user	盘点人姓名	是	[string]	张三
data>>real_check_uid	实际盘点人id	是	[int]	10002789
data>>real_check_user	实际盘点人姓名	是	[string]	张三
data>>check_time	盘点时间	是	[string]	2022-08-10 18:02:37
data>>commit_uid	提交人id	是	[int]	10002789
data>>commit_user	提交人姓名	是	[string]	张三
data>>commit_time	提交时间	是	[string]	2022-08-10 18:03:09
data>>cancel_uid	作废人id	是	[int]	
data>>cancel_user	作废人姓名	是	[string]	
data>>cancel_time	作废时间	是	[string]	
data>>cancel_reason	作废原因	是	[string]	
data>>remark	备注	是	[string]	
data>>request_status	单据状态：0 正常，1 处理中	是	[string]	
data>>file	上传附件信息	是	[array]	
data>>file>>file_id	附件id	否	[int]	
data>>file>>file_name	附件名称	否	[string]	
data>>file>>file_url	附件URL	否	[string]	 
 上一章节
创建已完成的SKU调整单
下一章节 
查询盘点单详情