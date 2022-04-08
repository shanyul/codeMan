package controller

import (
	"chujian-api/internal/service"
	"chujian-api/pkg/app"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
	IndexService service.IndexService
}

func (api *Index) Index(c *gin.Context) {
	app.ResponseHtml(c, "index.tmpl", http.StatusOK, 200, nil, "")
}

func (api *Index) GetCode(c *gin.Context) {
	content := c.PostForm("content")
	fmt.Println("content:", content)
	if content == "" {
		app.Response(c, http.StatusOK, 400, nil, "请输入生成内容")
		return
	}
	code, err := api.IndexService.GetCode(content)
	if err != nil {
		app.Response(c, http.StatusOK, 400, nil, "获得二维码错误")
		return
	}
	response := make(map[string]interface{})
	response["code"] = "data:image/png;base64," + base64.StdEncoding.EncodeToString(code)
	app.Response(c, http.StatusOK, 200, response, "")
}
