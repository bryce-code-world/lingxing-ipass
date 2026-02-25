SP广告活动

注：小时数据来源于Amazon Marketing Stream，只提供授权后的；因亚马逊会修正数据，小时和天数据可能会存在差异。

接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/spCampaigns	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	1.【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始
2.推荐使用X-API-VERSION:2	[int]	2
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	12
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
state	状态：【不传默认为所有】
enabled
paused
archived	否	[string]	enabled
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15	否	[int]	15
next_token	分页游标，上次分页结果中的next_token
(第一次分页无需填写，当next_token 和 offset同时存在时以next_token为主	否	[string]	"MTAx"
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	nd8492d4-7a15-11ed-b162-0242ac1c0004
response_time	响应时间	是	[string]	2022-11-03 14:09:15
total	总数	是	[int]	1
next_token	分页游标，填入下次请求中的next_token	是	[string]	"ODAwMDAwMDAwMDAwMDAyNDE3"
data	响应数据	是	[array]	
data>>campaign_id	广告活动id	是	[number]	1711935301143
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
data>>name	广告活动名称	是	[string]	PVC Pouch-Label (A4/20P)_auto_20210810
data>>campaign_type	广告活动类型	是	[string]	sponsoredProducts
data>>targeting_type	投放类型	是	[string]	auto
data>>premium_bid_adjustment	溢价报价调整	是	[number]	1
data>>daily_budget	每日预算	是	[number]	10
data>>start_date	起始日期	是	[string]	2021-08-09 00:00:00
data>>end_date	结束日期	是	[string]	2021-08-09 00:00:00
data>>creation_date	创建日期	是	[number]	1628567183000
data>>last_updated_date	最后更新时间	是	[number]	1632985461000
data>>state	状态	是	[string]	enabled
data>>serving_status	服务状态	是	[string]	CAMPAIGN_STATUS_ENABLED
data>>bidding	竞价策略【json格式】	是	[string]	{"strategy": "legacyForSales", "adjustments":[{"predicate":"placementProductPage","percentage":44}]}
data>>portfolio_id	广告组合id	是	[number]	935301143
附加说明

bidding 竞价策略 json 内容字段说明如下：

strategy：竞价策略
legacyForSales：动态竞价 – 降低 (Dynamic Bidding - Down Only)；
autoForSales：动态竞价 – 提高和降低 (Dynamic Bidding - Up and Down)；
manual：固定竞价 (Fixed Bidding)；

adjustments：竞价调整
adjustments>>predicate 根据展示位置调整竞价，枚举值：placementTop、placementProductPage；
adjustments>>percentage 调整竞价的百分比；

 上一章节
广告组合
下一章节 
SP广告组