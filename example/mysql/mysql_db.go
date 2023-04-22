package mysql

import (
	"database/sql"
)

var DB *sql.DB

func init() {
	//var mysqlErr error
	//DB, mysqlErr = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hero_story")
	//
	//if nil != mysqlErr {
	//	panic(mysqlErr)
	//}
	//
	//DB.SetMaxOpenConns(128)
	//DB.SetMaxIdleConns(16)
	//DB.SetConnMaxLifetime(2 * time.Minute)
	//
	//if mysqlErr = DB.Ping(); nil != mysqlErr {
	//	panic(mysqlErr)
	//}
}
