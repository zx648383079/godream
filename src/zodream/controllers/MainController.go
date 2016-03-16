package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "index"
	this.TplName = "index.tpl"
	//this.Ctx.WriteString("hello")
}