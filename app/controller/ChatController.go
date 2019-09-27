package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type ChatController struct {
	BaseController
}

func (this *ChatController) Connect(ctx *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func (r *http.Request) bool {
			return true
		},
	}

	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		// 读取 ws 中的数据
		mt, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping"{
			message = []byte("pong")
		}
		// 写入 ws 数据
		err = conn.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}