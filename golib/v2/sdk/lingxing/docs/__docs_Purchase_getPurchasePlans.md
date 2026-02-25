查询采购计划列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/data/local_inventory/getPurchasePlans	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search_field_time	时间搜索维度：
creator_time 创建时间
expect_arrive_time 预计到货时间
update_time 更新时间	是	[string]	creator_time
start_date	开始日期，Y-m-d，闭区间，当筛选update_time时，格式为：Y-m-d H:i:s	是	[string]	2024-07-30
end_date	结束日期，Y-m-d，闭区间，当筛选update_time时，格式为：Y-m-d H:i:s	是	[string]	2024-08-02
plan_sns	采购计划编号	否	[array]	["PP240730001","PP240730002","PP240730005"]
is_combo	是否为组合商品：0 否，1 是	否	[int]	0
is_related_process_plan	是否关联加工计划，0：否，1：是	否	[int]	0
status	状态：
2 待采购
-2 已完成
121 待审批
122 已驳回
-3、124 已作废	否	[array]	[-2]
sids	店铺id ，对应查询亚马逊店铺列表接口对应字段【sid】	否	[array]	[0]
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认500，上限500	否	[int]	20
请求示例
{
{
    "search_field_time": "creator_time",
    "start_date": "2024-07-30",
    "end_date": "2024-08-02",
    "plan_sns": [
        "PP240730001",
        "PP240730002",
        "PP240730005"
    ],
    "is_combo": 0,
    "is_related_process_plan": 0,
    "status": [
        -2
    ],
    "sids": [
        0
    ],
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
error_details	错误信息	是	[string]	
request_id	请求链路id	是	[string]	350C7176-A759-3517-9CDA-369C9E1A5743
response_time	响应时间	是	[string]	2022-05-20 11:48:25
total	总数	是	[int]	1
data	响应数据	是	[array]	
data>>plan_sn	采购计划编号	是	[string]	PP230712003
data>>ppg_sn	采购计划批次号	是	[string]	PPG230712003
data>>status_text	状态说明	是	[string]	待采购
data>>status	状态值	是	[int]	2
data>>creator_real_name	创建人名称	是	[string]	小明
data>>creator_uid	创建人id	是	[int]	297
data>>create_time	创建时间	是	[string]	2021-10-11 01:20:20
data>>file	附件	是	[array]	
data>>file>>name	附件名	是	[string]	导入计划采购商品模板news.xlsx
data>>file>>url	附件url	是	[string]	https://imagexxxx/d73aaf35f.xlsx
data>>plan_remark	备注	是	[string]	备注
data>>pic_url	产品图片	是	[string]	https://image-1254213xxxxx.com/xxx4%B5%E7%9B%92.jpg
data>>spu_name	款名	是	[string]	童黄土-橘白条
data>>spu	SPU	是	[string]	
data>>product_name	品名	是	[string]	[演示数据]三星真无线耳塞含无线充电盒
data>>product_id	商品id	是	[int]	1019
data>>sku	SKU	是	[string]	SKU627B4F6
data>>attribute	属性	是	[array]	
data>>sid	店铺id	是	[number]	1
data>>seller_name	店铺名称	是	[string]	店铺1
data>>marketplace	国家	是	[string]	美国
data>>fnsku	FNSKU	是	[string]	FN004
data>>msku	MSKU	是	[array]	
data>>supplier_id	供应商id	是	[object]	2
data>>supplier_name	供应商名称	是	[string]	测试供应商
data>>wid	仓库id	是	[int]	2
data>>warehouse_name	仓库名称	是	[string]	广州仓库
data>>purchaser_id	采购方id	是	[int]	1
data>>purchaser_name	采购方名称	是	[string]	默认采购方
data>>cg_box_pcs	单箱数量	是	[int]	10
data>>quantity_plan	计划采购量	是	[int]	100
data>>expect_arrive_time	期望到货时间	是	[string]	2022-04-15
data>>cg_uid	采购员id	是	[int]	1
data>>cg_opt_username	采购员名称	是	[string]	小米
data>>remark	产品备注	是	[string]	备注
data>>is_combo	是否为组合商品：0 否，1 是	是	[int]	1
data>>is_aux	是否为辅料：0 否，1 是	是	[int]	1
data>>is_related_process_plan	是否关联了加工计划：0 否，1 是	是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "350C7176-A759-3517-9CDA-369C9E1A5743",
    "response_time": "2022-05-20 11:48:25",
    "total": 1,
    "data": [
        {
            "plan_sn": "PP230712003",
            "ppg_sn": "PPG230712003",
            "product_name": "10369333测试",
            "sku": "10369333",
            "fnsku": "",
            "pic_url": "",
            "supplier_id": 590,
            "status_text": "待采购",
            "status": 2,
            "sid": 0,
            "expect_arrive_time": "",
            "remark": "",
            "quantity_plan": 1,
            "product_id": 18791,
            "group_id": 2689,
            "cg_uid": 0,
            "is_combo": 0,
            "is_aux": 0,
            "spu": "",
            "spu_name": "",
            "attribute": [],
            "creator_uid": 10317908,
            "wid": 1,
            "warehouse_name": "默认仓库",
            "purchaser_id": 113,
            "supplier_name": "好的供应商111",
            "cg_opt_username": "",
            "creator_real_name": "验证账户2",
            "file": [
                {
                    "name": "导入计划采购商品模板news.xlsx",
                    "url": "https://imagexxxx/d73aaf35f.xlsx"
                }
            ],
            "plan_remark": "",
            "create_time": "2021-10-11 01:20:20",
            "msku": [],
            "cg_box_pcs": 0,
            "seller_name": "",
            "marketplace": "",
            "purchaser_name": "基本abc",
            "is_related_process_plan": 0
        }
    ]
}
复制
错误
复制成功
 上一章节
查询采购方列表
下一章节 
创建待采购的采购计划