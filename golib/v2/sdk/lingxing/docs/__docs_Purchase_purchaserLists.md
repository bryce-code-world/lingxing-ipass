查询采购方列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/purchaser/lists	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认500	否	[int]	1000
请求示例
{
    "offset": 0,
    "length": 1000
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
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2024-08-02 17:32:21
data	响应数据	是	[object]	
data>>total	总数	是	[int]	64
data>>list	列表	是	[array]	
data>>list>>purchaser_id	采购方id	是	[int]	139
data>>list>>name	采购方名称	是	[string]	测试采购方
data>>list>>address	地址	是	[string]	深圳市南山区
data>>list>>contact_phone	联系方式	是	[string]	15632145622
data>>list>>contacter	联系人	是	[string]	小王
data>>list>>email	邮箱	是	[string]	dds@xxx.com
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "929A1D7D-D656-0EA0-C78A-A686790C7090",
    "response_time": "2024-08-02 17:32:21",
    "data": {
        "total": 64,
        "list": [
            {
                "purchaser_id": 139,
                "name": "测试采购方",
                "address": "深圳市南山区",
                "contact_phone": "15632145622",
                "contacter": "小王",
                "email": "    dds@xxx.com"
            }
        ]
    },
    "total": 64
}
复制
错误
复制成功
 上一章节
添加/修改供应商
下一章节 
查询采购计划列表