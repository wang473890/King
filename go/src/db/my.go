package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const CConnectString = "wg:wg2019@tcp(140.143.234.207:3600)/king?charset=latin1"

var My *sql.DB

func InitMy() {
	var err error
	My, err = sql.Open("mysql", CConnectString)
	if err != nil {
		//TODO 日志记录
		fmt.Println("mysql err =====", err)
	}
	My.SetMaxOpenConns(200)
	My.SetMaxIdleConns(100)
}
