package logger

//go:generate mockery --with-expecter --name Logger
type Logger interface {
	Fatalf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	WithError(err error) Logger
	WithData(data map[string]interface{}) Logger
}
