上传本地产品图片

支持本地产品图片上传到领星ERP系统内

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/product/uploadPictures	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sku	本地产品SKU	是	[string]	skuAdd
picture_list	产品图片信息	是	[array]	
picture_list>>pic_url	产品图片链接	是	[string]	http://www.news.cn/xx/1129400412e0h.jpg
picture_list>>is_primary	是否产品主图：0 否，1 是	是	[int]	1
请求示例
{
    "sku": "skuAdd",
    "picture_list": [{
        "pic_url": "http://www.news.cn/xx/1129400412e0h.jpg",
        "is_primary": 1
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
request_id	请求链路id	是	[string]	850E2256-7F4C-DE28-1444-99F5BFBA5C4F
response_time	响应时间	是	[string]	2023-03-22 16:03:33
data	响应数据	是	[object]	
data>>sku	本地产品SKU	是	[string]	123123AAA21231
data>>picture_list	产品图片信息	是	[array]	
data>>picture_list>>pic_url	已上传到领星的图片链接	是	[string]	https://image.distxx.com/xxx/3625f14b09704b7cb3953af04e1fbb51.jpg
data>>picture_list>>customer_url	客户提交的的图片链接	是	[string]	http://www.news.cn/xx/1129400412e0h.jpg
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "850E2256-7F4C-DE28-1444-99F5BFBA5C4F",
    "response_time": "2023-03-22 16:03:33",
    "data": {
        "sku": "123123AAA21231",
        "picture_list": [{
            "pic_url": "https://image.distxx.com/xxx/3625f14b09704b7cb3953af04e1fbb51.jpg",
            "customer_url": "http://www.news.cn/xx/1129400412e0h.jpg"
        }]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
添加/编辑产品分类
下一章节 
查询产品标签