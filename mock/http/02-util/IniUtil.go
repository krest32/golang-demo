package main

import (
	"fmt"
	"strings"
	myIo "demo-basic/mock/util/io"
	myHttp "demo-basic/mock/util/http"
	myIni "demo-basic/mock/util/ini"
)


type Request struct {
	method  	string
	url 		string
	requestBody string
	contentType string
}

func main()  {
	var responseTxt string
	filePath := "mock/http/02-util/requestBody.txt"
	iniFile := "mock/http/02-util/config.ini"
	writeFilePath := "src/day6-lib/Ini/02-util/response.txt"
	cfg := myIni.Parse(iniFile)
	res := Request{
		method: cfg.Section("request").Key("method").String(),
		url: cfg.Section("request").Key("url").String(),
		requestBody: myIo.ReadFile(filePath),
		contentType: cfg.Section("request").Key("contentType").String(),
	}

	if strings.EqualFold(res.method,"post") {
		responseTxt = myHttp.Post(res.url,res.contentType,responseTxt)
	}else if strings.EqualFold(res.method,"get") {
		fmt.Printf("get")
	}
	myIo.WriteFile(responseTxt,writeFilePath)
}

