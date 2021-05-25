package models

import (
	"chat-server/internal/util"
	"chat-server/internal/util/logs"
	"github.com/astaxie/beego/orm"
)

type Relation struct {
	Id       int    `json:"id" orm:"pk"`
	OwnerId  string `json:"owner_id"`
	FriendId string `json:"friend_id"`
}

func LinkRelation(user User, friend User) error {
	o := orm.NewOrm()
	re := &Relation{
		OwnerId:  user.Id,
		FriendId: friend.Id,
	}
	if err := o.Read(re); err != orm.ErrNoRows {
		if err == nil {
			return util.ERR_OneOrMultiRows
		}
		return err
	}
	_, err := o.Insert(re)
	if err != nil {
		logs.GetLogger().Errorf("Insert err = %v", err)
		return err
	}
	return nil
}

func DeleteRelationBetweenUsers(user User, friend User) error {
	o := orm.NewOrm()
	re := &Relation{
		OwnerId:  user.Id,
		FriendId: friend.Id,
	}
	if _, err := o.Delete(re); err != nil {
		logs.GetLogger().Errorf("Delete err = %v", err)
		return err
	}
	return nil
}

func IsFriend(user User, friend User) (bool, error) {
	o := orm.NewOrm()
	re := &Relation{
		OwnerId:  user.Id,
		FriendId: friend.Id,
	}
	err := o.Read(re)
	if err != nil {
		if err == orm.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
