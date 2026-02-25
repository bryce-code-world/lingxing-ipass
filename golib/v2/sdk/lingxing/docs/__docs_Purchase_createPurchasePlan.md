创建待采购的采购计划

支持创建采购计划，状态为“待采购”

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/createPurchasePlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
remark	计划备注	否	[string]	备注
data	产品信息	是	[array]	
data>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[string]	1
data>>supplier_id	供应商id	否	[int]	1
data>>sku	sku	是	[string]	SKU001
data>>fnsku	fnsku	否	[string]	FNSKU001
data>>wid	仓库id	否	[int]	1
data>>purchaser_id	采购方id	否	[int]	1
data>>expect_arrive_time	期望到货时间，格式：Y-m-d	否	[string]	2021-10-01
data>>cg_uid	采购员id	否	[int]	12
data>>quantity_plan	计划采购量	是	[int]	100
data>>remark	产品备注	否	[string]	xxxx
data>>options	可选参数	否	[object]	
data>>options>>is_auto_fill_fnsku	是否自动FNSKU：【默认0】 0 否，1 是	否	[int]	1
data>>options>>is_auto_fill_store	是否自动填充店铺：【默认0】 0 否，1 是	否	[int]	1
请求示例
{
    "remark": "备注",
    "data": [{
        "sid": "1",
        "supplier_id": 1,
        "sku": "SKU001",
        "fnsku": "FNSKU001",
        "wid": 1,
        "purchaser_id": 1,
        "expect_arrive_time": "2021-10-01",
        "cg_uid": 0,
        "quantity_plan": 100,
        "remark": "",
        "options": {
            "is_auto_fill_fnsku": 1,
            "is_auto_fill_store": 1
        }
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8DFEE1CA-EC9B-F401-5D41-9F95251D5D50
response_time	响应时间	是	[string]	2021-09-21 15:48:58
data	响应数据	是	[object]	
data>>plan_sn	计划编号	是	[array]	["PP211112002"]
data>>ppg_sn	计划批次号	是	[string]	PPG211112001
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D41373ED-AF54-CDF1-1658-B8F86286F209",
    "response_time": "2021-11-12 12:37:55",
    "data": {
        "ppg_sn": "PPG211112001",
        "plan_sn": ["PP211112002"]
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "仓库参数错误 [请求码:3CF77C]",
    "request_id": "31667FCC-24E5-0CBE-2FC8-23E6908D9473",
    "response_time": "2021-11-12 12:34:26",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询采购计划列表
下一章节 
查询采购单列表