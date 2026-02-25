作废费用单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/fee/management/open/feeManagement/otherFee/discard	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
numbers	费用单号，上限200	是	[array]	["FY231009000001"]
请求示例
{
    "numbers": [
        "FY231009000001"
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
msg	消息提示	是	[string]	操作成功
data	响应数据	是	[object]	 
返回成功示例
{
    "code": 0,
    "msg": "操作成功",
    "data": {}
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "msg": null,
    "data": {
        "FY231009000001": "已作废状态的费用单无法作废"
    }
}
复制
错误
复制成功
 上一章节
编辑费用单
下一章节 
删除费用单