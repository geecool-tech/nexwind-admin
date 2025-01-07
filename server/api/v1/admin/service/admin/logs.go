package admin

import (
	"nexwind.net/admin/server/api/model/sys"
	"nexwind.net/admin/server/api/v1/admin/vo/common"
	"nexwind.net/admin/server/app"
)

// QueryLogList 查询日志列表
func QueryLogList(param common.ListReq) ([]sys.AdminLog, int64, error) {
	var (
		list  []sys.AdminLog
		err   error
		total int64
	)
	query := app.Db.Model(&sys.AdminLog{}).
		Count(&total)
	if param.ApiName != "" {
		query = query.Where("api_name like ?", "%"+param.ApiName+"%")
	}
	if param.Path != "" {
		query = query.Where("path like ?", "%"+param.Path+"%")
	}
	if param.Nickname != "" {
		query = query.Where("nickname like ?", "%"+param.Nickname+"%")
	}
	err = query.Debug().Limit(param.PageSize).
		Offset((param.Current - 1) * param.PageSize).
		Find(&list).Error
	if err != nil {
		return list, 0, err
	}
	return list, total, nil
}
