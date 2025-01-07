package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	adminService "nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	appAuth "nexwind.net/admin/server/app/auth"
	"nexwind.net/admin/server/initialize/resp"
	"strconv"
	"strings"
)

// UserInfo admin 用户信息
func (a *Admin) UserInfo(c *gin.Context) {
	auth, err := app.GetAuth(c)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, auth.UserInfo)
}

// AdminList 管理员列表
func (a *Admin) AdminList(c *gin.Context) {
	var (
		err      error
		param    admin.UserListReq
		userList []appAuth.UserInfo
	)

	if err = c.BindQuery(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
	}
	userList = adminService.GetUserList(param)
	for k, v := range userList {
		if !strings.Contains(v.Avatar, "http") {
			userList[k].Avatar = fmt.Sprintf("http://%s%s", c.Request.Host, v.Avatar)
		}
	}
	app.Resp.PageRes(c, resp.ListRes{List: userList, Total: len(userList)})
}

// SaveAdmin 添加管理员
func (a *Admin) SaveAdmin(c *gin.Context) {
	var (
		param admin.UserInfoReq
		err   error
	)

	if err := c.ShouldBindJSON(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	err = adminService.SaveUser(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// DeleteAdmin 删除管理员
func (a *Admin) DeleteAdmin(c *gin.Context) {
	var (
		id = c.DefaultQuery("id", "")
	)
	if id == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: "id不能为空"})
		return
	}
	idInt, _ := strconv.Atoi(id)
	if err := adminService.DelUser(idInt, c); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// SwitchRole 切换角色
func (a *Admin) SwitchRole(c *gin.Context) {
	var (
		param admin.RoleSwitchReq
		err   error
	)
	auth, err := app.GetAuth(c)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	if err := adminService.SwitchRole(auth.UserInfo.ID, param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)

}

// CurrentRoles 当前用户的角色
func (a *Admin) CurrentRoles(c *gin.Context) {
	var err error
	auth, err := app.GetAuth(c)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	data, err := adminService.GetCurrentRoles(auth.UserInfo.ID)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, data)
}

// Logout 退出登录
func (a *Admin) Logout(c *gin.Context) {
	auth, err := app.GetAuth(c)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	jwt, err := adminService.ValidateJWT(token)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	err = adminService.Logout(jwt, token, auth.UserInfo.ID)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}
