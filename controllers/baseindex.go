package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/beego_blog/models"
	"github.com/lisijie/goblog/util"
	"strings"
	"strconv"
)

type BaseIndexController struct {
	beego.Controller
	userid         int
	username       string
	moduleName     string
	controllerName string
	actionName     string
	cache          *util.LruCache
}

func (this *BaseIndexController) Prepare() {
	this.auth()
	this.conf()
}

func (this *BaseIndexController) conf() {
	var result []*models.Conf
	new(models.Conf).Query().All(&result)
	ConfList := make(map[string]string)
	for _, v := range result {
		ConfList[v.Name] = v.Value
	}

	this.Data["conf"] = ConfList

}

func (this *BaseIndexController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.username = ""
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userid, _ := strconv.Atoi(idstr)
		if userid > 0 {
			var user models.User
			user.Id = userid
			authToken := beego.AppConfig.String("auth_token")
			if user.Read() == nil {
				authkey := util.Md5([]byte(authToken + "|" + user.Password))
				if password == authkey {
					this.userid = user.Id
					this.username = user.UserName
					this.Data["user"] = user
				}
			}
		}
	}

	if this.username == "" {
		this.Ctx.SetCookie("auth", "")
	}
}
