广告组合
接口信息
API Path	请求协议	请求方式	令牌桶容量
/pb/openapi/newad/portfolios	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	101
profile_id	VC广告店铺profile_id，对应查询广告账号列表接口对应字段【profile_id】，sid跟profile_id其中一个必填	是	[int]	123456
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
request_id	请求链路id	是	[string]	f02f8062-36c0-45f4-9aa8-32d295389b36
response_time	响应时间	是	[string]	2023-07-20 16:57:17
total	总数	是	[int]	109
next_token	分页游标，填入下次请求中的next_token	是	[string]	"ODAwMDAwMDAwMDAwMDAyNDE3"
data	响应数据	是	[array]	
data>>portfolio_id	广告组合id	是	[number]	64803903424665
data>>profile_id	亚马逊店铺数字id	是	[number]	121923590660074
data>>name	名称	是	[string]	广告组合01
data>>budget	广告组合预算信息【json格式】：
null 无预算上限
amout 预算金额
policy 预算策略【dateRange、monthlyRecurring】
startDate 应用预算开始日期
endDate 应用预算的结束日期	是	[string]	{"amount": 5.0, "policy": "dateRange", "startDate": "20221120", "currencyCode": "USD"}
data>>in_budget	当前广告组合是否在预算范围内：
0 超出预算
1 在范围内	是	[number]	1
data>>state	状态	是	[string]	enabled
data>>creation_date	创建时间	是	[number]	1639644442056
data>>last_updated_date	最后更新时间	是	[number]	1668998717687
data>>serving_status	服务状态	是	[string]	PORTFOLIO_STATUS_ENABLED
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [],
    "request_id": "f02f8062-36c0-45f4-9aa8-32d295389b36",
    "response_time": "2023-07-20 16:57:17",
    "total": 1,
    "next_token": "ODAwMDAwMDAwMDAwMDAyNDE3"
    "data": [
        {
            "portfolio_id": 251907479208076,
            "profile_id": 121923590660074,
            "name": "广告组合01",
            "budget": "{\"amount\": 3.0, \"policy\": \"dateRange\", \"endDate\": \"20230617\", \"startDate\": \"20230611\", \"currencyCode\": \"USD\"}",
            "in_budget": 1,
            "state": "enabled",
            "creation_date": 1635327388175,
            "last_updated_date": 1686541822051,
            "serving_status": "PORTFOLIO_ENDED"            
        }
    ]
}
复制
错误
复制成功
 上一章节
SB分摊
下一章节 
SP广告活动