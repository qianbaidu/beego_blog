package routers

import (
	"github.com/qianbaidu/beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/page/:page", &controllers.IndexController{},"*:Get")
	beego.Router("/:page", &controllers.IndexController{})
	beego.Router("/category/:cid", &controllers.CategoryController{})
	beego.Router("/category/:cid/:page", &controllers.CategoryController{})
	beego.Router("/article/:id([0-9]+).html", &controllers.ArticleController{})
	beego.Router("/article/:id", &controllers.ArticleController{})
	beego.Router("/tag/:id", &controllers.TagController{})


	//beego.Router("/api/pass", &controllers.UserController{}, "post:Pass")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/about", &controllers.AboutController{})

	//APIS
	ns :=
		beego.NewNamespace("/v1/",
			//此处正式版时改为验证加密请求
			//beego.NSCond(func(ctx *context.Context) bool {
			//	if ua := ctx.Input.Request.UserAgent(); ua != "" {
			//		return true
			//	}
			//	return false
			//}),
			beego.NSNamespace("/api",
				//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
				beego.NSNamespace("/user",
					// /api/ios/create/node/
					beego.NSRouter("/pass", &controllers.UserController{}, "post:Pass"),
				),
			),
		)

	beego.AddNamespace(ns)

}
