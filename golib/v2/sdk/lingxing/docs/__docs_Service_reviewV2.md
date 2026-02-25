查询评论管理 - Review(新)
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/service/v3/data/mws/reviews	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sort_field	排序类型	否	[string]	review_date
sort_type	排序	否	[string]	desc
sids	店铺id，多个用逗号分隔 ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[string]	1
mids	站点id，多个用逗号分隔	否	[string]	1
principal_uids	lisitng负责人，多个用逗号分隔	否	[string]	128402,128637
search_field	搜索字段:
asin ASIN
parent_asin 父ASIN
remark 备注
amazon_order_id 订单号
author 买家信息
review_id Review ID
buyer_email 买家
last_title 评价标题	否	[string]	amazon_order_id
search_value	搜索值	否	[string]	12345
date_field	时间搜索类型:
review_time 评价时间
create_time 创建时间
last_update_time 更新时间	是	[string]	review_time
start_date	开始时间，格式：Y-m-d	是	[string]	2024-06-06
end_date	结束时间，格式：Y-m-d	是	[string]	2024-09-04
status	状态，多个用逗号分隔:
0 待处理
1 处理中
2 已完成	否	[string]	0,1,2
star	星级，多个用逗号分隔	否	[string]	1,2,3,4,5
review_modified_status	内容，多个用逗号分隔:
-1 已删除
0 未标识
1 已变更	否	[string]	-1,1,0
mark	标识，多个用逗号分隔:
is_vp
is_er
is_topc
is_topr
is_vine	否	[string]	is_vp,is_er,is_topc,is_topr,is_vine
cs_principal_uids	处理人，多个用逗号分隔	否	[string]	10329601
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
cids	分类id，多个用逗号分隔	否	[string]	30
global_tag_ids	标签id，多个用逗号分隔	否	[string]	907464870663877373
match_types	匹配类型，多个用逗号分隔，默认传空字符串	否	[string]	1,2,3,4,5
请求示例
{
  "sort_field": "review_date",
  "sort_type": "desc",
  "search_field": "review_id",
  "search_value": "R1KKLEHWNZWH05",
  "start_date": "2021-06-01",
  "end_date": "2024-09-06",
  "date_field": "review_time",
  "offset": 0,
  "length": 20,
  "cids": "",
  "global_tag_ids": "",
  "match_types": ""
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	a8779aecad254b1ca0d34df46b00972b.1725518992094
response_time	响应时间	是	[string]	2024-09-05 14:49:52
data	响应数据	是	[array]	
data>>small_image_url	图片	是	[string]	https://image-1254213275.cos.ap-guangzhou.myqcloud.com/dev/20210914/USB%E5%A3%81%E5%BC%8F%E5%85%85%E7%94%B5%E5%99%A8.jpg
data>>asin	ASIN	是	[string]	B0ABC60875
data>>seller_sku	MSKU	是	[array]	["MSKU26323CA","MSKU4C1DA7B"]
data>>last_star	评级 - 星级	是	[number]	1
data>>last_title	评级 - 标签	是	[string]	Very Pleased
data>>review_likes	点赞数	是	[number]	0
data>>review_id	Review ID	是	[string]	3FA1E258B9D40D
data>>review_url	评价链接	是	[string]	https://www.amazon.com/gp/xxxx/xxxx7FLN9VVKY
data>>last_content	评价内容	是	[string]	Very easy to use. Works quickly. Love that it's lightweight and compact.
data>>parent_asin	父ASIN	是	[array]	["B0ABC53810"]
data>>item_name	标题	是	[array]	["lingxing amazon product title DBEC4F0A76"]
data>>local_info	本地信息	是	[array]	
data>>local_info>>local_sku	本地SKU	是	[string]	28520211130
data>>local_info>>local_name	本地品名	是	[string]	首页-测试用1
data>>local_info>>category_name	分类名	是	[string]	产品管理-Listing-退货订单-分类
data>>author	买家信息	是	[string]	评论者376035
data>>images	评论图片链接	是	[array]	["https://m.media-amazon.com/images/I/61SF10hpIIL.jpg"]
data>>videos	评论视频链接	是	[array]	["https://m.media-amazon.com/images/S/vse-vms-transcodiing-artifact-eu-west-1-prod/e86fbba7-b53e-42f1-ba1f-599df347a600/default.vertical.jobtemplate.mp4.480"]
data>>is_vp		是	[number]	1
data>>seller_name	店铺	是	[array]	["店铺2"]
data>>marketplace	国家	是	[string]	美国
data>>review_date	评价时间	是	[string]	2021-08-21
data>>create_time	创建时间	是	[string]	2021-08-23
data>>update_time	更新时间	是	[string]	2021-08-23 15:11:10
data>>crawl_date	操作时间	是	[string]	2021-08-23 15:11:10
data>>amazon_order_list	订单号列表	是	[array]	
data>>amazon_order_list>>seller_name	店铺	是	[string]	店铺2
data>>amazon_order_list>>amazon_order_id	订单号	是	[string]	394-8475169-8475563
data>>amazon_order_list>>buyer_email	买家邮箱	是	[string]	
data>>buyer_email	买家邮箱	是	[array]	
data>>remark	备注	是	[string]	
data>>status	处理状态：0 待处理，1 处理中，2 已完成	是	[number]	0
data>>tags	标签	是	[array]	
data>>cs_principals	处理人	是	[array]	
total	总数	是	[number]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a016098c47674ff99ca7ac177e54111c.1725616690741",
    "response_time": "2024-09-06 17:58:10",
    "data": [
        {
            "small_image_url": "https://m.media-amazon.com/images/I/81h9TDdkM-L._SL75_.jpg",
            "asin": "B07XKHF683",
            "seller_sku": [
                "YE-2XPZ-XKY1-A"
            ],
            "last_star": 5,
            "last_title": "Kinda unique for me",
            "review_likes": 0,
            "review_id": "R1KKLEHWNZWH05",
            "review_url": "https://www.amazon.com/gp/xxxx/xxxx7FLN9VVKY",
            "last_content": "  I love it,  so perfect for my room.",
            "parent_asin": [
                "B08CCRG337"
            ],
            "item_name": [
                "HEETA 2 Stück Duschablage Ohne Bohren, Rostfreies Badezimmer Regal mit Selbstklebender Wandmontage, Starke Duschregal ohne Bohren Küchenregale, Edelstahl-Organizer und Aufbewahrung (13 Inches)"
            ],
            "local_info": [
                {
                    "local_sku": "2323235",
                    "local_name": "sadgacv",
                    "category_name": "3C类"
                }
            ],
            "author": "Mae",
            "is_vp": 0,
            "seller_name": [
                "JHL-DE-2"
            ],
            "marketplace": "德国",
            "review_date": "2023-08-05",
            "create_time": "2021-11-27",
            "update_time": "2021-12-31 21:16:37",
            "crawl_date": "2021-12-31 21:16:37",
            "amazon_order_list": [],
            "buyer_email": [],
            "remark": "",
            "status": 0,
            "tags": [],
            "cs_principals": []
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询评价管理-Review
下一章节 
查询评价管理-1-3星Feedback列表