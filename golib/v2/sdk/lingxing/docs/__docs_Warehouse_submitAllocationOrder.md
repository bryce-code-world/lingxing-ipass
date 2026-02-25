创建待调拨的调拨单
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/inventoryReceipt/StorageAllocation/submitAllocationOrder	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
sys_wid	系统出库仓库ID	是	[int]	1
sys_to_wid	系统入库仓库ID	是	[int]	11
freight_fee	运费	否	[string]	50.00
other_fee	其他费用	否	[string]	
fee_part_type	费用分摊方式：0 不分摊【默认值】，2 按sku数量分摊，3 按重量，4 按体积，5 按自定义	否	[int]	0
remark	备注	否	[string]	
predict_time	预计到货时间	否	[string]	2025-03-10
type	默认为2-标准调拨	否	[string]	2
out_bin_type	默认0 出库仓位不为空时必传1	是	[string]	1
product_list	产品明细	是	[array]	
product_list>>product_id	产品id	是	[string]	1
product_list>>seller_id	店铺id，不传默认0【对应查询亚马逊店铺信息接口字段sid】	否	[string]	1
product_list>>fnsku	fnsku，不传默认为空	否	[string]	
product_list>>good_num	可用调拨量，和次品调拨量其中一项必填	否	[int]	48
product_list>>bad_num	次品调拨量，和可用调拨量其中一项必填	否	[int]	0
product_list>>cg_package_length	包装规格-长（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_package_width	包装规格-宽（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_package_height	包装规格-高（CM），不传或者传空则取产品管理的值	否	[string]	
product_list>>cg_product_gross_weight	单品净重，不传或者传空则取产品管理的值	否	[string]	
product_list>>freight_fee_unit	单位运费，若费用分摊方式为5-自定义则必填，其他则该字段无效	否	[string]	
product_list>>other_fee_unit	单位其他费用，若费用分摊方式为5-自定义则必填，其他则该字段无效	否	[string]	
product_list>>product_remark	明细备注，最大长度为255个字符	否	[string]	
product_list>>out_available_bin	出库可用仓位列表，不指定则传空数组 []	否	[array]	
product_list>>out_available_bin>>whb_code	出库可用仓位编码	否	[string]	
product_list>>out_available_bin>>num	出库可用仓位调拨量	否	[string]	
product_list>>out_inferior_bin	出库次品仓位列表，不指定则传空数组 []	否	[array]	
product_list>>out_inferior_bin>>whb_code	出库次品仓位编码	否	[string]	
product_list>>out_inferior_bin>>num	出库次品仓位调拨量	否	[string]	
product_list>>to_available_bin	入库可用仓位列表，不指定则传空数组 []	否	[array]	
product_list>>to_available_bin>>whb_code	入库可用仓位编码	否	[string]	
product_list>>to_available_bin>>num	入库可用仓位调拨量	否	[string]	
product_list>>to_inferior_bin	入库次品仓位列表，不指定则传空数组 []	否	[array]	
product_list>>to_inferior_bin>>whb_code	入库次品仓位编码	否	[string]	
product_list>>to_inferior_bin>>num	入库次品仓位调拨量	否	[string]	
请求curl示例
curl --location --globoff 'https://openapi.lingxing.com/erp/sc/routing/inventoryReceipt/StorageAllocation/submitAllocationOrder?access_token=value&app_key=value&sign=value&timestamp=value' \
--header 'Content-Type: application/json' \
--data '{
  "sys_wid": 1,
  "sys_to_wid": 11,
  "freight_fee": "50.00",
  "other_fee": "10.00",
  "fee_part_type": 0,
  "remark": "紧急调拨，请优先处理",
  "predict_time": "2025-03-10",
  "type": "2",
  "out_bin_type": "1",
  "product_list": [
    {
      "product_id": "1",
      "seller_id": "1",
      "fnsku": "X000001",
      "good_num": 48,
      "bad_num": 0,
      "cg_package_length": "30",
      "cg_package_width": "20",
      "cg_package_height": "10",
      "cg_product_gross_weight": "2.5",
      "freight_fee_unit": "",
      "other_fee_unit": "",
      "product_remark": "第一批次调拨",
      "out_available_bin": [
        {
          "whb_code": "A001",
          "num": "24"
        },
        {
          "whb_code": "A002",
          "num": "24"
        }
      ],
      "out_inferior_bin": [],
      "to_available_bin": [
        {
          "whb_code": "B001",
          "num": "48"
        }
      ],
      "to_inferior_bin": []
    },
    {
      "product_id": "2",
      "seller_id": "2",
      "fnsku": "X000002",
      "good_num": 30,
      "bad_num": 5,
      "cg_package_length": "25",
      "cg_package_width": "15",
      "cg_package_height": "5",
      "cg_product_gross_weight": "1.5",
      "freight_fee_unit": "",
      "other_fee_unit": "",
      "product_remark": "第二批次调拨",
      "out_available_bin": [
        {
          "whb_code": "A003",
          "num": "30"
        }
      ],
      "out_inferior_bin": [
        {
          "whb_code": "A004",
          "num": "5"
        }
      ],
      "to_available_bin": [
        {
          "whb_code": "B002",
          "num": "30"
        }
      ],
      "to_inferior_bin": [
        {
          "whb_code": "B003",
          "num": "5"
        }
      ]
    }
  ]
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	值可能性	限制	示例
code	状态码，0 成功	是	[int]			0
message	提示消息	是	[string]			success
error_details	错误信息	是	[array]			
request_id	请求链路id	是	[string]			81CA6A9A-A5C4-8B5E-5FED-E1C2EC04753C
response_time	响应时间	是	[string]			2020-12-01 12:11:12
data	响应数据	是	[object]			
data>>order_sn	调拨单号	是	[string]	;		TF220729002
成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "C70FD2A7-4E7D-6B32-7D72-FC42587CE97E",
    "response_time": "2022-07-29 14:33:42",
    "data": {
        "order_sn": "TF220729002"
    },
    "total": 0
}
复制
错误
复制成功
 上一章节
获取自定义出库类型
下一章节 
入库单确认入库