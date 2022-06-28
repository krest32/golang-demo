package main

import (
	"demo-basic/blog/code/api"
	"demo-basic/blog/code/config"
	"fmt"
	"os"
)

type Alien struct {
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Height string `json:"height"`
}

func main() {
	// 通过命令行启动服务器，调用Go服务进行启动
	var iniFilePath string
	//os.Args是一个[]string，获取命令的输入
	if len(os.Args) > 1 {

		// 获取第一个请求参数
		iniFilePath = os.Args[1]
		config.ParseConfig(iniFilePath)
	} else {
		fmt.Printf("error config filePath")
	}

	// 启动服务
	api.Start()
}
