package debugo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	namespace string
	color     string
	lastLog   time.Time
	forced    bool
}

type Options struct {
	ForceEnable bool
}

func NewWithOptions(namespace string, options *Options) (func(message ...any), *Logger) {
	f, logger := New(namespace)
	logger.configure(options)

	return f, logger
}

func New(namespace string) (func(message ...any), *Logger) {
	logger := &Logger{namespace: namespace, lastLog: time.Now(), forced: false}
	logger.color = logger.getNextColor()

	return func(message ...any) {
		logger.debug(message...)
	}, logger
}

func (l *Logger) debug(message ...any) {
	if l.matchNamespace() {
		stringMessages := make([]string, len(message))
		for i, v := range message {
			stringMessages[i] = fmt.Sprintf("%v", formatValue(v))
		}

		fmt.Fprintf(os.Stderr, "%s %s %s\n", l.applyColor(l.namespace), resetColor(strings.Join(stringMessages, " ")), l.applyColor(fmt.Sprintf("+%s", prettyPrintDuration(l.elapsed()))))
	}
}

func (l *Logger) configure(options *Options) {
	if options != nil {
		if options.ForceEnable {
			l.forced = true
		}
	}
}
