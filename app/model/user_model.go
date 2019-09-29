package model

import "time"

// 用户表【客服表】
type User struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Avatar            string    `json:"avatar"`
	Sex               int       `json:"sex"`
	Online            int       `json:"online"`
	CompanyId         int       `json:"company_id"`
	NotificationCount int       `json:"notification_count"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
