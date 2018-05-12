package main

import (
	"bufio"
	"fmt"
	"os"
	"go-imgdown/library/cmd"
	"io/ioutil"
	"runtime"
)

var ostype = runtime.GOOS

func init()  {

	//加载配置文件
	getConfig()

	fmt.Println("dddd")

}

func main() {

	//获取用户输入的命令
	commandstr := ""

	args := os.Args
	if args == nil || len(args)<2 || len(args[1])<2 {
		commandstr = command()
	} else {
		commandstr = args[1]
	}

	cmd.Command(commandstr)

}

//获取命令参数
func command() string  {
	//获取用户输入的
	fmt.Println("请输入命令:")
	firstInput := bufio.NewScanner(os.Stdin)

	firstPath := ""
	if firstInput.Scan() {
		firstPath = firstInput.Text()
	}
	if len(firstPath) <2 {
		return command()
	} else {
		return firstPath
	}
}

//获取配置文件
func getConfig()  {
	configFile := "config.ini"

	dir, _ := os.Getwd()
	path := "/"
	if ostype == "windows"{
		path = "\\"
	}

	allConfigPath := dir + path + "src" + path + "go-imgdown" + path + "conf" + path + configFile
	fmt.Println(allConfigPath)

	configbye, err := ioutil.ReadFile(allConfigPath)
	if err != nil {
		fmt.Print(err)
	}

	str := string(configbye)
	fmt.Println(str)
	//
	//inputFile, err := os.Open(allConfigPath)
	//if err != nil {
	//	fmt.Printf("请确认配置文件是否正确？")
	//	return
	//}
	//defer inputFile.Close()
	//
	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, err := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if err == io.EOF {
	//		return
	//	}
	//}
}
