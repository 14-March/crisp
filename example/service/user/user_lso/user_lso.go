package user_lso

import (
	"fmt"
	"github.com/hcraM41/crisp/async_op"
	"github.com/hcraM41/crisp/example/service/user/user_dao"
	"github.com/hcraM41/crisp/example/service/user/user_data"
)

type UserLso struct {
	*user_data.User
}

func (lso *UserLso) GetLsoId() string {
	return fmt.Sprintf("UserLso_%d", lso.UserId)
}

func (lso *UserLso) SaveOrUpdate(callback func()) {
	async_op.Process(
		int(lso.UserId),
		func() {
			user_dao.SaveOrUpdate(lso.User)

			if nil != callback {
				callback()
			}
		},
	)
}
