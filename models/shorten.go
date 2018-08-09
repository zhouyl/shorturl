package models

import (
    "time"
    "math/rand"

    h "shorturl/helpers"
    e "shorturl/errors"

    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/orm"
)

// TokenLength 默认 token 长度
const TokenLength = 6

// Shortens 短网址数据库 ORM 映射
type Shortens struct {
    Id          int64  `json:"id"`
    Token       string `json:"token"`
    Hash        string `json:"hash"`
    Business    string `json:"business"`
    Description string `json:"description"`
    LongUrl     string `json:"longUrl"`
    Visits      int64  `json:"visits"`
    VisitUsrs   int64  `json:"visitUsrs"`
    CreatedAt   int64  `json:"createdAt"`
    LastVisited int64  `json:"lastVisited"`
}

func init() {
    orm.RegisterModel(new(Shortens))
}

// RandomString 生成随机字符串
func RandomString(length int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < length; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
}

// ShortenInfo 根据 token 查询短网址信息
func ShortenInfo(token string) (*Shortens, error) {
    o := orm.NewOrm()
    s := &Shortens{}
    err := o.Raw("select * from shortens where token = ?", token).QueryRow(s)
    return s, err
}

// ShortenInfoByHash 根据 hash 值查询短网址信息
func ShortenInfoByHash(hash string) (*Shortens, error) {
    o := orm.NewOrm()
    s := &Shortens{}
    err := o.Raw("select * from shortens where hash = ?", hash).QueryRow(s)
    return s, err
}

// GenerateShorten 生成/获取一个短网址信息
func GenerateShorten(longUrl string, business string, description string) (*Shortens, error) {
    hash := h.Md5(longUrl + business)
    s, err := ShortenInfoByHash(hash)
    if err == nil {
        return s, nil
    }

    for i := 0; i < 100; i++ { // 循环 100 次尝试生成新的 token
        token := RandomString(TokenLength)
        if _, err := ShortenInfo(token); err != nil {
            s := &Shortens{
                Token: token,
                Hash: hash,
                LongUrl: longUrl,
                Business: business,
                Description: description,
                Visits: 0,
                VisitUsrs: 0,
                CreatedAt: time.Now().Unix(),
                LastVisited: 0,
            }

            o := orm.NewOrm()
            id, err := o.Insert(s)
            if err == nil {
                s.Id = id
                return s, nil
            }

            logs.Error(err) // 写入失败，日志记录一条错误信息
            return nil, err
        }
    }

    return nil, e.NewError(e.ShortenFail)
}
