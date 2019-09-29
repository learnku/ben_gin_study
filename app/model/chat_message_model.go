package model

import "time"

type Message struct {
	Id             int64     `json:"id"`
	UserId         int64     `json:"user_id"`
	UserRole       int       `json:"user_role"`
	AcceptUserId   int64     `json:"accept_user_id"`
	AcceptUserRole int       `json:"accept_user_role"`
	CookieName     string    `json:"cookie_name"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}
