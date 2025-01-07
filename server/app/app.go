package app

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexwind.net/admin/server/api/model/baseapi"
	"nexwind.net/admin/server/app/auth"
	uAuth "nexwind.net/admin/server/app/auth"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/config"
	"nexwind.net/admin/server/initialize/resp"
	"sync"
)

var (
	Db           *gorm.DB
	Log          *zap.Logger
	Conf         *config.Conf
	Resp         resp.Resp
	MySQLSet     map[string]*gorm.DB
	RedisSet     map[string]*redis.Client
	userAuth     *auth.Auth
	authLock     sync.Mutex
	ApiEnforcer  *casbin.Enforcer
	ApiListHash  map[string]baseapi.SysBaseApi
	ApiListIdMap map[int]baseapi.SysBaseApi
	MenuEnforcer *casbin.Enforcer
)

// SetAuth set user auth info.
func SetAuth(auth *auth.Auth) {
	authLock.Lock()
	defer authLock.Unlock()
	userAuth = auth
}

// GetAuth get user auth info.
func GetAuth(c *gin.Context) (*auth.Auth, error) {
	var (
		authInfo = &auth.Auth{}
		userInfo = &uAuth.UserInfo{}
		claims   = &constant.UserClaims{}
	)
	value, exists := c.Get("auth")
	if !exists {
		return nil, errors.New("auth not exists")
	}
	claims = value.(*constant.UserClaims)
	affected := Db.Model(&uAuth.UserInfo{}).Where("account_id=?", claims.AccountId).Find(&userInfo).RowsAffected
	if affected == 0 {
		return nil, errors.New("user not exists")
	}
	authInfo.UserInfo = userInfo
	return authInfo, nil
}
