package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 这个写法是全局变量， 包的任何地方都可以使用到
var (
	dbConn *sql.DB
	err    error
)

/**
init  是一个特殊方式， 当整个包被加载的时候， 第一个执行
*/
func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(192.168.205.10:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
