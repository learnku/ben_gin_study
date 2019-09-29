package chat

import (
	"ben_gin_study/app/controller"
	"ben_gin_study/utils/utils_cookie"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ChatController struct {
	controller.BaseController
}

func (this *ChatController) Connect(ctx *gin.Context) {
	isvalida := false
	// 检测接入是否合法
	exists, visitorCookie := utils_cookie.GetVisitorCookie(ctx)
	if exists {
		isvalida = true
	}

	// 握手（websocket）
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 获取 conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
	}

	// 形成与 Node 的绑定关系
	rwlocker.Lock()
	if visitorCookie == "vip" {
		ServiceClientMap[visitorCookie] = node
	} else {
		// 形成访客与Node的绑定关系
		VisitorClientMap[visitorCookie] = node
	}
	rwlocker.Unlock()

	// 向服务端发送消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("sdsdsd"))
	fmt.Println(err)
}
