package logger

import (
	"log"
	"os"
)

type logger struct {
	name string
}

func NewLogger(name string) LoggerInterface {
	return &logger{
		name: "[" + name + "]",
	}
}

func (l *logger) Log(text string) {
	log.Printf("%s: %s", yellow(l.name), green(text))
}

func (l *logger) Info(text string) {
	log.Printf("%s: %s", yellow(l.name), cyan(text))
}

func (l *logger) Error(text string) {
	log.Printf("%s: %s", yellow(l.name), red(text))
}

func (l *logger) Warn(text string) {
	log.Printf("%s: %s", yellow(l.name), yellow(text))
}

func (l *logger) Debug(text string) {
	log.Printf("%s: %s", yellow(l.name), magenta(text))
}

func (l *logger) Fatal(text string) {
	log.Printf("%s: %s", yellow(l.name), red(text))
	os.Exit(1)
}

var Default LoggerInterface = NewLogger("default")
