组织维度-批量新增/更新目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/user/batchOperate	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份(只允许去年、今年、明年)	是	[int]	2022
assessType	考核指标：1 销售额，2 销量	是	[int]	1
operateType	操作类型：
1 覆盖
2 更新	是	[int]	1
currencyCode	币种【仅支持USD、EUR、GBP、CNY、JPY】	是	[string]	USD
userGoalList	用户目标集合	是	[array]	
userGoalList>>uid	用户id	是	[int]	1001
userGoalList>>value1	1月目标	是	[number]	100
userGoalList>>value2	2月目标	是	[number]	100
userGoalList>>value3	3月目标	是	[number]	100
userGoalList>>value4	4月目标	是	[number]	100
userGoalList>>value5	5月目标	是	[number]	100
userGoalList>>value6	6月目标	是	[number]	100
userGoalList>>value7	7月目标	是	[number]	100
userGoalList>>value8	8月目标	是	[number]	100
userGoalList>>value9	9月目标	是	[number]	100
userGoalList>>value10	10月目标	是	[number]	100
userGoalList>>value11	11月目标	是	[number]	100
userGoalList>>value12	12月目标	是	[number]	100
请求示例
{
    "assessYear":2021,
    "assessType": 1,
    "operateType":2,
    "currencyCode":"CNY",
    "userGoalList":[
        {
            "uid": 100109,
            "value1": 33.00,
            "value2": 13.00,
            "value3": 13.00,
            "value4": 43.00,
            "value5": 13.00,
            "value6": 13.00,
            "value7": 613.30,
            "value8": 253.00,
            "value9": 213.00,
            "value10": 253.00,
            "value11": 213.00,
            "value12": 213.00
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，1 成功	是	[int]	1
msg	返回消息	是	[string]	操作成功
data	响应数据	是	[array]	
data>>uid	用户名	是	[number]	100109
data>>reason	失败原因	是	[string]	用户不存在
traceId	请求链路id	是	[string]	4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111
返回成功示例
{
    "code": 1,
    "msg": "操作成功",
    "data": [
        {
            "uid": 100109,
            "reason": "用户目标已存在"
        }
    ],
    "traceId": "4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111",
    "success": true
}
复制
错误
复制成功
 上一章节
组织维度-批量查询目标
下一章节 
组织维度-批量删除目标