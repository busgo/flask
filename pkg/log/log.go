package logback

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type LogLevel int32

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

const logFormatter = "timestamp=\"%s\",level=\"%s\",msg=\"%s\",source=\"%s\""

type Logback interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}
type logback struct {
	path      string
	level     LogLevel
	service   string
	debugFile *os.File
	infoFile  *os.File
	warnFile  *os.File
	errorFile *os.File
}

// new logback
func NewLogBack(path string, level LogLevel) (l Logback, err error) {
	logback := &logback{path: path, level: level}
	err = logback.init()
	return logback, err
}

// init log back
func (l *logback) init() error {

	exists, err := pathExists(l.path)
	if err != nil {
		return err
	}
	if !exists {

		err = os.MkdirAll(l.path, 644)
		if err != nil {
			return err
		}
	}

	levels := []LogLevel{DEBUG, INFO, WARN, ERROR}

	for _, logLevel := range levels {

		if l.matchLevel(logLevel) {

			levelDesc := getLevelDesc(logLevel)
			fileName := fmt.Sprintf("%s/%s.log", l.path, levelDesc)
			logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_CREATE|os.O_APPEND, 644)
			if err != nil {
				return err
			}
			switch logLevel {
			case DEBUG:
				l.debugFile = logFile
			case INFO:
				l.infoFile = logFile
			case WARN:
				l.warnFile = logFile
			case ERROR:
				l.errorFile = logFile
			}

		}

	}
	return nil

}
func pathExists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}
	return true, nil

}

func (l *logback) Debug(format string, args ...interface{}) {

	if !l.matchLevel(INFO) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	source := getSource()
	content := fmt.Sprintf(logFormatter, timestamp, getLevelDesc(DEBUG), msg, source)

	fmt.Fprintln(l.debugFile, content)

}

func (l *logback) Info(format string, args ...interface{}) {

	if !l.matchLevel(INFO) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	source := getSource()
	content := fmt.Sprintf(logFormatter, timestamp, getLevelDesc(INFO), msg, source)
	fmt.Println(content)
	fmt.Fprintln(l.infoFile, content)

}

func (l *logback) Warn(format string, args ...interface{}) {

	if !l.matchLevel(WARN) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	source := getSource()
	content := fmt.Sprintf(logFormatter, timestamp, getLevelDesc(WARN), msg, source)
	fmt.Println(content)
	fmt.Fprintln(l.warnFile, content)

}
func (l *logback) Error(format string, args ...interface{}) {

	if !l.matchLevel(ERROR) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	source := getSource()
	content := fmt.Sprintf(logFormatter, timestamp, getLevelDesc(ERROR), msg, source)
	fmt.Println(content)
	fmt.Fprintln(l.errorFile, content)

}

// check macth level
func (l logback) matchLevel(level LogLevel) bool {

	return l.level <= level
}

// get source
func getSource() string {

	ptr, file, line, _ := runtime.Caller(2)
	fun := runtime.FuncForPC(ptr)
	funcName := fun.Name()

	return fmt.Sprintf("%s:%d:%s()", file, line, funcName)
}

// get level desc
func getLevelDesc(level LogLevel) string {

	switch level {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "uknow"
	}

}
