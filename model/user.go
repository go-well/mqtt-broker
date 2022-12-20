package model

import "time"

type User struct {
	Id int64

	Created time.Time `xorm:"created"`
}
