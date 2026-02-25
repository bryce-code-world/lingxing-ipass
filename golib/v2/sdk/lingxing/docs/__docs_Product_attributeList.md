查询产品属性列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/attribute/attributeList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	是	[int]	0
length	分页长度，上限200	是	[int]	20
请求示例
{
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
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8D680B8B-3C6F-3E16-F055-73DC70EAD977
response_time	响应时间	是	[string]	2021-07-18 11:23:47
data	响应数据	是	[object]	
data>>total	总数	是	[int]	589
data>>list	数据列表	是	[array]	
data>>list>>pa_id	属性id	是	[int]	713
data>>list>>attr_name	属性名称	是	[string]	属性16892
data>>list>>create_time	属性创建时间	是	[string]	2021-07-13 21:43:01
data>>list>>item_list	属性值列表	是	[array]	
data>>list>>item_list>>pai_id	属性id	是	[int]	713
data>>list>>item_list>>pa_id	属性值id	是	[int]	1340
data>>list>>item_list>>attr_value	属性值	是	[string]	属性值845722
data>>list>>item_list>>create_time	属性值创建时间	是	[string]	2021-07-13 21:43:01
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "4CFE0142-B0FD-B840-3B3A-42997D251F97",
    "response_time": "2023-07-20 14:14:10",
    "data": {
        "total": 589,
        "list": [
            {
                "pa_id": 713,
                "attr_name": "属性16892",
                "create_time": "2021-07-13 21:43:01",
                "item_list": [
                    {
                        "pai_id": 1340,
                        "pa_id": 713,
                        "attr_value": "属性值845722",
                        "create_time": "2021-07-13 21:43:01"
                    }
                ]
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
添加/编辑本地产品
下一章节 
添加/编辑产品属性