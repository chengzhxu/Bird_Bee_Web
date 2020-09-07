package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["Website"] = "beego.me.cn"
	this.Data["Email"] = "astaxie@gmail.com!!!!"
	this.TplName = "index.tpl"
}

func (this *UserController) Home() {
	var req = this.Ctx.Request.URL.Query()

	var website = req.Get("website")
	var email = req.Get("email")
	if website == "" {
		website = "beego.me.cc"
	}
	if email == "" {
		email = "astaxie@gmail.cc"
	}

	this.Data["Website"] = website
	this.Data["Email"] = email
	this.TplName = "index.tpl"
}

// @router /:objectId [get]
func (this *UserController) hill() {
	this.Data["Website"] = "beego.me.hill"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
