查询TikTok在线商品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/tiktok/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
brandIds	品牌id列表	否	[array]	[1,2,3,4,5,6]
categoryIds	分类id列表	否	[array]	[5,4,3,2,1]
offset	分页偏移量	否	[int]	0
length	分页长度，上限1000	否	[int]	20
pairingStatus	配对状态	否	[int]	1
searchField	搜索维度
1、标题
2、品名
5、平台SPU
7、MSKU ID
8、SKU
9、MSKU
10、SPU货号	否	[string]	1
platformStatus	状态
DRAFT
PENDING
FAILED
ACTIVATE
SELLER_DEACTIVATED
PLATFORM_DEACTIVATED
FREEZE
DELETED	否	[array]	["ACTIVATE"]
storeIds	店铺id列表	否	[array]	["110458505618526720","110470479434055680"]
searchSingleValue	搜索值	否	[string]	BEL
searchValues	搜索值列表	否	[array]	["1","2"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/multiplatform/tiktok/list?access_token=value&sign=value&timestamp=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "brandIds": [value],
    "categoryIds": [value],
    "offset": value,
    "length": value,
    "pairingStatus": value,
    "searchField": "value",
    "platformStatus": ["value"],
    "storeIds": ["value"],
    "searchSingleValue": "value",
    "searchValues": ["value"],
    "req_time_sequence": "value"
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
request_id	请求链路id	是	[string]	
response_time	响应时间	是	[string]	2024-09-12 09:18:59
data	响应数据	是	[object]	
data>>count	总数	是	[int]	2
data>>list		是	[array]	
data>>list>>id	id	是	[string]	16
data>>list>>productUniqueId	商品uid	是	[string]	901467036615189504
data>>list>>principalUids	负责人uid列表	是	[array]	
data>>list>>principalName	负责人名称	是	[String]	
data>>list>>companyId	企业id	是	[string]	901157321452879872
data>>list>>storeId	店铺id	是	[string]	110458505618526720
data>>list>>storeName	店铺名称	是	[string]	美国平台物流
data>>list>>platformCode	平台代码	是	[string]	10011
data>>list>>mskuId	MSKU ID	是	[string]	1729498834797957654
data>>list>>msku	MSKU	是	[string]	SUTOWL-1
data>>list>>pname	品名	是	[string]	wjc13_pm
data>>list>>sku	SKU	是	[string]	wjc13_sku
data>>list>>pid	本地商品id	是	[string]	5763
data>>list>>title	本地商品id	是	[string]	
data>>list>>cid	分类id	是	[string]	13
data>>list>>bid	品牌id	是	[string]	9
data>>list>>notes	备注	是	[String]	
data>>list>>imgUrl	图片链接（原质量）	是	[string]	
data>>list>>spu	SPU	是	[String]	
data>>list>>spuId	SPU ID	是	[string]	1729498831843791382
data>>list>>platformStatus	状态	是	[string]	ACTIVATE
data>>list>>categoryName	分类名称	是	[string]	Bathroom Tumblers
data>>list>>usableInventory	可用库存	是	[int]	684
data>>list>>waitOutboundQuantity	待出库库存	是	[int]	0
data>>list>>siteCode	站点代码	是	[String]	
data>>list>>taxExclusivePrice	税前价格	是	[double]	20
data>>list>>salesPrice	售价	是	[string]	20.00
data>>list>>currency	货币单位	是	[string]	$
data>>list>>bname	品牌名称	是	[string]	品牌9
data>>list>>cname	叶子类目	是	[string]	分类13
total		是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.142.17261283521902653",
    "response_time": "2024-09-12 16:05:56",
    "data": {
        "count": 1,
        "list": [
            {
                "id": "6",
                "productUniqueId": "9014578083787
                    "10329319"
                ],
                "principalName": "丁",
                "companyId": "901157321452879872",
                "storeId": "110427040906481664",
                "storeName": "tiktok沙盒店铺",
                "platformCode": "10011",
                "mskuId": "10002",
                "msku": "test1-sku-name2",
                "pname": "ueeshop1111",
                "sku": "ueeshop1111",
                "pid": "5793",
                "title": "",
                "cid": "8",
                "bid": "8",
                "notes": "",
                "imgUrl": "",
                "spu": "172959296971062701",
                "spuId": "1729592969712207008",
                "platformStatus": "SELLERDEACTIVATED",
                "categoryName": "test1-Botol & Stoples Penyimpanan",
                "usableInventory": 101,
                "waitOutboundQuantity": 10,
                "siteCode": "US,UK",
                "taxExclusivePrice": null,
                "salesPrice": "117.50",
                "currency": "$",
                "bname": "品牌8",
                "cname": "分类8"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询Shein在线商品
下一章节 
查询Temu在线商品