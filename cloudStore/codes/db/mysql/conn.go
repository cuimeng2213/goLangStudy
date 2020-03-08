package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:cm2213@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("fail to connect mysql: %v \n", err)
		os.Exit(1)
		return
	}
}

//DBConn: 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
