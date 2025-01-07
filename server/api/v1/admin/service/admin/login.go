package admin

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	adminVo "nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/app/auth"
	"nexwind.net/admin/server/constant"
	"time"
)

// GenerateJWT 生成jwt
func GenerateJWT(accountId string) (string, error) {
	claims := constant.UserClaims{
		AccountId: accountId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.ExpireDays).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(constant.JwtKey)
}

// ValidateJWT  验证JWT
func ValidateJWT(tokenString string) (*constant.UserClaims, error) {
	claims := &constant.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return constant.JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// HashPassword 密码加盐
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyHashPassword 验证hash密码
func VerifyHashPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// VerifyPwd 用户登录逻辑
func VerifyPwd(param adminVo.UserLoginReq) (token string, err error) {
	var userinfo auth.UserInfo
	recordCount := app.Db.Model(&auth.UserInfo{}).Where("account = ?", param.Account).First(&userinfo).RowsAffected
	if recordCount == 0 {
		return "", errors.New("用户不存在")
	}
	if err := VerifyHashPassword(userinfo.Password, param.Password); err != nil {
		return "", errors.New("密码错误")
	}
	generateJWT, err := GenerateJWT(userinfo.AccountID)
	if err != nil {
		return "", errors.New("token生成错误")
	}
	return generateJWT, nil
}
