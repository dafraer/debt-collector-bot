package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string, keyVals ...interface{})
	Warn(msg string, keyVals ...interface{})
	Error(msg string, keyVals ...interface{})
	Debug(msg string, keyVals ...interface{})
	With(keyVals ...interface{}) Logger
}

const (
	Error Level = iota
	Warn
	Info
	Debug
)

func NewLogger(levelFilter Level, keyVals ...interface{}) Logger {
	return &logger{
		logger:      log.New(os.Stdout, "", 0),
		levelFilter: levelFilter,
		keyVals:     keyVals,
	}
}
