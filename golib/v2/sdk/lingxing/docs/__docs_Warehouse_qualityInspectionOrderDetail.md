查询质检单详情
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/qualityInspectionOrder/detail	HTTPS	POST	1
请求参数：
参数名	说明	必填	类型	示例
qc_sn	质检单号	是	[string]	QC220719001
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	ecfbe567-9abe-4dbd-9c4b-9a5929bcdee9.1677838481851
response_time	响应时间	是	[string]	2020-04-30 17:33:32
total	总数	是	[int]	6
data	响应数据	是	[object]	
data>>qc_id	质检单id	是	[string]	1_QC220818016
data>>qc_sn	质检单号	是	[string]	QC220818016
data>>qc_type	质检类型：
1 仓库质检
2 预检
3 免检	是	[string]	1
data>>qc_type_text	质检类型名称	是	[string]	仓库质检
data>>qc_method	质检方式：
1 抽检
2 全检	是	[string]	1
data>>qc_method_text	质检方式名称	是	[string]	抽检
data>>qc_image	质检标准图片id	是	[string]	19,20,21,22
data>>receive_time	收货时间	是	[string]	2022-08-18 16:51:33
data>>receive_uid	收货人id	是	[string]	560
data>>qc_uid	质检人id	是	[string]	0
data>>sid	店铺id	是	[string]	0
data>>product_receive_num	质检量	是	[int]	10
data>>product_good_num	良品量	是	[int]	3
data>>product_bad_num	次品量	是	[int]	0
data>>qc_time	质检时间	是	[string]	2023-03-18 10:19:56
data>>status	质检状态：
0 待质检(质检中)
1 已质检
2 已免检
10 已质检(已撤销)
20 已免检(已撤销)	是	[string]	0
data>>status_text	质检状态说明	是	[string]	质检中
data>>price	单价	是	[string]	1.000000
data>>product_id	本地产品id	是	[string]	2121907
data>>product_name	品名	是	[string]	AAAA001
data>>sku	SKU	是	[string]	AAAA001
data>>wid	仓库id	是	[int]	1
data>>order_id	采购单id/委外单id	是	[string]	1_PO220818002
data>>order_sn	采购单号/委外单号	是	[string]	PO220818002
data>>order_type	订单类型	是	[string]	1
data>>cg_uid	采购员id	是	[string]	560
data>>fnsku	FNSKU	是	[string]	X001JVQKZV
data>>file_id	附件id	是	[string]	
data>>qc_num	抽检数量	是	[int]	0
data>>qc_bad_num	质检次品量	是	[int]	0
data>>qc_rate	抽检比例	是	[string]	1.0000
data>>qc_rate_pass	抽检合格率	是	[string]	0.0000
data>>qc_remark	备注	是	[string]	this is a demo
data>>qc_pic_url	图片地址	是	[array]	
data>>whb_code_good	可用仓位	是	[string]	
data>>whb_code_bad	次品仓位	是	[string]	
data>>product_qc_num	质检量	是	[int]	10
data>>qc_realname	质检员	是	[string]	管理员
data>>receive_realname	收货员	是	[string]	管理员
data>>opt_realname	操作员	是	[string]	曹管理员
data>>pic_url	产品图片地址	是	[string]	
data>>is_combo	是否组合产品：
0 否
1 是	是	[int]	0
data>>is_aux	是否辅料：
0 否
1 是	是	[int]	0
data>>supplier_id	供应商id	是	[int]	1
data>>supplier_name	供应商名称	是	[string]	供应商1
data>>source	采购数据来源：
0 web
1 金蝶
2 openApi	是	[int]	2
data>>file	附件	是	[array]	
data>>image	质检标准图片	是	[array]	
data>>image>>file_id	图片ID	是	[string]	19
data>>image>>name	图片名称	是	[string]	1634625436(1).jpg
data>>image>>url	链接	是	[string]	/api/file/previewimg/id/19/type/0.html
data>>qc_standard	质检标准	是	[array]	
data>>qc_standard>>pqs_id	质检标准id	是	[string]	210212154628104220
data>>qc_standard>>qc_id	质检单id	是	[string]	1_QC220818016
data>>qc_standard>>qc_sn	质检单号	是	[string]	QC220818016
data>>qc_standard>>type	质检标准类型：
1 系统
2 自定义	是	[int]	1
data>>qc_standard>>qc_item	质检项	是	[string]	颜色
data>>qc_standard>>qc_content	质检内容	是	[string]	白色，无色差
data>>custom_receive_time	收货时间	是	[string]	2023-03-13 12:03:31
data>>custom_qc_time	质检时间	是	[string]	2023-03-17 23:59:59
data>>delivery_order_sn	收货单号	是	[string]	CR230313001
data>>source_custom_order_sn	自定义单号	是	[string]	PO230313002
data>>whb_code_good_list	良品仓位	是	[array]	
data>>whb_code_good_list>>whb_num	仓位数量	是	[int]	1
data>>whb_code_good_list>>whb_code	仓位编码	是	[string]	良品暂存
data>>whb_code_bad_list	次品仓位	是	[array]	
data>>whb_code_bad_list>>whb_num	仓位数量	是	[int]	1
data>>whb_code_bad_list>>whb_code	仓位编码	是	[string]	次品暂存
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a3995d6442db4fe48897228bc993632f.1693474878330",
    "response_time": "2023-08-31 17:41:20",
    "total": 0,
    "data": {
        "qc_id": "1_QC220719001",
        "qc_sn": "QC220719001",
        "qc_type": "1",
        "qc_type_text": "仓库质检",
        "qc_method": "1",
        "qc_method_text": "抽检",
        "qc_image": "",
        "receive_time": "2022-07-19 16:35:05",
        "receive_uid": "1",
        "qc_uid": "1",
        "sid": "35",
        "seller_id": "35",
        "product_receive_num": 3,
        "product_good_num": 3,
        "product_bad_num": 0,
        "qc_time": "2022-07-19 16:35:06",
        "status": "1",
        "status_text": "已质检",
        "price": "0.000000",
        "product_id": "128671",
        "product_name": "组合店铺（英德）",
        "sku": "0714-3",
        "wid": 3173,
        "order_id": "1_PO220719006",
        "order_sn": "PO220719006",
        "order_type": "1",
        "cg_uid": "1",
        "fnsku": "X001JVQKZV",
        "file_id": "",
        "qc_num": 3,
        "qc_bad_num": 0,
        "qc_rate": "1.0000",
        "qc_rate_pass": "1.0000",
        "qc_remark": "",
        "qc_pic_url": [],
        "whb_code_good": null,
        "whb_code_bad": null,
        "product_qc_num": 3,
        "qc_realname": "0超级管理员01",
        "receive_realname": "0超级管理员01",
        "opt_realname": "0超级管理员01",
        "pic_url": "",
        "is_combo": 1,
        "is_aux": 0,
        "supplier_id": 3452,
        "supplier_name": "34wefs修改gys",
        "source": 0,
        "file": [],
        "image": [],
        "qc_standard": [],
        "custom_receive_time": "2022-07-19 16:35:05",
        "custom_qc_time": "2022-07-19 16:35:06",
        "delivery_order_sn": "CR220719002",
        "source_custom_order_sn": "PO220719006",
        "whb_code_good_list": [
            {
                "whb_num": "3",
                "whb_code": "可用暂存"
            }
        ],
        "whb_code_bad_list": []
    }
}
复制
错误
复制成功
 上一章节
查询质检单列表
下一章节 
待收货退货单快捷入库