package api

import (
	"demo-basic/blog/code/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)


func AppsController(r *gin.Engine) {
	// 读取文件中的字符
	r.GET("/getAllApps", func(c *gin.Context) {
		apps := service.GetAllApps()
		data,err := json.Marshal(apps)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data" : "失败",
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"data":string(data),
			})
		}
	})
}
