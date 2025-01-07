package logs

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/adminlogs"
)

// RegisterLogs 注册日志相关路由
func RegisterLogs(g *gin.RouterGroup) {
	var c adminlogs.AdminLogs
	g.GET("/list", c.List)
}
