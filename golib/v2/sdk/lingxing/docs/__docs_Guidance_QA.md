常见问题案例
1. 获取access_token时报以下错误
{
    "code": "0001001",
    "msg": "rsa private decrypt error",
    "data": null
}
复制
错误
复制成功

原因：因为appSecret中可能有特殊字符，在传输中会转义，因此appSecret用Rsa解密时报错。
解决方法： appSecret利用urlencode.encode()进行转义后传输。

2. 生成签名不正确
{
    "code": "2001006",
    "msg": " api sign not correct",
    "data": {
        "throwable": " api sign not correct",
        "throwTime": "2022-01-20 15:37:55",
        "message": null,
        "stackTrace": null,
        "metadata": {}
    }
}
复制
错误
复制成功

原因：生成的签名不正确。
解决方法：
1、检查参与生成签名的参数是否正确。
2、当参数的值为空时，不需要把该参数代入计算签名。
3、参与签名的时间戳长度取10位。
4、检查签名是否进行了URLEncoder.encode()进行url字符转义。
5、post请求url上只能带appId，token，timestamp，sign【转义后的值】这四个参数，其他业务请求参数放在body里。
6、传入body参数若为数组、List等集合类型时，需要将其先转入为String类型后再进行传入从而生成sign；在生成sign之后，需要重新将数组、List集合本身再次进行覆盖，防止出现参数不合法的问题；参考SDK下载-Java代码示例com.asinking.com.openapi.sdk.samples.OrderListDemo中的写法
7、使用python调用接口时，为了防止大写False/True导致签名错误，Boolean字段统一用字符串'true'或’false‘;

3. 频繁出现access_token过期
{
    "code":"2001005",
    "msg":" access token not match"
}
复制
错误
复制成功

解决方法：检查自己本地是否存在别的应用一直在重复请求导致access_token频繁变换。

4. ip被限制访问
{
        "code": "3001002",
        "msg": "ip not permit, please add ip to white list first.",
        "data": {
                "throwable": "ip not permit, please add ip to white list first.",
                "throwTime": "2021-12-27 15:49:19",
                "message": null,
                "stackTrace": null,
                "metadata": {}
        }
}
复制
错误
复制成功

原因：应用所在服务器的ip不在白名单中。
解决方法：ip未加入白名单，确认发起ip地址后在【设置】>【业务配置】>【全局】>【开放接口】自行增加即可。

5. 服务器返回400
{
    "code": 400,
    "message": "服务不存在",
    "error_details": [],
    "request_id": "903BFE57-E7B3-81C2-4E4A-29EDA32C3B59",
    "response_time": "2021-12-27 16:21:14",
    "data": [],
    "total": 0
}
复制
错误
复制成功

原因：请求url路径不正确。
解决方法：检查url地址是否正确。

6. 服务器返回405
Response error, status code: 405, body: <html>
<head><title>405 Not Allowed</title></head>
<body>
<center><h1>405 Not Allowed</h1></center>
<hr><center>nginx</center>
</body>
</html>
复制
错误
复制成功

原因：请求url路径不正确。
解决方法：检查url地址是否正确。

7. 服务器返回500
{
    "code": "500",
    "msg": "404 NOT_FOUND",
    "data": {
        "throwable": "404 NOT_FOUND",
        "throwTime": "2021-12-27 16:22:09",
        "message": null,
        "stackTrace": null,
        "metadata": {}
    }
}
复制
错误
复制成功

原因：请求url路径不正确。
解决方法：检查请求地址是否正确，注意 Domain与 API Path 的拼接中不要有多余的 ”/“【反斜杠】。

8. 公共请求参数不完整
{
    "code": "3001001",
    "msg": "missing query param(access_token,sign,timestamp,app_key)",
    "data": {
        "throwable": "missing query param(access_token,sign,timestamp,app_key)",
        "throwTime": "2024-07-11 14:58:44",
        "message": null,
        "stackTrace": null,
        "metadata": {}
    }
}
复制
错误
复制成功

原因：公共请求参数不完整。
解决方法：在Query Params参数中，需要传入包括access_token, app_key, sign, timestamp四个公共请求参数

9. 参数不合法
{
    "code": "102",
    "msg": "参数不合法",
    "data": {
        "throwable": "参数不合法",
        "throwTime": "2024-07-11 14:58:44",
        "message": null,
        "stackTrace": null,
        "metadata": {}
    }
}
复制
错误
复制成功

原因：数组、List集合在生成sign过程中会将其数据类型转为string类型后传入body参数当中，而转义后的数据不符合接口文档传入参数类型的规范
解决方法：在生成sign之后，添加一段代码将body参数中对应的值重新以本身数组或List集合类型进行覆盖即可

10. 限流拦截
{
    "code":103,
    "message":"请求过于频繁,请稍后再试 [C0C1B855-48A2-6BB6-5636-5AC80042386C]",
    "error_details":[],
    "request_id":"4B750B70-2026-FE98-B88F-400B91D3F1E4",
    "response_time":"2024-07-21 03:10:37",
    "data":[],
    "total":0
}
复制
错误
复制成功

原因：请求速率过高导致被限流拦截 解决办法：根据限流规则去设计发送请求的频率，防止被系统拦截

11. appId被禁用
{
    "code": "401",
    "msg": "授权被禁用，请检查授权状态",
    "data": null
}
复制
错误
复制成功

解决办法：请求接口的appId被禁用了，请联系erp超级管理员在 设置-全局-开放接口 页面，为该appId解除禁用

12. appId已经失效
{
    "code": "403",
    "msg": "授权失效，请更新授权有效期",
    "data": null
}
复制
错误
复制成功

解决办法：请求接口的appId过期，请联系erp超级管理员在 设置-全局-开放接口 页面，为该appId更新授权有效期

13. appId没有接口权限
{
    "msg": "授权失效，请更新授权有效期或检查权限",
    "code": "403",
    "data": null
}
复制
错误
复制成功

解决办法：appId账号没有其请求接口的权限，请联系erp超级管理员在 设置-全局-开放接口 页面，为该appId授权接口

 上一章节
领星 API 接入指南
下一章节 
全局错误码说明