package models

import (
	"github.com/beego/beego/v2/client/orm"
)

const (
	UserDisable = iota
	UserEnable
)

func InsertOneUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindUserByName(username string) (*User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
