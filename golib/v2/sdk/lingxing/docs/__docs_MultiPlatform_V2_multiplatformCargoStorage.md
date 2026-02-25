WFS货件暂存
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/multiplatform/cargo/storage	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_id	店铺id	是	[string]	110000000020008001
cargo_goods_list	货件包含的商品	是	[array]	
cargo_goods_list>>box_num	箱子数量	是	[string]	1
cargo_goods_list>>single_box_num	单箱商品数量	是	[string]	1
cargo_goods_list>>declare_num	货件初始申报量	是	[string]	1
cargo_goods_list>>expected_arrival_time	预估到货时间	是	[string]	2023-09-15 00:00:00
cargo_goods_list>>msku	MSKU，查询添加商品列表接口对应的字段【msku】	是	[string]	WMUS-DLO-HairRoller-Blue-DLO00030
cargo_goods_list>>picture_url	图片地址，查询添加商品列表接口对应的字段【picture_url】	是	[string]	https://i5.xxx.com/b.jpeg
cargo_goods_list>>product_desc	商品描述	是	[string]	桌子
cargo_goods_list>>product_id	商品id，查询添加商品列表接口对应的字段【item_id】	是	[string]	09123107025675
cargo_goods_list>>product_name	商品名字，查询添加商品列表接口对应的字段【local_name】	是	[string]	test1109pm
cargo_goods_list>>sku	SKU，查询添加商品列表接口对应的字段【local_sku】	是	[string]	SKU_CAR
cargo_goods_list>>product_type	商品类型：
GTIN GTIN
UPC UPC
EAN EAN	是	[string]	GTIN
cargo_goods_list>>gtin	GTIN，查询添加商品列表接口对应的字段【gtin】	是	[string]	09123107025675
cargo_goods_list>>value_added_service	增值服务类型：
1 需要
0 不需要	是	[string]	0
cargo_remark	货件备注	否	[string]	xxx
inbound_order_id	入库订单id	否	[string]	
return_address	退件地址，查询退件地址列表 接口获取	是	[object]	
return_address>>address_alias	地址别名	是	[string]	退货地址001
return_address>>address_id	地址id	是	[string]	39
return_address>>city	城市	是	[string]	深圳市
return_address>>mobile	电话	是	[string]	177xxxxxxxx
return_address>>postal_code	邮政编码	是	[string]	466100
return_address>>province	州/省/地区	是	[string]	广东省
return_address>>receive_or_delivery_country	发货方/收货方所属国家(或地区)编码	是	[string]	AD
return_address>>receive_or_delivery_country_name	发货方/收货方所属国家(或地区)名称	是	[string]	安道尔
return_address>>street_detail	街道详细地址	是	[string]	XXX街道，xxx社区，x栋
请求示例
{
  "store_id": "110000000020008001",
  "cargo_goods_list": [
    {
      "box_num": "1",
      "declare_num": "1",
      "expected_arrival_time": "2023-09-15 00:00:00",
      "msku": "WMUS-DLO-HairRoller-Blue-DLO00030",
      "picture_url": "https://i5.walmartimages.com/asr/22121042-a896-4b06-a312-23570f0a6162.20427cf1d30b27865e119f2b8599cc9b.jpeg",
      "product_desc": "1",
      "product_id": "09123107025675",
      "gtin": "09123107025675",
      "product_name": "test1109pm",
      "product_type": "GTIN",
      "value_added_service": 0,
      "single_box_num": "1",
      "sku": "test1109sku"
    }
  ],
  "cargo_remark": "",
  "inbound_order_id": "",
  "return_address": {
    "address_alias": "123",
    "address_id": "39",
    "city": "1",
    "mobile": "",
    "postal_code": "1",
    "province": "1",
    "receive_or_delivery_country": "AD",
    "receive_or_delivery_country_name": "安道尔",
    "street_detail": "1,"
  }
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
request_id	请求链路id	是	[string]	77ac259a67d5462594c83b80669b6eae.1692331008758
response_time	响应时间	是	[string]	2023-08-18 11:56:49
total	总数	是	[int]	200
data	响应数据	是	[object]	
data>>inbound_order_id	入库订单编号	是	[string]	202309150022
返回成功示例

{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "ace6b6587de441df83115493251e4068.1694771809931",
    "response_time": "2023-09-15 17:56:51",
    "data": {
        "inbound_order_id":"202309150022"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询WFS货件可添加商品列表
下一章节 
删除暂存货件