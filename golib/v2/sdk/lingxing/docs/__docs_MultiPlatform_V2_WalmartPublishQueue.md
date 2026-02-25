Walmart修改库存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/walmart/publishQueue	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	数据字典	限制	示例
queue_type	调整类型,可用值:1	否	[number]			2
data	响应数据	否	[array]			[{"variant_unique_id":"901209431056917504","adjust_value":"50"}]
data>>variant_unique_id	变体Id	否	[string]			901209431056917504
data>>adjust_value	调整后的数据	否	[string]			50
请求示例
{
    "queue_type": 2,
    "data": [
        {
            "variant_unique_id": "901209431062705152",
            "adjust_value": "333"
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	数据字典	限制	示例
code	状态码，0：成功	是	[number]			0
message	消息提示	是	[string]			success
error_details	数据校验失败时的错误详情	是	[array]			[]
request_id	请求链路id	是	[string]			f83b6dc953f243b8ac7dcc138114efe9.1749696520938
response_time	响应时间	是	[string]			2025-06-12 10:48:41
data	响应数据	是	[object]			
data>>success_num	成功条数	是	[number]			0
data>>fail_num	失败条数	是	[number]			1
data>>failure_info	失败原因	是	[array]			[{"msku":"wal-0726-111","store_name":"test18自创walmart测试店铺2号1j j j","failure_reason":"当前库存调整未完成，不支持修改"}]
data>>failure_info>>msku	MSKU	是	[string]			wal-0726-111
data>>failure_info>>store_name	店铺名	是	[string]			test18自创walmart测试店铺2号1j j j
data>>failure_info>>failure_reason	失败原因	是	[string]			当前库存调整未完成，不支持修改
total	总数	是	[number]			0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "f83b6dc953f243b8ac7dcc138114efe9.1749696520638",
    "response_time": "2025-06-12 10:48:41",
    "data": {
        "success_num": 1,
        "fail_num": 0,
        "failure_info": []
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "f83b6dc953f243b8ac7dcc138114efe9.1749696520938",
    "response_time": "2025-06-12 10:48:41",
    "data": {
        "success_num": 0,
        "fail_num": 1,
        "failure_info": [
            {
                "msku": "wal-0726-111",
                "store_name": "test18自创walmart测试店铺2号1j j j",
                "failure_reason": "当前库存调整未完成，不支持修改"
            }
        ]
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
Shein修改库存
下一章节 
批量TEMU地址解密