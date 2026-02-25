查询产品分类列表

支持查询本地产品的分类列表，对应【产品】>【产品分类】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/category	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000，上限1000	否	[int]	1000
data>>ids	分类ID	否	[array]	1,2,3
请求示例
{
    "offset": 0,
    "length": 1000,
    "ids": [1,2,3]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	83FD7703-B573-4BE0-058B-7D34A94EBD4E
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	返回数据	是	[array]	
data>>cid	分类ID	是	[int]	2
data>>parent_cid	父级分类ID	是	[int]	
data>>title	分类名称	是	[string]	111
data>>category_code	分类简码	是	[string]	Bo135152
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "83FD7703-B573-4BE0-058B-7D34A94EBD4E",
    "response_time": "2020-05-18 11:23:47",
    "data": [
        {
            "cid": 1,
            "title": "分类1",
            "parent_cid": 0,
            "category_code": "Bo135152"
        },
        {
            "cid": 2,
            "title": "分类2",
            "parent_cid": 0,
            "category_code": "Bo135152"
        }
    ],
    "total": 130
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": [],
    "request_id": "9712C7C8-F677-247E-A908-2AB0790F676B",
    "response_time": "2021-11-11 18:04:11",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
添加/编辑产品品牌
下一章节 
添加/编辑产品分类