package article

import "github.com/gin-gonic/gin"

// RegisterArticle 注册文章相关路由
func RegisterArticle(g *gin.RouterGroup) {
	g.GET("/detail")
}
