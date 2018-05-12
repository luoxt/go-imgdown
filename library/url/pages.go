package url

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"golang.org/x/net/html"
	"go-imgdown/library/base"
)

var urlArr []string

func Pages() {
	//是否存在
	rows, err:= db.Query("select id, page_url from list_url where status = 0")
	if err!=nil {
		fmt.Println("【没有地址可以下载。。。】")
		os.Exit(0)
	}
	for rows.Next(){
		var page_url string
		var id int
		rows.Scan(&id, &page_url)

		pageLink(page_url, id)
	}
	defer db.Close()

}

//页面地址
func pageLink(urlstr string, id int) {

	//html document
	http, _ := http.Get(urlstr)
	doc, err := html.Parse(http.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	defer http.Body.Close()

	//获取链接地址
	hrefUrl := nodeUrl(nil, doc)

	for _, urlVal := range hrefUrl {

		if path.Ext(urlVal) == ".html" {

			haveJs := strings.Index(urlVal, "/html/article/")

			if haveJs > 0 {

				if !inArrays(urlArr, urlVal) {
					urlArr = append(urlArr, urlVal)
					insert(urlVal, "html")
				}
			}
		}
	}

	
	stmt, err := db.Prepare("UPDATE list_url set status = 1 where id = ?")
	base.CheckErr(err)

	stmt.Exec(id)
	db.Close()
	fmt.Println("【list_id】:", id, "【list_url】", urlstr)

}

//获取页面链接
func nodeUrl(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				haveJs := strings.Index(a.Val, "javascript")
				if haveJs != 0 {

					url := a.Val
					if strings.Index(url, "http") < 0 {
						url = "http://www.22qqjj.com" + url
					}

					haveJs := strings.Index(url, "/html/article/")
					if haveJs > 0 {
						links = append(links, url)
					}

				}

			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = nodeUrl(links, c)
	}

	return links
}

func inArrays(urlArr []string, url string) bool {
	for _, val := range urlArr {
		if val == url {
			return true
		}
	}
	return false
}

func insert(url string, url_type string) {
	//是否存在
	rows := db.QueryRow("select id from page_url where page_url = ?", url)
	var id int
	rows.Scan(&id)

	if id<=0 {
		//插入数据
		stmt, err := db.Prepare("INSERT INTO page_url(page_url, pid) values(?, ?)")
		base.CheckErr(err)

		res, err := stmt.Exec(url, 18)
		base.CheckErr(err)

		insert_id, err := res.LastInsertId()
		base.CheckErr(err)

		fmt.Println("【成功插入】ID：", insert_id)


	} else {
		fmt.Println("【已经存在】", id)
	}

	db.Close()

}
