上传备货单装箱信息

创建“待发货/待收货/已完成”状态备货单时，支持上传备货单装箱信息。

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/packing	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
overseas_order_no	备货单号	是	[string]	
packaging_type	装箱类型：1 每箱多个sku，2 每箱一个sku	是	[int]	
box_count	总箱数	是	[int]	
box_list	装箱数据	是	[array]	
box_list>>box_no	箱号	是	[int]	
box_list>>box_count	箱数，默认1	否	[int]	
box_list>>box_nos	自定义箱号数组，箱数大于1时必传	否	[array]	
box_list>>height	箱规-高（cm）	是	[number]	
box_list>>length	箱规-长（cm）	是	[number]	
box_list>>width	箱规-宽（cm）	是	[number]	
box_list>>weight	箱子毛重（KG）	是	[number]	
box_list>>items	商品详情	是	[array]	
box_list>>items>>product_id	商品id	是	[int]	
box_list>>items>>twp_id	三方商品id，查询系统产品与第三方海外仓产品映射列表接口对应字段【twp_id】，
注：收货仓库是第三方仓时必填	否	[int]	
box_list>>items>>quantity_shipped	装箱数：计算公式=备货数量/箱数	是	[int]	
box_list>>items>>oversea_product_code	三方商品编码（收货仓库为三方海外仓，可不传product_id）	否	[int]	
box_list>>items>>match_num	配对数量【默认1】：传 1 为单个配对，其他值均为整箱配对	是	[int]	
box_list>>items>>fnsku	fnsku	否	[string]	
box_list>>items>>barcode	万邑通SN码	否	[string]	
box_list>>items>>sid	店铺id，库存中心过渡版本后有fnsku必填 ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	 
请求示例
{
  "overseas_order_no": "FBA123456789",
  "packaging_type": 1,
  "box_count": 2,
  "box_list": [
    {
      "box_no": 1,
      "height": 30.5,
      "length": 40.0,
      "width": 35.2,
      "weight": 12.8,
      "items": [
        {
          "product_id": 1001,
          "twp_id": 2001,
          "quantity_shipped": 10,
          "match_num": 1,
          "fnsku": "X000001ABC",
          "sid": 3001
        },
        {
          "product_id": 1002,
          "quantity_shipped": 5,
          "match_num": 2,
          "sid": 3002
        }
      ]
    },
    {
      "box_no": 2,
      "height": 25.0,
      "length": 35.0,
      "width": 30.0,
      "weight": 8.5,
      "items": [
        {
          "product_id": 1003,
          "quantity_shipped": 8,
          "match_num": 1,
          "fnsku": "X000002DEF",
          "sid": 3001
        }
      ]
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
message	消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	BD36FCD8-E32E-E9E1-6BEC-06DE331C95AB
response_time	响应时间	是	[string]	2021-06-15 17:49:16
data	响应数据	是	[object]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C70FD2A7-4E7D-6B32-7D72-FC42587CE97E",
    "response_time": "2022-07-29 14:33:42",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
查询备货单装箱信息
下一章节 
查询备货单收货记录