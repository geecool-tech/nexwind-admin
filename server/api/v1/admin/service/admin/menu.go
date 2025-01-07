package admin

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/structs"
	"github.com/duke-git/lancet/v2/strutil"
	"nexwind.net/admin/server/api/model/basemodel"
	"nexwind.net/admin/server/app"
	"strconv"
)

// Meta 定义菜单标签的结构体
type Meta struct {
	Locale             string `json:"locale"`
	RequiresAuth       bool   `json:"requiresAuth"`
	Icon               string `json:"icon,omitempty"` // `omitempty` 表示该字段可选
	Order              int    `json:"order,omitempty"`
	HideInMenu         bool   `json:"hideInMenu,omitempty"`
	HideChildrenInMenu bool   `json:"hideChildrenInMenu,omitempty"`
}

// Value 定义序列化
func (m *Meta) Value() (driver.Value, error) {
	j, _ := json.Marshal(m)
	return string(j), nil
}

// Scan Implement the sql.Scanner interface for Meta struct
func (m *Meta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal Meta value: %v", value)
	}
	return json.Unmarshal(bytes, m)
}

// Menu admin 菜单结构体
type Menu struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id,omitempty"`
	Title    string `json:"title"`
	Meta     Meta   `json:"meta"`
	Order    int    `json:"order"`
	Type     uint8  `json:"type" gorm:"column:type"`
	Children []Menu `json:"children,omitempty" gorm:"-"`
	basemodel.BaseModel
}

func (*Menu) TableName() string {
	return "sys_admin_menu"
}

// DelMenu 删除目录
func DelMenu(menuId int) error {
	var (
		recordCount int64
	)
	app.Db.Model(&Menu{}).Where("parent_id=?", menuId).Count(&recordCount)
	if recordCount > 0 {
		return errors.New("该菜单下有子菜单，请先删除子菜单")
	}
	return app.Db.Model(&Menu{}).Where("id=?", menuId).Delete(&Menu{}).Error
}

// SaveMenu 保存菜单
func SaveMenu(param Menu) error {
	parsed := structs.New(param)
	parsedParam, _ := parsed.ToMap()
	marshal, _ := json.Marshal(param.Meta)
	parsedParam["meta"] = string(marshal)
	delete(parsedParam, "children")
	if param.ID == 0 {
		var err error
		err = app.Db.Model(&Menu{}).Create(&parsedParam).Error
		if err != nil {
			return err
		}
		_, err = app.MenuEnforcer.AddPolicy(fmt.Sprintf("role:%d", 1), fmt.Sprintf("menu:%d", parsedParam["id"]), "GET")
		if err != nil {
			return err
		}
		err = app.MenuEnforcer.LoadPolicy()
		return err

	} else {
		var err error
		if err != nil {
			return err
		}
		err = app.Db.Model(&Menu{}).Where("id=?", param.ID).Save(&parsedParam).Error
		return err
	}
}

func GetMenuList(id int) []Menu {
	rootMenus, _ := fetchMenus(id)
	return rootMenus
}

func fetchMenus(id int) ([]Menu, error) {
	// 1. 获取所有菜单项，一次性查询
	var (
		allMenus []Menu
		err      error
	)
	if id == 0 {
		err = app.Db.Order("`order` ASC, created_time ASC").Find(&allMenus).Error
	}
	if id != 0 {
		roles, _ := app.MenuEnforcer.GetRolesForUser(fmt.Sprintf("user:%d", id))
		if len(roles) == 0 {
			allMenus = []Menu{}
			return allMenus, errors.New("当前用户无权限")
		}
		trim := strutil.SplitAndTrim(roles[0], ":")

		if len(trim) == 0 {
			allMenus = []Menu{}
			return allMenus, errors.New("当前用户无权限")
		}
		roleId := trim[1]
		policy, err := app.MenuEnforcer.GetFilteredPolicy(0, fmt.Sprintf("role:%v", roleId))
		if err != nil {
			return allMenus, err
		}
		var menuId []int
		for _, v := range policy {
			menuIdInt, _ := strconv.Atoi(strutil.SplitAndTrim(v[1], ":")[1])
			menuId = append(menuId, menuIdInt)
		}
		err = app.Db.Where("id in ?", menuId).Order("`order` ASC, created_time ASC").Find(&allMenus).Error
	}
	if err != nil {
		return nil, err
	}
	// 2. 构建一个 map，key 是父菜单的 ID，value 是对应的子菜单
	menuMap := make(map[int][]Menu)
	var rootMenus []Menu
	for _, menu := range allMenus {
		if menu.ParentID == nil {
			// 没有 parent_id 的菜单是根菜单
			rootMenus = append(rootMenus, menu)
		} else {
			// 存在 parent_id 的菜单作为子菜单存储在 map 中
			parentID := *menu.ParentID
			menuMap[parentID] = append(menuMap[parentID], menu)
		}
	}
	// 3. 递归将子菜单放入父菜单的 Children 字段
	var buildTree func([]Menu) []Menu
	buildTree = func(menus []Menu) []Menu {
		for i := range menus {
			// 如果当前菜单有子菜单，将子菜单放入其 Children 字段
			if children, ok := menuMap[menus[i].ID]; ok {
				menus[i].Children = buildTree(children)
			}
		}
		return menus
	}
	// 4. 返回构建好的菜单树
	return buildTree(rootMenus), nil
}
