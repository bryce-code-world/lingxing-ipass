查询产品标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/label/operation/v1/label/product/list	HTTPS	GET	10
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	返回信息	是	[string]	success
request_id	请求链路id	是	[string]	550d352c-7a05-11ed-b0c7-0242ac1c0003
response_time	响应时间	是	[string]	2022-12-15 20:16:38
data	响应数据	是	[object]	
data>>list	列表数据	是	[array]	
data>>list>>label_id	标签id	是	[string]	907203021760103279
data>>list>>label_name	标签名称	是	[string]	1
data>>list>>gmt_created	创建时间	是	[int]	1657507264000
data>>total	总数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "550d352c-7a05-11ed-b0c7-0242ac1c0003",
    "response_time": 2022-12-15 20: 16: 38,
    "data": {
        "list": [
            {
                "label_id": "4864548789455617",
                "label_name": "精品",
                "gmt_created": 49
            }
        ],
        "total": 1
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "类型错误",
    "request_id": "0C02418E-80C3-5924-8F80-8CD931FFD939",
    "response_time": "2022-12-15 21:12:45",
    "data": null
}
复制
错误
复制成功
 上一章节
上传本地产品图片
下一章节 
创建产品标签