package img

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"io"
	"path"
	"strings"
	"golang.org/x/net/html"
	mysql "go-imgdown/library/db"
	"go-imgdown/library/base"
)

type Link struct {
	num int
	url string
	domain string
}

var L Link

var db *sql.DB

func init()  {
	db = mysql.Connect()
}

func Images() {

	//获取数据库数据
	var count int
	resCount:= db.QueryRow("select count(*) as count from page_url where status = 0")
	resCount.Scan(&count)
	if count==0 {
		fmt.Println("没有数据")
		return
	}

	rows, _:= db.Query("select id, page_url from page_url where status = 0")
	for rows.Next(){
		fmt.Println("结构")
		var id int
		var page_url string
		rows.Scan(&id, &page_url)

		urlVal := page_url
		paths := strings.Split(urlVal, "/")

		L.num = 0
		L.domain = paths[2]
		L.url = urlVal
fmt.Println(L)
		// GetImgUrl(urlVal, id)

		// //更新数据
		// stmt, _ := db.Prepare("UPDATE page_url set status = 1 where id = ?")
		// stmt.Exec(id)
		// defer db.Close()
	}

	i := 0
	for {
		i=i+1
	}

}

func GetImgUrl(urlstr string, id int) {

	http, _ := http.Get(urlstr)
	doc, err := html.Parse(http.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	//获取图片地址
	var imgArr []string
	imgfUrl := ImgUrl(nil, doc)
	imgcount := len(imgfUrl)
	fmt.Println("【需下载】张数：", imgcount)

	for _, urlVal := range imgfUrl {
		if path.Ext(urlVal) == ".jpg" {
			if !Inarrays(imgArr, urlVal) {
				imgArr = append(imgArr, urlVal)

				DownImg(urlVal, id)

				if L.num >= imgcount {
					fmt.Println("【已经下载完成】张数：", L.num)
				}
			}
		}
	}
}

//获取节点图片
func ImgUrl(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				haveJs := strings.Index(a.Val, "javascript")
				if haveJs != 0 {
					url := a.Val
					if strings.Index(url, "http") < 0 {
						url = "http://"+L.domain + url
					}
					ext := path.Ext(url)
					if ext == ".jpg" ||ext== ".jpeg" {
						links = append(links, url)
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = ImgUrl(links, c)
	}

	return links
}

/**
 *	获取URL
 */
func DownImg(url string, id int) {
	L.num++

	table_id := Insert(url, id)
	if table_id==0 {
		return
	}

	//保存路径
	var paths string = `/Users/luoxiantao/Downloads/pic/pic1/`
	baseName := path.Base(url)
	targName := paths + baseName

	//检查文件存在
	_, err := os.Stat(targName)
	if os.IsNotExist(err) {
		http, err := http.Get(url)
		if err != nil {
			fmt.Println("获取url", err)
			return
		}

		fileSize := http.ContentLength
		if fileSize > 7000 {

			fmt.Println("【图片下载】", "地址：", url)

			imgFile, _ := os.Create(targName)
			io.Copy(imgFile, http.Body)

			//更新数据
			
			stmt, _ := db.Prepare("UPDATE img_url set status = 1 where id = ?")
			stmt.Exec(table_id)
			defer db.Close()
		}
		defer http.Body.Close()
	} else {
		fmt.Println("【图片存在】",  "地址：", url)
	}
}

func Insert(url string, pid int)(table_id int64) {
	

	//是否存在
	rows := db.QueryRow("select id from img_url where page_url = ?", url)
	var id int64
	rows.Scan(&id)

	if id<=0 {
		//插入数据
		stmt, err := db.Prepare("Insert INTO img_url(page_url,pid) values(?,?)")
		base.CheckErr(err)

		res, err := stmt.Exec(url, pid)
		base.CheckErr(err)

		Insert_id, err := res.LastInsertId()
		base.CheckErr(err)

		fmt.Println("【记录成功插入】ID：", Insert_id)

		table_id = Insert_id

	} else {
		fmt.Println("【记录已经存在】", id)

		table_id = 0
	}

	defer db.Close()
	return table_id
}

func Inarrays(urlArr []string, url string) bool {

	for _, val := range urlArr {
		if val == url {
			return true
		}
	}
	return false

}
