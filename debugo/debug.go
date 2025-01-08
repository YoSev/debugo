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
}

func New(namespace string) *Logger {
	logger := &Logger{namespace: namespace}
	logger.color = logger.getNextColor()

	return logger
}

func (l *Logger) Debug(message ...any) {
	if l.matchNamespace() {

		stringMessages := make([]string, len(message))
		for i, v := range message {
			stringMessages[i] = fmt.Sprintf("%v", formatValue(v))
		}

		if len(l.namespace) > 0 {
			fmt.Fprintf(os.Stderr, "%s ", l.applyColor(l.namespace))
		}

		fmt.Fprintf(os.Stderr, "%+v ", resetColor(strings.Join(stringMessages, " ")))
		fmt.Fprintf(os.Stderr, "%s\n", l.applyColor(fmt.Sprintf("+%s", prettyPrintDuration(l.elapsed()))))
	}
}
