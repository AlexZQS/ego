package common

import (
	"database/sql"
	"fmt"
)

//数据库操作的三个对象
var (
	db   *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
)

//打开数据库连接
func openConn() (err error) {
	db, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/ego")
	if err != nil {
		fmt.Println("打开连接错误", err)
	}
}
