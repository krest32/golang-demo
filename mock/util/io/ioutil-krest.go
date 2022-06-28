package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
	读取文件内容，并且转换为string
*/
func ReadFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return ""
	}
	return string(content)
}


/**
	将string内容写入文件，会覆盖原有文件内容
	可以新建文件，但是不能够创建文件夹
*/
func WriteFile(content, filePath string)  {
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}

/**
	在文件中追加内容新内容
 */
func Append(content, filePath string) bool{
	file, errOpen := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if errOpen != nil {
		fmt.Printf("open file err=%v\n",errOpen)
		return false
	}
	defer file.Close()
	writer:= bufio.NewWriter(file)
	writer.WriteString(content)
	errWrite := writer.Flush()
	if errWrite != nil {
		fmt.Printf("writer file err=%v\n",errWrite)
		return false
	}
	return true
}

func CopyFilePath(oldFile, newFile string) bool  {

	srcFile, err := os.Open(oldFile)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(newFile, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return false
	}
	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	_, err = io.Copy(writer, reader)
	if err == nil {
		fmt.Printf("copy complete \n")
		return true
	}else {
		fmt.Printf("copy fail\n")
		return false
	}
}

func deleteFile(filePath string)  {
}