package routers

import (
	"github.com/qianbaidu/beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/test", &controllers.MainController{})
    beego.Router("/", &controllers.IndexController{})
}
