package admin

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
	"strconv"
)

// AuthMenu admin菜单列表
func (a *Admin) AuthMenu(c *gin.Context) {
	auth, err := app.GetAuth(c)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, admin.GetMenuList(auth.UserInfo.ID))
}

// MenuList 权限菜单列表
func (a *Admin) MenuList(c *gin.Context) {
	list := admin.GetMenuList(0)
	app.Resp.PageRes(c, resp.ListRes{List: list, Total: len(list)})
}

// DeleteMenu 删除菜单
func (a *Admin) DeleteMenu(c *gin.Context) {
	var (
		id  = c.DefaultQuery("id", "")
		err error
	)
	if id == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	idInt, _ := strconv.Atoi(id)
	err = admin.DelMenu(idInt)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}
func (a *Admin) SaveMenu(c *gin.Context) {
	var (
		param admin.Menu
		err   error
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_INVALID_PARAM})
		return
	}
	err = admin.SaveMenu(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}
