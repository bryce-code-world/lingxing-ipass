获取自定义出库类型
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/outbound/getCustomTypes	HTTPS	POST	1
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	D76CFAC8-3385-0A42-8916-A168F3449996
response_time	响应时间	是	[string]	2022-03-23 09:17:03
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>list	类型列表	是	[array]	
data>>list>>id	类型ID	是	[Long]	154525
data>>list>>name	类型名称	是	[string]	demo名称
data>>list>>is_delete	是否删除：
0 否
1 是	是	[string]	1
data>>list>>status	状态：
0 关闭
1 开启	是	[string]	0
data>>total	行数	是	[int]	5
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D5AB0A99-6F30-2DCD-BCB1-579331B96AE7",
    "response_time": "2024-08-01 16:02:37",
    "data": {
        "list": [
            {
                "id": "230461637276049920",
                "name": "移除",
                "status": 1,
                "is_delete": 0
            },
            {
                "id": "230470190590665216",
                "name": "库存初始化入库",
                "status": 1,
                "is_delete": 0
            }
        ],
        "total": 2
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
获取自定义入库类型
下一章节 
创建待调拨的调拨单