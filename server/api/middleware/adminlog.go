package middleware

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/model/sys"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
)

// AdminLogMiddleWare 访问日志中间件
func AdminLogMiddleWare(c *gin.Context) {
	var (
		apiInfo   = app.ApiListHash[c.Request.URL.Path]
		auth, err = app.GetAuth(c)
	)
	if apiInfo.NeedLogin != 1 {
		c.Next()
		return
	}
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_ILLEGAL_TOKEN})
		c.Abort()
	}
	log := &sys.AdminLog{
		UID:      auth.UserInfo.ID,
		Nickname: auth.UserInfo.Name,
		APIName:  apiInfo.Title,
		Path:     apiInfo.Path,
		ApiID:    apiInfo.ID,
		Method:   c.Request.Method,
		ClientIp: c.ClientIP(),
		Param:    "",
	}
	app.Db.Create(&log)
	c.Next()
}
