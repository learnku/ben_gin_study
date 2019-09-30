package model

import (
	"time"
)

// 用户表【客服表】
type User struct {
	Id                int64     `json:"id" validate:"required"`
	Name              string    `json:"name"`
	Email             string    `json:"email" validate:"required,email" vname:"邮箱"`
	Password          string    `json:"password"`
	Avatar            string    `json:"avatar"`
	Sex               int       `json:"sex"`
	Online            int       `json:"online"`
	CompanyId         int       `json:"company_id"`
	NotificationCount int       `json:"notification_count"`
	CreatedAt         time.Time `xorm:"created" json:"created_at"`
	UpdatedAt         time.Time `xorm:"updated" json:"updated_at"`
}
