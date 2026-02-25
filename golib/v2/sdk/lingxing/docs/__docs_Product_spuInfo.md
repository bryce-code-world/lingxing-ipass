查询多属性产品详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/spu/info	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
ps_id	SPU唯一id【ps_id 与 spu二选一必填	是	[int]	1
spu	SPU	是	[string]	spu
请求示例
{
    "ps_id": 1,
    "spu": "spu"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示消息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	7FB0A1FA-15DD-9F23-613C-856BA552C51A
response_time	响应时间	是	[string]	2022-09-13 15:14:47
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>ps_id	spu 唯一id	是	[int]	
data>>spu	SPU（添加时必填）	是	[string]	
data>>spu_name	款名（添加时必填）	是	[string]	
data>>model	型号	是	[string]	
data>>unit	单位	是	[string]	
data>>status	状态：0 停售，1 在售，2 开发中，3 清仓	是	[int]	
data>>cid	分类id	是	[int]	
data>>category_name	分类名	是	[string]	
data>>bid	品牌id	是	[int]	
data>>brand_name	品牌名	是	[string]	
data>>developer_uid	开发人id	是	[int]	
data>>product_duty_uids	产品负责人id	是	[array]	
data>>description	产品描述	是	[string]	
data>>attachmentFiles	附件信息	是	[array]	
data>>attachmentFiles>>file_id	附件id	是	[string]	
data>>attachmentFiles>>file_name	附件名称	是	[string]	
data>>attachmentFiles>>file_url	附件url	是	[string]	
data>>purchase_info	采购信息	是	[object]	
data>>purchase_info>>cg_uid	采购：采购员id	是	[int]	
data>>purchase_info>>purchase_remark	采购：采购备注	是	[string]	
data>>purchase_info>>cg_delivery	采购：采购交期（天）	是	[int]	
data>>purchase_info>>cg_product_length	采购：单品规格-长（CM）	是	[number]	
data>>purchase_info>>cg_product_width	采购：单品规格-宽（CM）	是	[number]	
data>>purchase_info>>cg_product_height	采购：单品规格-高（CM）	是	[number]	
data>>purchase_info>>cg_product_net_weight	采购：单品净重（g）	是	[number]	
data>>purchase_info>>cg_product_gross_weight	采购：单品毛重（g）	是	[number]	
data>>purchase_info>>cg_package_length	采购：包装规格-长（CM）	是	[number]	
data>>purchase_info>>cg_package_width	采购：包装规格-宽（CM）	是	[number]	
data>>purchase_info>>cg_package_height	采购：包装规格-高（CM）	是	[number]	
data>>purchase_info>>cg_box_length	采购：外箱规格-长（CM）	是	[number]	
data>>purchase_info>>cg_box_width	采购：外箱规格-宽（CM）	是	[number]	
data>>purchase_info>>cg_box_height	采购：外箱规格-高（CM）	是	[number]	
data>>purchase_info>>cg_product_box_weight	采购：单箱重量（KG）	是	[number]	
data>>purchase_info>>cg_box_pcs	采购：单箱数量（包装数量）	是	[int]	
data>>purchase_info>>cg_product_material	采购：产品材质	是	[string]	
data>>aux_relation_list	关联辅料	是	[array]	
data>>aux_relation_list>>aux_id	辅料id	是	[int]	
data>>aux_relation_list>>aux_name	辅料品名	是	[string]	
data>>aux_relation_list>>aux_sku	辅料SKU	是	[string]	
data>>aux_relation_list>>cg_price	单位成本	是	[string]	
data>>aux_relation_list>>quantity	数量	是	[int]	
data>>aux_relation_list>>remark	备注	是	[string]	
data>>logistics	物流报关清关	是	[object]	
data>>logistics>>declaration	报关数据	是	[object]	
data>>logistics>>declaration>>customs_export_name	报关：申报品名（中文）	是	[string]	
data>>logistics>>declaration>>customs_import_name	报关：申报品名（英文）	是	[string]	
data>>logistics>>declaration>>customs_import_price_currency	报关：申报单价的币种	是	[string]	
data>>logistics>>declaration>>customs_import_price	报关：申报单价	是	[string]	
data>>logistics>>declaration>>customs_declaration_unit	报关单位	是	[string]	
data>>logistics>>declaration>>customs_declaration_spec	规格型号	是	[string]	
data>>logistics>>declaration>>customs_declaration_origin_produce	报关：原厂国（地区）	是	[string]	
data>>logistics>>declaration>>customs_declaration_inlands_source	报关：境内货源地	是	[string]	
data>>logistics>>declaration>>customs_declaration_exempt	报关：征免	是	[string]	
data>>logistics>>clearance	清关数据	是	[object]	
data>>logistics>>clearance>>customs_clearance_material	清关：材质	是	[string]	
data>>logistics>>clearance>>customs_clearance_usage	清关：用途	是	[string]	
data>>logistics>>clearance>>customs_clearance_internal_code	清关：内部编码	是	[string]	
data>>logistics>>clearance>>customs_clearance_preferential	清关：出口享惠情况：
1 不享惠
2 享惠
3 不确定享惠情况	是	[int]	
data>>logistics>>clearance>>customs_clearance_brand_type	清关：品牌类型：
1 无品牌
2 境内自主品牌
3 境内收购品牌
4 境外品牌（贴牌生产）
5 境外品牌（其他）	是	[int]	
data>>logistics>>clearance>>customs_clearance_product_pattern	清关：产品型号	是	[string]	
data>>logistics>>clearance>>allocation_remark	清关：配货备注	是	[string]	
data>>logistics>>clearance>>customs_clearance_pic_url	清关：清关图片	是	[string]	
data>>logistics>>base	物流基础信息	是	[object]	
data>>logistics>>base>>bg_export_hs_code	报关：HS Code（中国）	是	[string]	
data>>logistics>>base>>special_attr	产品特殊属性：1 含电
2 纯电
3 液体
4 粉末
5 膏体
6 带磁	是	[array]	
data>>logistics>>fee	头程费用，支持国家：US、CA、MX、JP、UK、DE、FR、ES、IT、NL、AU、SG、IN、AE、SA、BR、SE、PL、BE、TR、UA、HU、PK、LB、AT、CH、CZ、DK、IE、LU、NO、PT、SK、RU、KZ、BY、CL、KR	是	[object]	
data>>logistics>>fee>>国家简码_cg_transport_costs	默认头程费用（含税）	是	[number]	
data>>logistics>>fee>>国家简码_currency	默认头程费用币种	是	[string]	
data>>logistics>>fee>>国家简码_clearance_price	清关价格	是	[number]	
data>>logistics>>fee>>国家简码_clearance_price_currency	清关价格币种	是	[string]	
data>>logistics>>fee>>国家简码_bg_import_hs_code	HS Code	是	[string]	
data>>logistics>>fee>>国家简码_bg_tax_rate	税率	是	[number]	
data>>logistics>>fee>>国家简码_clearance_remark	备注	是	[string]	
data>>sku_list	产品列表	是	[array]	
data>>sku_list>>sku	产品SKU	是	[string]	
data>>sku_list>>product_name	产品名称	是	[string]	
data>>sku_list>>attribute	属性列表	是	[array]	
data>>sku_list>>attribute>>pa_id	属性id	是	[int]	
data>>sku_list>>attribute>>pai_id	属性值id	是	[int]	
data>>sku_list>>picture_list	产品图片链接	是	[array]	 
aux_relation_list	辅料列表	是	[object]	
aux_relation_list>>aux_sku	辅料sku	是	[string]	
aux_relation_list>>aux_name	辅料名称	是	[string]	
aux_relation_list>>sku_qty	辅料比例（主料）	是	[string]	
aux_relation_list>>aux_qty	辅料比例（辅料）	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7FB0A1FA-15DD-9F23-613C-856BA552C51A",
    "response_time": "2022-09-13 15:14:47",
    "total": 0,
    "data": {
        "ps_id": 1,
        "spu": "spu-0001",
        "spu_name": "spu-0001",
        "model": "1",
        "unit": "1",
        "status": 1,
        "status_text": "在售",
        "cid": 1275,
        "category_name": "zdh1680246622.2350538",
        "bid": 1012,
        "brand_name": "苹果",
        "create_user": "小花",
        "developer_uid": 10397585,
        "developer": "chenjinbo1",
        "product_duty_users": [],
        "description": "<p>11111</p>",
        "attachmentFiles": [
            {
                "file_id": "5444",
                "file_name": "3.jpg",
                "file_url": "/api/type/1.html"
            }
        ],
        "purchase_info": {
            "purchase_remark": "1111",
            "cg_delivery": 1,
            "cg_price": "11.0000",
            "cg_product_length": "1.00",
            "cg_product_length_in": "0.39",
            "cg_product_width": "1.00",
            "cg_product_width_in": "0.39",
            "cg_product_height": "10.00",
            "cg_product_height_in": "3.94",
            "cg_package_length": "1.00",
            "cg_package_length_in": "0.39",
            "cg_package_width": "1.00",
            "cg_package_width_in": "0.39",
            "cg_package_height": "10.00",
            "cg_package_height_in": "3.94",
            "cg_box_length": "1.00",
            "cg_box_length_in": "0.39",
            "cg_box_width": "1.00",
            "cg_box_width_in": "0.39",
            "cg_box_height": "0.10",
            "cg_box_height_in": "0.04",
            "cg_product_net_weight": "0.01",
            "cg_product_net_weight_in": "0.00",
            "cg_product_gross_weight": "1.00",
            "cg_product_gross_weight_in": "0.00",
            "cg_box_weight": "1.00",
            "cg_box_weight_in": "2.20",
            "cg_box_pcs": 1,
            "cg_product_material": "1",
            "cg_product_net_weight_unit": "g",
            "cg_product_net_weight_in_unit": "lb",
            "cg_product_gross_weight_unit": "g",
            "cg_product_gross_weight_in_unit": "lb",
            "cg_box_weight_unit": "kg",
            "cg_box_weight_in_unit": "lb",
            "cg_uid": 10317902,
            "cg_user": "小红"
        },
        "aux_relation_list": [
            {
                "aux_id": 22072,
                "aux_name": "测试镜",
                "aux_sku": "hahaceshijing",
                "cg_price": "10.000000",
                "quantity": 1,
                "remark": "测试"
            }
        ],
        "logistics": {
            "base": {
                "special_attr": [],
                "bg_export_hs_code": "0"
            },
            "fee": {
                "US_cg_transport_costs": "0.0000",
                "US_currency": "CNY",
                "US_bg_import_hs_code": "",
                "US_bg_tax_rate": "0.0000",
                "US_clearance_price_currency": "CNY",
                "US_clearance_price": "0.00",
                "US_clearance_remark": "",
                "CA_cg_transport_costs": "0.0000",
                "CA_currency": "CNY",
                "CA_bg_import_hs_code": "",
                "CA_bg_tax_rate": "0.0000",
                "CA_clearance_price_currency": "CNY",
                "CA_clearance_price": "0.00",
                "CA_clearance_remark": "",
                "MX_cg_transport_costs": "0.0000",
                "MX_currency": "CNY",
                "MX_bg_import_hs_code": "",
                "MX_bg_tax_rate": "0.0000",
                "MX_clearance_price_currency": "CNY",
                "MX_clearance_price": "0.00",
                "MX_clearance_remark": "",
                "JP_cg_transport_costs": "0.0000",
                "JP_currency": "CNY",
                "JP_bg_import_hs_code": "",
                "JP_bg_tax_rate": "0.0000",
                "JP_clearance_price_currency": "CNY",
                "JP_clearance_price": "0.00",
                "JP_clearance_remark": "",
                "UK_cg_transport_costs": "0.0000",
                "UK_currency": "CNY",
                "UK_bg_import_hs_code": "",
                "UK_bg_tax_rate": "0.0000",
                "UK_clearance_price_currency": "CNY",
                "UK_clearance_price": "0.00",
                "UK_clearance_remark": "",
                "DE_cg_transport_costs": "0.0000",
                "DE_currency": "CNY",
                "DE_bg_import_hs_code": "",
                "DE_bg_tax_rate": "0.0000",
                "DE_clearance_price_currency": "CNY",
                "DE_clearance_price": "0.00",
                "DE_clearance_remark": "",
                "FR_cg_transport_costs": "0.0000",
                "FR_currency": "CNY",
                "FR_bg_import_hs_code": "",
                "FR_bg_tax_rate": "0.0000",
                "FR_clearance_price_currency": "CNY",
                "FR_clearance_price": "0.00",
                "FR_clearance_remark": "",
                "ES_cg_transport_costs": "0.0000",
                "ES_currency": "CNY",
                "ES_bg_import_hs_code": "",
                "ES_bg_tax_rate": "0.0000",
                "ES_clearance_price_currency": "CNY",
                "ES_clearance_price": "0.00",
                "ES_clearance_remark": "",
                "IT_cg_transport_costs": "0.0000",
                "IT_currency": "CNY",
                "IT_bg_import_hs_code": "",
                "IT_bg_tax_rate": "0.0000",
                "IT_clearance_price_currency": "CNY",
                "IT_clearance_price": "0.00",
                "IT_clearance_remark": "",
                "NL_cg_transport_costs": "0.0000",
                "NL_currency": "CNY",
                "NL_bg_import_hs_code": "",
                "NL_bg_tax_rate": "0.0000",
                "NL_clearance_price_currency": "CNY",
                "NL_clearance_price": "0.00",
                "NL_clearance_remark": "",
                "AU_cg_transport_costs": "0.0000",
                "AU_currency": "CNY",
                "AU_bg_import_hs_code": "",
                "AU_bg_tax_rate": "0.0000",
                "AU_clearance_price_currency": "CNY",
                "AU_clearance_price": "0.00",
                "AU_clearance_remark": "",
                "SG_cg_transport_costs": "0.0000",
                "SG_currency": "CNY",
                "SG_bg_import_hs_code": "",
                "SG_bg_tax_rate": "0.0000",
                "SG_clearance_price_currency": "CNY",
                "SG_clearance_price": "0.00",
                "SG_clearance_remark": "",
                "IN_cg_transport_costs": "0.0000",
                "IN_currency": "CNY",
                "IN_bg_import_hs_code": "",
                "IN_bg_tax_rate": "0.0000",
                "IN_clearance_price_currency": "CNY",
                "IN_clearance_price": "0.00",
                "IN_clearance_remark": "",
                "AE_cg_transport_costs": "0.0000",
                "AE_currency": "CNY",
                "AE_bg_import_hs_code": "",
                "AE_bg_tax_rate": "0.0000",
                "AE_clearance_price_currency": "CNY",
                "AE_clearance_price": "0.00",
                "AE_clearance_remark": "",
                "SA_cg_transport_costs": "0.0000",
                "SA_currency": "CNY",
                "SA_bg_import_hs_code": "",
                "SA_bg_tax_rate": "0.0000",
                "SA_clearance_price_currency": "CNY",
                "SA_clearance_price": "0.00",
                "SA_clearance_remark": "",
                "BR_cg_transport_costs": "0.0000",
                "BR_currency": "CNY",
                "BR_bg_import_hs_code": "",
                "BR_bg_tax_rate": "0.0000",
                "BR_clearance_price_currency": "CNY",
                "BR_clearance_price": "0.00",
                "BR_clearance_remark": "",
                "SE_cg_transport_costs": "0.0000",
                "SE_currency": "CNY",
                "SE_bg_import_hs_code": "",
                "SE_bg_tax_rate": "0.0000",
                "SE_clearance_price_currency": "CNY",
                "SE_clearance_price": "0.00",
                "SE_clearance_remark": "",
                "PL_cg_transport_costs": "0.0000",
                "PL_currency": "CNY",
                "PL_bg_import_hs_code": "",
                "PL_bg_tax_rate": "0.0000",
                "PL_clearance_price_currency": "CNY",
                "PL_clearance_price": "0.00",
                "PL_clearance_remark": "",
                "BE_cg_transport_costs": "0.0000",
                "BE_currency": "CNY",
                "BE_bg_import_hs_code": "",
                "BE_bg_tax_rate": "0.0000",
                "BE_clearance_price_currency": "CNY",
                "BE_clearance_price": "0.00",
                "BE_clearance_remark": "",
                "TR_cg_transport_costs": "0.0000",
                "TR_currency": "CNY",
                "TR_bg_import_hs_code": "",
                "TR_bg_tax_rate": "0.0000",
                "TR_clearance_price_currency": "CNY",
                "TR_clearance_price": "0.00",
                "TR_clearance_remark": "",
                "UA_cg_transport_costs": 0,
                "UA_currency": "CNY",
                "UA_bg_import_hs_code": "",
                "UA_bg_tax_rate": 0,
                "UA_clearance_price_currency": "CNY",
                "UA_clearance_price": 0,
                "UA_clearance_remark": "",
                "HU_cg_transport_costs": 0,
                "HU_currency": "CNY",
                "HU_bg_import_hs_code": "",
                "HU_bg_tax_rate": 0,
                "HU_clearance_price_currency": "CNY",
                "HU_clearance_price": 0,
                "HU_clearance_remark": "",
                "PK_cg_transport_costs": 0,
                "PK_currency": "CNY",
                "PK_bg_import_hs_code": "",
                "PK_bg_tax_rate": 0,
                "PK_clearance_price_currency": "CNY",
                "PK_clearance_price": 0,
                "PK_clearance_remark": "",
                "LB_cg_transport_costs": 0,
                "LB_currency": "CNY",
                "LB_bg_import_hs_code": "",
                "LB_bg_tax_rate": 0,
                "LB_clearance_price_currency": "CNY",
                "LB_clearance_price": 0,
                "LB_clearance_remark": "",
                "AT_cg_transport_costs": 0,
                "AT_currency": "CNY",
                "AT_bg_import_hs_code": "",
                "AT_bg_tax_rate": 0,
                "AT_clearance_price_currency": "CNY",
                "AT_clearance_price": 0,
                "AT_clearance_remark": "",
                "CH_cg_transport_costs": 0,
                "CH_currency": "CNY",
                "CH_bg_import_hs_code": "",
                "CH_bg_tax_rate": 0,
                "CH_clearance_price_currency": "CNY",
                "CH_clearance_price": 0,
                "CH_clearance_remark": "",
                "CZ_cg_transport_costs": 0,
                "CZ_currency": "CNY",
                "CZ_bg_import_hs_code": "",
                "CZ_bg_tax_rate": 0,
                "CZ_clearance_price_currency": "CNY",
                "CZ_clearance_price": 0,
                "CZ_clearance_remark": "",
                "DK_cg_transport_costs": 0,
                "DK_currency": "CNY",
                "DK_bg_import_hs_code": "",
                "DK_bg_tax_rate": 0,
                "DK_clearance_price_currency": "CNY",
                "DK_clearance_price": 0,
                "DK_clearance_remark": "",
                "IE_cg_transport_costs": 0,
                "IE_currency": "CNY",
                "IE_bg_import_hs_code": "",
                "IE_bg_tax_rate": 0,
                "IE_clearance_price_currency": "CNY",
                "IE_clearance_price": 0,
                "IE_clearance_remark": "",
                "LU_cg_transport_costs": 0,
                "LU_currency": "CNY",
                "LU_bg_import_hs_code": "",
                "LU_bg_tax_rate": 0,
                "LU_clearance_price_currency": "CNY",
                "LU_clearance_price": 0,
                "LU_clearance_remark": "",
                "NO_cg_transport_costs": 0,
                "NO_currency": "CNY",
                "NO_bg_import_hs_code": "",
                "NO_bg_tax_rate": 0,
                "NO_clearance_price_currency": "CNY",
                "NO_clearance_price": 0,
                "NO_clearance_remark": "",
                "PT_cg_transport_costs": 0,
                "PT_currency": "CNY",
                "PT_bg_import_hs_code": "",
                "PT_bg_tax_rate": 0,
                "PT_clearance_price_currency": "CNY",
                "PT_clearance_price": 0,
                "PT_clearance_remark": "",
                "SK_cg_transport_costs": 0,
                "SK_currency": "CNY",
                "SK_bg_import_hs_code": "",
                "SK_bg_tax_rate": 0,
                "SK_clearance_price_currency": "CNY",
                "SK_clearance_price": 0,
                "SK_clearance_remark": "",
                "RU_cg_transport_costs": 0,
                "RU_currency": "CNY",
                "RU_bg_import_hs_code": "",
                "RU_bg_tax_rate": 0,
                "RU_clearance_price_currency": "CNY",
                "RU_clearance_price": 0,
                "RU_clearance_remark": "",
                "KZ_cg_transport_costs": 0,
                "KZ_currency": "CNY",
                "KZ_bg_import_hs_code": "",
                "KZ_bg_tax_rate": 0,
                "KZ_clearance_price_currency": "CNY",
                "KZ_clearance_price": 0,
                "KZ_clearance_remark": "",
                "BY_cg_transport_costs": 0,
                "BY_currency": "CNY",
                "BY_bg_import_hs_code": "",
                "BY_bg_tax_rate": 0,
                "BY_clearance_price_currency": "CNY",
                "BY_clearance_price": 0,
                "BY_clearance_remark": "",
                "CL_cg_transport_costs": 0,
                "CL_currency": "CNY",
                "CL_bg_import_hs_code": "",
                "CL_bg_tax_rate": 0,
                "CL_clearance_price_currency": "CNY",
                "CL_clearance_price": 0,
                "CL_clearance_remark": "",
                "KR_cg_transport_costs": 0,
                "KR_currency": "CNY",
                "KR_bg_import_hs_code": "",
                "KR_bg_tax_rate": 0,
                "KR_clearance_price_currency": "CNY",
                "KR_clearance_price": 0,
                "KR_clearance_remark": ""
            },
            "declaration": {
                "customs_declaration_hs_code": "0",
                "customs_import_price": "0.00",
                "customs_import_price_currency": "USD",
                "customs_import_price_currency_icon": "$",
                "customs_export_name": "",
                "customs_import_name": "",
                "customs_declaration_unit": "",
                "customs_declaration_spec": "",
                "customs_declaration_origin_produce": "",
                "customs_declaration_inlands_source": "",
                "customs_declaration_exempt": ""
            },
            "clearance": {
                "customs_clearance_material": "",
                "customs_clearance_usage": "",
                "customs_clearance_internal_code": "",
                "customs_clearance_preferential": 0,
                "customs_clearance_preferential_text": "",
                "customs_clearance_brand_type": 0,
                "customs_clearance_brand_type_text": "",
                "customs_clearance_product_pattern": "",
                "customs_clearance_pic_url": "",
                "allocation_remark": "",
                "weaving_mode": 0,
                "weaving_mode_text": "",
                "customs_clearance_hs_code": "0",
                "customs_clearance_tax_rate": "0.0000",
                "customs_clearance_price": "0.0000",
                "customs_clearance_price_currency": "USD",
                "customs_clearance_price_currency_icon": "$",
                "customs_clearance_remark": ""
            }
        },
        "sku_list": [
            {
                "attribute": [
                    {
                        "pa_id": 2,
                        "pai_id": "5",
                        "pai_name": "34"
                    }
                ],
                "product_id": 10010,
                "sku": "test-0001",
                "pic_url": "https://xxx/7f421aede0720f.png",
                "product_name": "test-0001",
                "pic_list": [
                    {
                        "pp_id": 210205166952853504,
                        "pic_name": "01.png",
                        "pic_url": "https://xxx/7f421aede0720f.png",
                        "product_id": 10010,
                        "is_primary": 1,
                        "pic_space": 294942,
                        "pic_size_w": 505,
                        "pic_size_h": 506,
                        "pic_type": 1
                    }
                ]
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
查询多属性产品列表
下一章节 
添加/编辑多属性产品