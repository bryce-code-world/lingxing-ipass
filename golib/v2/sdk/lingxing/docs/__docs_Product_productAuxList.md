查询产品辅料列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/productAuxList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000，上限1000	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20
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
request_id	请求链路id	是	[string]	6D1422D5-B767-675D-42F6-F604319C2333
response_time	响应时间	是	[string]	2023-03-09 14:45:40
data	响应数据	是	[array]	
data>>id	辅料id	是	[int]	1
data>>sku	SKU	是	[string]	视觉111
data>>product_name	品名	是	[string]	经济
data>>cg_price	采购成本（人民币）	是	[number]	1.00
data>>cg_product_length	单品规格长（CM）	是	[number]	测试账号2
data>>cg_product_width	单品规格宽（CM）	是	[number]	
data>>cg_product_height	单品规格高（CM）	是	[number]	
data>>cg_product_net_weight	单品净重（G）	是	[number]	
data>>purchase_supplier_quote	供应商报价信息	是	[object]	
data>>purchase_supplier_quote>>product_id	产品ID	是	[int]	2
data>>purchase_supplier_quote>>cg_price	采购：采购成本（人民币）	是	[number]	33
data>>purchase_supplier_quote>>has_cg_permission	是否有采购成本权限：0-无，1-有	是	[int]	1
data>>purchase_supplier_quote>>suppliers		是	[object]	
data>>purchase_supplier_quote>>suppliers>>supplier_id	供应商ID	是	[int]	12
data>>purchase_supplier_quote>>suppliers>>is_primary	是否首选供应商：0-否，1-是	是	[int]	1
data>>purchase_supplier_quote>>suppliers>>supplier_name	供应商名称	是	[string]	供应商1
data>>purchase_supplier_quote>>suppliers>>contact_person	供应商联系人	是	[string]	XXX
data>>purchase_supplier_quote>>suppliers>>contact_number	供应商联系电话	是	[string]	18000000
data>>purchase_supplier_quote>>suppliers>>province	供应商地址：省	是	[string]	广东省
data>>purchase_supplier_quote>>suppliers>>city	供应商地址：城市	是	[string]	深圳市
data>>purchase_supplier_quote>>suppliers>>address	供应商地址：详细地址	是	[string]	西丽街道
data>>purchase_supplier_quote>>suppliers>>supplier_product_url	供应商链接	是	[array]	
data>>purchase_supplier_quote>>suppliers>>quotes	供应商报价	是	[array]	
data>>purchase_supplier_quote>>suppliers>>quotes>>currency	报价币种	是	[string]	
data>>purchase_supplier_quote>>suppliers>>quotes>>currency_icon	报价币种符号	是	[string]	
data>>purchase_supplier_quote>>suppliers>>quotes>>is_tax	是否含税：0-否，1-是	是	[int]	
data>>purchase_supplier_quote>>suppliers>>quotes>>tax_rate	税率（百分比）	是	[number]	
data>>purchase_supplier_quote>>suppliers>>quotes>>step_prices	报价梯度	是	[array]	
data>>purchase_supplier_quote>>suppliers>>quotes>>step_prices>>price_with_tax	含税单价	是	[number]	9.9
data>>purchase_supplier_quote>>suppliers>>quotes>>step_prices>>price	不含税单价	是	[number]	
data>>purchase_supplier_quote>>suppliers>>quotes>>step_prices>>moq	最小起购量	是	[int]	10
data>>aux_relation_product	关联产品	是	[array]	
data>>aux_relation_product>>pid	产品id	是	[int]	
data>>aux_relation_product>>product_name	产品名称	是	[string]	
data>>aux_relation_product>>sku	sku	是	[string]	
data>>aux_relation_product>>quantity	关联辅料的数量	是	[int]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "80E9323C-CF3A-7A1F-B9A1-D4A988635AC8",
    "response_time": "2023-03-09 15:58:12",
    "data": [
        {
            "aux_relation_product": [
                {
                    "pid": 2125093,
                    "product_name": "AUX11111",
                    "sku": "AUX11111",
                    "quantity": 2
                },
                {
                    "pid": 2125091,
                    "product_name": "钢化玻璃膜2",
                    "sku": "SKU2986DFA",
                    "quantity": 2
                }
            ],
            "purchase_supplier_quote": {
                "product_id": 2125093,
                "cg_price": "0.0000",
                "has_cg_permission": 1,
                "suppliers": {
                    "1": {
                        "supplier_id": 1,
                        "is_primary": 1,
                        "supplier_name": "供应商1",
                        "contact_person": "",
                        "contact_number": "",
                        "province": "",
                        "city": "",
                        "address": "",
                        "supplier_product_url": [
                            "https://detail.1688.com/offer/669376802464.html?spm=a260j.12536015.jr601u7p.1.29d4700ecBBrdJ",
                            "https://detail.1688.com/offer/669376802464.html?spm=a260j.12536015.jr601u7p.1.29d4700ecBBrdJ"
                        ],
                        "quotes": {
                            "CNY": {
                                "currency": "CNY",
                                "currency_icon": "￥",
                                "is_tax": 0,
                                "tax_rate": 0,
                                "step_prices": [
                                    {
                                        "price_with_tax": "0.0000",
                                        "price": "0.0000",
                                        "moq": 0
                                    }
                                ]
                            }
                        }
                    },
                    "245": {
                        "supplier_id": 245,
                        "is_primary": 0,
                        "supplier_name": "新增供应商NXthJByPZt",
                        "contact_person": "",
                        "contact_number": "",
                        "province": "",
                        "city": "",
                        "address": "",
                        "supplier_product_url": [],
                        "quotes": {
                            "CNY": {
                                "currency": "CNY",
                                "currency_icon": "￥",
                                "is_tax": 0,
                                "tax_rate": 0,
                                "step_prices": [
                                    {
                                        "price_with_tax": "0.0000",
                                        "price": "0.0000",
                                        "moq": 0
                                    }
                                ]
                            }
                        }
                    }
                }
            },
            "id": 4532,
            "sku": "asdfasd",
            "product_name": "chaolia",
            "cid": 0,
            "cg_price": "0.0000",
            "cg_product_length": "0.00",
            "cg_product_width": "0.00",
            "cg_product_height": "0.00",
            "cg_product_net_weight": "0.00",
            "remark": "",
            "is_aux": 1
        }
    ],
    "total": 42
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "offset 必须是非负数 [请求码:C365FF]",
    "request_id": "0A754F01-588E-C019-2693-5B72CF31D914",
    "response_time": "2021-11-11 18:26:35",
    "data": [],
    "total": 0
}
复制
错误
复制成功
 上一章节
添加/编辑捆绑产品
下一章节 
添加/编辑辅料