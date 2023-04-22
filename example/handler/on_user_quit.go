package handler

import (
	"fmt"
	"github.com/hcraM41/crisp/ciface"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/msg"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"github.com/hcraM41/crisp/example/service/user/user_lock"
	"github.com/hcraM41/crisp/example/service/user/user_lso"
	"github.com/hcraM41/crisp/lazy_save"
	"github.com/hcraM41/crisp/network/broadcaster"
)

// OnUserQuit 当用户退出游戏时执行
func OnUserQuit(ctx ciface.MyCmdContext) {
	if nil == ctx {
		return
	}

	clog.Info("用户离线, userId = %d", ctx.GetUserId())

	//
	// 加锁
	//
	key := fmt.Sprintf("UserQuit_%d", ctx.GetUserId())
	user_lock.TryLock(key)

	broadcaster.Broadcast(&msg.UserQuitResult{
		QuitUserId: uint32(ctx.GetUserId()),
	})

	user := user_data.GetUserGroup().GetByUserId(ctx.GetUserId())

	if nil == user {
		return
	}

	user_data.GetUserGroup().RemoveByUserId(user.UserId)

	userLso := user_lso.GetUserLso(user)
	lazy_save.Discard(userLso)

	clog.Info("用户离线, 立即保存数据! userId = %d", ctx.GetUserId())

	userLso.SaveOrUpdate(func() {
		//
		// 解锁
		//
		clog.Info("延迟保存对象解锁：", key)
		user_lock.UnLock(key)
	})
}
