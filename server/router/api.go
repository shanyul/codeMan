package router

import (
	"chujian-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resources/view/*")
	// 文件处理,上传操作需要登录
	r.Static("/static", "resources/static")

	api := new(controller.Base)
	r.POST("/get-code", api.IndexApi.GetCode)
	r.GET("/", api.IndexApi.Index)

	return r
}
