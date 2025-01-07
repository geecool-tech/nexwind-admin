package admin

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	userModel "nexwind.net/admin/server/api/model/user"
	adminVo "nexwind.net/admin/server/api/v1/admin/vo/admin"
	"nexwind.net/admin/server/app"
	"strconv"
)

func GetRoleAccess(roleId int) map[string]any {
	var (
		apiIds  []int
		menuIds []int
	)
	apiPolicy, _ := app.ApiEnforcer.GetFilteredPolicy(0, fmt.Sprintf("role:%d", roleId))
	for _, v := range apiPolicy {
		apiIdSpilt := strutil.SplitAndTrim(v[1], ":")
		apiIdInt, transErr := strconv.Atoi(apiIdSpilt[1])
		if transErr == nil {
			apiIds = append(apiIds, apiIdInt)
		}

	}
	menuPolicy, _ := app.MenuEnforcer.GetFilteredPolicy(0, fmt.Sprintf("role:%d", roleId))
	for _, v := range menuPolicy {
		var menuIdInt int
		menuIdSpilt := strutil.SplitAndTrim(v[1], ":")
		menuIdInt, _ = strconv.Atoi(menuIdSpilt[1])
		menuIds = append(menuIds, menuIdInt)
	}
	if apiIds == nil {
		apiIds = []int{}
	}
	if menuIds == nil {
		menuIds = []int{}
	}
	return map[string]any{
		"api":  apiIds,
		"menu": menuIds,
	}
}

// GetRoleList 获取角色列表
func GetRoleList(param adminVo.RoleListReq) []userModel.Role {
	var (
		query = app.Db.Model(&userModel.Role{})
		list  []userModel.Role
	)
	if param.Name != "" {
		query = query.Where("name=?", param.Name)
	}
	query.Find(&list).Limit(param.PageSize).Offset((param.Current - 1) * param.PageSize)
	return list

}

// SetMenuAuth 设置菜单权限
func SetMenuAuth(param adminVo.RoleAuthReq) error {
	var (
		addPolicies [][]string
		delPolicies [][]string
		err         error
	)
	for _, v := range param.Add {
		addPolicies = append(addPolicies, []string{
			fmt.Sprintf("role:%v", param.RoleId),
			fmt.Sprintf("menu:%v", v),
			"GET",
		})
	}
	for _, v := range param.Remove {
		delPolicies = append(delPolicies, []string{
			fmt.Sprintf("role:%v", param.RoleId),
			fmt.Sprintf("menu:%v", v),
			"GET",
		})
	}
	if len(addPolicies) > 0 {
		_, err = app.MenuEnforcer.AddPolicies(addPolicies)
		if err != nil {
			return err
		}
	}
	if len(delPolicies) > 0 {
		_, err = app.MenuEnforcer.RemovePolicies(delPolicies)
		if err != nil {
			return err
		}
	}
	err = app.MenuEnforcer.LoadPolicy()
	if err != nil {
		return err
	}
	if err = app.MenuEnforcer.LoadPolicy(); err != nil {
		return err
	}
	return nil
}

