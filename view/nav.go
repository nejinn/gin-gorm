package view

import (
	"errors"
	"note/global"
	"note/model"
)

func CreateNav(e *model.Nav) (err error) {
	if e.Name == "" {
		return errors.New("导航菜单不能为空")
	}
	// icon 默认为 oi oi-list-rich
	if e.Icon == "" {
		e.Icon = "nav-icon fas oi oi-list-rich"
	}
	if e.Order == 0 {
		return errors.New("排序应该大于等于1")
	}

	err = global.NLY_DB.Create(&e).Error
	return err
}

func GetNavList(e []model.Nav) (res []model.Nav, err error) {
	err = global.NLY_DB.Order("order").Order("create_date").Find(&e).Error
	return e, err
}
