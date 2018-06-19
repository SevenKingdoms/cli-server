package db

import (
	"database/sql"
	"fmt"
	//"time"

	"github.com/cli-server/conf"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	db, err := sql.Open("mysql",
		conf.USER+":"+conf.PASSWORD+"@tcp("+conf.HOST+":"+conf.PORT+")/"+conf.DB)
  fmt.Println(db)
  // A recommended way to check err with more situations considerated.
  // err, _ = db.Exec("DO 1")
  checkErr(err)
  defer db.Close()

	// Debug: 插入数据
	// stmt, err := db.Prepare("INSERT User SET openId=?,name=?,avatar=?,phone=?")
	// checkErr(err)
	// res, err := stmt.Exec("a12", "astaxie", "http://a.png", "13712321234")
	// checkErr(err)
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)
  
  // TODO: to learn basic mysql: https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
