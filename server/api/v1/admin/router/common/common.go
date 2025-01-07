package common

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/common"
)

// RegisterCommon 注册公共 api
func RegisterCommon(g *gin.RouterGroup) {
	var c common.Common
	g.POST("/upload", c.Upload)
}
