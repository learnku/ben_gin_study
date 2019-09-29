package chat

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Message struct {
	Id      int64  `json:"id,omitempty"`      // 消息ID
	Userid  int64  `json:"userid,omitempty"`  // 谁发的
	Dstid   int64  `json:"dstid,omitempty"`   // 对端用户ID
	Content string `json:"content,omitempty"` // 消息的内容
}

// 本核心在于形成 userid 和 Node 的映射关系
type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte // 并行转串行
}

// 客服 or 访客与 Node 的关系映射
var VisitorClientMap map[string]*Node = make(map[string]*Node) // 访客映射关系 map
var ServiceClientMap map[string]*Node = make(map[string]*Node) // 客服映射关系 map

// 读写锁
var rwlocker sync.RWMutex
