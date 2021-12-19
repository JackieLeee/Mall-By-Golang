package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     int       `json:"status"`
	Balance    float64   `json:"balance"`
}

func init() {
	orm.RegisterModel(
		new(User),
	)
}
