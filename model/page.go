package model

import "time"

// 访问页面
type VisitPage struct {
	ID        int        `json:"id" gorm:"primary_key"`
	UserId    int        `json:"user_id"`
	UrlId     int        `json:"url_id"`
	CreatedAt *time.Time `json:"create_date" gorm:"column:create_date"`
	VisitIp   string     `json:"visit_ip"`
}

// 访问接口
type VisitApi struct {
	ID        int        `json:"id" gorm:"primary_key"`
	UserId    int        `json:"user_id"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"create_date" gorm:"column:create_date"`
	VisitIp   string     `json:"visit_ip"`
}
