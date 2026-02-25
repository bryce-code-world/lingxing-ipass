创建费用单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/bd/fee/management/open/feeManagement/otherFee/create	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
submit_type	提交类型：1 暂存，2 提交	是	[int]	2
dimension	分摊维度：
1 msku
2 asin
3 店铺
4 父asin
5 sku
6 企业	是	[int]	2
apportion_rule	分摊规则：
0 无
1 按销售额
2 按销量
3 店铺均摊后按销售额占比分摊
4 店铺均摊后按销量占比分摊	是	[int]	1
is_request_pool	是否请款：0 否，1 是	是	[int]	0
remark	费用单备注	是	[string]	备注
fee_items	费用明细项	是	[array]	
fee_items>>sids	店铺id：
单选店铺传店铺id
全部店铺传[99999999]
企业费用传[88888888]
多选店铺传[77777777] ，对应查询亚马逊店铺列表接口对应字段【sid】	是	[array]	[1,2]
fee_items>>dimension_value	纬度值，例如ASIN值	是	[string]	B0ABCD001
fee_items>>date	分摊日期，格式：Y-m-d 或 Y-m	是	[string]	2023-02
fee_items>>other_fee_type_id	费用类型id，查询费用类型列表 接口对应字段【id】	是	[number]	1000014
fee_items>>fee	金额	是	[number]	-100
fee_items>>fee	原币金额（注意正负数）	是	[number]	-100
fee_items>>currency_code	币种代码	是	[string]	CNY
fee_items>>remark	费用子项备注	是	[string]	费用子项备注
请求示例
{
    "submit_type": 2,
    "dimension": 1,
    "apportion_rule": 1,
    "is_request_pool": 0,
    "remark": "备注",
    "fee_items": [
        {
            "sids": [
                106
            ],
            "dimension_value": "FO-F20Y-K0KC",
            "date": "2023-02",
            "other_fee_type_id": 1167,
            "fee": -100,
            "currency_code": "CNY",
            "remark": "费用子项备注"
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
msg	消息提示	是	[string]	操作成功
data	响应数据	是	[string]	true
返回成功示例
{
    "code": 0,
    "msg": "操作成功",
    "data": true
}
复制
错误
复制成功
返回失败示例
{
    "code": 0,
    "msg": "当前月份：2023-02利润报表已结账/结账中，不可提交该月费用。",
    "data": null,
    "traceId": "82e464239000465fad92ae52ac9ee173.1696846050343"
}
复制
错误
复制成功
 上一章节
查询费用明细列表
下一章节 
编辑费用单