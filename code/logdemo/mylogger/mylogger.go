package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel uint16

// 错误级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TARCE
	INFO
	WARMING
	ERROR
	FATAL
)

// 日志结构体
type Logger struct {
	Level LogLevel
}

// 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

// 解析日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s) // 转成小写
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TARCE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARMING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别") //返回error类型
		// err := fmt.Errorf("无效的日志级别")
		return UNKNOWN, err
	}
}

// 日志开关,大于了.level的才写入
func (l Logger) enable(lev LogLevel) bool {
	return lev >= l.Level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [Debug] %s\n", now.Format("2006-01-0-2 15:04:05"), msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-0-2 15:04:05"), msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARMING) {
		now := time.Now()
		fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-0-2 15:04:05"), msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-0-2 15:04:05"), msg)
	}
}
