package cmd_handler

import (
	"github.com/hcraM41/crisp/ciface"
	"google.golang.org/protobuf/types/dynamicpb"
)

// CmdHandlerFunc 自定义消息处理函数
type CmdHandlerFunc func(ctx ciface.MyCmdContext, pbMsgObj *dynamicpb.Message)

var CmdHandlerMap = make(map[uint16]CmdHandlerFunc)

// GetCmdHandler 获取cmd处理函数
func GetCmdHandler(msgCode uint16) CmdHandlerFunc {
	return CmdHandlerMap[msgCode]
}
