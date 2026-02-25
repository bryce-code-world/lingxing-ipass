查询调整单确认调整异步结果
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/adjustOrder/adjust/getAdjustStatus	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
taskNo	异步任务编号	否	[string]	292558315891339776
请求curl示例
curl --location --request POST 'https://openapi.lingxing.com/basicOpen/adjustOrder/adjust/getAdjustStatus?app_key=ak_SEpKDE0QPc9Md&access_token=73b3b68e-1692-46fd-86b5-0d236cdf5a30&timestamp=1743660157&sign=%2BpopkuJ4Dcxx%2BTadNm6uoBoXvX6gecBhfzqfP4%2FY4f7LhFQTjRgX%2FS%2Bgzn9j6L85' \
--header 'Content-Type: application/json' \
--data-raw '{
    "taskNo":"292558315891339776"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[array]	
data>>success	成功数量	是	[string]	0
data>>fail	失败数量	是	[string]	2
data>>total	总单数	是	[string]	2
data>>failList	失败明细	是	[array]	
data>>failList>>orderSn	单号	是	[string]	IB220927004
data>>failList>>detail	失败原因	是	[string]	
返回结果示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "d58586799b514e83ab3a8c74dbd4af01.141.17442661336650715",
    "response_time": "2025-04-10 14:22:13",
    "data": {
        "success": 0,
        "fail": 2,
        "total": 2,
        "failList": [
            {
                "orderSn": "AD250407008",
                "msg": "提交失败：【SKU：SKUOWP',店铺：'JC12-UK',fnsku:'X001JVMWSF',仓位:'可用暂存】存在盘点中记录"
            },
            {
                "orderSn": "AD250407007",
                "msg": "提交失败：【SKU：SKUOWP',店铺：'JC12-UK',fnsku:'X001JVMWSF',仓位:'可用暂存】存在盘点中记录"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
调整单确认调整
下一章节 
删除出库单