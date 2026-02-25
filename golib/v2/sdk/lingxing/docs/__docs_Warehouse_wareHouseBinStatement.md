查询仓位流水
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/wareHouseBinStatement	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
wid	仓库ID，多个仓库ID用英文逗号,分隔，传或者传空则默认所有仓库	否	[string]	1,578,765
type	流水类型：【多个流水类型用英文逗号分隔，不填默认全部类型】
16 换标入库
17 加工入库
18 拆分入库
19 其他入库
22 采购入库
23 委外入库
24 调拨入库
25 盘盈入库
26 退货入库
27 移除入库
28 采购质检
29 委外质检
32 委外出库
33 盘亏出库
34 换标出库
35 加工出库
36 拆分出库
37 FBA出库
38 FBM出库
39 退货出库
41 调拨出库
42 其他出库
65 WFS出库
71 采购上架
72 委外上架
100 库存调整
200 成本补录
30001 已撤销	否	[string]	19
bin_type_list	仓位类型：【多个类型用逗号分隔】
1 待检暂存
2 可用暂存
3 次品暂存
4 拣货暂存
5 可用
6 次品	否	[string]	5
start_date	操作开始时间，Y-m-d，闭区间，联合结束时间使用	否	[string]	2022-01-30
end_date	操作结束时间，Y-m-d，开区间，联合开始时间使用	否	[string]	2024-06-01
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认20	否	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "wids": "1,578,765",
    "type": "19",
    "bin_type_list": "5",
    "start_date": "2022-01-30",
    "end_date": "2024-06-01"
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
request_id	请求链路id	是	[string]	502B9DD9-1BA0-03C5-6C61-D77C830440A6
response_time	响应时间	是	[string]	2020-05-18 11:23:47
data	响应时间	是	[array]	
data>>wid	仓库ID	是	[int]	8
data>>ware_house_name	仓库名	是	[string]	测试仓库
data>>whb_id	仓位id	是	[string]	
data>>whb_name	仓位名称	是	[string]	
data>>whb_type_name	仓位类型名称	是	[string]	
data>>order_sn	单据号	是	[string]	OB220527028
data>>product_id	商品ID	是	[int]	18232
data>>product_name	品名	是	[string]	AA-组合3
data>>sku	SKU	是	[string]	AA-组合3
data>>fnsku	FNSKU	是	[string]	
data>>num	数量	是	[int]	
data>>type	流水类型	是	[int]	
data>>remark	备注	是	[string]	库存初始化
data>>opt_uid	操作人员ID	是	[int]	230
data>>opt_time	操作时间	是	[string]	2020-09-05 16:23
data>>type_text	流水类型文本	是	[string]	其他入库
data>>opt_realname	操作人员姓名	是	[string]	李X
total	总数目	是	[int]	143643
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "3BA348FF-0C5D-6E7E-BEAE-70873B219122",
    "response_time": "2024-07-31 15:46:18",
    "data": [
        {
            "wid": 1,
            "ware_house_name": "wcn测试仓库测试改名",
            "whb_id": 594,
            "whb_name": "A-1-2",
            "whb_type_name": "可用",
            "order_sn": "IB220829005",
            "product_id": 10001,
            "product_name": "0.0MDF/WOOD/MTL SIGN/CLEAN SINGLE.../7-7/8\\\"",
            "sku": "ceshi001",
            "seller_id": "0",
            "fnsku": "",
            "num": 100,
            "type": 19,
            "type_text": "其他入库",
            "opt_uid": 273,
            "opt_time": "2022-08-29 14:56",
            "opt_realname": "曾卓彬",
            "remark": ""
        }
    ],
    "total": 3471
}
复制
错误
复制成功
 上一章节
查询库存流水（新）
下一章节 
查询收货单列表