// SetApiAuth 设置Api权限
func SetApiAuth(param adminVo.RoleAuthReq) error {
	var (
		addPolicies [][]string
		delPolicies [][]string
		err         error
	)
	for _, v := range param.Add {
		act := ""
		switch v.(type) {
		case float64:
			apiInt := int(v.(float64))
			act = app.ApiListIdMap[apiInt].Method
		case string:
			act = "GET"
		}
		addPolicies = append(addPolicies, []string{
			fmt.Sprintf("role:%v", param.RoleId),
			fmt.Sprintf("api:%v", v),
			act,
		})
	}
	for _, v := range param.Remove {
		act := ""
		switch v.(type) {
		case float64:
			apiInt := int(v.(float64))
			act = app.ApiListIdMap[apiInt].Method
		case string:
			act = "GET"
		}
		delPolicies = append(delPolicies, []string{
			fmt.Sprintf("role:%v", param.RoleId),
			fmt.Sprintf("api:%v", v),
			act,
		})
	}
	if len(addPolicies) > 0 {
		_, err = app.ApiEnforcer.AddPolicies(addPolicies)
		if err != nil {
			return err
		}
	}
	if len(delPolicies) > 0 {
		_, err = app.ApiEnforcer.RemovePolicies(delPolicies)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUserRole 获取用户角色
func GetUserRole(uid int) []int {
	var (
		roleIds []int
	)
	app.Db.Table("sys_admin_role_relation").Where("user_id=?", uid).Select("role_id").Find(&roleIds)
	return roleIds
}

// SetUserRole 设置用户角色
func SetUserRole(param adminVo.SetRoleReq) error {
	var (
		err    error
		roleId string
	)
	roles, _ := app.MenuEnforcer.GetRolesForUser(fmt.Sprintf("user:%v", param.Uid))
	fmt.Println(roles)
	if len(roles) > 0 {
		currentUserRole := strutil.SplitAndTrim(roles[0], ":")
		roleId = currentUserRole[1]
	}
	if len(param.Add) > 0 {
		var relations []map[string]any
		for _, v := range param.Add {
			relations = append(relations, map[string]any{
				"user_id": param.Uid,
				"role_id": v,
			})
			_, err = app.ApiEnforcer.AddRoleForUser(fmt.Sprintf("user:%v", param.Uid), fmt.Sprintf("role:%v", v))
			if err != nil {
				return err
			}
			_, err = app.MenuEnforcer.AddRoleForUser(fmt.Sprintf("user:%v", param.Uid), fmt.Sprintf("role:%v", v))
			if err != nil {
				return err
			}
		}
		err = app.Db.Table("sys_admin_role_relation").Create(&relations).Error
		if err != nil {
			return err
		}

	}
	if len(param.Remove) > 0 {
		var relations []map[string]any
		for _, v := range param.Remove {
			relations = append(relations, map[string]any{
				"user_id": param.Uid,
				"role_id": v,
			})
			if roleId == strconv.Itoa(int(v.(float64))) {
				roleIdInt, _ := strconv.Atoi(roleId)
				err := SwitchRole(param.Uid, adminVo.RoleSwitchReq{
					Prev:  roleIdInt,
					After: 2,
				})
				if err != nil {
					return err
				}
			}
			_, err = app.ApiEnforcer.DeleteRoleForUser(fmt.Sprintf("user:%v", param.Uid), fmt.Sprintf("role:%v", v))
			if err != nil {
				return err
			}
			_, err = app.MenuEnforcer.DeleteRoleForUser(fmt.Sprintf("user:%v", param.Uid), fmt.Sprintf("role:%v", v))

			if err != nil {
				return err
			}
		}
		err = app.Db.Table("sys_admin_role_relation").Where("user_id=? and role_id in ?", param.Uid, param.Remove).Delete(&relations).Error
		if err != nil {
			return err
		}

	}
	return nil
}

// SaveRole 保存角色
func SaveRole(param userModel.Role) error {
	if param.ID == 0 {
		return app.Db.Model(&userModel.Role{}).Create(&param).Error
	}
	return app.Db.Model(&userModel.Role{}).Where("id=?", param.ID).Updates(&param).Error
}

// DelRole 删除角色
func DelRole(id string) error {
	var (
		relationCount int64
		err           error
	)
	if id == "1" || id == "2" {
		return errors.New("超管/访客 角色不可删除")
	}
	app.Db.Table("sys_admin_role_relation").Where("role_id=?", id).Count(&relationCount)
	if relationCount > 0 {
		return errors.New("该角色已关联用户,不可删除")
	}
	_, err = app.MenuEnforcer.DeleteRole(fmt.Sprintf("role:%v", id))
	if err != nil {
		return err
	}
	_, err = app.ApiEnforcer.DeleteRole(fmt.Sprintf("role:%v", id))
	if err != nil {
		return err
	}
	return app.Db.Model(&userModel.Role{}).Where("id=?", id).Delete(&userModel.Role{}).Error
}
