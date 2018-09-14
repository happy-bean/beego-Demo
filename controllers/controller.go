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

type AnnotationController struct {
	beego.Controller
}

func (c *AnnotationController) URLMapping() {
	c.Mapping("Antest", c.Antest)
}

// @router /antest/:key [get]
func (this *AnnotationController) Antest() {
	this.Ctx.Output.Body([]byte("annotation controller"))
}