批量设置规则 - MSKU

支持批量设置MSKU的补货建议规则

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/fbaSug/msku/setConfigs	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
msku_list	msku信息	是	[array]	
msku_list>>sid	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[int]	3
msku_list>>msku	MSKU	是	[string]	XXXXX
days_plan	采购计划时长	是	[string]	2
days_qc	质检时长	是	[string]	3
sm_fba_list	本地仓至FBA时效	是	[array]	
sm_fba_list>>sm_id	运输方式id	是	[string]	241250000631390721
sm_fba_list>>days	天数	是	[string]	25
sm_oversea_list	本地仓至海外仓时效	是	[array]	
sm_oversea_list>>sm_id	运输方式id	是	[string]	241250000631390721
sm_oversea_list>>days	天数	是	[string]	20
days_oversea_to_fba	海外仓至FBA天数	是	[number]	0
days_frequency_purchase	采购频率	是	[number]	0
days_frequency_local_send	本地仓发货频率	是	[number]	0
days_frequency_oversea_send	海外仓发货频率	是	[number]	0
safe_day	安全天数	是	[number]	14
is_ignore_certainly_short	建议量扣除必断货量：0 否，1 是	是	[number]	0
is_ignore_history_out_stock	历史销量排除断货数据：0 否，1 是	是	[number]	0
days_toucheng	已弃用（原本地至FBA天数-海运）	否	[number]	30
days_oversea	已弃用（原本地至海外仓天数-海运）	否	[number]	0
days_toucheng_air	已弃用（原本地至FBA时效-空运）	否	[number]	0
days_oversea_air	已弃用（原本地至海外仓时效-空运）	否	[number]	0
default_type_toucheng	已弃用（原默认头程物流类型）	否	[number]	0
default_type_oversea	已弃用（原默认本地发海外仓物流类型 ）	否	[number]	0
days_frequency	已弃用（原补货频率）	否	[number]	0
config_list	日销量设置	是	[array]	
config_list>>title	规则名称	是	[string]	默认规则
config_list>>is_default	是否默认：0 否，1 是	是	[int]	1
config_list>>type	类型： 0 自定义，1 动态销量	是	[int]	1
config_list>>weigth_3	权重：3天	是	[number]	0
config_list>>weigth_7	权重：7天	是	[number]	50
config_list>>weigth_14	权重：14天	是	[number]	30
config_list>>weigth_30	权重：30天	是	[number]	20
config_list>>weigth_60	权重：60天	是	[number]	0
config_list>>weigth_90	权重：90天	是	[number]	0
config_list>>volume	日销量（四舍五入，保留两位小数）	是	[string]	0.09
config_list>>date_start	开始日期	是	[string]	2023-11-01
config_list>>date_end	结束日期	是	[string]	2023-11-30
denoise_list	日销量去噪设置	是	[array]	
denoise_list>>title	名称	是	[string]	规则1
denoise_list>>date_start	配置起始日期	是	[string]	2023-11-01
denoise_list>>date_end	配置结束日期	是	[string]	2023-11-30
denoise_list>>type	类型：1 固定值去噪，2 百分比去噪	是	[int]	
denoise_list>>percent	去噪百分比	是	[number]	0
denoise_list>>volume	日销量	是	[number]	10
请求示例
{
    "msku_list": [
        {
            "sid": 3,
            "msku": "XXXXX"
        }
    ],
    "days_plan": "2",
    "days_qc": "3",
    "sm_fba_list": [
        {
            "sm_id": "241250000631390721",
            "days": "25"
        }
    ],
    "sm_oversea_list": [
        {
            "sm_id": "241250000631390721",
            "days": "20"
        }
    ],
    "days_oversea_to_fba": 0,
    "days_frequency_purchase": 0,
    "days_frequency_local_send": 0,
    "days_frequency_oversea_send": 0,
    "safe_day": 14,
    "is_ignore_certainly_short": 0,
    "is_ignore_history_out_stock": 0,
    "config_list": [
        {
            "title": "默认规则",
            "is_default": 1,
            "type": 1,
            "weigth_3": 0,
            "weigth_7": 50,
            "weigth_14": 30,
            "weigth_30": 20,
            "weigth_60": 0,
            "weigth_90": 0,
            "volume": "0.09",
            "date_start": "2023-11-01",
            "date_end": "2023-11-30"
        }
    ],
    "denoise_list": [
        {
            "title": "规则1",
            "date_start": "2023-11-01",
            "date_end": "2023-11-30",
            "type": 1,
            "percent": 0,
            "volume": 10
        }
    ]
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
error_details>>sid	店铺id	是	[string]	4
error_details>>msku	msku	是	[string]	xxxx
error_details>>error	错误提示	是	[string]	补货建议不存在
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
data	响应数据	是	[object]	
total	总数	是	[int]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "A5992A57-05DF-8A74-55BE-EC5689F44964",
    "response_time": "2021-11-11 17:53:36",
    "data": {},
    "total": 0
}
复制
错误
复制成功
 上一章节
批量设置规则-ASIN
下一章节 
查询报表型数据明细-ASIN