package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel int32

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func level2String(level LogLevel) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	case 4:
		return "FATAL"
	default:
		return ""
	}
}

type LoggingMessage struct {
	Level    LogLevel
	FileName string
	FuncName string
	Message  string
	NowTime  string
	Line     int
}

type Logging struct {
	level       string
	logFileName string
	logFilePath string
	fileHandler *os.File
	logChan     chan *LoggingMessage
}

func getCallInfo() (fileName, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	fileName = path.Base(fileName)
	funcName = runtime.FuncForPC(pc).Name()
	return
}

// 日志结构体构指针造函数
func NewLogger(level LogLevel, logFilePath, logFileName string) *Logging {
	log := &Logging{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
		logChan:     make(chan *LoggingMessage, 50000),
		// TODO 错误日志文件
	}
	log.initLogger()
	log.writeFile()
	return log
}

// 异步写入日志文件
func (l *Logging) writeFile() {
	for msgData := range l.logChan {
		funcName, fileName, line := getCallInfo()
		//fmt.Printf("%s %s %d \n", msgData.FuncName, msgData.FileName, msgData.Line)
		// 由于使用Fprintln写入文件自带换行符 则去掉foramat中的\n（如果存在\n）
		format = strings.Trim(format, "\n")
		format = fmt.Sprintf("[%s] [%s:%s] [%d] [%s] %s",
			msgData.NowTime, msgData.FileName, msgData.FuncName, msgData.line, level2String(msgData.Level), format)
		msgStr := fmt.Sprintf(format, args...)
		fmt.Fprintln(l.fileHandler, msgStr)
	}
}

// 日志文件指针初始化
func (l *Logging) initLogger() {
	filePath := path.Join(l.logFilePath, l.logFileName)
	fp, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%失败%v", filePath, err))
	}
	l.fileHandler = fp
}

func (l *Logging) log(level LogLevel, format string, args ...interface{}) {
	// 如果当前日志句柄日志级别大于当前调用方法的级别则不记录
	if l.level > level {
		return
	}
	// 由于获取函数和代码行号比较耗时，商用版本应不调用该方法(也就是仅在debug模式打印)
	funcName, fileName, line := getCallInfo()
	// fmt.Printf("%s %s %d \n", funcName, fileName, line)
	// // 由于使用Fprintln写入文件自带换行符 则去掉foramat中的\n（如果存在\n）
	format = strings.Trim(format, "\n")
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := level2String(level)
	format = fmt.Sprintf("[%s] [%s:%s] [%d] [%s] %s", nowStr, fileName, funcName, line, levelStr, format)
	msgStr := fmt.Sprintf(format, args...)
	// fmt.Fprintln(l.fileHandler, msgStr)

	// 1.向日志通道写入日志信息结构体
	mshInfo := &LoggingMessage{
		Level:    levelStr,
		FileName: fileName,
		FuncName: funcName,
		Message:  msgStr,
		NowTime:  nowStr,
		Line:     line,
	}
	select {
	case l.logChan <- msgInfo:
	default:

	}
}

// Debug method
func (l *Logging) Debug(format string, args ...interface{}) {
	l.log(DebugLevel, format, args...)
}

// Info method
func (l *Logging) Info(format string, args ...interface{}) {
	l.log(InfoLevel, format, args...)
}

// Warn method
func (l *Logging) Warn(format string, args ...interface{}) {
	l.log(WarnLevel, format, args...)
}

// Error method
func (l *Logging) Error(format string, args ...interface{}) {
	l.log(ErrorLevel, format, args...)
}

// Fatal method
func (l *Logging) Fatal(format string, args ...interface{}) {
	l.log(FatalLevel, format, args...)
}
