package logForGo

import (
	// "errors"
	"fmt"
	"os"
	"path"
	// "runtime"
	// "strings"
	"time"
)

type LogForGo struct {
	outPut *os.File
	Level  LogLevel
}

func NewLogForGo(level string) LogForGo {
	lev, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return LogForGo{
		outPut: os.Stdout,
		Level:  lev,
	}
}

// 要改变结构体内的字段，需要使用指针接收者
func (l *LogForGo) SetOutput(output *os.File) {
	l.outPut = output
}

func (l *LogForGo) format(format, level, funcName, fileName string, line int) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	fileName = path.Base(fileName)
	format = fmt.Sprintf("[%s][%s %s %d] %s %s", t, funcName, fileName, line, level, format)
	fmt.Printf("Format: %#v \n", format)
	return format
}

func (l *LogForGo) enable(level LogLevel) bool {
	return l.Level >= level
}

func (l *LogForGo) output(level LogLevel, format string, elem ...interface{}) {
	if !l.enable(level) {
		return
	}
	funcName, fileName, line := getInfo(3)
	fmt.Fprintf(l.outPut, l.format(format, getLevelName(level), funcName, fileName, line), elem...)
}

func (l *LogForGo) Info(format string, elem ...interface{}) {
	l.output(INFO, format, elem...)
}

func (l *LogForGo) Debug(format string, elem ...interface{}) {
	l.output(DEBUG, format, elem...)
}

func (l *LogForGo) Warning(format string, elem ...interface{}) {
	l.output(WARNING, format, elem...)
}

func (l *LogForGo) Fatal(format string, elem ...interface{}) {
	l.output(FATAL, format, elem...)
}

func (l *LogForGo) Error(format string, elem ...interface{}) {
	l.output(ERROR, format, elem...)
}
