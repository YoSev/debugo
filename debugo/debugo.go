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
}

// Options overwrites debugs default values
type Options struct {
	// Forces log output independent of given namespace matching (default: false)
	ForceEnable bool
	// Defines to use background colors over foreground colors (default: false)
	UseBackgroundColors bool
	// Defines a strict color (github.com/fatih/color) (default: random foreground color)
	Color *color.Color
	// Defines the pipe to output to, eg. stdOut (default: stdErr)
	Output *os.File
}

// Returns a log-function and an instance of Logger configured using options
func NewWithOptions(namespace string, options *Options) (func(message ...any), *Logger) {
	logger := new(namespace)
	logger.applyOptions(options)

	return func(message ...any) {
		logger.write(message...)
	}, logger
}

// Returns a log-function and an instance of Logger configured with default values
func New(namespace string) (func(message ...any), *Logger) {
	logger := new(namespace)
	logger.applyOptions(&Options{
		ForceEnable:         false,
		UseBackgroundColors: false,
		Color:               nil,
	})

	return func(message ...any) {
		logger.write(message...)
	}, logger
}

func new(namespace string) *Logger {
	return &Logger{namespace: namespace, lastLog: time.Now(), forced: false, output: os.Stderr}
}
