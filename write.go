package debugo

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func (l *Debugger) Debug(message ...any) {
	l.write(message...)
}

func (l *Debugger) Debugf(format string, message ...any) {
	l.write(fmt.Sprintf(format, message...))
}

func (l *Debugger) write(message ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.matchNamespace() {
		msg := fmt.Sprint(message...)

		timestamp := ""
		if globalTimestamp != nil {
			timestamp = time.Now().Format(globalTimestamp.Format) + " "
		} else if l.timestamp != nil {
			timestamp = time.Now().Format(l.timestamp.Format) + " "
		}

		var log string
		if !noColors {
			log = fmt.Sprintf("%s %s %s\n", l.color.Sprintf("%s%s", timestamp, l.namespace), noColor.Sprint(msg), l.color.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
		} else {
			log = fmt.Sprintf("%s %s %s\n", fmt.Sprintf("%s%s", timestamp, l.namespace), fmt.Sprint(msg), fmt.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
		}
		if l.channel != nil {
			l.channel <- log
		} else {
			mutex.Lock()
			defer mutex.Unlock()
			if l.output == nil {
				fmt.Fprintf(output, "%s", log)
			} else {
				fmt.Fprintf(l.output, "%s", log)
			}
		}
	}
}
