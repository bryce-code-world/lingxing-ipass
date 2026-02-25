查询供应商列表

支持查询【采购】>【供应商】中的供应商信息

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/local_inventory/supplier	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000	否	[int]	1000
请求示例
{
    "offset": 0,
    "length": 1000
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
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2022-05-20 16:57:03
data	响应数据	是	[array]	
data>>customer_supplier_id	客户供应商id【已停用】	是	[string]	
data>>supplier_id	系统供应商id	是	[int]	500
data>>supplier_name	供应商名称	是	[string]	供应商5870
data>>supplier_code	供应商编码【供应商代码只支持数字、英文字母、英文句号、-】	是	[string]	34-AK.pei
data>>employees	规模：员工数	是	[int]	1
data>>url	供应商网址	是	[string]	012010
data>>contact_person	联系人	是	[string]	012010
data>>contact_number	联系电话	是	[string]	012010
data>>qq	qq	是	[string]	012010
data>>email	邮箱	是	[string]	012010
data>>fax	传真	是	[string]	012010
data>>account_name	户名	是	[string]	012010
data>>open_bank	开户行	是	[string]	012010
data>>bank_card_number	银行卡号	是	[string]	012010
data>>remark	备注	是	[string]	0接口新增供应商5870
data>>purchaser	跟进人	是	[string]	1
data>>is_delete	是否已删除：0 否，1 是	是	[int]	0
data>>address_full	详细地址	是	[string]	
data>>payment_method_text	支付方式	是	[string]	网银转账
data>>pc_name	采购合同名称	是	[string]	
data>>settlement_method_text	结算方式	是	[string]	款到发货
data>>settlement_description	结算描述	是	[string]	结算描述1
data>>employees_text	规模大小	是	[string]	少于50人
data>>level_text	级别	是	[string]	★
data>>credit_code	统一社会信用代码	是	[string]	
data>>prepay_percent	预付比例	是	[string]	
data>>payment_account_group	收款账户	是	[array]	
data>>payment_account_group>>name	账户名称	是	[string]	
data>>payment_account_group>>account_name	户名	是	[string]	
data>>payment_account_group>>bank_name	开户行	是	[string]	
data>>payment_account_group>>account_id	账号	是	[string]	
data>>payment_account_group>>remark	备注	是	[string]	
data>>payment_account_group>>is_default	是否默认账户	是	[int]	
data>>payment_account_group>>is_open	是否启用账户	是	[int]	
data>>payment_account_group>>key	功能	是	[string]	
data>>payment_account_group>>version	功能版本	是	[string]	
data>>period_config_key	结算账期	是	[string]	
data>>period_config_text	结算账期文本	是	[string]	
data>>wid	仓库id	是	[int]	
data>>w_name	仓库名称	是	[string]	
data>>template_id	新采购合同模板id	是	[string]	
data>>template_name	新采购合同模板名称	是	[string]	
data>>purchaser_id	采购方id	是	[int]	
data>>purchaser_id_text	采购方	是	[string]	
data>>receipt_wid	采购收货仓库id	是	[int]	
data>>receipt_wid_text	采购收货仓库	是	[string]	
total	总数	是	[int]	1
返回成功示例
{
  "code": 0,
  "message": "success",
  "error_details": [],
  "request_id": "929A1D7D-D656-0EA0-C78A-A686790C7090",
  "response_time": "2022-05-20 16:57:03",
  "data": [
    {
      "customer_supplier_id": null,
      "supplier_id": 500,
      "supplier_name": "供应商5870",
      "supplier_code": "34-AK.pei",
      "employees": 1,
      "url": "012010",
      "contact_person": "012010",
      "contact_number": "012010",
      "qq": "012010",
      "email": "012010",
      "fax": "012010",
      "account_name": "012010",
      "open_bank": "012010",
      "bank_card_number": "012010",
      "remark": "0接口新增供应商5870",
      "purchaser": "1",
      "is_delete": 0,
      "address_full": "",
      "payment_method_text": "网银转账",
      "pc_name": "",
      "settlement_method_text": "款到发货",
      "settlement_description": "结算描述1",
      "employees_text": "少于50人",
      "level_text": "★",
      "credit_code": "",
      "prepay_percent": "",
      "payment_account_group": [
        {
          "name": "默认收款账户",
          "account_name": "张三",
          "bank_name": "中国工商银行",
          "account_id": "6222021000000000000",
          "remark": "主要收款账户",
          "is_default": 1,
          "is_open": 1,
          "key": "default_payment_account",
          "version": "1.0.0"
        }
      ],
      "period_config_key": "MONTHLY_SETTLEMENT",
      "period_config_text": "月结",
      "wid": 101,
      "w_name": "主仓库",
      "template_id": "template_abc_123",
      "template_name": "通用采购合同模板",
      "purchaser_id": 200,
      "purchaser_id_text": "采购部门A",
      "receipt_wid": 102,
      "receipt_wid_text": "采购收货仓库B"
    },
    {
      "customer_supplier_id": null,
      "supplier_id": 501,
      "supplier_name": "供应商6000",
      "supplier_code": "SUP-001.XYZ",
      "employees": 5,
      "url": "www.supplier6000.com",
      "contact_person": "李四",
      "contact_number": "13800138000",
      "qq": "123456789",
      "email": "lisi@supplier6000.com",
      "fax": "021-88889999",
      "account_name": "李四公司",
      "open_bank": "中国建设银行",
      "bank_card_number": "6227000000000000000",
      "remark": "新注册供应商",
      "purchaser": "2",
      "is_delete": 0,
      "address_full": "上海市浦东新区张江高科",
      "payment_method_text": "银行转账",
      "pc_name": "电子元件采购合同",
      "settlement_method_text": "货到付款",
      "settlement_description": "货品验收合格后3个工作日内付款",
      "employees_text": "50-100人",
      "level_text": "★★",
      "credit_code": "913100001234567890",
      "prepay_percent": "20%",
      "payment_account_group": [
        {
          "name": "主要收款账户",
          "account_name": "李四公司",
          "bank_name": "中国建设银行",
          "account_id": "6227000000000000000",
          "remark": "默认收款账户",
          "is_default": 1,
          "is_open": 1,
          "key": "main_payment_account",
          "version": "1.0.0"
        }
      ],
      "period_config_key": "WEEKLY_SETTLEMENT",
      "period_config_text": "周结",
      "wid": 103,
      "w_name": "辅料仓库",
      "template_id": "template_def_456",
      "template_name": "辅料采购合同模板",
      "purchaser_id": 201,
      "purchaser_id_text": "采购部门B",
      "receipt_wid": 104,
      "receipt_wid_text": "辅料收货仓库C"
    }
  ],
  "total": 2
}
复制
错误
复制成功
 上一章节
查询透明计划商品列表
下一章节 
添加/修改供应商