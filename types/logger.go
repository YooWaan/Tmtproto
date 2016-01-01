package types

import (
	"io"
	"log"
)

type AppLogger struct {
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Err   *log.Logger
}

func NewAppLogger(trace io.Writer, info io.Writer, warn io.Writer, err io.Writer) (*AppLogger) {
	var format int = log.Ldate|log.Ltime|log.Lshortfile
	return &AppLogger{
		Trace: log.New(trace, "TRACE: ", format),
		Info: log.New(info, "INFO: ", format),
		Warn: log.New(warn, "WARN: ", format),
		Err: log.New(err, "ERROR: ", format),
	}
}

/* define in applicatoin
var LOG AppLogger
func InitLogger(trace io.Writer, std io.Writer, err io.Writer) {
	LOG = NewAppLogger(trace, std, std, err)
}
*/
