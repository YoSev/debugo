package debugo

import (
	"fmt"
)

func (l *Debugger) listen() {
	for {
		select {
		case log := <-l.channel:
			fmt.Fprintf(l.output, "%s", log)
		}
	}
}
