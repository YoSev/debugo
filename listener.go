package debugo

import (
	"fmt"
)

func (l *Debugger) listen() {
	for {
		select {
		case log := <-l.channel:
			if l.output != nil {
				fmt.Fprintf(l.output, "%s", log)
			} else {
				fmt.Fprintf(output, "%s", log)
			}
		}
	}
}
