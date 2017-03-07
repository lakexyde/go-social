package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Timeline() {
	this.TplName = "user/index.tpl"
}

func (this *UserController) Profile(){
	n := this.Ctx.Input.Param(":user")
	v := this.GetSession("lakex")
	if v == nil {
		this.Redirect("/", 302)
	} else {
		this.Data["g"] = v
	}

	this.Data["u"] = n
	this.TplName = "user/profile.tpl"
}