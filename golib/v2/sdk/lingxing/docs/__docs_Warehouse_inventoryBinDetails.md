查询仓位库存明细
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/inventoryBinDetails	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库id，多个仓库用英文逗号分隔，默认所有仓库	否	[string]	1,5,18
bin_type_list	仓位类型，多个类型用英文逗号分隔：
1 待检暂存
2 可用暂存
3 次品暂存
4 拣货暂存
5 可用
6 次品	否	[string]	5
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20 ，上限500	否	[int]	20
请求示例
{
    "wid": "1,5,18",
    "bin_type_list": "5",
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
message	返回消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BDE02505-821E-2757-5938-650A21A6875B
response_time	响应时间	是	[string]	2021-07-16 11:01:34
data	响应数据	是	[array]	
data>>wid	仓库id	是	[int]	1771
data>>wh_name	仓库名称	是	[string]	
data>>whb_id	仓位id	是	[int]	
data>>whb_name	仓位名称	是	[string]	
data>>whb_type	仓位类型	是	[string]	
data>>whb_type_name	仓位类型名称	是	[string]	
data>>product_id	商品id	是	[int]	23551
data>>sku	SKU	是	[string]	导入大写PNG
data>>seller_id	店铺id	是	[string]	
data>>fnsku	FNSKU	是	[string]	
data>>total	总量	是	[int]	11
data>>lockNum	锁定量	是	[int]	
data>>validNum	未锁定量	是	[string]	
data>>third_inventory	海外仓第三方库存信息	是	[object]	
data>>third_inventory>>qty_sellable	可用量	是	[int]	35194
data>>third_inventory>>qty_pending	待上架库存	是	[int]	
data>>third_inventory>>qty_reserved	锁定量	是	[int]	7
data>>third_inventory>>qty_onway	备货在途	是	[int]	
total	总条数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "A884D2C7-186A-2810-072A-7A0123093BBA",
    "response_time": "2024-07-31 15:08:33",
    "data": [
        {
            "wid": 1,
            "product_id": 10001,
            "sku": "ceshi001",
            "seller_id": "0",
            "fnsku": "",
            "product_total": 55889,
            "product_valid_num": 0,
            "product_bad_num": 0,
            "product_qc_num": 3322,
            "product_lock_num": 52567,
            "good_lock_num": 52566,
            "bad_lock_num": 1,
            "stock_cost_total": "4214317.76",
            "quantity_receive": "10005050",
            "stock_cost": "75.4051",
            "product_onway": 187,
            "transit_head_cost": "0.00",
            "average_age": 610,
            "third_inventory": [],
            "stock_age_list": [
                {
                    "name": "0-15天库龄",
                    "qty": 0
                },
                {
                    "name": "16-45天库龄",
                    "qty": 119
                },
                {
                    "name": "61-120天库龄",
                    "qty": 14
                },
                {
                    "name": "121-150天库龄",
                    "qty": 1128
                },
                {
                    "name": "151-180天库龄",
                    "qty": 1
                },
                {
                    "name": "181-210天库龄",
                    "qty": 131
                },
                {
                    "name": "366天以上库龄",
                    "qty": 54111
                }
            ],
            "available_inventory_box_qty": "5256.6"
        }
    ],
    "total": 12093
}
复制
错误
复制成功
 上一章节
查询仓库库存明细
下一章节 
查询批次明细