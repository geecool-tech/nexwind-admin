package adminlogs

import (
	"github.com/gin-gonic/gin"
	"nexwind.net/admin/server/api/model/sys"
	"nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/api/v1/admin/vo/common"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
)

// AdminLogs Admin 操作日志
type AdminLogs struct {
}

// List 管理员日志列表
func (a *AdminLogs) List(c *gin.Context) {
	var (
		param common.ListReq
		list  []sys.AdminLog
		total int64
		err   error
	)
	if err = c.BindQuery(&param); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: resp.GetMsg(constant.CODE_FAIL)})
		return
	}
	list, total, err = admin.QueryLogList(param)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: resp.GetMsg(constant.CODE_FAIL)})
		return
	}
	app.Resp.PageRes(c, resp.ListRes{
		Total: int(total),
		List:  list,
	})

}
