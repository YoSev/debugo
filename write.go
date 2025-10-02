package debugo

import (
	"fmt"
	"strings"
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

	if !d.matchNamespace() {
		return
	}

	msg := fmt.Sprint(message...)
	if msg == "" {
		return
	}

	// Optional timestamp
	var timestamp string
	if t := GetTimestamp(); t != nil {
		timestamp = time.Now().Format(t.Format)
	}

	// Build log parts
	parts := []string{}
	if timestamp != "" {
		parts = append(parts, timestamp)
	}

	if GetUseColors() {
		parts = append(parts, d.color.Sprintf("%s", d.namespace))
	} else {
		parts = append(parts, d.namespace)
	}

	parts = append(parts, msg)
	parts = append(parts, fmt.Sprintf("+%s", prettyPrintDuration(d.elapsed())))

	log := strings.Join(parts, " ") + "\n"

	// Write to output
	if d.output != nil {
		fmt.Fprint(d.output, log)
	} else if o := GetOutput(); o != nil {
		fmt.Fprint(o, log)
	}
}
