package url

import (
	"fmt"
	"database/sql"
	mysql "go-imgs-down/app/db"
	"go-imgs-down/app/base/helper"
)

var db *sql.DB

func init()  {
	db = mysql.Connect()
}

func Lists(url_list []string) {
	
	for _, list := range url_list {

		func(url string){
			//是否存在
			rows := db.QueryRow("select id from list_url where page_url = ?", url)
			var id int
			rows.Scan(&id)

			if id<=0 {
				//插入数据
				stmt, err := db.Prepare("INSERT INTO list_url(page_url, pid) values(?, ?)")
				helper.CheckErr(err)

				res, err := stmt.Exec(url, 18)
				helper.CheckErr(err)

				insert_id, err := res.LastInsertId()
				helper.CheckErr(err)

				fmt.Println("【成功插入】ID：", insert_id)


			} else {
				fmt.Println("【已经存在】", id)
			}

			defer db.Close()
		}(list)
	}

	fmt.Println("【列表处理完成！】")
	return

}


