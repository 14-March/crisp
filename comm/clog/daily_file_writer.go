/*
Package clog
@Author：14March
@File：daily_file_writer.go
*/
package clog

import (
	"errors"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type dailyFileWriter struct {
	// 日志文件名称
	fileName string
	// 上一次写入日期
	lastYearDay int
	// 输出文件
	outputFile *os.File
	// 文件交换锁
	fileSwitchLock *sync.Mutex
}

// Write 输出日志
func (w *dailyFileWriter) Write(byteArray []byte) (n int, err error) {
	if nil == byteArray ||
		len(byteArray) <= 0 {
		return 0, nil
	}

	outputFile, err := w.getOutputFile()

	if nil != err {
		return 0, err
	}

	_, _ = os.Stderr.Write(byteArray)
	_, _ = outputFile.Write(byteArray)

	return len(byteArray), nil
}

// 获取输出文件
// 每天创建一个新得日志文件
func (w *dailyFileWriter) getOutputFile() (io.Writer, error) {
	yearDay := time.Now().YearDay()

	if w.lastYearDay == yearDay &&
		nil != w.outputFile {
		// 如果当前日期和上一次日期一样
		// 且输出文件也不为空
		return w.outputFile, nil
	}

	if nil == w.fileSwitchLock {
		return nil, errors.New("fileSwitchLock is nil")
	}

	w.fileSwitchLock.Lock()
	defer w.fileSwitchLock.Unlock()

	// 加完锁之后进行二次判断
	if w.lastYearDay == yearDay &&
		nil != w.outputFile {
		return w.outputFile, nil
	}

	w.lastYearDay = yearDay

	// 构建日志目录
	err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm)

	if nil != err {
		return nil, errors.New("创建目录失败")
	}

	// 定义新的日志文件
	newDailyFile := w.fileName + "." + time.Now().Format("20060102")

	outputFile, err := os.OpenFile(
		newDailyFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if nil != err ||
		nil == outputFile {
		return nil, errors.New("打开文件失败")
	}

	if nil != w.outputFile {
		// 关闭原来的文件
		_ = w.outputFile.Close()
	}

	w.outputFile = outputFile
	return outputFile, nil
}
