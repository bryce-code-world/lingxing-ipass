查询物流-头程物流商

支持查询物流-头程物流商，默认返回已启用现结api对接的物流商。

接口信息
API Path	请求协议	请求方式	令牌桶容量
/basicOpen/logistics/headLogisticsProvider/query/list	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
search	搜索参数对象	是	[object]	
search>>length	分页长度，每页显示的记录数	是	[int]	20
search>>page	页码，从1开始	是	[int]	1
search>>enabled	启用状态，枚举值：0-禁用, 1-启用，默认启用	否	[int]	1
search>>isAuth	是否api对接，枚举值：0-否, 1-是，默认是	否	[int]	1
search>>payMethod	结算方式，枚举值：1-现结, 2-月结，默认现结	否	[int]	1
search>>searchField	搜索字段，指定搜索的目标字段名称，code 代码 ，name 物流商，默认物流商	否	[string]	name
search>>searchValue	搜索值，用于模糊搜索物流商名称、编码等	否	[string]	顺丰
请求cURL示例
curl --location 'https://openapi.lingxing.com/basicOpen/logistics/headLogisticsProvider/query/list?access_token=value&timestamp=value&sign=value&app_key=value' \
--header 'Content-Type: application/json' \
--data '{
    "search": {
        "page": 1,
        "length": 20
    }
}'
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	响应码	是	[int]	
data	返回数据	是	[object]	
data>>total	总记录数	是	[int]	100
data>>providers	物流商列表	是	[array]	
data>>providers>>providerId	物流商id	否	[string]	
data>>providers>>name	物流商名	否	[string]	
data>>providers>>code	物流商代码	否	[string]	
data>>providers>>enabled	是否启用 0禁用 1 启用	否	[int]	
data>>providers>>logisticsType	类型
0 API物流
1 自定义物流
2 第三方仓物流
3 头程物流
4 平台物流	否	[int]	
data>>providers>>isAuth	是否api对接
0 否
1 是	否	[int]	
data>>providers>>supplierCode	授权方code	否	[int]	
data>>providers>>supplierName	授权方	否	[string]	
data>>providers>>status	授权状态
0 未授权
1 已授权	否	[int]	
data>>providers>>remark	备注	否	[string]	
data>>providers>>payMethod	结算方式	否	[int]	
data>>providers>>contactName	联系人	否	[string]	
data>>providers>>contactPhone	联系电话	否	[string]	
data>>providers>>creatorId	创建人id	否	[long]	
data>>providers>>creatorName	创建人名	否	[string]	
data>>providers>>createdAt	创建时间，Unix时间戳（秒）	否	[long]	
error_details	错误详情	是	[array]	
message	提示信息	是	[string]	
request_id	请求id	是	[string]	
response_time	响应时间	是	[string]	
total	总记录数	是	[int]	
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "a0d54debf93140f3b58d1ed81e8e3583.178.17255922733991817",
    "response_time": "2024-11-12 16:00:00",
    "data": {
        "total": 3,
        "providers": [
            {
                "providerId": "PROV001",
                "name": "顺丰速运",
                "code": "SFEXPRESS",
                "enabled": 1,
                "logisticsType": 1,
                "isAuth": 1,
                "supplierCode": 1001,
                "supplierName": "顺丰集团",
                "status": 1,
                "remark": "国内主要快递服务商",
                "payMethod": 1,
                "contactName": "王经理",
                "contactPhone": "13812345678",
                "creatorId": 10001,
                "creatorName": "管理员A",
                "createdAt": 1704067200
            },
            {
                "providerId": "PROV002",
                "name": "DHL国际快递",
                "code": "DHLINTL",
                "enabled": 1,
                "logisticsType": 3,
                "isAuth": 1,
                "supplierCode": 2001,
                "supplierName": "DHL集团",
                "status": 1,
                "remark": "全球领先的国际快递服务商",
                "payMethod": 2,
                "contactName": "李女士",
                "contactPhone": "13987654321",
                "creatorId": 10002,
                "creatorName": "管理员B",
                "createdAt": 1704067300
            },
            {
                "providerId": "PROV003",
                "name": "菜鸟裹裹",
                "code": "CAINIAO",
                "enabled": 1,
                "logisticsType": 4,
                "isAuth": 1,
                "supplierCode": 3001,
                "supplierName": "阿里巴巴",
                "status": 1,
                "remark": "平台聚合物流服务",
                "payMethod": 1,
                "contactName": "赵先生",
                "contactPhone": "13700112233",
                "creatorId": 10001,
                "creatorName": "管理员A",
                "createdAt": 1704067400
            }
        ]
    },
    "total": 3
}
复制
错误
复制成功
 上一章节
查询头程物流渠道列表
下一章节 
批量添加头程物流商