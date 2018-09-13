package controllers

import "github.com/astaxie/beego"

type WelcomeController struct {
	beego.Controller
}

func (this *WelcomeController) Get() {
	name := this.GetString("name")
	this.Data["name"] = name
	this.TplName = "welcome.tpl"
}
