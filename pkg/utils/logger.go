package utils

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	LevelError = iota
	LevelWaring
	LevelInfo
	LevelDebug
)

type logger struct {
	level int
	mutex sync.Mutex
}

func (l *logger) print(prefix, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fmt.Println(runtime.FuncForPC(pc).Name(), file, line)
	}
	fmt.Printf("[%s]%s %s\n", prefix, time.Now().Format("2006-01-02 15:04:05"), msg)
}
func (l *logger) Error(msg string, v ...interface{}) {
	if l.level < LevelError {
		return
	}
	out := fmt.Sprintf(msg, v...)
	l.print("Error", out)
	panic(out)
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
