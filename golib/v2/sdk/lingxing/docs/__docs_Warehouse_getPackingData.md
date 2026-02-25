查询备货单装箱信息
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/getPackingData	HTTPS	GET	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	OWS210303001
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	7BABECDC-65FE-43B1-20A3-E6BB6EDC928A
response_time	响应时间	是	[string]	2024-06-30 16:26:54
data	响应数据	是	[object]	
data>>overseas_order_no	备货单号	是	[string]	OWS240619003
data>>packaging_type	装箱类型：1 每箱多个sku，2 每箱一个sku	是	[number]	2
data>>box_count	总箱数	是	[number]	2
data>>box_list	装箱信息	是	[array]	
data>>box_list>>box_no	箱号	是	[int]	1
data>>box_list>>height	箱子高(CM)	是	[number]	23
data>>box_list>>length	箱子长(CM)	是	[number]	23
data>>box_list>>width	箱子宽(CM)	是	[number]	34
data>>box_list>>weight	箱子重(KG)	是	[number]	3.3
data>>box_list>>items	商品详情	是	[array]	
data>>box_list>>items>>product_id	商品id	是	[number]	2125186
data>>box_list>>items>>twp_id	三方商品id	是	[number]	0
data>>box_list>>items>>quantity_shipped	装箱数	是	[number]	1
data>>box_list>>items>>match_num	配对数量	是	[number]	1
data>>box_list>>items>>fnsku	FNSKU	是	[string]	
data>>box_list>>items>>sid	店铺id	是	[number]	0
total		是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "832A860C-F8D8-7719-9A55-6253E06C786E",
    "response_time": "2024-08-21 15:29:30",
    "data": {
        "overseas_order_no": "OWS210303001",
        "packaging_type": 2,
        "box_count": 15,
        "box_list": [
            {
                "box_no": 1,
                "height": 289,
                "length": 278,
                "width": 128,
                "weight": 129,
                "items": [
                    {
                        "product_id": 10014,
                        "twp_id": 0,
                        "quantity_shipped": 10,
                        "match_num": 1,
                        "fnsku": "",
                        "sid": 0
                    }
                ]
            },
            {
                "box_no": 2,
                "height": 289,
                "length": 278,
                "width": 128,
                "weight": 129,
                "items": [
                    {
                        "product_id": 10014,
                        "twp_id": 0,
                        "quantity_shipped": 10,
                        "match_num": 1,
                        "fnsku": "",
                        "sid": 0
                    }
                ]
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
更新备货单物流信息
下一章节 
上传备货单装箱信息