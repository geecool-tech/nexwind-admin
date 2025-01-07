package initialize

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	app "nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/apirbac"
	"nexwind.net/admin/server/initialize/config"
	"nexwind.net/admin/server/initialize/db"
	"nexwind.net/admin/server/initialize/log"
	menurbac "nexwind.net/admin/server/initialize/menurbac"
	"nexwind.net/admin/server/initialize/redis"
	"nexwind.net/admin/server/initialize/resp"
	"nexwind.net/admin/server/initialize/router"
)

func Init() {
	var (
		currentEnv = constant.PROD
	)
	debug := flag.Bool("d", false, "调试模式")
	flag.Parse()
	if *debug {
		currentEnv = constant.DEV
	}
	app.Resp = resp.InitResp()
	app.Log = log.InitZap()
	defer func(Log *zap.Logger) {
		err := Log.Sync()
		if err != nil {
			_ = fmt.Errorf("日志服务初始化失败: %v", err)
		}
	}(app.Log)
	app.Conf = config.InitConfig(debug, currentEnv, app.Log)
	app.Db = db.InitSysDb(app.Conf, app.Log)
	app.Log.Info("当前运行模式", zap.String("env", currentEnv), zap.Bool("is_debug", *debug))
	app.MySQLSet = db.InitMySQLSet(app.Conf, app.Log)
	app.RedisSet = redis.InitRedisSet(app.Conf, app.Log)
	app.ApiEnforcer = apirbac.InitApiRBAC(app.Db)
	app.MenuEnforcer = menurbac.InitMenuRBAC(app.Db)
	router.InitRouter(app.Conf)
}
