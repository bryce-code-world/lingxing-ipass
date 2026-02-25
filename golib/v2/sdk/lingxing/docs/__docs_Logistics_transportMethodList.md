查询运输方式列表

支持查询自定义运输方式，对应系统【设置】>【业务配置】>【物流】当中的运输方式

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/businessConfig/transportMethod/list	HTTPS	POST	1
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>method_id	运输方式id	是	[string]	241261034153378816
data>>code	序号	是	[string]	7
data>>name	运输方式名称	是	[string]	传统联运
data>>is_system	是否为系统运输方式：
true 是
false 否	是	[boolean]	true
data>>enabled	启动状态：
0 停用
1 启用	是	[int]	0
data>>remark	备注	是	[string]	
data>>creator_id	创建人id	是	[int]	10325110
data>>creator_name	创建人名称	是	[string]	张三
data>>updater_id	最后编辑人id	是	[int]	10393433
data>>updater_name	最后编辑人名称	是	[string]	李四
data>>created_at	创建时间，格式：秒级时间戳	是	[int]	1671431440
data>>updated_at	更新时间，格式：秒级时间戳	是	[int]	1671431440
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "4a703231393a44e0b50b958e6e9f33eb.1701334080222",
    "response_time": "2023-11-30 16:48:03",
    "data": [
        {
            "method_id": "241261034153378816",
            "code": "1",
            "name": "海派",
            "is_system": true,
            "enabled": 1,
            "remark": "",
            "creator_id": 10393433,
            "creator_name": "张三",
            "updater_id": 10393433,
            "updater_name": "李四",
            "created_at": 1671431440,
            "updated_at": 1671431440
        } ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询已启用的自发货物流方式
下一章节 
SP广告活动报表