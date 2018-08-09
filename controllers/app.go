package controllers

import (
	"strings"
    "fmt"
    "time"

    h "shorturl/helpers"
    m "shorturl/models"
    e "shorturl/errors"

    "github.com/astaxie/beego"
)

// AppController 默认控制器
type AppController struct {
    beego.Controller
}

// Response API 输出格式
type Response struct {
    Code        int         `json:"code"`
    Message     string      `json:"message"`
    Data        interface{} `json:"data,omitempty"`
    Timestamp   int64       `json:"timestamp"`
}

// RequireInputs 判断请求中是否包含指定的参数
func (c *AppController) RequireInputs(args ...string) (error) {
    for _, k := range args {
        v := c.GetString(k)
        if v == "" {
            return fmt.Errorf("Parameter [%s] is missing", k)
        }
    }

    return nil
}

// GetDefaultString 为 GetString 提供一个默认值
func (c *AppController) GetDefaultString(key string, def string) string {
    v := c.GetString(key)
    if v == "" {
        return def
    }
    return v
}

// NewResponse 构造一个规范的 api 输出响应
func (c *AppController) NewResponse(code int, data interface{}) {
    c.NewMsgResponse(code, e.GetMessage(code), data)
}

// NewMsgResponse 构造一个自定义 msg 的 api 输出响应
func (c *AppController) NewMsgResponse(code int, msg string, data interface{}) {
    c.Data["json"] = &Response{
        Code: code,
        Message: msg,
        Data: data,
        Timestamp: time.Now().Unix(),
    }
    c.ServeJSON()
}


// Shorten 为长链接生成一个 token
// curl -X GET 'http://127.0.0.1:8080/shorten?business=sms-marketing&longUrl=https%3A%2F%2Fwww.myhost.com%2Fregister&description=用户注册'
func (c *AppController) Shorten() {
    err := c.RequireInputs("longUrl")
    if err != nil {
        c.NewMsgResponse(e.ParamMissing, fmt.Sprint(err), nil)
        return
    }

    info, err := m.GenerateShorten(c.GetString("longUrl"), c.GetString("business"), c.GetString("description"))

    if err == nil {
        c.NewResponse(e.OK, info)
    } else {
        c.NewResponse(e.ShortenFail, nil)
    }
}

// Query 检测 token 的有效性，获取短链接信息
// curl -X GET 'http://127.0.0.1:8080/query?token=R1QJTocA'
func (c *AppController) Query() {
    err := c.RequireInputs("token")
    if err != nil {
        c.NewMsgResponse(e.ParamMissing, fmt.Sprint(err), nil)
        return
    }

    token := c.GetString("token")
    info, err := m.ShortenInfo(token)

    if err == nil {
        c.NewResponse(e.OK, info)
    } else {
        c.NewResponse(e.InvalidToken, nil)
    }
}

// Visit 根据 token，请求访问 longUrl
// curl -X GET 'http://127.0.0.1:8080/visit?token=R1QJTocA'
func (c *AppController) Visit() {
    err := c.RequireInputs("token", "requestUrl", "userHash","userAgent")
    if err != nil {
        c.NewMsgResponse(e.ParamMissing, fmt.Sprint(err), nil)
        return
    }

    l := m.VisitLogs{
        Token: c.GetString("token"),
        RequestUrl: c.GetString("requestUrl"),
        UserHash: c.GetString("userHash"),
        UserAgent: c.GetString("userAgent"),
        RefererUrl: c.GetString("refererUrl"),
        ClientAddress: c.GetString("clientAddress", "0.0.0.0"),
        HttpMethod: strings.ToUpper(c.GetDefaultString("httpMethod", "GET")),
        VisitTime: time.Now().Unix(),
    }

    if !h.IsHTTPMethod(l.HttpMethod) {
        c.NewResponse(e.InvalidHTTPMethod, nil)
        return
    }

    info, err := m.ShortenInfo(l.Token)
    if err != nil {
        c.NewResponse(e.InvalidToken, nil)
        return
    }

    m.AddVisitLog(&l)
    c.NewResponse(e.OK, info)
}
