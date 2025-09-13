package debugo

import (
	"fmt"
	"time"
)

func (d *Debugger) Debug(message ...any) *Debugger {
	d.write(message...)
	return d
}

func (d *Debugger) Debugf(format string, message ...any) *Debugger {
	d.write(fmt.Sprintf(format, message...))
	return d
}

func (d *Debugger) write(message ...any) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.matchNamespace() {
		msg := fmt.Sprint(message...)

		if msg == "" {
			return
		}

		t := GetTimestamp()
		var timestamp string
		if t != nil {
			timestamp = time.Now().Format(t.Format)
		}

		log := fmt.Sprintf("%s %s %s %s\n", timestamp, d.color.Sprintf("%s", d.namespace), msg, d.color.Sprintf("+%s", prettyPrintDuration(d.elapsed())))

		if d.output != nil {
			fmt.Fprintf(d.output, "%s", log)
		} else {
			o := GetOutput()
			if o != nil {
				fmt.Fprintf(o, "%s", log)
			}
		}
	}
}
