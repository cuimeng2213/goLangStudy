package logForGo

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
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

// 要改变结构体内的字段，需要使用指针接收者
func (l *LogForGo) SetOutput(output *os.File) {
	l.outPut = output
}

func (l *LogForGo) format(format, level string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("[%s] %s %s", t, level, format)
	fmt.Printf("Format: %#v \n", format)
	return format
}

func (l *LogForGo) enable(level LogLevel) bool {
	return l.Level >= level
}

func (l *LogForGo) Info(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(INFO) {
		return
	}
	funcName, fileName, line := getInfo(2)

	n, err = fmt.Fprintf(l.outPut, l.format(format, "info"), elem...)
	return
}

func (l *LogForGo) Debug(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(DEBUG) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "debug"), elem...)
	return
}

func (l *LogForGo) Warning(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(WARNING) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "warning"), elem...)
	return
}

func (l *LogForGo) Fatal(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(FATAL) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "fatal"), elem...)
	return
}

func (l *LogForGo) Error(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(ERROR) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "error"), elem...)
	return
}
