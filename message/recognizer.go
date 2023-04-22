package message

import (
	"errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
	"sync"
)

var MsgCodeAndMsgDescMap = make(map[int16]protoreflect.MessageDescriptor)
var MsgNameAndMsgCodeMap = make(map[string]int16)
var MsgLocker = &sync.Mutex{}

func getMsgDescByMsgCode(msgCode int16) (protoreflect.MessageDescriptor, error) {
	if msgCode < 0 {
		return nil, errors.New("消息代号无效")
	}

	if len(MsgCodeAndMsgDescMap) <= 0 {
		return nil, errors.New("msgCodeAndMsgDescMap 未初始化")
	}

	return MsgCodeAndMsgDescMap[msgCode], nil
}

func getMsgCodeByMsgName(msgName string) (int16, error) {
	if len(msgName) <= 0 {
		return -1, errors.New("消息名称为空")
	}

	if len(MsgNameAndMsgCodeMap) <= 0 {
		return 0, errors.New("msgNameAndMsgCodeMap 未初始化")
	}

	msgName = strings.ToLower(
		strings.Replace(msgName, "_", "", -1),
	)

	return MsgNameAndMsgCodeMap[msgName], nil
}
