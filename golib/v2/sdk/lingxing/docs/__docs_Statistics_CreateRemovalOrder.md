创建移除订单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/statistic/removalOrder/createAndCommit	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
lists	提交数据，支持批量，上限100个	是	[array]	
lists>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	5
lists>>order_type	移除类型：2 Return，3 Disposal	是	[int]	2
lists>>country_code	国家code，查询亚马逊市场列表接口对应【code】字段	是	[string]	US
lists>>state_or_region	省州code，查询亚马逊国家下地区列表接口对应【code】字段	是	[string]	MI
lists>>sys_wid	系统仓库id，查询仓库列表接口对应【wid】字段	否	[int]	1
lists>>removal_no	移除订单号，最长10位，为空则自动创建订单号	否	[string]	141hg16
lists>>city	城市	是	[string]	Michigan
lists>>address_line1	地址1	是	[string]	address_line1-test
lists>>address_line2	地址2	否	[string]	address_line1-test
lists>>address_line3	地址3	否	[string]	address_line1-test
lists>>postal_code	邮编	是	[string]	49444
lists>>phone	联系电话	是	[string]	2317982361
lists>>name	地址名称	是	[string]	Michael Morton
lists>>remark	备注	否	[string]	remark demo
lists>>items	items列表	是	[array]	
lists>>items>>msku	msku	是	[string]	Black_ Head_Rope
lists>>items>>sellable_quantity	移除可售数量	是	[int]	1
lists>>items>>unsellable_quantity	移除不可售数量	是	[int]	1
请求示例
{
    "lists": [{
        "sid": 5,
        "order_type": 2,
        "country_code": "US",
        "state_or_region": "MI",
        "sys_wid": 1,
        "removal_no": "141hg16",
        "city": "Michigan",
        "address_line1": "address_line1-test",
        "address_line2": "address_line1-test",
        "address_line3": "address_line1-test",
        "postal_code": "49444",
        "phone": "132xxxxxxxx",
        "name": "Michael Morton",
        "remark": "remark demo",
        "items": [{
            "msku": "Black_xxx",
            "sellable_quantity": 1,
            "unsellable_quantity": 1
        }]
    }]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	EF617C27-34B4-3596-BA2A-75F70F0B053C
response_time	响应时间	是	[string]	2023-04-13 14:38:53
data	响应数据	是	[array]	
total	总数	是	[int]	0
 上一章节
库存报表-FBA-历史报表-汇总-明细
下一章节 
查询亚马逊赔偿报告列表