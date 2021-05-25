package models

import (
	"chat-server/internal/util/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type State struct {
	Id          int       `json:"id" orm:"pk"`
	Uuid        string    `json:"uuid"`
	State       bool      `json:"state"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func GetOrCreateState(user User) (*State, error) {
	o := orm.NewOrm()
	state := &State{
		Uuid: user.Id,
	}
	err := o.Read(state)
	if err == orm.ErrNoRows {
		_, err = o.Insert(state)
		if err != nil {
			logs.GetLogger().Errorf("Insert State err:%v", err)
			return nil, err
		}
		return state, nil
	}
	if err != nil {
		logs.GetLogger().Errorf("Read State err: %v", err)
		return nil, err
	}
	return state, nil
}

func (s *State) Offline() error {
	s.OfflineTime = time.Now()
	s.State = false
	o := orm.NewOrm()
	_, err := o.Update(s)
	if err != nil {
		logs.GetLogger().Errorf("Update err: %v", err)
		return err
	}
	return nil
}
func (s *State) Online() error {
	s.OnlineTime = time.Now()
	s.State = true
	o := orm.NewOrm()
	_, err := o.Update(s)
	if err != nil {
		logs.GetLogger().Errorf("Update err: %v", err)
		return err
	}
	return nil
}

func (s *State) Delete() error {
	o := orm.NewOrm()
	state := &State{
		Uuid: s.Uuid,
	}
	_, err := o.Delete(state)
	if err != nil {
		logs.GetLogger().Errorf("Delete error:%v", err)
		return err
	}
	return nil
}
