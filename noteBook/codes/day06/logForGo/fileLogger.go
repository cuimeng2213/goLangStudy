package logForGo

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errObj      *os.File
	maxFileSize int64
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
	return f.Level >= level
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

func (f *FileLogger) output(level LogLevel, format string, elem ...interface{}) {
	if !f.enable(level) {
		return
	}
	funcName, fileName, line := getInfo(3)
	fmt.Println(funcName, fileName, line)
	format = f.format(format, getLevelName(level), funcName, fileName, line)
	// 错误日志大于ERROR写入错误日志文件
	if f.Level >= ERROR {
		fmt.Fprintf(f.errObj, format, elem...)
	} else {
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
