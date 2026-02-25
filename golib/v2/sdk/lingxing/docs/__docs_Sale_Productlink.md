批量添加/编辑Listing配对

支持推送Listing与本地仓库SKU的配对关系到领星ERP，实现SKU配对

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/storage/product/link	HTTPS	POST	3
请求参数
参数名	说明	必填	类型	示例
data		是	[array]	
data>>seller_id	亚马逊店铺id ,对应查询亚马逊店铺列表接口对应字段【seller_id】	否	[string]	A4373BD6018725
data>>marketplace_id	市场id	否	[string]	xxxxxxxxxxxxxxxx
data>>msku	msku	是	[string]	xxxx
data>>sku	本地sku	是	[string]	xxx
data>>is_sync_pic	是否同步listing图片：0 否，1 是	是	[int]	0
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	ADD67A0B-7619-D681-BE36-BE92261CF931
response_time	响应时间	是	[string]	2020-04-30 17:33:32
data	响应数据	是	[array]	
data>>total	总配对个数	是	[int]	3
data>>success	配对成功数	是	[int]	3
data>>error	配对失败数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "ADD67A0B-7619-D681-BE36-BE92261CF931",
    "response_time": "2020-07-24 10:32:19",
    "data": {
        "total": 1,
        "success": 1,
        "error": 0
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 1000,
    "message": "success",
    "error_details": [
        {
            "index": 0,
            "message": "错误：msku：QLSLR_001000000,不存在"
        }
    ],
    "request_id": "8450EC9F-F4E1-2DC5-23D8-E5A5E04B2251",
    "response_time": "2020-07-24 10:36:25",
    "data": {
        "total": 1,
        "success": 0,
        "error": 1
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
批量分配Listing负责人
销售
下一章节 
批量修改Listing价格
销售