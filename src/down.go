package main

import (
	"bufio"
	"fmt"
	"img-down/src/libs"
	"io"
	"os"
	"path/filepath"
)

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

	switch commandstr {
	case "-l":
		list()
		break

	case "-p":
		page()
		break

	case "-i":
		img()
		break

	case "-ic":
		imgConcur()
		break

	case "-h":
		help()
		break

	default:
		fmt.Println("对不起，参数错误！")
		help()
		break

	}

	return

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

func list()  {
	//获取用户输入的
	fmt.Println("请输入根路径:")
	firstInput := bufio.NewScanner(os.Stdin)

	firstPath := ""
	if firstInput.Scan() {
		firstPath = firstInput.Text()
	}
	fmt.Println(firstPath)

	fmt.Println("请输入文件列表和'end'并按回车结束输入：")
	var pathstr []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}

		textstr := input.Text()
		if len(textstr) > 5 {
			pathstr = append(pathstr, textstr)
		}
	}

	libs.Lists(pathstr)
	return

}

func page()  {
	fmt.Println("获取页面地址")
	libs.Pages()
}

func img()  {
	fmt.Println("逐条下载图片")
	libs.Images()
}

func imgConcur() {
	fmt.Println("并发下载图片")
	libs.ImagesConcur()
}

func help()  {
	fmt.Println("请参照下面说明:")
	fmt.Println(" -l：获取网页列表")
	fmt.Println(" -p：获取列表下页面")
	fmt.Println(" -i：获取页面下图片地址并下载")
	fmt.Println(" -ic：获取页面下图片地址并发下载")
}

func getConfig()  {
	configPath := "config/down.ini"
	dir, _ := os.Getwd()
	allConfigPath := dir+"/"+configPath
fmt.Println(allConfigPath)
	inputFile, err := os.Open(allConfigPath)
	if err != nil {
		fmt.Printf("请确认配置文件是否正确？")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if err == io.EOF {
			return
		}
	}
}
