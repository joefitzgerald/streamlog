package streamlog

import (
	"io"
	"log"
	"os"
	"sync"

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

// New creates a new StreamLogger that writes to os.Stdout and os.Stderr.
func New() StreamLogger {
	return NewWithWriters(os.Stdout, os.Stderr, "", log.LstdFlags)
}

// NewWithWriters creates a new StdOutErrStreamLogger. The out and err variables set the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func NewWithWriters(out io.Writer, err io.Writer, prefix string, flag int) StreamLogger {
	outLogger := log.New(os.Stdout, prefix, log.LstdFlags)
	errLogger := log.New(os.Stderr, prefix, log.LstdFlags)
	return &outErrStreamLogger{out: outLogger, err: errLogger, prefix: prefix, flag: flag}
}

// NewWithLoggers creates a new StdOutErrStreamLogger. The out and err variables
// set the destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func NewWithLoggers(out standardlog.Logger, err standardlog.Logger, prefix string, flag int) StreamLogger {
	return &outErrStreamLogger{out: out, err: err, prefix: prefix, flag: flag}
}

// StdOutErrStreamLogger is an OutErrLogger using stderr and stdout
type outErrStreamLogger struct {
	mu     sync.Mutex         // ensures atomic writes; protects the following fields
	out    standardlog.Logger // Logger for OutputStream.Out
	err    standardlog.Logger // Logger for OutputStream.Err
	prefix string             // prefix to write at beginning of each line
	flag   int                // properties
}

// Flags returns the output flags for the logger.
func (l *outErrStreamLogger) Flags() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flag
}

// Prefix returns the output prefix for the logger.
func (l *outErrStreamLogger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

// SetPrefix sets the output prefix for the logger.
func (l *outErrStreamLogger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
	l.out.SetPrefix(prefix)
	l.err.SetPrefix(prefix)
}

// SetFlags sets the output flags for the logger.
func (l *outErrStreamLogger) SetFlags(flag int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = flag
	l.out.SetFlags(flag)
	l.err.SetFlags(flag)
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *outErrStreamLogger) Printf(s OutputStream, format string, v ...interface{}) {
	switch s {
	case Out:
		l.out.Printf(format, v)
	case Err:
		l.err.Printf(format, v)
	}
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *outErrStreamLogger) Print(s OutputStream, v ...interface{}) {
	switch s {
	case Out:
		l.out.Print(v)
	case Err:
		l.err.Print(v)
	}
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *outErrStreamLogger) Println(s OutputStream, v ...interface{}) {
	switch s {
	case Out:
		l.out.Println(v)
	case Err:
		l.err.Println(v)
	}
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *outErrStreamLogger) PrintfErr(format string, v ...interface{}) {
	l.err.Printf(format, v)
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *outErrStreamLogger) PrintErr(v ...interface{}) {
	l.err.Print(v)
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *outErrStreamLogger) PrintlnErr(v ...interface{}) {
	l.err.Println(v)
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *outErrStreamLogger) PrintfOut(format string, v ...interface{}) {
	l.out.Printf(format, v)
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *outErrStreamLogger) PrintOut(v ...interface{}) {
	l.out.Print(v)
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *outErrStreamLogger) PrintlnOut(v ...interface{}) {
	l.out.Println(v)
}
