package model

import "time"

type Subscribe struct {
	Id int64

	ClientId string
	Topic    string
	Disabled bool
	Created  time.Time `xorm:"created"`
}

type Publish struct {
	Id       int64
	ClientId int64

	Topic    string
	Disabled bool
	Created  time.Time `xorm:"created"`
}
