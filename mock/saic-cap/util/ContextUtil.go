package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DealContext(context *gin.Context){
	appId := context.GetHeader("app-id");
	timeStamp := context.GetHeader("timestamp");
	requestSn := context.GetHeader("request-sn");
	signature := context.GetHeader("signature");
	signType := context.GetHeader("sign-type");
	// 显示请求头的内容
	fmt.Print(appId+"\n")
	fmt.Print(timeStamp+"\n")
	fmt.Print(requestSn+"\n")
	fmt.Print(signature+"\n")
	fmt.Print(signType+"\n")
}