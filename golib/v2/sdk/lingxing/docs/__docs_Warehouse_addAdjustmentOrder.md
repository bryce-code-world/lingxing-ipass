创建已完成的数量调整单

支持推送已完成的数量调整单至领星ERP，并调整对应数量

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAdjustment/addAdjustmentOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	系统仓库id	是	[int]	
remark	单据备注	否	[string]	
product_list	调整的产品明细数据	是	[array]	
product_list>>adjustment_valid_num	可用调整数量，不调整则传0，但不能与次品调整数量同时为0	是	[int]	
product_list>>adjustment_bad_num	次品调整数量，不调整则传0，但不能与可用调整数量同时为0	是	[int]	
product_list>>adjustment_available_bin	可用仓位编号，为空则默认可用暂存	是	[string]	
product_list>>adjustment_inferior_bin	次品仓位编号，为空则默认为次品暂存	是	[string]	
product_list>>adjustment_valid_sgn	可用增加标志符号，增加时+ ，减少是-	是	[string]	
product_list>>adjustment_bad_sgn	次品增加标志符号，增加时+ ，减少是-	是	[string]	
product_list>>sku	sku	是	[string]	
product_list>>fnsku	fnsku，没有可以为空	是	[string]	
product_list>>product_id	产品id	是	[int]	
product_list>>seller_id	店铺id，默认为0 ,对应查询亚马逊店铺列表接口对应字段【seller_id】	是	[int]	 
请求示例
{
    "wid": 1,
    "remark": "",
    "product_list": [
        {
            "adjustment_valid_num": 1,
            "adjustment_bad_num": 2,
            "adjustment_available_bin": "可用暂存",
            "adjustment_inferior_bin": "次品暂存",
            "adjustment_valid_sgn": "+",
            "adjustment_bad_sgn": "+",
            "sku": "SKU1",
            "fnsku": "FNSKU1",
            "product_id": 1,
            "seller_id": "1"
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
request_id	请求链路id	是	[string]	327BA547-EFD3-B788-678A-1776BCDDBDFF
response_time	响应时间	是	[string]	2023-01-10 17:21:54
data	响应数据	是	[object]	
data>>order_sn	生成的调整单的单据号	是	[string]	AD230110025
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "327BA547-EFD3-B788-678A-1776BCDDBDFF",
    "response_time": "2023-01-10 17:21:54",
    "data": {
        "order_sn": "AD230110025"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询调整单列表
下一章节 
创建已完成的换标调整单