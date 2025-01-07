package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	user_model "nexwind.net/admin/server/api/model/user"
	adminService "nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/api/v1/admin/service/baseapi"
	"nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
	"strconv"
)

// RoleAccess 获取角色已有权限
func (a *Admin) RoleAccess(c *gin.Context) {
	var (
		roleId    = c.DefaultQuery("id", "")
		roleIdInt int
	)
	if roleId == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	roleIdInt, _ = strconv.Atoi(roleId)
	access := adminService.GetRoleAccess(roleIdInt)
	app.Resp.SuccessRes(c, access)
}

// AuthNode 获取权限节点
func (a *Admin) AuthNode(c *gin.Context) {
	menuList := adminService.GetMenuList(0)
	apiList, _ := baseapi.GetBaseApiList()
	app.Resp.SuccessRes(c, map[string]any{
		"menu_list": menuList,
		"tree":      apiList["tree"],
	})
}

// RoleList 角色列表
func (a *Admin) RoleList(c *gin.Context) {
	var (
		param admin.RoleListReq
	)
	if err := c.BindQuery(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	roleList := adminService.GetRoleList(param)
	app.Resp.SuccessRes(c, resp.ListRes{List: roleList, Total: len(roleList)})
}

func (a *Admin) AccreditApi(c *gin.Context) {
	var (
		param admin.RoleAuthReq
		err   error
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	err = adminService.SetApiAuth(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	app.Resp.SuccessRes(c, nil)
}
func (a *Admin) AccreditMenu(c *gin.Context) {
	var (
		param admin.RoleAuthReq
		err   error
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	err = adminService.SetMenuAuth(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// UserRole 根据用户id获取角色列表
func (a *Admin) UserRole(c *gin.Context) {
	var uid = c.DefaultQuery("uid", "")
	if uid == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	uidInt, _ := strconv.Atoi(uid)
	roles := adminService.GetUserRole(uidInt)
	app.Resp.SuccessRes(c, roles)
}

// DelRole 删除角色
func (a *Admin) DelRole(c *gin.Context) {
	var (
		roleId = c.DefaultQuery("id", "")
	)
	if roleId == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	err := adminService.DelRole(roleId)
	if err != nil {
		fmt.Println(err)
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// SaveRole 新增&保存角色
func (a *Admin) SaveRole(c *gin.Context) {
	var param user_model.Role
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	err := adminService.SaveRole(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// SetUserRole 设置用户角色
func (a *Admin) SetUserRole(c *gin.Context) {
	var param admin.SetRoleReq
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	err := adminService.SetUserRole(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL})
		return
	}
	app.Resp.SuccessRes(c, nil)
}
