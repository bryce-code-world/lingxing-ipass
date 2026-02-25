操作日志（新）
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/apiLogStandard	HTTPS	POST	10
请求头
标签	必填	说明	类型	示例
X-API-VERSION	是	【兼容旧版本】不加此头部标签，offset 为分页页码，从1开始；值为 2 时，表示 offset 为分页偏移量，从0开始	[int]	2
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	16
log_source	日志来源：
all 包括ERP和亚马逊后台的操作
erp 仅ERP中调整广告的日志
amazon 亚马逊后台的日志	是	[string]	erp
sponsored_type	广告类型：
sp 返回sp操作日志
sb 返回sb操作日志
sd 返回sd操作日志	是	[string]	sp
operate_type	对象类型:
campaigns 广告活动
adGroups 广告组
productAds 广告
keywords 关键词
negativeKeywords 否定关键词
targets 商品投放
negativeTargets 否定商品投放
profiles 预算设置	是	[string]	campaigns
start_date	起始时间，格式：Y-m-d【日期间隔不能超过一个月】	是	[string]	2022-10-01
end_date	结束时间，格式：Y-m-d【日期间隔不能超过一个月】	是	[string]	2022-11-01
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认15	否	[int]	15
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	WEUHA89F-7996-81F8-8E2F-DE7534EB2AQ2
response_time	响应时间	是	[string]	2023-04-18 10:51:37
total	总数	是	[int]	8
data	响应数据	是	[array]	
data>>profile_id	亚马逊店铺数字ID	是	[number]	121923590660074
data>>sponsored_type	广告类型：sp、sb、sd	是	[string]	sp
data>>operate_type	对象类型	是	[string]	campaigns
data>>campaign_id	广告活动id	是	[number]	75412355765106
data>>campaign_name	广告活动名称	是	[number]	campaign-dk-0111
data>>ad_group_id	广告组id	是	[number]	800000000000137102
data>>object_id	广告对象id	是	[number]	75412355765106
data>>ad_group_name	广告组名称	是	[string]	AD_GROUP
data>>object_name	广告对象名称	是	[string]	campaign-dk-0111
data>>function_name	功能来源	是	[string]	WebPage
data>>change_type	操作类型：create 创建，update 更新	是	[string]	update
data>>operate_before	操作前	是	[array]	
data>>operate_before>>code	操作类型编码	是	[string]	BUDGET_AMOUNT
data>>operate_before>>value	操作对象值	是	[string]	1.01
data>>operate_after	操作后	是	[array]	
data>>operate_after>>code	操作类型编码	是	[string]	BUDGET_AMOUNT
data>>operate_after>>value	操作对象值	是	[string]	1.02
data>>user_id	用户id	是	[int]	3071
data>>user_name	用户名	是	[string]	谢xx
data>>operate_time	操作时间 - 站点时间	是	[string]	2023-04-18 09:56:34
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [],
    "request_id": "worweoeiruxlkjsrwUEWRk",
    "response_time": "2023-04-18 10:51:37",
    "total": 1
    "data": [
        {
            "profile_id": 121923590660074,
            "sponsored_type": "sp",
            "operate_type": "campaigns",
            "campaign_id": 75412355765106,
            "campaign_name": "campaign-1118",
            "ad_group_id": null,
            "object_id": 75412355765106,
            "ad_group_name": null,
            "object_name": "campaign-dk-0111",
            "function_name": "WebPage",
            "change_type": "更新",
            "operate_before": [
                {
                    "code": "BUDGET_AMOUNT",
                    "value": "1.01"
                },
                {
                    "code": "PORTFOLIO_ID",
                    "value": "0"
                }
            ],
            "operate_after": [
                {
                    "code": "BUDGET_AMOUNT",
                    "value": "1.02"
                },
                {
                    "code": "PORTFOLIO_ID",
                    "value": "--"
                }
            ],
            "user_id": 3071,
            "user_name": "谢xx",
            "operate_time": "2023-04-18 09:56:34"
        }
    ]
}
复制
错误
复制成功
 上一章节
ABA搜索词报告-按周维度
下一章节 
查询费用类型列表