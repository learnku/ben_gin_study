package dbserver

import (
	"ben_gin_study/app/model"
	. "ben_gin_study/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/
// https://github.com/go-xorm/xorm/blob/d0827bfc00eac71059fef291939d5b63aff42d5d/README_CN.md
// http://gobook.io/read/github.com/go-xorm/manual-zh-CN/

var (
	DbEngin *xorm.Engine
)

// Mysql 数据库初始化
func init() {
	drivename := "mysql"
	dsName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", MYSQL["USERNAME"], MYSQL["PASSWORD"], MYSQL["HOST"], MYSQL["PORT"], MYSQL["DATABASE"])
	var err error
	DbEngin, err = xorm.NewEngine(drivename, dsName)
	if err != nil {
		log.Fatal(err)
	}

	// 记录日志
	logFile, err := os.Create("logs/db/mysql.log")
	if err != nil {
		logrus.Error(err)
	} else {
		DbEngin.SetLogger(xorm.NewSimpleLogger(logFile))
	}

	// 显示 Sql 语句
	DbEngin.ShowSQL(true)

	// 数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(25)

	// 自动建表
	autoRegisterTablesSql()
}

// 自动建表
func autoRegisterTablesSql() {
	new(User).CreateTable()        // 用户表
	new(UserVisitor).CreateTable() // 访客用户表
	new(ChatMessage).CreateTable() // 对话消息表
}

// 自动建表
func autoRegisterTables() {
	err := DbEngin.Sync2(
		new(model.User),
	)
	if err != nil {
		logrus.Error(err.Error())
	}
}
