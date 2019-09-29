package dbserver

import (
	"ben_gin_study/app/model"
	"fmt"
	"github.com/sirupsen/logrus"
)

type User struct {
}

func (s *User) CreateTable() {
	sql := `CREATE TABLE IF NOT EXISTS user(
		id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT "用户id",
		name VARCHAR(20) NOT NULL UNIQUE COMMENT "用户名",
		email VARCHAR(32) NOT NULL UNIQUE COMMENT "邮箱",
		password VARCHAR(255) NOT NULL COMMENT "密码",
		avatar VARCHAR(255) NULL COMMENT "头像",
		sex TINYINT DEFAULT 0 COMMENT "性别",
		online TINYINT DEFAULT 0 COMMENT "在线：默认不在线",
		company_id BIGINT DEFAULT 0 COMMENT "公司id",
		notification_count INT DEFAULT 0 comment "待通知消息条目",
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
	) comment "用户表-客服"`
	_, err := DbEngin.Exec(sql)
	if err != nil {
		logrus.Error(err)
	}
}

// 插入
func (s *User) Insert(user *model.User) {
	fmt.Println(user.Email)
}