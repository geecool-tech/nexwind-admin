package baseapi

import (
	"github.com/gin-gonic/gin"
	baseApiModel "nexwind.net/admin/server/api/model/baseapi"
	"nexwind.net/admin/server/api/v1/admin/service/baseapi"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/initialize/resp"
	"strconv"
)

type BaseApi struct {
}

// List 获取 base_api 列表
func (b *BaseApi) List(c *gin.Context) {
	list, err := baseapi.GetBaseApiList()
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, list)
}

// Delete 删除 base_api
func (b *BaseApi) Delete(c *gin.Context) {
	var id = c.DefaultQuery("id", "")
	if id == "" {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: "id不能为空"})
		return
	}
	idInt, _ := strconv.Atoi(id)
	if err := baseapi.DelBaseApi(idInt); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	app.Resp.SuccessRes(c, nil)
}

// Save 保存 base_api
func (b *BaseApi) Save(c *gin.Context) {
	var (
		param baseApiModel.SysBaseApi
		err   error
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	err = baseapi.SaveBaseApi(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: -1, Msg: err.Error()})
		return
	}
	SyncApiList()
	app.Resp.SuccessRes(c, nil)
}

// SyncApiList 同步 Api 列表
func SyncApiList() {
	var (
		allApiList []baseApiModel.SysBaseApi
	)
	app.Db.Where("deleted_time is null").Find(&allApiList)
	app.ApiListHash = make(map[string]baseApiModel.SysBaseApi)
	app.ApiListIdMap = make(map[int]baseApiModel.SysBaseApi)
	for _, v := range allApiList {
		app.ApiListHash[v.Path] = v
		app.ApiListIdMap[v.ID] = v
	}
}
