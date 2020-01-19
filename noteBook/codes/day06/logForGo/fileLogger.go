package NewFileLoger

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
		Level:    logLevel,
		filePath: fp,
		fileName, fn,
		maxFileSize: maxSize,
	}
	err := f.initFileLogger()
	if err != nil {
		panic(err)
	}
	return f
}
func (l *FileLogger) initFileLogger() error {
	fullPath := path.join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file %s failed \n", fullPath)
		return err
	}
	l.fileObj = fileObj
	errFilePath := fullPath + ".err"
	errFileObj, err := os.OpenFile(errFilePath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file %s failed \n", errFilePath)
		return err
	}
	return nil
}
func (l *NewFileLoger) format(format, level string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("[%s] %s %s", t, level, format)

	return format
}

func (l *NewFileLoger) enable(level LogLevel) bool {
	return l.Level >= level
}

//查看文件大小是否超出阈值，进行切割文件
func (l *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	if fileInfo.Size() >= f.maxFileSize {
		return true
	}
	return false
}

func (l *NewFileLoger) Info(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(INFO) {
		return
	}
	funcName, fileName, line := getInfo(2)

	n, err = fmt.Fprintf(l.outPut, l.format(format, "info"), elem...)
	return
}

func (l *NewFileLoger) Debug(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(DEBUG) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "debug"), elem...)
	return
}

func (l *NewFileLoger) Warning(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(WARNING) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "warning"), elem...)
	return
}

func (l *NewFileLoger) Fatal(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(FATAL) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "fatal"), elem...)
	return
}

func (l *NewFileLoger) Error(format string, elem ...interface{}) (n int, err error) {
	if !l.enable(ERROR) {
		return
	}
	n, err = fmt.Fprintf(l.outPut, l.format(format, "error"), elem...)
	return
}

func (l *FileLogger) Close() {
	l.fileObj.Close()
	l.errObj.Close()
}
