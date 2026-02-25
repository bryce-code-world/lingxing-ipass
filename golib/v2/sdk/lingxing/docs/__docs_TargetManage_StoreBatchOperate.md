店铺维度-批量新增/更新目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/store/batchOperate	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份(只允许去年、今年、明年)	是	[int]	2022
currencyCode	币种【仅支持USD、EUR、GBP、CNY、JPY、原币种】	是	[string]	原币种
operateType	操作类型：
1 仅新增
2 新增并更新	是	[int]	1
goalList	目标列表	是	[array]	
goalList>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
goalList>>amount1	1月目标	是	[number]	953
goalList>>amount2	2月目标	是	[number]	153
goalList>>amount3	3月目标	是	[number]	726
goalList>>amount4	4月目标	是	[number]	100
goalList>>amount5	5月目标	是	[number]	100
goalList>>amount6	6月目标	是	[number]	100
goalList>>amount7	7月目标	是	[number]	100
goalList>>amount8	8月目标	是	[number]	100
goalList>>amount9	9月目标	是	[number]	100
goalList>>amount10	10月目标	是	[number]	100
goalList>>amount11	11月目标	是	[number]	100
goalList>>amount12	12月目标	是	[number]	100
请求示例
{
    "assessYear": 2022,
    "currencyCode": "原币种",
    "operateType": 1,
    "goalList": [
        {
            "sid": 1,
            "amount1": 953,
            "amount2": 153,
            "amount3": 726,
            "amount4": 442,
            "amount5": 633,
            "amount6": 439,
            "amount7": 567,
            "amount8": 384,
            "amount9": 455,
            "amount10": 589,
            "amount11": 313,
            "amount12": 385
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
data>>sid	店铺id	是	[number]	1
data>>reason	失败原因【返回操作失败对应数据的失败原因】	是	[string]	店铺目标已存在
traceId	请求链路id	是	[string]	4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111
返回成功示例
{
    "code": 1,
    "msg": "操作成功",
    "data": [
        {
            "sid": 1,
            "reason": "店铺目标已存在"
        }
    ],
    "traceId": "4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111",
    "success": true
}
复制
错误
复制成功
 上一章节
店铺维度-批量查询目标
下一章节 
店铺维度-批量删除目标