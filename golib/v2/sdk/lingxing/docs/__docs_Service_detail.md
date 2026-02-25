查询邮件详情

支持获取邮件内容详情

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/mail/detail	HTTPS	POST	3
请求参数
参数名	说明	必填	类型	示例
webmail_uuid	邮件唯一标识	是	[string]	1615637469510164901
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	73A0718B-AA68-E120-D28B-F7175018175D
response_time	响应时间	是	[string]	2022-04-19 15:16:18
data	响应数据	是	[object]	
data>>webmail_uuid	邮件唯一标识	是	[string]	1615637469510164901
data>>subject	邮件标题	是	[string]	标题
data>>from_name	发件人姓名	是	[string]	system
data>>from_address	发件人地址	是	[string]	system@qq.com
data>>to_address_all	所有收件人地址	是	[string]	testtest@qq.com;test2test2@qq.com;
data>>date	日期	是	[string]	2020-11-29 17:52:28
data>>cc	抄送	是	[string]	testtest@qq.com;test2test2@qq.com;
data>>bcc	密送地址	是	[string]	testtest@qq.com;test2test2@qq.com;
data>>text_html	邮件内容	是	[string]	

Failed to send email test


data>>text_plain	纯文本的邮件内容	是	[string]	
data>>attachments	附件	是	[array]	
data>>attachments>>name	附件名称	是	[string]	F42FD424@DF41B544.txt
data>>attachments>>size	附件大小（b）	是	[number]	97078
data>>type	邮件类型
0、QA
1、买家邮件
2、亚马逊邮件
3、站外邮件	是	[string]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7F1131D2-BB65-F77A-E29D-8C712D6EAB64",
    "response_time": "2022-03-31 17:18:07",
    "data": {
        "webmail_uuid": "1615637469510164901",
        "subject": "标题",
        "from_name": "system",
        "from_address": "system@qq.com",
        "to_address_all": "test<test@qq.com>;test2<test2@qq.com>;",
        "date": "2020-11-29 17:52:28",
        "cc": "test<test@qq.com>;test2<test2@qq.com>;",
        "bcc": "test<test@qq.com>;test2<test2@qq.com>;",
        "text_html": "<p><em>Failed to send email  test</em></p>\r\n\r\n",
        "text_plain": "",
        "attachments": [
            {
                "name": "F42FD424@DF41B544.txt",
                "size": 97078
            },
            {
                "name": "91C213FF@15C47C48.txt",
                "size": 287677
            }
        ]
        type: ""
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 102,
    "message": "参数不合法",
    "error_details": ["webmail_uuid => email不能为空"],
    "request_id": "8671A9D1-BC61-9398-FBD9-C90A672DAD11",
    "response_time": "2022-03-31 15:55:03",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询邮件列表
下一章节 
查询评价管理-Review