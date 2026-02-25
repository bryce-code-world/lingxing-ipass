批量TEMU地址解密
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/temu/temuAddressDecrypt	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
decryptSnList	系统单号数组	是	[array]	["103548778881066177","103548778881069102"]
请求示例
{
    "decryptSnList": [
        "103548778881451018",
        "103548778881471106"
    ]
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8f17ddf49d2342f5b26fd616fade6541.117.17431331675082831
response_time	响应时间	是	[string]	2025-03-28 11:39:27
data	响应数据主体	是	[object]	
data>>successList	成功单号列表(不存在的单号默认解密成功)	是	[array]	["103552277325680282", "103552234678332007"]
data>>failureList	失败列表	是	[array]	
data>>failureList>>globalOrderNo	失败单号	否	[string]	
data>>failureList>>msg	失败原因	否	[string]	
total	成功项总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "error_details": [],
    "request_id": "65b44051a1a842bfbefc095b811bfd8c.167.17436462497300033",
    "response_time": "2025-04-03 10:10:52",
    "data": {
        "successList": [
            "103548778881471106"
        ],
        "failureList": [
            {
                "globalOrderNo": "103548778881451018",
                "msg": "解密失败：店铺发货Token已过期或被刷新，请在店铺后台生成新的token，在ERP中更新店铺授权"
            }
        ]
    },
    "total": 1
}
复制
错误
复制成功
 上一章节
Walmart修改库存
下一章节 
平台仓发货单分配库存