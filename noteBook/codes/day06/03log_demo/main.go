package main

import (
	"fmt"
	"goLangStudy/noteBook/codes/day06/logForGo"
	"os"
	"time"
)

func main() {
	logger := logForGo.NewLogForGo("warning")
	fileObj, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file xx.log failed")
		return
	}
	logger.SetOutput(fileObj)
	for {
		logger.Info("=%s=\n", "这是一条测试语句")
		logger.Debug("=%s=\n", "这是一条测试语句")
		logger.Warning("=%s=\n", "这是一条测试语句")
		logger.Error("=%s=\n", "这是一条测试语句")
		logger.Fatal("=%s=\n", "这是一条测试语句")
		time.Sleep(3 * time.Second)
	}

	defer fileObj.Close()
}
