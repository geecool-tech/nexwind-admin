package constant

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JwtKey jwt秘钥
var JwtKey = []byte("nexwind_jwt_0718")

// ExpireDays 设置过期时间
var ExpireDays = 7 * 24 * time.Hour

// UserClaims 用于存储用户信息
type UserClaims struct {
	AccountId string `json:"account_id"`
	jwt.StandardClaims
}
