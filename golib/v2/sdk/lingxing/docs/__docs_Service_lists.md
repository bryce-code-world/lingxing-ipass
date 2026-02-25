查询邮件列表

支持查询【邮件消息】列表

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/mail/lists	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
flag	类型：
sent 发件
receive 收件	是	[string]	sent
email	店铺绑定邮箱	是	[string]	xxx@qq.com
start_date	开始日期，格式：yyyy-mm-dd	是	[string]	2024-10-30
end_date	开始日期，格式：yyyy-mm-dd	是	[string]	2024-10-30
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
{
    "flag": "sent",
    "email": "xxx@qq.com",
    "offset": 0,
    "length": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	73A0718B-AA68-E120-D28B-F7175018175D
response_time	响应时间	是	[string]	2022-04-19 15:16:18
data	响应数据	是	[array]	
data>>webmail_uuid	邮件唯一标识	是	[string]	118155928087207936
data>>date	日期	是	[string]	2022-02-25 14:33:00
data>>subject	邮件标题	是	[string]	标题1
data>>from_name	发件人姓名	是	[string]	system
data>>from_address	发件人地址	是	[string]	system@tapd.cn
data>>to_name	接收人	是	[string]	auto-communication
data>>to_address	接收人地址	是	[string]	auto-communication@qq.com
data>>has_attachment	是否存在附件：
1 存在
0 不存在	是	[int]	1
total	总数	是	[int]	187
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "27BB5B29-31C0-FA02-6A88-7134DED5A8C6",
    "response_time": "2022-03-31 15:49:35",
    "data": [
        {
            "webmail_uuid": "118155928087207936",
            "date": "2022-02-25 14:33:00",
            "subject": "标题1",
            "from_name": "system",
            "from_address": "system@tapd.cn",
            "to_name": "auto-communication",
            "to_address": "auto-communication@qq.com",
            "has_attachment": 1
        },
        {
            "webmail_uuid": "118155927439331328",
            "date": "2022-02-24 22:30:22",
            "subject": "标题2",
            "from_name": "system",
            "from_address": "system@tapd.cn",
            "to_name": "auto-communication",
            "to_address": "auto-communication@qq.com",
            "has_attachment": 0
        }
    ],
    "total": 187
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": [
        "email => email不能为空"
    ],
    "request_id": "8671A9D1-BC61-9398-FBD9-C90A672DAD11",
    "response_time": "2022-03-31 15:55:03",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
应收报告-详情-基础信息
下一章节 
查询邮件详情