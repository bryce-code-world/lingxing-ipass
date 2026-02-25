查询评价管理-Review
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/v2/data/mws/reviews	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
start_date	开始评论时间，闭区间，格式：Y-m-d	是	[string]	2024-08-05
end_date	结束评论时间，闭区间，格式：Y-m-d	是	[string]	2024-08-05
offset	分页偏移量，默认0	是	[int]	0
length	分页长度	是	[int]	20
date_field	时间类型:
review_date 评价时间【默认值】
create_time 创建时间	否	[string]	review_date
请求示例
{
    "offset": 0,
    "length": 20,
    "sid": 1,
    "start_date": "2024-08-05",
    "end_date": "2024-08-05",
    "date_field": "review_date"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	73A0718B-AA68-E120-D28B-F7175018175D
response_time	响应时间	是	[string]	2022-04-19 15:16:18
data	响应数据	是	[array]	
data>>asin	ASIN	是	[string]	B0ABC6760875
data>>last_star	星级	是	[int]	5
data>>last_title	评论标题	是	[string]	Very Pleased
data>>last_content	评价内容	是	[string]	Very easy to use. Works quickly. Love that it's lightweight and compact.
data>>review_likes	点赞数	是	[int]	1000
data>>author	评论客户	是	[string]	评论者363934
data>>author_id	评论客户id	是	[string]	
data>>review_date	评论时间	是	[string]	2021-04-09
data>>review_id	Review ID	是	[string]	0002118826800D
data>>status	处理状态：0 待处理，1 处理中，2 已完成	是	[int]	2
data>>update_time	更新时间，时间戳	是	[int]	1624856476
data>>create_time	创建时间，时间戳	是	[int]	1618164429
data>>review_modified_status	评论状态：-1 已删除，0 无标识，1 已变更	是	[int]	1
data>>remark	备注	是	[string]	
data>>amazon_order_list	匹配订单列表	是	[array]	
data>>marketplace	国家	是	[string]	美国
data>>is_vp	是否是VP：0 否，1 是	是	[int]	
data>>is_er	是否是ER：0 否，1 是	是	[int]	0
data>>is_topc	是否是TOPC：0 否，1 是	是	[int]	0
data>>is_topr	是否是TOPR：0 否，1 是	是	[int]	0
data>>is_vine	是否是VINE：0 否，1 是	是	[int]	0
data>>asin_url	ASIN 链接	是	[string]	https://www.amazon.com/dp/B0AB76C60875
data>>review_url	评价链接	是	[string]	https://www.amazon.com/xx/ref=cmttl?ie=UTF8&ASIN=B0ABC7605
data>>seller_name	店铺名	是	[array]	["店铺2","店铺5"]
data>>sids	店铺ID	是	[array]	["2","5"]
data>>attachments	附件	是	[array]	
data>>attachments>>type	附件类型：image 图片，video 视频	是	[string]	image
data>>attachments>>url	链接地址	是	[string]	https://images-na.ssl-images-amazon.com/images/xxx.jpg
total	总数	是	[int]	6556
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "42293103-574D-CF2A-1767-822AE1942984",
    "response_time": "2022-04-19 15:29:47",
    "data": [
        {
            "asin": "ASIN",
            "last_star": "星级",
            "last_title": "评价标题",
            "last_content": "评价内容",
            "review_likes": "点赞数",
            "author": "评价客户",
            "author_id": "评价客户ID",
            "review_date": "评论时间",
            "review_id": "Review ID",
            "status": "处理状态",
            "update_time": "更新时间",
            "create_time": "创建时间",
            "review_modified_status": "评价状态",
            "remark": "备注",
            "amazon_order_list": ["订单号列表","订单号列表"],
            "marketplace": "国家",
            "is_vp": "是否VP",
            "is_er": "是否ER",
            "is_topc": "是否TOPC",
            "is_topr": "是否TOPR",
            "is_vine": "是否VINE",
            "asin_url": "ASIN 链接",
            "review_url": "评价链接",
            "seller_name": ["店铺","店铺"],
            "sids": ["店铺id","店铺id"],
            "attachments": [
                {
                    "type": "image",
                    "url": "https://images-na.ssl-images-amazon.com/images/I/71w0PBmui+L._SY88.jpg"
                },
                {
                    "type": "video",
                    "url": "https://images-na.ssl-images-amazon.com/images/I/C1OFT0iGsoS.mp4"
                },
            ]
        }
    ],
    "total": "1"
}
复制
错误
复制成功
 上一章节
查询邮件详情
下一章节 
查询评价管理-Review(新)