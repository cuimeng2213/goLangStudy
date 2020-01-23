package logForGo

import (
	"fmt"
	"os"
	"path"
	"time"
)

const (
	ERROR_FILE = iota
	STD_FILE
)

type logMsg struct{
	Level LogLevel
	msg string
	funcName string
	fileName string
	line int
}

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errObj      *os.File
	maxFileSize int64
	//异步写日志
	//增加一个chan字段，所有日志写入该channel，然后开启多个goroutine去读取chanel 写入磁盘。
	logChan chan logMsg
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan : make(chan *logMsg, 10000)
	}
	err = f.initFileLogger()
	if err != nil {
		panic(err)
	}
	return f
}

// 初始化文件句柄
func (f *FileLogger) initFileLogger() error {
	fullPath := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file %s failed \n", fullPath)
		return err
	}
	f.fileObj = fileObj
	errFilePath := fullPath + ".err"
	errFileObj, err := os.OpenFile(errFilePath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file %s failed \n", errFilePath)
		return err
	}
	f.errObj = errFileObj
	return nil
}
func (f *FileLogger) format(format, level, funcName, fileName string, line int) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("[%s] [%s %s %d]%s %s", t, funcName, fileName, line, level, format)
	return format
}

func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

//查看文件大小是否超出阈值，进行切割文件
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	if fileInfo.Size() >= f.maxFileSize {
		return true
	}
	return false
}

func (f *FileLogger) changeFileObj(t int) {
	fullPath := path.Join(f.filePath, f.fileName)
	var err error
	if t == ERROR_FILE { // 切换错误日志文件句柄
		//1.关闭文件句柄
		f.errObj.Close()
		//2.重命名原来的文件 在原始文件后加时间戳
		oldErrFileName := fullPath + ".err"
		stuff := time.Now().Format("20060102150405")
		newErrorFileName := oldErrFileName + stuff
		os.Rename(oldErrFileName, newErrorFileName)
		//3.重新打开原始的文件
		//4.赋值给文件句柄
		f.errObj, err = os.OpenFile(oldErrFileName, os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Open file %s err ", oldErrFileName)
			return
		}

	} else {
		//1.关闭文件句柄
		f.fileObj.Close()
		//2.重命名原来的文件 在原始文件后加时间戳
		oldErrFileName := fullPath
		stuff := time.Now().Format("20060102150405")
		newErrorFileName := oldErrFileName + stuff
		os.Rename(oldErrFileName, newErrorFileName)

		f.fileObj, err = os.OpenFile(oldErrFileName, os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Open file %s err ", oldErrFileName)
			return
		}
	}
}

func (f *FileLogger) output(level LogLevel, format string, elem ...interface{}) {
	if !f.enable(level) {
		return
	}
	funcName, fileName, line := getInfo(3)
	//fmt.Println(funcName, fileName, line)
	format = f.format(format, getLevelName(level), funcName, fileName, line)

	// 错误日志大于ERROR写入错误日志文件
	if level >= ERROR {
		// 检测文件大小，如果超出阈值则切割文件
		if f.checkSize(f.errObj) {
			// 切换errObj文件句柄
			f.changeFileObj(ERROR_FILE)
		}
		fmt.Printf("AAAAAAAAAAAAAAA: error log \n")
		fmt.Fprintf(f.errObj, format, elem...)

	} else {
		// 切换fileObj文件句柄
		if f.checkSize(f.fileObj) {
			// 切换errObj文件句柄
			f.changeFileObj(STD_FILE)
		}
		fmt.Fprintf(f.fileObj, format, elem...)
	}
}

func (f *FileLogger) Info(format string, elem ...interface{}) {
	f.output(INFO, format, elem...)
}

func (f *FileLogger) Debug(format string, elem ...interface{}) {
	f.output(DEBUG, format, elem...)
}

func (f *FileLogger) Warning(format string, elem ...interface{}) {
	f.output(WARNING, format, elem...)
}

func (f *FileLogger) Fatal(format string, elem ...interface{}) {
	f.output(FATAL, format, elem...)
}

func (f *FileLogger) Error(format string, elem ...interface{}) {
	f.output(ERROR, format, elem...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errObj.Close()
}
