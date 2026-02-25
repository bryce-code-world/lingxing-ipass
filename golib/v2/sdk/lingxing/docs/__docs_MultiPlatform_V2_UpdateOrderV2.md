编辑/更新自发货订单

批量编辑更新【订单管理】模块“待审核”状态订单

API Path	请求协议	请求方式	令牌桶容量
/pb/mp/order/v2/updateOrder	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
order_list	订单列表	是	[array]	
order_list>>address_info	收货信息	否	[object]	
order_list>>address_info>>address_line1	详细地址1	否	[string]	
order_list>>address_info>>address_line2	详细地址2	否	[string]	
order_list>>address_info>>address_line3	详细地址3	否	[string]	
order_list>>address_info>>city	城市	否	[string]	
order_list>>address_info>>district	区/县	否	[string]	
order_list>>address_info>>doorplate_no	门牌号	否	[string]	
order_list>>address_info>>postal_code	邮编	否	[string]	
order_list>>address_info>>receiver_company_name	公司名	否	[string]	
order_list>>address_info>>receiver_country_code	国家/地区二字码	否	[string]	
order_list>>address_info>>receiver_mobile	手机	否	[string]	
order_list>>address_info>>receiver_name	收件人	否	[string]	
order_list>>address_info>>receiver_tel	电话	否	[string]	
order_list>>address_info>>state_or_region	省/州	否	[string]	
order_list>>global_order_no	全局系统单号	否	[int]	
order_list>>logistics	物流信息	否	[object]	
order_list>>logistics>>cod_type	是否COD订单（是 or 否）	否	[string]	
order_list>>logistics>>sender_tax_no	税号	否	[string]	
order_list>>logistics>>sender_tax_type	税号类型（VAT/CPF/IOSS/EORI/收件人税号）	否	[string]	
order_list>>order_item_list	商品信息
备注：【可以传空列表】,不传空列表时order_list>>order_item_list>>type必传	是	[array]	
order_list>>order_item_list>>mark	商品备注	否	[string]	
order_list>>order_item_list>>msku	msku	否	[string]	
order_list>>order_item_list>>price	单价	否	[int)]	
order_list>>order_item_list>>quantity	数量	否	[int]	
order_list>>order_item_list>>sku	sku	否	[string]	
order_list>>order_item_list>>type	编辑类型：
1 新增
2 删除
3 覆盖	否	[int]	1
order_list>>order_item_list>>id	系统订单商品ID	否	[string]	1
order_list>>order_item_list>>platformOrderNo	平台单号，该字段在type=1 新增时生效	否	[string]	1
请求示例
{
    "order_list": [
        {
            "address_info": {
                "address_line1": "",
                "city": "",
                "receiver_country_code": "",
                "receiver_mobile": "",
                "receiver_name": ""
            },
            "global_order_no": 1032790274616770561,
            "order_item_list": [
                {
                    "quantity": 1,
                    "price": 1,
                    "sku": "wjc10_sku",
                    "type": 3
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
code	10002：订单更新成功
10000：订单更新失败
10001：订单更新部分成功	是	[int]	10001
message	消息提示	是	[string]	订单更新部分成功
data	响应数据	是	[array]	
data>>error_details	错误信息	是	[array]	[{global_order_no:xx,error_message:xxx},..]
data>>error_details>>global_order_no	系统单号		[string]	
data>>error_details>>error_message	失败原因		[string]	
request_id	请求链路id	是	[string]	ad9c1a8625c74f3ea0525e216374e31a1676514819613
response_time	响应时间	是	[string]	2023-02-16 10:33:40
返回成功示例
{
    "code": 10001,
    "message": "订单更新部分成功",
    "request_id": "d8e36903-8679-4372-8d09-21d80934d87a.1689907710574"
    "response_time": "2023-07-21 10:48:30",
    "data": {
        "error_details": [
            {
                "global_order_no": "1032790274616770561",
                "error_message": "系统不存在单号1032790274616770561"
            }
        ]
    }
}
复制
错误
复制成功
附加说明
编辑订单信息
文档包含字段均支持编辑。处商品信息外，请求体body中包含的字段名传入的，都将覆盖对应字段系统数据，不传字段名的将保留原订单数据。
即需要修改则传字段名，不需要修改则不传。
给订单中未配对的MSKU写入配对本地SKU
在order_item_list，传msku字段用于匹配订单中的msku，传需配对的sku，type=3覆盖。线上商品（一般有MSKU的）不支持修改数量和单价
给订单中已有配对本地SKU的MSKU更换SKU
在order_item_list中,传msku字段用于匹配订单中的msku，传新的sku，tpye=3覆盖。线上商品不支持修改数量和单价。
给订单新增一个本地sku
在order_item_list中，传sku、quantity、type=1新增，无需传msku。需指定新增的sku的数量。
仅手工订单类型的订单才能修改sku单价，线上单新增的sku视为单价为0的赠品。
order_list>>order_item_list>>price字段，实际业务含义=接口请求值*1000000
 上一章节
编辑订单（新版）
下一章节 
更新订单客服备注