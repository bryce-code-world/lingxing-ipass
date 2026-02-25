添加/编辑多属性产品
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/spu/set	HTTPS	POST	5
请求参数
参数名	说明	必填	类型	示例
spu	SPU（添加时必填）	是	[string]	
spu_name	款名（添加时必填）	是	[string]	
model	型号	否	[string]	
unit	单位	否	[string]	
status	状态【默认1】：0 停售，1 在售，2 开发中，3 清仓	否	[int]	
cid	分类id	否	[int]	
bid	品牌id	否	[int]	
create_uid	创建人id	否	[int]	
developer_uid	开发人id	否	[int]	
product_duty_uids	产品负责人id	否	[array]	
description	产品描述	否	[string]	
use_spu_template	是否应用SPU信息至新生成的SKU：0 否，1 是	否	[int]	
sku_list	产品列表【提交的sku不存在时系统会自动创建】	是	[array]	
sku_list>>sku	本地产品SKU	是	[string]	
sku_list>>product_name	产品名称【提交的sku不存在时为必填项】	否	[string]	
sku_list>>attribute	属性列表	是	[array]	
sku_list>>attribute>>pa_id	属性id	是	[int]	
sku_list>>attribute>>pai_id	属性值id	是	[int]	
sku_list>>picture_list	产品图片信息	否	[array]	 
sku_lis>>picture_list>>pic_url	产品图片链接	是	[string]	
sku_list>>picture_list>>is_primary	是否产品主图:0否,1是	是	[int]	
purchase_info	采购相关信息	否	[object]	
purchase_info>>cg_uid	采购：采购员id	否	[int]	
purchase_info>>purchase_remark	采购：采购备注	否	[string]	
purchase_info>>cg_delivery	采购：采购交期（天）	否	[int]	
purchase_info>>cg_product_length	采购：单品规格-长（CM）	否	[number]	
purchase_info>>cg_product_width	采购：单品规格-宽（CM）	否	[number]	
purchase_info>>cg_product_height	采购：单品规格-高（CM）	否	[number]	
purchase_info>>cg_product_net_weight	采购：单品净重（G）	否	[number]	
purchase_info>>cg_product_gross_weight	采购：单品毛重（G）	否	[number]	
purchase_info>>cg_package_length	采购：包装规格-长（CM）	否	[number]	
purchase_info>>cg_package_width	采购：包装规格-宽（CM）	否	[number]	
purchase_info>>cg_package_height	采购：包装规格-高（CM）	否	[number]	
purchase_info>>cg_box_length	采购：外箱规格-长（CM）	否	[number]	
purchase_info>>cg_box_width	采购：外箱规格-宽（CM）	否	[number]	
purchase_info>>cg_box_height	采购：外箱规格-高（CM）	否	[number]	
purchase_info>>cg_box_weight	采购：单箱重量（KG）	否	[number]	
purchase_info>>cg_box_pcs	采购：单箱数量（包装数量）	否	[string]	
purchase_info>>cg_product_material	采购：产品材质	否	[string]	
logistics	物流报关相关信息	否	[object]	
logistics>>declaration	报关数据	否	[object]	
logistics>>declaration>>customs_export_name	报关：申报品名（中文）	否	[string]	
logistics>>declaration>>customs_import_name	报关：申报品名（英文）	否	[string]	
logistics>>declaration>>customs_import_price_currency	报关：申报单价的币种	否	[string]	
logistics>>declaration>>customs_import_price	报关：申报单价	否	[number]	
logistics>>declaration>>customs_declaration_unit	报关单位	否	[string]	
logistics>>declaration>>customs_declaration_spec	规格型号	否	[string]	
logistics>>declaration>>customs_declaration_origin_produce	报关：原厂国（地区）	否	[string]	
logistics>>declaration>>customs_declaration_inlands_source	报关：境内货源地	否	[string]	
logistics>>declaration>>customs_declaration_exempt	报关：征免	否	[string]	
logistics>>clearance	清关数据	否	[object]	
logistics>>clearance>>customs_clearance_material	清关：材质	否	[string]	
logistics>>clearance>>customs_clearance_usage	清关：用途	否	[string]	
logistics>>clearance>>customs_clearance_internal_code	清关：内部编码	否	[string]	
logistics>>clearance>>customs_clearance_preferential	清关：出口享惠情况：
1 不享惠
2 享惠
3 不确定享惠情况	否	[string]	
logistics>>clearance>>customs_clearance_brand_type	清关：品牌类型：
1 无品牌
2 境内自主品牌
3 境内收购品牌
4 境外品牌（贴牌生产）
5 境外品牌（其他）	否	[string]	
logistics>>clearance>>customs_clearance_product_pattern	清关：产品型号	否	[string]	
logistics>>clearance>>allocation_remark	清关：配货备注	否	[string]	
logistics>>clearance>>customs_clearance_pic_url	清关：清关图片	否	[string]	
logistics>>base	物流基础信息	否	[object]	
logistics>>base>>bg_export_hs_code	报关：HS Code（中国）	否	[string]	
logistics>>base>>special_attr	产品特殊属性：1 含电
2 纯电
3 液体
4 粉末
5 膏体
6 带磁	否	[array]	
logistics>>fee	头程费用，支持国家：US、CA、MX、JP、UK、DE、FR、ES、IT、NL、AU、SG、IN、AE、SA、BR、SE、PL、BE、TR、UA、HU、PK、LB、AT、CH、CZ、DK、IE、LU、NO、PT、SK、RU、KZ、BY、CL、KR	否	[object]	
logistics>>fee>>国家简码_cg_transport_costs	默认头程费用（含税）	否	[number]	
logistics>>fee>>国家简码_currency	默认头程费用币种	否	[string]	
logistics>>fee>>国家简码_clearance_price	清关价格	否	[number]	
logistics>>fee>>国家简码_clearance_price_currency	清关价格币种	否	[string]	
logistics>>fee>>国家简码_bg_import_hs_code	HS Code	否	[string]	
logistics>>fee>>国家简码_bg_tax_rate	税率	否	[number]	 
aux_relation_list	辅料列表	是	[array]	
aux_relation_list>>aux_sku	辅料sku	是	[string]	
aux_relation_list>>sku_qty	辅料比例（主料）	是	[string]	
aux_relation_list>>aux_qty	辅料比例（辅料）	是	[string]
请求示例
{
    "spu": "AAAAAABBB",
    "spu_name": "AAAAAAAAA",
    "model": "",
    "unit": "",
    "status": 0,
    "cid": "",
    "bid": "",
    "create_uid": 0,
    "developer_uid": "",
    "product_duty_uids": [],
    "description": "",
    "use_spu_template": 0,
    "sku_list": [
        {
            "sku": "CWZS20A-S",
            "product_name": "儿童沙滩罩衫",
            "picture_list": [
                {"pic_url":"http://hhjskbnd.png","is_primary":1}
            ],
            "attribute": [
                {
                    "pa_id": "",
                    "pai_id": ""
                }
            ]
        }
    ],
    "purchase_info": {
        "cg_uid": 1,
        "purchase_remark": "",
        "cg_delivery": 0,
        "cg_product_length": 0,
        "cg_product_width": 0,
        "cg_product_height": 0,
        "cg_product_net_weight": 0,
        "cg_product_gross_weight": 0,
        "cg_package_length": 0,
        "cg_package_width": 0,
        "cg_package_height": 0,
        "cg_box_length": 0,
        "cg_box_width": 0,
        "cg_box_height": 0,
        "cg_box_weight": 0,
        "cg_box_pcs": "",
        "cg_product_material": ""
    },
    "logistics": {
        "declaration": {
            "customs_export_name": "",
            "customs_import_name": "",
            "customs_import_price_currency": "",
            "customs_import_price": 0,
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
            "customs_clearance_preferential": "",
            "customs_clearance_brand_type": "",
            "customs_clearance_product_pattern": "",
            "allocation_remark": "",
            "customs_clearance_pic_url": ""
        },
        "base": {
            "bg_export_hs_code": "",
            "special_attr": []
        },
        "fee": {
            "US_cg_transport_costs": 0,
            "US_currency": "",
            "US_clearance_price": 0,
            "US_clearance_price_currency": "",
            "US_bg_import_hs_code": 0,
            "US_bg_tax_rate": 0
        }
    }
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
request_id	请求链路id	是	[string]	6D1422D5-B767-675D-42F6-F604319C2333
response_time	响应时间	是	[string]	2023-03-09 14:45:40
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>ps_id	spu唯一id	是	[int]	113
data>>sku_list	spu下对应的sku数据	是	[array]	
data>>sku_list>>product_id	本地产品id	是	[int]	48528
data>>sku_list>>sku	本地产品sku	是	[string]	CWZS20A-S
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "6D1422D5-B767-675D-42F6-F604319C2333",
    "response_time": "2023-03-09 14:45:40",
    "total": 0,
    "data": {
        "ps_id": 113,
        "sku_list": [
            {
                "product_id": 48528,
                "sku": "CWZS20A-S"
            },
            {
                "product_id": 48529,
                "sku": "CWZS20A-M"
            },
            {
                "product_id": 48530,
                "sku": "CWZS20A-L"
            }
        ]
    }
}
复制
错误
复制成功
 上一章节
查询多属性产品详情
下一章节 
查询捆绑产品关系列表