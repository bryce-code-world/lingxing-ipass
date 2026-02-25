添加 / 编辑捆绑产品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/product/setBundled	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sku	SKU（添加时必填）	是	[string]	skuAdd
product_name	品名（添加时必填）	是	[string]	skuAddProductName
picture_list	产品图片信息	否	[array]	
picture_list>>pic_url	产品图片链接	否	[string]	http://www.news.cn/xxx/1129400title0h.jpg
picture_list>>is_primary	是否产品主图：0 否，1 是	否	[int]	1
model	型号	否	[string]	model s
unit	单位（商品单位：套、个、台）	否	[string]	个
status	状态【默认1】：0 停售，1 在售，2 开发中，3 清仓	否	[int]	1
category_id	分类id,与分类同时存在时，优先取分类id	否	[int]	1
category	分类	否	[string]	分类1
brand_id	品牌id，与品牌同时存在时，优先取品牌id	否	[int]	2
brand	品牌	否	[string]	品牌2
product_developer	开发者名称	否	[string]	张三
product_developer_uid	开发者id，与开发者名称同时填写时，以开发者id为准	否	[int]	12
product_duty_uids	负责人id	否	[array]	[1,2,3]
is_append_product_duty	负责人是否追加创建人：0 否，1 是；默认1，该字段只有编辑SKU时该才生效	否	[int]	1
product_creator_uid	创建人ERP id，默认 api 用户id	否	[int]	13
description	商品描述	否	[string]	有logo 白盒
group_list	组合商品列表，捆绑产品子产品的总数量要大于1	否	[array]	
group_list>>sku	子商品	否	[string]	skuAdd1-1
group_list>>quantity	商品比例数	否	[int]	2
group_list>>cost_ratio	费用比例，默认为空，若填写则每项必填，且总和为1	否	[int]	2
请求示例
{
    "sku": "skuAdd",
    "product_name": "skuAddProductName",
    "picture_list": [{
        "pic_url": "http://www.news.cn/xxx/1129400title0h.jpg",
        "is_primary": 1
    }],
    "model": "model s",
    "unit": "个",
    "status": 1,
    "category_id": 1,
    "category": "分类1",
    "brand_id": 2,
    "brand": "品牌2",
    "product_developer": "张三",
    "product_developer_uid": 12,
    "product_duty_uids": [1, 2, 3],
    "is_append_product_duty": 1,
    "product_creator_uid": 13,
    "description": "有logo 白盒",
    "group_list": [{
        "sku": "skuAdd1-1",
        "quantity": 2
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	81CA6A9A-A5C4-8B5E-5FED-E1C2EC04753C
response_time	响应时间	是	[string]	2020-12-01 12:11:12
data	响应数据	是	[object]	
data>>product_id	本地产品id	是	[int]	2125067
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8290CE90-93C5-4529-3A51-CA45B7C0CF0D",
    "response_time": "2023-02-27 15:41:30",
    "data": {
        "product_id": 2125067
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询捆绑产品关系列表
下一章节 
查询产品辅料列表