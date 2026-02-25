获取UPC编码列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/listing/publish/api/upc/upcList	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	200
请求示例
{
    "offset": 0,
    "length": 200
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
msg	状态说明	是	[string]	成功
data	响应数据	是	[array]	
data>>total	商品编码总数	是	[int]	2
data>>list	商品编码数据列表	是	[array]	
data>>list>>id	记录唯一id	是	[number]	220
data>>list>>commodity_code	商品编码	是	[string]	88888899
data>>list>>code_type	商品编码类型	是	[string]	EAN
data>>list>>is_used	商品编码使用状态：0 否 ，1 是	是	[number]	1
data>>list>>created_user_id	创建人uid	是	[number]	502
data>>list>>use_user_id	使用人uid	是	[number]	502
data>>list>>use_time	商品编码被使用的时间	是	[string]	2021-09-26 17:13:55
data>>list>>remark	备注	是	[string]	
data>>list>>gmt_create	商品编码创建时间	是	[string]	2022-12-09 19:46:10
data>>list>>is_used_desc	商品编码使用状态说明	是	[string]	已使用
request_id	请求链路id	是	[string]	550d352c-7a05-11ed-b0c7-0242ac1c0004
返回成功示例
{
    "code": 0,
    "msg": "成功",
    "data": {
        "total": 2,
        "list": [
            {
                "id": 97,
                "commodity_code": "88886666",
                "code_type": "EAN",
                "is_used": 1,
                "created_user_id": 129100,
                "use_user_id": 547,
                "use_time": "2021-09-26 17:13:55",
                "remark": "",
                "gmt_create": "2021-09-24 19:05:40",
                "is_used_desc": "已使用"
            },
            {
                "id": 221,
                "commodity_code": "56565560",
                "code_type": "EAN",
                "is_used": 0,
                "created_user_id": 502,
                "use_user_id": 0,
                "use_time": "",
                "remark": "",
                "gmt_create": "2022-12-09 19:46:10",
                "is_used_desc": "未使用"
            }
        ]
    },
    "request_id": "550d352c-7a05-11ed-b0c7-0242ac1c0004"
}
复制
错误
复制成功
返回失败示例
{
    "code": 1000,
    "msg": "未知错误",
    "data": null,
    "request_id": "b528898e-7a14-11ed-8a7d-0242ac1c0004"
}
复制
错误
复制成功
 上一章节
创建UPC编码
下一章节 
查询本地产品列表