查询多平台店铺信息

支持查询多平台店铺基础信息，其中store_id为多平台店铺唯一

接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/mp/shop/v2/getSellerList	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	否	[int]	0
length	分页长度，上限200	否	[int]	200
platform_code	平台code：
10001 AMAZON
10002 Shopify
10003 eBay
10004 Wish
10005 AliExpress
10006 Shopee
10007 Lazada
10008 Walmart
10009 自定义平台
10010 Wayfair
10011 TikTok
10012 MERCADO
10013 CDISCOUNT
10014 NEWEGG
10015 RAKUTEN
10016 SHOPLINE
10017 TEAPPLIX
10018 SHOPLAZZA
10019 UEESHOP
10020 COUPANG
10021 SHEIN
10022 Temu全托管
10024 Temu半托管
10025 OTTO
10026 OZON
10027 SHEIN全托管
10028 SHEIN半托管
10029 AliExpress半托管
10030 AliExpress全托管
10033 Qoo10
10034 Mirakl
10038 line shopping	否	[array]	[10008,10011]
is_sync	店铺同步状态：
1 启用
0 停用	否	[int]	1
status	店铺授权状态：
1 正常授权
0 授权失败	否	[int]	1
请求示例
{
    "offset": 0,
    "length": 200,
    "platform_code": [10008,10011],
    "is_sync": 1,
    "status": 1
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	操作成功
request_id	请求链路id	是	[string]	fa58d84e-c843-4616-9fff-9b4065964465.1721704107513
response_time	响应时间	是	[string]	2024-07-23 11:08:27
data	响应数据	是	[array]	
data>>total	总数	是	[int]	1
data>>list	店铺数据	是	[array]	
data>>list>>currency	店铺币种	是	[string]	CNY
data>>list>>platform_code	平台code	是	[int]	10008
data>>list>>platform_name	平台名称	是	[string]	Walmart
data>>list>>sid	店铺id【对应获取亚马逊店铺接口sid】	是	[string]	1519
data>>list>>store_id	店铺ID	是	[string]	1108413237731xxxx
data>>list>>store_name	店铺名称	是	[string]	Walmart测试店铺
data>>list>>is_sync	店铺同步状态：
1 启用
0 停用	是	[int]	1
data>>list>>status	店铺授权状态：
1 正常授权
0 授权失败	是	[int]	0
返回请求示例
{
    "code": 0,
    "data": {
        "total": "1",
        "list": [
            {
                "store_id": "1108413237731xxxx",
                "sid": "",
                "store_name": "Walmart测试店铺",
                "platform_code": "10008",
                "platform_name": "Walmart",
                "currency": "USD",
                "is_sync": 1,
                "status": 0
            },
        ]
    },
    "response_time": "2024-07-23 11:08:27",
    "message": "操作成功",
    "request_id": "fa58d84e-c843-4616-9fff-9b4065964465.1721704107513"
}
复制
错误
复制成功
 上一章节
组织维度-批量删除目标
下一章节 
查询订单管理订单列表