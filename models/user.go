package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   string `json:"id" orm:"pk"`
	Name string `json:"name" orm:"size(50)"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func init() {
}
func AddUser(u *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return err
}

func GetUserById(uuid string) (user User, err error) {
	o := orm.NewOrm()
	queryUser := User{
		Id: uuid,
	}
	err = o.Read(&queryUser)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return
	}
	return queryUser, nil
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	o := orm.NewOrm()
	_, err := o.QueryTable(new(User)).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
