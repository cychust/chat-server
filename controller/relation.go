package controller

import (
	"chat-server/internal/util/logs"
	"chat-server/models"
	"encoding/json"
	"net/http"
)

type RelationController struct {
	BaseController
}

type AddFriendData struct {
	FriendId string `json:"friend_id"`
}

type DeleteFriendData struct {
	FriendId string `json:"friend_id"`
}

type SearchFriendData struct {
	FriendId string `json:"friend_id"`
}

func (this *RelationController) AddFriend() {
	queryId := this.GetString(":id")
	logs.GetLogger().Infof("query id is %s", queryId)
	if queryId == "" {
		this.ErrorMsg(http.StatusBadRequest, "id can not be null", 1000)
		return
	}
	data := new(AddFriendData)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, data)
	if err != nil {
		logs.GetLogger().Errorf("unmarshal err = %v", err)
		this.ErrorMsg(http.StatusBadRequest, "request body err", 1000)
		return
	}
	user, err := models.GetUserById(queryId)
	if err != nil {
		logs.GetLogger().Errorf("user can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "user can not find", 1000)
		return
	}
	friend, err := models.GetUserById(data.FriendId)
	if err != nil {
		logs.GetLogger().Errorf("friend can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "friend can not find", 1000)
		return
	}
	if err := models.LinkRelation(user, friend); err != nil {
		logs.GetLogger().Errorf("link relationship between user and friend err: %v", err)
		this.ErrorMsg(http.StatusInternalServerError, "internal error", 1001)
		return
	}
	this.RespMsg(http.StatusOK, "success", nil)
}

func (this *RelationController) DeleteFriend() {
	queryId := this.GetString(":id")
	logs.GetLogger().Infof("query id is %s", queryId)
	if queryId == "" {
		this.ErrorMsg(http.StatusBadRequest, "id can not be null", 1000)
		return
	}
	data := new(DeleteFriendData)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, data)
	if err != nil {
		logs.GetLogger().Errorf("unmarshal err = %v", err)
		this.ErrorMsg(http.StatusBadRequest, "request body err", 1000)
		return
	}
	user, err := models.GetUserById(queryId)
	if err != nil {
		logs.GetLogger().Errorf("user can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "user can not find", 1000)
		return
	}
	friend, err := models.GetUserById(data.FriendId)
	if err != nil {
		logs.GetLogger().Errorf("friend can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "friend can not find", 1000)
		return
	}
	if err := models.DeleteRelationBetweenUsers(user, friend); err != nil {
		logs.GetLogger().Errorf("unlink relationship between user[%s] and friend[%s]", user.Id, friend.Id)
		this.ErrorMsg(http.StatusInternalServerError, "unlink relationship error", 1000)
		return
	}
	this.RespMsg(http.StatusOK, "success", nil)
}

func (this *RelationController) IsFriend() {
	queryId := this.GetString(":id")
	logs.GetLogger().Infof("query id is %s", queryId)
	if queryId == "" {
		this.ErrorMsg(http.StatusBadRequest, "id can not be null", 1000)
		return
	}
	data := new(SearchFriendData)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, data)
	if err != nil {
		logs.GetLogger().Errorf("unmarshal err = %v", err)
		this.ErrorMsg(http.StatusBadRequest, "request body err", 1000)
		return
	}
	user, err := models.GetUserById(queryId)
	if err != nil {
		logs.GetLogger().Errorf("user can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "user can not find", 1000)
		return
	}
	friend, err := models.GetUserById(data.FriendId)
	if err != nil {
		logs.GetLogger().Errorf("friend can not find: %v", err)
		this.ErrorMsg(http.StatusBadRequest, "friend can not find", 1000)
		return
	}
	var isFriend bool
	if isFriend, err = models.IsFriend(user, friend); err != nil {
		logs.GetLogger().Errorf("unlink relationship between user[%s] and friend[%s]", user.Id, friend.Id)
		this.ErrorMsg(http.StatusInternalServerError, "unlink relationship error", 1000)
		return
	}
	retData := make(map[string]bool)
	retData["IsFriend"] = isFriend
	this.RespMsg(http.StatusOK, "success", retData)
}
