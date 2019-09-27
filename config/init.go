package config

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var (
	APP   map[string]string // 应用程序
	MYSQL map[string]string // mysql数据库
)

func init() {
	config, err := goconfig.LoadConfigFile("config/env.ini")
	if err != nil {
		log.Fatal("Error loading 'config/env.ini' file : ", err.Error())
		return
	}

	app(config)   // 基础配置
	mysql(config) // Mysql
}

// 基础配置
func app(config *goconfig.ConfigFile) {
	APP = make(map[string]string, 4)
	APP["NAME"], _ = config.GetValue("APP", "NAME") // 应用名称
	APP["ENV"], _ = config.GetValue("APP", "ENV")   // 运行模式
	APP["PORT"], _ = config.GetValue("APP", "PORT") // 程序运行端口号
	APP["URL"], _ = config.GetValue("APP", "URL")   // 域名
}

// Mysql
func mysql(config *goconfig.ConfigFile) {
	MYSQL = make(map[string]string, 5)
	MYSQL["HOST"], _ = config.GetValue("MYSQL", "HOST")
	MYSQL["PORT"], _ = config.GetValue("MYSQL", "PORT")
	MYSQL["DATABASE"], _ = config.GetValue("MYSQL", "DATABASE")
	MYSQL["USERNAME"], _ = config.GetValue("MYSQL", "USERNAME")
	MYSQL["PASSWORD"], _ = config.GetValue("MYSQL", "PASSWORD")
}
