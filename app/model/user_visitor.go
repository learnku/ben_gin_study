package model

import "time"

// 访客表
type UserVisitor struct {
	Id                int64                  `json:"id"`
	CookieName        string                 `json:"cookie_name"`
	Ip                string                 `json:"ip"`
	Extra             map[string]interface{} `json:"extra"`
	NotificationCount int                    `json:"notification_count"`
	CreatedAt         time.Time              `xorm:"created" json:"created_at"`
}
