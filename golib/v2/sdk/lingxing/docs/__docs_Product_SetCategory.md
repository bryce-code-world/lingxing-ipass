添加 / 编辑产品分类

支持添加/编辑本地产品分类

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/category/set	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
data	请求数据	是	[array]	
data>>id	为空时新增，不为空时编辑，查询本地产品分类列表对应cid字段	否	[int]	1
data>>parent_cid	父级分类id	否	[int]	0
data>>title	分类名称	是	[string]	衣服
data>>category_code	分类简码	是	[string]	Bo135152
请求示例
{
    "data": [
        {
            "id":1,
            "parent_id":0,
            "title": "衣服",
            "category_code": "Bo135152"
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
message	返回信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	2021-01-11 14:49:12
data	成功的数据	是	[array]	
data>>id	分类id	是	[string]	3
data>>title	分类名称	是	[string]	衣服
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "71936B27-D0AE-8436-7F4F-62BEC7201AAE",
    "response_time": "2021-11-11 18:21:45",
    "data": [
        {
            "title": "华为2",
            "parent_cid": "",
            "id": "41"
        }
    ],
    "total": 1
}
复制
错误
复制成功
返回失败示例
{
    "code": 1000,
    "message": "业务处理失败",
    "error_details": [
        {
            "params": {
                "id": "112",
                "parent_cid": ""
            },
            "message": "错误：未找到对应的分类"
        }
    ],
    "request_id": "B7527DF4-4518-3EA4-DC53-8FBECA93571B",
    "response_time": "2021-11-11 18:22:57",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询产品分类列表
下一章节 
上传本地产品图片