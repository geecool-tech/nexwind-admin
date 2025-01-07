package menurbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitMenuRBAC(db *gorm.DB) *casbin.Enforcer {
	var err error
	adapter, err := gormAdapter.NewAdapterByDBUseTableName(db, "", "sys_rbac_menu")
	if err != nil {
		panic(err)
	}
	rbacModel, err := model.NewModelFromString(getRbacModel())
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer(rbacModel, adapter)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	return enforcer
}

// getRbacModel 返回rbac_model
func getRbacModel() string {
	return `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`
}
