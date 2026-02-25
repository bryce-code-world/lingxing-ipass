查询预警消息列表-商品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/settings/warningMessage/goodsList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	否	[int]	0
length	分页长度，默认50，上限200	否	[int]	50
model_id_list	预警模型：
1 Listing调价预警
2 FBA费变更预警
3 Listing下架预警
6 FBA费异常预警
7 折扣异常预警
18 业务指标预警
20 折扣叠加预警
21 buybox丢失预警
26 父ASIN变更预警	否	[array]	[1,2]
sids	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[array]	[1,136]
start_date	开始日期【提醒时间】，闭区间，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2024-06-05
end_date	结束日期【提醒时间】，闭区间，格式：Y-m-d，时间间隔最长不超过90天	是	[string]	2024-08-05
search_field	搜索类型：
rule_name 规则名称
asin ASIN
msku MSKU	否	[string]	asin
search_value	搜索值	否	[string]	B0CWS8MNW1
show_status	处理状态：
0 待处理
1 全部	是	[int]	1
请求示例
{
    "offset": 0,
    "length": 50,
    "model_id_list": [1,2],
    "sids": [1,136],
    "start_date": "2024-06-05",
    "end_date": "2024-08-05",
    "search_field": "asin",
    "search_value": "B0CWS8MNW1",
    "show_status": 1
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
data>>message_id	用户消息id	是	[string]	906327600703779312
data>>image_url	图片地址	是	[string]	https://m.xxx.com/images/I/71y9Dp8v.jpg
data>>asin	ASIN【父ASIN变更预警时，该值是父ASIN】	是	[string]	B09MT9BKGH
data>>asin_url	ASIN地址	是	[string]	https://www.xxx.com/dp/B09MT9BxxxKGH
data>>msku_list	mksu列表	是	[array]	
data>>msku_list>>msku	MSKU	是	[string]	Pink_Head_Rope_1
data>>title	标题	是	[string]	Love Pearl Bottom Hair
data>>sku_list	sku列表	是	[array]	
data>>sku_list>>local_sku	本地产品SKU	是	[string]	HO
data>>sku_list>>local_name	品名	是	[string]	HO
data>>country	国家名称	是	[string]	美国
data>>model_id	模型id	是	[string]	1
data>>model_name	预警模型	是	[string]	listing调价
data>>rule_id	规则id	是	[string]	906204786149642240
data>>rule_name	规则名称	是	[string]	Listing 调价预警BLT0001-Green
data>>metric	监控指标	是	[string]	monitor_listing_price
data>>notify_way_str	通知方式说明	是	[string]	站内消息,强提醒
data>>notify_time	提醒时间	是	[string]	2023-06-25 16:50:41
data>>receive_uid	接收人id	是	[string]	1
data>>receiver	接收人名称	是	[string]	jack
data>>handle_status	处理状态：
0 待处理
1 已处理	是	[string]	0
data>>handle_status_str	处理状态说明	是	[string]	待处理
data>>read_status	阅读状态：
0 未读
1 已读	是	[string]	0
data>>read_status_str	阅读状态说明	是	[string]	未读
data>>monitor_time	预警时间	是	[string]	2023-08-11
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "eb112b2f45434b7d82d1b8c768d604ad.1699523180992",
    "response_time": "2023-11-09 17:46:22",
    "data": [
        {
            "message_id": "906327600703779312",
            "image_url": "https://m.xxx.com/images/I/71y9Dp8v.jpg",
            "asin": "B09MT9BKGH",
            "asin_url": "https://www.xxx.com/dp/B09MT9BxxxKGH",
            "msku_list": [
                {
                    "": "Pink_Head_Rope_1"
                }
            ],
            "title": "Love Pearl Bottom Hair",
            "sku_list": [
                {
                    "local_sku": "HO",
                    "local_name": "HO"
                }
            ],
            "country": "美国",
            "model_id": "1",
            "model_name": "listing调价",
            "rule_id": "906204786149642240",
            "rule_name": "Listing 调价预警BLT0001-Green",
            "metric": "monitor_listing_price",
            "notify_way_str": "站内消息,强提醒",
            "notify_time": "2023-06-25 16:50:41",
            "receive_uid": "1",
            "receiver": "jack",
            "handle_status": "0",
            "handle_status_str": "待处理",
            "read_status": "0",
            "read_status_str": "未读",
            "monitor_time": null
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询竞品监控列表
下一章节 
查询预警消息列表-库存