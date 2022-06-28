package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloController(r *gin.Engine) {
	// 读取文件中的字符
	r.GET("/hello", func(c *gin.Context) {
		// 直接返回string信息
		c.String(http.StatusOK, "hello world")
	})
}
