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
func Register(r *model.User) (err error) {
	var registerUser model.User
	global.NLY_DB.Where("username = ?", r.Username).First(&registerUser)
	if !errors.Is(global.NLY_DB.Where("username = ?", r.Username).First(&registerUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经注册")
	}
	password := utils.NewMD5([]byte(r.Password))
	avatar, avatarErr := utils.CreateDefaultAvatar(r.Username)
	if avatarErr != nil {
		return errors.New("生成头像失败")
	}
	user := model.User{
		Username: r.Username,
		Password: password,
		UserPic:  avatar,
	}
	err = global.NLY_DB.Create(&user).Error
	return err
}

// 登录校验账号密码 （admin）
func Login(r *model.User) (err error, userInfo *model.User) {
	var loginUser model.User
	if r.Username == "" {
		return errors.New("用户名不能为空"), nil
	}
	if r.Password == "" {
		return errors.New("密码不能为空"), nil
	}
	r.Password = utils.NewMD5([]byte(r.Password))

	global.NLY_DB.Where("username = ? AND password = ?", r.Username, r.Password).First(&loginUser)

	if errors.Is(global.NLY_DB.Where("username = ? AND password = ?", r.Username, r.Password).First(&loginUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("登录信息错误"), nil
	}
	return nil, &loginUser
}
