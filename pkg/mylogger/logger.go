package logger

import (
	"fmt"
	"log"

)

type LoggerInterface interface {
	Infof()
	Warnf()
	Errorf()
	Debugf()
}

type Logger struct {
	prefix string
	LoggerInterface
}

func Default() *Logger {
	return &Logger{}
}

func (l *Logger) Infof(format string, values ...interface{}) {
	log.Printf(fmt.Sprintf("\033[36m[INFO] %s\033[0m", format), values...)
}

func (l *Logger) Warnf(format string, values ...interface{}) {
	log.Printf(fmt.Sprintf("\033[33m[WARN] %s\033[0m", format), values...)
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	log.Printf(fmt.Sprintf("\033[31m[ERROR] %s\033[0m", format), values...)
}


func (l *Logger) Debugf(format string, values ...interface{}) {
	log.Printf(fmt.Sprintf("[DEBUG]%s", format), values...)
}


