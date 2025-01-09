package debugo

import (
	"fmt"
	"strings"
)

func (l *Logger) write(message ...any) {
	if l.matchNamespace() {
		stringMessages := make([]string, len(message))
		for i, v := range message {
			stringMessages[i] = fmt.Sprintf("%v", formatValue(v))
		}

		log := fmt.Sprintf("%s %s %s\n", l.color.Sprint(l.namespace), noColor.Sprint(strings.Join(stringMessages, " ")), l.color.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
		if l.channel != nil {
			l.channel <- log
		} else {
			fmt.Fprintf(l.output, "%s", log)
		}
	}
}
