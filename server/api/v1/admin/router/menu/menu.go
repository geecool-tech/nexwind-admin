package menu

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/controller/admin"
)

func RegisterMenu(g *gin.RouterGroup) {
	var c = &admin.Admin{}
	{
		g.GET("/list", c.MenuList)
		g.POST("/save", c.SaveMenu)
		g.GET("/delete", c.DeleteMenu)
	}
}
