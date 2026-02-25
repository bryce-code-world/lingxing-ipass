店铺维度-批量删除目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/store/batchDelete	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份【只允许去年、今年、明年】	是	[int]	2024
sids	需要删除的店铺id列表 ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[array]	[135,102]
请求示例
{
    "assessYear": "2024",
    "sids": [
        135,
        102
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，1 成功	是	[int]	1
msg	返回消息	是	[string]	操作成功
data	删除条数	是	[array]	2
traceId	请求链路id	是	[string]	4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111
返回成功示例
{
    "code": 1,
    "msg": "操作成功",
    "data": 2,
    "traceId": "cf6d9586d8074aec92af0f1a85d18d9c.1670328945764",
    "success": true
}
复制
错误
复制成功
 上一章节
店铺维度-批量新增/更新目标
下一章节 
组织维度-批量查询目标