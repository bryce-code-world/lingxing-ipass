修改B2B价格
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/b2bPrice/modifyPrice	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
content	B2B售价	是	[array]	
content>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	1
content>>msku	MSKU	是	[string]	MSKU_XSXS_001
content>>asin	ASIN	是	[string]	B085254D
content>>b2b_price	B2B价格	是	[string]	1.00
请求示例
{
  "content": [
    {
      "sid": 31,
      "msku": "7K-YYWO-O4GB",
      "asin": "B09S9XCP1T",
      "b2b_price": "1.00"
    }
  ]
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
data>>success_num	成功数量	是	[int]	0
data>>failure_num	失败数量	是	[int]	1
data>>failure_detail_list	修改失败数据	是	[array]	
data>>failure_detail_list>>sid	店铺id	是	[int]	107
data>>failure_detail_list>>msku	MSKU	是	[string]	MSKU_XSXS_001
data>>failure_detail_list>>asin	ASIN	是	[string]	B085254D
data>>failure_detail_list>>msg	失败原因	是	[string]	请求重复
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "cdd69cf62be740158933e36dbf615c9d.1698390066049",
    "response_time": "2023-10-27 15:01:07",
    "data": {
        "success_num": 1,
        "failure_num": null,
        "failure_detail_list": null
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询Listing操作日志列表
销售
下一章节 
Listing新增商品标签
销售