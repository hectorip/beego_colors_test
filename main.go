package main

import (
	// _ "colors_test/routers"
	beego "github.com/beego/beego/v2/server/web"
)


type MainController struct {
    beego.Controller
}

type color struct {
	ColorName string `json:"color_name"`
}
  
  
func (this *MainController) Get() {
	color_one := color{ColorName:"White"}
	this.Data["json"] = &color_one
	this.ServeJSON()
}

// func (this *MainController) Get() {
//     this.Ctx.WriteString("Colors Beego test")
// }

func main() {
    beego.Router("/color", &MainController{})
    beego.Run()
}