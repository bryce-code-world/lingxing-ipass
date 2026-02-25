创建产品标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/label/operation/v1/label/product/create	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
label	标签名称，最长15个字符，中间不能有空格	是	[string]	精品
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
request_id	请求链路id	是	[string]	550d352c-7a05-11ed-b0c7-0242ac1c0004
response_time	响应时间	是	[string]	2022-12-15 20:16:38
data	响应数据	是	[object]	
data>>label_name	标签名称	是	[string]	精品
data>>label_id	标签id	是	[string]	4864548789455617
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "550d352c-7a05-11ed-b0c7-0242ac1c0004",
    "response_time": "2022-12-15 20:16:38",
    "data": {
        "label_name": "精品",
        "label_id": "4864548789455617"
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
查询产品标签
下一章节 
标记产品标签