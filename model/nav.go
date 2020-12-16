package model

import "time"

type UserNav struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"column:name"`                        // 导航名称
	Icon        string    `json:"icon" gorm:"column:icon;default:null"`           // 导航icon
	ParentId    int       `json:"parent_id" gorm:"column:parent_id;default:null"` // 父级菜单id
	Router      string    `json:"router" gorm:"column:router;default:null"`       // 路由
	Order       int       `json:"order" gorm:"column:order;default:null"`         // 排序
	Type        int       `json:"type" gorm:"column:type;default:1"`              // 类型 1 nly-sidebar-nav-item 2 nly-sidebar-nav-tree 3 nly-sidebar-nav-header
	CreatedAt   time.Time `json:"create_date" gorm:"column:create_date"`
	UpdatedAt   time.Time `json:"update_date" gorm:"column:update_date"`
	DeletedTime time.Time `json:"deleted_date" gorm:"column:deleted_date;default:null"`
	IsDelete    bool      `json:"is_delete" gorm:"column:is_delete;default:false"`
	CreateId    int       `json:"create_id" gorm:"column:create_id;default:null"`
	UpdateId    int       `json:"update_id" gorm:"column:update_id;default:null"`
}
