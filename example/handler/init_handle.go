package handler

import (
	"github.com/hcraM41/crisp/cmd_handler"
	"github.com/hcraM41/crisp/example/msg"
)

func InitHandle() {
	cmd_handler.CmdHandlerMap[uint16(msg.MsgCode_USER_ATTK_CMD.Number())] = handleUserAttkCmd
	cmd_handler.CmdHandlerMap[uint16(msg.MsgCode_USER_ENTRY_CMD.Number())] = handleUserEntryCmd
	cmd_handler.CmdHandlerMap[uint16(msg.MsgCode_USER_LOGIN_CMD.Number())] = handleUserLoginCmd
	cmd_handler.CmdHandlerMap[uint16(msg.MsgCode_USER_MOVE_TO_CMD.Number())] = handleUserMoveToCmd
	cmd_handler.CmdHandlerMap[uint16(msg.MsgCode_WHO_ELSE_IS_HERE_CMD.Number())] = handleWhoElseIsHereCmd
}
