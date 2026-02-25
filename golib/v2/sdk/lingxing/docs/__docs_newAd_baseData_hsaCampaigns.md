SB广告活动
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/hsaCampaigns	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始	[int]	2
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	12
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
state	状态：【不传默认为所有】
enabled
paused
archived	否	[string]	archived
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15	否	[int]	15
next_token	分页游标，上次分页结果中的next_token
(第一次分页无需填写，当next_token 和 offset同时存在时以next_token为主	否	[string]	"MTAx"
返回结果

Json Object

参数名	
说明
	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	nd8492d4-7a15-11ed-b162-0242ac1c0004
response_time	响应时间	是	[string]	2022-11-03 14:09:15
total	总数	是	[int]	1
next_token	分页游标，填入下次请求中的next_token	是	[string]	"ODAwMDAwMDAwMDAwMDAyNDE3"
data	响应数据	是	[array]	
data>>campaign_id	广告活动ID	是	[number]	112910451299387
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
data>>name	广告活动名称	是	[string]	L038-Video-Test-Produkt-Blue
data>>budget	预算	是	[number]	10
data>>budget_type	预算类型	是	[string]	daily
data>>start_date	开始日期	是	[datetime]	2021-06-11 00:00:00
data>>state	状态	是	[string]	enabled
data>>serving_status	服务状态	是	[string]	running
data>>bid_optimization	自动竞价	是	[number]	
data>>creative	广告创意结构	是	[string]	{"type": "video", "asins": ["B07Q8KTP65"], "videoMediaIds": ["amzn1.adex.media1.d9cd46ae-1a24-4ffe-b29f-74d7e131e16d"]}
data>>landing_page	广告着陆页	是	[string]	{"url": "https://www.amazon.de/dp/B07Q8KTP65", "pageType": "detailPage"}
data>>bid_multiplier	自定义竞价调整	是	[number]	
data>>end_date	结束日期	是	[string]	
data>>portfolio_id	广告组合ID	是	[number]	136707579327356
data>>creative_type	创意类型	是	[string]	COLLECTION
附加说明

creative 广告创意结构 json 内容字段说明如下：

type：创意类型，值为video；
asin：asin；
videoMediaIds：视频文件资源id；

landing_page 广告着陆页 json 内容字段说明如下：

url：url；
pageType：着陆页类型；
 上一章节
SP否定投放
下一章节 
SB广告组