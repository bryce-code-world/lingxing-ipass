查询Walmart Review列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/walmart/queryCommentList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
endDate	结束日期	是	[string]	2024-12-31
pageNum	页码	否	[int]	1
pageSize	每页大小	否	[int]	20
ratings	评分列表	否	[array]	[1,2,3,4,5]
searchDateField	搜索日期字段	否	[string]	commentDate
searchField	搜索字段	否	[string]	msku
searchValue	搜索值列表	否	[array]	["value1","value2"]
startDate	开始日期	是	[string]	2024-01-01
storeIds	店铺ID列表	否	[array]	["123","456"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/multiplatform/walmart/queryCommentList?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "searchField": "msku",
    "searchDateField": "commentDate",
    "startDate": "2023-10-03",
    "endDate": "2025-11-02",
    "pageSize": 20,
    "pageNum": 1
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	响应码	是	[int]	
data	返回数据	是	[object]	
data>>current	当前页码	是	[int]	1
data>>currentSize	当前页记录数	是	[int]	20
data>>hasNextPage	是否有下一页	是	[boolean]	1
data>>hasPreviousPage	是否有上一页	是	[boolean]	
data>>pageCount	总页数	是	[int]	5
data>>records	评论记录列表	是	[array]	
data>>records>>buyerName	买家名称	是	[string]	张三
data>>records>>commentContent	评论内容	是	[string]	商品质量很好
data>>records>>commentDate	评论日期	是	[string]	2024-01-01
data>>records>>commentStar	评论星级	是	[int]	5
data>>records>>commentTitle	评论标题	是	[string]	非常满意
data>>records>>globalTagIds	全局标签ID列表	是	[array]	
data>>records>>id	ID	是	[string]	1001
data>>records>>image	商品图片URL	是	[string]	https://example.com/image.jpg
data>>records>>platformProductNo	平台商品编号	是	[string]	WM123456
data>>records>>productMsku	产品MSKU列表	是	[array]	
data>>records>>remark	备注信息	是	[string]	需要跟进
data>>records>>siteCode	站点代码	是	[string]	US
data>>records>>siteText	站点文本描述	是	[string]	美国站
data>>records>>storeId	店铺ID	是	[string]	123456
data>>records>>storeName	店铺名称	是	[string]	旗舰店
data>>records>>tags	标签列表	是	[array]	
data>>records>>tags>>color	标签颜色	是	[string]	#FF0000
data>>records>>tags>>globalTagId	全局标签ID	是	[long]	10001
data>>records>>tags>>tagName	标签名称	是	[string]	重点关注
data>>records>>title	商品标题	是	[string]	优质商品
data>>size	每页大小	是	[int]	20
data>>total	总记录数	是	[int]	100
error_details	错误详情	是	[array]	
message	提示信息	是	[string]	
request_id	请求id	是	[string]	
response_time	响应时间	是	[string]	
total	总记录数	是	[int]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "f1bfaa74aa1647189f78a4aba95f2174.172.17624866546150449",
    "response_time": "2025-11-07 11:37:35",
    "data": {
        "total": 474,
        "pageCount": 24,
        "current": 1,
        "size": 20,
        "records": [{
                "image": "i",
                "siteCode": "",
                "platformProductNo": "",
                "siteText": "",
                "remark": "",
                "productMsku": [
                    {
                        "msku": "",
                        "localSku": "",
                        "localName": ""
                    }
                ],
                "commentContent": "",
                "storeId": "",
                "title": " ",
                "buyerName": "T",
                "tags": [
                    {
                        "color": "#",
                        "globalTagId": 1,
                        "tagName": ""
                    }
                ],
                "commentStar": 1,
                "globalTagIds": [
                    ""
                ],
                "commentTitle": "",
                "commentDate": "2023-12-26",
                "storeName": "",
                "id": ""
            }],
        "hasNextPage": true,
        "hasPreviousPage": false,
        "currentSize": 20
    },
    "total": 474
}
复制
错误
复制成功
 上一章节
平台仓发货单发货
下一章节 
查询平台订单列表