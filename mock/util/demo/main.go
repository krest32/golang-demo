package main

import "log"

func main()  {
	log.Printf("this is normal log" )

	/*
		ini 測試部分
		folder :="mock/recourses/"
		iniFileName := "kyConfig.ini"
		iniFilePath := folder+iniFileName
		value:= myini.GetValue(iniFilePath, "request", "method")
		fmt.Printf(value.String())

	 */




	/*
		Io 測試部分
	folder :="mock/recourses/"
	iniFileName := "kyConfig.ini"
	oldFileName :="demo-io.txt"
	newFileName := "demo-new.txt"
	oldFilePath := folder+oldFileName
	newFilePath := folder+newFileName

	myio.WriteFile("hello world 2\n",oldFilePath)
	myio.Append("haha hello\n", oldFilePath)
	myio.CopyFilePath(oldFilePath, newFilePath)

	*/

}
