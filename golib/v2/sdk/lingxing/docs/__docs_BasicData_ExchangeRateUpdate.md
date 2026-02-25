修改我的汇率
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/settings/exchangeRate/update	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
my_rate	我的汇率【小数位数最多10位】，查询汇率列表 接口对应字段【my_rate】	是	[string]	1.0000
date	汇率年月，查询汇率列表 接口对应字段【date】	是	[string]	2021-08
code	币种，查询汇率列表 接口对应字段【code】	是	[string]	CNY
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "842d7733ef5d4d968bc6877e06c34453.1697004896752",
    "response_time": "2023-10-11 14:14:56",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
查询汇率
基础数据
下一章节 
下载附件
基础数据