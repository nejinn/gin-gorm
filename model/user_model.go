package model

import (
	"github.com/jinzhu/gorm"
)

// 构造 用户 model
type User struct {
	gorm.Model
	Username string `gorm:"column:'username'"`
	Password string `gorm:"column:'password'"`
	Gender   bool   `gorm:"column:'gender';default:false"`
	IsDelete bool   `gorm:"column:'Is_delete';default:false"`
	IsSuper  bool   `gorm:"column:'is_super';default:false"`
}
