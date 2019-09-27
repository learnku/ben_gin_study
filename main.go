package main

import (
	"ben_gin_study/app/controller"
	"ben_gin_study/config"
	_ "ben_gin_study/logs"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	gin.SetMode(config.APP["ENV"])
	r := gin.Default()

	// 静态文件
	r.Static("/assets", "resources/assets")
	r.StaticFile("/favicon.ico", "resources/favicon.ico")
	r.LoadHTMLGlob("resources/view/*")

	// 测试 websocket 页面
	r.GET("/index.html", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	var chatController = new(controller.ChatController)
	r.GET("/ws", chatController.Connect)

	// 启动
	if err := r.Run(config.APP["PORT"]); err != nil {
		panic(fmt.Sprintf("服务启动失败：%v", err))
	}
}