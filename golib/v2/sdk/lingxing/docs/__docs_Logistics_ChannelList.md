查询头程物流渠道列表

对应系统【物流】-》【头程物流】-》【物流渠道】数据

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/local_inventory/channelList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	是	[int]	0
length	分页长度	是	[int]	20
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
request_id	请求链路id	是	[string]	C4B16647-7899-6178-F08A-79B20AE0E7F7
response_time	响应时间	是	[string]	2021-03-23 16:55:23
total	总数	是	[int]	1
data	响应数据	是	[array]	
data>>id	物流渠道id【对应ERP页面“物流方案代码”】	是	[int]	10000405
data>>channel_name	物流渠道	是	[string]	测试7777
data>>method_id	运输方式id	是	[string]	132456
data>>method_name	运输方式名称	是	[string]	海派
data>>billing_type	计费类型：0 计费重，1 体积	是	[int]	0
data>>volume_calc_param	材积计算参数	是	[string]	154
data>>zip_code	邮编	是	[string]	323
data>>valid_period	时效天数	是	[int]	3
data>>remark	备注	是	[string]	
data>>enabled	状态：0 停用、1 启用	是	[int]	1
data>>last_modify_uid	最后更新数据用户id	是	[int]	626
data>>gmt_modified	更新时间	是	[string]	2022-12-09 17:16:46
data>>provider	物流商信息	是	[object]	
data>>provider>>id	所属头程物流商id	是	[string]	69
data>>provider>>logistics_provider_name	所属头程物流商名称	是	[string]	默认物流商1
data>>freight	运费规则	是	[array]	
data>>freight>>country_code	国家	是	[string]	US
data>>freight>>region_code	分区	是	[string]	EAST
data>>freight>>billing_weight_start	计费重量范围开始	是	[number]	21
data>>freight>>billing_price	计费价格	是	[number]	41
data>>send_place_code	提货地代码	是	[string]	82
data>>receive_country_code	目的国家二字码	是	[string]	IT
data>>is_include_tax	是否包税：0 否，1 是	是	[int]	0
data>>is_points_behind	是否分抛：0 否，1 是	是	[int]	1
data>>points_behind_coeffient	分抛系数，不带%	是	[number]	10
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C86C5A38-E441-3C6A-9D68-BCD59CBB9183",
    "response_time": "2023-07-06 18:43:46",
    "data": [
        {
            "id": "1",
            "channel_name": "测试7777",
            "billing_type": 0,
            "volume_calc_param": 5000,
            "zip_code": "efef",
            "valid_period": 0,
            "remark": "1111111",
            "last_modify_uid": 10317908,
            "gmt_modified": "2023-04-21 17:44:28",
            "provider": {
                "id": "3",
                "logistics_provider_name": "test10"
            },
            "freight": [
                {
                    "country_code": "",
                    "region_code": "",
                    "billing_weight_start": "0.00",
                    "billing_price": "5.00"
                }
            ],
            "send_place_codes": [],
            "receive_country_codes": [],
            "is_include_tax": 0,
            "is_points_behind": 0,
            "points_behind_coeffient": 0
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
销售出库单截单
下一章节 
查询头程物流商列表