添加/编辑产品品牌

支持添加/编辑本地产品品牌信息

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/storage/brand/set	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
data	请求数据	是	[array]	
data>>id	为空时表新增，不为空时表编辑，查询本地产品品牌列表对应bid字段	否	[int]	12
data>>title	品牌名称	是	[string]	华为2
data>>brand_code	品牌简码	否	[string]
请求示例
{
    "data": [
        {
            "title": "PM1610as01",
            "brand_code": "code"            
        },
        {
            "id":12,
            "title": "PM1610as02"
            "brand_code": "code"
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	
message	返回信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	E75A1E60-A812-6230-4182-2680FC1BC58A
response_time	响应时间	是	[string]	2021-01-11 14:49:12
data	成功的数据	是	[array]	
data>>id	品牌id	是	[string]	12
data>>title	品牌名称	是	[string]	华为2
data>>brand_code	品牌简码	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "E75A1E60-A812-6230-4182-2680FC1BC58A",
    "response_time": "2021-11-11 18:18:50",
    "data": [
        {
            "id": 12,
            "title": "华为2",
            "brand_code": "code"           
        }
    ]
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "data => data不能为空"
    ],
    "request_id": "8892C7F3-2083-4001-0AA1-57A2BAAE3543",
    "response_time": "2021-11-11 18:17:08",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询产品品牌列表
下一章节 
查询产品分类列表