package log

import (
	"log"
	"os"
	"strconv"
	"time"
)

var logs *log.Logger

var Prefix = "[Debug]"

func init() {
	if logs == nil {
		now := time.Now()
		logName := strconv.Itoa(now.Year()) + "-" + now.Month().String() + "-" + strconv.Itoa(now.Day())
		fileInfo, err := os.Stat("log" + logName + ".log")
		if err == nil {
			file, _ := os.Open(fileInfo.Name())
			logs = log.New(file, Prefix, log.Ltime)
		} else {
			file, _ := os.Create("log-" + logName + ".log")
			logs = log.New(file, Prefix, log.Ltime)
		}
	}
}

func SetPrefix(prefix string) {
	Prefix = prefix
	return
}

func Fatal(v interface{}) {
	logs.Fatalln(v)
}

func Info(v interface{}) {
	logs.Println(v)
}
