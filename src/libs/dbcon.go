package libs

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const(
	HOST = "127.0.0.1"
	DBNAME = "link"
	USER = "root"
	PASS = "123456"
	PORT = "3306"
)

func Dbcon() *sql.DB {
	db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+HOST+":"+PORT+")/"+DBNAME+"?charset=utf8")
	//defer db.Close()

	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	db.SetMaxOpenConns(120)
	db.SetMaxIdleConns(120)
	return db
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
