package user

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/admin"
)

func RegisterUser(g *gin.RouterGroup) {
	var c admin.Admin
	{
		g.POST("/login", c.Login)
		g.GET("/info", c.UserInfo)
		g.GET("/authMenu", c.AuthMenu)
		g.GET("/list", c.AdminList)
		g.POST("/save", c.SaveAdmin)
		g.GET("/delete", c.DeleteAdmin)
		g.GET("/roles", c.CurrentRoles)
		g.POST("/switchRole", c.SwitchRole)
		g.GET("logout", c.Logout)
	}
}
