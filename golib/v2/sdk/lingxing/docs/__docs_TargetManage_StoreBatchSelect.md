店铺维度-批量查询目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/store/batchSelect	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份	是	[string]	2024
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，1 成功	是	[int]	1
msg	提示消息	是	[string]	操作成功
data	响应数据	是	[array]	
data>>goalName	目标名	是	[string]	店铺1 2024 月度销售目标
data>>sid	店铺id	是	[number]	1
data>>name	店铺名	是	[string]	店铺1
data>>currencyCode	币种	是	[string]	USD
data>>icon	币种符号	是	[string]	$
data>>assessYear	目标年份	是	[int]	2022
data>>goalAmount1	1月目标值	是	[number]	669.00
data>>goalAmount2	2月目标值	是	[number]	958.25
data>>goalAmount3	3月目标值	是	[number]	958.25
data>>goalAmount4	4月目标值	是	[number]	100.25
data>>goalAmount5	5月目标值	是	[number]	100.25
data>>goalAmount6	6月目标值	是	[number]	100.25
data>>goalAmount7	7月目标值	是	[number]	100.25
data>>goalAmount8	8月目标值	是	[number]	100.25
data>>goalAmount9	9月目标值	是	[number]	100.25
data>>goalAmount10	10月目标值	是	[number]	100.25
data>>goalAmount11	11月目标值	是	[number]	100.25
data>>goalAmount12	12月目标值	是	[number]	100.25
data>>realAmount1	1月完成值	是	[number]	58.25
data>>realAmount2	2月完成值	是	[number]	58.25
data>>realAmount3	3月完成值	是	[number]	58.25
data>>realAmount4	4月完成值	是	[number]	58.25
data>>realAmount5	5月完成值	是	[number]	58.25
data>>realAmount6	6月完成值	是	[number]	58.25
data>>realAmount7	7月完成值	是	[number]	58.25
data>>realAmount8	8月完成值	是	[number]	58.25
data>>realAmount9	9月完成值	是	[number]	58.25
data>>realAmount10	10月完成值	是	[number]	58.25
data>>realAmount11	11月完成值	是	[number]	58.25
data>>realAmount12	12月完成值	是	[number]	58.25
data>>completeRateAmount1	1月完成率	是	[number]	5.13
data>>completeRateAmount2	2月完成率	是	[number]	8.13
data>>completeRateAmount3	3月完成率	是	[number]	8.13
data>>completeRateAmount4	4月完成率	是	[number]	58.13
data>>completeRateAmount5	5月完成率	是	[number]	58.13
data>>completeRateAmount6	6月完成率	是	[number]	58.13
data>>completeRateAmount7	7月完成率	是	[number]	58.13
data>>completeRateAmount8	8月完成率	是	[number]	58.13
data>>completeRateAmount9	9月完成率	是	[number]	58.13
data>>completeRateAmount10	10月完成率	是	[number]	58.13
data>>completeRateAmount11	11月完成率	是	[number]	58.13
data>>completeRateAmount12	12月完成率	是	[number]	58.13
data>>totalGoalAmount	累计目标值	是	[number]	3523.25
data>>totalRealAmount	累计完成值	是	[number]	763.23
data>>totalCompleteRate	累计完成率	是	[number]	20.34
data>>createUserId	创建用户id	是	[number]	10001
data>>createUserName	创建用户名	是	[string]	张三
data>>updateUserId	更新用户id	是	[number]	10002
data>>updateUserName	更新用户名	是	[string]	李四
data>>gmtCreate	创建时间	是	[string]	2022-08-01 10:00:00
data>>gmtModified	更新时间	是	[string]	2022-08-01 10:00:00
traceId	请求链路id	是	[string]	4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111
返回成功示例
{
  "code": 1,
  "msg": "请求成功",
  "data": [
    {
      "goalName": "店铺1 2024 月度销售目标",
      "zid": 1,
      "sid": 1,
      "name": "店铺1",
      "companyId": 90001,
      "currencyCode": "USD",
      "icon": "$",
      "assessType": 1,
      "assessObject": 1,
      "assessYear": 2022,
      "goalAmount1": 669,
      "goalAmount2": 958,
      "goalAmount3": 403,
      "goalAmount4": 460,
      "goalAmount5": 410,
      "goalAmount6": 359,
      "goalAmount7": 640,
      "goalAmount8": 587,
      "goalAmount9": 468,
      "goalAmount10": 77,
      "goalAmount11": 19,
      "goalAmount12": 645,
      "realAmount1": 900,
      "realAmount2": 416,
      "realAmount3": 794,
      "realAmount4": 714,
      "realAmount5": 326,
      "realAmount6": 38,
      "realAmount7": 364,
      "realAmount8": 524,
      "realAmount9": 804,
      "realAmount10": 815,
      "realAmount11": 487,
      "realAmount12": 61,
      "completeRateAmount1": 2.23,
      "completeRateAmount2": 2.23,
      "completeRateAmount3": 12.23,
      "completeRateAmount4": 2.23,
      "completeRateAmount5": 16.23,
      "completeRateAmount6": 2.33,
      "completeRateAmount7": 2.33,
      "completeRateAmount8": 22.23,
      "completeRateAmount9": 62.26,
      "completeRateAmount10": 2.23,
      "completeRateAmount11": 24.23,
      "completeRateAmount12": 2.23,
      "totalGoalAmount": 786,
      "totalRealAmount": 37,
      "totalCompleteRate": 2.23,
      "createUserId": 1001,
      "createUserName": "user1",
      "updateUserId": 1001,
      "updateUserName": "user1",
      "isDeleted": 0,
      "gmtCreate": "2022-08-01 10:00:00",
      "gmtModified": "2022-08-01 10:00:00"
    }
  ],
  "traceId": "cf6d9586d8074aec92af0f1a85d18d9c.1670328945764",
  "success": true
}
复制
错误
复制成功
 上一章节
查询运营日志(新)
下一章节 
店铺维度-批量新增/更新目标