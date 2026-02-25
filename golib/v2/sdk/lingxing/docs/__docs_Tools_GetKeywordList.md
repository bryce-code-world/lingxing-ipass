关键词列表
接口信息
API Path	请求协议	请求方式	令牌桶容量
/erp/sc/routing/tool/toolKeywordRank/getKeywordList	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
mid	国家id	否	[int]	1
start_date	开始日期，格式：Y-m-d	否	[string]	2024-08-01
end_date	结束日期，格式：Y-m-d	否	[string]	2024-08-01
offset	分页偏移量，默认0	是	[int]	0
length	分页长度，默认20，最大值为2000	是	[int]	20
请求示例
{
    "offset": 0,
    "length": 20,
    "mid": 1,
    "start_date": "2024-08-01",
    "end_date": "2024-08-01"

}
复制
错误
复制成功
返回结果

Json Object

参数名	说明	必填	类型	示例
code	状态码，0 成功	是	[int]	0
message	提示信息	是	[string]	success
error_details	错误信息	是	[array]	
request_id	请求链路id	是	[string]	8D680B8B-3C6F-3E16-F055-73DC70EAD977
response_time	响应时间	是	[string]	2022-11-23 19:52:36
data	响应数据	是	[array]	
data>>id	记录唯一id	是	[int]	3796
data>>key_word	关键词	是	[string]	11
data>>rank	排名	是	[number]	10
data>>page	页码	是	[number]	100
data>>create_time	开始监控时间	是	[string]	2022-09-09 10:52
data>>monitor_time	更新时间	是	[string]	2022-12-01 12:01
data>>keyword_remark	关键词备注	是	[string]	测试备注
data>>asin	监控的asin	是	[string]	B0ABC43215
data>>parent_asin	父asin	是	[string]	B0ABC43215
data>>title	标题	是	[string]	标题237
data>>keyword_num	关键词数量	是	[number]	16
data>>asin_remark	监控asin的备注	是	[string]	1221
data>>country	国家	是	[string]	美国
data>>creator	创建人	是	[string]	用户1
data>>monitors	监控人	是	[array]	["用户1","用户2"]
data>>asin_create_time	监控asin的创建时间	是	[string]	2020-03-31 09:10
data>>current_page_rank	当前页排名:
0 在获取中
1000 未进前6页	是	[number]	1
data>>sbv_page	sbv排名:
0 在获取中
1000 未进前6页
-1 没有sbv排名	是	[number]	1
data>>rank_text	排名说明	是	[string]	1
data>>sbv_text	sbv排名说明	是	[string]	1
data>>is_sponsored	监控范围:
1 广告排名
0 自然排名	是	[number]	1
data>>type	监控指标:
1 PC端
2 移动端	是	[number]	1
total		是	[int]	1
返回成功示例
{
    "code": 0,
    "message": "success",
    "error_details": [],
    "request_id": "8D680B8B-3C6F-3E16-F055-73DC70EAD977",
    "response_time": "2022-11-23 19:52:36",
    "data": [
          {
            "id": 386,
            "key_word": "Size Blenders",
            "rank": 1000,
            "page": 1000,
            "create_time": "2024-08-01 16:38",
            "current_page_rank": 1000,
            "sbv_page": -1,
            "rank_text": "-(未进前6页)",
            "sbv_text": "",
            "is_sponsored": 0,
            "type": 1,
            "monitor_time": "2024-12-19 08:00",
            "keyword_remark": "",
            "asin": "B0CT6P5N4V",
            "parent_asin": "",
            "title": "",
            "keyword_num": 4,
            "asin_remark": "",
            "country": "美国",
            "creator": "李坤莲",
            "monitors": [
                "李坤莲"
            ],
            "asin_create_time": "2024-08-01 16:38"
        }
    ],
    "total": 1
}
复制
错误
复制成功
 上一章节
查询业绩通知列表
下一章节 
查询竞品监控列表