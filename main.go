package main

import (
	"github.com/hprose/hprose-go/hprose"
	"github.com/migege/list/models"

	"github.com/astaxie/beego"
)

func main() {
	service := hprose.NewHttpService()
	service.AddMethods(models.NewRoomModel())
	beego.Handler("/", service)
	beego.Run()
}
