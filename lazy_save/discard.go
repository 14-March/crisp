package lazy_save

import (
	"github.com/hcraM41/crisp/comm/clog"
)

func Discard(lso LazySaveObj) {
	if nil == lso {
		return
	}

	clog.Info("放弃延迟保存, lsoId = %+v", lso.GetLsoId())

	lsoMap.Delete(lso.GetLsoId())
}
