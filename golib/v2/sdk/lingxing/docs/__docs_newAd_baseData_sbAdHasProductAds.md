SB广告创意
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/hsaProductAds	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
sid	店铺id	是	[int]	9013
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
offset	分页偏移量，默认0	否	[int]	1
length	分页长度，默认15	否	[int]	1
next_token	分页游标，上次分页结果中的next_token
(第一次分页无需填写，当next_token 和 offset同时存在时以next_token为主	否	[string]	"MTAx"
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	6bb694e1-3d25-4821-8db8-d55dc903f6ba
response_time	响应时间	是	[string]	2023-02-17 09:59:19
total	总数	是	[int]	1
next_token	分页游标，填入下次请求中的next_token	是	[string]	"ODAwMDAwMDAwMDAwMDAyNDE3"
data	响应数据	是	[array]	
data>>campaign_id	广告活动id	是	[number]	800000000000137200
data>>ad_group_id	广告组id	是	[number]	800000000000137700
data>>ad_creative_id	广告创意id（亚马逊）	是	[number]	448868552891455
data>>creative_id	广告创意id（已作废）	是	[number]	26
data>>name	广告创意名称	是	[string]	demo sb ads 2023-02-27 7db7
data>>state	广告创意状态：
ENABLED 启用
PAUSED 暂停
ARCHIVED 归档	是	[string]	enabled
data>>creation_date	创建时间	是	[number]	1704643200000
data>>last_updated_date	最后一次更新时间	是	[number]	1704643200000
data>>profile_id	亚马逊店铺数字id	是	[number]	9000000000000013
data>>serving_status	服务状态：
AD_POLICING_SUSPENDED
INELIGIBLE
REJECTED
CAMPAIGN_OUT_OF_BUDGET
CAMPAIGN_ARCHIVED
ADVERTISER_PAYMENT_FAILURE
CAMPAIGN_PAUSED
AD_STATUS_LIVE
PORTFOLIO_OUT_OF_BUDGET
AD_GROUP_PAUSED
PORTFOLIO_ENDED
AD_PAUSED	是	[string]	AD_STATUS_LIVE
data>>asin	广告创意基础数据中ASIN字段	是	[array]	["B09MT9BKGH","B0BB389BKQ","B09MYZ614S"]
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "data": [
        {
            "campaign_id": 800000000000137259,
            "ad_group_id": 800000000000137743,
            "creative_id": 26,
            "name": "demo sb ads 2023-02-27 7db7",
            "state": "enabled",
            "creation_date": 1704643200000,
            "last_updated_date": 1704643200000,
            "profile_id": 9000000000000013,
            "serving_status": "AD_STATUS_LIVE"
            "asin": [
                "B09MT9BKGH",
                "B0BB389BKQ",
                "B09MYZ614S"
            ]
        }
    ],
    "total": 120,
    "next_token": "ODAwMDAwMDAwMDAwMDAyNDE3",
    "error_details": [],
    "request_id": "6bb694e1-3d25-4821-8db8-d55dc903f6ba",
    "response_time": "2024-05-22 15:52:53"
}
复制
错误
复制成功
 上一章节
SB广告的投放
下一章节 
SB否定关键词