/*
Package crisp
@Author：14March
@File：server.go
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

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

func main() {
	fmt.Println("hello crisp~")

	ex, err := os.Executable()

	if nil != err {
		panic(err)
	}

	clog.Config(path.Dir(ex) + "/log/crisp.log")

	http.HandleFunc("/websocket", websocketHandshake)
	_ = http.ListenAndServe("127.0.0.1:12345", nil)
}

// websocketHandshake : WebSocket 握手, 处理 WebSocket Upgrade 协议
func websocketHandshake(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if nil != err {
		clog.Error("WebSocket Upgrade Error, %v", err)
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	clog.Info("新客户端连入")

	for {
		_, msgData, err := conn.ReadMessage()

		if nil != err {
			clog.Error("%v", err)
			break
		}

		clog.Info("%v", msgData)
	}
}
