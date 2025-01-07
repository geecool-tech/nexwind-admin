package admin

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	user_model "nexwind.net/admin/server/api/model/user"
	"nexwind.net/admin/server/constant"
	"time"

	"nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/app/auth"
	"nexwind.net/admin/server/utils"
)

// GetUserList 获取用户列表
func GetUserList(param admin.UserListReq) []auth.UserInfo {
	var (
		list  []auth.UserInfo
		query = app.Db
	)
	if param.Mobile != "" {
		query = query.Where("mobile=?", param.Mobile)
	}
	if param.Account != "" {
		query = query.Where("account=?", param.Account)
	}
	if param.Name != "" {
		query = query.Where("name like '%?%'", param.Nickname)
	}
	if param.Email != "" {
		query = query.Where("email=?", param.Email)
	}
	query.Find(&list).Limit(param.PageSize).Offset((param.Current - 1) * param.PageSize)

	return list

}

// SaveUser 保存用户
func SaveUser(param admin.UserInfoReq) error {
	var record auth.UserInfo
	if param.ID == 0 {
		rows := app.Db.Model(&auth.UserInfo{}).Where("account=?", param.Account).Find(&record).RowsAffected
		if rows > 0 {
			return errors.New("用户已存在")
		}
		randStr := utils.GenerateShortUUID(8)
		password, err := HashPassword(param.Password)
		if err != nil {
			return err
		}
		param.Password = password
		param.AccountID = randStr
		err = app.Db.Create(&param).Error
		var relation = map[string]any{"user_id": param.ID, "role_id": 2}
		app.Db.Table("sys_admin_role_relation").Create(&relation)
		if err != nil {
			return err
		}
		_, err = app.MenuEnforcer.AddRoleForUser(fmt.Sprintf("user:%v", param.ID), fmt.Sprintf("role:%v", 2))
		if err != nil {
			return err
		}
		_, err = app.ApiEnforcer.AddRoleForUser(fmt.Sprintf("user:%v", param.ID), fmt.Sprintf("role:%v", 2))
		if err != nil {
			return err
		}
		return nil
	}
	exitRecord := auth.UserInfo{}
	exitUserCount := app.Db.Model(&auth.UserInfo{}).Where("account=?", param.Account).Where("id<>?", param.ID).Find(&exitRecord).RowsAffected
	if exitUserCount > 0 {
		return errors.New(fmt.Sprintf("用户[ %s ]已存在", param.Account))
	}
	app.Db.Model(&auth.UserInfo{}).Where("id=?", param.ID).Find(&exitRecord)
	if param.Password != "" {
		param.Password, _ = HashPassword(param.Password)
	} else {
		param.Password = exitRecord.Password
	}
	return app.Db.Model(&auth.UserInfo{}).Where("id=?", param.ID).Updates(&param).Error
}

// DelUser 删除用户
func DelUser(id int, c *gin.Context) error {
	var (
		getAuth, _ = app.GetAuth(c)
		delUser    = &auth.UserInfo{ID: id}
	)
	if getAuth.UserInfo.ID == id {
		return errors.New("不可删除自己")
	}
	return app.Db.Model(&auth.UserInfo{}).Delete(&delUser).Error
}

// GetCurrentRoles 获取当前用户的角色
func GetCurrentRoles(id int) (map[string]any, error) {
	var (
		selectableIds   []int
		selectableRoles []user_model.Role
		currentRole     user_model.Role
		err             error
	)
	err = app.Db.Table("sys_admin_role_relation").Where("user_id=?", id).Select("role_id").Find(&selectableIds).Error
	if err != nil {
		return nil, err
	}
	err = app.Db.Model(&user_model.Role{}).Where("id in ?", selectableIds).Find(&selectableRoles).Error
	if err != nil {
		return nil, err
	}
	roles, _ := app.MenuEnforcer.GetRolesForUser(fmt.Sprintf("user:%d", id))
	if len(roles) == 0 {
		return nil, errors.New("当前用户无权限")
	}
	fmt.Println(roles)
	trim := strutil.SplitAndTrim(roles[0], ":")

	if len(trim) == 0 {
		return nil, errors.New("当前用户无权限")
	}
	roleId := trim[1]
	err = app.Db.Where("id=?", roleId).Find(&currentRole).Error
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"current":    currentRole,
		"selectable": selectableRoles,
	}, nil

}

// SwitchRole 切换角色
func SwitchRole(uid int, param admin.RoleSwitchReq) error {
	var (
		selectableRoleIds []int
		err               error
	)
	app.Db.Table("sys_admin_role_relation").Where("user_id=?", uid).Select("role_id").Find(&selectableRoleIds)
	if !slice.Contain(selectableRoleIds, param.After) {
		return errors.New("无权切换到该角色")
	}
	fmt.Println(fmt.Sprintf("user:%d", uid), fmt.Sprintf("role:%d", param.Prev))
	_, err = app.MenuEnforcer.DeleteRoleForUser(fmt.Sprintf("user:%d", uid), fmt.Sprintf("role:%d", param.Prev))
	if err != nil {
		return err
	}
	_, err = app.ApiEnforcer.DeleteRoleForUser(fmt.Sprintf("user:%d", uid), fmt.Sprintf("role:%d", param.Prev))
	if err != nil {
		return err
	}
	_, err = app.ApiEnforcer.AddRoleForUser(fmt.Sprintf("user:%d", uid), fmt.Sprintf("role:%d", param.After))
	if err != nil {
		return err
	}
	_, err = app.MenuEnforcer.AddRoleForUser(fmt.Sprintf("user:%d", uid), fmt.Sprintf("role:%d", param.After))
	if err != nil {
		return err
	}
	err = app.MenuEnforcer.LoadPolicy()
	if err != nil {
		return err
	}
	err = app.ApiEnforcer.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

// Logout 登出
func Logout(claims *constant.UserClaims, token string, uid int) error {
	logout := user_model.Logout{
		Uid:         uid,
		Token:       token,
		ExpiredTime: time.Unix(claims.ExpiresAt, 0),
	}
	return app.Db.Create(&logout).Error
}
