创建已完成的盘点单

支持创建已完成的盘点单，盘点类型默认为“SKU+仓位”盘点

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/InventoryCheck/addOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	盘点仓库id,对应领星系统的仓库id	是	[int]	42
is_display_check	是否明盘：0 否，1 是【默认值】	是	[int]	1
check_uid	盘点人id	是	[int]	10609364
remark	单据备注	否	[string]	gogogo
product_list	盘点明细	是	[array]	
product_list>>product_id	本地产品id	是	[int]	18418
product_list>>seller_id	店铺id，传空或者不传则默认0 ,对应查询亚马逊店铺列表接口对应字段【seller_id】	否	[string]	0
product_list>>fnsku	FNSKU，不传则默认空	否	[string]	xxx
product_list>>whb_code	仓位，传空或者不传则默认可用暂存	否	[string]	ts_valid
product_list>>actual_inventory	实盘库存	是	[int]	0
product_list>>remark	盘点明细备注	否	[string]	gogogo
请求示例
{
    "wid": 42,
    "is_display_check": 1,
    "check_uid": 10609364,
    "remark": "gogogo",
    "product_list": [
        {
            "product_id": 18418,
            "seller_id": "0",
            "fnsku": "xxx",
            "whb_code": "ts_valid",
            "actual_inventory": "",
            "remark": "gogogo"
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
request_id	请求链路id	是	[string]	E487053A-3868-093F-1DFE-021B0867E205
response_time	响应时间	是	[string]	2022-08-23 18:37:20
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>order_sn	盘点单号	是	[string]	IC220823001
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "5E058CF8-2297-CBF5-127C-571123BE4D69",
    "response_time": "2024-08-01 15:54:49",
    "data": {
        "order_sn": "IC240801002"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询盘点单详情
下一章节 
获取自定义入库类型