组织维度-批量删除目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/user/batchDelete	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份【只允许去年、今年、明年】	是	[int]	2024
assessType	考核指标：1 销售额，2 销量	是	[int]	1
uidList	用户id集合	是	[array]	[1001,1002]
请求示例
{
    "assessYear": "2024",
    "assessType": 1,
    "uidList": [
        1001,
        1002
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
data	操作成功数据条数	是	[int]	2
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
组织维度-批量新增/更新目标
下一章节 
查询多平台店铺信息