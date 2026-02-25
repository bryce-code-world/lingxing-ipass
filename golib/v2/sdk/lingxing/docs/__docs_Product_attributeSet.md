添加 / 编辑产品属性

1、接口对属性数据为覆盖式操作，入参属性值会全量覆盖系统里已存在属性内容
2、属性值有关联SPU的情况下，不允许对该属性值有编辑、删除操作
3、如需对已存在属性新增属性值，入参为：该属性下已存在属性值 + 新增属性值
4、如 pa_id 与 pai_id 都不传，视为新增属性

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/attribute/set	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
pa_id	领星属性id	否	[int]	1
attr_name	属性名	是	[string]	颜色
attr_values	属性值数组	是	[array]	
attr_values>>pai_id	领星属性值id	否	[int]	12
attr_values>>attr_value	属性值名称	是	[string]	白色
请求示例
{
    "pa_id": 1,
    "attr_name": "颜色",
    "attr_values": [{
        "pai_id": 12,
        "attr_value": "白色"
    }]
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
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[array]	 
data>>pa_id	属性名	是	[int]	100
data>>pai_id	属性值	是	[int]	1000
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8D680B8B-3C6F-3E16-F055-73DC70EAD977",
    "response_time": "2024-07-29 15:35:49",
    "data": [
        {
            "pa_id": 100,
            "pai_id": 1000
        }
    ]
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 301000,
    "message": "",
    "error_details": "错误：属性已存在 [请求码:081827]",
    "request_id": "46C01A56-D622-55C3-1973-63ABDA2F12F8",
    "response_time": "2024-07-29 15:37:22",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
查询产品属性列表
下一章节 
查询多属性产品列表