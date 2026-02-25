删除产品标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/label/operation/v1/label/product/unmarkLabel	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
type	操作类型：
1 删除SKU指定的标签
2 删除SKU全部的标签【此类型下对应sku的label_list为空数组即可】	是	[int]	1
detail_list	标签信息，上限200	是	[array]	
detail_list>>sku	本地产品sku	是	[string]	test-sku-time
detail_list>>label_list	标签名称，上限10	是	[array]	["标签-1","标签-2"]
请求示例
{
    "type": 1,
    "detail_list": [
        {
            "sku": "test-sku-time",
            "label_list": [
                "标签-1",
                "标签-2"
            ]
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
message	返回信息	是	[string]	success
request_id	请求链路id	是	[string]	550d352c-7a05-11ed-b0c7-0242ac1c0004
response_time	响应时间	是	[string]	2022-12-15 20:16:38
data	响应数据	是	[object]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "550d352c-7a05-11ed-b0c7-0242ac1c0004",
    "response_time": "2022-12-15 20:16:38",
    "data": null
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
标记产品标签
下一章节 
查询操作日志