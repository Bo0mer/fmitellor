package logger

import (
	"io"
	"log"
)

// Logger can be used to log stuff.
type Logger struct {
	log *log.Logger
}

// NewLogger returns new Logger writing to the specified output.
func NewLogger(output io.Writer) *Logger {
	return &Logger{
		log: log.New(output, "", 0),
	}
}

// Printf writes to the logger's output. Arguments are handled in the manner of
// fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.log.Printf(format, v...)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf(format, v...)
}
