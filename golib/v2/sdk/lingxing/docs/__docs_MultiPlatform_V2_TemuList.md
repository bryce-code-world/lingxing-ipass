查询Temu在线商品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/temu/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
brandIds	品牌id列表	否	[array]	[1,2,3,4,5,6,7]
categoryIds	分类id列表	否	[array]	[10,9,8,7,6,5,4,3,2,1]
offset	分页偏移量	否	[int]	0
length	分页长度，上限1000	否	[int]	20
pairingStatus	配对状态
0、未配对
1、已配对	否	[int]	1
searchField	搜索维度
1、标题
2、品名
4、SKC货号
5、平台SPU
6、平台SKC
7、MSKU ID
8、SKU
9、MSKU	是	[string]	9
status	状态
0、删除
2、正常	否	[int]	2
storeIds	店铺id列表	否	[array]	["110479125448229376","110480765502464512"]
searchValues	精确搜索值列表	否	[array]	["8"]
searchSingleValue	模糊搜索值	否	[string]	["1","2"]
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
request_id	请求链路id	是	[string]	a0d54debf93140f3b58d1ed81e8e3583.114.17261215395432737
response_time	响应时间	是	[string]	2024-09-12 14:12:20
data	响应数据	是	[object]	
data>>count	总数	是	[int]	2
data>>list		是	[array]	
data>>list>>id	id	是	[string]	11
data>>list>>productUniqueId	商品uid	是	[string]	901457876655013888
data>>list>>principalUids	负责人uid列表	是	[array]	["100003"]
data>>list>>principalName	负责人名称	是	[string]	林小芳111
data>>list>>companyId	企业id	是	[string]	901157321452879872
data>>list>>storeId	店铺id	是	[string]	110443249963082240
data>>list>>storeName	店铺名称	是	[string]	欧区测试--1111_半7
data>>list>>platformCode	平台代码	是	[string]	10024
data>>list>>mskuId	MSKU ID	是	[string]	27162024062701
data>>list>>msku	MSKU	是	[string]	test2-msku-b8888
data>>list>>pname	品名	是	[string]	test2407094pm
data>>list>>sku	SKU	是	[string]	test2407094sku
data>>list>>pid	本地商品id	是	[string]	5858
data>>list>>title	标题	是	[string]	test-semi_goods_add_test_by_api
data>>list>>cid	分类id	是	[string]	16
data>>list>>bid	品牌id	是	[string]	1
data>>list>>day7SaleCnt	7天销量	是	[int]	0
data>>list>>day30SaleCnt	30天销量	是	[int]	0
data>>list>>wareHouseDataList		是	[array]	
data>>list>>wareHouseDataList>>warehouseId	仓库id	是	[string]	WH-03852452270231627
data>>list>>wareHouseDataList>>warehouseName	仓库名称	是	[string]	美东仓库
data>>list>>wareHouseDataList>>quantity	库存数量	是	[int]	90
data>>list>>notes	备注	是	[string]	
data>>list>>imgUrl	图片链接	是	[string]	
data>>list>>skcId	SKC ID	是	[string]	2024062701
data>>list>>skc	SKC	是	[string]	test2-skc-b1
data>>list>>spuId	SPU ID	是	[string]	252845852024062701
data>>list>>status	上架状态 0、删除 2、正常	是	[int]	2
data>>list>>attribute	SKU规格	是	[string]	test-红色-test-白色
data>>list>>jitMode	JIT发货模式
0、无发货模式
1、Jit发货模式
2、非Jit发货模式	是	[int]	2
data>>list>>catName	叶子类目	是	[int]	1
data>>list>>selectStatus	商品状态
枚举类，详情见附加说明	是	[string]	test2-男士项链
data>>list>>virtualQuantity	虚拟库存	是	[int]	150
data>>list>>usableInventory	可用库存	是	[null]	
data>>list>>bname	品牌名称	是	[string]	品牌12
data>>list>>cname	分类名称	是	[string]	篮球鞋
data>>list>>supplyPriceList		是	[array]	
data>>list>>supplyPriceList>>price	供货价格	是	[double]	77
data>>list>>supplyPriceList>>currency	货币	是	[string]	¥
data>>list>>supplyPriceList>>beginStr	供货价格执行开始时间	是	[string]	2024-08-07
data>>list>>supplyPriceList>>endStr	供货价格执行结束时间	是	[string]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.205.17261288401451103",
    "response_time": "2024-09-12 16:14:01",
    "data": {
        "count": 1,
        "list": [
            {
                "id": "11",
                "productUniqueId": "901457813888",
                "principalUids": [
                    "100003"
                ],
                "principalName": "林小",
                "companyId": "9011572879872",
                "storeId": "1104432082240",
                "storeName": "欧区测试",
                "platformCode": "10024",
                "mskuId": "2712701",
                "msku": "test2-msku-b8888",
                "pname": "test2407094pm",
                "sku": "test2407094sku",
                "pid": "5858",
                "title": "test-semi_goods_add_test_by_api",
                "cid": "16",
                "bid": "1",
                "day7SaleCnt": 0,
                "day30SaleCnt": 0,
                "wareHouseDataList": [
                    {
                        "warehouseId": "WH-03852452270231627",
                        "warehouseName": "美东仓库",
                        "quantity": 90
                    },
                    {
                        "warehouseId": "WH-03852452270231628",
                        "warehouseName": "美西仓库",
                        "quantity": 60
                    }
                ],
                "notes": "",
                "imgUrl": "",
                "skcId": "2024062701",
                "skc": "test2-skc-b1",
                "spuId": "2528458062701",
                "status": 2,
                "attribute": "test-红色-test-白色",
                "jitMode": 2,
                "catName": "test2-男士项链",
                "virtualQuantity": 150,
                "usableInventory": null,
                "bname": "品牌12",
                "cname": "篮球鞋",
                "supplyPriceList": [
                    {
                        "price": 77.00,
                        "currency": "¥",
                        "beginStr": "2024-08-07",
                        "endStr": null
                    },
                    {
                        "price": 47.00,
                        "currency": "¥",
                        "beginStr": "2024-05-30",
                        "endStr": "2024-08-06"
                    },
                    {
                        "price": 88.00,
                        "currency": "¥",
                        "beginStr": "2024-05-21",
                        "endStr": "2024-05-29"
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
附加说明：

selectStatus枚举类如下：
0, 已弃用
1, 待平台选品
2, 待上传生产资料
3, 待寄样
4, 寄样中
5, 待平台审版
6, 审版不合格
7, 平台核价中
8, 待修改生产资料
9, 核价未通过
10, 待下首单
11, 已下首单
12, 已加入站点
13, 已下架
14, 待卖家修改
15, 已修改
16, 服饰可加色
17, 已终止

 上一章节
查询TikTok在线商品
下一章节 
查询eBay在线商品列表