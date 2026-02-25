组织维度-批量查询目标
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/goal/management/open/user/batchSelect	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
assessYear	目标年份	是	[int]	2024
assessType	考核指标：1 销售额，2 销量	是	[int]	1
请求示例
{
    "assessYear": "2024",
    "assessType": 1
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
data>>realName	用户名	是	[string]	user1
data>>uid	uid	是	[string]	10001
data>>defaultOrg	默认部门	是	[string]	技术部
data>>defaultOrgId	默认部门id	是	[string]	40023
data>>orgs	所有部门	是	[array]	
data>>orgs>>orgId	部门id	是	[string]	40023
data>>orgs>>orgName	部门名称	是	[string]	技术部
data>>currencyCode	币种	是	[string]	USD
data>>icon	货币符号	是	[string]	$
data>>assessType	考核指标：1 销售额，2 销量	是	[string]	1
data>>goalValue1	1月目标值	是	[number]	100.00
data>>goalValue2	2月目标值	是	[number]	100.00
data>>goalValue3	3月目标值	是	[number]	100.00
data>>goalValue4	4月目标值	是	[number]	100.00
data>>goalValue5	5月目标值	是	[number]	100.00
data>>goalValue6	6月目标值	是	[number]	100.00
data>>goalValue7	7月目标值	是	[number]	100.00
data>>goalValue8	8月目标值	是	[number]	100.00
data>>goalValue9	9月目标值	是	[number]	100.00
data>>goalValue10	10月目标值	是	[number]	100.00
data>>goalValue11	11月目标值	是	[number]	100.00
data>>goalValue12	12月目标值	是	[number]	100.00
data>>realValue1	1月完成值	是	[number]	60.23
data>>realValue2	2月完成值	是	[number]	60.23
data>>realValue3	3月完成值	是	[number]	60.23
data>>realValue4	4月完成值	是	[number]	60.23
data>>realValue5	5月完成值	是	[number]	60.23
data>>realValue6	6月完成值	是	[number]	60.23
data>>realValue7	7月完成值	是	[number]	60.23
data>>realValue8	8月完成值	是	[number]	60.23
data>>realValue9	9月完成值	是	[number]	60.23
data>>realValue10	10月完成值	是	[number]	60.23
data>>realValue11	11月完成值	是	[number]	60.23
data>>realValue12	12月完成值	是	[number]	60.23
data>>completeRate1	1月完成率	是	[number]	60.23
data>>completeRate2	2月完成率	是	[number]	60.23
data>>completeRate3	3月完成率	是	[number]	60.23
data>>completeRate4	4月完成率	是	[number]	60.23
data>>completeRate5	5月完成率	是	[number]	60.23
data>>completeRate6	6月完成率	是	[number]	60.23
data>>completeRate7	7月完成率	是	[number]	60.23
data>>completeRate8	8月完成率	是	[number]	60.23
data>>completeRate9	9月完成率	是	[number]	60.23
data>>completeRate10	10月完成率	是	[number]	60.23
data>>completeRate11	11月完成率	是	[number]	60.23
data>>completeRate12	12月完成率	是	[number]	60.23
data>>yearGoalValue	年度目标值	是	[number]	1200.00
data>>yearRealValue	年度完成值	是	[number]	728.23
data>>completeProcess	年度完成进度	是	[number]	60.23
data>>createUserId	目标创建用户id	是	[number]	10005
data>>createUser	目标创建用户	是	[string]	张三
data>>updateUserId	目标最后更新用户id	是	[number]	10006
data>>updateUser	目标最后更新用户	是	[string]	李四
data>>gmtCreate	创建时间	是	[string]	2022-08-01 10:00:00
data>>gmtModified	更新时间	是	[string]	2022-10-01 10:00:00
traceId	请求链路id	是	[string]	4f7d4a69b6a54f9898881c70f60a5dd9.1670328817111
返回成功示例
{
    "code": 1,
    "msg": "操作成功",
    "data": [
        {
            "zid": 1,
            "companyId": 900001,
            "realName": "user1",
            "uid": 100108,
            "defaultOrg": null,
            "defaultOrgId": null,
            "orgs": [],
            "currencyCode": "CNY",
            "icon": "￥",
            "assessType": 1,
            "goalValue1": 33.00,
            "goalValue2": 13.00,
            "goalValue3": 13.00,
            "goalValue4": 43.00,
            "goalValue5": 13.00,
            "goalValue6": 13.00,
            "goalValue7": 613.30,
            "goalValue8": 253.00,
            "goalValue9": 213.00,
            "goalValue10": 253.00,
            "goalValue11": 213.00,
            "goalValue12": 213.00,
            "realValue1": 0.00,
            "realValue2": 0.00,
            "realValue3": 0.00,
            "realValue4": 0.00,
            "realValue5": 0.00,
            "realValue6": 0.00,
            "realValue7": 140.86,
            "realValue8": 0.00,
            "realValue9": 70.43,
            "realValue10": 0.00,
            "realValue11": 0.00,
            "realValue12": 0.00,
            "completeRate1": 0.00,
            "completeRate2": 0.00,
            "completeRate3": 0.00,
            "completeRate4": 0.00,
            "completeRate5": 0.00,
            "completeRate6": 0.00,
            "completeRate7": 22.97,
            "completeRate8": 0.00,
            "completeRate9": 33.07,
            "completeRate10": 0.00,
            "completeRate11": 0.00,
            "completeRate12": 0.00,
            "yearGoalValue": 1886.30,
            "yearRealValue": 211.29,
            "completeProcess": 11.20,
            "isDeleted": 0,
            "createUserId": 10002,
            "createUser": "user2",
            "updateUserId": 10002,
            "updateUser": "user2",
            "gmtCreate": "2022-11-18 15:12:55",
            "gmtModified": "2022-12-06 20:13:40"
        }
    ],
    "traceId": "b51db5a7b7a64d029879be4fa26cd7f3.1670328998326",
    "success": true
}
复制
错误
复制成功
 上一章节
店铺维度-批量删除目标
下一章节 
组织维度-批量新增/更新目标