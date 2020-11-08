package view

import (
	"errors"
	"github.com/jinzhu/gorm"
	"note/global"
	"note/model"
	"note/utils"
)

type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//用户注册
func Register(r *RegisterStruct)(err error) {
	x := model.User{}
	global.NLY_DB.Where("username = ?", r.Username).First(&x)
	if !errors.Is(global.NLY_DB.Where("username = ?", r.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound){
		return errors.New("用户名已经注册")
	}
	password := utils.NewMD5([]byte(r.Password))
	user := model.User{
		Username:    r.Username,
		Password:    password,
	}
	err = global.NLY_DB.Create(&user).Error
	return err
}

// 登录校验账号密码 （admin）