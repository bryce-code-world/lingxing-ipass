打印AWD入库货件箱子标签
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/awd/inbound-shipment/uploadPacking	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
shipmentIdInfo		是	[array]	
shipmentIdInfo>>pageType	纸张类型：THERMAL_NONPCP：热敏纸1个标签；THERMAL_NONPCP_01：热敏纸（152 x 108 mm）1个标签；THERMAL_NONPCP_01_WATERMARK：热敏纸（152 x 108 mm）1个标签（带水印）；PLAIN_PAPER：每张美国信纸1个标签；LETTER_6：每张美国信纸6个标签	是	[string]	THERMAL_NONPCP_01
shipmentIdInfo>>shipmentId	货件单号	是	[string]	AWD20230715001
shipmentIdInfo>>sid	领星店铺ID 对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	123456789
请求示例

json

{
  "shipmentIdInfo": [
    {
      "pageType": "THERMAL_NONPCP_01",
      "shipmentId": "AWD20230715001",
      "sid": 123456789
    },
    {
      "pageType": "LETTER_6", 
      "shipmentId": "AWD20230715002",
      "sid": 987654321
    }
  ]
}
复制
错误
复制成功
 上一章节
更新AWD货件跟踪编号
FBA
下一章节 
查询补货列表