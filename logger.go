package streamlog

import (
	"io"
	"log"
	"os"

	"github.com/joefitzgerald/standardlog"
)

// OutputStream represents standard output streams
type OutputStream int

// The standard output streams
const (
	Out OutputStream = iota
	Err
)

// StreamLogger allows logging to stderr and stdout streams
type StreamLogger interface {
	Flags() int
	Prefix() string
	Print(s OutputStream, v ...interface{})
	Printf(s OutputStream, format string, v ...interface{})
	Println(s OutputStream, v ...interface{})
	PrintErr(v ...interface{})
	PrintfErr(format string, v ...interface{})
	PrintlnErr(v ...interface{})
	PrintOut(v ...interface{})
	PrintfOut(format string, v ...interface{})
	PrintlnOut(v ...interface{})
	SetFlags(flag int)
	SetPrefix(prefix string)
}

// NewWithWriters creates a new StdOutErrStreamLogger. The out and err variables set the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func NewWithWriters(out io.Writer, err io.Writer, prefix string, flag int) *StdOutErrStreamLogger {
	outLogger := log.New(os.Stdout, prefix, log.LstdFlags)
	errLogger := log.New(os.Stderr, prefix, log.LstdFlags)
	return &StdOutErrStreamLogger{out: outLogger, err: errLogger, prefix: prefix, flag: flag}
}

// StdOutErrStreamLogger is an OutErrLogger using stderr and stdout
type StdOutErrStreamLogger struct {
	out    standardlog.Logger // Logger for OutputStream.Out
	err    standardlog.Logger // Logger for OutputStream.Err
	prefix string             // prefix to write at beginning of each line
	flag   int                // properties
}

// New creates a new StdOutErrStreamLogger that writes to os.Stdout and
// os.Stderr.
func (l *StdOutErrStreamLogger) New() *StdOutErrStreamLogger {
	return NewWithWriters(os.Stdout, os.Stderr, "", log.LstdFlags)
}
