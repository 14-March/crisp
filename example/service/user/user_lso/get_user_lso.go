package user_lso

import "github.com/hcraM41/crisp/example/service/user/user_data"

func GetUserLso(user *user_data.User) *UserLso {
	if nil == user {
		return nil
	}

	existComp, _ := user.GetComponentMap().Load("UserLso") // map.get("UserLso")

	if nil != existComp {
		return existComp.(*UserLso)
	}

	existComp = &UserLso{
		User: user,
	}

	existComp, _ = user.GetComponentMap().LoadOrStore("UserLso", existComp)

	return existComp.(*UserLso)
}
