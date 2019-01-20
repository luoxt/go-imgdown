package db

import (
	"database/sql"
	"log"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"go-imgs-down/app/base/helper"
)

var db *sql.DB

func init()  {
	db = Connect()
}

func Connect() *sql.DB {
	type database struct {
		Host string `json:"host"`
		Dbname string `json:"dbname"`
		User string `json:"user"`
		Pass string `json:"pass"`
		Port string `json:"port"`
	}
	var dbconfig database
	str := helper.GetConfig()

	errs := json.Unmarshal([]byte(str), &dbconfig)
	if errs != nil {
		log.Fatal("JSON解析错误")
	}

	db, err := sql.Open("mysql", dbconfig.User+":"+dbconfig.Pass+"@tcp("+dbconfig.Host+":"+dbconfig.Port+")/"+dbconfig.Dbname+"?charset=utf8")

	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	db.SetMaxOpenConns(120)
	db.SetMaxIdleConns(120)
	return db
}

