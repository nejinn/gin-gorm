package view

import (
	"errors"
	"note/global"
)

type UserNav struct {
	Name     string `json:"name"`      // 导航名称
	Icon     string `json:"icon"`      // 导航icon
	ParentId int    `json:"parent_id"` // 父级菜单id
	Router   string `json:"router"`    // 路由
	Order    int    `json:"order"`     // 排序
	CreateId int    `json:"create_id"`
	UpdateId int    `json:"update_id"`
}

func (e *UserNav) CreateUserNav() (err error) {
	if e.Name == "" {
		return errors.New("导航菜单不能为空")
	}
	// icon 默认为 oi oi-list-rich
	if e.Icon == "" {
		e.Icon = "oi oi-list-rich"
	}
	if e.Order == 0 {
		return errors.New("排序应该大于等于1")
	}
	err = global.NLY_DB.Create(&e).Error
	return err
}
