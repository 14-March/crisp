package ctx

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"github.com/hcraM41/crisp/cmd_handler"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/handler"
	"github.com/hcraM41/crisp/main_thread"
	"github.com/hcraM41/crisp/message"
	"google.golang.org/protobuf/reflect/protoreflect"
	"time"
)

const oneSecond = 1000
const readMsgCountPerSecond = 16

// CmdExampleCtxImpl 就是 MyCmdContext 的 WebSocket 实现
type CmdExampleCtxImpl struct {
	userId       int64
	clientIpAddr string
	Conn         *websocket.Conn
	sendMsgQ     chan protoreflect.ProtoMessage // BlockingQueue
	SessionId    int32
}

func (ctx *CmdExampleCtxImpl) BindUserId(val int64) {
	ctx.userId = val
}

func (ctx *CmdExampleCtxImpl) GetUserId() int64 {
	return ctx.userId
}

func (ctx *CmdExampleCtxImpl) GetClientIpAddr() string {
	return ctx.clientIpAddr
}

func (ctx *CmdExampleCtxImpl) Write(msgObj protoreflect.ProtoMessage) {
	if nil == msgObj ||
		nil == ctx.Conn ||
		nil == ctx.sendMsgQ {
		return
	}

	ctx.sendMsgQ <- msgObj // queue.push
}

func (ctx *CmdExampleCtxImpl) SendError(errorCode int, errorInfo string) {
}

func (ctx *CmdExampleCtxImpl) Disconnect() {
	if nil != ctx.Conn {
		clog.Info("Disconnect ===>", ctx.userId)
		// DIY用户离线处理
		handler.OnUserQuit(ctx)

		_ = ctx.Conn.Close()
	}
}

// LoopSendMsg 循环发送消息
func (ctx *CmdExampleCtxImpl) LoopSendMsg() {
	clog.Info("LoopSendMsg ===>")
	// 首先构建发送队列
	ctx.sendMsgQ = make(chan protoreflect.ProtoMessage, 64)

	go func() { // new Thread().start(() -> { ... })
		for {
			msgObj := <-ctx.sendMsgQ // queue.pop

			if nil == msgObj {
				continue
			}

			func() {
				defer func() {
					if err := recover(); nil != err {
						clog.Error("发生异常, %+v", err)
					}
				}()

				byteArray, err := message.Encode(msgObj)

				if nil != err {
					clog.Error("%+v", err)
					return
				}

				if err := ctx.Conn.WriteMessage(websocket.BinaryMessage, byteArray); nil != err {
					clog.Error("%+v", err)
				}
			}()
		}
	}()
}

// LoopReadMsg 循环读取消息
func (ctx *CmdExampleCtxImpl) LoopReadMsg() {
	clog.Info("LoopReadMsg ===>")
	if nil == ctx.Conn {
		return
	}

	// 设置读取字节数限制
	ctx.Conn.SetReadLimit(64 * 1024)

	t0 := int64(0)
	counter := 0

	for {
		_, msgData, err := ctx.Conn.ReadMessage()

		if nil != err {
			// XXX 注意: 用户断线的时候也会触发一个异常,
			// 遇到这个异常直接停止循环即可...
			clog.Error("%+v", err)
			break
		}

		t1 := time.Now().UnixMilli()

		if (t1 - t0) > oneSecond {
			t0 = t1
			counter = 0
		}

		if counter >= readMsgCountPerSecond {
			clog.Error("消息过于频繁")
			continue
		}

		counter++

		func() {
			defer func() {
				if err := recover(); nil != err {
					clog.Error("发生异常, %+v", err)
				}
			}()

			msgCode := binary.BigEndian.Uint16(msgData[2:4])
			newMsgX, err := message.Decode(msgData[4:], int16(msgCode))

			if nil != err {
				clog.Error(
					"消息解码错误, msgCode = %d, error = %+v",
					msgCode, err,
				)
				return
			}

			clog.Info(
				"收到客户端消息, msgCode = %d, msgName = %s",
				msgCode,
				newMsgX.Descriptor().Name(),
			)

			// 创建指令处理器
			cmdHandler := cmd_handler.GetCmdHandler(msgCode)

			if nil == cmdHandler {
				clog.Error(
					"未找到指令处理器, msgCode = %d",
					msgCode,
				)
				return
			}

			main_thread.Process(func() {
				cmdHandler(ctx, newMsgX)
			})
		}()
	}

	// 处理用户离线逻辑
	ctx.Disconnect()
}
