package main

import (
	"cuimeng/homework/myLogger/logging"
	"fmt"
)

func main() {
	fmt.Println(logging.InfoLevel)
	logger := logging.NewLogger(logging.ErrorLevel, "./", "log.txt")
	s := "白日依山尽"
	//s2 := fmt.Sprintf("this is a str %s\n", s)
	for {

		logger.Debug("这是一个测试日志%s", s)
		logger.Info("这是一个测试日志%s", s)
		logger.Warn("这是一个测试日志%s", s)
		logger.Error("这是一个测试日志%s", s)

	}
}
