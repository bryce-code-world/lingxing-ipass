修改 FBM库存&处理时间
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/FbmManagement/modifyFbmInventory	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
fbmInventoryList	修改库存列表（支持批量修改，单次最多传200个元素）	是	[array]	
fbmInventoryList>>storeId	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	6
fbmInventoryList>>msku	要修改FBM库存或处理时间的MSKU	是	[string]	
fbmInventoryList>>fbmInventory	FBM库存（此项必填）	是	[int]	4213
fbmInventoryList>>shipDays	处理时间	否	[string]	"1"
请求示例
{
  "fbmInventoryList": [
    {
      "storeId": 6,
      "msku": "EXAMPLE-MSKU-123",
      "fbmInventory": 4213,
      "shipDays": "1"
    },
    {
      "storeId": 6,
      "msku": "EXAMPLE-MSKU-456",
      "fbmInventory": 200,
      "shipDays": "2"
    }
  ]
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	操作成功
error_details	错误信息	是	[array]	[ ]
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	2023-02-16 10:33:40
total	总数(默认为0)	是	[int]	0
data	响应数据	是	[object]	
data>>successNum	成功数量	是	[int]	1
data>>failureNum	失败数量	是	[int]	
data>>failureDetail	错误信息详情	是	[array]	
data>>failureDetail>>storeId	店铺id	是	[int]	9143
data>>failureDetail>>asin	asin	是	[string]	
data>>failureDetail>>msku	msku	是	[string]	
data>>failureDetail>>msg	错误信息	是	[string]	
返回示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "c18c0dc31a394deca0b9fe57a1ba5cf2.1744873289340",
    "response_time": "2025-04-17 15:01:31",
    "data": {
        "successNum": 0,
        "failureNum": 2,
        "failureDetail": [
            {
                "storeId": 17,
                "asin": "B0B6G8GY5R",
                "msku": "GELPEN001",
                "msg": "店铺状态异常，请于【店铺授权】页面查看授权状态"
            },
            {
                "storeId": 6,
                "asin": "",
                "msku": "PEE-9",
                "msg": "该listing不支持修改库存"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
Listing删除商品标签
销售
下一章节 
解除Listing配对
销售