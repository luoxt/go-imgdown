package libs

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"io"
	"path"
	"strings"
	"golang.org/x/net/html"
)

const (
	DOWN_PATH  = "./data/downloads/"
	GOCOUNT = 10
)

type ImgLink struct {
	num int
	url string
	domain string
}

type ROW struct {
	Id int
	Url string
}

var IL ImgLink
var db *sql.DB

func init()  {
	db = Dbcon()
}

func ImagesConcur() {

	//获取数据库数据
	rows, err:= db.Query("select id, page_url from page_url where pid=18 and status = 0 ")
	if err != nil {
		fmt.Println("查询出错", err)
		os.Exit(1)
	}
	for rows.Next(){
		var id int
		var page_url string
		rows.Scan(&id, &page_url)

		GetImgConcurConcur(page_url, id)

		stmt, _ := db.Prepare("UPDATE page_url set status = 1 where id = ?")
		stmt.Exec(id)
		stmt.Close()
	}
	rows.Close()

	i := 0
	for {
		i++
	}

}

//获取地址
//并发下载
func GetImgConcurConcur(urlstr string, id int) {

	http, _ := http.Get(urlstr)
	doc, err := html.Parse(http.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findImgLinks: %v\n", err)
		os.Exit(1)
	}

	//获取图片地址
	imgfUrl := ImgConcur(nil, doc)
	imgcount := len(imgfUrl)
	fmt.Println("【需下载】张数：", imgcount)

	img_list := make(chan ROW, GOCOUNT)
	st_count := make(chan int, GOCOUNT)

	go func() {
		for _, urlVal := range imgfUrl {

			var ro = ROW{id, urlVal}
			img_list <- ro
			st_count <- 1
		}
	}()

	//并发下载
	for i:= 1; i<=GOCOUNT; i++ {
		go func(num int) {
			fmt.Println("协程序号：", num)
			for {
				img_row := <- img_list

				down_status := DownImgConcur(img_row.Url, img_row.Id)
				if down_status {
					<- st_count
				} else {
					img_list <- img_row
				}
			}

		}(i)
	}

}

//获取节点图片
func ImgConcur(ImgLinks []string, n *html.Node) []string {

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
						ImgLinks = append(ImgLinks, url)
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ImgLinks = ImgConcur(ImgLinks, c)
	}

	return ImgLinks
}

/**
 *	获取URL
 */
func DownImgConcur(url string, id int) bool {
	L.num++

	table_id := ImgInsert(url, id)
	if table_id==0 {
		return false
	}

	//保存路径
	paths :=  DOWN_PATH
	baseName := path.Base(url)
	targName := paths + baseName

	http, err := http.Get(url)
	if err != nil {
		fmt.Println("获取url", err)
		return false
	}
	downSize := http.ContentLength
	defer http.Body.Close()

	//检查文件存在
	fileinfo, err := os.Stat(targName)
	if downSize > 7000 {
		if os.IsNotExist(err) || fileinfo.Size() != downSize {

			fmt.Println("【图片下载】", "地址：", url)

			imgFile, err := os.Create(targName)
			if err != nil {
				fmt.Println("创建文件失败", err)
				return false
			}

			_, err = io.Copy(imgFile, http.Body)
			if err == nil {
				fileinfo, err := os.Stat(targName)
				if err != nil {
					return false
				}
				filesize := fileinfo.Size()

				//判断是否下载完成
				if filesize == downSize {
					//更新数据
					stmt, err := db.Prepare("UPDATE img_url set size=?, status = 1 where id = ?")
					if err == nil {
						stmt.Exec(downSize, table_id)
					} else {
						fmt.Println(err.Error())
					}
					stmt.Close()
					return true
				} else {

					return false
				}

			} else {
				return false
			}

		} else {
			fmt.Println("【图片存在】", "地址：", url)
			return true
		}
		return false
	}
	return true
}

func ImgInsert(url string, pid int)(table_id int64) {

	//是否存在
	rows := db.QueryRow("select id, status from img_url where page_url = ?", url)
	var id int64
	var status int
	rows.Scan(&id, &status)

	if id<=0 {
		//插入数据
		stmt, err := db.Prepare("ImgInsert INTO img_url(page_url,pid) values(?,?)")
		CheckErr(err)

		res, err := stmt.Exec(url, pid)
		CheckErr(err)
		stmt.Close()

		ImgInsert_id, err := res.LastInsertId()
		CheckErr(err)

		fmt.Println("【记录成功插入】ID：", ImgInsert_id)

		table_id = ImgInsert_id

	} else {

		if status == 1 {
			fmt.Println("【记录已经存在】", id)
			table_id = 0
		} else {
			table_id = id
		}
	}

	return table_id
}

