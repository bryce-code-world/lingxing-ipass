查询FBA库存列表-v2

支持查询FBA库存，对应系统【仓库】>【FBA库存明细】数据,数量维度展示

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/openapi/storage/fbaWarehouseDetail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20,取值范围[20,200]	否	[int]	20
search_field	搜索维度:
sku
product_name
seller_sku
fnsku
asin
parent_asin
spu
spu_name	否	[string]	seller_sku
search_value	搜索值	否	[string]	MSKUFDA5E30
cid	分类	否	[string]	
bid	品牌	否	[string]	
attribute	属性	否	[string]	
asin_principal	Listing负责人uid，对应查询ERP用户信息列表uid字段
多个使用,分隔	否	[string]	
status	在售状态:
0 停售
1 在售	否	[string]	
senior_search_list	高级搜索列表，详情见附加说明	否	[string]	
fulfillment_channel_type	配送方式:
FBA
FBM	否	[string]	
is_hide_zero_stock	是否隐藏零库存行:
0 不隐藏零库存行
1 隐藏零库存行	否	[string]	0
is_parant_asin_merge	是否合并父ASIN:
0 不合并父ASIN
1 合并父ASIN	否	[string]	0
is_contain_del_ls	是否显示已删除Listing:
0 显示已删除Listing
1 不显示已删除Listing	否	[string]	0
query_fba_storage_quantity_list	true 是、false 否；默认false，如果传入true,则出参数据中的欧洲共享仓会将出参字段-fba_storage_quantity_list的值返回	否	[Boolean]	true
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/openapi/storage/fbaWarehouseDetail?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "cid": "value",
    "bid": "value",
    "attribute": "value",
    "asin_principal": "value",
    "search_field": "value",
    "search_value": "value",
    "is_cost_page": value,
    "status": "value",
    "senior_search_list": "value",
    "offset": value,
    "length": value,
    "fulfillment_channel_type": "value",
    "is_hide_zero_stock": "value",
    "is_parant_asin_merge": "value",
    "is_contain_del_ls": "value"
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	数据校验失败时的错误详情	是	[array]	
request_id	请求链路id	是	[string]	a0d54debf93140f3b58d1ed81e8e3583.178.17255922733991817
response_time	响应时间	是	[string]	2024-09-06 11:11:13
data	响应数据	是	[array]	
data>>name	仓库名	是	[string]	仓库4
data>>seller_group_name	共享仓店铺名	是	[string]	
data>>sid	店铺id【当仓库为共享仓时，sid为0返回】	是	[int]	3
data>>asin	ASIN	是	[string]	B0ABC80784
data>>product_name	品名	是	[string]	Sony 儿童耳机
data>>small_image_url	预览图链接	是	[string]	
data>>seller_sku	MSKU	是	[string]	MSKUFDA5E30
data>>fnsku	FNSKU	是	[string]	FNC020467
data>>sku	SKU	是	[string]	SKU9E12CC2
data>>category_text	分类文本	是	[string]	分类2
data>>cid	分类Id	是	[int]	2
data>>product_brand_text	品牌文本	是	[string]	品牌3
data>>bid	品牌Id	是	[int]	3
data>>share_type	共享类型:
0 非共享
1 北美共享
2 欧洲共享	是	[int]	0
data>>total	总数	是	[int]	0
data>>total_price	总价	是	[double]	0
data>>available_total	可用总数	是	[int]	0
data>>available_total_price	可用总数成本价	是	[string]	
data>>afn_fulfillable_quantity	FBA可售	是	[int]	0
data>>afn_fulfillable_quantity_price	FBA可售成本价	是	[string]	0
data>>reserved_fc_transfers	待调仓	是	[int]	0
data>>reserved_fc_transfers_price	待调仓成本价	是	[string]	0
data>>reserved_fc_processing	调仓中	是	[int]	0
data>>reserved_fc_processing_price	调仓中成本价	是	[string]	0
data>>reserved_customerorders	待发货	是	[int]	0
data>>reserved_customerorders_price	待发货成本价	是	[string]	0
data>>quantity	FBM可售	是	[int]	0
data>>quantity_price	FBM可售成本价	是	[string]	0
data>>afn_unsellable_quantity	不可售	是	[int]	0
data>>afn_unsellable_quantity_price	不可售成本价	是	[string]	0
data>>afn_inbound_working_quantity	计划入库	是	[int]	0
data>>afn_inbound_working_quantity_price	计划入库成本价	是	[string]	0
data>>afn_inbound_shipped_quantity	在途	是	[int]	0
data>>afn_inbound_shipped_quantity_price	在途成本价	是	[string]	0
data>>afn_inbound_receiving_quantity	入库中	是	[int]	0
data>>afn_inbound_receiving_quantity_price	入库中成本价	是	[string]	0
data>>stock_up_num	实际在途	是	[int]	0
data>>stock_up_num_price	实际在途成本价	是	[string]	0
data>>afn_researching_quantity	调查中数量	是	[int]	0
data>>afn_researching_quantity_price	调查中数量成本价	是	[string]	0
data>>total_fulfillable_quantity	总可用库存:
可售+待调仓+调仓中
【非ERP页面对应总库存】	是	[int]	0
data>>inv_age_0_to_30_days	0-1个月库龄	是	[int]	0
data>>inv_age_0_to_30_price	0-1个月库龄成本价	是	[string]	0
data>>inv_age_31_to_60_days	1-2个月库龄	是	[int]	0
data>>inv_age_31_to_60_price	1-2个月库龄成本价	是	[string]	0
data>>inv_age_61_to_90_days	2-3个月库龄	是	[int]	0
data>>inv_age_61_to_90_price	2-3个月库龄成本价	是	[string]	0
data>>inv_age_0_to_90_days	0-3个月库龄	是	[int]	0
data>>inv_age_0_to_90_price	0-3个月库龄成本价	是	[string]	0
data>>inv_age_91_to_180_days	3-6个月库龄	是	[int]	0
data>>inv_age_91_to_180_price	3-6个月库龄成本价	是	[string]	0
data>>inv_age_181_to_270_days	6-9个月库龄	是	[int]	0
data>>inv_age_181_to_270_price	6-9个月库龄成本价	是	[string]	0
data>>inv_age_271_to_330_days	9-11个月库龄	是	[int]	0
data>>inv_age_271_to_330_price	9-11个月库龄成本价	是	[string]	0
data>>inv_age_271_to_365_days	9-12个月库龄	是	[int]	0
data>>inv_age_271_to_365_price	9-12个月库龄成本价	是	[string]	0
data>>inv_age_331_to_365_days	11-12个月库龄	是	[int]	0
data>>inv_age_331_to_365_price	11-12个月库龄成本价	是	[string]	0
data>>inv_age_365_plus_days	12个月以上库龄	是	[int]	0
data>>inv_age_365_plus_price	12个月以上库龄成本价	是	[string]	0
data>>recommended_action	推荐操作	是	[string]	
data>>sell_through	售出率	是	[double]	0
data>>estimated_excess_quantity	预计冗余数量	是	[double]	0
data>>estimated_storage_cost_next_month	预计30天仓储费用	是	[double]	0
data>>fba_minimum_inventory_level	最低库存水平	是	[double]	0
data>>fba_inventory_level_health_status	库存水平健康度	是	[string]	
data>>historical_days_of_supply	历史供货天数	是	[double]	0
data>>historical_days_of_supply_price	历史供货天数成本价	是	[string]	0
data>>low_inventory_level_fee_applied	低库存水平费收取情况	是	[string]	
data>>fulfillment_channel	配送方式	是	[string]	AMAZON_NA
data>>cg_price	单位采购成本	是	[string]	
data>>cg_transport_costs	单位头程费用	是	[string]	
data>>fba_storage_quantity_list	FBA可售信息列表，当仓库为共享仓库时，该字段才返回	否	[array]	
data>>fba_storage_quantity_list>>sid	店铺id	是	[int]	
data>>fba_storage_quantity_list>>name	店铺名称	是	[string]	
data>>fba_storage_quantity_list>>quantity_for_local_fulfillment	FBA可售数量	是	[int]	
total	总数	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "4683c92968ca47ccbff50f1ea8e4fb7a.165.17299260210002027",
    "response_time": "2024-10-26 15:00:21",
    "data": [
        {
            "name": "韧啸-US美国仓",
            "sid": 136,
            "asin": "B0D6XVLTJL",
            "product_name": "无线鼠标",
            "small_image_url": "https://image.distributetop.com/erp-vue/90128554873982976/20240802/027e8b7698414d899f7f0a2e77f97268.jpeg",
            "seller_sku": "20240613-1",
            "fnsku": "X004A0PW6J",
            "sku": "shubiao_1",
            "category_text": "电子设备\\鼠标",
            "cid": 3318,
            "product_brand_text": "罗技",
            "bid": 1173,
            "share_type": 0,
            "total": 2,
            "total_price": 40.00,
            "available_total": 42,
            "available_total_price": "840.00",
            "afn_fulfillable_quantity": 0,
            "afn_fulfillable_quantity_price": "0.00",
            "reserved_fc_transfers": 0,
            "reserved_fc_transfers_price": "0.00",
            "reserved_fc_processing": 0,
            "reserved_fc_processing_price": "0.00",
            "reserved_customerorders": 0,
            "reserved_customerorders_price": "0.00",
            "quantity": 0,
            "quantity_price": "0.00",
            "afn_unsellable_quantity": 0,
            "afn_unsellable_quantity_price": "0.00",
            "afn_inbound_working_quantity": 47,
            "afn_inbound_working_quantity_price": "940.00",
            "afn_inbound_shipped_quantity": 2,
            "afn_inbound_shipped_quantity_price": "40.00",
            "afn_inbound_receiving_quantity": 0,
            "afn_inbound_receiving_quantity_price": "0.00",
            "stock_up_num": 42,
            "stock_up_num_price": "840.00",
            "afn_researching_quantity": 0,
            "afn_researching_quantity_price": "0.00",
            "total_fulfillable_quantity": 0,
            "inv_age_0_to_90_days": 0,
            "inv_age_0_to_90_price": "0.00",
            "inv_age_271_to_365_days": 0,
            "inv_age_271_to_365_price": "0.00",
            "inv_age_0_to_30_days": 0,
            "inv_age_0_to_30_price": "0.00",
            "inv_age_31_to_60_days": 0,
            "inv_age_31_to_60_price": "0.00",
            "inv_age_61_to_90_days": 0,
            "inv_age_61_to_90_price": "0.00",
            "inv_age_91_to_180_days": 0,
            "inv_age_91_to_180_price": "0.00",
            "inv_age_181_to_270_days": 0,
            "inv_age_181_to_270_price": "0.00",
            "inv_age_271_to_330_days": 0,
            "inv_age_271_to_330_price": "0.00",
            "inv_age_331_to_365_days": 0,
            "inv_age_331_to_365_price": "0.00",
            "inv_age_365_plus_days": 0,
            "inv_age_365_plus_price": "0.00",
            "recommended_action": "",
            "sell_through": 0.0,
            "estimated_excess_quantity": 0.0,
            "estimated_storage_cost_next_month": 0.0,
            "fba_minimum_inventory_level": 0.0,
            "fba_inventory_level_health_status": "",
            "historical_days_of_supply": 0.00,
            "historical_days_of_supply_price": "0.00",
            "low_inventory_level_fee_applied": "",
            "fulfillment_channel": "AMAZON_NA",
            "fba_storage_quantity_list": [
                {
                    "sid": 39,
                    "name": "ouzhou001-MX",
                    "quantity_for_local_fulfillment": 10
                },
                {
                    "sid": 38,
                    "name": "ouzhou001-IT",
                    "quantity_for_local_fulfillment": 10
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
附加说明

senior_search_list 高级搜索列表字段说明如下：

senior_search_list 是一个 json 字符串，用于多条件组合搜索，每个元素表示一个搜索条件（字段+值列表）。 需严格符合 JSON 格式，字段名和值需用双引号包裹。

字段名	类型	必填	说明
name	string	是	搜索条件的显示名称（如“SKU”“品名”等），与 search_field 为固定枚举映射
search_field	string	是	搜索条件的显示名称（如“SKU”“品名”等），与 search_field 为固定枚举映射
search_value	array	是	搜索值列表，支持多值（逻辑“OR”关系）。

使用示例：

    "senior_search_list": "[{ \"name\":         \"SKU\", 
                              \"search_field\": \"sku\", 
                              \"search_value\": [\"sku1\", \"sku2\", \"sku3\"]},
                            { \"name\":         \"品名\", 
                              \"search_field\": \"product_name\", 
                              \"search_value\": [\"衣服\", \"裤子\", \"鞋子\"]},
                            { \"name\":         \"MSKU\", 
                              \"search_field\": \"seller_sku\", 
                              \"search_value\": [\"msku1\", \"msku2\", \"msku3\"]},
                            { \"name\":         \"FNSKU\", 
                              \"search_field\": \"fnsku\", 
                              \"search_value\": [\"fnsku1\", \"fnsku2\", \"fnsku3\"]},
                            { \"name\":         \"ASIN\", 
                              \"search_field\": \"asin\", 
                              \"search_value\": [\"asin1\", \"asin2\", \"asin3\"]},
                            { \"name\":         \"父ASIN\", 
                              \"search_field\": \"parent_asin\", 
                              \"search_value\": [\"fuAsin1\", \"fuAsin2\", \"fuAsin3\"]},
                            { \"name\":         \"SPU\", 
                              \"search_field\": \"spu\", 
                              \"search_value\": [\"spu1\", \"spu2\", \"spu3\"]},
                            { \"name\":         \"款名\", 
                              \"search_field\": \"spu_name\", 
                              \"search_value\": [\"kuan1\", \"kuan2\", \"kuan3\"]}]"
复制
错误
复制成功

name, search_field 字段说明如下： name 和 search_field 是一一对应的枚举值。

name	search_field
SKU	sku
品名	product_name
MSKU	seller_sku
FNSKU	fnsku
ASIN	asin
父ASIN	parent_asin
SPU	spu
款名	spu_name
 上一章节
查询FBA库存列表
下一章节 
查询AWD库存列表