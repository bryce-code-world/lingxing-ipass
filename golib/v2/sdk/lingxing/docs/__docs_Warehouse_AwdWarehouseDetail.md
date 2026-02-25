查询AWD库存列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/storage/awdWarehouseDetail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wids	仓库ID列表，使用逗号分隔	否	[string]	
cid	分类ID列表，使用逗号分隔	否	[string]	
bid	品牌ID列表，使用逗号分隔	否	[string]	
attribute	属性值	否	[int]	
asin_principal	ASIN负责人UID列表，使用逗号分隔 * 0、负责人为空	否	[string]	
search_field	搜索字段，指定进行搜索的列 sku product_name seller_sku fnsku asin parent_asin spu spu_name	是	[string]	sku
search_value	搜索值	否	[string]	asdkpasdk
status	状态列表，使用逗号分隔 0、停售 1、在售	否	[string]	
is_hide_zero_stock	是否隐藏零库存 0、不隐藏 1、隐藏	否	[number]	1
offset	分页偏移量	否	[number]	0
length，上限200	分页长度	否	[number]	20
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/storage/awdWarehouseDetail?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "wids": "",
    "cid": "",
    "bid": "",
    "attribute": "",
    "asin_principal": "",
    "search_field": "sku",
    "search_value": "asdkpasdk",
    "status": "",
    "is_hide_zero_stock": 1,
    "offset": 0,
    "length": 20
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	a0d54debf93140f3b58d1ed81e8e3583.157.17255237213100783
response_time	响应时间	是	[string]	2024-09-05 16:08:42
data		是	[object]	
data>>total	总数	是	[int]	1
data>>list		是	[array]	
data>>list>>wname	仓库名	是	[string]	韧啸-US美国仓(AWD)
data>>list>>sid	sid	是	[int]	50
data>>list>>mid	mid	是	[int]	1
data>>list>>seller_name		是	[string]	韧啸-US
data>>list>>nation	国家	是	[string]	美国
data>>list>>asin	asin	是	[string]	B0D3F2W56C
data>>list>>asin_url	asin链接	是	[string]	https://www.amazon.com/dp/B0D3F2W56C
data>>list>>parent_asin	父asin	是	[string]	B0D3F2W56C
data>>list>>seller_sku	MSKU	是	[string]	CN0002
data>>list>>fnsku	FNSKU	是	[string]	X00489S3FJ
data>>list>>sku	SKU	是	[string]	asdkpasdk
data>>list>>product_name	品名	是	[string]	随便spu-3
data>>list>>spu	SPU	是	[string]	asdkpasdk
data>>list>>spu_name	款名	是	[string]	随便spu
data>>list>>small_image_url	预览图链接	是	[string]	https://m.media-amazon.com/images/I/51YvC-1dvXL._SL75_.jpg
data>>list>>pic_url	原图链接	是	[string]	
data>>list>>product_id	产品id	是	[int]	84834
data>>list>>bid	品牌id	是	[int]	0
data>>list>>product_brand_text	品牌文本	是	[string]	
data>>list>>cid	分类id	是	[int]	0
data>>list>>category_level1	分类层级1	是	[string]	
data>>list>>category_level2	分类层级2	是	[string]	
data>>list>>category_level3	分类层级3	是	[string]	
data>>list>>category_Arr	类别数组	是	[array]	
data>>list>>category_text	类别文本	是	[string]	
data>>list>>attribute	属性列表	是	[array]	
data>>list>>attribute>>attr_id	属性id	是	[int]	137
data>>list>>attribute>>attr_name	属性名称	是	[string]	尺寸
data>>list>>attribute>>attr_value	属性值	是	[string]	S
data>>list>>asin_principal_list	Listing负责人	是	[array]	
data>>list>>total_onhand_quantity	AWD在库数量	是	[int]	22
data>>list>>available_distributable_quantity	AWD可用量	是	[int]	10
data>>list>>reserved_distributable_quantity	AWD待发货量	是	[int]	12
data>>list>>awd_quantity_shipped	AWD标发在途数量	是	[int]	0
data>>list>>awd_actual_quantity_shipped	AWD实际在途数量	是	[int]	68
data>>list>>awd_to_fba_quantity_shipped	AWD补货至FBA在途数量	是	[int]	0
data>>list>>total_onhand_quantity_price	AWD在库成本	是	[string]	201.99
data>>list>>available_distributable_quantity_price	AWD可用成本	是	[string]	91.81
data>>list>>reserved_distributable_quantity_price	AWD待发成本	是	[string]	110.17
data>>list>>awd_quantity_shipped_price	AWD标发在途成本	是	[string]	0.00
data>>list>>awd_actual_quantity_shipped_price	AWD实际在途成本	是	[string]	624.32
data>>list>>awd_to_fba_quantity_shipped_price	AWD补货至FBA在途成本	是	[string]	0.00
total	总数	是	[number]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.155.17286133576874025",
    "response_time": "2024-10-11 10:22:38",
    "data": {
        "total": 1,
        "list": [
            {
                "wname": "韧啸-US美国仓(AWD)",
                "sid": 50,
                "mid": 1,
                "seller_name": "韧啸-US",
                "nation": "美国",
                "asin": "B0CPT1TVV7",
                "asin_url": "https://www.amazon.com/dp/B0CPT1TVV7",
                "parent_asin": "B0CPT1TVV7",
                "seller_sku": "HM300-BLACK",
                "fnsku": "",
                "sku": null,
                "product_name": null,
                "spu": "",
                "spu_name": "",
                "small_image_url": "https://m.media-amazon.com/images/I/41yXqZiD0JL._SL75_.jpg",
                "pic_url": "",
                "product_id": 0,
                "bid": 0,
                "product_brand_text": "",
                "cid": 0,
                "category_level1": "",
                "category_level2": "",
                "category_level3": "",
                "category_Arr": [],
                "category_text": "",
                "attribute": [],
                "asin_principal_list": [],
                "total_onhand_quantity": 24,
                "available_distributable_quantity": 14,
                "reserved_distributable_quantity": 10,
                "awd_quantity_shipped": 0,
                "awd_actual_quantity_shipped": 0,
                "awd_to_fba_quantity_shipped": 0,
                "total_onhand_quantity_price": "0.00",
                "available_distributable_quantity_price": "0.00",
                "reserved_distributable_quantity_price": "0.00",
                "awd_quantity_shipped_price": "0.00",
                "awd_actual_quantity_shipped_price": "0.00",
                "awd_to_fba_quantity_shipped_price": "0.00"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询FBA库存列表v2
下一章节 
查询仓库库存明细