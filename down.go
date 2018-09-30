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



<<<<<<< HEAD
//获取配置文件
func getConfig()  {
	dir, _ := os.Getwd()

	//配置路径
	configFile := "config.ini"
	path := "/"
	if ostype == "windows"{
		path = "\\"
	}
	//configPath := dir + path + "conf" + path + configFile
	configPath := dir + path + "src" + path + "go-imgdown" +  path + "conf" + path + configFile

	//配置内容
	configbye, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Print(err)
	}
	configStr := string(configbye)

	//////////////////////////////////////////
	type database struct {
		Host string `json:"host"`
		Dbname string `json:"dbname"`
		User string `json:"user"`
		Pass string `json:"pass"`
		Port string `json:"port"`
	}
	
	type Config struct {
		Database *database `json:"database"`
		Path  string  `json:"path"`
		Gocount int  `json:"gocount"`
	}

	///////////////////////////////
	var goconfig Config

	fmt.Println(configStr)
	if err := json.Unmarshal([]byte(configStr), &goconfig); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(goconfig)
	}

	////////////////////////////////////
	//type ConfigStruct struct {
	//	Host              string   `json:"host"`
	//	Port              int      `json:"port"`
	//	AnalyticsFile     string   `json:"analytics_file"`
	//	StaticFileVersion int      `json:"static_file_version"`
	//	StaticDir         string   `json:"static_dir"`
	//	TemplatesDir      string   `json:"templates_dir"`
	//	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	//	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	//	Fruits            []string `json:"fruits"`
	//}
	//
	//jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	//
	////json str 转struct
	//var config ConfigStruct
	//if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
	//	fmt.Println("================json str 转struct==")
	//	fmt.Println(config)
	//	fmt.Println(config.Host)
	//}

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
=======

>>>>>>> 2913d28f904ddeeee7dfaffb0bf4a96173829819
