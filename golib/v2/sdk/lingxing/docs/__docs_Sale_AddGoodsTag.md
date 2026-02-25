Listing新增商品标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/listingManage/bindListingAndTag	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
bindDetail	配对信息	是	[array]	
bindDetail>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	17
bindDetail>>relationId	msku，查询亚马逊Listing 接口对应字段【seller_sku】	是	[string]	HOLDER001
tagIds	标签id数组	是	[array]	
请求示例
{
    "bindDetail": [
        {
            "sid": 42,
            "relationId": "MX-1CIO-Z1TG"
        },
        {
            "relationId": "QG-WQQW-YRK2",
            "sid": 43
        }
    ],
    "tagIds": ["907398794822668564"]
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	错误信息	是	[array]	
response_time	响应时间	是	[string]	2025-03-28 11:20:52
request_id	请求链路id	是	[string]	8f17ddf49d2342f5b26fd616fade6541.152.17431320520980181
data	响应数据	是	[object]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8f17ddf49d2342f5b26fd616fade6541.152.17431320520980181",
    "response_time": "2025-03-28 11:20:52",
    "data": null,
    "total": 0
}
复制
错误
复制成功
 上一章节
修改B2B价格
销售
下一章节 
Listing删除商品标签
销售