package models

import (
    "time"

    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/orm"
)

// VisitLogs 短网址访问日志
type VisitLogs struct {
    Id              int64
    Token           string
    RequestUrl      string
    UserHash        string
    UserAgent       string
    RefererUrl      string
    ClientAddress   string
    HttpMethod      string
    VisitTime       int64
}

// ShortenStat 短网址统计信息
type ShortenStat struct {
    Visits      int64
    VisitUsrs   int64
}

func init() {
    orm.RegisterModel(new(VisitLogs))
}

// AddVisitLog 写入一条请求日志
func AddVisitLog(l *VisitLogs) int64 {
    o := orm.NewOrm()
    id, err := o.Insert(l)

    if err != nil {
        logs.Error(err) // 写入失败，日志记录一条错误信息
    } else {
        var s ShortenStat
        err := o.Raw(`select count(0) as visits, count(distinct user_hash) as visit_usrs
            from visit_logs where token = ?`, l.Token).QueryRow(&s)
        if err == nil {
            o.Raw(`update shortens set visits = ?, visit_usrs = ?, last_visited = ? where token=?`,
                s.Visits, s.VisitUsrs, time.Now().Unix(), l.Token).Exec()
        }
    }

    return id
}
