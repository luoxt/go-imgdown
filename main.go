package main

import (
	"bufio"
	"fmt"
	"os"
	"go-imgs-down/app/cmd"
)


func main() {
	//获取配置

	//初始化数据库

	//初始化缓存

	//执行任务
	command()

	//img.Images()
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




