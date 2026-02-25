查询Shein在线商品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/shein/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
brandIds	品牌ID列表	否	[array]	[1,2,3,4,5,6,7,8,9,10,11,12,13,16]
categoryIds	分类ID列表	否	[array]	[0,18,15,16,14,13,12,11,10,9,8,7,6,5,4,3,2,1]
offset	偏移量	否	[int]	0
length	分页长度，上限1000	否	[int]	20
pairingStatus	配对状态
0、未配对
1、已配对	否	[int]	1
searchField	搜索字段
1、标题
2、品名
3、SPU货号
4、SKC货号
5、平台SPU
6、平台SKC
7、MSKU ID
8、SKU
9、MSKU	否	[string]	1
status	状态
0、删除
1、在售
2、停售	否	[int]	1
storeIds	店铺ID列表	否	[array]	[110461728895488000","110470632053572096"]
searchSingleValue	单一值搜索	否	[string]	白色
searchValues	精确搜索值列表	否	[array]	["1","2"]
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/multiplatform/shein/list?access_token=value&sign=value&timestamp=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "brandIds": [value],
    "categoryIds": [value],
    "offset": value,
    "length": value,
    "pairingStatus": value,
    "searchField": "value",
    "status": value,
    "storeIds": ["value"],
    "searchSingleValue": "value",
    "searchValues": ["value"]
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
request_id	请求链路id	是	[string]	a0d54debf93140f3b58d1ed81e8e3583.209.17259613819113307
response_time	响应时间	是	[string]	2024-09-10 17:43:03
data	响应数据	是	[object]	
data>>count	总数	是	[int]	18
data>>list		是	[array]	
data>>list>>id	id	是	[string]	136
data>>list>>productUniqueId	商品uid	是	[string]	901457819926044672
data>>list>>principalUids	负责人uid列表	是	[array]	["100003","10329149"]
data>>list>>principalName	负责人名称	是	[string]	
data>>list>>companyId	企业id	是	[string]	901157321452879872
data>>list>>storeId	店铺id	是	[string]	110000000018027003
data>>list>>storeName	店铺名称	是	[string]	SHINE测试店铺-发货单
data>>list>>platformCode	平台代码	是	[string]	10027
data>>list>>mskuId	MSKU ID	是	[string]	test7-I82lggftsddy
data>>list>>msku	MSKU	是	[string]	test7-PH-30-BAI
data>>list>>pname	品名	是	[string]	ueeshop1111
data>>list>>sku	SKU	是	[string]	ueeshop1111
data>>list>>pid	本地商品id	是	[string]	5793
data>>list>>title	标题	是	[string]	test1-1件白色family家庭相框组合连体挂墙照片墙经典款婚妙影1
data>>list>>cid	分类id	是	[string]	8
data>>list>>bid	品牌id	是	[string]	8
data>>list>>day7SaleCnt	7天销量	是	[int]	5
data>>list>>day30SaleCnt	30天销量	是	[int]	20
data>>list>>wareHouseDataList		是	[array]	
data>>list>>wareHouseDataList>>warehouseId	仓库id	是	[string]	PS1618418400
data>>list>>wareHouseDataList>>warehouseName	仓库名称	是	[string]	美国1
data>>list>>wareHouseDataList>>quantity	库存数量	是	[int]	8587
data>>list>>notes	备注	是	[string]	
data>>list>>imgSmallUrl	图片链接（低质量）	是	[string]	
data>>list>>imgMediumUrl	图片链接（中质量）	是	[string]	
data>>list>>imgUrl	图片链接（原质量）	是	[string]	
data>>list>>skcId	SKC ID	是	[string]	sh2312290829294899
data>>list>>skc	SKC	是	[string]	PH-30白色家庭相框组合1
data>>list>>spuId	SPU ID	是	[string]	test-h23122908292
data>>list>>spu	SPU	是	[string]	test-PH-30黑色家庭相框组合
data>>list>>status	状态	是	[int]	1
data>>list>>site	站点	是	[string]	shein-www,shein-fr,shein-es
data>>list>>siteList		是	[array]	
data>>list>>siteList>>site	站点名称	是	[string]	SHEIN阿根廷站
data>>list>>siteList>>isOnSale	是否在售	是	[boolean]	false
data>>list>>attribute	属性	是	[string]	test-白色-test-均码
data>>list>>categoryName	大分类名称	是	[string]	装饰框和相框
data>>list>>salePrice	售价	是	[double]	
data>>list>>saleCurrency	货币	是	[string]	
data>>list>>inventoryQuantity	总库存	是	[int]	50
data>>list>>lockedQuantity	出库中库存	是	[int]	20
data>>list>>tempLockQuantity	待支付库存	是	[int]	10
data>>list>>usableInventory	可用库存	是	[int]	5
data>>list>>bname	品牌名称	是	[string]	品牌8
data>>list>>cname	细分类名称	是	[string]	分类8
data>>list>>supplyPriceList		是	[array]	
data>>list>>supplyPriceList>>price	供货价格	是	[double]	17.8
data>>list>>supplyPriceList>>currency	货币	是	[string]	¥
data>>list>>supplyPriceList>>beginStr	供货价格执行开始时间	是	[string]	2024-06-27
data>>list>>supplyPriceList>>endStr	供货价格执行结束时间	是	[String]	2024-07-27
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.191.17261286463909883",
    "response_time": "2024-09-12 16:10:47",
    "data": {
        "count": 1,
        "list": [
            {
                "id": "145",
                "productUniqueId": "901458096084689408",
                "principalUids": [
                    "1"
                ],
                "principalName": "1",
                "companyId": "901157321452879872",
                "storeId": "110000000018021002",
                "storeName": "shein测试店铺211-11",
                "platformCode": "10021",
                "mskuId": "test16-I82lggftsddy",
                "msku": "shein_wjc11_skut2n2l",
                "pname": "wjc11_pm",
                "sku": "wjc11_sku",
                "pid": "5761",
                "title": "家庭相框组合连体挂墙照片墙经典款婚妙影1",
                "cid": "12",
                "bid": "10",
                "day7SaleCnt": 0,
                "day30SaleCnt": 0,
                "wareHouseDataList": [],
                "notes": null,
                "imgSmallUrl": "",
                "imgMediumUrl": "",
                "imgUrl": "",
                "skcId": "sh2312290829294899",
                "skc": "PH-30白色家庭相框组合1",
                "spuId": "test-h23122908292",
                "spu": "test-PH-30黑色家庭相框组合",
                "status": 1,
                "site": "shein-www,shein-fr,shein-es",
                "siteList": null,
                "attribute": "test-白色-test-均码",
                "categoryName": "装饰框和相框",
                "salePrice": 19.80,
                "saleCurrency": "¥",
                "inventoryQuantity": 50,
                "lockedQuantity": 20,
                "tempLockQuantity": 10,
                "usableInventory": 5,
                "bname": "品牌10",
                "cname": "分类12",
                "supplyPriceList": []
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询Shopify在线商品-子体数据
下一章节 
查询TikTok在线商品