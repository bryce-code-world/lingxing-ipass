查询竞品监控列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/tool/competitiveMonitor/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
levels	竞品等级：
1 A
2 B
3 C
4 D	否	[array]	[1,2,3,4]
update_time_start	开始时间【更新时间】，闭区间，格式：Y-m-d	否	[string]	2024-08-05
update_time_end	结束时间【更新时间】，闭区间，格式：Y-m-d	否	[string]	2024-08-05
search_field	搜索字段：asin ASIN	否	[string]	asin
search_value	搜索值：多个使用英文逗号分隔，上限200	否	[string]	B0CWS8MNW1
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
请求示例
{
    "levels": [1,2,3,4],
    "update_time_start": "2024-08-05",
    "update_time_end": "2024-08-05",
    "search_field": "asin",
    "search_value": "B0CWS8MNW1",
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	类型	描述	示例
code	[int]	状态码，0 成功	0
message	[string]	提示信息	success
error_details	[array]	错误信息	
request_id	[string]	请求链路id	60CEFCDF-BF57-6E6E-541A-C31E83B266BE
response_time	[string]	响应时间	2023-08-16 18:18:31
total	[int]	总数	2
data	[array]	响应数据	
data>>mid	[int]	国家id	1
data>>title	[string]	商品标题	Super J Running Belt for Gym Workout Exercise Stable Sports Waist Bag Ultra Light Bounce Free
data>>asin	[string]	ASIN	B0BBB1QWER
data>>asin_url	[string]	ASIN地址	https://www.amazon.com/xxxxxxx
data>>price	[string]	商品价格	1.01
data>>currency	[string]	货币	$
data>>level_name	[string]	竞品等级名称	B
data>>category_list	[array]	竞品分类	
data>>category_list>>category_name	[string]	分类名称	分类1
data>>rating	[string]	rating	0
data>>star	[string]	评分	0.0
data>>review_num	[string]	review数量	123
data>>big_category_rank	[string]	大类排名名次	1
data>>big_category	[string]	大类排名名称	kitchen electricties
data>>init_big_category_rank	[int]	初始大类排名	0
data>>small_ranks	[array]	小类排名	
data>>small_ranks>>small_rank	[int]	小类排名	12
data>>small_ranks>>init_small_rank	[int]	初始小类排名	9
data>>small_ranks>>small_category_text	[string]	小类排名名称	xxx
data>>monitor_status	[int]	监控状态：
0 暂停监控
1 监控中	1
data>>creator_uid	[string]	监控记录创建人uid	1
data>>monitor_uids	[array]	监控用户ID	542
data>>creator	[string]	竞品记录创建人	超级管理员
data>>last_update_event	[array]	最近更新事件	
data>>search_term	[string]	Search Term	
data>>main_image	[string]	主图	https://m.media/xxxx.jpg
data>>thumbnail	[array]	商品轮播图	["https://m.media/xxxx.jpg"]
data>>featurebullets	[array]	商品要点	["专业级的厨师刀套装"]
data>>item_weight	[string]	商品重量	0.634 ounces
data>>product_dimensions	[string]	商品尺寸	2.44 x 1.3 x 0.35 inches
data>>init_fbm_seller_num	[int]	初始fbm卖家数量	0
data>>fbm_seller_num	[int]	当前fbm卖家数量	1
data>>init_fba_seller_num	[int]	初始fba卖家数量	0
data>>fba_seller_num	[int]	当前fba卖家数量	0
data>>init_buybox_price	[string]	初始buybox价格	0.00
data>>buybox_price	[string]	当前buybox价格	0.00
data>>buybox_currency	[string]	当前buybox货币符号	$
data>>buybox_usd_price	[string]	当前buybox的USD价格	0.00
data>>avg_price	[string]	平均价格	0.00
data>>avg_currency	[string]	平均价格的货币符号	$
返回示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "ede6e9708a054d45967c82bd8dec31ec.1692944366375",
    "response_time": "2023-08-25 14:19:26",
    "data": [
        {
            "mid": 1,
            "title": "NP Phone Holder",
            "asin": "B0BB3SDDE34",
            "asin_url": "https://www.amazon.com/xxxx",
            "price": "0.00",
            "currency": "",
            "level_name": "B",
            "category_list": [
                {
                    "category_name": "分类1"
                },
                {
                    "category_name": "分类2"
                }
            ],
            "rating": "0",
            "star": "0.0",
            "review_num": "0",
            "big_category_rank": "0",
            "big_category": "",
            "init_big_category_rank": 0,
            "small_ranks": [],
            "monitor_status": 1,
            "creator_uid": "1",
            "monitor_uids": [],
            "creator": "管理员",
            "last_update_event": [],
            "search_term": "",
            "main_image": "https://m.media-amazon.com/xxxxxx.jpg",
            "thumbnail": [],
            "featurebullets": [
                "PHONE HOLDER-ar_ae"
            ],
            "item_weight": "0.634 ounces",
            "product_dimensions": "2.44 x 1.3 x 0.35 inches",
            "init_fba_seller_num": 0,
            "fba_seller_num": 0,
            "init_fbm_seller_num": 0,
            "fbm_seller_num": 1,
            "init_buybox_price": "0.00",
            "buybox_price": "0.00",
            "buybox_currency": "",
            "buybox_usd_price": "0.00",
            "avg_currency": "$",
            "avg_price": "0.00"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
关键词列表
下一章节 
查询预警消息列表-商品