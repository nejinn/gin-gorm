package model

import (
	"time"
)

// 构造 用户 model
type User struct {
	ID                int       `json:"id" gorm:"primary_key"`
	Username          string    `json:"username" gorm:"column:username"`
	Password          string    `json:"password" gorm:"column:password"`
	UserPic           string    `json:"user_pic" gorm:"column:user_pic"`
	UserType          bool      `json:"user_type" gorm:"column:user_type;default:false"`
	UserPhone         string    `json:"user_phone" gorm:"column:user_phone;default:null"`
	UserEmail         string    `json:"user_email" gorm:"column:user_email;default:null"`
	UserExp           int       `json:"user_exp" gorm:"column:user_exp;default:0"`
	UserLevel         int       `json:"user_level" gorm:"column:user_level;default:1"`
	UserBirthday      string    `json:"user_birthday" gorm:"column:user_birthday;default:null"`
	UserGender        int       `json:"user_gender" gorm:"column:user_gender;default:1"` // 1 男 2 女
	UserIp            string    `json:"user_ip" gorm:"column:user_ip;default:null"`
	UserLastLoginIp   string    `json:"user_last_login_ip" gorm:"column:user_last_login_ip;default:null"`
	UserLastLoginTime time.Time `json:"user_last_login_time" gorm:"column:user_last_login_time;default:null"`
	CreatedAt         time.Time `json:"create_date" gorm:"column:create_date"`
	UpdatedAt         time.Time `json:"update_date" gorm:"column:update_date"`
	DeletedTime       time.Time `json:"deleted_date" gorm:"column:deleted_date;default:null"`
	IsDelete          bool      `json:"is_delete" gorm:"column:is_delete;default:false"`
	CreateId          int       `json:"create_id" gorm:"column:create_id;default:null"`
	UpdateId          int       `json:"update_id" gorm:"column:update_id;default:null"`
}
