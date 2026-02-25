查询采购变更单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/purchase/purchaseChangeOrder/changeOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search_field_time	筛选时间类型，创建时间:create_time, 更新时间：update_time，不填时默认创建时间	否	[string]	create_time
start_date	开始时间	否	[string]	2024-08-02
end_date	结束时间	否	[string]	2024-08-02
offset	分页偏移量	是	[int]	0
length	分页长度	是	[int]	20
multi_search_field	搜索单号字段，变更单号：order_sn；采购单号：purchase_order_sn	否	[string]	purchase_order_sn
multi_search_value	批量搜索的单号值	否	[array]	["test-01","test-02"]
请求示例
{
    "search_field_time":"create_time",
    "offset": 0,
    "length": 20,
    "start_date": "2024-08-02",
    "end_date": "2024-08-02",
    "multi_search_field":"purchase_order_sn",
    "multi_search_value":["test-01","test-02"]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	响应消息	是	[string]	操作成功
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8DFEE1CA-EC9B-F401-5D41-9F95251D5D50
response_time	响应时间	是	[string]	2020-09-21 15:48:58
data	数据	是	[array]	
data>>total	总数	是	[int]	338
data>>list	列表数据	是	[array]	
data>>list>>order_sn	变更单号	是	[string]	POC220505001
data>>list>>create_time	创建时间	是	[string]	2022-05-05 15:08:43
data>>list>>supplier_name	供应商	是	[string]	默认供应商
data>>list>>old_supplier_name	旧供应商	是	[string]	默认旧供应商
data>>list>>wid	仓库id	是	[string]	1
data>>list>>old_wid	旧仓库id	是	[string]	1
data>>list>>ware_house_name	仓库	是	[string]	fq本地仓库1
data>>list>>old_ware_house_name	旧仓库	是	[string]	旧仓库
data>>list>>create_realname	创建人	是	[string]	冯x
data>>list>>opt_realname	采购员	是	[string]	安安
data>>list>>remark	备注	是	[string]	
data>>list>>status	状态标识码：
-1 已驳回
0 待审核
1 已处理	是	[int]	1
data>>list>>status_text	状态文本	是	[string]	已处理
data>>list>>icon	货币符号	是	[string]	￥
data>>list>>amount	金额	是	[number]	
data>>list>>old_amount	旧金额	是	[number]	
data>>list>>item_list	变更单商品子项	是	[array]	
data>>list>>item_list>>wid	仓库id	是	[array]	
data>>list>>item_list>>old_wid	旧仓库id	是	[array]	
data>>list>>item_list>>ware_house_name	仓库	是	[string]	本地仓库
data>>list>>item_list>>old_ware_house_name	旧仓库	是	[string]	旧仓库
data>>list>>item_list>>spu	spu	是	[string]	
data>>list>>item_list>>spu_name	spu名称	是	[string]	
data>>list>>item_list>>product_name	品名	是	[string]	木牛流马
data>>list>>item_list>>product_id	产品id	是	[string]	
data>>list>>item_list>>sku	SKU	是	[string]	mu123
data>>list>>item_list>>attribute	属性信息	是	[string]	
data>>list>>item_list>>attribute>>attr_id	属性id	是	[int]	
data>>list>>item_list>>attribute>>attr_name	属性名	是	[string]	
data>>list>>item_list>>attribute>>attr_value	属性值	是	[string]	
data>>list>>item_list>>fnsku	FNSKU	是	[string]	X002JICY5VCR
data>>list>>item_list>>price	含税单价	是	[string]	0.00
data>>list>>item_list>>old_price	旧含税单价	是	[string]	0.00
data>>list>>item_list>>price_without_tax	不含税单价	是	[number]	
data>>list>>item_list>>old_price_without_tax	旧不含税单价	是	[number]	
data>>list>>item_list>>tax_rate	税率	是	[number]	
data>>list>>item_list>>old_tax_rate	旧税率	是	[number]	
data>>list>>item_list>>quantity_real	实际采购量	是	[number]	250
data>>list>>item_list>>old_quantity_real	旧实际采购量	是	[number]	200
data>>list>>item_list>>amount	金额	是	[number]	0.00
data>>list>>item_list>>old_amount	旧金额	是	[number]	0.00
data>>list>>item_list>>seller	店铺列表	是	[array]	
data>>list>>item_list>>seller>>mid	站点id	是	[int]	
data>>list>>item_list>>seller>>sid	店铺id	是	[int]	
data>>list>>item_list>>seller>>seller_name	店铺名称	是	[string]	
data>>list>>item_list>>seller>>country_name	国家	是	[string]	
data>>list>>item_list>>msku	msku	是	[array]	
data>>list>>item_list>>is_tax	是否含税	是	[int]	1
data>>list>>item_list>>is_aux	是否辅料	是	[int]	0
data>>list>>item_list>>expect_arrive_time	预计到货时间	是	[string]	
data>>list>>item_list>>old_expect_arrive_time	原预计到货时间	是	[string]
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "AABB891A-8A17-94C0-4DBA-0399C4641320",
    "response_time": "2022-05-26 10:54:41",
    "data": [
        {
            "order_sn": "POC230524011",
            "create_time": "2023-05-24 15:35:26",
            "supplier_name": "A0516供应商115",
            "old_supplier_name": "A0516供应商115",
            "wid": 1,
            "old_wid": 1,
            "ware_house_name": "默认仓库",
            "old_ware_house_name": "默认仓库",
            "create_realname": "超级管理员",
            "opt_realname": "超级管理员",
            "old_opt_realname": "超级管理员",
            "remark": "备注xx",
            "status": 1,
            "status_text": "已处理",
            "icon": "￥",
            "amount": "200.00",
            "old_amount": "100.00",
            "item_list": [
                {
                    "wid": 1,
                    "old_wid": 1,
                    "ware_house_name": "默认仓库",
                    "old_ware_house_name": "默认仓库",
                    "spu": "140141010",
                    "spu_name": "770",
                    "product_name": "调拨单删除",
                    "product_id": "17797",
                    "sku": "0007-1",
                    "attribute": [
                        {
                            "attr_id": 590,
                            "attr_name": "ZDH属性",
                            "attr_value": "ZDH属性值"
                        }
                    ],
                    "is_aux": 0,
                    "seller": [],
                    "msku": [],
                    "fnsku": "",
                    "price": "10.0000",
                    "old_price": "10.0000",
                    "price_without_tax": "8.8496",
                    "old_price_without_tax": "8.8496",
                    "tax_rate": "13.00",
                    "old_tax_rate": "13.00",
                    "amount": "200.00",
                    "old_amount": "100.00",
                    "is_tax": 1,
                    "quantity_real": 20,
                    "old_quantity_real": 10
                }
            ]
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
作废采购/委外退货单
下一章节 
创建已完成的采购变更单