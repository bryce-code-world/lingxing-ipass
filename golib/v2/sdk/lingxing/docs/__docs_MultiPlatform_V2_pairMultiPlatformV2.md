批量添加、编辑多平台配对关系
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/listing/v2/pairMultiPlatform	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
pair_multi_platform_list	配对数据，上限为5000条	是	[array]	
pair_multi_platform_list>>msku	msku	是	[string]	231428567
pair_multi_platform_list>>store_id	店铺id	是	[string]	110000000018010004
pair_multi_platform_list>>sku	sku	是	[string]	wjc10_sku
请求示例
{
    "pair_multi_platform_list": [
        {
            "msku": "231428567",
            "store_id": "110000000018010004",
            "sku": "wjc10_sku"
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
message	提示信息	是	[string]	
data	响应数据	是	[object]	
data>>message	配对结果信息:
全部配对成功
部分配对失败
全部配对失败	是	[string]	全部配对失败
data>>error_detail_list	返回错误数据【只返回配对失败数据】	是	[array]	
data>>error_detail_list>>msku	msku	否	[string]	msku123456
data>>error_detail_list>>store_id	店铺id	否	[string]	110101363085123584
data>>error_detail_list>>error_message	错误信息	否	[string]	msku123456+110101363085123584配对失败，SKU不存在
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "response_time": "2023-03-07 15:19:49",
    "request_id": "8f6a4ad0-73b9-4998-a0ff-817e4da4d561.1670916113614",
    "data": {
        "message": "全部配对失败",
        "error_detail_list": [
            {
                "msku": "231428567",
                "store_id": "110000000018010004",
                "error_message": "231428567+110000000018010004配对失败，店铺不存在"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
查询多平台配对列表
下一章节 
预发货