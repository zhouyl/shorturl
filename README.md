# 短网址服务

就职于某司时，尝试用 go 开发的一个短网址服务，服务于各项业务，例如短信发送、消息推送等。

## 一、项目说明

### 1. Makefile 命令

* make build / 重新编译文件到 ./shorturl
* make testing / 发布到 t0111 测试环境
* make release / 发布到生产环境

### 2. shorturl.sh 脚本

请使用以下命令管理

`shorturl.sh start|stop|status|restart`

## 二、API 文档

### 1. 请求生成短网址

`curl -X GET 'http://127.0.0.1:8080/shorten?longUrl=https%3A%2F%2Fwww.myhost.com%2Fregister&business=sms-marketing'`

POST 参数：

    longUrl     <必须>  需要生成短网址的 url
    business    <必须>  用于标识业务来源，以便数据统计，例如：sms-marketing

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "Id": 1,
    "token": "tP1ABC",
    "hash": "23f86214b729cbd1d2959f9d14188ce6",
    "longUrl": "https://www.myhost.com/register",
    "business": "sms-marketing",
    "visits": 0,
    "createdAt": 1529049275,
    "lastVisited": 1529049275
  },
  "timestamp": 1529049791
}
```

### 2. 查询短网址

`curl -X GET 'http://host:port/query?token=tP1ABC'`

GET 参数：

    token  <必须>  短网址 token

成功：

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "Id": 1,
    "token": "tP1ABC",
    "hash": "23f86214b729cbd1d2959f9d14188ce6",
    "longUrl": "https://www.myhost.com/register",
    "business": "sms-marketing",
    "visits": 0,
    "createdAt": 1529049275,
    "lastVisited": 1529049275
  },
  "timestamp": 1529049727
}
```

失败：

```json
{
  "code": 10002,
  "message": "Invalid token",
  "timestamp": 1529049777
}
```

### 2. 访问短网址

`curl -X POST 'http://host:port/visit' -d 'token=tP1ABC&...'`

POST 参数：

    token           <必须>  短网址 token
    business        <必须>  用于标识业务来源，以便数据统计，例如：sms-marketing
    requestUrl      <必须>  请求页面的 url，例如: http://jrq.cn/tP1ABC
    userHash        <必须>  获取页面 sid 然后 md5 得到，用于分析 UV
    userAgent       <必须>  HTTP_USER_AGENT
    clientAddress   <必须>  客户端 IP 地址
    method          [可选]  默认 GET
    refererUrl      [可选]  HTTP_REFERER
