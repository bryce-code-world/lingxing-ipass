查询产品品牌列表

支持查询本地产品品牌列表，对应系统【产品】>【品牌管理】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/local_inventory/brand	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000，上限1000	否	[int]	100
请求示例
{
    "offset": 0,
    "length": 100
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
request_id	请求链路id	是	[string]	66A95CC1-C8E0-2CA8-13F8-2D0F3C7E86A6
response_time	响应时间	是	[string]	2021-11-11 18:14:00
data	返回数据	是	[array]	
data>>bid	品牌id	是	[int]	2
data>>title	品牌名称	是	[string]	111
data>>brand_code	品牌简码	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "66A95CC1-C8E0-2CA8-13F8-2D0F3C7E86A6",
    "response_time": "2021-11-11 18:14:00",
    "data": [
        {
            "bid": 1,
            "title": "品牌1",
            "brand_code": "code"
        },
        {
            "bid": 2,
            "title": "品牌2"
            "brand_code": "code"         
        }
    ]
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误,数据库异常",
    "error_details": [],
    "request_id": "6EBFF08B-032A-3D90-DE7A-18186924F1C2",
    "response_time": "2021-11-11 18:14:47",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
添加/编辑辅料
下一章节 
添加/编辑产品品牌