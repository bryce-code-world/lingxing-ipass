启用、禁用仓位
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/wareHouseBin/switchStatus	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库id	是	[string]	1
whbCode	仓位名称	是	[string]	xx
status	仓位状态：0 禁用，1 启用	是	[int]	1
请求示例
{
    "wid": 1678,
    "whbCode": "Charon",
    "status": 1
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	信息提示	是	[string]	success
error_details	错误信息	是	[array]	
error_details>>errorMsg	错误信息	是	[string]	1:失败,:
error_details>>is_error	是否失败	是	[string]	1:失败,0:操作成功
error_details>>ids	失败的仓位id	是	[array]	1:失败,:
request_id	请求链路id	是	[string]	D17B0688-C52D-1BA9-1BCC-50AA00058A3F
response_time	响应时间	是	[string]	2022-08-23 16:56:31
data		是	[array]	
data>>successMsg	成功信息	是	[string]	1:失败,
 上一章节
添加仓位
下一章节 
海外仓sku配对