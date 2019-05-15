package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	return nil
}

//关闭连接,首字母大写，需要跨包访问的
func CloseConn() {
	if rows != nil {
		rows.Close()
	}

	if stmt != nil {
		stmt.Close()
	}

	if db != nil {
		db.Close()
	}
}

//执行DML 新增，删除，修改操作
func Dml(sql string, args ...interface{}) (int64, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DML时出现错误, 打开连接失败")
		return 0, err
	}

	//此处也是等号
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出现错误, 预处理出现错误")
		return 0, err
	}

	//此处要有... 表示切片，如果没有表示数组，会报错
	result, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println("执行DML时出现错误, 执行错误")
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("执行DML时出现错误, 获取受影响行数错误")
		return 0, err
	}

	//关闭连接
	CloseConn()
	return count, err
}

//执行DQL查询
func Dql(sql string, args ...interface{}) (*sql.Rows, error) {

	err := openConn()
	if err != nil {
		fmt.Println("执行DML时出现错误, 打开连接失败")
		return nil, err
	}

	//此处也是等号
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出现错误, 预处理出现错误")
		return nil, err
	}

	//此处要有... 表示切片，如果没有表示数组，会报错
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("执行DML时出现错误, 执行错误")
		return nil, err
	}

	//此处没有关闭,调用此函数结束时，记得关闭连接
	return rows, nil
}
