package main

import "github.com/astaxie/beego"

type LoginController struct {
  beego.Controller
}

func (this *LoginController) Get() {
  this.Ctx.WriteString("User Login Page in GET.")
}

func (this *LoginController) Post() {
  this.Ctx.WriteString("User Login Page in Post.")
}

func main(){
    beego.Router("/login", &LoginController{})
    beego.Run("0.0.0.0:3000")
}
