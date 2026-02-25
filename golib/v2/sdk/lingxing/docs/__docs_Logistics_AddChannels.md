批量添加头程物流方式

支持批量添加头程物流方式到领星ERP，对应系统功能【物流】>【头程物流】>【物流渠道】

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/tms/FirstVessel/addChannels	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
channelsData	头程物流方式数据，每次请求限制20条	是	[array]	
channelsData>>channel_name	头程物流方式名称	否	[string]	
channelsData>>volume_calc_param	材积计算参数	否	[string]	
channelsData>>zip_code	邮编	是	[int]	
channelsData>>valid_period	时效天数	是	[string]	
channelsData>>remark	备注	是	[string]	
channelsData>>billing_type	计费类型：0 重量，1 体积	否	[int]	
channelsData>>logistics_provider_id	所属头程物流商id	否	[string]	
channelsData>>billing	运费信息，格式：【注意逗号使用英文逗号，多条运费以竖线分隔】
重量范围开始(kg),重量范围结束(kg),价格(元/kg)	否	[string]	 
请求示例
{
    "channelsData": [{
        "channel_name": "test-0826",
        "volume_calc_param": 5000,
        "zip_code": "",
        "valid_period": 1,
        "remark": "",
        "billing_type": 0,
        "logistics_provider_id": "36",
        "billing": "1,10,2|11,15,2.8"
    }]
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
request_id	请求链路id	是	[string]	FFE89896-B9C5-1634-5589-6E0F89F0E4A2
response_time	响应时间	是	[string]	2020-04-30 17:33:32
data	响应数据	是	[array]	 
data>>id	物流方式对应的id	是	[int]	
total	总数	是	[int]	 
 上一章节
批量添加头程物流商
下一章节 
查询已启用的自发货物流方式