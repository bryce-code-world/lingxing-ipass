SP否定投放
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/spNegativeTargetsOrKeywords	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始	[int]	2
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	101
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
campaign_id	广告活动id	否	[number]	141775982358867
target_type	投放类型：keyword target	是	[string]	target
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15	否	[int]	15
next_token	分页游标，上次分页结果中的next_token
(第一次分页无需填写，当next_token 和 offset同时存在时以next_token为主	否	[string]	"MTAx"
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	AA8EA3F5-7996-81F8-8E2F-DE7534EB2AE0
response_time	响应时间	是	[string]	2023-05-09 17:02:10
total	总数	是	[int]	239
next_token	分页游标，填入下次请求中的next_token	是	[string]	"ODAwMDAwMDAwMDAwMDAyNDE3"
data	响应数据	是	[array]	
data>>campaign_id	广告活动id	是	[number]	1640591546
data>>ad_group_id	广告组id	是	[number]	1640591547
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
data>>negative_type	否定类型：
negativeKeyword 否定关键词
negativeAsin 否定ASIN
negativeBrand 否定品牌	是	[string]	negativeKeyword
data>>negative_text	否定内容：【按类型显示】
否定关键词：显示否定关键词的文本
否定ASIN：显示ASIN码
否定品牌：显示品牌名称	是	[string]	关键词A
data>>negative_match_type	否定匹配方式：
否定关键词：negativeExact、negativePhrase
否定ASIN和否定品牌：返回空字符串	是	[string]	negativeExact
data>>state	状态	是	[string]	enabled
data>>creation_date	创建时间	是	[number]	1655548464
data>>last_updated_date	最后更新时间	是	[number]	1655548467
data>>serving_status	服务状态	是	[string]	TARGETING_CLAUSE_STATUS_LIVE
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [],
    "request_id": "SD7EA3F5-7996-81F8-8E2F-DE7534EB2AE1",
    "response_time": "2023-05-08 12:02:10",
    "total": 1,
    "next_token": "ODAwMDAwMDAwMDAwMDAyNDE3",
    "data": [
        {
            "campaign_id": 83080499191276,
            "ad_group_id": 214029880198148,
            "profile_id": 121923590660074,
            "negative_type": "negativeAsin",
            "negative_text": "B077PWXCYS",
            "negative_match_type": "",
            "state": "enabled",
            "creation_date": 1634815522980,
            "last_updated_date": 1634815522980,
            "serving_status": "CAMPAIGN_PAUSED"
        }
    ]
}
复制
错误
复制成功
 上一章节
SP商品定位
下一章节 
SB广告活动