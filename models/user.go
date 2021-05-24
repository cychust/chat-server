package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `json:"id" orm:"column(id);pk"`
	Name  string `json:"name" orm:"size(50)"`
	Age   int    `json:"age"`
	Sex   int    `json:"sex"`
	Phone string `json:"phone"`
}

func AddUser(u *User) (err error) {
	o := orm.NewOrm()

	o.Begin()
	defer func() {
		if err != nil {
			_ = o.Rollback()
			return
		}
		_ = o.Commit()
	}()

	count, err := o.QueryTable(new(User)).Filter("phone__in", u.Phone).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("err")
	}
	_, err = o.Insert(u)
	if err != nil {
		fmt.Printf("sss%v\n", err)
		return err
	}
	return nil
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

func ModifyUserById(uuid string, user User) (err error) {
	o := orm.NewOrm()
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	user.Id = uuid
	queryUser := User{
		Id: uuid,
	}
	if err = o.Read(&queryUser); err != nil {
		return err
	}
	if _, err = o.Update(&user); err != nil {
		return err
	}
	return nil
}
