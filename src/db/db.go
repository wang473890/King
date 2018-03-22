package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"config"
)

type MysqlConf struct {
	Db     string
	Coding string
}

var Conf = MysqlConf{
	Db:     "king",
	Coding: "utf8",
}

type Context struct {
	Db *sql.DB
}

var DevContext Context

func SetDb(conf MysqlConf) {
	Conf = conf
}
func InitDb() {
	var err error
	var connect string
	connect = config.DevCtx.MysqlAdmin + ":" + config.DevCtx.MysqlPass + "@tcp(" + config.DevCtx.MysqlIp + ")/" + Conf.Db + "?charset=" + Conf.Coding
	DevContext.Db, err = sql.Open("mysql", connect)
	if err != nil {
		//TODO 日志记录
		fmt.Println("mysql err =====", err)
	}
	DevContext.Db.SetMaxOpenConns(200)
	DevContext.Db.SetMaxIdleConns(100)
}
