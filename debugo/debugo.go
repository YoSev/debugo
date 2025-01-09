package debugo

import (
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	namespace string
	color     *color.Color
	lastLog   time.Time
	forced    bool
	output    *os.File
	channel   chan string
	timestamp *Timestamp
}

type Timestamp struct {
	Format string
}

// Options overwrites debugs default values
type Options struct {
	// Force log output independent of given namespace matching (default: false)
	ForceEnable bool
	// Use background colors over foreground colors (default: false)
	UseBackgroundColors bool
	// Use a static color (github.com/fatih/color) (default: random foreground color)
	Color *color.Color
	// Defines the pipe to output to, eg. stdOut (default: stdErr)
	Output *os.File
	// Write log files in their own go routine (maintains order)
	Threaded bool
	// Enable leading timestamps by adding a time format
	Timestamp *Timestamp
}

// Returns a log-function and an instance of Logger configured using options
func NewWithOptions(namespace string, options *Options) (func(message ...any), *Logger) {
	logger := new(namespace)
	logger.applyOptions(options)
	return logFunc(logger)
}

// Returns a log-function and an instance of Logger configured with default values
func New(namespace string) (func(message ...any), *Logger) {
	logger := new(namespace)
	logger.applyOptions(&Options{
		ForceEnable:         false,
		UseBackgroundColors: false,
		Color:               nil,
		Output:              os.Stderr,
		Threaded:            false,
		Timestamp:           nil,
	})

	return logFunc(logger)
}

func new(namespace string) *Logger {
	return &Logger{namespace: namespace, lastLog: time.Now(), forced: false, output: os.Stderr}
}

var logFunc = func(logger *Logger) (func(message ...any), *Logger) {
	return func(message ...any) {
		logger.write(message...)
	}, logger
}
