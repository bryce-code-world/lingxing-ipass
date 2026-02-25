查询盘点单详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/InventoryCheck/getOrderDetail	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
order_sn	盘点单号	是	[string]	IC240726001
search_field	搜索字段：
sku SKU
fnsku FNSKU
product_name 品名
whb_code_text 仓位
whb_type_text 仓位类型	否	[string]	sku
search_value	搜索值	否	[string]	diannaotaozhuang
sort_field	排序字段：
book_inventory 账面库存
actual_inventory 实盘库存
different_count 库存差异	否	[string]	book_inventory
sort_type	排序规则：desc 降序【默认】，asc 升序	否	[string]	desc
page	分页页码，默认1【控制 product_list 返回数目】	否	[int]	1
page_size	分页长度，默认20【控制 product_list 返回数目】	否	[int]	20
请求示例
{
    "order_sn": "IC240726001",
    "search_field": "sku",
    "search_value": "diannaotaozhuang",
    "sort_field": "book_inventory",
    "sort_type": "desc",
    "page": 1,
    "page_size": 20
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	信息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	D17B0688-C52D-1BA9-1BCC-50AA00058A3F
response_time	响应时间	是	[string]	2022-08-23 16:56:31
data	响应数据	是	[object]	
data>>order_sn	盘点单号	是	[string]	IC220810001
data>>status	盘点状态：
10 待盘点
20 预锁
30 盘点中
40 已盘点
121 待审核
122 已驳回
123 通过
124 作废	是	[int]	40
data>>status_text	状态文本	是	[string]	已盘点
data>>wid	盘点仓库id	是	[int]	1
data>>ware_house_name	盘点仓库名称	是	[string]	仓库1
data>>check_type	盘点类型：
1 整仓盘点
2 SKU盘点
3 仓位盘点
4 SKU+仓位盘点	是	[int]	2
data>>check_type_text	盘点类型说明	是	[string]	SKU盘点
data>>is_display_check	是否明盘：0 否，1 是	是	[int]	1
data>>display_check_name	是否明盘文本	是	[string]	
data>>is_zero	是否零库存参与盘点：0 否，1 是	是	[int]	1
data>>product_type	产品种类	是	[int]	1
data>>create_uid	创建人id	是	[int]	10002789
data>>create_user	创建人姓名	是	[string]	张三
data>>create_time	创建时间	是	[string]	2022-08-10 18:02:37
data>>check_uid	盘点人id	是	[int]	10002789
data>>check_user	盘点人姓名	是	[string]	张三
data>>real_check_uid	实际盘点人id	是	[int]	10002789
data>>real_check_user	实际盘点人姓名	是	[string]	张三
data>>check_time	盘点时间	是	[string]	2022-08-10 18:02:37
data>>commit_uid	提交人id	是	[int]	10002789
data>>commit_user	提交人姓名	是	[string]	张三
data>>commit_time	提交时间	是	[string]	2022-08-10 18:03:09
data>>cancel_uid	作废人id	是	[int]	
data>>cancel_user	作废人姓名	是	[string]	
data>>cancel_time	作废时间	是	[string]	
data>>cancel_reason	作废原因	是	[string]	
data>>remark	备注	是	[string]	
data>>request_status	单据状态：0 正常，1 处理中	是	[int]	
data>>file	上传附件信息	是	[array]	
data>>file>>file_id	附件id	否	[int]	
data>>file>>file_name	附件名称	否	[string]	
data>>file>>file_url	附件URL	否	[string]	
data>>product_list	盘点明细列表	是	[array]	
data>>product_list>>product_id	本地产品id	是	[int]	9
data>>product_list>>product_name	品名	是	[string]	[演示数据]Sony 儿童耳机
data>>product_list>>sku	SKU	是	[string]	1164812
data>>product_list>>fnsku	FNSKU	是	[string]	
data>>product_list>>seller_id	店铺id	是	[string]	
data>>product_list>>seller_name	店铺名称	是	[string]	
data>>product_list>>country_name	店铺所属国家名称	是	[string]	
data>>product_list>>whb_id	仓位id	是	[int]	2
data>>product_list>>whb_code	仓位	是	[string]	ts_valid
data>>product_list>>whb_code_text	仓位名称	是	[string]	可用暂存
data>>product_list>>whb_type	仓位类型：
1 待检暂存
2 可用暂存
3 次品暂存
4 拣货暂存
5 可用
6 次品
7 可用在途
8 次品在途	是	[int]	2
data>>product_list>>whb_type_text	仓位类型文本	是	[string]	可用暂存
data>>product_list>>book_inventory	账面库存	是	[int]	80
data>>product_list>>actual_inventory	实盘库存	是	[int]	70
data>>product_list>>different_count	盘点差异	是	[int]	-10
data>>product_list>>remark	明细备注	是	[string]	
data>>product_list>>pic_url	产品图片链接	是	[string]	https://xxx/xxx/.jpg
data>>product_list>>is_combo	是否组合产品：0 否，1 是	是	[int]	
data>>product_list>>is_aux	是否辅料：0 否，1 是	是	[int]	
data>>total	盘点明细总数	是	[int]	1
data>>item_total	盘点明细信息	是	[object]	
data>>item_total>>book_inventory	账面库存总数	是	[int]	80
data>>item_total>>actual_inventory	实盘库存总数	是	[int]	70
data>>item_total>>different_count	盘点差异总数	是	[int]	-10
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "7C458B32-D931-12D9-70F8-E6DB82598908",
    "response_time": "2024-08-01 15:49:15",
    "data": {
        "order_sn": "IC240726001",
        "status": 20,
        "wid": 42,
        "check_type": 1,
        "check_uid": 10609364,
        "is_display_check": 1,
        "is_zero": 1,
        "product_type": 15,
        "create_time": "2024-07-26 16:32:20",
        "create_uid": 10609364,
        "check_time": "",
        "real_check_uid": 0,
        "commit_time": "",
        "commit_uid": 0,
        "cancel_time": "",
        "cancel_uid": 0,
        "cancel_reason": "",
        "remark": "",
        "request_status": 0,
        "increment_time": "2024-07-26 16:32:51",
        "check_type_text": "整仓盘点",
        "status_text": "待盘点_预锁",
        "check_user": "邓金琳",
        "real_check_user": "",
        "commit_user": "",
        "create_user": "邓金琳",
        "cancel_user": "",
        "ware_house_name": "BTC",
        "display_check_name": "明盘",
        "product_list": [
            {
                "id": 511209,
                "product_id": 18418,
                "product_name": "电脑套装（电脑＋鼠标＋键盘）",
                "sku": "diannaotaozhuang",
                "fnsku": "",
                "seller_id": "0",
                "whb_id": 582,
                "whb_code": "ts_valid",
                "whb_type": 2,
                "remark": "",
                "actual_inventory": "",
                "book_inventory": 0,
                "different_count": "",
                "whb_type_text": "可用暂存",
                "whb_code_text": "可用暂存",
                "pic_url": "https://image.umaicloud.com/erp-vue/90128554873982976/20220309/a51149090e274723a408d18bc0aad887.jpeg",
                "seller_name": "",
                "country_name": "",
                "is_combo": 1,
                "is_aux": 0
            }
        ],
        "total": 1,
        "item_total": {
            "book_inventory": "0",
            "actual_inventory": 0,
            "different_count": "0"
        },
        "file": []
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
查询盘点单列表
下一章节 
创建已完成的盘点单