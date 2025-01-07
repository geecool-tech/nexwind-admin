package role

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/admin"
)

func RegisterRole(g *gin.RouterGroup) {
	var c = admin.Admin{}
	g.GET("/list", c.RoleList)
	g.GET("/authNode", c.AuthNode)
	g.GET("/roleAccess", c.RoleAccess)
	g.POST("/accreditMenu", c.AccreditMenu)
	g.POST("/accreditApi", c.AccreditApi)
	g.GET("/userRoles", c.UserRole)
	g.POST("/setUserRole", c.SetUserRole)
	g.POST("/save", c.SaveRole)
	g.GET("/delete", c.DelRole)

}
