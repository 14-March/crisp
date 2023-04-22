package handler

import (
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"github.com/hcraM41/crisp/network/broadcaster"
	"google.golang.org/protobuf/types/dynamicpb"
)

// 用户入场指令处理器
func handleUserEntryCmd(ctx ciface.MyCmdContext, _ *dynamicpb.Message) {
	if nil == ctx ||
		ctx.GetUserId() <= 0 {
		return
	}

	clog.Info(
		"收到用户入场消息! userId = %d",
		ctx.GetUserId(),
	)

	// 获取用户数据
	user := user_data.GetUserGroup().GetByUserId(ctx.GetUserId())

	if nil == user {
		clog.Error(
			"未找到用户数据, userId = %d",
			ctx.GetUserId(),
		)
		return
	}

	userEntryResult := &msg.UserEntryResult{
		UserId:     uint32(ctx.GetUserId()),
		UserName:   user.UserName,
		HeroAvatar: user.HeroAvatar,
	}

	broadcaster.Broadcast(userEntryResult)
}
