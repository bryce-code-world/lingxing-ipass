查询Listing列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/listingManage/vcListing/pageList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
vc_store_ids	vc店铺id，查询VC店铺列表 接口对应字段【vc_store_id】	否	[array]	["134225003201380860"]
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>vc_store_id	VC店铺id	是	[string]	134225003201380860
data>>small_min_image_url	在线商品略缩图地址	是	[string]	https://m.sss-amazon.com/images/I/ss.jpg
data>>asin	ASIN	是	[string]	B097MP26YP
data>>asin_url	ASIN地址	是	[string]	https://www.XXX.com/dp/B097MP26YPXPS
data>>msku	MSKU	是	[string]	0590
data>>upc	UPC	是	[string]	750918095383
data>>ean	EAN	是	[string]	0750918095383
data>>item_name	标题	是	[string]	This is a title
data>>parent_asin	父ASIN	是	[string]	B0B53KJRKW
data>>local_sku	SKU	是	[string]	SKUZHCG
data>>local_name	品名	是	[string]	组合采购链接
data>>category_name	本地产品分类名	是	[string]	v22
data>>brand_id	本地产品品牌ID	是	[string]	0
data>>product_id	本地产品ID	是	[string]	3137
data>>classification_rank	小类排名	是	[array]	
data>>classification_rank>>classification_id	分类ID	是	[string]	3480728011
data>>classification_rank>>title	分类名	是	[string]	Shade Sails
data>>classification_rank>>link	分类链接	是	[string]	https://www.aaa.com/48
data>>classification_rank>>rank	排名	是	[string]	16
data>>display_group_rank	大类排名	是	[array]	
data>>display_group_rank>>website_display_group	分类组名	是	[string]	private_brands_display_on_website
data>>display_group_rank>>title	分类名	是	[string]	Our Brands
data>>display_group_rank>>link	分类链接	是	[string]	https://www.xxx.com/rands
data>>display_group_rank>>rank	排名	是	[string]	6615
data>>reviews_num	评论数	是	[string]	12978
data>>stars	星级	是	[string]	4.5
data>>principal_list	负责人列表	是	[array]	
data>>principal_list>>uid	负责人uid	是	[string]	123
data>>principal_list>>real_name	负责人姓名	是	[string]	jack
data>>remark	备注	是	[string]	this is a remark
data>>on_sale_time	开售时间	是	[string]	1970-01-01
data>>status	在线商品状态：
-1已删除
0 停售
1 在售	是	[int]	1
data>>price	优惠金额	是	[string]	59.99
data>>price_currency_icon	优惠金额货币符号	是	[string]	$
返回成功示例

{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "2cda9d50b57e43aa9e951e9890b6c2d6.1697438796964",
    "response_time": "2023-10-16 14:46:39",
    "data": [
        {
            "vc_store_id": "134225003201380864",
            "small_min_image_url": "https://www.xxx.com/images/I/41Z.jpg",
            "asin": "B0BG2CMB5J",
            "asin_url": "https://www.xxx.com/dp/B0BxxxG2CMB5J",
            "msku": "0590",
            "upc": "750918095383",
            "ean": "0750918095383",
            "item_name": "MacBook, iPAD, Dell, HP, Huawei, Phones",
            "parent_asin": "",
            "local_sku": "SKUZHCG",
            "local_name": "组合采购链接",
            "category_name": "",
            "brand_id": "0",
            "category_id": "0",
            "product_id": "3137",
            "classification_rank": [],
            "display_group_rank": [
                {
                    "website_display_group": "fashion_display_on_website",
                    "title": "Clothing, Shoes & Jewelry",
                    "link": "https://www.xxx.com/gp/fashion",
                    "rank": "214440"
                }
            ],
            "reviews_num": "0",
            "stars": "0",
            "principal_list": [
                {
                    "uid": "100003",
                    "real_name": "jack"
                },
            ],
            "remark": "1010",
            "on_sale_time": "1970-01-01",
            "status": 1,
            "price": "123",
            "price_currency_icon": ""
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询VC店铺列表
下一章节 
查询VC订单列表