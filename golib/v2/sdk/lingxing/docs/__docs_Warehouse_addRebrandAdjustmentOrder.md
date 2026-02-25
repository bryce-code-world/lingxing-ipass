创建已完成的换标调整单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAdjustment/addRebrandAdjustmentOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	系统仓库id	是	[int]	1
remark	单据备注	否	[string]	单据备注
bin_type	出库仓位方式：【默认1】
1 系统自定选择
2 指定出库仓位	否	[int]	2
product_list	调整的产品明细数据	是	[array]	
product_list>>product_id	产品id	是	[int]	1
product_list>>seller_id	原店铺id，默认为0 ,对应查询亚马逊店铺列表接口对应字段【seller_id】	否	[string]	1
product_list>>fnsku	原FNSKU，默认为空	否	[string]	XXX111
product_list>>to_seller_id	新店铺id，默认为0	否	[int]	2
product_list>>to_fnsku	新FNSKU，默认空	否	[string]	XXX222
product_list>>adjustment_valid_num	调整量	是	[int]	1
product_list>>product_remark	产品备注	否	[string]	产品备注
product_list>>out_available_bin	出库仓位列表，默认可用暂存【仅当bin_type = 2 生效】	否	[array]	
product_list>>out_available_bin>>whb_code	出库仓位	否	[string]	T-1
product_list>>out_available_bin>>num	出库数量，默认按ERP页面逻辑出库
【out_available_bin>>num 之和 等于 adjustment_valid_num】	否	[int]	1
product_list>>in_available_bin	入库仓位列表，默认可用暂存	否	[array]	
product_list>>in_available_bin>>whb_code	入库仓位	否	[string]	T-2
product_list>>in_available_bin>>num	入库数量，当填写入库仓位时，该字段必填
【in_available_bin>>num 之和 等于 adjustment_valid_num】	否	[int]	1
请求示例
{
    "wid": 1,
    "remark": "单据备注",
    "bin_type": 2,
    "product_list": [
        {
            "product_id": 1,
            "seller_id": "1",
            "fnsku": "XXX111",
            "to_seller_id": 2,
            "to_fnsku": "XXX222",
            "adjustment_valid_num": 1,
            "product_remark": "产品备注",
            "out_available_bin": [
                {
                    "whb_code": "T-1",
                    "num": 1
                }
            ],
            "in_available_bin": [
                {
                    "whb_code": "T-2",
                    "num": 1
                }
            ]
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
response_time	响应时间	是	[string]	2023-01-10 16:22:24
data	响应数据	是	[object]	
data>>order_sn	生成的调整单的单据号	是	[string]	AD230110022
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "327BA547-EFD3-B788-678A-1776BCDDBDFF",
    "response_time": "2023-01-10 16:22:24",
    "data": {
        "order_sn": "AD230110022"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
创建已完成的数量调整单
下一章节 
创建已完成的SKU调整单