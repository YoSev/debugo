package debugo

import (
	"fmt"
	"strings"
)

func (l *Logger) debug(message ...any) {
	if l.matchNamespace() {
		stringMessages := make([]string, len(message))
		for i, v := range message {
			stringMessages[i] = fmt.Sprintf("%v", formatValue(v))
		}

		fmt.Fprintf(l.output, "%s %s %s\n", l.color.Sprint(l.namespace), noColor.Sprint(strings.Join(stringMessages, " ")), l.color.Sprintf("+%s", prettyPrintDuration(l.elapsed())))
	}
}
