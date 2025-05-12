package debugo

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

var debug = os.Getenv("DEBUGO")
var output io.Writer = os.Stderr
var noColors bool = false
var globalTimestamp *Timestamp = nil

type Debugger struct {
	namespace string
	color     *color.Color
	lastLog   time.Time
	forced    bool
	output    io.Writer
	channel   chan string
	timestamp *Timestamp
	options   *Options

	mutex *sync.Mutex
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
	Output io.Writer
	// Write log files in their own go routine (maintains order)
	Threaded bool
	// Enable leading timestamps by adding a time format
	Timestamp *Timestamp
}

// Returns an instance of Debugger configured using options
func NewWithOptions(namespace string, options *Options) *Debugger {
	logger := new(namespace, options)
	logger.applyOptions()
	return logger
}

// Returns an instance of Debugger configured with default values
func New(namespace string) *Debugger {
	logger := new(namespace, &Options{
		ForceEnable:         false,
		UseBackgroundColors: false,
		Color:               nil,
		Output:              nil,
		Threaded:            false,
		Timestamp:           nil,
	})
	logger.applyOptions()

	return logger
}

// Programmatically set the namespace(s) to debug during runtime
func SetDebug(namespace string) {
	debug = namespace
}

// Get the namespace(s) active for debug
func GetDebug() string {
	return debug
}

// Set the output writer
func SetOutput(w io.Writer) {
	output = w
}

// Get the output writer
func GetOutput() io.Writer {
	return output
}

// Disable all colors
func DisableColors(b bool) {
	noColors = b
}

// Set global timestamp (nil to disable)
func SetTimestamp(timestamp *Timestamp) {
	globalTimestamp = timestamp
}

// Check if instance would match the currently active debug namespace(s)
func (l *Debugger) Enabled() bool {
	return l.matchNamespace()
}

func new(namespace string, options *Options) *Debugger {
	return &Debugger{namespace: namespace, lastLog: time.Now(), forced: false, output: nil, options: options, mutex: &sync.Mutex{}}
}
