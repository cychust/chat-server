package router

import (
	"chat-server/controller"
	"github.com/astaxie/beego"
)

func InitWebRouter() {
	beego.Router("/user", &controller.UserController{}, "GET:GetAllUser")
	beego.Router("/user", &controller.UserController{}, "POST:AddUser")
}
