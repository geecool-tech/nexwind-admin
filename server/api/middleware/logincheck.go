package middleware

import (
	"github.com/gin-gonic/gin"
	user_model "nexwind.net/admin/server/api/model/user"
	"nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
	"strings"
)

// LoginCheck 登录检测
func LoginCheck(c *gin.Context) {
	var (
		apiInfo       = app.ApiListHash[c.Request.URL.Path]
		banTokenCount int64
		logoutRes     = user_model.Logout{}
	)
	if apiInfo.NeedLogin != 1 {
		c.Next()
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_TOKEN_NOT_EXIST})
		c.Abort()
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	banTokenCount = app.Db.Model(&user_model.Logout{}).Where("token=?", token).Find(&logoutRes).RowsAffected
	if banTokenCount > 0 {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_ILLEGAL_TOKEN})
		c.Abort()
		return
	}
	jwt, err := admin.ValidateJWT(token)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_ILLEGAL_TOKEN})
		c.Abort()
		return
	}
	c.Set("auth", jwt)
	c.Set("accountId", jwt.AccountId)
	c.Next()
}
