package cmd

import(
	"bufio"
	"fmt"
	"os"
	"go-imgdown/library/url"
	"go-imgdown/library/img"
)

func Command(comstr string)  {
	switch comstr {
	case "-l":
		list()
		break

	case "-p":
		page()
		break

	case "-i":
		imgs()
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
}

func list()  {
	//获取用户输入的
	// fmt.Println("请输入根路径:")
	// firstInput := bufio.NewScanner(os.Stdin)

	// var firstPath string
	// if firstInput.Scan() {
	// 	firstPath = firstInput.Text()
	// }

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

	url.Lists(pathstr)
	return

}

func page()  {
	fmt.Println("获取页面地址")
	url.Pages()
}

func imgs()  {
	fmt.Println("逐条下载图片")
	img.Images()
}

func imgConcur() {
	fmt.Println("并发下载图片")
	img.ImagesConcur()
}

func help()  {
	fmt.Println("请参照下面说明:")
	fmt.Println(" -l：获取网页列表")
	fmt.Println(" -p：获取列表下页面")
	fmt.Println(" -i：获取页面下图片地址并下载")
	fmt.Println(" -ic：获取页面下图片地址并发下载")
}
