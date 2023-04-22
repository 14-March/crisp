package handler

import (
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"github.com/hcraM41/crisp/example/service/user/user_lso"
	"github.com/hcraM41/crisp/lazy_save"
	"github.com/hcraM41/crisp/network/broadcaster"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// 用户攻击指令处理器
func handleUserAttkCmd(ctx ciface.MyCmdContext, pbMsgObj *dynamicpb.Message) {
	if nil == ctx ||
		nil == pbMsgObj {
		return
	}

	userAttkCmd := &msg.UserAttkCmd{}

	pbMsgObj.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		userAttkCmd.ProtoReflect().Set(f, v)
		return true
	})

	userAttkResult := &msg.UserAttkResult{
		AttkUserId:   uint32(ctx.GetUserId()),
		TargetUserId: userAttkCmd.TargetUserId,
	}

	broadcaster.Broadcast(userAttkResult)

	user := user_data.GetUserGroup().GetByUserId(int64(userAttkCmd.GetTargetUserId()))

	if nil == user {
		return
	}

	var subtractHp int32 = 10
	user.CurrHp -= subtractHp

	userSubtractHpResult := &msg.UserSubtractHpResult{
		SubtractHp:   uint32(subtractHp),
		TargetUserId: userAttkCmd.TargetUserId,
	}

	broadcaster.Broadcast(userSubtractHpResult)

	lso := user_lso.GetUserLso(user)

	// 执行延迟保存
	lazy_save.SaveOrUpdate(lso)
}
