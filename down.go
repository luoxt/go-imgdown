package main

import (
	"bufio"
	"fmt"
	"os"
	"go-imgdown/library/cmd"
	// "go-imgdown/library/img"
)


func main() {
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




