package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/qianbaidu/beego_blog/util"
	"github.com/qianbaidu/beego_blog/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}



func (this *UserController) Login() {
	par := make(map[string]string)
	par["username"] = strings.TrimSpace(this.GetString("username"))
	par["password"] = strings.TrimSpace(this.GetString("password"))
	remember := strings.TrimSpace(this.GetString("remember"))
	err := util.IsRequire(par)
	if err != nil {
		this.Data["json"] = util.JsonMsg(err)
		this.ServeJSON()
		return
	}
	var user models.User
	user.UserName = par["username"]
	err = user.Read("user_name")

	if err == nil && user.Password == util.Md5([]byte(par["password"])) {
		this.Data["json"] = util.JsonMsg("")
		authToken := beego.AppConfig.String("auth_token")
		authkey := util.Md5([]byte(authToken + "|" + user.Password))
		if remember == "forever" {
			this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
			println("-----auth---")
			println(this.Ctx.GetCookie("auth"))
		} else {
			this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
		}
	} else {
		this.Data["json"] = util.JsonMsg("Incorrect username and password")
	}
	this.ServeJSON()
	return
}



func (this *UserController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/", 302)
}


func (this *UserController) Register() {
	par := make(map[string]string)

	par["username"] = strings.TrimSpace(this.GetString("username"))
	par["password"] = strings.TrimSpace(this.GetString("password"))
	par["confirm_password"] = strings.TrimSpace(this.GetString("confirm_password"))
	err := util.IsRequire(par)
	if err != nil {
		this.Data["json"] = util.JsonMsg(err)
		this.ServeJSON()
		return
	}
	if (par["password"] != par["confirm_password"]) {
		this.Data["json"] = util.JsonMsg("The password must be the same.")
		this.ServeJSON()
		return
	}

	var user models.User
	user.UserName = par["username"]
	user.Password = util.Md5([]byte(par["password"]))
	if err := user.Insert(); err != nil {
		this.Data["json"] = util.JsonMsg(err)
		this.ServeJSON()
		return
	}
	this.Data["json"] = util.JsonMsg("success")
	this.ServeJSON()
	return

}

//配合前端模板接口通用，这里2个接口合并一起
func (this *UserController) Pass() {
	action := strings.TrimSpace(this.GetString("action"))
	if (action == "signin") {
		this.Login()
	} else if (action == "signup") {
		this.Data["json"] = map[string]interface{}{"ok": false, "msg": "Disable the registration", "error": ""}
		this.ServeJSON()
		return
		//this.Register()
	} else {
		this.Data["json"] = map[string]interface{}{"ok": false, "msg": "Disable the registration", "error": ""}
		this.ServeJSON()
		return
	}
	return
}
