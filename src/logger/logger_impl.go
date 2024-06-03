package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
	"time"
)

type Level uint8
type values []interface{}

// String returns log level in a string format
func (l Level) String() string {
	switch l {
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

type logger struct {
	logger      *log.Logger
	levelFilter Level
	keyVals     []interface{}
}

// With returns a new logger which logs keyVals
func (l *logger) With(keyVals ...interface{}) Logger {
	newKeyVals := l.keyVals

	newKeyVals = append(newKeyVals, keyVals...)

	return &logger{
		levelFilter: l.levelFilter,
		logger:      l.logger,
		keyVals:     newKeyVals,
	}
}

// log outputs log to Stdout
func (l *logger) log(level Level, msg string, keyVals []interface{}) {
	if level > l.levelFilter {
		return
	}

	if len(keyVals)%2 == 1 {
		keyVals = append(keyVals, "VALUE_MISSING")
	}

	keyVals = append(keyVals, "message", msg)

	res := make(map[string]values, (len(keyVals)+1)/2)

	//Adding key-value pairs to the map
	for i := 0; i < len(keyVals); i += 2 {
		key := reflect.ValueOf(keyVals[i])

		if key.Kind() == reflect.Pointer {
			key = key.Elem()
		}

		if key.Kind() != reflect.String {
			l.Error("Key is not a string", "key", keyVals[i], "logger_log", true)
			continue
		}

		res[key.String()] = append(res[key.String()], keyVals[i+1])
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println("Unable to marshal json", err.Error())
	}

	l.logger.Printf("%s\t[%s]\t%s\n", time.Now().String(), level.String(), string(jsonData))
}

func (l *logger) Info(msg string, keyVals ...interface{}) {
	l.log(Info, msg, keyVals)
}

func (l *logger) Error(msg string, keyVals ...interface{}) {
	l.log(Error, msg, keyVals)
}

func (l *logger) Warn(msg string, keyVals ...interface{}) {
	l.log(Warn, msg, keyVals)
}

func (l *logger) Debug(msg string, keyVals ...interface{}) {
	l.log(Debug, msg, keyVals)
}

func (v values) MarshalJSON() ([]byte, error) {
	if len(v) < 1 {
		return []byte("null"), nil
	}

	if len(v) < 2 {
		return json.Marshal(v[0])
	}

	var buf []byte
	bufWriter := bytes.NewBuffer(buf)
	if err := json.NewEncoder(bufWriter).Encode([]interface{}(v)); err != nil {
		return nil, err
	}

	return bufWriter.Bytes(), nil
}
