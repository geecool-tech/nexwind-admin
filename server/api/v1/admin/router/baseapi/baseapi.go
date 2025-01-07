package baseapi

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/baseapi"
)

// RegisterBaseApi 注册 base_api 相关路由
func RegisterBaseApi(g *gin.RouterGroup) {
	var c baseapi.BaseApi
	g.GET("/list", c.List)
	g.POST("/save", c.Save)
	g.GET("/delete", c.Delete)
}
