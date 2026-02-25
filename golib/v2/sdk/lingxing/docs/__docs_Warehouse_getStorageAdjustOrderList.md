查询调整单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAdjustment/getStorageAdjustOrderList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search_date_type	时间类型：
1 创建时间
2 调整时间
3 更新时间	否	[int]	1
start_date	开始日期，格式：Y-m-d	否	[string]	2024-07-30
end_date	结束日期，格式：Y-m-d	否	[string]	2024-07-30
order_sn	调整单号，多个使用英文逗号分隔	否	[string]	AD240730002
adjust_status	单据状态：
5 待提交
10 待调整
20 已完成
30 已删除
121 待审批
122 已驳回	否	[int]	20
wid	系统仓库id，多个使用英文逗号分隔	否	[string]	1469
type	调整类型：
0 数量调整
1 换标调整
2 sku调整	否	[int]	0
page	当前页码，默认1	否	[int]	1
page_size	分页条数，默认20	否	[int]	20
请求示例
{
    "wid": "1469",
    "type": 0,
    "order_sn": "AD240730002",
    "search_date_type": 1,
    "start_date": "2024-07-30",
    "end_date": "2024-07-30",
    "adjust_status": 20,
    "page": 1,
    "page_size": 20
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
request_id	请求链路id	是	[string]	80B5BD25-FD74-62D2-84C1-75A6D838A2CB
response_time	响应时间	是	[string]	2022-07-01 17:17:21
total	查询总数	是	[int]	498
data	响应数据	是	[array]	
data>>order_sn	单据号	是	[string]	AD220624005
data>>wid	仓库id	是	[int]	1
data>>ware_house_name	仓库名称	是	[string]	仓库1
data>>type	调整类型：
0 数量调整
1 换标调整
2 sku调整	是	[int]	1
data>>type_text	调整类型文本	是	[string]	换标调整
data>>status	单据状态：
5 待提交
10 待调整
20 已完成
30 已删除
121 待审批
122 已驳回	是	[int]	20
data>>status_text	单据状态说明	是	[string]	已完成
data>>remark	单据备注	是	[string]	
data>>create_uid	创建人id	是	[int]	540
data>>create_realname	创建人名称	是	[string]	张三
data>>create_time	创建时间	是	[string]	2022-06-24 19:17:22
data>>commit_uid	提交人id	是	[int]	540
data>>commit_realname	提交人名称	是	[string]	张三
data>>commit_time	提交时间	是	[string]	2022-06-24 19:17:22
data>>adjustment_uid	调整人id	是	[int]	540
data>>adjustment_realname	调整人名称	是	[string]	张三
data>>adjustment_time	调整时间	是	[string]	2022-06-24 19:17:22
data>>opt_uid	单据最后操作人id	是	[int]	540
data>>opt_realname	单据最后操作人名称	是	[string]	张三
data>>opt_time	单据最后操作时间	是	[string]	1656069442
data>>increment_time	单据增量时间	是	[string]	
data>>item_list	单据明细列表	是	[array]	
data>>item_list>>sku	本地产品sku	是	[string]	SKU72711C6
data>>item_list>>product_id	本地产品id	是	[int]	2045
data>>item_list>>product_name	品名	是	[string]	[演示数据]SanDisk Extreme microSD UHS-I带适配器
data>>item_list>>pic_url	图片链接	是	[string]	https://XXXX/xxxx.jpg
data>>item_list>>product_remark	产品备注	是	[string]	产品备注
data>>item_list>>fnsku	fnsku	是	[string]	FNAE6BDA8
data>>item_list>>seller_id	店铺id	是	[string]	
data>>item_list>>seller_name	店铺名称	是	[string]	店铺2
data>>item_list>>country_name	店铺所属国家	是	[string]	加拿大
data>>item_list>>adjustment_valid_num	可用调整量	是	[int]	2
data>>item_list>>available_bin_list	出库的可用仓位列表【只有换标调整才有该字段】	是	[array]	
data>>item_list>>available_bin_list>>whb_code	仓位编码	是	[string]	ts_valid
data>>item_list>>available_bin_list>>whb_code_name	仓位编码名称	是	[string]	可用暂存
data>>item_list>>available_bin_list>>product_num	出库数量	是	[int]	2
data>>item_list>>to_sku	调整的sku【只有SKU调整才有该字段】	是	[string]	SKU72711C62
data>>item_list>>to_product_name	调整的品名【只有SKU调整才有该字段】	是	[string]	[演示数据]SanDisk Extreme microSD UHS-I带适配器
data>>item_list>>to_fnsku	调整的fnsku【数量调整没有该字段】	是	[string]	
data>>item_list>>to_seller_id	调整的店铺id【数量调整没有该字段】	是	[string]	1
data>>item_list>>to_seller_name	调整的店铺名称【数量调整没有该字段】	是	[string]	店铺1
data>>item_list>>to_country_name	调整的店铺所属国家【数量调整没有该字段】	是	[string]	美国
data>>item_list>>to_available_bin	【废弃字段】换标的入库仓位【只有换标调整才有该字段】	是	[string]	ts_valid
data>>item_list>>to_available_bin_name	【废弃字段】换标的入库仓位名称【只有换标调整才有该字段】	是	[string]	可用暂存
data>>item_list>>available_bin	可用仓位【只有数量调整有该字段】	是	[string]	ts_valid
data>>item_list>>available_bin_name	可用仓位名称【只有数量调整有该字段】	是	[string]	可用暂存
data>>item_list>>inferior_bin	次品仓位【只有数量调整有该字段】	是	[string]	ts_bad
data>>item_list>>inferior_bin_name	次品仓位名称【只有数量调整有该字段】	是	[string]	次品暂存
data>>item_list>>adjustment_bad_num	次品调整量【只有数量调整有该字段】	是	[int]	2
data>>item_list>>to_available_bin_list	换标的入库仓位列表	是	[array]	
data>>item_list>>to_available_bin_list>>whb_code	仓位编码	是	[string]	T-1-2
data>>item_list>>to_available_bin_list>>whb_code_name	仓位名称	是	[string]	T-1-2
data>>item_list>>to_available_bin_list>>product_num	入库数量	是	[int]	4
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "99927EBF-3866-D31E-EB2B-5129DA597DAD",
    "response_time": "2023-12-21 10:55:16",
    "data": [
        {
            "order_sn": "AD231220001",
            "type": 1,
            "wid": 1,
            "status": 20,
            "remark": "",
            "create_uid": 10317908,
            "create_time": "2023-12-20 15:48:11",
            "commit_uid": 10317908,
            "commit_time": "2023-12-20 15:48:11",
            "adjustment_uid": 10317908,
            "adjustment_time": "2023-12-20 15:48:11",
            "opt_uid": 10317908,
            "opt_time": "2023-12-20 15:48:11",
            "company_id": 901202594714743296,
            "increment_time": "2023-12-20 15:48:11",
            "type_text": "换标调整",
            "ware_house_name": "默认仓库29",
            "status_text": "已完成",
            "create_realname": "张三",
            "commit_realname": "张三",
            "adjustment_realname": "张三",
            "opt_realname": "张三",
            "item_list": [
                {
                    "product_id": 23060,
                    "sku": "12-200",
                    "product_name": "产品1",
                    "pic_url": "",
                    "adjustment_valid_num": 2,
                    "product_remark": "",
                    "fnsku": "",
                    "seller_name": "",
                    "country_name": "",
                    "seller_id": "0",
                    "available_bin_list": [],
                    "to_fnsku": "B0BB389BKQ",
                    "to_seller_name": "8pspeed-11",
                    "to_country_name": "美国",
                    "to_seller_id": "101",
                    "to_available_bin": "",
                    "to_available_bin_name": "",
                    "to_available_bin_list": [
                        {
                            "whb_code": "ts_valid",
                            "whb_code_name": "可用暂存",
                            "product_num": 2
                        }
                    ]
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
撤销调拨单
下一章节 
创建已完成的数量调整单