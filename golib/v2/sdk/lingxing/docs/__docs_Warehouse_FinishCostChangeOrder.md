创建已完成的成本补录单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/CostChangeOrder/finishCostChangeOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
type	补录类型—只支持入库成本(1)	是	[int]	1
wid	仓库ID	是	[int]	1
remark	备注	是	[string]	成本补录单备注
list	成本补录子项	是	[array]	
list>>product_id	产品ID	是	[int]	1
list>>seller_id，对应查询亚马逊店铺列表接口对应字段【sid】	否	[int]	0
list>>fnsku	FNSKU，默认为空	否	[string]	
list>>relation_order_out	入库单号	是	[string]	IB250514006
list>>unit_cost_price	变更后采购单价	是	[float]	10.0000
list>>unit_fee_price	变更后单位费用	是	[float]	6.0000
请求示例
{
  "type": 1,
  "wid": 1,
  "remark": "成本补录单备注",
  "list": [
    {
      "product_id": 1,
      "seller_id": 0,
      "fnsku": "",
      "relation_order_out": "IB250514006",
      "unit_cost_price": 10.0000,
      "unit_fee_price": 6.0000
    },
    {
      "product_id": 2,
      "seller_id": 0,
      "fnsku": "X000001",
      "relation_order_out": "IB250514007",
      "unit_cost_price": 15.5000,
      "unit_fee_price": 5.2000
    }
  ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	数据字典	限制	示例
code	状态码，0 成功 参数示例： 0	是	[string]			0
message	提示消息	是	[string]			success
error_details	错误信息	是	[string]			error_msg
request_id	请求链路id	是	[string]			81CA6A9A-A5C4-8B5E-5FED-E1C2EC04753C
response_time	响应时间	是	[string]			2025-05-20 00:00:00
data	响应数据	是	[array]			
data>>order_sn	成本补录单号	是	[string]			CA2505220016
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "6DB9CEC4-C9B6-9175-36C7-CE7A2CFC3968",
    "response_time": "2021-09-03 10:38:21",
    "data": {"order_sn":"CA2505220016"}
}
复制
错误
复制成功
 上一章节
删除出库单
下一章节 
查询海外仓备货单列表