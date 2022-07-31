package logger

type ILoggerProvider interface {
	Error(namespace, message string)
	Info(namespace, message string)
}
