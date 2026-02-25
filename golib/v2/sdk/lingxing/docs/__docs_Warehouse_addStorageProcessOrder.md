创建加工单 / 拆分单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageProcess/addStorageProcessOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
type	单据类型：1 加工单，2 拆分单	是	[int]	1
wid	系统仓库id	是	[int]	1
remark	备注	否	[string]	xxx
product_list	产品信息	是	[array]	
product_list>>combo_sku	组合品sku	是	[string]	
product_list>>combo_sid	组合品店铺id，没有传0即可	是	[int]	
product_list>>combo_fnsku	组合品fnsku	否	[string]	
product_list>>quantity_num	加工量 / 拆分量	是	[int]	
product_list>>combo_whb_code	加工单组合品入库仓位，不传默认暂存可用；拆分单不用传
【查询本地仓位列表 接口对应字段【storage_bin】】	否	[string]	
product_list>>process_fee	加工费【默认0】；加工单专有，拆分单可不传	否	[number]	
product_list>>single_product_list	单品明细项	是	[array]	
product_list>>single_product_list>>sku	单品sku	是	[string]	
product_list>>single_product_list>>fnsku	单品fnsku	否	[string]	
product_list>>single_product_list>>sid	单品店铺id，组合品店铺id为0时此字段值传0	是	[int]	
product_list>>single_product_list>>price_scale	单品拆分比例	是	[number]	
product_list>>single_product_list>>whb_code	拆分单的单品入库仓位，不传默认可用暂存；加工单不用传
【查询本地仓位列表 接口对应字段【storage_bin】】	否	[string]	
product_list>>single_product_list>>remark	remark	否	[string]	xx
请求示例
{
    "type": 1,
    "wid": 1,
    "remark": "备注",
    "product_list": [
        {
            "combo_sku": "xxxxxx",
            "combo_sid": 0,
            "combo_fnsku": "xxx",
            "quantity_num": 0,
            "combo_whb_code": "1",
            "process_fee": "1",
            "single_product_list": [
                {
                    "sku": "xxx",
                    "fnsku": "",
                    "sid": 0,
                    "price_scale": "",
                    "whb_code": "1",
                    "remark": ""
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
request_id	请求链路id	是	[string]	D4F7833C-53E7-207E-E4AB-E0BA96D0DA13
response_time	响应时间	是	[string]	2022-08-23 16:56:31
data	响应数据	是	[object]	
data>>order_sn	加工单/拆分单 单号	是	[string]	PT230321003
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D4F7833C-53E7-207E-E4AB-E0BA96D0DA13",
    "response_time": "2022-08-23 16:56:31",
    "data": {
        "order_sn": "PT230321003"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
撤销出库单
下一章节 
加工单列表