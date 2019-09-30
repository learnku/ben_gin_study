package dbserver

import (
	"github.com/sirupsen/logrus"
)

type ChatMessage struct {
}

var UserRole = map[string]int{
	"admin":   0, // 管理员
	"serve":   1, //
	"visitor": 2, // 访客
}

func (s *ChatMessage) CreateTable() {
	sql := `CREATE TABLE IF NOT EXISTS chat_message(
		id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT "消息id",
		user_id BIGINT NOT NULL COMMENT "发送人的消息 id",
		user_role tinyint NOT NULL DEFAULT 0 COMMENT "发送人的角色",
		accept_user_id BIGINT NOT NULL COMMENT "接收人的消息 id",
		accept_user_role tinyint NOT NULL DEFAULT 0 COMMENT "接收人的角色",
		cookie_name VARCHAR(255) NULL DEFAULT "" COMMENT "访客唯一标识",
		content VARCHAR(500) NULL DEFAULT "" COMMENT "消息内容",
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	) COMMENT "图书信息表"`
	_, err := DbEngin.Exec(sql)
	if err != nil {
		logrus.Error(err)
	}
}
