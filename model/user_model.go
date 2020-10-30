package model

import (
	"github.com/jinzhu/gorm"
)

//type UserGenderType bool

const (
	UserGenderMan   bool = false
	UserGenderWomen bool = true
	UserIsDelete    bool = false
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;comment:用户名"`
	Password string `gorm:"column:password;comment:密码"`
	Gender   *bool  `gorm:"column:gender;comment:性别;default:false"`
	IsDelete *bool  `gorm:"column:gender;comment:已删除;default:false"`
	IsSuper  *bool  `gorm:"column:gender;comment:是否超级管理员;default:false"`
}
