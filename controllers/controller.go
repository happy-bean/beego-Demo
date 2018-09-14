package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-Demo/models"
	"fmt"
)

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
	c.Mapping("Jsontest", c.Jsontest)
}

// @router /antest/:key [get]
func (this *AnnotationController) Antest() {
	this.Ctx.Output.Body([]byte("annotation controller"))

	//session使用
	//v := this.GetSession("sess")
	//if v == nil {
	//	this.SetSession("sess", int(1))
	//} else {
	//	this.SetSession("sess", v.(int)+1)
	//}
}

// @router /jsontest/:key [post]
func (this *AnnotationController) Jsontest() {
	var user models.User
	println(string(this.Ctx.Input.RequestBody))
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	println(user.Name)
	jsons, errs := json.Marshal(user)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	this.Ctx.Output.Body(jsons)
}
