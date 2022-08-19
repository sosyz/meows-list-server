package utils

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	LevelPanic = iota
	LevelError
	LevelWaring
	LevelInfo
	LevelDebug
)

type logger struct {
	level   int
	mutex   sync.Mutex
	errFile *os.File
	out     io.Writer
}

func (l *logger) writeErrLog(msg string) {
	if l.errFile == nil {
		//	打开文件 不存在则创建
		file, err := os.OpenFile(fmt.Sprintf("%s/%s.error.log", "./", time.Now().Format("2006_01_02_15_04_05")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("open err.log failed, err:", err)
			return
		}
		l.errFile = file
	}
	_, _ = l.errFile.WriteString(msg)
}

func (l *logger) print(prefix, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	pc, file, line, ok := runtime.Caller(2)
	var stack string
	if ok {
		stack = fmt.Sprintf("%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
	}
	// 取出文件名
	fileName := strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
	out := fmt.Sprintf("%s [%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), prefix, fileName, msg)
	_, _ = l.out.Write([]byte(out))
	if prefix == "Error" || prefix == "Panic" {
		l.writeErrLog(stack + " " + out)
	}
}

func (l *logger) Panic(msg string, v ...interface{}) {
	if l.level < LevelPanic {
		return
	}
	out := fmt.Sprintf(msg, v...)
	l.print("Panic", out)
	panic(out)
}

func (l *logger) Error(msg string, v ...interface{}) {
	if l.level < LevelError {
		return
	}
	l.print("Error", fmt.Sprintf(msg, v...))
}

func (l *logger) Waring(msg string, v ...interface{}) {
	if l.level < LevelWaring {
		return
	}
	l.print("Waring", fmt.Sprintf(msg, v...))
}

func (l *logger) Info(msg string, v ...interface{}) {
	if l.level < LevelInfo {
		return
	}
	l.print("Info", fmt.Sprintf(msg, v...))
}

func (l *logger) Debug(msg string, v ...interface{}) {
	if l.level < LevelDebug {
		return
	}
	l.print("Debug", fmt.Sprintf(msg, v...))
}

func (l *logger) Write() io.Writer {
	return l.out
}
