package helper

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Gtime() int  {
	return time.Now().Second()
}

func Gdate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//创建目录
func Mkdir(path string) {

	codePath := path
	_, err := os.Stat(codePath)
	if os.IsNotExist(err) {
		os.Mkdir(codePath, os.ModePerm)
	}
}

func Gtype(a interface{}) interface{} {
	return reflect.TypeOf(a)
}

//组装JSON
type Rejson struct {
	Status string `json:"status"`
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{}  `json:"data"`
}
func Gjson(status string, code string, msg string, data interface{}) *Rejson {

	re := &Rejson{
		status,
		code,
		msg,
		data,
	}
	//str, _ := json.Marshal(re)
	return re
}
