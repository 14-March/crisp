package handler

import (
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"github.com/hcraM41/crisp/network/broadcaster"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"

	"time"
)

// 用户移动到指令处理器
func handleUserMoveToCmd(ctx ciface.MyCmdContext, pbMsgObj *dynamicpb.Message) {
	if nil == ctx ||
		ctx.GetUserId() <= 0 ||
		nil == pbMsgObj {
		return
	}

	// 获取用户数据
	user := user_data.GetUserGroup().GetByUserId(ctx.GetUserId())

	if nil == user {
		return
	}

	userMoveToCmd := &msg.UserMoveToCmd{}

	pbMsgObj.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		userMoveToCmd.ProtoReflect().Set(f, v)
		return true
	})

	if nil == user.MoveState {
		user.MoveState = &user_data.MoveState{}
	}

	nowTime := time.Now().UnixMilli()

	user.MoveState.FromPosX = userMoveToCmd.MoveFromPosX
	user.MoveState.FromPosY = userMoveToCmd.MoveFromPosY
	user.MoveState.ToPosX = userMoveToCmd.MoveToPosX
	user.MoveState.ToPosY = userMoveToCmd.MoveToPosY
	user.MoveState.StartTime = nowTime

	userMoveToResult := &msg.UserMoveToResult{
		MoveUserId:    uint32(ctx.GetUserId()),
		MoveFromPosX:  userMoveToCmd.MoveFromPosX,
		MoveFromPosY:  userMoveToCmd.MoveFromPosY,
		MoveToPosX:    userMoveToCmd.MoveToPosX,
		MoveToPosY:    userMoveToCmd.MoveToPosY,
		MoveStartTime: uint64(nowTime),
	}

	broadcaster.Broadcast(userMoveToResult)
}
