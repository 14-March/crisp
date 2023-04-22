package main_thread

import (
	"github.com/hcraM41/crisp/comm/clog"
	"sync"
)

// 主队列大小
const mainQSize = 2048

// 主队列
var mainQ = make(chan func(), mainQSize)

// 开始标记
var started = false
var startLocker = &sync.Mutex{}

// Process 处理任务,
// 只将任务添加到队列而不是马上执行...
func Process(task func()) {
	if nil == task {
		return
	}

	mainQ <- task

	if !started {
		startLocker.Lock()
		defer startLocker.Unlock()

		if !started {
			started = true
			go execute()
		}
	}
}

// 执行 task
func execute() {
	for {
		task := <-mainQ

		if nil == task {
			continue
		}

		func() {
			//
			// task 是放在 for 循环里执行的,
			// 我希望每次执行 task 的时候都去检查一下异常.
			// 但是直接用 defer 是不可以的!
			// defer 是函数退出之前才被执行的...
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
