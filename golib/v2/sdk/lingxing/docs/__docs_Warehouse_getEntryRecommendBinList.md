查询产品仓位列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/warehouseConfig/warehouseBin/getEntryRecommendBinList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
list		是	[object]	
list>>wid	仓库id	是	[string]	
list>>productId	产品id	是	[string]	
list>>fnsku	fnsku	否	[string]	
list>>sid	店铺id	否	[string]	
withHistory	是否查询历史仓位，false-否true-是;默认否	否	[boolean]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/warehouseConfig/warehouseBin/getEntryRecommendBinList?access_token=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "list": [
    {
        "wid": "1", 
        "productId": 39258, 
        "fnsku": "", 
        "sid": 103
    },
    {
        "wid": "1", 
        "productId": 31386, 
        "fnsku": "", 
        "sid": 103
    }
        ], 
    "withHistory": false
 }'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0成功	否	[string]	
message	提示信息	是	[string]	
errorDetails	错误信息	是	[string]	
requestId	请求链路id	是	[string]	
data		是	[object]	
data>>fnsku	fnsku	是	[string]	
data>>productId	产品id	是	[string]	
data>>wid	仓库id	是	[string]	
data>>sid	店铺id	是	[string]	
data>>validStorageArr	返回可用仓位列表	是	[array]	
data>>validStorageArr>>productNum	可用量	是	[string]	
data>>validStorageArr>>whbId	仓位id	是	[string]	
data>>validStorageArr>>whbType	仓位类型	是	[string]	
data>>validStorageArr>>whbTypeName	仓位	是	[string]	
data>>validStorageArr>>whbCode	仓位编码	是	[string]	
data>>validStorageArr>>whbCodeName	仓位名称	是	[string]	
data>>validStorageArr>>name	仓位编码	是	[string]	
data>>validStorageArr>>value	仓位名称	是	[string]	
data>>invalidStorageArr	返回次品仓位列表	是	[array]	
data>>invalidStorageArr>>productNum	可用量	是	[string]	
data>>invalidStorageArr>>whbId	仓位id	是	[string]	
data>>invalidStorageArr>>whbType	仓位类型	是	[string]	
data>>invalidStorageArr>>whbTypeName	仓位	是	[string]	
data>>invalidStorageArr>>whbCode	仓位编码	是	[string]	
data>>invalidStorageArr>>whbCodeName	仓位名称	是	[string]	
data>>invalidStorageArr>> name	仓位编码	是	[string]	
data>>invalidStorageArr>>value	仓位名称	是	[string]	
total	总数	是	[string]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "5545f7272e03483caeff02cd4cb9c592.1757574480275",
    "response_time": "2025-09-11 15:08:07",
    "data": [
        {
            "fnsku": "",
            "product_id": 39258,
            "wid": 1,
            "sid": "103",
            "valid_storage_arr": [
                {
                    "product_num": 0,
                    "whb_id": 2,
                    "whb_type": 2,
                    "whb_type_name": "可用暂存",
                    "whb_code": "ts_valid",
                    "whb_code_name": "可用暂存",
                    "name": "ts_valid",
                    "value": "可用暂存"
                }
            ],
            "invalid_storage_arr": []
        },
        {
            "fnsku": "",
            "product_id": 31386,
            "wid": 1,
            "sid": "103",
            "valid_storage_arr": [
                {
                    "product_num": 69,
                    "whb_id": 2,
                    "whb_type": 2,
                    "whb_type_name": "可用暂存",
                    "whb_code": "ts_valid",
                    "whb_code_name": "可用暂存",
                    "name": "ts_valid",
                    "value": "可用暂存"
                }
            ],
            "invalid_storage_arr": [
                {
                    "product_num": 12,
                    "whb_id": 3,
                    "whb_type": 3,
                    "whb_type_name": "次品暂存",
                    "whb_code": "ts_bad",
                    "whb_code_name": "次品暂存",
                    "name": "ts_bad",
                    "value": "次品暂存"
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
海外sku取消配对
下一章节 
查询FBA库存列表