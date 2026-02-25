创建STA任务
接口信息
API Path	请求协议	请求方式	令牌桶容量
/amzStaServer/openapi/inbound-plan/createInboundPlan	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
addressLine1	详细街道地址1	是	[string]	
addressLine2	详细街道地址2	否	[string]	
city	城市	是	[string]	
companyName	公司名称	否	[string]	
countryCode	国家(地区）	是	[string]	
email	邮箱	否	[string]	
inboundPlanItems	计划明细列表	是	[array]	
inboundPlanItems>>expiration	有效期	否	[string]	
inboundPlanItems>>labelOwner	标签类型
AMAZON
SELLER
NONE	是	[string]	AMAZON
inboundPlanItems>>msku	msku	是	[string]	
inboundPlanItems>>prepOwner	预处理提供方
AMAZON
SELLER
NONE	是	[string]	AMAZON
inboundPlanItems>>quantity	申报量	是	[int]	
inboundPlanItems>>prepCategory	预处理分类：（当商品未在卖家中心设置过预处理指导时可填写）

ADULT-成人：对应的预处理类型通常为无需预处理[ITEM_NO_PREP]，也可传其他预处理类型

HANGER-悬挂在衣架上的服装：对应的预处理类型通常为聚乙烯塑料袋包装、取下衣架[ITEM_RMOVHANG, ITEM_POLYBAGGING]，也可传其他预处理类型

TEXTILE-服装、面料、毛绒玩具和纺织品：对应的预处理类型通常为聚乙烯塑料袋包装[ITEM_POLYBAGGING]，也可传其他预处理类型

BABY-母婴用品：对应的预处理类型通常为无需预处理[ITEM_NO_PREP]，也可传其他预处理类型

FRAGILE-易碎品：对应的预处理类型通常为气泡膜包装、贴标签[ITEM_BUBBLEWRAP, ITEM_LABELING]，也可传其他预处理类型

LIQUID-液体（未存放在玻璃容器中）：对应的预处理类型通常为无需预处理[ITEM_NO_PREP]，也可传其他预处理类型

PERFORATED-打孔包装：对应的预处理类型通常为聚乙烯塑料袋包装、窒息警告[ITEM_POLYBAGGING, ITEM_SUFFOSTK]，也可传其他预处理类型

GRANULAR-粉末、球状或颗粒状物品：对应的预处理类型通常为聚乙烯塑料袋包装、窒息警告[ITEM_POLYBAGGING, ITEM_SUFFOSTK]，也可传其他预处理类型

SHARP-尖利物品：对应的预处理类型通常为气泡膜包装、贴标签[ITEM_BUBBLEWRAP, ITEM_LABELING]，也可传其他预处理类型

SMALL-小号：对应的预处理类型通常为聚乙烯塑料袋包装、贴标签[ITEM_POLYBAGGING, ITEM_LABELING]，也可传其他预处理类型

SET-套装销售：对应的预处理类型通常为气泡膜包装、贴标签、套装创建[ITEM_BUBBLEWRAP,ITEM_SETCREAT, ITEM_LABELING]，也可传其他预处理类型

NONE-无需进行预处理：对应的预处理类型通常为无需预处理[ITEM_NO_PREP]，也可传其他预处理类型	否，但如果要填写prepCategory或者prepTypes，两者需同时填写	[String]	
inboundPlanItems>>prepTypes	预处理类型：（当商品未在卖家中心设置过预处理指导时可填写）

ITEM_POLYBAGGING-聚乙烯塑料袋包装：将商品放在透明袋中，以防破损、落尘或泄漏。包装袋必须印有窒息警告，必须密封（如果不是自粘袋）。商品也可以使用收缩包装。

ITEM_BUBBLEWRAP-气泡膜包装：如果商品本身不具有保护性包装，请将每件商品都用保护性气泡膜包装或者缠裹起来，也可以将商品放在保护性套箱或气泡袋中（气泡袋必须有窒息警告）。确保包装严密且完全覆盖裸露的边缘。

ITEM_LABELING-贴标签：必须在商品上应用FNSKU标签。

ITEM_SUFFOSTK-窒息警告：开口12厘米或更大的塑料袋（平放时测量）必须有窒息警告。此警告必须印在袋子上或作为标签附加。

ITEM_RMOVHANG-取下衣架：商品不能使用衣架发货。

ITEM_SETCREAT-套装创建：成套销售的商品必须在包装上贴上套装标签。例如，“套装销售”、“准备发货”或“套装勿拆”。单件商品上的条形码不得朝外，否则应予以遮盖。

ITEM_BLACK_SHRINKWRAP-不透明袋包装：将商品放入黑色或不透明的包装袋中，包装袋必须印有窒息警告，必须封紧（如果不是自粘袋）。

ITEM_SIOC-原包装发货：商品以原始产品包装发货。

ITEM_HANG_GARMENT-悬挂衣服：取下衣架，然后将商品放入透明袋中，以防损坏或落尘。商品也可以使用收缩包装。

ITEM_SETSTK-套装销售：套装产品必须在包装上标明为套装。在子套装产品上添加一个标签，清楚地说明产品必须作为整体单位接受和销售。例如，如果一套六个独特的玩具车作为整体单位销售，每辆车的包装必须表明它是套装的一部分

ITEM_TAPING-胶带固定：表示需要用胶带固定。

ITEM_BOXING-装箱：对于锋利物品、易碎物品、危险液体和黑胶唱片，可能需要装箱。对于超过4.5公斤的物品，请使用双壁瓦楞箱。

ITEM_BLANKSTK-空白贴纸：商品需要一个空白贴纸来遮盖无法被另一个贴纸覆盖的不良条形码

ITEM_CAP_SEALING-二次密封：为了防止泄漏，产品需要有以下类型的二次密封：感应密封、安全环、夹子、热收缩塑料袋或装箱。

ITEM_DEBUNDLE-拆包：需要拆开标有单独销售的一组商品。	否，但如果要填写prepCategory或者prepTypes，两者需同时填写	[array]	
phoneNumber	电话号码	是	[string]	
planName	计划名称	否	[string]	
positionType	分仓方式(1-先装箱再分仓，2-先分仓再装箱)	是	[string]	
postalCode	邮政编码	是	[string]	
remark	备注	否	[string]	
shipperName	发货方名称	是	[string]	
sid	领星店铺ID，对应查询亚马逊店铺列表接口对应字段【sid】	是	[long]	
stateOrProvinceCode	州/省/地区	是	[string]	
请求cURL示例
curl --location 'https://openapi.lingxing.com/amzStaServer/openapi/inbound-plan/createInboundPlan?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "addressLine1": "详细街道地址1",
    "addressLine2": "详细街道地址2",
    "city": "城市",
    "companyName": "公司名称",
    "countryCode": "国家(地区）",
    "email": "邮箱",
    "inboundPlanItems": [{
        "expiration": "2023-11-20",
        "labelOwner": "AMAZON",
        "msku": "msku",
        "prepOwner": "NONE",
        "quantity": 1
    }],
    "phoneNumber": "",
    "planName": "计划名称",
    "positionType": "1",
    "postalCode": "邮政编码",
    "remark": "备注",
    "shipperName": "发货方名称",
    "sid": 1,
    "stateOrProvinceCode": ""
}'
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	消息提示	是	[string]	success
errorDetails	错误信息	是	[array]	
requestId	请求链路id	是	[string]	
responseTime	响应时间	是	[string]	2020-05-18 11:23:47
data	响应数据	是	[object]	
data>>errorMsg	错误信息	是	[string]	
data>>inboundPlanId	亚马逊任务编号	是	[string]	
data>>taskId	任务id	是	[string]	
data>>taskStatus	任务状态
process
success
failure
local_failure	是	[string]	
返回成功示例
{
    "code": 0,
    "message": "操作成功",
    "errorDetails": [],
    "requestId": "3b3d867e7d014971a580549f107c8c5a.1732773886069",
    "responseTime": "2024-11-28T14:04:46.069",
    "data": {
        "errorMsg": "错误信息",
        "inboundPlanId": "亚马逊任务编号",
        "taskId": "任务id",
        "taskStatus": "任务状态"
    }
}
复制
错误
复制成功
返回失败示例
{
    "code": 500,
    "message": "内部错误",
    "errorDetails": [{ 
        "code": "BadRequest", 
        "message": "商品CN0001需要提交prepCategory预处理分类、prepTypes预处理类型"
    }],
    "requestId": "0A754F01-588E-C019-2693-5B72CF31D914",
    "responseTime": "2021-11-11 18:26:35",
    "data": {
        "errorMsg": "错误信息",
        "inboundPlanId": "亚马逊任务编号",
        "taskId": "任务id",
        "taskStatus": "任务状态"
    }
}
复制
错误
复制成功
 上一章节
获取发货单头程物流-其他费类型
FBA
下一章节 
查询包装组
FBA