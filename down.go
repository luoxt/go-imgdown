package main

import (
	"bufio"
	"fmt"
	"os"
	"go-imgdown/library/cmd"
	"io/ioutil"
	"runtime"
	"encoding/json"
)

var ostype = runtime.GOOS

func init()  {

	//加载配置文件
	getConfig()

}

func main() {

	command()

	

}

//获取命令参数
func command()  {
	//获取用户输入的命令
	commandstr := ""

	args := os.Args
	if args == nil || len(args)<2 || len(args[1])<2 {
		//获取用户输入的
		fmt.Println("请输入命令:")
		firstInput := bufio.NewScanner(os.Stdin)

		if firstInput.Scan() {
			commandstr = firstInput.Text()
		}
	} else {
		commandstr = args[1]
	}

	cmd.Command(commandstr)
}



//获取配置文件
func getConfig()  {
	configFile := "config.ini"

	dir, _ := os.Getwd()
	path := "/"
	if ostype == "windows"{
		path = "\\"
	}

	allConfigPath := dir + path + "conf" + path + configFile
	fmt.Println(allConfigPath)

	configbye, err := ioutil.ReadFile(allConfigPath)
	if err != nil {
		fmt.Print(err)
	}

	str := string(configbye)
	fmt.Println(str)

	//////////////////////////////////////////
	type database struct {
		host string
		dbname string
		user string
		pass string
		port string
	}
	
	type Config struct {
		datebases database 
		path  string
		gocount int
	}

	///////////////////////////////
	type Book struct { 
		Title string 
		Authors []string 
		Publisher string 
		IsPublished bool 
		Price float32
	} 
	
	b := []byte(`{"Title": "Go语言编程", "Sales": 1000000}`) 
	var gobook Book

	err = json.Unmarshal(b, &gobook)
    fmt.Println(gobook)

	//
	//inputFile, err := os.Open(allConfigPath)
	//if err != nil {
	//	fmt.Printf("请确认配置文件是否正确？")
	//	return
	//}

	//defer inputFile.Close()

	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, err := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if err == io.EOF {
	//		return
	//	}
	//}
}
