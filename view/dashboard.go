package view

import (
	"note/global"
	"note/model"
	"time"
)

type V struct {
	Num        int    `json:"num"`
	GrowthRate int    `json:"growth_rate"`
	Name       string `json:"name"`
	Day        string `json:"day"`
	Variant    string `json:"variant"`
	Icon       string `json:"icon"`
}

func InfoBox() (err error, res interface{}) {
	var vp []model.VisitPage
	var vpCount int
	var vpOldCount int
	var va []model.VisitApi
	var vaCount int
	var vaOldCount int
	var vu []model.VisitUser
	var vuCount int
	var vuOldCount int
	var vd []V
	var vd1 []V
	var user []model.User
	var userCount int
	var userOldCount int
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -1)
	global.NLY_DB.Where("create_date > ?", oldTime).Find(&vp).Count(&vpCount)
	global.NLY_DB.Where("create_date > ?", oldTime).Find(&va).Count(&vaCount)
	global.NLY_DB.Where("create_date > ?", oldTime).Find(&vu).Count(&vuCount)
	global.NLY_DB.Where("create_date > ?", oldTime).Find(&user).Count(&userCount)
	global.NLY_DB.Where("create_date <= ?", oldTime).Find(&vp).Count(&vpOldCount)
	global.NLY_DB.Where("create_date <= ?", oldTime).Find(&va).Count(&vaOldCount)
	global.NLY_DB.Where("create_date <= ?", oldTime).Find(&vu).Count(&vuOldCount)
	global.NLY_DB.Where("create_date <= ?", oldTime).Find(&user).Count(&userOldCount)
	vd = append(vd, V{
		Name:       "总访问量",
		Num:        vuCount + vuOldCount,
		Day:        "3",
		GrowthRate: vuCount * 100 / (vuOldCount + vuCount),
		Variant:    "success",
		Icon:       "far oi oi-home",
	})
	vd = append(vd, V{
		Name:       "页面总访问量",
		Num:        vpCount + vpOldCount,
		Day:        "3",
		GrowthRate: vpCount * 100 / (vpOldCount + vpCount),
		Variant:    "info",
		Icon:       "far oi oi-image",
	})
	vd = append(vd, V{
		Name:       "接口总访问量",
		Num:        vaCount + vaOldCount,
		Day:        "3",
		GrowthRate: vaCount * 100 / (vaOldCount + vaCount),
		Variant:    "danger",
		Icon:       "far oi oi-compass",
	})
	vd = append(vd, V{
		Name:       "总注册用户",
		Num:        userCount + userOldCount,
		Day:        "3",
		GrowthRate: userCount * 100 / (userOldCount + userCount),
		Variant:    "primary",
		Icon:       "far oi oi-people",
	})

	vd1 = append(vd1, V{
		Name:    "今天访问量",
		Num:     vuCount,
		Variant: "success",
		Icon:    "far oi oi-home",
	})
	vd1 = append(vd1, V{
		Name:    "今天页面访问量",
		Num:     vpCount,
		Variant: "info",
		Icon:    "far oi oi-image",
	})
	vd1 = append(vd1, V{
		Name:    "今天接口访问量",
		Num:     vaCount,
		Variant: "danger",
		Icon:    "far oi oi-compass",
	})
	vd1 = append(vd1, V{
		Name:    "今天注册用户",
		Num:     userCount,
		Variant: "primary",
		Icon:    "far oi oi-people",
	})
	type Data struct {
		Vd  []V `json:"vd"`
		Vd1 []V `json:"vd1"`
	}
	data := Data{
		Vd:  vd,
		Vd1: vd1,
	}
	return nil, data
}
