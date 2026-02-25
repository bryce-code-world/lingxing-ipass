创建UPC编码
接口信息
API Path	请求协议	请求方式	令牌桶容量
/listing/publish/api/upc/addCommodityCode	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
commodity_codes	编码-最多支持两百个	是	[array]	["1234567890123","123124342114"]
code_type	编码类型：支持UPC、EAN、ISBN	是	[string]	UPC
请求示例
{
    "commodity_codes": ["1234567890123","123124342114"],
    "code_type": "UPC"
}
复制
错误
复制成功

创建UPC编码

返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
msg	提示信息	是	[string]	成功
data	此次请求的code编码参数	是	[array]	[88888899,56565560]
request_id	请求链路id	是	[string]	118ed13e-77b7-11ed-a6af-0242ac1c0004
返回成功示例
{
    "code": 0,
    "msg": "成功",
    "data": [
        88888892,
        56565561
    ],
    "request_id": "bd8492d4-7a15-11ed-b162-0242ac1c0004"
}
复制
错误
复制成功
返回失败示例
{
    "code": 11004,
    "msg": "code_string已存在，请勿重复添加",
    "data": null,
    "request_id": "b528898e-7a14-11ed-8a7d-0242ac1c0004"
}
复制
错误
复制成功
 上一章节
查询补货限制列表
下一章节 
查询UPC编码列表