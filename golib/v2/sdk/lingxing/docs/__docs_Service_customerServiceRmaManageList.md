查询RMA管理
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/rmaManage/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id，支持多选，数组	是	[array]	[114,115]
searchTimeFiled	搜索时间类型：1创建时间 2.操作时间 createTime operationTime	是	[string]	operationTime
startTime	创建或完成时间（开始），精确到年月日时 分秒，无默认	是	[string]	2024-04-10
endTime	创建或完成时间（开始），精确到年月日时 分秒，无默认	是	[string]	2024-04-10
searchValue	搜索值，msku和asin支持多个搜索，数组	是	[array]	["PEE-618"]
searchField	搜索字段：msku，asin，sku	是	[string]	msku
sortColumn	排序字段	是	[string]	operationTime
sortType	排序方式	是	[string]	desc
pageNum	页码	是	[number]	1
pageSize	每页数量	是	[number]	20
请求示例
{
    "sid": [114, 115],
    "searchTimeFiled": "operationTime",
    "startTime": "2024-04-10",
    "endTime": "2024-04-10",
    "searchValue": ["PEE-618"],
    "searchField": "msku",
    "sortColumn": "operationTime",
    "sortType": "desc",
    "pageNum": 1,
    "pageSize": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	3754c992a9c6436b9e9eb788881963b8.1752197082022
response_time	响应时间	是	[string]	2025-07-11 09:24:42
data		是	[object]	
data>>total	总数	是	[number]	207
data>>size		是	[number]	20
data>>pageCount		是	[number]	11
data>>current		是	[number]	1
data>>currentSize		是	[number]	20
data>>hasNextPage		是	[boolean]	true
data>>hasPreviousPage		是	[boolean]	false
data>>records	返回结果数据	是	[array]	[]
data>>records>>id	id	是	[string]	229
data>>records>>createTime	创建时间	是	[string]	2024-04-10 16:32:15
data>>records>>rmaNo	rma编号	是	[string]	RMA0000229
data>>records>>sid	店铺id	是	[number]	114
data>>records>>creator	创建人	是	[string]	自动化测试专用
data>>records>>amazonOrderId	订单号	是	[string]	
data>>records>>asin	asin	是	[string]	B08D7KMRT5
data>>records>>sellerSku	msku	是	[string]	PEE-618
data>>records>>itemName	商品描述	是	[string]	Fit for BMW 3 Series E90 E91 E92 E93 F30 F34 320i 325i 328i 330i 335i 320d 325d 2005-2019 Custom Luxury Waterproof Floor mats(Fit for BMW 3 Series Tourer 2013-2016,Black/red)
data>>records>>sku	本地sku	是	[string]	
data>>records>>localName	本地品名	是	[string]	
data>>records>>sellerName	店铺名称	是	[string]	8speed-CA
data>>records>>country	国家	是	[string]	加拿大
data>>records>>channelSourceName	渠道来源名称	是	[string]	站内信
data>>records>>channelSource	渠道来源id	是	[number]	2
data>>records>>afterSaleTypeName	售后类型名称	是	[string]	336
data>>records>>afterSaleType	售后类型id	是	[string]	136269578119356928
data>>records>>afterSaleCount	售后数量	是	[number]	1
data>>records>>processWayName	处理方式名称	是	[string]	
data>>records>>processWay	处理方式id	是	[string]	
data>>records>>buyerName	买家名	是	[string]	
data>>records>>buyerEmail	买家邮箱	是	[string]	
data>>records>>purchaseDateLocal	订购时间	是	[null]	null
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "3754c992a9c6436b9e9eb788881963b8.1752197082022",
    "response_time": "2025-07-11 09:24:42",
    "data": {
        "total": 207,
        "size": 20,
        "pageCount": 11,
        "current": 1,
        "currentSize": 20,
        "hasNextPage": true,
        "hasPreviousPage": false,
        "records": [{
            "id": "229",
            "createTime": "2024-04-10 16:32:15",
            "rmaNo": "RMA0000229",
            "sid": 114,
            "creator": "自动化测试专用",
            "amazonOrderId": "",
            "asin": "B08D7KMRT5",
            "sellerSku": "PEE-618",
            "itemName": "Fit for BMW 3 Series E90 E91 E92 E93 F30 F34 320i 325i 328i 330i 335i 320d 325d 2005-2019 Custom Luxury Waterproof Floor mats(Fit for BMW 3 Series Tourer 2013-2016,Black/red)",
            "sku": "",
            "localName": "",
            "sellerName": "8speed-CA",
            "country": "加拿大",
            "channelSourceName": "站内信",
            "channelSource": 2,
            "afterSaleTypeName": "336",
            "afterSaleType": "136269578119356928",
            "afterSaleCount": 1,
            "processWayName": "",
            "processWay": "",
            "buyerName": "",
            "buyerEmail": "",
            "purchaseDateLocal": null,
            "remark": "1哈哈哈\n2修改\n3再次u噶\n4\n5\n6\n7\n8\n9\n10\n11\n12\n13\n14\n15\n16\n17\n18\n19\n20\n21\n22\n23\n24\n25\n26\n27\n28\n29\n30",
            "operationTime": "2024-04-10 16:38:06",
            "rmaMasterUuid": "136430225554780672",
            "createUid": 10327479,
            "addEntrySource": 1,
            "tags": [{
                "tagName": "",
                "globalTagId": "",
                "color": ""
            }]
        }]
    },
    "total": 207
}
复制
错误
复制成功
 上一章节
查询买家之声列表
下一章节 
查询客户列表（新）