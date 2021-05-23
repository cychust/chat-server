package controller

import (
	"chat-server/internal/util"
	"chat-server/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) GetAllUser() {
	u, err := models.GetAllUsers()
	if err != nil {
		fmt.Printf("err = %v\n", err)
		this.Data["Content"] = "error"
		this.ServeJSON()
	}
	this.Data["json"] = u
	this.ServeJSON()
}

func (this *UserController) AddUser() {
	u := models.User{}
	data := this.Ctx.Input.RequestBody
	fmt.Printf("%v\n", string(data))
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Printf("err = %v\n", err)
		this.Data["Content"] = "params is error"
		this.ServeJSON()
		return
	}
	u.Id = util.NewUUID()
	_ = models.AddUser(&u)
	ret, err := models.GetUserById(u.Id)
	if err != nil {
		fmt.Printf("err = %v", err)
		this.Data["Content"] = "Inner  error"
		this.ServeJSON()
		return
	}
	fmt.Printf("%v", ret)
	this.Data["json"] = ret
	this.ServeJSON()
}
