package main

import (
	"demo-basic/mock/saic-cap/util"
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"
)

func main() {

	r := gin.Default()
	// 读取文件中的字符
	r.POST("/EE/saicComp", func(c *gin.Context) {
		// 获取请求头内容
		util.DealContext(c)
		// 等待30秒
		//time.Sleep(10 * time.Second)
		// 直接返回string信息
		c.String(http.StatusOK, util.GetFileString("F:\\project\\go\\GoPath\\src\\demo-basic\\mock\\saic-cap\\rescourses\\demoCompStr1.json"))
	})

	r.POST("/EE/saicIndi", func(c *gin.Context) {
		// 获取请求头内容
		util.DealContext(c)
		// 等待30秒
		//time.Sleep(10 * time.Second)

		// 直接返回string信息
		c.String(http.StatusOK, util.GetFileString("F:\\project\\go\\GoPath\\src\\demo-basic\\mock\\saic-cap\\rescourses\\demoIndiStr1.json"))
	})

	r.POST("/EE/CAP/Token", func(c *gin.Context) {
		// 获取请求头内容
		util.DealContext(c)
		// 直接返回string信息
		c.String(http.StatusOK, util.GetFileString("F:\\project\\go\\GoPath\\src\\demo-basic\\mock\\saic-cap\\rescourses\\CapToken.json"))
	})

	r.POST("/EE/CAP/Detail", func(c *gin.Context) {
		// 获取请求头内容
		//util.DealContext(c)
		// 等待30秒
		//time.Sleep(5*time.Second)
		// 直接返回string信息
		c.String(http.StatusOK, util.GetFileString("F:\\project\\go\\GoPath\\src\\demo-basic\\mock\\saic-cap\\rescourses\\CapDetail.json"))
	})

	r.Run(":8001")
}
