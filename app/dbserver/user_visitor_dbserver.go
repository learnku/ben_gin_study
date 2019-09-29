package dbserver

import "github.com/sirupsen/logrus"

type UserVisitor struct {

}

func (s *UserVisitor) CreateTable() {
	sql := `CREATE TABLE IF NOT EXISTS user_visitor(
		id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT "用户id",
		cookie_name VARCHAR(255) NOT NULL COMMENT "访客唯一标识",
		ip" VARCHAR(32) NOT NULL COMMENT "访客ip",
		extra JSON COMMENT "冗余字段",
		notification_count INT DEFAULT 0 COMMENT "待通知消息条目",
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	) comment "访客表"`
	_, err := DbEngin.Exec(sql)
	if err != nil {
		logrus.Error(err)
	}
}