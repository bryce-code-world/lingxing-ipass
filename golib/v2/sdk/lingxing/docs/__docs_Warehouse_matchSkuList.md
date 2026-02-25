查询系统产品与第三方海外仓产品映射列表

支持查询本地产品与第三方海外仓产品映射列表

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/owms/inbound/matchSkuList	HTTPS	POST	1
请求参数

Json Object

参数名	说明	必填	类型	示例
wid	仓库id，多个用英文逗号分隔	是	[string]	1,578,765
is_matched	是否配对：【空表示都返回】
0 未配对
1 配对	否	[int]	1
offset	分页偏移量	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
{
    "wid": "1,578,765",
    "is_matched": 1,
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
error_details	错误消息	是	[array]	
request_id	请求链路id	是	[string]	7E80FC0B-55C1-BAC5-0E4B-7AE59706F051
response_time	响应时间	是	[string]	2021-06-15 20:16:38
data	响应数据	是	[array]	
data>>warehouse_name	仓库名称	是	[string]	
data>>warehouse_code	仓库代码code	是	[string]	
data>>wid	仓库id	是	[int]	
data>>local_sku	本地sku编码	是	[string]	
data>>local_name	本地sku名称	是	[string]	
data>>match_num	配对数量，默认为1	是	[int]	
data>>oversea_product_name	三方产品名称	是	[string]	
data>>oversea_unique_code	三方产品唯一code，可能为空	是	[string]	
data>>oversea_product_code	三方sku编码	是	[string]	
data>>is_matched	是否已配对：
false 未配对
true 已配对	是	[boolean]	
data>>is_matched_text	对is_matched的文本展示	是	[string]	
data>>twp_id	三方产品id	是	[int]	
data>>sid	店铺id 库存中心过渡版本之后返回	是	[int]	
total		是	[int]	 
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "B5C43D09-0EAC-FCCF-D88F-B638CFCA4D24",
    "response_time": "2024-08-01 16:22:17",
    "data": [
        {
            "wpm_id": 475,
            "tw_id": 135,
            "warehouse_name": "美国仓库",
            "warehouse_name_local": "赛诺仓 美国仓1",
            "wid": 765,
            "warehouse_code": "US01",
            "twp_id": 3428,
            "match_num": 1,
            "match_third_status": null,
            "match_msg": null,
            "oversea_product_name": "1+5",
            "oversea_unique_code": "1+5",
            "oversea_product_code": "1+5",
            "oversea_spec": "",
            "local_sku": "00022",
            "local_name": "手套",
            "product_id": 17712,
            "fnsku": "X002Y7Q47L",
            "seller_id": 101,
            "seller_name": "Ak-quanqiu-JP-US2&3",
            "country_name": "美国",
            "is_matched": true,
            "sid": 101,
            "is_matched_text": "已配对"
        }
    ],
    "total": 22
}
复制
错误
复制成功
 上一章节
创建待发货/待收货/已完成的备货单
下一章节 
查询海外仓sku配对列表