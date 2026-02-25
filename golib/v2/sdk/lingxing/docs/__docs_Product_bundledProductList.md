查询捆绑产品关系列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/bundledProductList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000，上限1000	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	
message	消息提示	是	[string]	
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	6D1422D5-B767-675D-42F6-F604319C2333
response_time	响应时间	是	[string]	2023-03-09 14:45:40
data	响应数据	是	[array]	
data>>id	捆绑产品ID	是	[int]	
data>>sku	捆绑产品SKU	是	[string]	
data>>product_name	捆绑产品名	是	[string]	
data>>cg_price	捆绑产品采购成本	是	[number]	
data>>status_text	产品状态：停售、在售、开发中、清仓	是	[string]	
data>>bundled_products	捆绑产品关系	是	[array]	
data>>bundled_products>>productId	子产品ID	是	[int]	
data>>bundled_products>>sku	子产品SKU	是	[string]	
data>>bundled_products>>bundledQty	捆绑数量	是	[int]	
data>>bundled_products>>cost_ratio	费用比例	是	[int]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "83594BE0-1486-5614-5097-514722893FAF",
    "response_time": "2024-07-29 15:50:36",
    "data": [
        {
            "id": 30226,
            "sku": "111sdASD",
            "product_name": "11111111SSSS",
            "cg_price": "307.0500",
            "status_text": "在售",
            "bundled_products": [
                {
                    "productId": 30164,
                    "sku": "dgd",
                    "bundledQty": 1
                },
                {
                    "productId": 30165,
                    "sku": "110098",
                    "bundledQty": 1
                },
                {
                    "productId": 30166,
                    "sku": "CS9-12",
                    "bundledQty": 1
                },
                {
                    "productId": 30168,
                    "sku": "AK912",
                    "bundledQty": 1
                },
                {
                    "productId": 30169,
                    "sku": "danpin11111111111",
                    "bundledQty": 1
                },
                {
                    "productId": 30170,
                    "sku": "danpin222222222",
                    "bundledQty": 1
                },
                {
                    "productId": 30172,
                    "sku": "danpin3333333",
                    "bundledQty": 1
                },
                {
                    "productId": 30173,
                    "sku": "zuheppin123",
                    "bundledQty": 1
                },
                {
                    "productId": 30174,
                    "sku": "21111111111111111111111111111111111111111111111111",
                    "bundledQty": 1
                },
                {
                    "productId": 30177,
                    "sku": "101-3",
                    "bundledQty": 1
                },
                {
                    "productId": 30179,
                    "sku": "9999uh",
                    "bundledQty": 1
                },
                {
                    "productId": 30180,
                    "sku": "100110111",
                    "bundledQty": 1
                },
                {
                    "productId": 30181,
                    "sku": "SKU2023091401",
                    "bundledQty": 1
                },
                {
                    "productId": 30182,
                    "sku": "20230914",
                    "bundledQty": 1
                },
                {
                    "productId": 30183,
                    "sku": "720",
                    "bundledQty": 1
                },
                {
                    "productId": 30185,
                    "sku": "tutu",
                    "bundledQty": 1
                },
                {
                    "productId": 30187,
                    "sku": "920",
                    "bundledQty": 1
                },
                {
                    "productId": 30188,
                    "sku": "2342342sdfs",
                    "bundledQty": 1
                },
                {
                    "productId": 30189,
                    "sku": "SKU998874444",
                    "bundledQty": 1
                },
                {
                    "productId": 30190,
                    "sku": "111231312312",
                    "bundledQty": 1
                },
                {
                    "productId": 30191,
                    "sku": "JJJJJJJ",
                    "bundledQty": 1
                },
                {
                    "productId": 30192,
                    "sku": "SKU12345678777",
                    "bundledQty": 1
                },
                {
                    "productId": 30193,
                    "sku": "111aaa",
                    "bundledQty": 1
                },
                {
                    "productId": 30194,
                    "sku": "1111aaa-20pcs",
                    "bundledQty": 1
                },
                {
                    "productId": 30196,
                    "sku": "wwwww",
                    "bundledQty": 1
                },
                {
                    "productId": 30197,
                    "sku": "ASDDFFGHHJJKK",
                    "bundledQty": 1
                },
                {
                    "productId": 30199,
                    "sku": "simple",
                    "bundledQty": 1
                },
                {
                    "productId": 30200,
                    "sku": "mmmmmm",
                    "bundledQty": 1
                },
                {
                    "productId": 30201,
                    "sku": "PDACESHI",
                    "bundledQty": 1
                },
                {
                    "productId": 30202,
                    "sku": "zjcp",
                    "bundledQty": 1
                },
                {
                    "productId": 30203,
                    "sku": "0925test",
                    "bundledQty": 1
                },
                {
                    "productId": 30204,
                    "sku": "cai12",
                    "bundledQty": 1
                },
                {
                    "productId": 30205,
                    "sku": "asdfasdfsadfsadasdfasdfsadf",
                    "bundledQty": 1
                },
                {
                    "productId": 30206,
                    "sku": "adfsdafdfsadfasdf",
                    "bundledQty": 1
                },
                {
                    "productId": 30207,
                    "sku": "fybx",
                    "bundledQty": 1
                },
                {
                    "productId": 30208,
                    "sku": "HYC01",
                    "bundledQty": 1
                },
                {
                    "productId": 30209,
                    "sku": "HYC02",
                    "bundledQty": 1
                },
                {
                    "productId": 30210,
                    "sku": "HYC",
                    "bundledQty": 1
                },
                {
                    "productId": 30213,
                    "sku": "sku963554",
                    "bundledQty": 1
                },
                {
                    "productId": 30214,
                    "sku": "cceshiquanxian",
                    "bundledQty": 1
                },
                {
                    "productId": 30215,
                    "sku": "wdw-222-",
                    "bundledQty": 1
                },
                {
                    "productId": 30216,
                    "sku": "shangyilvse",
                    "bundledQty": 1
                },
                {
                    "productId": 30218,
                    "sku": "DMS-1580TBV",
                    "bundledQty": 1
                },
                {
                    "productId": 30219,
                    "sku": "111444",
                    "bundledQty": 1
                },
                {
                    "productId": 30220,
                    "sku": "rtyuiop",
                    "bundledQty": 1
                },
                {
                    "productId": 30221,
                    "sku": "A2023102101",
                    "bundledQty": 1
                }
            ]
        },
        {
            "id": 31238,
            "sku": "A2024040903",
            "product_name": "A2024040903",
            "cg_price": "0.0000",
            "status_text": "在售",
            "bundled_products": [
                {
                    "productId": 31236,
                    "sku": "A2024040901",
                    "bundledQty": 1
                },
                {
                    "productId": 31237,
                    "sku": "A2024040902",
                    "bundledQty": 1
                }
            ]
        }
    ],
    "total": 2
}
复制
错误
复制成功
 上一章节
添加/编辑多属性产品
下一章节 
添加/编辑捆绑产品