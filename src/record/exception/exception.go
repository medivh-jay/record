package exception

import (
	"record/log"
)

func Errors(v interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("[INFO]")
			log.Info(err)
		}
	}()
	panic(v)
}

func Fatal(v interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("[FATAL]")
			log.Fatal(err)
		}
	}()
	panic(v)
}
