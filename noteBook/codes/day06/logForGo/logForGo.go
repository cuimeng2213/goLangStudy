package logForGo

import (
	"errors"
	"path"
	"runtime"
	"strings"
)

// 自定义日志级别
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

//获取行号
func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	lineNo = line
	return
}

// 获取日志级别字符串
func getLevelName(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	}
	return "UNKNOWN"
}
