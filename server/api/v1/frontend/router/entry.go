package router

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/frontend/router/article"
)

// InitV1FrontendRouter 初始化V1 Frontend 路由
func InitV1FrontendRouter(group *gin.RouterGroup) {
	// register for frontend-related routers.
	article.RegisterArticle(group.Group("/article"))
}
