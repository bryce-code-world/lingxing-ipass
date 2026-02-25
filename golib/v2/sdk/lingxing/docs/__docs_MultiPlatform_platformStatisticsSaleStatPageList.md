查询销量统计列表（04.30下线）

该接口预计04.30日下线，现已迁移升级提供V2版本-查询销量统计列表v2支持查询全平台销量

API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformStatistics/saleStat/pageList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
start_date	开始日期【下单时间】，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2021-09-07
end_date	结束日期【下单时间】，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2021-09-23
data_type	统计数据维度：
1 单体
2 父体
3 MSKU
4 SKU
5 SPU	是	[string]	4
result_type	汇总类型：
1 销量
2 订单量
3 销售额	是	[string]	1
date_unit	统计时间指标：
1 年
2 月
3 周
4 日	是	[string]	2
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
{
    "start_date": "2023-06-07",
    "end_date": "2023-09-04",
    "data_type": "4",
    "result_type": "1",
    "date_unit": "2",
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
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>pic_url	图片地址	是	[string]	http://www.xxx.com/d/xx.jpg
data>>sku	SKU	是	[array]	[ "wjc9_sku", "wjc5_sku","wjc8_sku" ]
data>>spu	SPU	是	[array]	["wjc1_spu"]
data>>spu_name	SPU名称	是	[array]	["wjc1_kmk"]
data>>msku	MSKU	是	[array]	["shop_wjc9_sku2WTVtRW9"]
data>>variant_id	单体id	是	[array]	["145244552278"]
data>>varian_title	单体名称	是	[array]	["branco222"]
data>>brand_id	品牌id	是	[array]	["8"]
data>>cid	分类id	是	[array]	["9"]
data>>product_id	本地产品id	是	[array]	[]
data>>product_name	品名	是	[array]	["wjc9_pm"]
data>>develop_uid	开发人id	是	[array]	["10062027"]
data>>develop_name	开发人名称	是	[array]	["jack"]
data>>platform_code	平台编码	是	[array]	["10006"]
data>>platform_name	平台名称	是	[array]	["Shopee"]
data>>site_code	站点编码	是	[array]	["10006-SG"]
data>>site_name	站点名称	是	[array]	["新加坡"]
data>>store_id	店铺id	是	[array]	["110000000018006002"]
data>>store_name	店铺名称	是	[array]	["自创Shopee店铺5号"]
data>>country_code	国家编码	是	[string]	
data>>currency_code	币种	是	[string]	USD
data>>icon	币种符号	是	[string]	$
data>>date_collect	数据明细	是	[string]	{"2021-11-30": 0.0,"2021-11-29": 0.0}
data>>volume_total	明细小计	是	[string]	9
data>>platform_product_id	平台商品id	是	[array]	["2023010798"]
data>>platform_product_title	标题	是	[array]	["Collapsible Storage Boxes for Closet"]
data>>statistics_list_child	与父级结构一致	是	[array]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "b171ce9140f64903983b5dcb012b1b18.1697765325548",
    "response_time": "2023-10-20 09:28:45",
    "data": [
        {
            "pic_url": "",
            "sku": [
                "wjc10_sku",
                "wjc6_sku",
                "wjc11_sku"
            ],
            "spu": [
                "wjc0_spu"
            ],
            "spu_name": [
                "wjc0_km"
            ],
            "msku": [
                "shop_wjc10_sku2NHjm9uO",
                "shop_wjc6_sku2Qq8nqPa",
                "shop_wjc11_sku2XmnexYs"
            ],
            "variant_id": [
                "145244552278"
            ],
            "varian_title": [
                "branco222"
            ],
            "brand_id": [
                "9",
                "5",
                "10"
            ],
            "cid": [
                "10",
                "6",
                "11"
            ],
            "product_id": [],
            "product_name": [
                "wjc10_pm",
                "wjc6_pm",
                "wjc11_pm"
            ],
            "develop_uid": [
                "10062027"
            ],
            "develop_name": [
                "jack"
            ],
            "platform_code": [
                "10006"
            ],
            "platform_name": [
                "Shopee"
            ],
            "site_code": [
                "10006-SG"
            ],
            "site_name": [
                "新加坡"
            ],
            "store_id": [
                "110000000018006002"
            ],
            "store_name": [
                "自创Shopee店铺2号"
            ],
            "country_code": null,
            "currency_code": "USD",
            "icon": "$",
            "date_collect": "{\"2021-11\":9.0}",
            "volume_total": "9",
            "platform_product_id": [
                "2023010784",
                "2023010759",
                "2023010774"
            ],
            "platform_product_title": [
                "biaoti222"
            ],
            "statistics_list_child": [
                {
                    "pic_url": "https://xxx.com/10775.png",
                    "sku": [
                        "wjc11_sku"
                    ],
                    "spu": [
                        "wjc0_spu"
                    ],
                    "spu_name": [
                        "wjc0_km"
                    ],
                    "msku": [
                        "shop_wjc11_sku2XmnexYs"
                    ],
                    "variant_id": [
                        "145244552278"
                    ],
                    "varian_title": [
                        "branco222"
                    ],
                    "brand_id": [
                        "10"
                    ],
                    "cid": [
                        "11"
                    ],
                    "product_id": [],
                    "product_name": [
                        "wjc11_pm"
                    ],
                    "develop_uid": [
                        "10062027"
                    ],
                    "develop_name": [
                        "jack"
                    ],
                    "platform_code": [
                        "10006"
                    ],
                    "platform_name": [
                        "Shopee"
                    ],
                    "site_code": [
                        "10006-SG"
                    ],
                    "site_name": [
                        "新加坡"
                    ],
                    "store_id": [
                        "110000000018006002"
                    ],
                    "store_name": [
                        "自创Shopee店铺2号"
                    ],
                    "country_code": null,
                    "currency_code": "USD",
                    "icon": "$",
                    "date_collect": "{\"2021-11\":3.0}",
                    "volume_total": "3",
                    "platform_product_id": [
                        "2023010774"
                    ],
                    "platform_product_title": [
                        "Collapsible Storage Boxes for Closet"
                    ],
                    "statistics_list_child": null
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
获取快速出库结果
下一章节 
查询销量统计列表V2