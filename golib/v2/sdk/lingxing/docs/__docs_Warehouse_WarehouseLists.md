查询仓库列表

支持查询本地仓库列表信息，对应系统【设置】>【仓库设置】仓库列表

接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/data/local_inventory/warehouse	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
type	仓库类型：
1 本地仓【默认值】
3 海外仓
4 亚马逊平台仓
6 AWD仓	否	[int]	1
sub_type	海外仓子类型：
1 无API海外仓
2 有API海外仓【此参数只在type=3生效】	否	[int]	2
is_delete	是否删除，多个使用英文逗号分隔：
0 未删除【默认值】
1 已删除	否	[string]	0
offset	分页偏移量，默认0	否	[int]	0
length	分页长度，默认1000条	否	[int]	1000
请求示例
{
    "type": 1,
    "sub_type": 2,
    "is_delete": 0,
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
request_id	请求链路id	是	[string]	502B9DD9-1BA0-03C5-6C61-D77C830440A6
response_time	响应时间	是	[string]	2020-05-18 11:23:47
total	总数	是	[int]	5
data	响应数据	是	[array]	
data>>wid	系统仓库id	是	[int]	1300
data>>name	仓库名	是	[string]	00000
data>>type	仓库类型：
1 本地仓
3 海外仓
4 平台仓
6 AWD仓	是	[int]	1
data>>is_delete	是否删除：0 未删除，1 已删除	是	[string]	0
data>>t_country_area_name	第三方仓库国家/地区	是	[string]	英国
data>>t_status	状态:0 未启用 ，1 启用	是	[int]	1
data>>t_warehouse_code	第三方仓库代码	是	[string]	UKBM
data>>t_warehouse_name	第三方仓库名	是	[string]	UKBM Warehouse
data>>country_code	国家代码	是	[string]	UK
data>>wp_id	服务商ID，仅type=3且仓库为第三方海外仓时有值	是	[int]	1
data>>wp_name	系统服务商名称	是	[string]	0
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "D7FFE246-BE0D-D8C9-355D-ADB1EB860C3C",
    "response_time": "2024-07-31 14:33:14",
    "data": [
        {
            "wid": 1300,
            "type": 1,
            "name": "00000",
            "is_delete": 0,
            "wp_id":1,
            "wp_name":""
            "t_country_area_name": "英国",
              "t_status": 1,
              "t_warehouse_code": "UKBM",
             "t_warehouse_name": "UKBM Warehouse",
              "country_code": "UK"
        },
        {
            "wid": 858,
            "type": 1,
            "name": "测试仓",
            "is_delete": 0,
            "wp_id":1,
            "wp_name":""
              "t_country_area_name": "英国",
              "t_status": 1,
              "t_warehouse_code": "UKBM",
              "t_warehouse_name": "UKBM Warehouse",
              "country_code": "UK"
        }
    ],
    "total": 2
}
复制
错误
复制成功
 上一章节
采购单整单结束到货
下一章节 
添加/修改仓库