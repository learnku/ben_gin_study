package main

import (
	"ben_gin_study/app/controller"
	chat_controller "ben_gin_study/app/controller/chat"
	"ben_gin_study/config"
	_ "ben_gin_study/logs"
	"ben_gin_study/utils/utils_cookie"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.APP["ENV"])
	r := gin.Default()

	// 静态文件
	r.Static("/assets", "resources/assets")
	r.StaticFile("/favicon.ico", "resources/favicon.ico")
	r.LoadHTMLGlob("resources/view/*")

	// 测试 websocket 页面
	r.GET("/index.html", func(ctx *gin.Context) {
		// 设 cookie
		if exists, _ := utils_cookie.GetVisitorCookie(ctx); !exists {
			utils_cookie.SetVisitorCookie(ctx, utils_cookie.MakeCookieValue())
		}
		ctx.HTML(200, "index.html", nil)
	})

	// version = v1
	v1 := r.Group("/v1")
	{
		// websocket
		ws := v1.Group("/ws")
		{
			var chatController = new(chat_controller.ChatController)
			ws.GET("/chat", chatController.Connect)
		}

		// 用户
		user := v1.Group("/user")
		{
			var userController = new(controller.UserController)
			user.GET("/index", userController.Get)
		}
	}

	// 启动
	if err := r.Run(config.APP["PORT"]); err != nil {
		panic(fmt.Sprintf("服务启动失败：%v", err))
	}
}
