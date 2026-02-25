删除暂存货件
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/deleteCargoStorage	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
id	WFS货件id，查询WFS货件列表 接口对应字段【id】	是	[string]	200498
请求示例
{
    "id": "200498"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>is_success	是否删除成功：
true 删除成功
false 删除失败	是	[boolean]	true
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "842d7733ef5d4d968bc6877e06c34453.1697004896752",
    "response_time": "2023-10-11 14:14:56",
    "data": {
        "is_success": true
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
WFS货件暂存
下一章节 
查询利润报表-MSKU(旧版，将于04.30下线)