获取自定义入库类型
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/inbound/getCustomTypes	HTTPS	POST	1
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	否	[int]	0
message	消息提示	否	[string]	success
error_details	错误信息	否	[array]	
request_id	请求链路id	否	[string]	D76CFAC8-3385-0A42-8916-A168F3449996
response_time	响应时间	否	[string]	2022-03-23 09:17:03
total	总数	否	[int]	0
data	响应数据	否	[object]	
data>>list	类型列表	否	[array]	
data>>list>>id	类型ID	否	[Long]	154525
data>>list>>name	类型名称	否	[string]	demo名称
data>>list>>is_delete	是否删除：
0 否
1 是	否	[string]	1
data>>list>>status	状态：
0 关闭
1 开启	否	[string]	0
data>>total	行数	否	[int]	5
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "794DB00B-76D7-3336-8712-235C2F2A2859",
    "response_time": "2024-08-01 15:57:53",
    "data": {
        "list": [
            {
                "id": "230454385165754880",
                "name": "test",
                "status": 1,
                "is_delete": 1
            },
            {
                "id": "230454385542291968",
                "name": "借用",
                "status": 1,
                "is_delete": 0
            },
            {
                "id": "230461637062251008",
                "name": "采用",
                "status": 1,
                "is_delete": 0
            }
        ],
        "total": 3
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
创建已完成的盘点单
下一章节 
获取自定义出库类型