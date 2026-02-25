查询VC店铺列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/platformAuth/vcSeller/pageList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20，上限200	否	[int]	20
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	C3D9F541-8083-E376-EB5C-606A872F5C89
response_time	响应时间	是	[string]	2022-12-08 18:27:13
total	总数	是	[int]	0
data	响应数据	是	[array]	
data>>account_id	账号ID	是	[int]	1
data>>seller_id	SELLER_ID	是	[string]	TEST12345
data>>account_name	账号名称	是	[string]	VC01
data>>region	站点简称	是	[string]	na
data>>region_name	站点名称	是	[string]	北美洲
data>>vc_store_id	VC店铺id	是	[string]	134225003201380864
data>>name	店铺名称	是	[string]	VC01-美国
data>>status	店铺授权服务状态：
-1 删除
0 暂停同步
1 正常同步
2 授权异常
	是	[int]	1
data>>mid	站点id	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "491b9227a802465a88077a54f4be02d4.1697714035725",
    "response_time": "2023-10-19 19:13:56",
    "data": [
        {
            "account_id": 1,
            "seller_id": "TEST12345",
            "account_name": "VC01",
            "region": "na",
            "region_name": "北美洲",
            "vc_store_id": "134225003201380864",
            "name": "VC01-美国",
            "status": 1,
            "mid": 1
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询平台订单列表
下一章节 
查询VC-Listing列表