package logger

type LoggerFields map[string]interface{}
type Logger interface {
	Info(v string)
	InfoWithFields(v string, fields LoggerFields)
	InfoError(v string, err error)
}
