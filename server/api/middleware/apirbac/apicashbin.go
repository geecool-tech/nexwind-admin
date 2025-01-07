package apirbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
)

// ApiRBAC api rbac 中间件
func ApiRBAC(c *gin.Context) {
	var (
		auth, err = app.GetAuth(c)
		apiInfo   = app.ApiListHash[c.Request.URL.Path]
	)
	if apiInfo.NeedAuth != 1 {
		c.Next()
		return
	}
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		c.Abort()
		return
	}
	allowed, err := app.ApiEnforcer.Enforce(fmt.Sprintf("user:%d", auth.UserInfo.ID), fmt.Sprintf("api:%d", apiInfo.ID), c.Request.Method)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		c.Abort()
		return
	}
	if allowed == true {
		c.Next()
		return
	}
	app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_UN_AUTH, Msg: "unauthorized"})
	c.Abort()
}
