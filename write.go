package debugo

import (
	"fmt"
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
		msg := fmt.Sprint(message...)

		timestamp := ""
		if globalTimestamp != nil {
			timestamp = time.Now().Format(globalTimestamp.Format) + " "
		} else if l.timestamp != nil {
			timestamp = time.Now().Format(l.timestamp.Format) + " "
		}

		log := fmt.Sprintf("%s %s %s\n", l.color.Sprintf("%s%s", timestamp, l.namespace), noColor.Sprint(msg), l.color.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
		if l.channel != nil {
			l.channel <- log
		} else {
			if l.output == nil {
				fmt.Fprintf(output, "%s", log)
			} else {
				fmt.Fprintf(l.output, "%s", log)
			}
		}
	}
}
