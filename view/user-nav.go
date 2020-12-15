package view

import "errors"

type UserNav struct {
	Name     string `json:"name"`       // 导航名称
	Icon     string `json:"icon"`       // 导航icon
	ParentId int    `json:"data_group"` // 父级菜单id
	Router   string `json:"router"`     // 路由
	Order    string `json:"order"`      // 排序
}

func (e UserNav) createUserNav() (err error) {
	if e.Name == "" {
		return errors.New("导航菜单不能为空")
	}
	// icon 默认为 oi oi-list-rich
	if e.Icon == "" {
		e.Icon = "oi oi-list-rich"
	}

	if e.ParentId ==""
}
