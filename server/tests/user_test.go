package tests

import (
	"fmt"
	"nexwind.net/admin/server/api/v1/admin/service/admin"
	"nexwind.net/admin/server/utils"
	"testing"
)

// TestUserGenPwd 生成密码
func TestUserGenPwd(t *testing.T) {
	password, err := admin.HashPassword("123456") // $2a$10$L320F1NchzpMByEcsjq4z.ru/Tn0yqkonIsCgAjd50qhdtTguTe9m
	fmt.Println(password, err)
}

// TestVerifyPwd 验证密码
func TestVerifyPwd(t *testing.T) {
	var hashPwdStr = "$2a$10$L320F1NchzpMByEcsjq4z.ru/Tn0yqkonIsCgAjd50qhdtTguTe9m"
	err := admin.VerifyHashPassword(hashPwdStr, "Adaoz18781371379")
	fmt.Println(err)
}

// TestUUID 验证密码
func TestUUID(t *testing.T) {
	fmt.Println(utils.GenerateShortUUID(8))
}
