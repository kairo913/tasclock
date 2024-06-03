package utility

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var Layout = "2006-01-02 15:04:05"

type TrashScanner struct {}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func Init() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")

	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	var err error
	if Db, err = sql.Open("mysql", path); err != nil {
		log.Fatal("Db open error: ", err.Error())
	}
	checkConnect(100)

	fmt.Println("db connected!")
}

func checkConnect(count uint) {
	if err := Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		if count == 0 {
			log.Fatal("Db connection failed!")
		}
		fmt.Printf("Retry connect to db, %d times left\n", count)
		checkConnect(count)
	}
}

