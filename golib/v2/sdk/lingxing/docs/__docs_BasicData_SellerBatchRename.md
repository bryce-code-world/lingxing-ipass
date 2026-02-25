批量修改店铺名称
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/seller/batchEditSellerName	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sid_name_list	批量修改店铺数组，最多可批量修改10个	是	[array]	
sid_name_list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
sid_name_list>>name	店铺名称	是	[string]	店铺1
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
data	返回数据	是	[object]	
data>>success_num	成功个数	是	[int]	19
data>>failure_num	失败个数	是	[int]	1
data>>failure_detail	失败详情	是	[array]	
data>>failure_detail>>sid	店铺id	是	[string]	10
data>>failure_detail>>name	店铺名	是	[string]	店铺10
data>>failure_detail>>error	失败原因	是	[string]	店铺记录不存在
request_id	请求id	是	[string]	752F0A0A-F113-1192-52E9-C04627835AC2
response_time	响应时间	是	[string]	2022-12-26 15:18:26
totoal	总记录数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C84C3ECA-BC04-7A8E-7211-DBE9C0A91C64",
    "response_time": "2022-02-28 20:40:12",
    "data": {
        "sid": 4,
        "name": "店铺4-JP"
    },
    "total": 1
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": [
        "错误：店铺名称不能为空"
    ],
    "request_id": "CDBBA4F7-9D62-1A48-B71C-FE75A93AD561",
    "response_time": "2022-03-01 16:07:42",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询亚马逊概念店铺列表
基础数据
下一章节 
查询ERP用户信息列表
基础数据