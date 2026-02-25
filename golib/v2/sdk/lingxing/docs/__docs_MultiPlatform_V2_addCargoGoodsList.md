查询WFS货件可添加商品列表

支持查询可以在货件内添加所选店铺下可用于WFS配送的商品（需要商品的配送方式为Walmart Fulfilled以及WFS Eligible），数据来源为领星系统下的Walmart在线商品

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/cargo/addCargoGoods/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	是	[string]	110418202566107648
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
请求示例
{
    "store_id": "110418202566107648",
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	77ac259a67d5462594c83b80669b6eae.1692331008758
response_time	响应时间	是	[string]	2023-08-18 11:56:49
total	总数	是	[int]	200
data	响应数据	是	[array]	
data>>gtin	GTIN	是	[string]	09123107025675
data>>item_id	商品id	是	[string]	09123107025675
data>>local_name	品名	是	[string]	test1109pm
data>>local_sku	本地商品sku	是	[string]	test1109sku
data>>msku	MSKU	是	[string]	WMUS-DLO-HairRoller-Blue-DLO00030
data>>picture_url	图片地址	是	[string]	https://i5.xxx.com/asr/2cc9b.jpeg
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "2a7df1a6e7e449b49c1bb85294eb2c37.1694769619141",
    "response_time": "2023-09-15 17:20:20",
    "data": [
        {
            "gtin": "09123107025675",
            "item_id": "09123107025675",
            "local_name": "test1109pm",
            "local_sku": "test1109sku",
            "msku": "WMUS-DLO-HairRoller-Blue-DLO00030",
            "picture_url": "https://i5.b.jpeg"
        }
    ],
    "total": 276
}
复制
错误
复制成功
 上一章节
查询退件地址列表
下一章节 
WFS货件暂存