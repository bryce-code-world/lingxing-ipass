查询操作日志
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/product/getPagingLogLists	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
businessId	businessId，对应查询本地产品列表data>>id字段	是	[long]	1234567890
endTime	结束时间	是	[string]	示例值
startTime	开始时间	是	[string]	示例值
page	页码	否	[int]	1
size	每页大小	否	[int]	20
请求cURL示例
curl --location 'https://openapi.lingxing.com/product/getPagingLogLists?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "businessId": 1234567890,
    "endTime": "示例值",
    "startTime": "示例值",
    "page": 1,
    "size": 20
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	响应码	是	[int]	
data	返回数据	是	[array]	
data>>action	操作类型	是	[string]	示例值
data>>datetime	datetime（日期格式：yyyy-MM-dd hh:mm:ss）	是	[string]	2024-01-01 16:35:53
data>>detail	操作详情	是	[string]	示例值
data>>userId	操作人ID	是	[long]	1234567890
data>>userName	操作人	是	[string]	示例名称
error_details	错误详情	是	[array]	
message	提示信息	是	[string]	
request_id	请求id	是	[string]	
response_time	响应时间	是	[string]	
total	总记录数	是	[int]	
返回成功示例
{
  "code": 0,
  "data": [
    {
      "action": "创建商品",
      "datetime": "2024-01-01 16:35:53",
      "detail": "SKU: P-001 被用户A创建",
      "userId": 1234567890,
      "userName": "示例用户A"
    },
    {
      "action": "修改库存",
      "datetime": "2024-01-02 10:20:15",
      "detail": "SKU: P-001 库存由100修改为90",
      "userId": 1234567891,
      "userName": "示例用户B"
    }
  ],
  "error_details": [],
  "message": "请求成功",
  "request_id": "req_1234567890abcdef",
  "response_time": "2024-05-18 14:30:00",
  "total": 2
}
复制
错误
复制成功
 上一章节
删除产品标签
下一章节 
查询透明计划商品列表