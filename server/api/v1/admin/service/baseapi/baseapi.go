package baseapi

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"nexwind.net/admin/server/api/model/baseapi"
	"nexwind.net/admin/server/app"
)

type ApiListRes struct {
	Scope string
	Group []struct {
		Group   string
		ApiList []baseapi.SysBaseApi
	}
}

// DelBaseApi 删除 api
func DelBaseApi(id int) (err error) {
	var record baseapi.SysBaseApi
	err = app.Db.Model(&baseapi.SysBaseApi{}).Where("id=?", id).Delete(&record).Error
	if err != nil {
		return
	}
	//router.SyncApiList()
	//_, err = app.ApiEnforcer.GetAllRoles()
	return
	//for _, role := range roles {
	//	app.ApiEnforcer.
	//}
	return
}

// GetBaseApiList 获取 base_api 列表
func GetBaseApiList() (map[string]any, error) {
	var (
		tree             []map[string]any
		allApiList       []baseapi.SysBaseApi
		allScope         = map[string][]string{"回收站": {"未分组"}}
		err              error
		formattedApiList = make(map[string]map[string][]baseapi.SysBaseApi)
		apiListGroupHash = make(map[string][]baseapi.SysBaseApi)
	)
	err = app.Db.Find(&allApiList).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allApiList {
		if _, ok := formattedApiList[v.Scope]; !ok {
			formattedApiList[v.Scope] = make(map[string][]baseapi.SysBaseApi)
		}
		if !slice.Contain(allScope[v.Scope], v.Group) {
			allScope[v.Scope] = append(allScope[v.Scope], v.Group)
		}
		formattedApiList[v.Scope][v.Group] = append(formattedApiList[v.Scope][v.Group], v)
		apiListGroupHash[fmt.Sprintf("%s-%s", v.Scope, v.Group)] = append(apiListGroupHash[fmt.Sprintf("%s-%s", v.Scope, v.Group)], v)

	}
	for sk, scope := range allScope {
		t := map[string]any{
			"title":    sk,
			"id":       fmt.Sprintf("id-%s", sk),
			"children": "",
		}
		var c []map[string]any
		for _, group := range scope {

			c = append(c, map[string]any{
				"title":    group,
				"id":       fmt.Sprintf("id-%s-%s", sk, group),
				"children": apiListGroupHash[fmt.Sprintf("%s-%s", sk, group)],
			})

		}
		t["children"] = c
		tree = append(tree, t)
	}
	return map[string]any{
		"all_scope":      allScope,
		"formatted_list": formattedApiList,
		"tree":           tree,
	}, nil

}

// SaveBaseApi 保存 api
func SaveBaseApi(param baseapi.SysBaseApi) error {
	return app.Db.Model(&baseapi.SysBaseApi{}).Where("id=?", param.ID).Updates(&param).Error
}
