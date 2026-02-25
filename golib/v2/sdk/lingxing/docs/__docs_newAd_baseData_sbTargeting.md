SB广告的投放
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/sbTargeting	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始	[int]	2
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	12
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
ads_type	广告类型：
SB 返回SB广告数据
SBV 返回SBV广告数据
ALL 同时返回SB和SBV广告数据	是	[string]	ALL
targeting_type	投放类型：
keyword 返回关键词数据
producttarget 返回商品定位数据
ALL：同时返回关键词和商品定位数据	是	[string]	ALL
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000	否	[int]	1000
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
data>>campaign_id	广告活动ID	是	[number]	
data>>ad_group_id	广告组ID	是	[number]	
data>>keyword_id	关键词ID	是	[number]	
data>>keyword_text	关键词文本	是	[string]	
data>>keyword_bid	关键词竞价	是	[number]	
data>>keyword_state	关键词状态	是	[string]	enabled
data>>match_type	匹配类型	是	[string]	exact
data>>target_id	商品定位ID	是	[number]	
data>>expression	表达式	是	[string]	
data>>target_bid	商品定位竞价	是	[number]	
data>>target_state	商品定位状态	是	[string]	
data>>resolved_expression	商品定位解析表达式	是	[string]	
data>>targeting_type	投放类型	是	[string]	
data>>ads_type	广告类型	是	[string]	
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
 上一章节
SB广告组
下一章节 
SB广告创意