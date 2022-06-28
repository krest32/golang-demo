package util

import (
	"fmt"
	"io/ioutil"
)

func GetFileString(filePath string)(result string){
	content,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file failed , err : ", err)
		return
	}
	result = string(content)
	return result
}



