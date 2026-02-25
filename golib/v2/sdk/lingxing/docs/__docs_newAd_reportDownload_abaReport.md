ABA搜索词报告-按周维度
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/abaReport	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始	[int]	2
请求参数
参数名	说明	必填	类型	示例
country	国家代码：如US	是	[string]	US
data_start_time	报表开始日期：每周周日的日期，仅支持最近45天	是	[string]	2023-09-01
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	nd8492d4-7a15-11ed-b162-0242ac1c0004
response_time	响应时间	是	[string]	2022-11-03 14:09:15
data	响应数据	是	[object]	
data>>url	文件下载地址
【下载的文件为zip格式，常用解压工具解压即可】	是	[string]	https://dg-xxxd.com/xxx/xxx/A1AM785xxxffa9bae6
data>>data_start_time	报表日期	是	[string]	2023-09-01
data>>report_period	报表类型：WEEK	是	[string]	WEEK
data>>country	国家代码	是	[string]	US
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "request_id": "nd8492d4-7a15-11ed-b162-0242ac1c0004",
    "response_time": "2023-01-05 11:16:52"
    "data": {
        "url": "https://dg-xxxd.com/xxx/xxx/A1AM785xxxffa9bae6",
        "data_start_time": "2023-09-01",
        "report_period": "WEEK",
        "country": "US"
    }
}
复制
错误
复制成功
 上一章节
查询广告账号列表
下一章节 
操作日志（新）