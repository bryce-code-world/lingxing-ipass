添加 / 编辑辅料
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/product/setAux	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sku	SKU	是	[string]	
product_name	品名	是	[string]	
cg_price	采购：采购成本（人民币）	否	[number]	
cg_product_length	采购：单品规格-长（CM）	否	[number]	
cg_product_width	采购：单品规格-宽（CM）	否	[number]	
cg_product_height	采购：单品规格-高（CM）	否	[number]	
cg_product_net_weight	采购：单品净重（G）	否	[number]	
supplier_quote	供应商报价信息（不传该参数则清空产品供应商报价）	否	[array]	
supplier_quote>>erp_supplier_id	领星ERP供应商id	否	[int]	
supplier_quote>>supplier_id	客户系统供应商id，没有填这个值或者对应供应商不存在，则取erp_supplier_id	否	[int]	
supplier_quote>>supplier_product_url	采购链接，字符串数组，最多20个，没有则传空数组	否	[array]	
supplier_quote>>is_primary	是否首选供应商：0-否，1:是	否	[int]	
supplier_quote>>quotes	报价信息	否	[array]	
supplier_quote>>quotes>>currency	报价币种，目前只有CNY和USD	否	[string]	
supplier_quote>>quotes>>is_tax	是否含税：0-否，1-是	否	[int]	
supplier_quote>>quotes>>tax_rate	税率，为空则表示为0	否	[int]	
supplier_quote>>quotes>>step_prices	阶梯价信息	否	[array]	
supplier_quote>>quotes>>step_prices>>moq	最小起订量，最小值为1	否	[int]	
supplier_quote>>quotes>>step_prices>>price_with_tax	含税单价，4位小数	否	[number]	
remark	辅料描述	是	[string]	 
请求示例
{
    "sku": "AUX11111",
    "product_name": "AUX11111",
    "cg_price": "",
    "cg_product_length": "",
    "cg_product_width": "",
    "cg_product_height": "",
    "cg_product_net_weight": "",
    "supplier_quote": [{
        "erp_supplier_id": 1,
        "supplier_id": 0,
        "supplier_product_url": [],
        "quote_remark": "",
        "is_primary": 1,
        "quotes": [{
            "currency": "CNY",
            "is_tax": 0,
            "tax_rate": 0,
            "step_prices": [{
                "moq": 0,
                "price_with_tax": 0
            }]
        }]
    }],
    "remark": "备注"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	错误信息	是	[string]	success
error_details	具体错误信息	是	[array]	
request_id	请求链路	是	[string]	CCF38623-DE1A-004F-0961-90A6AF265013
response_time	响应时间	是	[string]	2020-04-30 17:33:32
data	响应数据	是	[object]	
data>>product_id	辅料产品id	是	[int]	1010
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "CCF38623-DE1A-004F-0961-90A6AF265013",
    "response_time": "2021-11-11 11:14:33",
    "data": {
        "product_id": 1010
    }
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
查询产品辅料列表
下一章节 
查询产品品牌列表