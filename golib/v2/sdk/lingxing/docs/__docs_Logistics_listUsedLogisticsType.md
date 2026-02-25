查询已启用的自发货物流方式

支持查询已启用的自发货物流方式，对应系统【物流】>【物流管理】当中的 API 物流、三方仓物流、平台物流、自定义物流列表

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/wms/WmsLogistics/listUsedLogisticsType	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
param	查询条件	是	[object]	
param>>provider_type	物流商类型：
0 API物流
1 自定义物流
2 海外仓物流
4 平台物流	是	[int]	0
param>>page	分页页码	否	[int]	1
param>>length	分页长度	否	[int]	20
请求示例
{
    "param": 
        {
            "provider_type": 2,
            "page": 1,
            "length": 30
        }

}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[string]	
request_id	请求链路id	是	[string]	C4B16647-7899-6178-F08A-79B20AE0E7F7
response_time	响应时间	是	[string]	2021-03-23 16:55:23
total	总数	是	[int]	1
data	响应数据	是	[array]	
data>>type	物流商类型	是	[int]	
data>>logistics_provider_id	物流商id	是	[int]	
data>>logistics_provider_name	物流商名称	是	[string]	
data>>type_id	物流方式id	是	[int]	
data>>name	物流方式名称	是	[string]	
data>>is_used	物流渠道是否启用：0 停用，1 启用	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "52EABD11-8DBC-4F0C-58C7-28F36676BF39",
    "response_time": "2024-08-05 11:49:26",
    "data": [
        {
            "type_id": "3054",
            "name": "中邮上海挂号小包",
            "code": "166",
            "is_used": 1,
            "logistics_provider_id": "2",
            "fee_template_id": 0,
            "updated_at": 1720576981,
            "address_list": {
                "sender_address": {
                    "country_code": "CN",
                    "country_name": "中国 内地",
                    "province": "广东",
                    "city": "深圳",
                    "district": "南山",
                    "address1": "创智云城A7栋2101",
                    "postcode": "81500",
                    "contact": "JT",
                    "telephone": "400-600-880",
                    "company": "AK",
                    "email": "JT@JT.COM",
                    "post_code": "81500"
                },
                "get_address": null,
                "return_address": null
            },
            "warehouse_type": 1,
            "logistics_provider_name": "燕文",
            "provider_is_used": 1,
            "supplier_id": 3,
            "type": 0,
            "is_platform_provider": 0,
            "supplier_code": null,
            "is_support_domestic_provider": false,
            "is_need_marking": true,
            "is_combine_channel": false
        },
        {
            "type_id": "203450729659600896",
            "is_used": 1,
            "name": "无需物流服务",
            "code": "No_Shipping_Service",
            "logistics_provider_id": "203449733648203776",
            "order_type": 3,
            "fee_template_id": 241461475853718016,
            "billing_type": 0,
            "volume_param": 5000,
            "warehouse_type": 1,
            "logistics_provider_name": "星链（咚咚） 咚咚美西仓",
            "provider_is_used": 1,
            "is_platform_provider": 0,
            "supplier_code": "Xl",
            "wp_code": "Xl",
            "type": 2,
            "wid": 1582,
            "is_support_domestic_provider": false,
            "is_need_marking": true,
            "is_combine_channel": false
        }
    ],
    "total": 460
}
复制
错误
复制成功
 上一章节
批量添加头程物流方式
下一章节 
查询运输方式列表