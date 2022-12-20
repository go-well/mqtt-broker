package db

import (
	"github.com/go-well/mqtt-broker/model"
	"xorm.io/xorm"
	"xorm.io/xorm/log"

	//加载数据库驱动
	//_ "github.com/mattn/go-sqlite3" //CGO版本
	_ "github.com/glebarez/go-sqlite" //纯Go版本 使用ccgo翻译的，偶有问题
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var Engine *xorm.Engine

type Options struct {
	Type     string `yaml:"type" json:"type"`
	URL      string `yaml:"url" json:"url"`
	Debug    bool   `yaml:"debug" json:"debug"`
	LogLevel int    `json:"log_level" json:"log_level"`
	Sync     bool   `yaml:"sync" json:"sync"`
}

func Open(cfg Options) error {
	var err error
	Engine, err = xorm.NewEngine(cfg.Type, cfg.URL)
	if err != nil {
		return err
	}
	if cfg.Debug {
		Engine.ShowSQL(true)
	}

	Engine.SetLogLevel(log.LogLevel(cfg.LogLevel))
	//Engine.SetLogger(logrus.StandardLogger())

	//同步表
	if cfg.Sync {
		err = Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return Engine.Close()
}

func Sync() error {
	return Engine.Sync2(
		new(model.User),
	)
}
