package user_dao

import (
	"github.com/hcraM41/crisp/comm/clog"
	"github.com/hcraM41/crisp/example/service/user/user_data"
)

const sqlGetUserByName = `select user_id, user_name, password, hero_avatar, curr_hp from t_user where user_name = ?`

// GetUserByName 根据用户名获得用户数据
//func GetUserByName(userName string) *user_data.User {
//	if len(userName) <= 0 {
//		return nil
//	}
//
//	row := mysql.DB.QueryRow(sqlGetUserByName, userName)
//
//	if nil == row {
//		return nil
//	}
//
//	user := &user_data.User{}
//
//	err := row.Scan(
//		&user.UserId,
//		&user.UserName,
//		&user.Password,
//		&user.HeroAvatar,
//		&user.CurrHp,
//	)
//
//	if nil != err {
//		clog.Error("%+v", err)
//		return nil
//	}
//
//	return user
//}

func GetUserByName(userName string) *user_data.User {
	clog.Info("GetUserByName测试ing，每次均为重建数据...", userName)
	return nil
}
