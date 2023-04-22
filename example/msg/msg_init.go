package msg

import (
	"github.com/hcraM41/crisp/message"
	"strings"
)

func InitMsg() {
	message.MsgLocker.Lock()
	defer message.MsgLocker.Unlock()

	if len(message.MsgCodeAndMsgDescMap) > 0 &&
		len(message.MsgNameAndMsgCodeMap) > 0 {
		return
	}

	// 先往 msgNameAndMsgCodeMap "名称 --> 代号" 这个字典里填数据

	for k, v := range MsgCode_value {
		// USER_LOGIN_CMD ==> userlogincmd
		msgName := strings.ToLower(
			strings.Replace(k, "_", "", -1),
		)

		message.MsgNameAndMsgCodeMap[msgName] = int16(v)
	}

	msgDescList := File_GameMsgProtocol_proto.Messages()

	for i := 0; i < msgDescList.Len(); i++ {
		msgDesc := msgDescList.Get(i)
		msgName := strings.ToLower(
			strings.Replace(string(msgDesc.Name()), "_", "", -1),
		)

		msgCode := message.MsgNameAndMsgCodeMap[msgName]
		message.MsgCodeAndMsgDescMap[msgCode] = msgDesc
	}
}
