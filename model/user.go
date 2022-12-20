package model

import "time"

type User struct {
	Id int64

	Username string
	Password string

	Disabled bool
	Created  time.Time `xorm:"created"`
}
