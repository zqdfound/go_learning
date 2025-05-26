package main

import (
	"fmt"
)

// 文件日志，扩展＋组合
type FileLog struct {
	ConsoleLog //扩展
	FilePath   string
}

func (f FileLog) Log(msg string) {
	//复用ConsoleLog
	f.ConsoleLog.Log(msg)
	//扩展
	fmt.Println("[File] restore log %s to path %s ", msg, f.FilePath)
}

// 使用接口的函数
func Process(logger Log) {
	logger.Log("log started...")
}
