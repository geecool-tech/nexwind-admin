package router

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/router/baseapi"
	"nexwind.net/admin/server/api/v1/admin/router/common"
	"nexwind.net/admin/server/api/v1/admin/router/logs"
	"nexwind.net/admin/server/api/v1/admin/router/menu"
	"nexwind.net/admin/server/api/v1/admin/router/role"
	"nexwind.net/admin/server/api/v1/admin/router/user"
)

// InitV1AdminRouter 初始化V1 Admin 路由
func InitV1AdminRouter(group *gin.RouterGroup) {
	// register for admin-related routers.
	user.RegisterUser(group.Group("/user"))
	menu.RegisterMenu(group.Group("/menu"))
	role.RegisterRole(group.Group("/role"))
	common.RegisterCommon(group.Group("/common"))
	baseapi.RegisterBaseApi(group.Group("/baseApi"))
	logs.RegisterLogs(group.Group("/adminLogs"))
}
