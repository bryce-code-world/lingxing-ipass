查询多属性产品列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/spu/spuList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量	是	[int]	0
length	分页长度，上限200	是	[int]	20
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
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	68E5504B-B027-72D4-336B-0A59825AC259
response_time	响应时间	是	[string]	2022-09-13 15:04:37
total	总数	是	[int]	1
data	响应数据	是	[array]	
data>>ps_id	SPU 唯一id	是	[int]	
data>>spu	SPU	是	[string]	spu-001
data>>spu_name	款名	是	[string]	spu-001
data>>model	型号	是	[string]	1
data>>cid	分类id	是	[int]	1275
data>>bid	品牌id	是	[int]	1012
data>>developer_uid	开发人id	是	[int]	1
data>>cg_uid	采购员id	是	[int]	1
data>>purchase_remark	采购备注	是	[string]	xx
data>>cg_price	采购成本	是	[string]	11.000000
data>>cg_delivery	交期	是	[int]	1
data>>create_uid	创建人id	是	[string]	1
data>>create_time	创建时间	是	[string]	2022-01-30 00:02:00
data>>status	状态：
0 停售，
1 在售，
2 开发中，
3 清仓	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "68E5504B-B027-72D4-336B-0A59825AC259",
    "response_time": "2022-09-13 15:04:37",
    "data": [
        {
            "ps_id": 1,
            "spu": "spu-0001",
            "spu_name": "spu-0001",
            "model": "1",
            "cid": 1275,
            "bid": 1012,
            "developer_uid": 10397585,
            "cg_uid": 10317902,
            "cg_price": "11.000000",
            "purchase_remark": "1111",
            "cg_delivery": 1,
            "create_uid": 10319730,
            "create_time": "2022-07-14 17:47:18",
            "status": 1
        }
    ],
    "total": 66
}
复制
错误
复制成功
 上一章节
添加/编辑产品属性
下一章节 
查询多属性产品详情