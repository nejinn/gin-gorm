package view

import (
	"errors"
	"note/global"
	"note/model"
)

func CreateNav(e *model.UserNav) (err error) {
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

func GetNavList(e []model.UserNav) (res []model.UserNav, err error) {
	err = global.NLY_DB.Find(&e).Error
	return e, err
}
