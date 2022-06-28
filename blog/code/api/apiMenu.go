package api

import (
	"demo-basic/blog/code/config"
	"github.com/gin-gonic/gin"
)

func Start(){
	r := gin.Default()

	// 测试路径
	HelloController(r)

	// 用户接口
	AppsController(r)

	r.Run(":"+config.GetBootPort())
}




