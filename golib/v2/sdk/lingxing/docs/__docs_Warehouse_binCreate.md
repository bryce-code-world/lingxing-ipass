添加仓位
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/wareHouseBin/create	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库id	是	[int]	1599
code	仓位名称	是	[string]	Charon
type	仓位类型：
5 可用
6 次品	是	[int]	5
请求示例
{
    "wid": 1599,
    "code": "Charon",
    "type": 5
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
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2020-04-30 17:33:32
data		是	[array]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": {
        "is_error": 0
    },
    "request_id": "9CF2D7AE-E452-1556-3C24-B3AE29BE4567",
    "response_time": "2024-07-31 14:46:34",
    "data": {
        "code": 1,
        "msg": "操作成功",
        "succ_count": 1,
        "error_list": []
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询本地仓位列表
下一章节 
启用、禁用仓位