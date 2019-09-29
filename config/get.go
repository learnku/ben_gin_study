package config

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

// 获取：[VISITOR]CookieMaxAge
func GetVisitor_CookieMaxAge() int {
	i, err := strconv.Atoi(VISITOR["CookieMaxAge"])
	if err != nil {
		logrus.Error(err)
	}
	return i
}