package db

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"go-imgs-down/app/base/helper"
)

func fetchAll()  {

	log.Println()
}

func fetch(table string, url string)  {
	rows := db.QueryRow("select id from list_url where page_url = ?", url)
	var id int
	rows.Scan(&id)
}

func insert(url string, pid int64) (id int64)  {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO list_url(page_url, pid) values(?, ?)")
	helper.CheckErr(err)

	res, err := stmt.Exec(url, pid)
	helper.CheckErr(err)

	insert_id, err := res.LastInsertId()
	helper.CheckErr(err)
	id = insert_id
	return insert_id
}

func update(id int64)  {
	stmt, err := db.Prepare("UPDATE list_url set status = 1 where id = ?")
	helper.CheckErr(err)

	stmt.Exec(id)
}

