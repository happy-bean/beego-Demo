package routers

import (
	"beego-Demo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Get("/hello",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Post("/post",func(ctx *context.Context){
		ctx.Output.Body([]byte("post"))
	})

	beego.Any("/any",func(ctx *context.Context){
		ctx.Output.Body([]byte("any"))
	})

}
