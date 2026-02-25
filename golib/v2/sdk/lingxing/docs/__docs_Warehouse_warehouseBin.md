查询本地仓位列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/warehouseBin	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库ID，字符串id，多个使用英文逗号分隔	否	[string]	5
id	仓位ID，字符串id，多个使用英文逗号分隔	否	[string]	13
status	仓位状态：
1 禁用
2 启用	否	[string]	2
type	仓位类型：
5 可用
6 次品	否	[string]	5
offset	分页偏移量，默认为0	否	[int]	0
limit	限制条数，默认20条	否	[int]	20
请求示例
{
    "wid": 5,
    "id": 13,
    "status": 2,
    "type": 5,
    "offset": 0,
    "limit": 20
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
request_id	请求链路id	是	[string]	502B9DD9-1BA0-03C5-6C61-D77C830440A6
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[array]	
data>>id	仓位id	是	[int]	125
data>>wid	仓库ID	是	[int]	125
data>>Ware_house_name	仓库名	是	[string]	1
data>>storage_bin	仓位	是	[int]	3
data>>whb_status	仓位状态	是	[string]	
data>>type	仓位类型	是	[string]	
data>>sku_fnsku	仓位商品关系	是	[array]	
data>>sku_fnsku>>SKU	sku	是	[string]	
data>>sku_fnsku>>FNSKU	fnsku	是	[string]	
total		是	[int]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "574FB0B8-FD6E-8D29-91FD-295B55CF5398",
    "response_time": "2024-07-31 14:41:25",
    "data": [
        {
            "id": 20,
            "wid": 5,
            "Ware_house_name": "自发货仓库呀",
            "storage_bin": "A-1-3",
            "status": 2,
            "type": 5,
            "sku_fnsku": [
                {
                    "sku": "hh",
                    "product_id": 10752,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "HH11-1"
                },
                {
                    "sku": "WYZTKSJTZ",
                    "product_id": 29375,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "吴彦祖同款手机套装"
                },
                {
                    "sku": "YF-Product-001",
                    "product_id": 29856,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-001"
                },
                {
                    "sku": "YF-Product-002",
                    "product_id": 29908,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-002"
                },
                {
                    "sku": "YF-Product-003",
                    "product_id": 29909,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-003"
                },
                {
                    "sku": "YF-Product-004",
                    "product_id": 29910,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-004"
                },
                {
                    "sku": "YF-Product-005",
                    "product_id": 29911,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-005"
                },
                {
                    "sku": "YF-Product-002_MUTI",
                    "product_id": 29912,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "YF-Product-002（组）"
                },
                {
                    "sku": "YF-Product-006",
                    "product_id": 29913,
                    "fnsku": "B08D3MPS2V",
                    "store_id": "119",
                    "seller_name": "uboss123456-US",
                    "product_name": "YF-Product-006"
                },
                {
                    "sku": "REDMI",
                    "product_id": 31064,
                    "fnsku": "",
                    "store_id": "0",
                    "seller_name": "",
                    "product_name": "红米耳机"
                }
            ]
        }
    ],
    "total": 2
}
复制
错误
复制成功
 上一章节
添加/修改仓库
下一章节 
添加仓位