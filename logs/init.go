package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05", // 时间格式
		DisableTimestamp: false,                 // 禁止显示时间
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      false,
	})

	// 日志写入文件并分割日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/runtime/runtime.log", // 要写入日志的文件
		MaxSize:    10,                         // 日志文件最大允许的大小，默认 100M
		MaxBackups: 30,                         // 要保留多少个旧日志文件
		MaxAge:     30,                         // 日志文件保留多少天后删除，单位：天
		Compress:   false,                      // disabled by default
	}
	logrus.SetOutput(lumberjackLogger)

	// 设置显示报错文件和行号
	filenameHook := NewHook()
	filenameHook.Field = "line"
	logrus.AddHook(filenameHook)

	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
}

// 创建日志目录
func makeLogFile(_dir string) (*os.File, error) {
	_now := time.Now()
	y := _now.Format("2006")
	m := _now.Format("01")
	logPath := path.Join(_dir, y)
	logFile := path.Join(logPath, m+".log")
	fmt.Println(logFile)
	err := os.MkdirAll(logPath, 0666)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
}
