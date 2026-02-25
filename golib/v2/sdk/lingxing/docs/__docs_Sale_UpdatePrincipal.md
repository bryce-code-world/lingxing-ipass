批量分配Listing负责人
接口信息
API Path	请求协议	请求方式	令牌桶容量
/listing/listing/open/api/asin/updatePrincipal	HTTPS	POST	3
请求参数
参数名	说明	必填	类型	示例
sid_asin_list	asin负责人分配信息，最多支持200个	是	[array]	
sid_asin_list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	16
sid_asin_list>>asin	asin	是	[string]	B09JK94H12
sid_asin_list>>principal_name	负责人姓名，最多支持10个负责人，传空或者不传表示清空负责人	否	[array]	["小明"]
请求示例
{
    "sid_asin_list": [{
        "sid": 16,
        "asin": "B09JK94H12",
        "principal_name": ["小明"]
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	错误说明	是	[string]	success
request_id	请求id	是	[string]	69a2ddc2-f3bb-11ed-aed5-0242ac130003
response_time	请求时间	是	[string]	2023-05-16 15:29:41
error_details	错误信息	是	[array]	
error_details>>index	传参数组下标	是	[int]	1
error_details>>message	错误信息	是	[string]	sid+asin重复,取第一条分配
data	响应数据	是	[object]	
data>>total	总数	是	[int]	3
data>>success	成功数	是	[int]	1
data>>error	失败数	是	[int]	2
返回成功示例
{
    "code": 0,
    "message": "success",
    "request_id": "69a2ddc2-f3bb-11ed-aed5-0242ac130003",
    "response_time": "2023-05-16 15:29:41",
    "error_details": [
        {
            "index": 1,
            "message": "sid+asin重复,取第一条分配"
        },
        {
            "index": 2,
            "message": "sid+asin重复,取第一条分配"
        }
    ],
    "data": {
        "total": 3,
        "success": 1,
        "error": 2
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 1,
    "message": "fail",
    "request_id": "9e09d24c-f3bf-11ed-a652-0242ac130003",
    "response_time": "2023-05-16 15:59:44",
    "error_details": [
        "sid_asin_list.0.sid 类型错误,必须是正整数"
    ],
    "data": []
}
复制
错误
复制成功
 上一章节
查询亚马逊Listing
销售
下一章节 
批量添加/编辑Listing配对
销售