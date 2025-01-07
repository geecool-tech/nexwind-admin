package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"nexwind.net/admin/server/api/middleware"
	"nexwind.net/admin/server/api/middleware/apirbac"
	"nexwind.net/admin/server/api/model/baseapi"
	v1AdminRouter "nexwind.net/admin/server/api/v1/admin/router"
	v1FrontendRouter "nexwind.net/admin/server/api/v1/frontend/router"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/initialize/config"
)

func InitRouter(conf *config.Conf) {
	router := gin.Default()
	router.Use(middleware.Cors(conf))
	// init v1 version router.
	v1AdminRouter.InitV1AdminRouter(router.Group("/api/v1/admin", middleware.LoginCheck, apirbac.ApiRBAC, middleware.AdminLogMiddleWare))
	v1FrontendRouter.InitV1FrontendRouter(router.Group("/api/v1/frontend", middleware.LoginCheck, apirbac.ApiRBAC))
	if conf.Server.Port == 0 {
		conf.Server.Port = 8888
	}
	go syncApi(router)
	go SyncApiList()
	router.Static("/uploads", "./uploads")
	err := router.Run(fmt.Sprintf(":%d", conf.Server.Port))
	if err != nil {
		fmt.Println("启动错误....")
	}
}

// syncApi 同步 Api 到表内
func syncApi(e *gin.Engine) {
	var (
		count = 0
	)
	for _, route := range e.Routes() {
		api := &baseapi.SysBaseApi{
			Path:      route.Path,
			Method:    route.Method,
			Scope:     "未划分",
			Group:     "未分组",
			NeedAuth:  1,
			NeedLogin: 1,
		}
		app.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&api)
		if api.ID != 0 {
			count += 1
			_, _ = app.ApiEnforcer.AddPolicy("role:1", fmt.Sprintf("api:%d", api.ID), api.Method)
		}
	}
	app.Log.Info(fmt.Sprintf("成功同步 %d 条新 Api 到数据库 ", count), zap.Int("count", count))
}

// SyncApiList 同步 Api 列表
func SyncApiList() {
	var (
		allApiList []baseapi.SysBaseApi
	)
	app.Db.Where("deleted_time is null").Find(&allApiList)
	app.ApiListHash = make(map[string]baseapi.SysBaseApi)
	app.ApiListIdMap = make(map[int]baseapi.SysBaseApi)
	for _, v := range allApiList {
		app.ApiListHash[v.Path] = v
		app.ApiListIdMap[v.ID] = v
	}
}
