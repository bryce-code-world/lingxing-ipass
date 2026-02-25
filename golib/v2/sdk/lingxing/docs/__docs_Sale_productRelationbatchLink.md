配对/批量配对
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/vcservice/productRelation/batchLink	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sidAsins	配对的sid和asin对象数组	是	[array]	[{"sid":"134228919447351298","asin":"B09N15BYDM"},{"sid":"134228919447351298","asin":"B09N15ZMVW"},{"sid":"134228919447351298","asin":"B08ZKG8WRD"}]
sidAsins>>sid	店铺id(seller表主键)	是	[string]	134228919447351298
sidAsins>>asin	asin	是	[string]	B09N15BYDM
productId	本地商品表主键ID	是	[number]	5913
isSyncPic	是否同步图片到本地商品	是	[number]	0
请求示例
{
    "sidAsins": [{
        "sid": "134228919447351298",
        "asin": "B09N15BYDM"
    }],
    "productId": 5913,
    "isSyncPic": 0
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	消息提示	是	[number]	0
message	消息	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求id	是	[string]	30e6ccc08ef5455595188bd9819197a3.1753340999741
response_time	响应时间	是	[string]	2025-07-24 15:10:00
data	结果data	是	[boolean]	true
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "30e6ccc08ef5455595188bd9819197a3.1753340999741",
    "response_time": "2025-07-24 15:10:00",
    "data": true,
    "total": 0
}
复制
错误
复制成功
 上一章节
VC订单-打印标签【DF】
下一章节 
查询VC发货单列表