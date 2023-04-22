package user

import (
	"fmt"
	"github.com/hcraM41/crisp/async_op"
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/service/user/user_dao"
	"github.com/hcraM41/crisp/example/service/user/user_data"
	"github.com/hcraM41/crisp/example/service/user/user_lock"
	"time"
)

var TestUserId = 0

// LoginByPasswordAsync 根据用户名称和密码进行登录,
// 将返回一个异步的业务结果
func LoginByPasswordAsync(userName string, password string) *async_op.AsyncBizResult {
	// 要说下面这两种写法有什么不同么?
	// func LoginByPasswordAsync(userName string, password string, callback func(user *user_data.User)) { ... }
	// func LoginByPasswordAsync(userName string, password string) *base.AsyncBizResult { ... }
	// 第一个是回调方式的, 第二个是返回 Future 方式的,
	// 这两个有什么不一样么?
	// 可以参考其他语言中的 async / await 相关知识...
	//
	if len(userName) <= 0 ||
		len(password) <= 0 {
		return nil
	}

	bizResult := &async_op.AsyncBizResult{}

	async_op.Process(
		StrToBindId(userName),
		func() {
			// 通过 DAO 获得用户数据
			user := user_dao.GetUserByName(userName)

			nowTime := time.Now().UnixMilli()

			if nil == user {
				// 如果用户数据为空，
				//
				TestUserId++
				user = &user_data.User{
					UserId:     int64(TestUserId),
					UserName:   userName,
					Password:   password,
					CreateTime: nowTime,
					HeroAvatar: "Hero_Hammer",
				}
			}

			clog.Info("LoginByPasswordAsync===>", user)
			//
			// 是否有登出锁,
			// 如果有锁,
			// 那就直接退出吧...
			//
			key := fmt.Sprintf("UserQuit_%d", user.UserId)
			if user_lock.HasLock(key) {
				clog.Info("有登出锁！！！！！！！！！", key)
				bizResult.SetReturnedObj(nil)
				return
			}

			// 更新最后登录时间
			user.LastLoginTime = nowTime
			user_dao.SaveOrUpdate(user)

			// 将用户添加到字典
			user_data.GetUserGroup().Add(user)

			bizResult.SetReturnedObj(user)
		},
	)

	return bizResult
}
