package debugo

import (
	"fmt"
)

func (l *Debugger) listen() {
	for {
		log := <-l.channel
		if l.output != nil {
			fmt.Fprintf(l.output, "%s", log)
		} else {
			outputMutex.Lock()
			fmt.Fprintf(output, "%s", log)
			outputMutex.Unlock()
		}
	}
}
