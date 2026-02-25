VC订单-打印标签【DF】
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformOrder/vcOrderDf/getShippingLabel	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
ids	订单ID，查询VC订单列表接口对应字段【id】	是	[array]	["107"]
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[object]	
data>>label_list	标签数据	是	[array]	
data>>label_list>>id	订单ID	是	[string]	107
data>>label_list>>purchase_order_number	订单编号	是	[string]	XqW90wh9r
data>>label_list>>label_count	标签数量	是	[int]	0
data>>label_list>>error_msg	错误信息	是	[string]	请求亚马逊异常,请检查店铺授权情况
data>>pdf_url	PDF下载链接	是	[string]	https://www.xxx.com/xxx
data>>download_url	压缩包下载链接	是	[string]	https://www.xxx.com/xxx
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "114efe4e0455446da7008f236e5d6110.1698318663923",
    "response_time": "2023-10-26 19:11:07",
    "data": {
        "label_list": [
            {
                "id": "107",
                "purchase_order_number": "XqW90wh9r",
                "label_count": 0,
                "error_msg": "请求亚马逊异常,请检查店铺授权情况"
            }
        ],
        "pdf_url": "",
        "download_url": ""
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
VC订单-请求标签【DF】
下一章节 
配对/批量配对