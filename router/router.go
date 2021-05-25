package router

import (
	"chat-server/controller"
	"github.com/astaxie/beego"
)

func InitWebRouter() {
	InitUserRouter()
	InitRelationRouter()
}

func InitUserRouter() {
	beego.Router("/user/:id", &controller.UserController{}, "GET:FindUserById")
	beego.Router("/user/:id", &controller.UserController{}, "PATCH:ModifyUserById")
	beego.Router("/user", &controller.UserController{}, "GET:GetAllUsers")
	beego.Router("/user", &controller.UserController{}, "POST:AddUser")
	beego.Router("/user", &controller.UserController{}, "DELETE:DeleteAllUsers")
}

func InitRelationRouter() {
	// beego.Router("")
	beego.Router("/user/addFriend/:id", &controller.RelationController{}, "POST:AddFriend")
	beego.Router("/user/deleteFriend/:id", &controller.RelationController{}, "POST:DeleteFriend")
	beego.Router("/user/isFriend/:id", &controller.RelationController{}, "POST:IsFriend")
}
