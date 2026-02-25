查询业绩通知列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/customerService/performanceNotice/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid	店铺id	是	[number]	121
status	处理状态：0（无），1（待处理），2（已处理），3（无需处理）	否	[array]	[0,1]
startDate	开始时间 YYYY-MM-DD	否	[string]	2025-09-11
endDate	结束时间 YYYY-MM-DD	否	[string]	2025-09-10
searchField	搜索字段,subject 邮件主题,content 邮件内容	否	[string]	subject
searchValue	搜索值	否	[string]	123
mailTagIds	邮件标签 id	否	[array]	["288325866709888888"]
isRead	是否已读，-1 全部，0 未读，1 已读	否	[number]	1
offset	偏移量	否	[number]	0
length	分页长度	否	[number]	50
请求示例
{
    "sid": 121,
    "status": [0, 1],
    "startDate": "2025-09-11",
    "endDate": "2025-09-10",
    "searchField": "subject",
    "searchValue": "123",
    "mailTagIds": ["2883251111888888"],
    "isRead": 1,
    "offset": 0,
    "length": 50
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0成功	是	[number]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	93093b23e6c64089b7ba9cb233dc11e5.1757576342640
response_time	响应时间	是	[string]	2025-09-11 15:39:03
data		是	[object]	
data>>list		是	[array]	
data>>list>>id	主键ID	是	[number]	19
data>>list>>companyId	企业ID	是	[number]	901113104767778800
data>>list>>performanceNoticeUuid	唯一标识ID	是	[string]	119551640807743488
data>>list>>sid	店铺ID	是	[number]	121
data>>list>>mailId	亚马逊邮件唯一标识	是	[string]	4d9e345b-c9f5-4639-be54-3bb26d14a12116988
data>>list>>subject	主题	是	[string]	第1封 sid：12
data>>list>>status	处理状态：0（无），1（待处理），2（已处理），3（无需处理）	是	[number]	0
data>>list>>isRead	是否已读，-1 全部，0 未读，1 已读	是	[number]	1
data>>list>>content	内容	是	[string]	
data>>list>>mailCreateDate	亚马逊邮件创建时间	是	[string]	2024-12-29
data>>list>>tagList	邮件标签信息列表	是	[array]	
data>>list>>tagList>>cid	分类ID	是	[number]	0
data>>list>>tagList>>color	标签颜色	是	[string]	A972CF
data>>list>>tagList>>tagName	标签名称	是	[string]	紫色
data>>list>>tagList>>webMailTagUuid	标签ID	是	[string]	123165465498498
data>>total	总数	是	[number]	1
data>>lastUpdateDate	最后更新时间	是	[string]	2025-03-19 18:32:11
total	总数	是	[number]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "93093b23e6c64089b7ba9cb233dc11e5.1757576342640",
    "response_time": "2025-09-11 15:39:03",
    "data": {
        "list": [{
            "id": 19,
            "companyId": 800,
            "performanceNoticeUuid": "488",
            "sid": 121,
            "mailId": "116988",
            "subject": "第1封 sid：12",
            "status": 0,
            "isRead": 1,
            "content": "",
            "mailCreateDate": "2024-12-29",
            "tagList": [{
                "cid": 0,
                "color": "A972CF",
                "tagName": "紫色",
                "webMailTagUuid": "12498"
            }]
        }],
        "total": 1,
        "lastUpdateDate": "2025-03-19 18:32:11"
    },
    "total": 1
}
复制
错误
复制成功
 上一章节
查询客户列表（新）
下一章节 
关键词列表