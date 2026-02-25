待收货退货单快捷入库
接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/return/order/fastStorageIn	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
reqs		是	[array]	
reqs>>rmaOrderNo	退货单号	是	[string]	RT901582009478087168
reqs>>storeId	店铺ID	是	[string]	110556144355708928
reqs>>wid	退货仓库	是	[string]	1
reqs>>itemReqs		是	[array]	
reqs>>itemReqs>>id	退货商品行id	是	[string]	1933404623618236417
reqs>>itemReqs>>picIds	售后图片	否	[array]	[{"accessUrl":"","fileName":"","mappingKey":""}]
reqs>>itemReqs>>picIds>>accessUrl	访问地址	否	[string]	
reqs>>itemReqs>>picIds>>fileName	图片名	否	[string]	
reqs>>itemReqs>>picIds>>mappingKey	映射键	否	[string]	
reqs>>itemReqs>>availableQuantity	可用量	是	[number]	2
reqs>>itemReqs>>availableWhbCode	可用仓位编码	是	[string]	ts_valid
reqs>>itemReqs>>defectiveQuantity	次品量	是	[number]	0
reqs>>itemReqs>>defectiveWhbCode	次品仓位编码	是	[string]	ts_bad
reqs>>itemReqs>>destroyedQuantity	销毁量	是	[number]	0
请求示例
{
    "reqs": [
        {
            "rmaOrderNo": "RT901582009478087168",
            "storeId": "110556144355708928",
            "wid": "1",
            "itemReqs": [
                {
                    "id": "1933404623618236417",
                    "picIds": [
                        {
                            "accessUrl": "",
                            "fileName": "",
                            "mappingKey": ""
                        }
                    ],
                    "availableQuantity": 2,
                    "availableWhbCode": "ts_valid",
                    "defectiveQuantity": 0,
                    "defectiveWhbCode": "ts_bad",
                    "destroyedQuantity": 0

                }
            ]
        }
    ]
}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0：成功	是	[number]	0
message	消息提示	是	[string]	success
error_details	错误详情	是	[array]	[]
request_id	请求id	是	[string]	3ccb4bf801654cb5bf952c581e0f9ef3.1754532975972
response_time	响应时间	是	[string]	2025-08-07 10:16:18
data	数据	是	[object]	
data>>suc	成功列表	是	[array]	[{"rmaOrderNo":"","storageInOrderNo":"","thirdReturnNo":""}]
data>>suc>>rmaOrderNo	退货单号	是	[string]	
data>>suc>>storageInOrderNo	入库单号	是	[string]	
data>>suc>>thirdReturnNo	三方仓推荐单号	是	[string]	
data>>fail	失败列表	是	[array]	[{"rmaOrderNo":"RT901582009478087168","failReason":null}]
data>>fail>>rmaOrderNo	退货单号	是	[string]	RT901582009478087168
data>>fail>>cause	失败原因	是	[string]	初始状态状态不能快捷入库
total	总数	是	[number]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "3ccb4bf801654cb5bf952c581e0f9ef3.1754532975972",
    "response_time": "2025-08-07 10:16:18",
    "data": {
        "suc": [
            {
                "rmaOrderNo": "",
                "storageInOrderNo": "",
                "thirdReturnNo": ""
            }
        ],
        "fail": [
            {
                "rmaOrderNo": "RT901582009478087168",
                "cause": "初始状态状态不能快捷入库"
            }
        ]
    },
    "total": 1
}
复制
错误
复制成功
 上一章节
查询质检单详情
下一章节 
订单退款