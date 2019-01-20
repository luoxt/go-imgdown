package helper

import (
	"os"
	"log"
	"io/ioutil"
	"runtime"
)

var ostype = runtime.GOOS

//获取配置文件
func GetConfig() string  {
	configFile := "config.ini"

	dir, _ := os.Getwd()
	path := "/"
	if ostype == "windows" {
		path = "\\"
	}

	allConfigPath := dir + path + "conf" + path + configFile

	if _, err := os.Stat(allConfigPath);  err != nil {
		allConfigPath = dir + path + "conf" + path + configFile
	}

	configbye, err := ioutil.ReadFile(allConfigPath)
	if err != nil {
		log.Panic(err)
	}

	return string(configbye)
	
}