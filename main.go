package main

import (
    "fmt"
    "os"
    "path/filepath"

    "shorturl/helpers"
    _ "shorturl/routers"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqldsn"))
}

func main() {
    // 写入 pid 文件
    pidfile := beego.AppConfig.DefaultString("pidfile", "/tmp/shorturl.pid")
    pid := os.Getpid()
    helpers.WriteFile(pidfile, fmt.Sprintf("%d", pid), 0644)
    beego.Info(fmt.Sprintf("pid file: %s, pid: %d", pidfile, pid))

    // 设定日志配置
    logFile, _ := filepath.Abs(beego.AppConfig.DefaultString("logfile", "logs/app.log"))
    logConf := `{
        "filename": "%s",
        "maxlines": 1000000,
        "maxsize": 100000000,
        "maxdays": 30
    }`
    helpers.MakedirIfNotExists(filepath.Dir(logFile), 0777)
    beego.SetLogger("file", fmt.Sprintf(logConf, logFile))

    beego.Run()
}
