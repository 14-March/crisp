package handler

import (
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"google.golang.org/protobuf/types/dynamicpb"
)

// 还有谁指令处理器
func handleWhoElseIsHereCmd(ctx ciface.MyCmdContext, _ *dynamicpb.Message) {
	if nil == ctx ||
		ctx.GetUserId() <= 0 {
		return
	}

	clog.Info(
		"收到“还有谁”消息! userId = %d",
		ctx.GetUserId(),
	)

	whoElseIsHereResult := &msg.WhoElseIsHereResult{}

	// 获得所有用户
	userALL := user_data.GetUserGroup().GetUserALL()

	for _, user := range userALL {
		if nil == user {
			continue
		}

		userInfo := &msg.WhoElseIsHereResult_UserInfo{
			UserId:     uint32(user.UserId),
			UserName:   user.UserName,
			HeroAvatar: user.HeroAvatar,
		}

		if nil != user.MoveState {
			// 将数据中的移动状体 同步到 消息上的移动状态
			userInfo.MoveState = &msg.WhoElseIsHereResult_UserInfo_MoveState{
				FromPosX:  user.MoveState.FromPosX,
				FromPosY:  user.MoveState.FromPosY,
				ToPosX:    user.MoveState.ToPosX,
				ToPosY:    user.MoveState.ToPosY,
				StartTime: uint64(user.MoveState.StartTime),
			}
		}

		whoElseIsHereResult.UserInfo = append( // List<UserInfo> userInfoList; userInfoList.add(userInfo);
			whoElseIsHereResult.UserInfo,
			userInfo,
		)
	}

	ctx.Write(whoElseIsHereResult)
}
