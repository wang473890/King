package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const CConnectString = "root:wang473890@tcp(localhost:3306)/wang?charset=utf8"

type Context struct {
	Db *sql.DB
}

var DevContext Context

func InitDb() {
	var err error
	DevContext.Db, err = sql.Open("mysql", CConnectString)
	if err != nil {
		//TODO 日志记录
		fmt.Println("mysql err =====", err)
	}
	DevContext.Db.SetMaxOpenConns(200)
	DevContext.Db.SetMaxIdleConns(100)
}