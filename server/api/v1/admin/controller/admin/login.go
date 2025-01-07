package admin

import (
	"github.com/gin-gonic/gin"
	adminService "nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
)

// Login admin 用户登录接口
func (a *Admin) Login(c *gin.Context) {
	var param admin.UserLoginReq
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}

	token, err := adminService.VerifyPwd(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, map[string]any{
		"token": token,
	})
}
