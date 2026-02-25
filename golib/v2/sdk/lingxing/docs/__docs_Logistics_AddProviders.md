批量添加头程物流商

支持批量创建头程物流商到领星ERP

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/tms/FirstVessel/addProviders	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
providersData	物流商数据，限制20条	是	[array]	
providersData>>logistics_provider_name	物流商名称，不能重复，限制30个字符	是	[string]	Charon物流3456788
providersData>>code	物流商代码，限制20个字符	否	[string]	gogogo
providersData>>remark	备注，限制200个字符	否	[string]	xxxxx
请求示例
{
    "providersData": [
        {
            "logistics_provider_name": "Charon物流3456788",
            "code": "gogogo",
            "remark": "xxxxx"
        }
    ]
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
request_id	请求链路id	是	[string]	2A3B5BEE-CFAC-546B-AD92-C4152BE1B278
response_time	响应时间	是	[string]	2024-08-05 11:13:59
data	响应数据	是	[array]	 
data>>id	物流商对应的id	是	[int]	1712
total	总数	是	[int]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "2A3B5BEE-CFAC-546B-AD92-C4152BE1B278",
    "response_time": "2024-08-05 11:13:59",
    "data": [
        {
            "id": "1712",
            "name": "Charon物流3456788"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询头程物流商列表
下一章节 
批量添加头程物流方式