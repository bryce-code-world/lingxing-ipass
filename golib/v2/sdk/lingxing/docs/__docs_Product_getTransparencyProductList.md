产品管理-查询透明计划商品列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/product/getTransparencyProductList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
isRelateMsku	是否关联MSKU，枚举值：1-是, 2-否	否	[int]	1
length	分页长度，默认20，最大200	否	[int]	20
offset	分页偏移量，默认0	否	[int]	0
productStatus	产品状态，枚举值：all-全部, Enrolled-已注册, In OPR-OPR中, Protected-受保护, NoStatus-无状态	否	[string]	all
searchField	搜索字段，指定搜索的字段名	否	[string]	asin
searchValue	搜索值，用于模糊搜索	否	[string]	B08XYZ1234
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/product/getTransparencyProductList?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "isRelateMsku": 1,
    "length": 20,
    "offset": 0,
    "productStatus": "all",
    "searchField": "asin",
    "searchValue": "B08XYZ1234"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	响应码	是	[int]	
data	返回数据	是	[object]	
data>>pageList	分页列表	是	[array]	
data>>pageList>>accountName	账号名称	是	[string]	店铺账号1
data>>pageList>>asin	asin	是	[string]	B08XYZ1234
data>>pageList>>brandName	品牌名称	是	[string]	品牌A
data>>pageList>>gtin	gtin	是	[string]	0123456789012
data>>pageList>>id	产品id	是	[long]	123456
data>>pageList>>labelType	标签类型	是	[string]	2D
data>>pageList>>picUrl	图片url	是	[string]	https://example.com/product.jpg
data>>pageList>>productStatus	产品状态	是	[string]	Enrolled
data>>pageList>>sellerSku	卖家sku	是	[string]	SKU-001
data>>pageList>>taId	账号id	是	[long]	789012
data>>pageList>>tcodeNotUsedTotal	未使用tcode	是	[int]	500
data>>pageList>>tcodeTotal	合计tcode	是	[int]	1000
data>>pageList>>title	标题	是	[string]	产品标题示例
data>>total	合计，总记录数	是	[int]	100
error_details	错误详情	是	[array]	
message	提示信息	是	[string]	
request_id	请求id	是	[string]	
response_time	响应时间	是	[string]	
total	总记录数	是	[int]	
返回成功示例
{
  "code": 0,
  "data": {
    "pageList": [
      {
        "accountName": "店铺账号1",
        "asin": "B08XYZ1234",
        "brandName": "品牌A",
        "gtin": "0123456789012",
        "id": 123456,
        "labelType": "2D",
        "picUrl": "https://example.com/product.jpg",
        "productStatus": "Enrolled",
        "sellerSku": "SKU-001",
        "taId": 789012,
        "tcodeNotUsedTotal": 500,
        "tcodeTotal": 1000,
        "title": "产品标题示例"
      },
      {
        "accountName": "店铺账号2",
        "asin": "B08ABC5678",
        "brandName": "品牌B",
        "gtin": "9876543210987",
        "id": 654321,
        "labelType": "3D",
        "picUrl": "https://example.com/product2.jpg",
        "productStatus": "Pending",
        "sellerSku": "SKU-002",
        "taId": 210987,
        "tcodeNotUsedTotal": 100,
        "tcodeTotal": 200,
        "title": "另一个产品标题"
      }
    ],
    "total": 2
  },
  "error_details": [],
  "message": "success",
  "request_id": "a0d54debf93140f3b58d1ed81e8e3583.178.17255922733991817",
  "response_time": "2024-11-12 16:00:00",
  "total": 2
}
复制
错误
复制成功
 上一章节
查询操作日志
下一章节 
查询供应商列表