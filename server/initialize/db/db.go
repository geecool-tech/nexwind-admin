package db

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nexwind.net/admin/server/initialize/config"
)

var _db *gorm.DB

// InitSysDb init system db.
func InitSysDb(conf *config.Conf, l *zap.Logger) *gorm.DB {
	var err error
	if _db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.SysDb.User,
			conf.SysDb.Pwd,
			conf.SysDb.Host,
			conf.SysDb.Port,
			conf.SysDb.Database,
		)
		marshal, _ := json.Marshal(conf.SysDb)
		fmt.Println(string(marshal))
		if _db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error), // 修改调试模式为error
		}); err != nil {
			l.Panic("sys_db connect failed.", zap.Any("conf", conf.SysDb))
		}
	}
	return _db
}

// InitMySQLSet init MySQL set.
func InitMySQLSet(conf *config.Conf, l *zap.Logger) map[string]*gorm.DB {
	var (
		mysqlSet = make(map[string]*gorm.DB)
		db       *gorm.DB
		err      error
	)
	for _, v := range conf.MySQLSet {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			v.User,
			v.Pwd,
			v.Host,
			v.Port,
			v.Database,
		)
		if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error), // 修改调试模式为error
		}); err != nil {
			l.Panic("failed init mysql instance.", zap.Any("conf", v))
		}
		mysqlSet[v.Key] = db
	}
	return mysqlSet
}
