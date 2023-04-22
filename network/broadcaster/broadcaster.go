package broadcaster

import (
	"github.com/hcraM41/crisp/ciface"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var innerMap = make(map[int32]ciface.MyCmdContext)

// AddCmdCtx 添加指令上下文分组
func AddCmdCtx(sessionId int32, ctx ciface.MyCmdContext) {
	if nil == ctx {
		return
	}

	innerMap[sessionId] = ctx
}

// RemoveCmdCtxBySessionId 移除指令上下文分组
func RemoveCmdCtxBySessionId(sessionId int32) {
	if sessionId <= 0 {
		return
	}

	delete(innerMap, sessionId)
}

// Broadcast 广播消息
func Broadcast(msgObj protoreflect.ProtoMessage) {
	if nil == msgObj {
		return
	}

	for _, ctx := range innerMap {
		if nil != ctx {
			ctx.Write(msgObj)
		}
	}
}
