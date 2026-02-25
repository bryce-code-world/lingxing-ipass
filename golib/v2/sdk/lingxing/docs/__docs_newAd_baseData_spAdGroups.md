SP广告组
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/spAdGroups	HTTPS	POST	10
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
data>>campaign_id	广告活动ID	是	[number]	57765830151100
data>>ad_group_id	广告组ID	是	[number]	800000000000080800
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
data>>name	广告组名称	是	[string]	test0331
data>>default_bid	默认竞价	是	[number]	0.02
data>>state	状态	是	[string]	enabled
data>>creation_date	创建时间	是	[number]	1623588941000
data>>last_updated_date	最后更新时间	是	[number]	1626506998000
data>>serving_status	服务状态	是	[string]	CAMPAIGN_PAUSED
data>>bid_optimization	优化建议	是	[string]	clicks
 上一章节
SP广告活动
下一章节 
SP广告商品