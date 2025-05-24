package debugo

import (
	"fmt"
	"time"
)

func (d *Debugger) Debug(message ...any) {
	d.write(message...)
}

func (d *Debugger) Debugf(format string, message ...any) {
	d.write(fmt.Sprintf(format, message...))
}

func (d *Debugger) write(message ...any) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.matchNamespace() {
		msg := fmt.Sprint(message...)

		t := GetTimestamp()
		var timestamp string
		if t != nil {
			timestamp = time.Now().Format(GetTimestamp().Format)
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
