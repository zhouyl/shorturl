package errors

import (
    "errors"
)

const (
    // OK 请求成功
    OK = 200

    // ParamMissing 缺少参数
    ParamMissing = 601

    // ShortenFail 生成短网址失败
    ShortenFail = 10001

    // InvalidToken 无效 token
    InvalidToken = 10002

    // InvalidHTTPMethod 无效的 HTTP method
    InvalidHTTPMethod = 10003
)

// ErrorCodes 错误代码对应消息内容
var ErrorCodes = map[int] string{
    OK: "OK",
    ParamMissing: "Missing required parameter",
    ShortenFail: "Generate shorten url failed",
    InvalidToken: "Invalid token",
    InvalidHTTPMethod: "Invalid HTTP method",
}

// GetMessage 根据错误代码，获取消息
func GetMessage(code int) string {
    msg, ok := ErrorCodes[code]

    if !ok {
        msg = "Unknown error"
    }

    return msg
}

// NewError 创建一个新的 error
func NewError(code int) error {
    return errors.New(GetMessage(code))
}
