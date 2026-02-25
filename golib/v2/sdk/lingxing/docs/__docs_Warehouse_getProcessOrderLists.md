加工单列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageProcess/getOrderLists	HTTPS	POST	1
请求参数
参数名	说明	类型	必填	示例
type	单据类型：1加工单，2拆分单	[int]	是	1
wid	仓库id，多个用英文逗号分隔	[string]	否	53
process_sn	加工单号，多个用英文逗号分隔	[string]	否	PT240730001
status	加工状态：
0 待配货
1 待完成
2 已完成	[int]	否	2
search_field_time	时间搜索维度：
create_time 创建时间
finish_time 完成时间
update_time 更新时间	[string]	否	create_time
start_date	开始时间，格式：Y-m-d	[string]	否	2024-07-30
end_date	结束时间，格式：Y-m-d	[string]	否	2024-07-30
offset	分页偏移量，默认0	[int]	是	0
length	分页长度，默认500	[int]	是	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wid": 53,
    "type": 1,
    "process_sn": "PT240730001",
    "status": 2,
    "search_field_time": "create_time",
    "start_date": "2024-07-30",
    "end_date": "2024-07-30"
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	结果	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	160B52C9-2B20-2A28-75E9-548C96A90565
response_time	响应时间	是	[string]	2023-03-28 09:59:00
total	单据总数	是	[int]	200
data	响应数据	是	[array]	
data>>process_sn	加工/拆分单号	是	[string]	PT230324003
data>>status	订单状态：
0 待配货
1 待完成
2 已完成	是	[int]	2
data>>type	单据类型：
1 加工单
2 拆分单	是	[string]	1
data>>ware_house_name	仓库名称	是	[string]	默认仓库
data>>wid	仓库id	是	[int]	1
data>>create_by	创建人id	是	[int]	111
data>>create_realname	创建人名称	是	[string]	验证账户2
data>>create_time	创建时间	是	[string]	2023-03-24 00:55:36
data>>finish_realname	最后操作人	是	[string]	验证账户2
data>>finish_time	最后操作时间	是	[string]	2023-03-26 22:46:51
data>>finish_uid	最后操作人id	是	[int]	111
data>>remark	备注	是	[string]	备注
data>>update_time	更新时间	是	[string]	2023-03-24 00:55:36
data>>product_list	组合品项	是	[array]	
data>>product_list>>sku	组合品sku	是	[string]	niubao
data>>product_list>>product_name	组合品产品名称	是	[string]	早餐
data>>product_list>>fnsku	组合品fnksu	是	[string]	fnsku
data>>product_list>>seller_id	组合品店铺id	是	[int]	3
data>>product_list>>seller_name	组合品店铺名称	是	[string]	店铺1
data>>product_list>>process_fee	加工费	是	[string]	0.000000
data>>product_list>>remark	备注	是	[string]	默认备注
data>>product_list>>whb_code_good	加工单组合品入库仓位	是	[array]	[]
data>>product_list>>whb_code_good>>whb_code	仓位编码	否	[string]	ts_valid
data>>product_list>>whb_code_good>>whb_code_name	仓位名	否	[string]	可用暂存
data>>product_list>>whb_code_good>>whb_num	数量	否	[int]	1
data>>product_list>>pic_url	组合品图片	是	[string]	https://image.distxxx.com/xxx/9dc0c077bc5c2.jpg
data>>product_list>>quantity	加工量/拆分量	是	[int]	1
data>>product_list>>item_list	单品明细项	是	[array]	
data>>product_list>>item_list>>sku	单品sku	是	[string]	sku1
data>>product_list>>item_list>>product_name	单品名称	是	[string]	鸡蛋
data>>product_list>>item_list>>fnsku	单品fnsku	是	[string]	fnsku
data>>product_list>>item_list>>quantity_process	单品加工量/拆分量	是	[int]	2
data>>product_list>>item_list>>seller_id	单品店铺id	是	[int]	33
data>>product_list>>item_list>>seller_name	单品店铺名	是	[string]	店铺1
data>>product_list>>item_list>>whb_code_good	拆分单品入库仓位	是	[array]	[]
data>>product_list>>item_list>>whb_code_good>>whb_code	仓位编码	否	[string]	ts_valid
data>>product_list>>item_list>>whb_code_good>>whb_code_name	仓位名	否	[string]	可用暂存
data>>product_list>>item_list>>whb_code_good>>whb_num	数量	否	[int]	1
返回成功示例
{
    "code":0,
    "message":"success",
    "error_details":[],
    "request_id":"160B52C9-2B20-2A28-75E9-548C96A90565",
    "response_time":"2023-03-28 09:59:00",
    "data":[
        {
            "process_sn":"PT230324003",
            "wid":1,
            "remark":"",
            "status":2,
            "type":1,
            "create_by":10317908,
            "create_time":"2023-03-24 00:55:36",
            "finish_uid":10317908,
            "finish_time":"2023-03-26 22:46:51",
            "create_realname":"验证账户2",
            "finish_realname":"验证账户2",
            "ware_house_name":"默认仓库",
            "product_list":[
                {
                    "sku":"niubao",
                    "fnsku":"",
                    "quantity":44,
                    "remark":"",
                    "process_fee":"0.000000",
                    "seller_id":0,
                    "pic_url":"https://image.distxxx.com/xxx/9dc0c077bc5c2.jpg",
                    "product_name":"早餐套餐",
                    "seller_name":"",
                    "whb_code_good":[],
                    "item_list":[
                        {
                            "fnsku":"",
                            "quantity_process":44,
                            "seller_id":0,
                            "product_name":"肉包",
                            "sku":"baozi",
                            "seller_name":"",
                            "whb_code_good":[]
                        },
                        {
                            "fnsku":"",
                            "quantity_process":44,
                            "seller_id":0,
                            "product_name":"纯牛奶",
                            "sku":"niunai",
                            "seller_name":"",
                            "whb_code_good":[]
                        }
                    ]
                }
            ]
        }
    ],
    "total":999
}
复制
错误
复制成功
 上一章节
创建加工单 / 拆分单
下一章节 
查询调拨单列表