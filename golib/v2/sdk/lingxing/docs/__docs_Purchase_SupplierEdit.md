添加/修改供应商

支持添加、修改供应商信息到系统【采购】>【供应商】中

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/storage/supplier/edit	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
supplier_id	客户供应商id,为空或者对应的值不存在时，取sys_supplier_id【已停用】	否	[string]	25
sys_supplier_id	系统供应商id，取该值且该值为空时，新增供应商	否	[int]	
supplier_name	供应商名称	是	[string]	xxx
supplier_code	供应商编码【供应商代码只支持数字、英文字母、英文句号、-】	否	[string]	114514
employees	员工数：
1=>少于50人
2=>50-150人
3=>150-500人
4=>500-1000人
5 =>1000人以上	否	[int]	1
url	供应商网址	否	[string]	/xxxx/xxx
contact_person	联系人
备注：支持传空值	是	[string]	小强
contact_number	联系电话
备注：支持传空值	是	[string]	15700000000
qq	QQ	否	[string]	1111111111
email	email	否	[string]	1111111111@qq.com
fax	传真	否	[string]	123
account_name	户名	否	[string]	1
open_bank	开户行	否	[string]	lingxxx
bank_card_number	银行卡号	否	[string]	123465
address	详细地址	否	[string]	shenzhen
remark	备注	否	[string]	gogogo
level	级别：
1=>★
2=>★★
3=>★★★
4=>★★★★
5=>★★★★★,	否	[int]	1
settlement_method	结算方式：
7 现结
8 月结	是	[int]	7
settlement_description	结算描述	否	[string]	go go go!!!
payment_method	支付方式：
1=>网银转账
2=>网上支付	否	[int]	1
purchaser	跟进人uid，最多支持10个	否	[array]	[154,623]
credit_code	统一社会信用代码	否	[string]	
prepay_percent	预付比例	否	[string]	
payment_account_group	收款账户列表	否	[array]	
payment_account_group>>name	账户名称	是	[string]	
payment_account_group>>account_name	户名	是	[string]	
payment_account_group>>bank_name	开户行	是	[string]	
payment_account_group>>account_id	账号	是	[string]	
payment_account_group>>is_default	是否默认
0=>否
1=>是	否	[int]	
payment_account_group>>is_open	是否开启
0=>否
1=>是	否	[int]	
payment_account_group>>remark	备注	否	[string]	
payment_account_group>>key	编辑时必须	否	[string]	
payment_account_group>>version	编辑时必须	否	[string]	
请求示例
{
    "supplier_id": 25,
    "supplier_name": "xxx",
    "supplier_code": "114514",
    "employees": 1,
    "url": "/xxxx/xxx",
    "contact_person": "小强",
    "contact_number": "15700000000",
    "qq": "1111111111",
    "email": "1111111111@qq.com",
    "fax": "123",
    "account_name": "1",
    "open_bank": "lingxxx",
    "bank_card_number": "123465",
    "address": "shenzhen",
    "remark": "gogogo",
    "level": "1",
    "settlement_method": 7,
    "settlement_description": "go go go!!!",
    "payment_method": 1,
    "purchaser": [154,623],
    "payment_account_group": [{
        "name": "",
        "account_name": "",
        "bank_name": "",
        "account_id": "",
        "is_default": 0,
        "is_open": "",
        "remark": "",
        "key": "",
        "version": ""
    }],
    "credit_code": "",
    "prepay_percent": ""
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
error_details	错误信息	是	[array]	[]
request_id	请求链路id	是	[string]	929A1D7D-D656-0EA0-C78A-A686790C7090
response_time	响应时间	是	[string]	2022-05-20 16:57:03
data	响应数据	是	[array]	
data>>customer_supplier_id	客户供应商ID【已停用】	是	[int]	235
data>>erp_supplier_id	系统供应商ID	是	[int]	238
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "BCB699A2-CEA9-FEDE-7DF3-28BCD3C71727",
    "response_time": "2021-11-11 19:25:38",
    "data": {
        "customer_supplier_id": "235",
        "erp_supplier_id": "238"
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "error_details": "规模：员工数必须在 1,2,3,4,5 范围内 [请求码:1D58EB]",
    "request_id": "143E9E96-744E-CCDE-B31D-A00C697F3C8F",
    "response_time": "2021-11-11 19:21:56",
    "data": [],
    "total": 0
}
复制
错误
复制成功
附加说明
客户供应商ID，指客户通过领星开放接口创建的供应商，在传值时向领星传入的供应商ID，用于记录客户系统供应商ID，仅支持从开放接口创建时录入，可为空；（已停用，新创建供应商不支持录入客户供应商ID，已有数据不受影响）
系统供应商ID，指领星ERP自动生成的供应商ID，唯一标识，每个供应商均有自己的系统ID；
 上一章节
查询供应商列表
下一章节 
查询采购方列表