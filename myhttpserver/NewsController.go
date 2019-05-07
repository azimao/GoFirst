package main

import "com.azimao/myhttpserver/core"

type NewsController struct {
	core.MyController
}

func (this *NewsController) GET() {
	//this.Ctx.WriteString("this is GET NewsController")
	//this.Ctx.WriteJson(core.Map{"username":"azimao"})
	p := this.GetParam("username", "cs", "ts")
	this.Ctx.WriteString(p)
}

func (this *NewsController) POST() {
	//this.Ctx.WriteString("this is POST NewsController")
	//p := this.PostParam("username", "cs", "ts")
	//this.Ctx.WriteString(p)
	user := UserModel{}
	this.JsonParam(&user)
	this.Ctx.WriteJson(user)
}
