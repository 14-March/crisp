package async_op

import (
	"github.com/hcraM41/crisp/comm/clog"
)

// 理解为其中的一个线程 + 队列
type worker struct {
	taskQ chan func() // LinkedBlockingQueue<Function> taskQ
}

// 处理异步过程,
// XXX 注意: 这里只是将异步操作加入到队列里, 并不立即执行...
func (w *worker) process(asyncOp func()) {
	// 这个 w *worker 就相当于 this

	if nil == asyncOp {
		clog.Error("异步操作为空")
		return
	}

	if nil == w.taskQ {
		clog.Error("任务队列尚未初始化")
		return
	}

	w.taskQ <- func() { // taskQ.offer(new Function() { ... })
		// 执行异步操作
		asyncOp()

	}
}

// 循环执行任务
func (w *worker) loopExecTask() {
	if nil == w.taskQ {
		clog.Error("任务队列尚未初始化")
		return
	}

	for {
		task := <-w.taskQ

		if nil == task {
			continue
		}

		func() {
			defer func() {
				if err := recover(); nil != err {
					clog.Error("发生异常, %+v", err)
				}
			}()

			// 执行任务
			task()
		}()
	}
}
