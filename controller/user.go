package controller

import (
	"chat-server/internal/util"
	"chat-server/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	BaseController
}

func (this *UserController) FindUserById() {
	queryId := this.GetString(":id")
	fmt.Printf("id = %v\n", queryId)
	if queryId == "" {
		this.ErrorMsg(http.StatusBadRequest, "id can not be null", 1000)
		return
	}
	user, err := models.GetUserById(queryId)
	if err != nil {
		this.ErrorMsg(http.StatusBadRequest, "can not find", 1000)
		return
	}
	this.RespMsg(http.StatusOK, "success", user)
}

func (this *UserController) GetAllUsers() {
	u, err := models.GetAllUsers()
	if err != nil {
		fmt.Printf("err = %v\n", err)
		this.ErrorMsg(http.StatusBadRequest, "err", 1000)
	}
	this.RespMsg(http.StatusOK, "", u)
}

func (this *UserController) AddUser() {
	u := models.User{}
	data := this.Ctx.Input.RequestBody
	fmt.Printf("%v\n", string(data))
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Printf("err = %v\n", err)
		this.ErrorMsg(http.StatusBadRequest, "err", 1000)
		return
	}
	if u.Phone == "" {
		this.ErrorMsg(http.StatusBadRequest, "phone cannot be null", 1000)
		return
	}
	u.Id = util.NewUUID()
	err := models.AddUser(&u)
	if err != nil {
		this.ErrorMsg(http.StatusInternalServerError, "internal server error", 1001)
		return
	}
	ret, err := models.GetUserById(u.Id)
	if err != nil {
		fmt.Printf("err = %v", err)
		this.ErrorMsg(http.StatusBadRequest, "err", 1000)
		return
	}
	this.RespMsg(http.StatusOK, "", ret)
}

func (this *UserController) ModifyUserById() {
	uuid := this.GetString(":id")
	if uuid == "" {
		this.ErrorMsg(http.StatusBadRequest, "id can not be null", 1000)
		return
	}
	user := models.User{}
	data := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &user); err != nil {
		this.ErrorMsg(http.StatusBadRequest, "payload is error", 1000)
		return
	}
	if err := models.ModifyUserById(uuid, user); err != nil {
		this.ErrorMsg(http.StatusBadRequest, "Inner Error", 1000)
		return
	}
	retUser, err := models.GetUserById(uuid)
	if err != nil {
		fmt.Printf("err = %v", err)
		this.ErrorMsg(http.StatusBadRequest, "err", 1000)
		return
	}
	this.RespMsg(http.StatusOK, "", retUser)
}

func (this *UserController) DeleteAllUsers() {

}
