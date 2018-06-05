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
  // TODO: to learn basic mysql: https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
  
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
