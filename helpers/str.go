package helpers

import (
    "fmt"
    "strings"
    "crypto/md5"
)

// HTTPMethods 有效的 http 请求方式
var HTTPMethods = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "OPTIONS", "CONNECT", "PATCH"}

// Md5 获取一个字符串的 md5 值
func Md5(str string) string {
    data := []byte(str)
    hash := md5.Sum(data)
    return fmt.Sprintf("%x", hash)
}

// IsHTTPMethod 判断是否有效的 HTTP method
func IsHTTPMethod(method string) bool {
    method = strings.ToUpper(method)
    for _, m := range HTTPMethods {
        if method == m {
            return true
        }
    }
    return false
}
