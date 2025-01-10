package debugo

import (
	"fmt"
	"strings"
	"time"
)

func (l *Debugger) Debug(message ...any) {
	l.write(message...)
}

func (l *Debugger) Debugf(format string, message ...any) {
	l.write(fmt.Sprintf(format, message...))
}

func (l *Debugger) write(message ...any) {
	if l.matchNamespace() {
		stringMessages := make([]string, len(message))
		for i, v := range message {
			stringMessages[i] = fmt.Sprintf("%v", formatValue(v))
		}

		timestamp := ""
		if globalTimestamp != nil {
			timestamp = time.Now().Format(l.timestamp.Format) + " "
		} else if l.timestamp != nil {
			timestamp = time.Now().Format(l.timestamp.Format) + " "
		}

		log := fmt.Sprintf("%s %s %s\n", l.color.Sprintf("%s%s", timestamp, l.namespace), noColor.Sprint(strings.Join(stringMessages, " ")), l.color.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
		if l.channel != nil {
			l.channel <- log
		} else {
			fmt.Fprintf(l.output, "%s", log)
		}
	}
}
