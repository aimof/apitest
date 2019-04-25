package logger

import (
	"log"
	"os"
)

const (
	ALL = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var LEVEL int = INFO

func Debug(str string) {
	if LEVEL <= DEBUG {
		log.New(os.Stderr, "debug: ", log.Ldate|log.Ltime).Output(2, str)
	}
}

func Info(str string) {
	if LEVEL <= INFO {
		log.New(os.Stderr, "info: ", log.Ldate|log.Ltime).Output(2, str)
	}
}

func Error(str string) {
	if LEVEL <= ERROR {
		log.New(os.Stderr, "error: ", log.Ldate|log.Ltime).Output(2, str)
	}
}

func Fatal(str string) {
	if LEVEL <= FATAL {
		log.New(os.Stderr, "fatal: ", log.Ldate|log.Ltime).Output(2, str)
	}
}
