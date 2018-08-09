package routers

import (
    "shorturl/controllers"

    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/shorten", &controllers.AppController{}, "get:Shorten")
    beego.Router("/query", &controllers.AppController{}, "get:Query")
    beego.Router("/visit", &controllers.AppController{}, "*:Visit")
}
