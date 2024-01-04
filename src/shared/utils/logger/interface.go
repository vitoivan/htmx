package logger

type LoggerInterface interface {
	Log(text string)
	Info(text string)
	Error(text string)
	Warn(text string)
	Debug(text string)
	Fatal(text string)
}
