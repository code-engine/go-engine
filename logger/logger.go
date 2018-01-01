package logger

import (
	"log"
	"os"
)

type outputWriter interface {
	Print(...interface{})
}

var output outputWriter

func init() {
	SetDefaultOutput()
}

func write(message interface{}) {
	output.Print(message)
}

func Warn(message interface{}) {
	write(message)
}

func Fatal(message interface{}) {
	write(message)
}

func Print(message interface{}) {
	write(message)
}

func SetOutput(o outputWriter) {
	output = o
}

func SetDefaultOutput() {
	output = log.New(os.Stderr, "", log.LstdFlags)
}
