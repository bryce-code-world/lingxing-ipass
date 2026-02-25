添加/编辑本地产品

支持添加/编辑系统本地产品信息

支持国家：US（美国）、CA（加拿大）、MX（墨西哥）、JP（日本）、UK（英国）、DE（德国）、FR（法国）、ES（西班牙）、IT（意大利）、NL（荷兰）、 AU（澳洲）、SG（新加坡）、IN（印度）、AE（阿联酋）、SA（沙特阿拉伯）、BR（巴西）、SE（瑞典）、PL（波兰）、BE（比利时）、TR（土耳其）、 UA（乌克兰）、HU（匈牙利）、PK（巴基斯坦）、LB（黎巴嫩）、AT（奥地利）、CH（瑞士）、CZ（捷克）、DK（丹麦）、IE（爱尔兰）、LU（卢森堡）、 NO（挪威）、PT（葡萄牙）、SK（斯洛伐克）、RU（俄罗斯）、KZ（哈萨克斯坦）、BY（白俄罗斯）、CL（智利）、KR（韩国）

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/product/set	HTTPS	POST	10
请求参数
参数名	说明	必填	类型	示例
sku	SKU	是	[string]	
product_name	品名【添加时必填】	是	[string]	
sku_identifier	SKU识别码	否	[string]	
picture_list	产品图片信息	否	[array]	
picture_list>>pic_url	产品图片链接	否	[string]	
picture_list>>is_primary	是否产品主图：0 否，1 是	否	[int]	
unit	单位（商品单位：套、个、台）	否	[string]	
category_id	分类id，与分类同时存在时，优先取分类id	否	[int]	
category	分类	否	[string]	
model	型号	否	[string]	
brand_id	品牌id，与品牌同时存在时，优先取品牌id	否	[int]	
brand	品牌	否	[string]	
open_status	开启状态：0 停用，1 启用	否	[int]	
status	状态【默认1】：0 停售，1 在售，2 开发中，3 清仓	否	[int]	
description	商品描述	否	[string]	
group_list	组合商品列表	否	[array]	
group_list>>sku	子商品	否	[string]	
group_list>>quantity	商品比例数	否	[int]	
cg_opt_uid	采购：采购员id，与采购员名同时填写时，以采购员id为准	否	[int]	
cg_opt_username	采购：采购员名	否	[string]	
product_developer_uid	开发者id，与开发者名称同时填写时，以开发者id为准	否	[int]	
product_developer	开发者名称	否	[string]	
product_creator_uid	创建人id，默认API账号id	否	[int]	
product_duty_uids	负责人id	否	[array]	
is_append_product_duty	负责人是否追加创建人：0 否，1 是；默认1，只有编辑SKU时才生效	否	[int]	
purchase_remark	采购备注	否	[string]	
cg_price	采购：采购成本（RMB）	否	[string]	
is_related	是否关联单品成本：0 否，1 是	否	[int]	
cg_delivery	采购：采购交期	否	[int]	
cg_product_material	采购：商品材质	否	[string]	
cg_product_length	采购：单品规格-长（CM）	否	[string]	
cg_product_width	采购：单品规格-宽（CM）	否	[string]	
cg_product_height	采购：单品规格-高（CM）	否	[string]	
cg_product_net_weight	采购：单品净重（G）	否	[string]	
cg_product_gross_weight	采购：单品毛重（G）	否	[string]	
cg_package_length	采购：包装规格-长（CM）	否	[string]	
cg_package_width	采购：包装规格-宽（CM）	否	[string]	
cg_package_height	采购：包装规格-高（CM）	否	[string]	
cg_box_length	采购：外箱规格-长（CM）	否	[string]	
cg_box_width	采购：外箱规格-宽（CM）	否	[string]	
cg_box_height	采购：外箱规格-高（CM）	否	[string]	
cg_box_weight	采购：单箱重量（KG）	否	[string]	
cg_box_pcs	采购：单箱数量（包装数量）	否	[int]	
bg_customs_export_name	报关：申报品名(中文)	否	[string]	
bg_export_hs_code	报关：HS Code(中国)	否	[string]	
bg_customs_import_name	报关：申报品名(英文)	否	[string]	
currency	报关：申报金额的币种	否	[string]	
bg_customs_import_price	报关：申报金额	否	[string]	
qc_standard	质检标准	否	[object]	
qc_standard>>custom_qc_template	自定义质检标准	否	[object]	
qc_standard>>custom_qc_template>>qc_image	质检图片【最多十张图】	否	[array]	
qc_standard>>custom_qc_template>>qc_image>>file_id	质检图片文件id	否	[string]	
qc_standard>>custom_qc_template>>qc_image>>customer_url	客户的质检图片URL	否	[string]	
product_logistics_list	报关清关费用信息
支持国家：US、CA、MX、JP、UK、DE、FR、ES、IT、NL、AU、SG、IN、AE、SA、BR、SE、PL、BE、TR、UA、HU、PK、LB、AT、CH、CZ、DK、IE、LU、NO、PT、SK、RU、KZ、BY、CL、KR	否	[object]	
product_logistics_list>>US_cg_transport_costs	默认头程费用（含税）	否	[string]	
product_logistics_list>>US_currency	默认头程费用币种	否	[string]	
product_logistics_list>>US_clearance_price	清关价格	否	[string]	
product_logistics_list>>US_clearance_price_currency	清关价格币种，默认CNY	否	[string]	
product_logistics_list>>US_bg_import_hs_code	HS Code	否	[string]	
product_logistics_list>>US_bg_tax_rate	税率	否	[string]	
supplier_quote	供应商报价信息（该参数传空值则清空产品供应商报价）	否	[array]	
supplier_quote>>erp_supplier_id	领星ERP供应商id，查询本地产品详情接口对应字段【supplier_id】，与supplier_id必填其一	是	[int]	
supplier_quote>>supplier_id	客户系统供应商id，没有填这个值或者对应供应商不存在，则取erp_supplier_id，与erp_supplier_id必填其一	否	[int]	
supplier_quote>>supplier_product_url	采购链接，字符串数组，最多20个，没有则传空数组	否	[array]	
supplier_quote>>quote_remark	报价备注	否	[string]	
supplier_quote>>quote_cg_delivery	供应商交期	否	[int]	
supplier_quote>>is_primary	首选供应商：0 否，1 是	是	[int]	
supplier_quote>>quotes	报价信息，传供应商报价时必传	是	[array]	
supplier_quote>>quotes>>currency	报价币种，目前只有CNY和USD	是	[string]	
supplier_quote>>quotes>>is_tax	是否含税：0 否，1 是	是	[int]	
supplier_quote>>quotes>>tax_rate	税率，为空则表示为0	是	[string]	
supplier_quote>>quotes>>step_prices	阶梯价信息	是	[array]	
supplier_quote>>quotes>>step_prices>>moq	最小起订量，最小值为1	是	[int]	
supplier_quote>>quotes>>step_prices>>price_with_tax	含税单价，4位小数	是	[string]	
special_attr	产品特殊属性：1 含电，2 纯电，3 液体，4 粉末，5 膏体，6 带磁	否	[array]	
declaration	报关数据	否	[object]	
declaration>>customs_import_price	报关：报关单价	否	[int]	
declaration>>customs_import_price_currency	报关：报关单价币种	否	[string]	
declaration>>customs_export_name	报关：中文报关名	否	[string]	
declaration>>customs_import_name	报关：英文报关名	否	[string]	
declaration>>customs_declaration_unit	报关：报关单位	否	[string]	
declaration>>customs_declaration_spec	报关：规格型号	否	[string]	
declaration>>customs_declaration_origin_produce	报关：原厂国（地区）	否	[string]	
declaration>>customs_declaration_inlands_source	报关：境内货源地	否	[string]	
declaration>>customs_declaration_hs_code	报关：报关HSCODE	否	[string]	
declaration>>other_declare_element	报关：其他申报要素	否	[string]	
declaration>>customs_declaration_exempt	报关：征免	否	[string]	
clearance	清关数据	否	[object]	
clearance>>customs_clearance_material	清关：中文材质	否	[string]	
clearance>>customs_clearance_en_material	清关：英文材质	否	[string]	
clearance>>customs_clearance_usage	清关：中文用途	否	[string]	
clearance>>customs_clearance_en_usage	清关：英文用途	否	[string]	
clearance>>customs_clearance_internal_code	清关：内部编码	否	[string]	
clearance>>customs_clearance_preferential	清关：出口享惠情况：1 不享惠，2 享惠，3 不确定享惠情况	否	[int]	
clearance>>customs_clearance_brand_type	清关：品牌类型：1 无品牌，2 境内自主品牌，3 境内收购品牌，4 境外品牌（贴牌生产），5 境外品牌（其他）	否	[int]	
clearance>>customs_clearance_product_pattern	清关：产品型号	否	[string]	
clearance>>allocation_remark	清关：配货备注	否	[string]	
clearance>>weaving_mode	织造方式：1 针织，2 梭织	否	[int]	1
clearance>>customs_clearance_pic_url	清关：清关图片	否	[string]	 
aux_relation_list	辅料列表	否	[array]	
aux_relation_list>>aux_sku	辅料sku	是	[string]	
aux_relation_list>>sku_qty	辅料比例（主料）	否	[int]	
aux_relation_list>>aux_qty	辅料比例（辅料）	否	[int]	
spec_pack_list	采购：更多箱规（非默认箱规）	是	[array]	
spec_pack_list>>spec_title	采购：更多箱规-箱规名称	是	[string]	
spec_pack_list>>cg_box_pcs	采购：更多箱规-单箱数量	否	[int]	
spec_pack_list>>cg_box_length	采购：更多箱规-外箱规格-长（CM）	否	[string]	
spec_pack_list>>cg_box_width	采购：更多箱规-外箱规格-宽（CM）	否	[string]	
spec_pack_list>>cg_box_height	采购：更多箱规-外箱规格-高（CM）	否	[string]	
spec_pack_list>>cg_package_length	采购：更多箱规-包装规格-长（CM）	否	[string]	
spec_pack_list>>cg_package_width	采购：更多箱规-包装规格-宽（CM）	否	[string]	
spec_pack_list>>cg_package_height	采购：更多箱规-包装规格-高（CM）	否	[string]	
spec_pack_list>>cg_box_weight	采购：更多箱规-单箱重量（KG）	否	[string]	
spec_pack_list>>cg_product_gross_weight	采购：更多箱规-单品毛重（G）	否	[string]
请求示例
    {
    "sku": "SKU2986DFA",
    "product_name": "钢化玻璃膜2",
    "picture_list": [
        {
            "pic_url": "https://xxxx.com/123.jpg",
            "is_primary": 1
        }
    ],
    "unit": "个",
    "category_id": 2,
    "category": "分类5",
    "model": "",
    "brand_id": 0,
    "brand": "",
    "status": 0,
    "description": "",
    "group_list": [
        {
            "sku": "CWZS20A-L",
            "quantity": 1
        }
    ],
    "cg_opt_uid": 0,
    "cg_opt_username": "",
    "product_developer_uid": 0,
    "product_developer": "",
    "purchase_remark": "",
    "cg_price": "",
    "is_related": 0,
    "cg_delivery": 0,
    "cg_product_material": "",
    "cg_product_length": "",
    "cg_product_width": "",
    "cg_product_height": "",
    "cg_product_net_weight": "",
    "cg_product_gross_weight": "",
    "cg_package_length": "",
    "cg_package_width": "",
    "cg_package_height": "",
    "cg_box_length": "",
    "cg_box_width": "",
    "cg_box_height": "",
    "cg_box_weight": "",
    "cg_box_pcs": 0,
    "bg_customs_export_name": "",
    "bg_export_hs_code": "",
    "bg_customs_import_name": "",
    "currency": "CNY",
    "bg_customs_import_price": "",
    "product_creator_uid": 0,
    "product_duty_uids": [],
    "is_append_product_duty": 0,
    "qc_standard": {
        "custom_qc_template": {
            "qc_image": [
                {
                    "file_id": "",
                    "customer_url": ""
                }
            ]
        }
    },
    "product_logistics_list": {
        "US_cg_transport_costs": 0,
        "US_currency": "",
        "US_clearance_price": 0,
        "US_clearance_price_currency": "",
        "US_bg_import_hs_code": "",
        "US_bg_tax_rate": 0
    },
    "supplier_quote": [
        {
            "erp_supplier_id": 1,
            "supplier_id": 0,
            "supplier_product_url": [],
            "quote_remark": "",
            "is_primary": 1,
            "quotes": [
                {
                    "currency": "CNY",
                    "is_tax": 0,
                    "tax_rate": 0,
                    "step_prices": [
                        {
                            "moq": 0,
                            "price_with_tax": 10
                        }
                    ]
                }
            ]
        }
    ],
    "special_attr": [
        1,
        2,
        3
    ],
    "declaration": {
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
        "customs_clearance_brand_type": 0,
        "customs_clearance_product_pattern": "",
        "allocation_remark": "",
        "weaving_mode": 0,
        "customs_clearance_pic_url": ""
    },
    "aux_relation_list": [
        {
            "aux_sku": "AUX-GLUE-01",
            "sku_qty": 1,
            "aux_qty": 2
        }
    ],
    "spec_pack_list": [
        {
            "spec_title": "小箱包装",
            "cg_box_pcs": 50,
            "cg_box_length": "25.0",
            "cg_box_width": "15.0",
            "cg_box_height": "10.0",
            "cg_package_length": "12.0",
            "cg_package_width": "6.0",
            "cg_package_height": "1.0",
            "cg_box_weight": "6.0",
            "cg_product_gross_weight": "25.0"
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
request_id	请求链路id	是	[string]	 
response_time	响应时间	是	[string]	2020-04-30 17:33:32
data	响应数据	是	[object]	
data>>product_id	本地产品id	是	[int]	1010
data>>sku	本地产品sku	是	[string]	SKU2986DFA1
data>>sku_identifier	SKU识别码	是	[string]	SKU2986DFA1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "CCF38623-DE1A-004F-0961-90A6AF265013",
    "response_time": "2021-11-11 11:14:33",
    "data": {
        "product_id": 1010,
        "sku": "SKU2986DFA1",
        "sku_identifier": "SKU2986DFA1"
    },
    "total": 0
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "is_primary不能为空 [请求码:4B70FE]",
    "request_id": "D00AB182-BC9B-4859-854E-BEFCE2C1411C",
    "response_time": "2021-11-11 11:01:13",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
产品启用、禁用
下一章节 
查询产品属性列表