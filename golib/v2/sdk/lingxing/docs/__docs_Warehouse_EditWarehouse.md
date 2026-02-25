添加/修改仓库

创建客户自定义仓库id与领星系统本地仓仓库id映射关系，添加/修改本地仓库基础信息

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/storage/wareHouse/edit	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
sys_wid	领星系统仓库id，编辑时必传	否	[int]	1335
wid	客户自定义仓库id【非领星系统ERP内仓库id】	否	[string]	1567
name	仓库名称	是	[string]	Charon
contact	负责人	否	[string]	Charon
telephone	联系电话	否	[string]	15767777777
address	仓库地址	否	[string]	xxxx
remark	备注	否	[string]	ceshi
type	仓库属性：1 -本地仓 3 -海外自建仓，不传默认 1	否	[int]	1
请求示例
{
    "wid": 1567,
    "name": "Charon",
    "contact": "Charon",
    "telephone": "15767777777",
    "address": "xxxx",
    "remark": "ceshi"
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
data	响应数据	是	[array]	 
data->wid	领星系统仓库id	是	[string]	1678
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "E37BFD54-8BEA-3763-85B5-EE025CC27C48",
    "response_time": "2024-07-31 14:38:40",
    "data": {
        "wid": 1678
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询仓库列表
下一章节 
查询本地仓位列表