package model

import "time"

type Client struct {
	Id       string `xorm:"pk"` //clientid
	Username string
	Password string

	ProductId string

	Disabled bool
	Created  time.Time `xorm:"created"`
}

type ClientHistory struct {
	ClientId string    `xorm:"index"`
	Event    string    //上线，下线
	Extra    string    //ip 。。。
	Created  time.Time `xorm:"created"`
}
