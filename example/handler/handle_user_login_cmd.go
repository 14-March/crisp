package handler

import (
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// 用户登录指令处理器
func handleUserLoginCmd(ctx ciface.MyCmdContext, pbMsgObj *dynamicpb.Message) {
	if nil == ctx ||
		nil == pbMsgObj {
		return
	}

	userLoginCmd := &msg.UserLoginCmd{}

	pbMsgObj.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		userLoginCmd.ProtoReflect().Set(f, v)
		return true
	})

	clog.Info(
		"收到用户登录消息! userName = %s, password = %s",
		userLoginCmd.GetUserName(),
		userLoginCmd.GetPassword(),
	)

	// 根据用户名称和密码登录
	bizResult := user.LoginByPasswordAsync(userLoginCmd.GetUserName(), userLoginCmd.GetPassword())

	if nil == bizResult {
		clog.Error(
			"业务结果返回空值, userName = %s",
			userLoginCmd.GetUserName(),
		)
		return
	}

	// 执行了一大堆别的操作...

	bizResult.OnComplete(func() {
		returnedObj := bizResult.GetReturnedObj()

		if nil == returnedObj {
			clog.Error(
				"用户不存在, userName = %s",
				userLoginCmd.GetUserName(),
			)
			return
		}

		resUser := returnedObj.(*user_data.User)

		userLoginResult := &msg.UserLoginResult{
			UserId:     uint32(resUser.UserId),
			UserName:   resUser.UserName,
			HeroAvatar: resUser.HeroAvatar,
		}

		ctx.BindUserId(resUser.UserId)
		ctx.Write(userLoginResult)
	})
}
