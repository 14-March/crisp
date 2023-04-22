/*
Package crisp
@Author：14March
@File：example_server.go
*/
package main

import (
	"fmt"
	"github.com/hcraM41/crisp/example/handler"
	"github.com/hcraM41/crisp/example/msg"
	"net/http"
	"os"
	"path"

	//myWebsocket "github.com/hcraM41/crisp/network/websocket"
	"github.com/hcraM41/crisp/example/ctx"
	"github.com/hcraM41/crisp/network/broadcaster"

	"github.com/gorilla/websocket"
	"github.com/hcraM41/crisp/comm/clog"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var sessionId int32 = 0

func main() {
	fmt.Println("hello crisp~")

	ex, err := os.Executable()

	if nil != err {
		panic(err)
	}

	clog.Config(path.Dir(ex) + "/log/crisp.log")
	// 注册路由
	msg.InitMsg()
	handler.InitHandle()

	http.HandleFunc("/websocket", websocketHandshake)
	_ = http.ListenAndServe("127.0.0.1:12345", nil)
}

// websocketHandshake : WebSocket 握手, 处理 WebSocket Upgrade 协议
func websocketHandshake(w http.ResponseWriter, r *http.Request) {
	if nil == w ||
		nil == r {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if nil != err {
		clog.Error("WebSocket upgrade error, %+v", err)
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	clog.Info("有新客户端连入")

	sessionId += 1

	//exampleCtx := &myWebsocket.CmdContextImpl{
	//	Conn:      conn,
	//	SessionId: sessionId,
	//}

	exampleCtx := &ctx.CmdExampleCtxImpl{
		Conn:      conn,
		SessionId: sessionId,
	}

	// 将指令上下文添加到分组,
	// 当断开连接时移除指令上下文...
	broadcaster.AddCmdCtx(sessionId, exampleCtx)
	defer broadcaster.RemoveCmdCtxBySessionId(sessionId)

	// 循环发送消息
	exampleCtx.LoopSendMsg()
	// 开始循环读取消息
	exampleCtx.LoopReadMsg()
}
