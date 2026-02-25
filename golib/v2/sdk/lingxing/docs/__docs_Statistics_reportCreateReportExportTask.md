报告导出 - 创建导出任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/report/create/reportExportTask	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
seller_id	亚马逊店铺id，查询亚马逊店铺列表接口对应字段【seller_id】	是	[string]	A1MQMW3JWPNCBX
report_type	亚马逊报表类型【具体类型参看下方附加说明】	是	[string]	Amazon VAT Transactions Report
data_start_time	亚马逊报表请求开始时间，时间格式：YYYY-MM-DDTHH:MM:SSZ	否	[string]	2022-06-21T18:17:13+00:00
data_end_time	亚马逊报表请求结束时间，时间格式：YYYY-MM-DDTHH:MM:SSZ	否	[string]	2022-06-25T18:17:16+00:00
marketplace_ids	亚马逊市场id	是	[array]	["ATVPDHSKDCJER"]
region	店铺所在的地区【对应区域值支持国家见附加说明】：
na 北美
eu 欧洲
fe 远东	是	[string]	na
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>task_id	任务id	是	[string]	f5345297-07e2-4b08-becf-a4c29335246b
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "9aef4b474dff4f02bcfcd353a02b255c.1696764983947",
    "response_time": "2023-10-08 19:36:25",
    "data": {
        "task_id": "f5345297-07e2-4b08-becf-a4c29335246b"
    },
    "total": 0
}
复制
错误
复制成功
附加说明

report_type说明如下：


请查阅此链接查看可使用的report type https://developer-docs.amazon.com/sp-api/docs/report-type-values

以下枚举仅供参考，实际以亚马逊提供的文档链接中为准

GET_AMAZON_FULFILLED_SHIPMENTS_DATA_GENERAL 【亚马逊配送货件】
GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_GENERAL 【所有订单】
GET_FLAT_FILE_SALES_TAX_DATA 【销售税报告】
SC_VAT_TAX_REPORT 【VAT计算报告】
GET_VAT_TRANSACTION_DATA 【VAT交易报告】
GET_RESERVED_INVENTORY_DATA 【预留库存】
GET_AFN_INVENTORY_DATA 【亚马逊库存】
GET_FBA_STORAGE_FEE_CHARGES_DATA 【月度仓储费】
GET_STRANDED_INVENTORY_UI_DATA 【无在售商品的亚马逊库存报告】
GET_RESERVED_INVENTORY_DATA 【预留库存报告】
GET_FBA_MYI_UNSUPPRESSED_INVENTORY_DATA 【管理亚马逊物流库存报告】
GET_FBA_MYI_ALL_INVENTORY_DATA 【管理亚马逊物流库存报告-已存档】
GET_AFN_INVENTORY_DATA_BY_COUNTRY 【多国库存报告】
GET_FBA_INVENTORY_PLANNING_DATA 【库存状况报告】
GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA 【亚马逊物流买家退货】
GET_FBA_FULFILLMENT_CUSTOMER_SHIPMENT_REPLACEMENT_DATA 【换货】
GET_FBA_REIMBURSEMENTS_DATA 【赔偿】
GET_FBA_FULFILLMENT_LONGTERM_STORAGE_FEE_CHARGES_DATA 【长期仓储费】
GET_REFERRAL_FEE_PREVIEW_REPORT （已废弃）亚马逊官方已改为 GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA【费用预览报告】
GET_FBA_OVERAGE_FEE_CHARGES_DATA 【超量仓储费报告】
GET_FBA_RECOMMENDED_REMOVAL_DATA 【建议移除】
GET_FBA_FULFILLMENT_REMOVAL_ORDER_DETAIL_DATA 【移除订单详情】
GET_FBA_FULFILLMENT_REMOVAL_SHIPMENT_DETAIL_DATA 【移除货件详情】


region 说明如下：
na【North America】：包括国家为 CA、US、MX、BR

eu【Europe】：包括国家为 ES、UK、FR、BE、NL、DE、IT、SE、ZA、PL、EG、TR、SA、AE、IN

fe【Far East】：包括国家为 SG、AU、JP
 上一章节
查询亚马逊源报表- 盘存记录
下一章节 
报告导出-查询导出任务结果