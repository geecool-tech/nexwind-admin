package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexwind.net/admin/server/initialize/config"
)

// Cors cors middleware.
func Cors(conf *config.Conf) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 当Access-Control-Allow-Credentials为true时
			for _, allowedOrigin := range conf.Cors {
				if allowedOrigin == origin {
					c.Header("Access-Control-Allow-Origin", origin)
					break
				}
			}
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 可选
			c.Set("content-type", "application/json")   // 可选
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
